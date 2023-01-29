// UNIX CALENDAR
/*

Supports times that are seconds since CE 1970-01-01 00:00:00 UTC, commonly used
by computer systems for storing date/time values, internally.

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

	"github.com/go-errors/errors"
)

func init() {
	RegisterElements(
		// name
		"unix",
		// toInternal
		func(date interface{}, format string) (stamp TAI64NARUXTime, err error) {
			var tmp TAI64NARUXTime
			// parse the value
			tmp, err = unixParseDate(date)
			if err != nil {
				return
			}

			stamp = UTCtoTAI(tmp)
			return
		},
		// fromInternal
		func(stamp TAI64NARUXTime, format string) (date string, err error) {
			stamp = TAItoUTC(stamp)

			// format the value
			date = fmt.Sprintf(format, stamp.Float())
			return
		},
		// offset
		func(in TAI64NARUXTime, offset interface{}) (out TAI64NARUXTime, err error) {
			var tmp1, tmp2 TAI64NARUXTime

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

func unixParseDate(date interface{}) (stamp TAI64NARUXTime, err error) {
	switch date := date.(type) {
	// TODO - other types
	case int:
		stamp = TAI64NARUXTimeFromDecimalString(fmt.Sprintf("%d", date))
	case float64:
		stamp = TAI64NARUXTimeFromDecimalString(fmt.Sprintf("%f", date))
	case big.Float:
		stamp = TAI64NARUXTimeFromFloat(date)
	case *big.Float:
		stamp = TAI64NARUXTimeFromFloat(*date)
	case []byte:
		stamp = TAI64NARUXTimeFromDecimalString(string(date))
	case string:
		stamp = TAI64NARUXTimeFromDecimalString(date)
	default:
		err = errors.Wrap(ErrUnsupportedInput, 1)
	}

	return
}
