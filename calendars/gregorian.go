// GREGORIAN CALENDAR
/*

Supports dates and times in the Gregorian calendar system, the current
international standard for communicating dates and times.

Supported Input Types:
  - string
  - []byte
  - time.Time (time.Duration for Offset)

Supported Format Strings:
  - any format supported by the time library or
    github.com/knz/strtime (or github.com/olebedev/when for Offset)

*/
package calendars

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-errors/errors"
	"github.com/knz/strtime"
	when "github.com/olebedev/when"
	when_common "github.com/olebedev/when/rules/common"
	when_en "github.com/olebedev/when/rules/en"
)

func init() {
	RegisterElements(
		// name
		"gregorian",
		// toInternal
		func(date interface{}, format string) (stamp TAI64NARUXTime, err error) {
			var in time.Time
			var str string

			switch date := date.(type) {
			case time.Time:
				in = date
			case string:
				str = date
			case []byte:
				str = string(date)
			default:
				err = errors.Wrap(ErrUnsupportedInput, 1)
				return
			}

			if str != "" {
				if strings.ContainsRune(format, '%') {
					in, err = strtime.Strptime(str, format)
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
		func(stamp TAI64NARUXTime, format string) (date string, err error) {
			tmp := time.Unix(TAItoUTC(stamp).Seconds, int64(TAItoUTC(stamp).Nano)).UTC()
			if strings.ContainsRune(format, '%') {
				date, err = strtime.Strftime(tmp, format)
			} else {
				date = tmp.Format(format)
			}

			return
		},
		// offset
		func(in TAI64NARUXTime, offset interface{}) (out TAI64NARUXTime, err error) {
			var str string

			switch offset := offset.(type) {
			case time.Duration:
				r := time.Unix(in.Seconds, int64(in.Nano)).UTC().Add(offset)
				out.Seconds = r.Unix()
				out.Nano = uint32(r.Nanosecond())
				return
			case string:
				str = offset
			case []byte:
				str = string(offset)
			default:
				err = errors.Wrap(ErrUnsupportedInput, 1)
			}
			if err != nil {
				return
			}

			w := when.New(nil)
			w.Add(when_en.All...)
			w.Add(when_common.All...)

			r, err := w.Parse(str, time.Unix(in.Seconds, int64(in.Nano)).UTC())
			if err != nil {
				return
			}
			if r == nil {
				err = errors.Wrap(ErrUnsupportedInput, 1)
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
