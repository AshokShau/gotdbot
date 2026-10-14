package gotdbot

import (
	"encoding/binary"
	"errors"
	"os"
	"strings"
	"sync"
	"testing"
)

func newTestQueue(t *testing.T, capacity int) *updateQueue {
	t.Helper()
	q := newUpdateQueue(capacity)
	t.Cleanup(func() {
		if err := q.close(); err != nil {
			t.Errorf("queue cleanup: %v", err)
		}
	})
	return q
}

func queueTestUpdate(id int64) TlObject {
	return &UpdateNewMessage{Message: &Message{Id: id, ChatId: 42}}
}

func TestUpdateQueueFIFOAcrossMemoryAndDisk(t *testing.T) {
	q := newTestQueue(t, 2)
	for i := int64(1); i <= 5; i++ {
		if err := q.push(queueTestUpdate(i)); err != nil {
			t.Fatal(err)
		}
	}
	dir := q.dir
	if dir == "" || len(q.memory) != 2 || q.spilled != 3 {
		t.Fatal("queue did not enforce memory bound and spill")
	}
	// Make space in memory, then append. New updates must stay behind the spool.
	first, err := q.take()
	if err != nil || first.(*UpdateNewMessage).Message.Id != 1 {
		t.Fatalf("first = (%v, %v)", first, err)
	}
	if err := q.push(queueTestUpdate(6)); err != nil {
		t.Fatal(err)
	}
	for want := int64(2); want <= 6; want++ {
		obj, err := q.next(nil)
		if err != nil {
			t.Fatal(err)
		}
		if got := obj.(*UpdateNewMessage).Message.Id; got != want {
			t.Fatalf("ID = %d, want %d", got, want)
		}
	}
	if _, err := os.Stat(dir); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("spool not removed: %v", err)
	}
	if q.dir != "" || q.spilled != 0 {
		t.Fatal("spool not reset after drain")
	}
	if obj, err := q.take(); obj != nil || err != nil {
		t.Fatalf("empty = (%v, %v)", obj, err)
	}
	if err := q.push(queueTestUpdate(7)); err != nil || len(q.memory) != 1 {
		t.Fatalf("did not return to memory: %v", err)
	}
}

func TestUpdateQueueSegmentRotation(t *testing.T) {
	q := newTestQueue(t, 0)
	large := &UpdateNewMessage{Message: testMessage(42, strings.Repeat("x", updateSegmentSize))}
	large.Message.Id = 1
	if err := q.push(large); err != nil {
		t.Fatal(err)
	}
	if err := q.push(queueTestUpdate(2)); err != nil {
		t.Fatal(err)
	}
	if q.tail != 2 {
		t.Fatalf("tail = %d, want two segments", q.tail)
	}
	firstSegment := q.segment(1)
	obj, err := q.take()
	if err != nil || len(obj.(*UpdateNewMessage).Message.GetText()) != updateSegmentSize {
		t.Fatalf("large record = (%T, %v)", obj, err)
	}
	if _, err := os.Stat(firstSegment); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("retired segment remains: %v", err)
	}
	obj, err = q.take()
	if err != nil || obj.(*UpdateNewMessage).Message.Id != 2 {
		t.Fatalf("next segment = (%v, %v)", obj, err)
	}
}

func TestUpdateQueueCloseWakesReadersAndCleansSpool(t *testing.T) {
	q := newTestQueue(t, 1)
	done := make(chan error, 8)
	for range 8 {
		go func() { _, err := q.next(nil); done <- err }()
	}
	if err := q.close(); err != nil {
		t.Fatal(err)
	}
	for range 8 {
		if err := receiveTest(t, done); !errors.Is(err, ErrClientClosed) {
			t.Fatal(err)
		}
	}
	if err := q.close(); err != nil {
		t.Fatal(err)
	}
	if err := q.push(queueTestUpdate(1)); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
	q2 := newTestQueue(t, 1)
	for i := range int64(3) {
		if err := q2.push(queueTestUpdate(i)); err != nil {
			t.Fatal(err)
		}
	}
	dir := q2.dir
	if err := q2.close(); err != nil {
		t.Fatal(err)
	}
	if len(q2.memory) != 0 {
		t.Fatal("closed queue retained memory backlog")
	}
	if _, err := os.Stat(dir); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("closed queue retained spool")
	}
	stop := make(chan struct{})
	close(stop)
	if _, err := newTestQueue(t, 1).next(stop); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
}

func TestUpdateQueueCorruptSpool(t *testing.T) {
	for _, corrupt := range []string{"zero length", "oversized length", "truncated header", "invalid JSON", "unknown type", "missing segment"} {
		t.Run(corrupt, func(t *testing.T) {
			q := newTestQueue(t, 0)
			if err := q.push(queueTestUpdate(1)); err != nil {
				t.Fatal(err)
			}
			path := q.segment(q.head)
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			switch corrupt {
			case "zero length":
				binary.BigEndian.PutUint32(data[:4], 0)
			case "oversized length":
				binary.BigEndian.PutUint32(data[:4], ^uint32(0))
			case "truncated header":
				data = data[:2]
			case "invalid JSON":
				data[4] = '!'
			case "unknown type":
				body := []byte(`{"@type":"unknownFutureUpdate"}`)
				data = binary.BigEndian.AppendUint32(nil, uint32(len(body)))
				data = append(data, body...)
			case "missing segment":
				if err := os.Remove(path); err != nil {
					t.Fatal(err)
				}
			}
			if corrupt != "missing segment" {
				if err := os.WriteFile(path, data, 0600); err != nil {
					t.Fatal(err)
				}
			}
			if obj, err := q.take(); obj != nil || err == nil {
				t.Fatalf("corrupt spool = (%v, %v)", obj, err)
			}
		})
	}
}

func TestUpdateQueueConcurrentProducersAndConsumers(t *testing.T) {
	q := newTestQueue(t, 4)
	const producers, perProducer, consumers = 8, 32, 4
	stop := make(chan struct{})
	results := make(chan int64, producers*perProducer)
	errCh := make(chan error, producers+consumers)
	var workers sync.WaitGroup
	t.Cleanup(func() { close(stop); workers.Wait() })
	for range consumers {
		workers.Go(func() {
			for {
				obj, err := q.next(stop)
				if errors.Is(err, ErrClientClosed) {
					return
				}
				if err != nil {
					errCh <- err
					return
				}
				results <- obj.(*UpdateNewMessage).Message.Id
			}
		})
	}
	var writers sync.WaitGroup
	for p := range producers {
		writers.Go(func() {
			for i := range perProducer {
				if err := q.push(queueTestUpdate(int64(p*perProducer + i))); err != nil {
					errCh <- err
					return
				}
			}
		})
	}
	seen := make(map[int64]bool)
	for range producers * perProducer {
		id := receiveTest(t, results)
		if id < 0 || id >= producers*perProducer || seen[id] {
			t.Fatalf("lost/duplicate update ID %d", id)
		}
		seen[id] = true
	}
	writers.Wait()
	if err := q.close(); err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	go func() { workers.Wait(); close(done) }()
	receiveTest(t, done)
	select {
	case err := <-errCh:
		t.Fatal(err)
	default:
	}
}

func TestManagerStopAndClientRegistry(t *testing.T) {
	m := NewClientManager("unused")
	// Avoid starting the native receiver while exercising synchronized registry.
	m.once.Do(func() {})
	if err := m.AddClient(nil); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
	const count = 32
	clients := make([]*Client, count)
	for i := range clients {
		clients[i] = newTestClient(t)
		clients[i].clientID = i + 1
	}
	var wg sync.WaitGroup
	for _, c := range clients {
		wg.Go(func() {
			if err := m.AddClient(c); err != nil {
				t.Error(err)
			}
			_ = m.GetClients()
			_ = m.GetClient(c.clientID)
		})
	}
	wg.Wait()
	if len(m.GetClients()) != count {
		t.Fatalf("registry size = %d", len(m.GetClients()))
	}
	snapshot := m.GetClients()
	snapshot[0] = nil
	if len(m.GetClients()) != count {
		t.Fatal("GetClients exposed mutable registry")
	}
	other := NewClientManager("unused")
	other.once.Do(func() {})
	if err := other.AddClient(clients[0]); err == nil {
		t.Fatal("client moved between managers")
	}
	for _, c := range clients {
		c.closedOnce.Do(func() { close(c.closed) })
	}
	for range 8 {
		wg.Go(func() { ; m.Stop() })
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	receiveTest(t, done)
	if len(m.GetClients()) != 0 {
		t.Fatal("Stop retained clients")
	}
	for _, c := range clients {
		if c.manager.Load() != nil {
			t.Fatal("Stop retained client manager")
		}
	}
	if err := m.Start(); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
	if err := m.AddClient(clients[0]); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
}

func TestProtocolCompletionBypassesBlockedApplicationHandlers(t *testing.T) {
	c := newTestClient(t)
	entered, release, handled := make(chan struct{}), make(chan struct{}), make(chan struct{}, 1)
	defer close(release)
	c.AddHandler(testHandler{handle: func(_ *Client, u TlObject) error {
		if message, ok := u.(*UpdateNewMessage); ok && message.Message.Id == 1 {
			close(entered)
			<-release
			handled <- struct{}{}
		}
		return nil
	}})
	startTestProcessors(c, 1)
	if err := c.updates.push(queueTestUpdate(1)); err != nil {
		t.Fatal(err)
	}
	receiveTest(t, entered)
	// Saturate memory while the only application worker is blocked.
	for i := int64(2); i < 18; i++ {
		if err := c.updates.push(queueTestUpdate(i)); err != nil {
			t.Fatal(err)
		}
	}
	ready := make(chan struct{}, 1)
	c.sendRequest = func(_ int, data []byte) {
		c.acceptResponse(&Message{ChatId: 42, Id: -1, SendingState: &MessageSendingStatePending{}}, decodeTestRequest(t, data).Extra, nil)
		ready <- struct{}{}
	}
	done := make(chan error, 1)
	go func() { _, err := c.Send(&SendMessage{ChatId: 42}); done <- err }()
	receiveTest(t, ready)
	c.acceptResponse(&UpdateMessageSendSucceeded{OldMessageId: -1, Message: &Message{ChatId: 42, Id: 99}}, "", nil)
	if err := receiveTest(t, done); err != nil {
		t.Fatal(err)
	}
	assertNoPendingCalls(t, c)
	// Native closure must also be signalled without a free worker.
	c.acceptResponse(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateClosed{}}, "", nil)
	receiveTest(t, c.closed)
}

type testUnserializableUpdate struct{}

func (testUnserializableUpdate) GetType() string { return "updateTest" }
func (testUnserializableUpdate) MarshalJSON() ([]byte, error) {
	return nil, errors.New("cannot serialize")
}

func TestUpdateQueueFailedSpillPreservesBacklog(t *testing.T) {
	q := newTestQueue(t, 1)
	if err := q.push(queueTestUpdate(1)); err != nil {
		t.Fatal(err)
	}
	if err := q.push(testUnserializableUpdate{}); err == nil {
		t.Fatal("failed serialization accepted")
	}
	if q.dir != "" || q.spilled != 0 {
		t.Fatal("failed serialization created spool")
	}
	if obj, err := q.take(); err != nil || obj.(*UpdateNewMessage).Message.Id != 1 {
		t.Fatalf("backlog = (%v,%v)", obj, err)
	}
	if err := q.push(queueTestUpdate(2)); err != nil {
		t.Fatal(err)
	}
	if err := q.push(queueTestUpdate(3)); err != nil {
		t.Fatal(err)
	}
	// Simulate a broken writer after a valid disk record. A failed append must
	// leave the earlier memory/disk records intact and readable.
	if err := q.writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := q.push(queueTestUpdate(4)); err == nil {
		t.Fatal("closed writer accepted append")
	}
	q.writer = nil // Already closed; prevent an intentional double-close cleanup error.
	for want := int64(2); want <= 3; want++ {
		obj, err := q.take()
		if err != nil || obj.(*UpdateNewMessage).Message.Id != want {
			t.Fatalf("record after failed write = (%v,%v)", obj, err)
		}
	}
}
