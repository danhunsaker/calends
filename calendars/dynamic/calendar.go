package dynamic

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
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

func (self *Calendar) ToTimestamp(date interface{}, format string) (stamp *big.Float, err error) {
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

	stamp = self.unitsToTime(self.dateToUnits(dateString, formatFromString(self, format)))

	return
}

func (self *Calendar) FromTimestamp(stamp *big.Float, format string) (string, error) {
	return self.unitsToDate(self.timeToUnits(stamp), formatFromString(self, format)), nil
}

func (self *Calendar) Offset(in *big.Float, offset interface{}) (out *big.Float, err error) {
	var offsetString string
	var units map[string]int

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

	units, err = self.unitsWithOffset(self.timeToUnits(in), offsetString)
	if err != nil {
		out = self.unitsToTime(units)
	}

	return
}

func (self *Calendar) dateToUnits(in string, format *Format) (out map[string]int) {
	formats := append([]*Format{format, self.DefaultFormat}, self.Formats...)

	for _, f := range formats {
		out, err := f.parse(in)
		if err == nil {
			return out
		}
	}

	return self.epochUnits(true)
}

func (self *Calendar) unitsToDate(in map[string]int, format *Format) (out string) {
	return format.format(in)
}

func (self *Calendar) unitsWithOffset(in map[string]int, offset string) (out map[string]int, err error) {
	var names map[string]int
	var matchUnit *Unit
	var matchValue int64

	re := regexp.MustCompile("(?P<value>[-+]?[0-9]+)\\s*(?P<unit>\\S+)")
	for idx, name := range re.SubexpNames() {
		if name == "" {
			continue
		}
		names[name] = idx
	}

	for _, match := range re.FindAllStringSubmatch(offset, -1) {
		matchUnit = nil
	FindUnit:
		for _, unit := range self.Units {
			if unit.InternalName == match[names["unit"]] {
				matchUnit = unit
				break FindUnit
			}

			for _, name := range unit.Names {
				if name.UnitName == match[names["unit"]] {
					matchUnit = unit
					break FindUnit
				}
			}
		}
		if matchUnit == nil {
			continue
		}

		matchValue, err = strconv.ParseInt(match[names["value"]], 0, 0)
		if err != nil {
			return
		}

		unitName, unitValue := matchUnit.reduceAuxiliary(int(matchValue))
		if currentValue, ok := out[unitName]; ok {
			unitValue += currentValue
		}

		out[unitName] = unitValue
	}

	out = self.BaseUnit.carryOver(out)

	return
}

func (self *Calendar) unitsToTime(in map[string]int) *big.Float {
	return self.BaseUnit.toSeconds(self.sumUnits(self.epochUnits(false), in))
}

func (self *Calendar) timeToUnits(in *big.Float) map[string]int {
	return self.sumUnits(self.epochUnits(true), self.BaseUnit.fromSeconds(in))
}

func (self *Calendar) sumUnits(a, b map[string]int) (out map[string]int) {
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

	out = self.BaseUnit.carryOver(out)

	return
}

func (self *Calendar) epochUnits(positive bool) (out map[string]int) {
	for _, unit := range self.Units {
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
