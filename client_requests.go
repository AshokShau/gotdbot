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
func (c *Client) SendWithContext(ctx context.Context, req TlObject) (obj TlObject, err error) {
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
	started := time.Now()
	defer func() {
		if errors.Is(err, context.DeadlineExceeded) && !errors.Is(err, SendTimeout) {
			err = fmt.Errorf("%w: %w", SendTimeout, err)
		}
		args := []any{"user_id", c.Me().Id, "method", req.GetType(), "elapsed", time.Since(started)}
		if err == nil {
			c.Logger.Debug("TDLib request completed", args...)
			return
		}
		args = append(args, "error", err)
		if errors.Is(err, context.DeadlineExceeded) {
			c.Logger.Warn("TDLib request timed out", args...)
		} else {
			c.Logger.Debug("TDLib request failed", args...)
		}
	}()
	// Serialize without mutating the request; correlation belongs to the transport.
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var wire map[string]json.RawMessage
	if err := json.Unmarshal(data, &wire); err != nil {
		return nil, err
	}
	if c.config.RequestTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.config.RequestTimeout)
		defer cancel()
	}
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
		c.Logger.Debug("Sending TDLib request", "user_id", c.Me().Id, "method", req.GetType(), "request_id", extra)
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
	if e.Code == 429 {
		seconds := e.GetRetryAfter()
		reason := ""
		switch {
		case *floodRetried:
			reason = "retry already attempted"
		case policy.MaxFloodWait <= 0:
			reason = "automatic flood retry disabled"
		case seconds <= 0:
			reason = "invalid retry delay"
		case int64(seconds) > int64(policy.MaxFloodWait/time.Second):
			reason = "retry delay exceeds MaxFloodWait"
		}
		if reason == "" {
			*floodRetried = true
			delay := time.Duration(seconds) * time.Second
			c.Logger.Warn("TDLib flood wait; delaying request retry", "user_id", c.Me().Id, "method", req.GetType(), "retry_after", delay)
			timer := time.NewTimer(delay)
			defer timer.Stop()
			select {
			case <-timer.C:
				c.Logger.Debug("Retrying TDLib request after flood wait", "user_id", c.Me().Id, "method", req.GetType())
				return true, nil
			case <-ctx.Done():
				return false, ctx.Err()
			case <-c.closed:
				return false, ErrClientClosed
			case <-c.stop:
				return false, ErrClientClosed
			}
		}
		c.Logger.Warn("TDLib flood wait; request will not be retried", "user_id", c.Me().Id, "method", req.GetType(), "retry_after_seconds", seconds, "max_flood_wait", policy.MaxFloodWait, "reason", reason)
		return false, nil
	}
	if e.Code != 400 {
		return false, nil
	}
	if e.Message == "Message not found" && !*messageLoaded && policy.MessageNotFound {
		*messageLoaded = true
		chatID, messageID := extractChatAndMessageID(req)
		if chatID != 0 && messageID != 0 {
			c.Logger.Debug("Loading missing message before request retry", "user_id", c.Me().Id, "method", req.GetType(), "chat_id", chatID, "message_id", messageID)
			_, err := c.SendWithContext(ctx, &GetMessage{ChatId: chatID, MessageId: messageID})
			if err != nil {
				c.Logger.Warn("Cannot load missing message for request retry", "user_id", c.Me().Id, "method", req.GetType(), "error", err)
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) || errors.Is(err, ErrClientClosed) {
					return false, err
				}
				return false, nil
			}
			c.Logger.Debug("Retrying TDLib request after loading message", "user_id", c.Me().Id, "method", req.GetType())
			return true, nil
		}
	}
	if e.Message == "Chat not found" && !*chatLoaded && policy.ChatNotFound {
		*chatLoaded = true
		if id, _ := extractChatAndMessageID(req); id != 0 {
			c.Logger.Debug("Loading missing chat before request retry", "user_id", c.Me().Id, "method", req.GetType(), "chat_id", id)
			_, err := c.SendWithContext(ctx, &GetChat{ChatId: id})
			if err != nil {
				c.Logger.Warn("Cannot load missing chat for request retry", "user_id", c.Me().Id, "method", req.GetType(), "error", err)
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) || errors.Is(err, ErrClientClosed) {
					return false, err
				}
				return false, nil
			}
			c.Logger.Debug("Retrying TDLib request after loading chat", "user_id", c.Me().Id, "method", req.GetType())
			return true, nil
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
