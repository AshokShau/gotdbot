package gotdbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/AshokShau/gotdbot/internal/tdjson"
)

// Send dispatches req to TDLib and waits for the response.
func (c *Client) Send(req TlObject) (TlObject, error) {
	return c.SendWithContext(context.Background(), req)
}

// SendWithContext dispatches req to TDLib and waits for the response, honouring ctx.
func (c *Client) SendWithContext(ctx context.Context, req TlObject) (TlObject, error) {
	if ctx == nil {
		return nil, fmt.Errorf("nil context")
	}
	if req == nil || (reflect.ValueOf(req).Kind() == reflect.Pointer && reflect.ValueOf(req).IsNil()) {
		return nil, fmt.Errorf("nil TDLib request")
	}
	if c.closing.Load() && req.GetType() != "close" {
		return nil, ErrClientClosed
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if _, isFn := req.(tlFunction); !isFn {
		return nil, fmt.Errorf("TDLib requests must be function pointers, got %T", req)
	}
	// Serialize without mutating the request; correlation belongs to the transport.
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var wire map[string]json.RawMessage
	if err := json.Unmarshal(data, &wire); err != nil {
		return nil, err
	}
	// Bound retries, lookup requests, send completion and flood waits together.
	ctx, cancel := context.WithTimeout(ctx, c.config.RequestTimeout)
	defer cancel()
	chatLoaded := req.GetType() == "getChat"
	messageLoaded := strings.HasPrefix(req.GetType(), "getMessage")
	floodRetried := false
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		select {
		case <-c.stop:
			return nil, ErrClientClosed
		default:
		}
		extra := strconv.FormatUint(c.requestID.Add(1), 10)
		wire["@extra"], _ = json.Marshal(extra)
		data, err := json.Marshal(wire)
		if err != nil {
			return nil, err
		}
		call := &pendingCall{response: make(chan TlObject, 1), waitForSend: waitsForSend(req)}
		c.requestMu.Lock()
		c.pendingRequests[extra] = call
		c.requestMu.Unlock()
		if c.sendRequest != nil {
			c.sendRequest(c.clientID, data)
		} else {
			tdjson.SendBytes(c.clientID, data)
		}
		var result TlObject
		select {
		case result = <-call.response:
		case <-ctx.Done():
			c.finishCall(extra)
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return nil, errors.Join(SendTimeout, ctx.Err())
			}
			return nil, ctx.Err()
		case <-c.closed:
			c.finishCall(extra)
			return nil, ErrClientClosed
		case <-c.stop:
			c.finishCall(extra)
			return nil, ErrClientClosed
		}
		if errObj, ok := result.(*Error); ok {
			c.finishCall(extra)
			retry, retryErr := c.retryRequest(ctx, req, errObj, &chatLoaded, &messageLoaded, &floodRetried)
			if retryErr != nil {
				return nil, retryErr
			}
			if retry {
				continue
			}
			return nil, errObj
		}
		if waitsForSend(req) {
			result, err = c.waitSend(ctx, result, call)
			if errors.Is(err, context.DeadlineExceeded) {
				err = errors.Join(SendTimeout, err)
			}
		}
		c.finishCall(extra)
		return result, err
	}
}

// Retry only explicit server rejections; never retry a transport timeout.
func (c *Client) retryRequest(ctx context.Context, req TlObject, e *Error, chatLoaded, messageLoaded, floodRetried *bool) (bool, error) {
	if c.config == nil || c.config.AutoRetry == nil {
		return false, nil
	}
	policy := c.config.AutoRetry
	if e.Code == 429 && !*floodRetried && policy.MaxFloodWait > 0 {
		seconds := e.GetRetryAfter()
		if seconds > 0 && int64(seconds) <= int64(policy.MaxFloodWait/time.Second) {
			*floodRetried = true
			timer := time.NewTimer(time.Duration(seconds) * time.Second)
			defer timer.Stop()
			select {
			case <-timer.C:
				return true, nil
			case <-ctx.Done():
				return false, ctx.Err()
			case <-c.stop:
				return false, ErrClientClosed
			}
		}
	}
	if e.Code != 400 {
		return false, nil
	}
	if e.Message == "Message not found" && !*messageLoaded && policy.MessageNotFound {
		*messageLoaded = true
		chatID, messageID := extractChatAndMessageID(req)
		if chatID != 0 && messageID != 0 {
			_, err := c.SendWithContext(ctx, &GetMessage{ChatId: chatID, MessageId: messageID})
			return err == nil, nil
		}
	}
	if e.Message == "Chat not found" && !*chatLoaded && policy.ChatNotFound {
		*chatLoaded = true
		if id, _ := extractChatAndMessageID(req); id != 0 {
			_, err := c.SendWithContext(ctx, &GetChat{ChatId: id})
			return err == nil, nil
		}
	}
	return false, nil
}

func extractChatAndMessageID(req TlObject) (int64, int64) {
	data, err := json.Marshal(req)
	if err != nil {
		return 0, 0
	}
	// TDLib accepts int64 values as JSON numbers or strings. json.Number keeps
	// their precision in either form; decoding through float64 loses large IDs.
	var ids struct {
		ChatID    json.Number `json:"chat_id"`
		MessageID json.Number `json:"message_id"`
	}
	if err := json.Unmarshal(data, &ids); err != nil {
		return 0, 0
	}
	chatID, _ := ids.ChatID.Int64()
	messageID, _ := ids.MessageID.Int64()
	return chatID, messageID
}
