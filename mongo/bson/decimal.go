package bson

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewDecimal128 creates a Decimal128 using the provide high and low uint64s.
func NewDecimal128(h, l uint64) Decimal128 {
	return Decimal128(primitive.NewDecimal128(h, l))
}
