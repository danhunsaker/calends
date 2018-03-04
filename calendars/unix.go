// UNIX CALENDAR
/*

Supports times that are seconds since 1970-01-01 00:00:00 UTC, commonly used by
computer systems for storing date/time values, internally.

Supported Input Types:
  - string
  - []byte
  - math.big.Float

Supported Format Strings:
  - any format supported by math.big.Float

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
			// parse the value
			switch date.(type) {
			// TODO - other types
			case int:
				stamp = TAI64NAXURTimeFromDecimalString(string(date.(int)))
			case float64:
				stamp = TAI64NAXURTimeFromDecimalString(fmt.Sprintf("%f", date.(float64)))
			case big.Float:
				stamp = TAI64NAXURTimeFromFloat(date.(big.Float))
			case []byte:
				stamp = TAI64NAXURTimeFromDecimalString(string(date.([]byte)))
			case string:
				stamp = TAI64NAXURTimeFromDecimalString(date.(string))
			default:
				err = UnsupportedInputError
			}
			if err != nil {
				return
			}

			stamp, err = UTCtoTAI(stamp)
			return
		},
		// fromInternal
		func(stamp TAI64NAXURTime, format string) (date string, err error) {
			stamp, err = TAItoUTC(stamp)
			if err != nil {
				return
			}

			// format the value
			date = fmt.Sprintf(format, stamp.Float())
			return
		},
		// offset
		func(in TAI64NAXURTime, offset interface{}) (out TAI64NAXURTime, err error) {
			out, err = TAItoUTC(in)
			if err != nil {
				return
			}

			// do the math
			tmp, err := ToInternal("unix", offset, "")
			if err != nil {
				return
			}

			out, err = UTCtoTAI(out.Add(tmp))
			return
		},
		// defaultFormat
		"%0.9f",
	)
}
