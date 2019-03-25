// DYNAMIC CALENDAR
/*

Supports custom, user-defined calendar systems. Each calendar system is
specified according to its components and rules for construction. More details
can be found in the `dynamic` sub-package.

Supported Input Types:
  - string
  - []byte
  - int
  - float64
  - math/big.Float

Supported Format Strings:
  - varies by actual (user-defined) calendar system

*/
package calendars

import (
	"github.com/danhunsaker/calends/calendars/dynamic"
)

type dynamicCalendar struct {
	info dynamic.Calendar
}

func init() {
	for _, info := range loadDynamicCalendars() {
		RegisterDynamic(info)
	}
}

func loadDynamicCalendars() (list []dynamic.Calendar) {
	// need to explore DB options...
	// use dynamic.Calendar.UnmarshalJSON() for this?
	return
}

func RegisterDynamic(info dynamic.Calendar) {
	RegisterObject(
		info.Name,
		dynamicCalendar{info},
		info.DefaultFormat.String,
	)
}

func (self dynamicCalendar) ToInternal(date interface{}, format string) (stamp TAI64NAXURTime, err error) {
	time, err := self.info.ToTimestamp(date, format)
	if err != nil {
		return
	}

	stamp = TAI64NAXURTimeFromFloat(*time)
	return
}

func (self dynamicCalendar) FromInternal(stamp TAI64NAXURTime, format string) (string, error) {
	return self.info.FromTimestamp(stamp.Float(), format)
}

func (self dynamicCalendar) Offset(in TAI64NAXURTime, offset interface{}) (out TAI64NAXURTime, err error) {
	time, err := self.info.Offset(in.Float(), offset)
	if err != nil {
		return
	}

	out = TAI64NAXURTimeFromFloat(*time)
	return
}
