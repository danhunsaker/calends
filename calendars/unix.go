// UNIX CALENDAR
/*

Supports times that are seconds since 1970-01-01 00:00:00 UTC, commonly used by
computer systems for storing date/time values, internally.

Supported Input Types:
  - string
  - []byte
	- int
	- float64
  - math/big.Float

Supported Format Strings:
  - any format supported by math/big.Float

*/
package calendars

import (
	"fmt"
	"math/big"
)

func init() {
	RegisterElements(
		// name
		"unix",
		// toInternal
		func(date interface{}, format string) (stamp TAI64NAXURTime, err error) {
			var tmp TAI64NAXURTime
			// parse the value
			tmp, err = unixParseDate(date)
			if err != nil {
				return
			}

			stamp = UTCtoTAI(tmp)
			return
		},
		// fromInternal
		func(stamp TAI64NAXURTime, format string) (date string, err error) {
			stamp = TAItoUTC(stamp)

			// format the value
			date = fmt.Sprintf(format, stamp.Float())
			return
		},
		// offset
		func(in TAI64NAXURTime, offset interface{}) (out TAI64NAXURTime, err error) {
			var tmp1, tmp2 TAI64NAXURTime

			tmp1 = TAItoUTC(in)

			// do the math
			tmp2, err = unixParseDate(offset)
			if err != nil {
				return
			}

			out = UTCtoTAI(tmp1.Add(tmp2))
			return
		},
		// defaultFormat
		"%0.9f",
	)
}

func unixParseDate(date interface{}) (stamp TAI64NAXURTime, err error) {
	switch date.(type) {
	// TODO - other types
	case int:
		stamp = TAI64NAXURTimeFromDecimalString(fmt.Sprintf("%d", date.(int)))
	case float64:
		stamp = TAI64NAXURTimeFromDecimalString(fmt.Sprintf("%f", date.(float64)))
	case big.Float:
		stamp = TAI64NAXURTimeFromFloat(date.(big.Float))
	case *big.Float:
		stamp = TAI64NAXURTimeFromFloat(*date.(*big.Float))
	case []byte:
		stamp = TAI64NAXURTimeFromDecimalString(string(date.([]byte)))
	case string:
		stamp = TAI64NAXURTimeFromDecimalString(date.(string))
	default:
		err = ErrUnsupportedInput
	}

	return
}
