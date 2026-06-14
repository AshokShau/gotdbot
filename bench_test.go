package gotdbot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func newBenchClient() *Client {
	c := &Client{}
	c.handlers.Store(&handlersData{
		handlers: make(map[int][]Handler),
	})
	c.waiters = make(map[int64]map[string]*Waiter)
	return c
}

// BenchmarkSend measures the performance of sending a request.
func BenchmarkSend(b *testing.B) {
	c := newBenchClient()
	req := &GetChat{ChatId: 123456789}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		extra := strconv.FormatUint(c.requestID.Add(1), 10)
		req.setExtra(extra)
		_, err := json.Marshal(req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkHandlerDispatch measures the overhead of the processor dispatching an update to handlers.
func BenchmarkHandlerDispatch(b *testing.B) {
	counts := []int{10, 50, 100, 500}
	for _, count := range counts {
		b.Run(fmt.Sprintf("Handlers_%d", count), func(b *testing.B) {
			c := newBenchClient()
			for i := 0; i < count; i++ {
				c.AddHandlerGroup(&benchHandler{}, i%10)
			}

			update := &UpdateNewMessage{Message: &Message{Id: 1}}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hData := c.handlers.Load()
				groups := hData.groups
				handlersMap := hData.handlers
				for _, group := range groups {
					groupHandlers := handlersMap[group]
					for _, h := range groupHandlers {
						if h.CheckUpdate(c, update) {
							_ = h.HandleUpdate(c, update)
							break
						}
					}
				}
			}
		})
	}
}

// BenchmarkWaiterMatching measures the cost of matching an update against active waiters.
func BenchmarkWaiterMatching(b *testing.B) {
	counts := []int{10, 100, 1000, 10000}
	for _, count := range counts {
		b.Run(fmt.Sprintf("Waiters_%d", count), func(b *testing.B) {
			c := newBenchClient()
			chatId := int64(12345)

			for i := 0; i < count; i++ {
				id := fmt.Sprintf("w%d", i)
				cid := int64(i)
				c.wMu.Lock()
				if c.waiters[cid] == nil {
					c.waiters[cid] = make(map[string]*Waiter)
				}
				c.waiters[cid][id] = &Waiter{
					Filter: func(cl *Client, u TlObject) bool { return false },
				}
				c.wMu.Unlock()
			}

			update := &UpdateNewMessage{Message: &Message{ChatId: chatId}}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				c.wMu.RLock()
				if len(c.waiters) > 0 {
					var matchedWaiters []*Waiter
					collectWaiters := func(id int64) {
						if inner, ok := c.waiters[id]; ok {
							for _, w := range inner {
								if w.Filter(c, update) {
									matchedWaiters = append(matchedWaiters, w)
								}
							}
						}
					}
					collectWaiters(0)
					collectWaiters(chatId)
					_ = matchedWaiters
				}
				c.wMu.RUnlock()
			}
		})
	}
}

// BenchmarkUpdateUnmarshal measures the standard unmarshaling cost of an incoming TDLib update.
func BenchmarkUpdateUnmarshal(b *testing.B) {
	data := []byte(`{"@type":"updateNewMessage","@client_id":1,"@extra":"12345","message":{"id":1,"chat_id":123,"content":{"@type":"messageText","text":{"text":"hello"}}}}`)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, err := UnmarshalWithClient(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

type benchHandler struct{}

func (h *benchHandler) CheckUpdate(c *Client, u TlObject) bool   { return true }
func (h *benchHandler) HandleUpdate(c *Client, u TlObject) error { return nil }
