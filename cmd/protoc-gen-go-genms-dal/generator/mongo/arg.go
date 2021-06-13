package mongo

import (
	"errors"

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

// ToMongo indicates what type the field value should be converted to.
func (a *Arg) ToMongo() string {
	if a.field == nil {
		return a.raw.Kind
	}

	return a.field.ToMongo()
}

// ToGo indicates what type the field value should be converted to.
func (a *Arg) ToGo() (string, error) {
	return a.field.ToGo()
}

// QueryName returns the name of the field in the remote system.
func (a *Arg) QueryName() (string, error) {
	if a.field != nil {
		fopts := a.field.Options()

		if f := fopts.GetMongo().GetName(); f != "" {
			return f, nil
		}

		if f := fopts.GetName(); f != "" {
			return f, nil
		}
	}

	if n := a.raw.GetMongo().GetName(); n != "" {
		return n, nil
	}

	return a.Arg.QueryName()
}

// Comparison returns the check to perform
func (a *Arg) Comparison() (string, error) {
	switch a.raw.GetComparison() {
	case annotations.Comparator_EQ:
		return "$eq", nil
	case annotations.Comparator_NE:
		return "$ne", nil
	case annotations.Comparator_GT:
		return "$gt", nil
	case annotations.Comparator_LT:
		return "$lt", nil
	case annotations.Comparator_GTE:
		return "$gte", nil
	case annotations.Comparator_LTE:
		return "$lte", nil
	}
	return "", errors.New("invalid comparison")
}
