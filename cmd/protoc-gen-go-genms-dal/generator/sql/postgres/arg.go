package postgres

import (
	"errors"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
)

// Arg wraps a Arg so additional logic can be applied.
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

// Comparison returns the check to perform
func (a *Arg) Comparison() (string, error) {
	switch a.raw.GetComparison() {
	case annotations.DalOptions_Query_Arg_EQ:
		return "=", nil
	case annotations.DalOptions_Query_Arg_NE:
		return "!=", nil
	case annotations.DalOptions_Query_Arg_GT:
		return ">", nil
	case annotations.DalOptions_Query_Arg_LT:
		return "<", nil
	case annotations.DalOptions_Query_Arg_GTE:
		return ">=", nil
	case annotations.DalOptions_Query_Arg_LTE:
		return "<=", nil
	}
	return "", errors.New("invalid comparison")
}
