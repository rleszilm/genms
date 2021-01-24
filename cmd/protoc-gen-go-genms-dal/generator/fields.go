package generator

import (
	"log"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// Fields is a struct that contains data about the messages fields.
type Fields struct {
	fieldNames      []string
	fieldsByName    map[string]*protogen.Field
	queryFieldNames []string
}

// Names returns the field names.
func (f *Fields) Names() []string {
	return f.fieldNames
}

// ByName returns the specified field.
func (f *Fields) ByName(n string) *protogen.Field {
	return f.fieldsByName["f:-"+n]
}

// QueryNames returns the field names.
func (f *Fields) QueryNames() []string {
	return f.queryFieldNames
}

// ByQueryName returns the field based off the query name.
func (f *Fields) ByQueryName(n string) *protogen.Field {
	return f.fieldsByName["q:-"+n]
}

// NewFields returns a new Fields.
func NewFields(msg *protogen.Message, be annotations.DalOptions_Backend) *Fields {
	fieldNames := []string{}
	fieldsByName := map[string]*protogen.Field{}
	queryFieldNames := []string{}

	for _, f := range msg.Fields {
		fName := string(f.Desc.Name())
		fieldNames = append(fieldNames, fName)
		fieldsByName["f:-"+fName] = f

		fOpts := f.Desc.Options()
		if proto.HasExtension(fOpts, annotations.E_GenmsDalField) {
			ext := proto.GetExtension(fOpts, annotations.E_GenmsDalField).(*annotations.DalFieldOptions)
			if ext != nil {
				if ext.Ignore {
					continue
				}

				log.Println(ext)
				var beOpts *annotations.DalFieldOptions_BackendFieldOptions
				switch be {
				case annotations.DalOptions_BackEnd_Postgres:
					beOpts = ext.Postgres
				case annotations.DalOptions_BackEnd_Rest:
					beOpts = ext.Rest
				}

				if beOpts != nil && beOpts.Field != "" {
					queryFieldNames = append(queryFieldNames, beOpts.Field)
					fieldsByName["q:-"+beOpts.Field] = f
					continue
				}

				if ext.Field != "" {
					queryFieldNames = append(queryFieldNames, ext.Field)
					fieldsByName["q:-"+ext.Field] = f
					continue
				}
			}
		}
		queryFieldNames = append(queryFieldNames, fName)
		fieldsByName["q:-"+fName] = f
	}

	return &Fields{
		fieldNames:      fieldNames,
		fieldsByName:    fieldsByName,
		queryFieldNames: queryFieldNames,
	}
}
