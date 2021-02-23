package generator

import (
	protocgenlib "github.com/rleszilm/gen_microservice/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Message adds functionality to the underlying message.
type Message struct {
	protocgenlib.Message
}

// NewMessage returns a new Message.
func NewMessage(outfile *protogen.GeneratedFile, message *protogen.Message) *Message {
	return &Message{
		Message: protocgenlib.Message{
			Outfile: outfile,
			Message: message,
		},
	}
}

// QualifiedDalType returns the fully qualified type within the dal.
func (m *Message) QualifiedDalType() string {
	i := protogen.GoIdent{
		GoName:       m.Message.Message.GoIdent.GoName,
		GoImportPath: m.Message.Message.GoIdent.GoImportPath + "/dal",
	}
	return m.Message.Outfile.QualifiedGoIdent(i)
}
