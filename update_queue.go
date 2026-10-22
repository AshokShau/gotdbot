package gotdbot

import "sync"

// updateQueue is an in-memory FIFO that grows without blocking the receiver.
type updateQueue struct {
	mu      sync.Mutex
	memory  chan TlObject
	backlog []TlObject
	ready   chan struct{}
	closed  bool
}

func newUpdateQueue(size int) *updateQueue {
	return &updateQueue{memory: make(chan TlObject, size), ready: make(chan struct{}, 1)}
}

func (q *updateQueue) notify() {
	select {
	case q.ready <- struct{}{}:
	default:
	}
}

func (q *updateQueue) push(update TlObject) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return ErrClientClosed
	}
	if len(q.backlog) == 0 {
		select {
		case q.memory <- update:
			q.notify()
			return nil
		default:
		}
	}
	q.backlog = append(q.backlog, update)
	q.notify()
	return nil
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
		if len(q.memory) > 0 || len(q.backlog) > 0 {
			q.notify()
		}
	}()
	select {
	case update := <-q.memory:
		return update, nil
	default:
	}
	if len(q.backlog) == 0 {
		return nil, nil
	}
	update := q.backlog[0]
	q.backlog[0] = nil
	q.backlog = q.backlog[1:]
	if len(q.backlog) == 0 {
		q.backlog = nil
	}
	return update, nil
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
	q.backlog = nil
	return nil
}
