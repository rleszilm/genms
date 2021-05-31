package generator

import (
	"errors"
	"log"
	"strings"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

// Arg wraps a QueryArg so additional logic can be applied.
type Arg struct {
	arg   *annotations.Arg
	file  *File
	field *Field
}

// NewArg returns a new Arg
func NewArg(file *File, fields *Fields, arg *annotations.Arg) *Arg {
	field := fields.ByName(arg.GetName())
	return &Arg{
		arg:   arg,
		file:  file,
		field: field,
	}
}

// Name returns the name of the arg.
func (a *Arg) Name() (string, error) {
	if a.field != nil {
		return a.field.Name(), nil
	}

	if a.arg.GetName() == "" {
		return "", errors.New("arg has no name")
	}
	return a.arg.Name, nil
}

// QueryName returns the name of the field in the remote system.
func (a *Arg) QueryName() (string, error) {
	if a.field != nil {
		return a.field.QueryName(), nil
	}

	if a.arg.GetName() == "" {
		return "", errors.New("arg has no name")
	}
	return a.arg.Name, nil
}

// QualifiedKind returns the kind of the arg.
func (a *Arg) QualifiedKind() (string, error) {
	if a.field != nil {
		return a.field.QualifiedKind(), nil
	}

	if a.arg.GetKind() == "" {
		log.Println("no kind", a.arg)
		return "", errors.New("arg has no kind")
	}

	tokens := strings.Split(a.arg.GetKind(), ".")
	if len(tokens) == 1 {
		return tokens[0], nil
	}

	ident := protogen.GoIdent{
		GoImportPath: protogen.GoImportPath(strings.Join(tokens[:len(tokens)-1], ".")),
		GoName:       tokens[len(tokens)-1],
	}
	return a.file.QualifiedKind(ident), nil
}

// ToRef returns the string needed to make a reference of the field.
func (a *Arg) ToRef() (string, error) {
	if a.field != nil {
		return a.field.ToRef(), nil
	}

	str, err := a.QualifiedKind()
	if err != nil {
		return "", err
	}

	if str[0] == '*' {
		return "&", nil
	}
	return "", nil
}
