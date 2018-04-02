// JDC (JULIAN DAY COUNT) CALENDAR
/*

A count of days since BCE 4713 Jan 01 12:00:00 UTC on the proleptic Julian
Calendar. Yes, that's noon. This calendar system is used mostly for astronomy
purposes, though there is a modified variant with a narrower scope which counts
from midnight instead.

Supported Input Types:
  - string
  - []byte
  - int
  - float64
  - math/big.Float

Supported Format Strings:
  - full     - the full, canonical Day Count
  - fullday  - the full Day Count, without the fractional time part
  - fulltime - just the fractional time part of the full Day Count
  - modified - an abbreviated Day Count, 2400000.5 less than the full (starts at
               midnight instead of noon)
  - day      - the modified Day Count, without the fractional time part
  - time     - just the fractional time part of the modified Day Count

*/
package calendars

import (
	"fmt"
	"math/big"
)

var jdcBaseDay = big.NewFloat(40587.)     // Modified Julian Date at January 1, 1970
var jdcModifier = big.NewFloat(2400000.5) // Amount by which MJD is modified from JD

func init() {
	RegisterElements(
		// name
		"JDC",
		jdcToInternal,
		jdcFromInternal,
		jdcOffset,
		// defaultFormat
		"modified",
	)
}

func jdcToInternal(date interface{}, format string) (stamp TAI64NAXURTime, err error) {
	var jdc, mjd big.Float
	var jdcP, mjdP *big.Float
	var in string

	switch date.(type) {
	// TODO - other types
	case big.Float:
		tmp := date.(big.Float)
		in = tmp.String()
	case *big.Float:
		in = date.(*big.Float).String()
	case float64:
		in = fmt.Sprintf("%f", date.(float64))
	case int:
		in = fmt.Sprintf("%d", date.(int))
	case []byte:
		in = string(date.([]byte))
	case string:
		in = date.(string)
	default:
		err = ErrUnsupportedInput
		return
	}

	switch format {
	case "full", "fullday", "fulltime":
		jdcP, _, err = big.ParseFloat(in, 10, 188, big.ToNearestAway)
		jdc = *jdcP
	case "modified", "day", "time":
		mjdP, _, err = big.ParseFloat(in, 10, 196, big.ToNearestAway)
		mjd = *mjdP
		jdc.Add(mjdP, jdcModifier)
	default:
		err = ErrInvalidFormat
	}
	if err != nil {
		return
	}

	switch format {
	case "day", "fullday":
		jdcInt, _ := jdc.Sub(&jdc, big.NewFloat(0.5)).Int(nil)
		jdc.Add(jdc.SetInt(jdcInt), big.NewFloat(0.5))
	case "time":
		mjdInt, _ := mjd.Int(nil)
		mjdP.SetInt(mjdInt)
		jdc.Add(jdc.Add(mjd.Sub(&mjd, mjdP), jdcModifier), jdcBaseDay)
	case "fulltime":
		jdcInt, _ := jdc.Int(nil)
		jdc.Sub(&jdc, big.NewFloat(0.5))
		jdcP.SetInt(jdcInt)
		jdc.Add(jdc.Add(jdc.Sub(&jdc, jdcP), jdcModifier), jdcBaseDay)
	}

	stamp = TAI64NAXURTimeFromFloat(*jdc.Mul(jdc.Sub(jdc.Sub(&jdc, jdcModifier), jdcBaseDay), big.NewFloat(86400)))

	return
}

func jdcFromInternal(stamp TAI64NAXURTime, format string) (date string, err error) {
	var mjd, jdc big.Float
	timestamp := stamp.Float()
	mjd.Add(mjd.Quo(timestamp, big.NewFloat(86400)), jdcBaseDay)
	jdc.Add(&mjd, jdcModifier)

	switch format {
	case "full":
		date = fmt.Sprintf("%f", &jdc)
	case "fullday":
		date = fmt.Sprintf("%0.0f", jdc.Sub(&jdc, big.NewFloat(0.5)))
	case "fulltime":
		jdcInt, _ := jdc.Int(nil)
		date = fmt.Sprintf("%f", jdc.Sub(&jdc, mjd.SetInt(jdcInt)))
	case "modified":
		date = fmt.Sprintf("%f", &mjd)
	case "day":
		date = fmt.Sprintf("%0.0f", &mjd)
	case "time":
		mjdInt, _ := mjd.Int(nil)
		date = fmt.Sprintf("%f", mjd.Sub(&mjd, jdc.SetInt(mjdInt)))
	default:
		err = ErrInvalidFormat
	}

	return
}

func jdcOffset(in TAI64NAXURTime, offset interface{}) (out TAI64NAXURTime, err error) {
	var jdc float64
	var mod string

	switch offset.(type) {
	// TODO - other types
	case big.Float:
		tmp := offset.(big.Float)
		mod = tmp.String()
	case *big.Float:
		mod = offset.(*big.Float).String()
	case float64:
		mod = fmt.Sprintf("%f", offset.(float64))
	case int:
		mod = fmt.Sprintf("%d", offset.(int))
	case []byte:
		mod = string(offset.([]byte))
	case string:
		mod = offset.(string)
	default:
		err = ErrUnsupportedInput
		return
	}

	_, err = fmt.Sscanf(mod, "%f", &jdc)
	if err != nil {
		return
	}

	out = in.Add(TAI64NAXURTimeFromFloat(*big.NewFloat(jdc * 86400.)))

	return
}
