package generator

import (
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// Fields is a struct that contains data about the messages fields.
type Fields struct {
	fieldNames      []string
	fieldsByName    map[string]*protogen.Field
	queryByField    map[string]string
	fieldByQuery    map[string]string
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

func (f *Fields) FieldNameByQueryName(n string) string {
	return f.fieldByQuery[n]
}

func (f *Fields) QueryNameByFieldName(n string) string {
	return f.queryByField[n]
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
	fieldByQuery := map[string]string{}
	queryByField := map[string]string{}

	for _, f := range msg.Fields {
		fName := string(f.Desc.Name())
		fieldNames = append(fieldNames, fName)
		fieldsByName["f:-"+fName] = f

		fOpts := f.Desc.Options()
		if proto.HasExtension(fOpts, annotations.E_FieldOptions) {
			ext := proto.GetExtension(fOpts, annotations.E_FieldOptions).(*annotations.DalFieldOptions)
			if ext != nil {
				if ext.Ignore {
					continue
				}

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
					fieldByQuery[beOpts.Field] = fName
					queryByField[fName] = beOpts.Field
					continue
				}

				if ext.Field != "" {
					queryFieldNames = append(queryFieldNames, ext.Field)
					fieldsByName["q:-"+ext.Field] = f
					fieldByQuery[ext.Field] = fName
					queryByField[fName] = ext.Field
					continue
				}
			}
		}
		queryFieldNames = append(queryFieldNames, fName)
		fieldsByName["q:-"+fName] = f
		fieldByQuery[fName] = fName
		queryByField[fName] = fName
	}

	return &Fields{
		fieldNames:      fieldNames,
		fieldsByName:    fieldsByName,
		queryFieldNames: queryFieldNames,
		fieldByQuery:    fieldByQuery,
		queryByField:    queryByField,
	}
}
