package mongo

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
)

// Arg wraps a QueryArg so additional logic can be applied.
type Arg struct {
	*generator.Arg

	field *Field
	raw   *annotations.Arg
}

// NewArg returns a new Arg
func NewArg(file *File, fields *Fields, arg *annotations.Arg) *Arg {
	field := fields.ByName(arg.GetName())

	return &Arg{
		Arg:   generator.NewArg(file.Generator(), fields.Generator(), arg),
		field: field,
		raw:   arg,
	}
}

// QueryName returns the name of the field in the remote system.
func (a *Arg) QueryName() (string, error) {
	if a.field != nil {
		fopts := a.field.Options()

		if f := fopts.GetMongo().GetField(); f != "" {
			return f, nil
		}

		if f := fopts.GetField(); f != "" {
			return f, nil
		}
	}

	if n := a.raw.GetMongo().GetName(); n != "" {
		return n, nil
	}

	return a.Arg.QueryName()
}
