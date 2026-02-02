package gotdbot

import (
	"fmt"
	"time"
)

// WaitMessage waits for the message to be sent and returns the final message.
func (c *Client) WaitMessage(msg *Message) (*Message, error) {
	if msg.SendingState != nil && msg.SendingState.Type() == "messageSendingStatePending" {
		key := fmt.Sprintf("%d:%d", msg.ChatId, msg.Id)
		ch := make(chan TlObject, 1)
		c.pendingMessages.Store(key, ch)
		defer c.pendingMessages.Delete(key)

		select {
		case res := <-ch:
			if errObj, ok := res.(*Error); ok {
				return nil, fmt.Errorf("TDLib error %d: %s", errObj.Code, errObj.Message)
			}
			if finalMsg, ok := res.(*Message); ok {
				return finalMsg, nil
			}
			return nil, fmt.Errorf("unexpected response type from waiter: %T", res)
		case <-time.After(30 * time.Second):
			return msg, nil
		}
	}

	return msg, nil
}
