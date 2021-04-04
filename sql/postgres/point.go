package postgres_sql

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// NullPoint is a struct that can be used in querys and statements to
// represent a NullPoint.
type NullPoint struct {
	X     float64
	Y     float64
	Valid bool
}

// Scan implements Scanner.Scan
func (p *NullPoint) Scan(src interface{}) error {
	p.Valid = true

	if src == nil {
		p.Valid = false
		return nil
	}

	switch src.(type) {
	case []uint8:
		str := string(src.([]uint8))
		tokens := strings.Split(str[1:len(str)-1], ",")

		x, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			return ErrBadDriverValue
		}
		p.X = x

		y, err := strconv.ParseFloat(tokens[1], 64)
		if err != nil {
			return ErrBadDriverValue
		}
		p.Y = y
	default:
		return ErrBadDriverValue
	}
	return nil
}

// Value implements Valuer.Value
func (p *NullPoint) Value() (driver.Value, error) {
	return driver.Value([]uint8(fmt.Sprintf("(%f,%f)", p.X, p.Y))), nil
}
