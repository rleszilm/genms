package rest

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

		if f := fopts.GetRest().GetName(); f != "" {
			return f, nil
		}

		if f := fopts.GetName(); f != "" {
			return f, nil
		}
	}

	if a.raw.GetRest().GetName() != "" {
		return a.raw.GetRest().GetName(), nil
	}

	return a.Arg.QueryName()
}

// IsQuery returns true if the field should be populated in the query.
func (a *Arg) IsQuery() bool {
	return a.raw.GetRest().GetLocation() == annotations.Arg_RestOptions_Query
}

// IsPath returns true if the field should be populated in the path.
func (a *Arg) IsPath() bool {
	return a.raw.GetRest().GetLocation() == annotations.Arg_RestOptions_Path
}

// IsBody returns true if the field should be populated in the body.
func (a *Arg) IsBody() bool {
	return a.raw.GetRest().GetLocation() == annotations.Arg_RestOptions_Body
}

// IsHeader returns true if the field should be populated as a header.
func (a *Arg) IsHeader() bool {
	return a.raw.GetRest().GetLocation() == annotations.Arg_RestOptions_Header
}
