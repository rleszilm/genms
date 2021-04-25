package rest

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
)

// Arg wraps a QueryArg so additional logic can be applied.
type Arg struct {
	*generator.Arg

	raw *annotations.DalOptions_Query_Arg
}

// NewArg returns a new Arg
func NewArg(file *File, fields *Fields, arg *annotations.DalOptions_Query_Arg) *Arg {
	return &Arg{
		Arg: generator.NewArg(file.Generator(), fields.Generator(), arg),
		raw: arg,
	}
}

// QueryName returns the name of the field in the remote system.
func (a *Arg) QueryName() (string, error) {
	if a.raw.GetRest().GetName() != "" {
		return a.raw.GetRest().GetName(), nil
	}

	return a.Arg.QueryName()
}

// IsQuery returns true if the field should be populated in the query.
func (a *Arg) IsQuery() bool {
	return a.raw.GetRest().GetLocation() == annotations.DalOptions_Query_Arg_Rest_Query
}

// IsPath returns true if the field should be populated in the path.
func (a *Arg) IsPath() bool {
	return a.raw.GetRest().GetLocation() == annotations.DalOptions_Query_Arg_Rest_Path
}

// IsBody returns true if the field should be populated in the body.
func (a *Arg) IsBody() bool {
	return a.raw.GetRest().GetLocation() == annotations.DalOptions_Query_Arg_Rest_Body
}

// IsHeader returns true if the field should be populated as a header.
func (a *Arg) IsHeader() bool {
	return a.raw.GetRest().GetLocation() == annotations.DalOptions_Query_Arg_Rest_Header
}
