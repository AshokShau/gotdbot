package gotdbot

// UpdateFilter is a function that returns true if the update matches the filter.
type UpdateFilter func(client *Client, update TlObject) bool

type Handler interface {
	CheckUpdate(client *Client, update TlObject) bool
	HandleUpdate(client *Client, update TlObject) error
}

type internalHandler struct {
	client     *Client
	handleFunc func(client *Client, update TlObject) error
	updateType string
}

func (h *internalHandler) CheckUpdate(client *Client, update TlObject) bool {
	return (h.client == nil || h.client == client) && update.GetType() == h.updateType
}

func (h *internalHandler) HandleUpdate(client *Client, update TlObject) error {
	return h.handleFunc(client, update)
}
