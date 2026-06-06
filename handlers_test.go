package gotdbot

import (
	"errors"
	"sort"
	"testing"
)

type mockHandler struct {
	handled bool
	err     error
	id      string
}

func (h *mockHandler) CheckUpdate(client *Client, update TlObject) bool {
	return true
}

func (h *mockHandler) HandleUpdate(client *Client, update TlObject) error {
	h.handled = true
	return h.err
}

func TestHandlerGroups(t *testing.T) {
	c := &Client{}
	handlersMap := make(map[int][]Handler)
	c.handlers.Store(&handlersMap)

	h1 := &mockHandler{id: "h1"}
	h2 := &mockHandler{id: "h2"}
	h3 := &mockHandler{id: "h3"}

	c.AddHandlerGroup(h1, 10)
	c.AddHandlerGroup(h2, 0)
	c.AddHandlerGroup(h3, -1)

	m := *c.handlers.Load()
	if len(m) != 3 {
		t.Fatalf("expected 3 groups, got %d", len(m))
	}

	groups := make([]int, 0, len(m))
	for k := range m {
		groups = append(groups, k)
	}
	sort.Ints(groups)

	if groups[0] != -1 || groups[1] != 0 || groups[2] != 10 {
		t.Errorf("groups not in correct order: %v", groups)
	}
}

func TestGroupFlowControl(t *testing.T) {
	c := &Client{
		Logger: DefaultClientConfig().Logger,
	}
	handlersMap := make(map[int][]Handler)
	c.handlers.Store(&handlersMap)

	h1 := &mockHandler{err: ContinueGroups, id: "h1"}
	h2 := &mockHandler{id: "h2"} // Should be skipped because h1 matched and succeeded/ContinueGroups
	h3 := &mockHandler{id: "h3"}

	c.AddHandlerGroup(h1, 0)
	c.AddHandlerGroup(h2, 0)
	c.AddHandlerGroup(h3, 1)

	// Simulate processor logic
	update := &UpdateNewMessage{}
	m := *c.handlers.Load()
	groups := make([]int, 0, len(m))
	for k := range m {
		groups = append(groups, k)
	}
	sort.Ints(groups)

	for _, group := range groups {
		groupHandlers := m[group]
		for _, h := range groupHandlers {
			if h.CheckUpdate(c, update) {
				err := h.HandleUpdate(c, update)
				if errors.Is(err, EndGroups) {
					goto end
				}
				if errors.Is(err, ContinueGroups) {
					break // Next group
				}
				// Success, next group
				goto nextGroup
			}
		}
	nextGroup:
	}
end:

	if !h1.handled {
		t.Error("h1 should have been handled")
	}
	if h2.handled {
		t.Error("h2 should NOT have been handled")
	}
	if !h3.handled {
		t.Error("h3 should have been handled")
	}
}
