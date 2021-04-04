package postgres

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Message adds functionality to the underlying message.
type Message struct {
	*generator.Message
}

// NewMessage returns a new Message.
func NewMessage(file *File, msg *protogen.Message) *Message {
	return AsMessage(protocgenlib.NewMessage(file.ProtocGenLib(), msg))
}

// AsMessage wraps a Message.
func AsMessage(msg *protocgenlib.Message) *Message {
	return &Message{
		Message: generator.AsMessage(msg),
	}
}
