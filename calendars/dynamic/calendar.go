package dynamic

import (
	"errors"
	"fmt"
	"math/big"
	// "strings"
)

type Calendar struct {
	Name          string
	Description   string
	Units         []*Unit
	BaseUnit      *Unit
	Formats       []*Format
	DefaultFormat *Format
	Fragments     []*Fragment
}

func (c *Calendar) ToTimestamp(date interface{}, format string) (stamp *big.Float, err error) {
	var dateString string

	switch date.(type) {
	case big.Float:
		tmp := date.(big.Float)
		dateString = tmp.String()
	case *big.Float:
		dateString = date.(*big.Float).String()
	case float64:
		dateString = fmt.Sprintf("%f", date.(float64))
	case int:
		dateString = fmt.Sprintf("%d", date.(int))
	case string:
		dateString = date.(string)
	case []byte:
		dateString = string(date.([]byte))
	default:
		err = errors.New("Unsupported Value")
		return
	}

	stamp = c.unitsToTime(c.dateToUnits(dateString, formatFromString(c, format)))

	return
}

func (c *Calendar) FromTimestamp(stamp *big.Float, format string) (string, error) {
	return c.unitsToDate(c.timeToUnits(stamp), formatFromString(c, format)), nil
}

func (c *Calendar) Offset(in *big.Float, offset interface{}) (out *big.Float, err error) {
	var offsetString string

	switch offset.(type) {
	case big.Float:
		tmp := offset.(big.Float)
		offsetString = tmp.String()
	case *big.Float:
		offsetString = offset.(*big.Float).String()
	case float64:
		offsetString = fmt.Sprintf("%f", offset.(float64))
	case int:
		offsetString = fmt.Sprintf("%d", offset.(int))
	case string:
		offsetString = offset.(string)
	case []byte:
		offsetString = string(offset.([]byte))
	default:
		err = errors.New("Unsupported Value")
		return
	}

	out = c.unitsToTime(c.unitsWithOffset(c.timeToUnits(in), offsetString))

	return
}

func (c *Calendar) dateToUnits(in string, format *Format) (out map[string]int) {
	formats := append([]*Format{format, c.DefaultFormat}, c.Formats...)

	for _, f := range formats {
		out, err := f.parse(in)
		if err == nil {
			return out
		}
	}

	return c.epochUnits(true)
}

func (c *Calendar) unitsToDate(in map[string]int, format *Format) (out string) {
	return
}

func (c *Calendar) unitsWithOffset(in map[string]int, offset string) (out map[string]int) {
	return
}

func (c *Calendar) unitsToTime(in map[string]int) *big.Float {
	return c.BaseUnit.toSeconds(c.sumUnits(c.epochUnits(false), in))
}

func (c *Calendar) timeToUnits(in *big.Float) map[string]int {
	return c.sumUnits(c.epochUnits(true), c.BaseUnit.fromSeconds(in))
}

func (c *Calendar) sumUnits(a, b map[string]int) (out map[string]int) {
	for aKey, aVal := range a {
		out[aKey] = aVal

		if bVal, ok := b[aKey]; ok {
			out[aKey] = aVal + bVal
		}
	}

	for bKey, bVal := range b {
		if _, ok := out[bKey]; !ok {
			out[bKey] = bVal
		}
	}

	out = c.BaseUnit.carryOver(out)

	return
}

func (c *Calendar) epochUnits(positive bool) (out map[string]int) {
	for _, unit := range c.Units {
		if unit.IsAuxiliary {
			continue
		}

		if positive {
			out[unit.InternalName] = unit.UnixEpoch
		} else {
			zero := 1
			if unit.UsesZero {
				zero = 0
			}

			out[unit.InternalName] = -1 * (unit.UnixEpoch - zero)
		}
	}

	return
}
