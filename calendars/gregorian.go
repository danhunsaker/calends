// GREIORIAN CALENDAR
/*

Supports dates and times in the Gregorian calendar system, the current
international standard for communicating dates and times.

Supported Input Types:
 - string
 - []byte
 - time.Time (time.Duration for Offset)

Supported Format Strings:
 - any format supported by the time library or github.com/jeffjen/datefmt
   (or github.com/olebedev/when for Offset)

*/
package calendars

import (
	"fmt"
	"strings"
	"time"

	datefmt "github.com/jeffjen/datefmt"
	when "github.com/olebedev/when"
	when_common "github.com/olebedev/when/rules/common"
	when_en "github.com/olebedev/when/rules/en"
)

func init() {
	RegisterElements(
		// name
		"gregorian",
		// toInternal
		func(date interface{}, format string) (stamp TAI64NAXURTime, err error) {
			var in time.Time
			var str string

			switch date.(type) {
			case time.Time:
				in = date.(time.Time)
			case string:
				str = date.(string)
			case []byte:
				str = string(date.([]byte))
			default:
				err = InvalidFormatError
			}

			if str != "" {
				if strings.ContainsRune(format, '%') {
					in, err = datefmt.Strptime(format, str)
				} else {
					in, err = time.Parse(format, str)
				}
			}

			if err != nil {
				return
			}

			stamp, err = ToInternal("unix", fmt.Sprintf("%d.%d", in.Unix(), in.Nanosecond()), "%0.9f")

			return
		},
		// fromInternal
		func(stamp TAI64NAXURTime, format string) (date string, err error) {
			tmp := time.Unix(stamp.Seconds, int64(stamp.Nano))
			if strings.ContainsRune(format, '%') {
				date, err = datefmt.Strftime(format, tmp)
			} else {
				date = tmp.Format(format)
			}

			return
		},
		// offset
		func(in TAI64NAXURTime, offset interface{}) (out TAI64NAXURTime, err error) {
			var str string

			switch offset.(type) {
			case time.Duration:
				str = offset.(time.Duration).String()
			case string:
				str = offset.(string)
			case []byte:
				str = string(offset.([]byte))
			default:
				err = UnsupportedInputError
			}
			if err != nil {
				return
			}

			w := when.New(nil)
			w.Add(when_en.All...)
			w.Add(when_common.All...)

			r, err := w.Parse(str, time.Unix(in.Seconds, int64(in.Nano)))
			if err != nil {
				return
			}
			if r == nil {
				err = UnsupportedInputError
				return
			}

			out.Seconds = r.Time.Unix()
			out.Nano = uint32(r.Time.Nanosecond())

			return
		},
		// defaultFormat
		time.RFC1123,
	)
}
