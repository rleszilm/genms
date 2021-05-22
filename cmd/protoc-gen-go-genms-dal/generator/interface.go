package generator

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Interface generates the interface for the dal type.
type Interface struct {
	File    *File
	Message *Message
	Fields  *Fields
	Queries *Queries
	Opts    *annotations.DalOptions
}

// NewInterface returns a new Interfaces
func NewInterface(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) *Interface {
	base := path.Base(file.GeneratedFilenamePrefix)
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/%s.genms.dal.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
	outfile := plugin.NewGeneratedFile(filename, ".")

	ifile := NewFile(outfile, file)
	imsg := NewMessage(ifile, msg)
	ifields := NewFields(imsg)
	iqueries := NewQueries(ifile, ifields, opts)

	return &Interface{
		File:    ifile,
		Message: imsg,
		Fields:  ifields,
		Queries: iqueries,
		Opts:    opts,
	}
}

// GenerateInterface generates the dal interface for the collection
func GenerateInterface(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	i := NewInterface(plugin, file, msg, opts)
	return i.render()
}

func (i *Interface) render() error {
	steps := []func() error{
		i.definePackage,
		i.defineErrors,
		i.defineInterface,
		i.defineFieldValues,
		i.defineUnimplemented,
	}

	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	return nil
}

func (i *Interface) definePackage() error {
	tmplSrc := `// Package {{ .File.DalPackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ .File.DalPackageName }}

`

	tmpl, err := template.New("defineInterfacePackage").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, i); err != nil {
		return err
	}

	if _, err := i.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (i *Interface) defineErrors() error {
	tmplSrc := `var (
	// Err{{ .I.Message.Name }}CollectionMethodImpl is returned when the called method is not implemented.
	Err{{ .I.Message.Name }}CollectionMethodImpl = {{ .P.Errors }}.New("{{ .I.Message.Name }}Collection method is not implemented")
)

`

	tmpl, err := template.New("defineInterfaceErrors").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Errors": i.File.QualifiedPackageName("errors"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"I": i,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := i.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (i *Interface) defineInterface() error {
	tmplSrc := `{{- $I := .I -}}
{{- $P := .P -}}
{{- $Generate := "go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 ." -}}
// {{ .I.Message.Name }}Collection is an autogenerated interface that can be used to interact with a collection of {{ .I.Message.Name }} objects.
//{{ $Generate }} {{ .I.Message.Name }}Collection
type {{ .I.Message.Name }}Collection interface {
	{{ .P.Service }}.Service

	{{ .I.Message.Name }}CollectionReader
	{{ .I.Message.Name }}CollectionWriter
}

// {{ .I.Message.Name }}CollectionWriter is an autogenerated interface that can be used to write to a collection of {{ .I.Message.Name }} objects.
//{{ $Generate }} {{ .I.Message.Name }}CollectionWriter
type {{ .I.Message.Name }}CollectionWriter interface {
	// Insert runs the command to generate a new object within the data store.
	Insert({{ .P.Context }}.Context, *{{ .I.Message.QualifiedKind }}) (*{{ .I.Message.QualifiedKind }}, error)
	// Upsert runs the command to overwrite the object in the datastore, or write it if it does nto already exist.
	Upsert({{ .P.Context }}.Context, *{{ .I.Message.QualifiedKind }}) (*{{ .I.Message.QualifiedKind }}, error)
	// Update runs the command to make changes to the given record.
	Update({{ .P.Context }}.Context, *{{ .I.Message.QualifiedKind }}, *{{ .I.Message.Name }}FieldValues) (*{{ .I.Message.QualifiedKind }}, error)
}

// {{ .I.Message.Name }}CollectionReader is an autogenerated interface that can be used to query a collection
// of {{ .I.Message.Name }} objects. The queries and their values are taken from the representative proto message.
//{{ $Generate }} {{ .I.Message.Name }}CollectionReader
type {{ .I.Message.Name }}CollectionReader interface {
	All({{ .P.Context }}.Context) ([]*{{ .I.Message.QualifiedKind }}, error)
	Filter({{ .P.Context }}.Context, *{{ .I.Message.Name }}FieldValues) ([]*{{ .I.Message.QualifiedKind }}, error)
	{{ range $qn := .I.Queries.Names -}}
		{{- $q := ($I.Queries.ByName $qn) -}}
		{{ ToTitleCase $qn }}(_ {{ $P.Context }}.Context
			{{- range $a := $q.Args -}}
				{{- $arg := (Arg $I.File $I.Fields $a) -}}
				, _ {{ $arg.QualifiedKind }}
			{{- end }}) ([]*{{ $I.Message.QualifiedKind }}, error)
	{{ end }}	
}

`

	tmpl, err := template.New("defineInterfaceInterface").
		Funcs(template.FuncMap{
			"Arg":         NewArg,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Context": i.File.QualifiedPackageName("context"),
		"Service": i.File.QualifiedPackageName("github.com/rleszilm/genms/service"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"I": i,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := i.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (i *Interface) defineFieldValues() error {
	tmplSrc := `{{- $i := .I -}}
// {{ $i.Message.Name }}FieldValues is an autogenerated struct that can be used in the generic queries against {{ $i.Message.Name }}Collection.
type {{ $i.Message.Name }}FieldValues struct {
	{{ range $name := $i.Fields.Names -}}
		{{- $f := ($i.Fields.ByName $name) -}}
		{{- if not $f.Ignore }}
			{{ ToTitleCase $f.Name }} {{ AsPointer $f.QualifiedKind -}}
		{{ end -}}
	{{ end -}}
}

`

	tmpl, err := template.New("defineInterfaceFieldValues").
		Funcs(template.FuncMap{
			"AsPointer":   protocgenlib.AsPointer,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Errors": i.File.QualifiedPackageName("errors"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"I": i,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := i.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (i *Interface) defineUnimplemented() error {
	tmplSrc := `{{- $I := .I -}}
{{- $P := .P -}}
// Unimplemented{{ .I.Message.Name }}Collection is an autogenerated implementation of {{ .I.Message.Name }}Collection that returns an error when any
// method is called.
type Unimplemented{{ .I.Message.Name }}Collection struct {
	{{ .P.Service }}.Dependencies
}

// Insert implements {{ .I.Message.Name }}Collection.Insert
func (x *Unimplemented{{ .I.Message.Name }}Collection) Insert(_ {{ .P.Context }}.Context, _ *{{ .I.Message.QualifiedKind }}) (*{{ .I.Message.QualifiedKind }}, error) {
	return nil, Err{{ .I.Message.Name }}CollectionMethodImpl
}

// Upsert implements {{ .I.Message.Name }}Collection.Upsert
func (x *Unimplemented{{ .I.Message.Name }}Collection) Upsert(_ {{ .P.Context }}.Context, _ *{{ .I.Message.QualifiedKind }}) (*{{ .I.Message.QualifiedKind }}, error) {
	return nil, Err{{ .I.Message.Name }}CollectionMethodImpl
}

// Update implements {{ .I.Message.Name }}Collection.Update
func (x *Unimplemented{{ .I.Message.Name }}Collection) Update(_ {{ .P.Context }}.Context, _ *{{ .I.Message.QualifiedKind }}, _ *{{ .I.Message.Name }}FieldValues) (*{{ .I.Message.QualifiedKind }}, error) {
	return nil, Err{{ .I.Message.Name }}CollectionMethodImpl
}

// Filter implements {{ .I.Message.Name }}Collection.Filter
func (x *Unimplemented{{ .I.Message.Name }}Collection) Filter(_ {{ .P.Context }}.Context, _ *{{ .I.Message.Name }}FieldValues) ([]*{{ .I.Message.QualifiedKind }}, error) {
	return nil, Err{{ .I.Message.Name }}CollectionMethodImpl
}

{{ range $qn := .I.Queries.Names -}}
	{{- $q := ($I.Queries.ByName $qn) -}}
	// {{ ToTitleCase $q.Name }} implements {{ $I.Message.Name }}Collection.{{ ToTitleCase $q.Name }}
	func (x *Unimplemented{{ $I.Message.Name }}Collection){{ ToTitleCase $q.Name }}(_ {{ $P.Context }}.Context
		{{- range $a := $q.Args -}}
			{{- $arg := (Arg $I.File $I.Fields $a) -}}
			, _ {{ $arg.QualifiedKind }}
		{{- end }}) ([]*{{ $I.Message.QualifiedKind }}, error) {
		return nil, Err{{ $I.Message.Name }}CollectionMethodImpl
	}
{{ end }}

func ReturnsOne{{ .I.Message.Name }}(xs []*{{ .I.Message.QualifiedKind }}, err error) (*{{ .I.Message.QualifiedKind }}, error) {
	if err != nil {
		return nil, err
	}

	switch len(xs) {
	case 0:
		return nil, err
	case 1:
		return xs[0], err
	default:
		return nil, {{ .P.Fmt }}.Errorf("{{ ToLowerCase .I.Message.QualifiedKind }}: more than 1 value returned - %w", err)
	}
}

`

	tmpl, err := template.New("defineInterfaceUnimplemented").
		Funcs(template.FuncMap{
			"Arg":         NewArg,
			"ToLowerCase": strings.ToLower,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Context": i.File.QualifiedPackageName("context"),
		"Errors":  i.File.QualifiedPackageName("errors"),
		"Fmt":     i.File.QualifiedPackageName("fmt"),
		"Service": i.File.QualifiedPackageName("github.com/rleszilm/genms/service"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"I": i,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := i.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
