package postgres

import (
	"errors"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
)

// Arg wraps a Arg so additional logic can be applied.
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

		if f := fopts.GetPostgres().GetField(); f != "" {
			return f, nil
		}

		if f := fopts.GetField(); f != "" {
			return f, nil
		}
	}

	if a.raw.GetPostgres().GetName() != "" {
		return a.raw.GetPostgres().GetName(), nil
	}

	return a.Arg.QueryName()
}

// Comparison returns the check to perform
func (a *Arg) Comparison() (string, error) {
	switch a.raw.GetComparison() {
	case annotations.Comparator_EQ:
		return "=", nil
	case annotations.Comparator_NE:
		return "!=", nil
	case annotations.Comparator_GT:
		return ">", nil
	case annotations.Comparator_LT:
		return "<", nil
	case annotations.Comparator_GTE:
		return ">=", nil
	case annotations.Comparator_LTE:
		return "<=", nil
	}
	return "", errors.New("invalid comparison")
}
