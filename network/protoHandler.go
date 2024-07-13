package network

import (
	"strings"
	"time"

	"go.uber.org/zap"

	noise "github.com/GoldenDeals/noise-edits"
)

type ChatMessage struct {
	content string
}

// Marshal serializes a chat message into bytes.
func (m ChatMessage) Marshal() []byte {
	return []byte(m.content)
}

// Unmarshal deserializes a slice of bytes into a chat message, and returns an error should deserialization
// fail, or the slice of bytes be malformed.
func UnmarshalChatMessage(buf []byte) (ChatMessage, error) {
	return ChatMessage{content: strings.ToValidUTF8(string(buf), "")}, nil
}

func (n *Node) HandlerF(c noise.HandlerContext) error {
	obj, err := c.DecodeMessage()
	if err != nil {
		return nil
	}

	msg, ok := obj.(ChatMessage)
	if !ok {
		return nil
	}

	n.Log.Debug("Got Message", zap.String("id", c.ID().String()), zap.String("data", msg.content))
	time.Sleep(1 * time.Second)

	return c.SendMessage(msg)
}
