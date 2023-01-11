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
  - TAI64NARUXTime
  - math/big.Float for Offset

Supported Format Strings:
  - decimal     - decimal; full resolution
  - tai64       - hexadecimal; just seconds
  - tai64n      - hexadecimal; with nanoseconds
  - tai64na     - hexadecimal; with attoseconds
  - tai64nar    - hexadecimal; with rontoseconds
  - tai64naru   - hexadecimal; with udectoseconds
  - tai64narux  - hexadecimal; with xindectoseconds

*/
package calendars

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/go-errors/errors"
)

func init() {
	RegisterElements(
		// name
		"TAI64",
		// toInternal
		func(date interface{}, format string) (stamp TAI64NARUXTime, err error) {
			var dateString string
			switch date.(type) {
			// TODO - other types
			case []byte:
				dateString = string(date.([]byte))
			case string:
				dateString = date.(string)
			case TAI64NARUXTime:
				stamp = date.(TAI64NARUXTime)
				return
			default:
				err = errors.Wrap(ErrUnsupportedInput, 1)
				return
			}

			switch format {
			case "decimal":
				tmp := strings.Split(dateString, ".")
				if len(tmp) < 2 {
					tmp = append(tmp, "0")
				}
				_, err = fmt.Sscanf(fmt.Sprintf("%s %-045s", tmp[0], tmp[1]), "%d %09d%09d%09d%09d%09d", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Ronto, &stamp.Udecto, &stamp.Xindecto)
			case "tai64narux":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X%08X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Ronto, &stamp.Udecto, &stamp.Xindecto)
			case "tai64naru":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Ronto, &stamp.Udecto)
			case "tai64nar":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto, &stamp.Ronto)
			case "tai64na":
				_, err = fmt.Sscanf(dateString, "%016X%08X%08X", &stamp.Seconds, &stamp.Nano, &stamp.Atto)
			case "tai64n":
				_, err = fmt.Sscanf(dateString, "%016X%08X", &stamp.Seconds, &stamp.Nano)
			case "tai64":
				_, err = fmt.Sscanf(dateString, "%016X", &stamp.Seconds)
			default:
				err = errors.Wrap(ErrInvalidFormat, 1)
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
		func(stamp TAI64NARUXTime, format string) (date string, err error) {
			switch format {
			case "decimal":
				date = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%0d.%09d%09d%09d%09d%09d", stamp.Seconds, stamp.Nano, stamp.Atto, stamp.Ronto, stamp.Udecto, stamp.Xindecto), "0"), ".")
			case "tai64narux":
				date = fmt.Sprintf("%016X%08X%08X%08X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto, stamp.Ronto, stamp.Udecto, stamp.Xindecto)
			case "tai64naru":
				date = fmt.Sprintf("%016X%08X%08X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto, stamp.Ronto, stamp.Udecto)
			case "tai64nar":
				date = fmt.Sprintf("%016X%08X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto, stamp.Ronto)
			case "tai64na":
				date = fmt.Sprintf("%016X%08X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano, stamp.Atto)
			case "tai64n":
				date = fmt.Sprintf("%016X%08X", stamp.Seconds+0x4000000000000000, stamp.Nano)
			case "tai64":
				date = fmt.Sprintf("%016X", stamp.Seconds+0x4000000000000000)
			default:
				err = errors.Wrap(ErrInvalidFormat, 1)
			}

			return
		},
		// offset
		func(in TAI64NARUXTime, offset interface{}) (out TAI64NARUXTime, err error) {
			var adjust TAI64NARUXTime
			switch offset.(type) {
			// TODO - other types
			case big.Float:
				adjust = TAI64NARUXTimeFromFloat(offset.(big.Float))
			case *big.Float:
				adjust = TAI64NARUXTimeFromFloat(*offset.(*big.Float))
			case []byte:
				adjust = TAI64NARUXTimeFromDecimalString(string(offset.([]byte)))
			case string:
				adjust = TAI64NARUXTimeFromDecimalString(offset.(string))
			case TAI64NARUXTime:
				adjust = offset.(TAI64NARUXTime)
			default:
				err = errors.Wrap(ErrUnsupportedInput, 1)
			}

			out = in.Add(adjust)

			return
		},
		// defaultFormat
		"tai64narux",
	)
}
