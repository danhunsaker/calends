// TAI64 CALENDAR
/*

Supports times that are seconds since CE 1970-01-01 00:00:00 TAI (plus 2**62,
when in hexadecimal), as defined at https://cr.yp.to/libtai/tai64.html (though
this library includes extensions to the formats described there). These values
are also used internally, so this calendar system can be used to directly expose
the underlying internal values in a manner that allows them to be used
elsewhere.

Supported Input Types:
  - string
  - []byte
  - TAI64NAXURTime
  - math/big.Float for Offset

Supported Format Strings:
  - decimal     - decimal; full resolution
  - tai64       - hexadecimal; just seconds
  - tai64n      - hexadecimal; with nanoseconds
  - tai64na     - hexadecimal; with attoseconds
  - tai64nax    - hexadecimal; with xictoseconds
  - tai64naxu   - hexadecimal; with uctoseconds
  - tai64naxur  - hexadecimal; with roctoseconds

*/
package calendars

import (
	"fmt"
	"math/big"
	"strings"
)

func init() {
	RegisterElements(
		// name
		"TAI64",
		// toInternal
		func(date interface{}, format string) (stamp TAI64NAXURTime, err error) {
			var dateString string
			switch date.(type) {
			// TODO - other types
			case []byte:
				dateString = string(date.([]byte))
			case string:
				dateString = date.(string)
			case TAI64NAXURTime:
				stamp = date.(TAI64NAXURTime)
				return
			default:
				err = ErrUnsupportedInput
				return
			}

			switch format {
			case "decimal":
				tmp := strings.Split(dateString, ".")
				if len(tmp) < 2 {
					tmp = append(tmp, "0")
				}
				_, err = fmt.Sscanf(fmt.Sprintf("%s %-045s", tmp[0], tmp[1]), "%d %09d%09d%09d%09d%09d", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Xicto, &stamp.Ucto, &stamp.Rocto)
			case "tai64naxur":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X%08X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Xicto, &stamp.Ucto, &stamp.Rocto)
			case "tai64naxu":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Xicto, &stamp.Ucto)
			case "tai64nax":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Xicto)
			case "tai64na":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto)
			case "tai64n":
				_, err = fmt.Sscanf(dateString, "%016X%08X", &stamp.Seconds, &stamp.Nano)
			case "tai64":
				_, err = fmt.Sscanf(dateString, "%016X", &stamp.Seconds)
			default:
				err = ErrInvalidFormat
			}

			if err != nil && err.Error() == "EOF" {
				err = nil
			}

			if format != "decimal" && err == nil {
				stamp.Seconds -= 0x4000000000000000
			}

			return
		},
		// fromInternal
		func(stamp TAI64NAXURTime, format string) (date string, err error) {
			switch format {
			case "decimal":
				date = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%0d.%09d%09d%09d%09d%09d", stamp.Seconds, stamp.Nano, stamp.Atto, stamp.Xicto, stamp.Ucto, stamp.Rocto), "0"), ".")
			case "tai64naxur":
				date = fmt.Sprintf("%016X%08X%08X%08X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto, stamp.Xicto, stamp.Ucto, stamp.Rocto)
			case "tai64naxu":
				date = fmt.Sprintf("%016X%08X%08X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto, stamp.Xicto, stamp.Ucto)
			case "tai64nax":
				date = fmt.Sprintf("%016X%08X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto, stamp.Xicto)
			case "tai64na":
				date = fmt.Sprintf("%016X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto)
			case "tai64n":
				date = fmt.Sprintf("%016X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano)
			case "tai64":
				date = fmt.Sprintf("%016X", stamp.Seconds+0x4000000000000000)
			default:
				err = ErrInvalidFormat
			}

			return
		},
		// offset
		func(in TAI64NAXURTime, offset interface{}) (out TAI64NAXURTime, err error) {
			var adjust TAI64NAXURTime
			switch offset.(type) {
			// TODO - other types
			case big.Float:
				adjust = TAI64NAXURTimeFromFloat(offset.(big.Float))
			case *big.Float:
				adjust = TAI64NAXURTimeFromFloat(*offset.(*big.Float))
			case []byte:
				adjust = TAI64NAXURTimeFromDecimalString(string(offset.([]byte)))
			case string:
				adjust = TAI64NAXURTimeFromDecimalString(offset.(string))
			case TAI64NAXURTime:
				adjust = offset.(TAI64NAXURTime)
			default:
				err = ErrUnsupportedInput
			}

			out = in.Add(adjust)

			return
		},
		// defaultFormat
		"tai64naxur",
	)
}
