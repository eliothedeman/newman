package newman

import "encoding"

// Message is something that can be marshaled to/from binary
type Message interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// MessageFactory returns a new Message
type MessageFactory func() Message
