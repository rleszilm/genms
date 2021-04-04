package generator

import (
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Message adds functionality to the underlying message.
type Message struct {
	*protocgenlib.Message
}

// NewMessage returns a new Message.
func NewMessage(file *File, msg *protogen.Message) *Message {
	return AsMessage(protocgenlib.NewMessage(file.File, msg))
}

// AsMessage wraps a Message.
func AsMessage(msg *protocgenlib.Message) *Message {
	return &Message{
		Message: msg,
	}
}

// QualifiedDalKind returns the fully qualified type within the dal.
func (m *Message) QualifiedDalKind() string {
	i := protogen.GoIdent{
		GoName:       m.Proto().GoIdent.GoName,
		GoImportPath: m.Proto().GoIdent.GoImportPath + "/dal",
	}
	return m.Message.Outfile().QualifiedGoIdent(i)
}
