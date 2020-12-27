package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	interfaceTemplateSource = `// Generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ DalPackageName .file }}

{{ with $state := . }}

var (
	// Err{{ MessageName .msg }}CollectionMethodImpl is returned when the called method is not implemented.
	Err{{ MessageName .msg }}CollectionMethodImpl = {{ .P.Errors }}.New("{{ MessageName .msg }}Collection method is not implemented")
)

// {{ MessageName .msg }}Collection is an autogenerated interface that can be used to interact with a collection of {{ MessageName .msg }} objects.
type {{ MessageName .msg }}Collection interface {
	{{ .P.Service }}.Service

	{{ MessageName .msg }}CollectionReader
	{{ MessageName .msg }}CollectionWriter
}

// {{ MessageName .msg }}CollectionWriter is an autogenerated interface that can be used to write to a collection of {{ MessageName .msg }} objects.
type {{ MessageName .msg }}CollectionWriter interface {
	Upsert({{ .P.Context }}.Context, *{{ QualifiedType .outfile .msg }}) (*{{ QualifiedType .outfile .msg }}, error)
}

// {{ MessageName .msg }}CollectionReader is an autogenerated interface that can be used to query a collection
// of {{ MessageName .msg }} objects. The queries and their values are taken from the representative proto message.
type {{ MessageName .msg }}CollectionReader interface {
	All({{ .P.Context }}.Context) ([]*{{ QualifiedType .outfile .msg }}, error)
	Filter({{ .P.Context }}.Context, *{{ MessageName .msg }}Fields) ([]*{{ QualifiedType .outfile .msg }}, error)
	{{ range .opts.Queries -}}
		{{ ToTitleCase .Name }}({{ $state.P.Context }}.Context{{ range .Args -}}, {{ GoFieldType $state.outfile ($state.fields.ByName .) }} {{- end }}) ([]*{{ QualifiedType $state.outfile $state.msg }}, error)
	{{ end }}

}

// {{ MessageName .msg }}Fields is an autogenerated struct that can be used in the generic queries against {{ MessageName .msg }}Collection.
type {{ MessageName .msg }}Fields struct {
	{{ range .fields.Names -}}
		{{ GoFieldName ($state.fields.ByName .) }} {{ Pointer (GoFieldType $state.outfile ($state.fields.ByName .)) }}
	{{ end -}}
}

// Unimplemented{{ MessageName .msg }}Collection is an autogenerated implementation of {{ MessageName .msg }}Collection that returns an error when any
// method is called.
type Unimplemented{{ MessageName .msg }}Collection struct {
	{{ .P.Service }}.Deps
}

// Upsert implements {{ MessageName .msg }}Collection.Upsert
func (x *Unimplemented{{ MessageName $state.msg }}Collection) Upsert(_ {{ .P.Context }}.Context, _ *{{ QualifiedType .outfile .msg }}) (*{{ QualifiedType .outfile .msg }}, error) {
	return nil, Err{{ MessageName $state.msg }}CollectionMethodImpl
}

// Filter implements {{ MessageName .msg }}Collection.Filter
func (x *Unimplemented{{ MessageName $state.msg }}Collection) Filter(_ {{ .P.Context }}.Context, _ *{{ MessageName $state.msg }}Fields) ([]*{{ QualifiedType .outfile .msg }}, error) {
	return nil, Err{{ MessageName $state.msg }}CollectionMethodImpl
}

{{ range .opts.Queries -}}
	// {{ ToTitleCase .Name }} implements {{ MessageName $state.msg }}Collection.{{ ToTitleCase .Name }}
	func (x *Unimplemented{{ MessageName $state.msg }}Collection){{ ToTitleCase .Name }}(_ {{ $state.P.Context }}.Context{{ range .Args -}}, _ {{ GoFieldType $state.outfile ($state.fields.ByName .) }} {{- end }}) ([]*{{ QualifiedType $state.outfile $state.msg }}, error) {
		return nil, Err{{ MessageName $state.msg }}CollectionMethodImpl
	}
{{ end }}
{{ end }}
`
)

var (
	interfaceTemplate = template.Must(
		template.New("interface").
			Funcs(template.FuncMap{
				"DalPackageName": DalPackageName,
				"MessageName":    MessageName,
				"GoFieldName":    GoFieldName,
				"GoFieldType":    GoFieldType,
				"ToTitleCase":    ToTitleCase,
				"ToSnakeCase":    ToSnakeCase,
				"QualifiedType":  QualifiedType,
				"Pointer": func(s string) string {
					if s[0] == '*' {
						return s
					}
					return "*" + s
				},
			}).
			Parse(interfaceTemplateSource))
)

// GenerateInterface generates the dal interface for the collection
func GenerateInterface(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	filename := fmt.Sprintf("dal/%s.genms.dal.%s.go", file.GeneratedFilenamePrefix, strings.ToLower(msg.GoIdent.GoName))
	outfile := plugin.NewGeneratedFile(filename, ".")

	p := map[string]string{
		"Errors":  QualifiedPackageName(outfile, "errors"),
		"Service": QualifiedPackageName(outfile, "github.com/rleszilm/gen_microservice/service"),
		"Context": QualifiedPackageName(outfile, "context"),
	}

	fields := NewFields(msg, annotations.DalOptions_BackEnd_None)

	model := map[string]interface{}{
		"file":    file,
		"msg":     msg,
		"opts":    opts,
		"fields":  fields,
		"outfile": outfile,
		"P":       p,
	}

	buf := &bytes.Buffer{}
	if err := interfaceTemplate.Execute(buf, model); err != nil {
		return err
	}

	if _, err := outfile.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}
