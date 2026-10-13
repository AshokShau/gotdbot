package gotdbot

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// updateQueue keeps a bounded number of decoded updates in memory.
type updateQueue struct {
	mu     sync.Mutex
	memory chan TlObject
	ready  chan struct{}
	closed bool

	dir                     string
	reader, writer          *os.File
	head, tail              uint64
	readOffset, writeOffset int64
	spilled                 int
}

const updateSegmentSize = 8 << 20

func newUpdateQueue(size int) *updateQueue {
	return &updateQueue{memory: make(chan TlObject, size), ready: make(chan struct{}, 1)}
}

func (q *updateQueue) notify() {
	select {
	case q.ready <- struct{}{}:
	default:
	}
}

func (q *updateQueue) segment(id uint64) string {
	return filepath.Join(q.dir, strconv.FormatUint(id, 10))
}

func (q *updateQueue) push(update TlObject) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return ErrClientClosed
	}

	// Once spilling starts, append every new update to disk until it drains.
	// Otherwise newer updates could overtake the older disk backlog.
	if q.dir == "" {
		select {
		case q.memory <- update:
			q.notify()
			return nil
		default:
		}
	}
	data, err := json.Marshal(update)
	if err != nil {
		return err
	}
	if uint64(len(data)) > uint64(^uint32(0)) {
		return fmt.Errorf("update exceeds spool record size")
	}

	if q.dir == "" {
		q.dir, err = os.MkdirTemp("", "gotdbot-updates-*")
		if err != nil {
			return err
		}
		q.head, q.tail = 1, 1
	}

	if q.writer == nil || q.writeOffset >= updateSegmentSize {
		id := q.tail
		if q.writer != nil {
			id++
		}
		writer, openErr := os.OpenFile(q.segment(id), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if openErr != nil {
			if q.spilled == 0 {
				openErr = errors.Join(openErr, q.clearSpool())
			}
			return openErr
		}
		if q.writer != nil {
			err = q.writer.Close()
		}
		q.writer, q.tail, q.writeOffset = writer, id, 0
	}
	record := binary.BigEndian.AppendUint32(nil, uint32(len(data)))
	record = append(record, data...)
	n, writeErr := q.writer.WriteAt(record, q.writeOffset)
	if writeErr == nil && n != len(record) {
		writeErr = io.ErrShortWrite
	}
	if writeErr != nil {
		writeErr = errors.Join(err, writeErr, q.writer.Truncate(q.writeOffset))
		if q.spilled == 0 {
			writeErr = errors.Join(writeErr, q.clearSpool())
		}
		return writeErr
	}
	q.writeOffset += int64(n)
	q.spilled++
	q.notify()
	return err
}

func (q *updateQueue) next(stop <-chan struct{}) (TlObject, error) {
	for {
		select {
		case <-stop:
			return nil, ErrClientClosed
		default:
		}
		update, err := q.take()
		if update != nil || err != nil {
			return update, err
		}
		select {
		case <-q.ready:
		case <-stop:
			return nil, ErrClientClosed
		}
	}
}

func (q *updateQueue) take() (TlObject, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return nil, ErrClientClosed
	}
	defer func() {
		if len(q.memory) > 0 || q.spilled > 0 {
			q.notify()
		}
	}()
	select {
	case update := <-q.memory:
		return update, nil
	default:
	}
	if q.spilled == 0 {
		return nil, nil
	}
	for {
		if q.reader == nil {
			var err error
			q.reader, err = os.Open(q.segment(q.head))
			if err != nil {
				return nil, err
			}
		}
		info, err := q.reader.Stat()
		if err != nil {
			return nil, err
		}
		if q.readOffset == info.Size() && q.head < q.tail {
			if err := q.retireSegment(); err != nil {
				return nil, err
			}
			continue
		}
		var header [4]byte
		if _, err := q.reader.ReadAt(header[:], q.readOffset); err != nil {
			return nil, err
		}
		size := int64(binary.BigEndian.Uint32(header[:]))
		if size == 0 || size > info.Size()-q.readOffset-4 {
			return nil, fmt.Errorf("invalid update spool record length %d", size)
		}
		data := make([]byte, int(size))
		if _, err := q.reader.ReadAt(data, q.readOffset+4); err != nil {
			return nil, err
		}
		update, _, err := Unmarshal(data)
		if err != nil {
			return nil, err
		}
		q.readOffset += 4 + size
		q.spilled--
		if q.spilled == 0 {
			err = q.clearSpool()
		} else if q.readOffset == info.Size() && q.head < q.tail {
			err = q.retireSegment()
		}
		return update, err
	}
}

func (q *updateQueue) retireSegment() error {
	err := q.reader.Close()
	err = errors.Join(err, os.Remove(q.segment(q.head)))
	q.reader = nil
	q.head++
	q.readOffset = 0
	return err
}

func (q *updateQueue) clearSpool() error {
	var err error
	for _, file := range []*os.File{q.reader, q.writer} {
		if file != nil {
			err = errors.Join(err, file.Close())
		}
	}
	if q.dir != "" {
		err = errors.Join(err, os.RemoveAll(q.dir))
	}
	q.dir, q.reader, q.writer = "", nil, nil
	q.head, q.tail, q.readOffset, q.writeOffset, q.spilled = 0, 0, 0, 0, 0
	return err
}

func (q *updateQueue) close() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return nil
	}
	q.closed = true
	close(q.ready)
	for len(q.memory) > 0 {
		<-q.memory
	}
	return q.clearSpool()
}
