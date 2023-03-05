package queue

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// For using when Publish new message to topic
func NewMessage(payload []byte) *message.Message {
	return message.NewMessage(watermill.NewUUID(), payload)
}
