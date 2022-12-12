package calends

import (
	"github.com/danhunsaker/calends/calendars"
	"github.com/go-errors/errors"
)

// Add creates a new Calends object a given offset after the current start
// point.
func (c Calends) Add(offset interface{}, calendar string) (out Calends, err error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return c, err
	}

	stamp, err := calendars.Offset(calendar, c.startTime, offset)
	if err != nil {
		return
	}

	out, err = c.SetDate(stamp, "tai64", "")

	return
}

// Subtract creates a new Calends object a given offset before the current start
// point.
func (c Calends) Subtract(offset interface{}, calendar string) (Calends, error) {
	return c.Add(negateOffset(offset), calendar)
}

// AddFromEnd creates a new Calends object a given offset after the current end
// point.
func (c Calends) AddFromEnd(offset interface{}, calendar string) (out Calends, err error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return c, err
	}

	stamp, err := calendars.Offset(calendar, c.endTime, offset)
	if err != nil {
		return
	}

	out, err = c.SetEndDate(stamp, "tai64", "")

	return
}

// SubtractFromEnd creates a new Calends object a given offset before the
// current end point.
func (c Calends) SubtractFromEnd(offset interface{}, calendar string) (Calends, error) {
	return c.AddFromEnd(negateOffset(offset), calendar)
}

// Next creates a new Calends object with a range spanning a given offset, and
// starting at the current end point.
func (c Calends) Next(offset interface{}, calendar string) (Calends, error) {
	if emptyValue(offset) {
		offset = c.duration.Text('f', 45)
		calendar = "tai64"
	}

	if calendar == "" {
		calendar = "unix"
	}

	newEndTime, err := calendars.Offset(calendar, c.endTime, offset)
	if err != nil {
		return c, err
	}

	return Create(map[string]interface{}{
		"start": c.endTime,
		"end":   newEndTime,
	}, "tai64", "")
}

// Previous creates a new Calends object with a range spanning a given offset,
// and ending at the current start point.
func (c Calends) Previous(offset interface{}, calendar string) (Calends, error) {
	if emptyValue(offset) {
		offset = c.duration.Text('f', 45)
		calendar = "tai64"
	}

	if calendar == "" {
		calendar = "unix"
	}

	newStartTime, err := calendars.Offset(calendar, c.startTime, negateOffset(offset))
	if err != nil {
		return c, err
	}

	return Create(map[string]interface{}{
		"start": newStartTime,
		"end":   c.startTime,
	}, "tai64", "")
}

// SetDate creates a new Calends object starting at a given date.
func (c Calends) SetDate(stamp interface{}, calendar, format string) (Calends, error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return c, err
	}

	if format == "" {
		format = calendars.DefaultFormat(calendar)
	}

	startTime, err := calendars.ToInternal(calendar, stamp, format)
	if err != nil {
		return c, err
	}

	return Create(map[string]interface{}{
		"start": startTime,
		"end":   c.endTime,
	}, "tai64", "")
}

// SetEndDate creates a new Calends object ending at a given date.
func (c Calends) SetEndDate(stamp interface{}, calendar, format string) (Calends, error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return c, err
	}

	if format == "" {
		format = calendars.DefaultFormat(calendar)
	}

	endTime, err := calendars.ToInternal(calendar, stamp, format)
	if err != nil {
		return c, err
	}

	return Create(map[string]interface{}{
		"start": c.startTime,
		"end":   endTime,
	}, "tai64", "")
}

// SetDuration creates a new Calends object spanning from the current start
// point for a given duration.
func (c Calends) SetDuration(duration interface{}, calendar string) (Calends, error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return c, err
	}

	offset, err := c.Add(duration, calendar)
	if err != nil {
		return c, err
	}

	return c.SetEndDate(offset.startTime, "tai64", "")
}

// SetDurationFromEnd creates a new Calends object spanning for a given duration
// to the current end point.
func (c Calends) SetDurationFromEnd(duration interface{}, calendar string) (Calends, error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return c, err
	}

	offset, err := c.SubtractFromEnd(duration, calendar)
	if err != nil {
		return c, err
	}

	return c.SetDate(offset.endTime, "tai64", "")
}

// Merge two Calends objects.
/*

Creates a new Calends object with the earlier start and later end points of the
current object and another one

*/
func (c Calends) Merge(z Calends) (Calends, error) {
	startTime := c.startTime
	if z.StartsBefore(c) {
		startTime = z.startTime
	}

	endTime := c.endTime
	if z.EndsAfter(c) {
		endTime = z.endTime
	}

	return Create(map[string]interface{}{
		"start": startTime,
		"end":   endTime,
	}, "tai64", "")
}

// Intersect gets the intersection of two Calends objects.
/*

Creates a new Calends object with the overlapping time between the current
object and another one

*/
func (c Calends) Intersect(z Calends) (Calends, error) {
	if !c.Overlaps(z) {
		return c, errors.New("Times do not overlap; no intersection exists")
	}

	startTime := c.startTime
	if z.StartsDuring(c) {
		startTime = z.startTime
	}

	endTime := c.endTime
	if z.EndsDuring(c) {
		endTime = z.endTime
	}

	return Create(map[string]interface{}{
		"start": startTime,
		"end":   endTime,
	}, "tai64", "")
}

// Gap gets the gap between two Calends objects.
/*

Creates a new Calends object with the gap in time between the current object and
another one

*/
func (c Calends) Gap(z Calends) (Calends, error) {
	if c.Overlaps(z) {
		return c, errors.New("Times overlap; no gap exists")
	}

	startTime := c.endTime
	if z.EndsBefore(c) {
		startTime = z.endTime
	}

	endTime := c.startTime
	if z.StartsAfter(c) {
		endTime = z.startTime
	}

	return Create(map[string]interface{}{
		"start": startTime,
		"end":   endTime,
	}, "tai64", "")
}

func emptyValue(in interface{}) (out bool) {
	switch in := in.(type) {
	case []byte:
		out = emptyValue(string(in))
	case string:
		out = (in == "")
	case int:
		out = false
	case float64:
		out = false
	case calendars.TAI64NAXURTime:
		out = false
	default:
		out = true
	}

	if in == nil {
		out = true
	}

	return
}

func negateOffset(in interface{}) (out interface{}) {
	switch in := in.(type) {
	case []byte:
		out = negateOffset(string(in))
	case string:
		out = "-" + in
	case int:
		out = 0 - in
	case float64:
		out = 0.0 - in
	case calendars.TAI64NAXURTime:
		out = (calendars.TAI64NAXURTime{}).Sub(in)
	default:
		out = ""
	}

	return
}
