package calends

import (
	"errors"

	"github.com/danhunsaker/calends/calendars"
)

// Add creates a new Calends object a given offset after the current start
// point.
func (c Calends) Add(offset, calendar string) (out Calends, err error) {
	stamp, err := calendars.Offset(calendar, c.startTime, offset)
	if err != nil {
		return
	}

	out, err = c.SetDate(stamp, "tai64", "")

	return
}

// Subtract creates a new Calends object a given offset before the current start
// point.
func (c Calends) Subtract(offset, calendar string) (Calends, error) {
	return c.Add("-"+offset, calendar)
}

// AddFromEnd creates a new Calends object a given offset after the current end
// point.
func (c Calends) AddFromEnd(offset, calendar string) (out Calends, err error) {
	stamp, err := calendars.Offset(calendar, c.endTime, offset)
	if err != nil {
		return
	}

	out, err = c.SetEndDate(stamp, "tai64", "")

	return
}

// SubtractFromEnd creates a new Calends object a given offset before the
// current end point.
func (c Calends) SubtractFromEnd(offset, calendar string) (Calends, error) {
	return c.AddFromEnd("-"+offset, calendar)
}

// Next creates a new Calends object with a range spanning a given offset, and
// starting at the current end point.
func (c Calends) Next(offset, calendar string) (Calends, error) {
	if offset == "" {
		offset = c.duration.Text('f', 45)
		calendar = "tai64"
	}

	if calendar == "" {
		calendar = "tai64"
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
func (c Calends) Previous(offset, calendar string) (Calends, error) {
	if offset == "" {
		offset = c.duration.Text('f', 45)
		calendar = "tai64"
	}

	if calendar == "" {
		calendar = "tai64"
	}

	newStartTime, err := calendars.Offset(calendar, c.startTime, "-"+offset)
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
		err := calendars.UnknownCalendarError
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
		err := calendars.UnknownCalendarError
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
func (c Calends) SetDuration(duration, calendar string) (Calends, error) {
	offset, err := c.Add(duration, calendar)
	if err != nil {
		return c, err
	}

	return c.SetEndDate(offset.startTime, "tai64", "")
}

// SetDurationFromEnd creates a new Calends object spanning for a given duration
// to the current end point.
func (c Calends) SetDurationFromEnd(duration, calendar string) (Calends, error) {
	offset, err := c.SubtractFromEnd(duration, calendar)
	if err != nil {
		return c, err
	}

	return c.SetDate(offset.endTime, "tai64", "")
}

// Merges two Calends objects.
/*

Creates a new Calends object with the earlier start and later end points of the
current object and another one

*/
func (a Calends) Merge(b Calends) (Calends, error) {
	startTime := a.startTime
	if b.StartsBefore(a) {
		startTime = b.startTime
	}

	endTime := a.endTime
	if b.EndsAfter(a) {
		endTime = b.endTime
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
func (a Calends) Intersect(b Calends) (Calends, error) {
	if !a.Overlaps(b) {
		return a, errors.New("Times do not overlap; no intersection exists")
	}

	startTime := a.startTime
	if b.StartsDuring(a) {
		startTime = b.startTime
	}

	endTime := a.endTime
	if b.EndsDuring(a) {
		endTime = b.endTime
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
func (a Calends) Gap(b Calends) (Calends, error) {
	if a.Overlaps(b) {
		return a, errors.New("Times overlap; no gap exists")
	}

	startTime := a.endTime
	if b.EndsBefore(a) {
		startTime = b.endTime
	}

	endTime := a.startTime
	if b.StartsAfter(a) {
		endTime = b.startTime
	}

	return Create(map[string]interface{}{
		"start": startTime,
		"end":   endTime,
	}, "tai64", "")
}
