package protocgenlib

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Message adds functionality to the underlying message.
type Message struct {
	Outfile *protogen.GeneratedFile
	Message *protogen.Message
}

// NewMessage returns a new Message.
func NewMessage(outfile *protogen.GeneratedFile, message *protogen.Message) *Message {
	return &Message{
		Outfile: outfile,
		Message: message,
	}
}

// Name returns the name of the message.
func (m *Message) Name() string {
	return m.Message.GoIdent.GoName
}

// QualifiedType returns the fully qualified type within the dal.
func (m *Message) QualifiedType() string {
	return m.Outfile.QualifiedGoIdent(m.Message.GoIdent)
}
