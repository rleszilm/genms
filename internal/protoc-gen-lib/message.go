package protocgenlib

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Message adds functionality to the underlying message.
type Message struct {
	file    *File
	message *protogen.Message
}

// NewMessage returns a new Message.
func NewMessage(file *File, message *protogen.Message) *Message {
	return &Message{
		file:    file,
		message: message,
	}
}

// Proto returns the base protogen object.
func (m *Message) Proto() *protogen.Message {
	return m.message
}

// File returns the base File object.
func (m *Message) File() *File {
	return m.file
}

// Outfile returns the file to which this field would be written.
func (m *Message) Outfile() *protogen.GeneratedFile {
	return m.file.Outfile()
}

// Name returns the name of the message.
func (m *Message) Name() string {
	return m.message.GoIdent.GoName
}

// Kind returns the fields go type.
func (m *Message) Kind() string {
	return m.message.GoIdent.GoName
}

// QualifiedKind returns the fully qualified type.
func (m *Message) QualifiedKind() string {
	return m.file.QualifiedKind(m.message.GoIdent)
}
