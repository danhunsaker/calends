// Package calends is a library for handling dates and times across arbitrary
// calendar systems.
/*

Dates and times are converted to "TAI64NARUX instants", values that
unambiguously encode moments over 146 billion years into the past or future, in
increments as small as 10**-45 seconds (internally called a "xindectosecond", but
there's no actual prefix for units this small, even though the Planck Time - the
smallest meaningful period of time in quantum physics - is 54 of them).
Calculations and comparisons are done using these instants, to maintain the
highest accuracy possible while working with such values. A single Calends value
can hold either a single instant, or a time span between two of them; the
duration of such a value is automatically calculated during object creation, for
easy use elsewhere. When you need a version of the instant itself (either of
them, for a span) that you can use elsewhere, it is converted back to a
date/time string in whatever calendar system you need at the time.

Several calendar systems are supported officially, and the library is easily
extended to support others without having to modify the core in any way.

*/
package calends

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/danhunsaker/calends/calendars"
	"github.com/go-errors/errors"
)

// The Calends type is the core of the library, and the primary interface with
// it.
type Calends struct {
	startTime calendars.TAI64NARUXTime
	duration  *big.Float
	endTime   calendars.TAI64NARUXTime
}

// Version of the library
var Version = "0.0.6"

// Create is the mechanism for constructing new Calends objects.
/*

It is preferred over using `make`, `new`, or `Calends{}` directly. It takes a
date/time value, a calendar name, and a format string, and returns a Calends
object representing the instant that date/time value took place. It also returns
an error object, if something goes wrong.

If the calendar passed is the empty string (`""`), Calends will use the
`"unix"` calendar automatically. If the format is the empty string, Calends
will use a default format provided by the calendar system itself.

The date/time value can be one of many types, and the exact list of types
supported varies by calendar system. At a minimum, string values should always
be supported. The documentation for each calendar system should provide more
detail on what other types are supported, and what ways the values can be
presented with each.

In any case, the date/time value can either be a single interface{} value, or a
map[string]interface{} containing two or three of them. The valid map keys are
'start', 'end', and 'duration'. Any combination of two will create the Calends
object with the associated time span. If all three are provided, the 'duration'
is ignored, and recalculated from the 'start' and 'end' values exclusively.

*/
func Create(stamp interface{}, calendar, format string) (instance Calends, err error) {
	instance.duration = big.NewFloat(0.)

	if stamp == nil {
		return
	}

	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err = errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return
	}

	if format == "" {
		format = calendars.DefaultFormat(calendar)
	}

	switch stamp := stamp.(type) {
	default:
		var internal calendars.TAI64NARUXTime

		internal, err = calendars.ToInternal(calendar, stamp, format)
		instance = Calends{
			startTime: internal,
			duration:  big.NewFloat(0.),
			endTime:   internal,
		}
	case map[string]interface{}:
		instance, err = retrieveInstance(calendar, stamp, format)
	}

	return
}

func retrieveInstance(calendar string, stamp map[string]interface{}, format string) (instance Calends, err error) {
	start, hasStart, e := retrieveStart(calendar, stamp, format)
	if e != nil {
		err = e
		return
	}
	end, hasEnd, e := retrieveEnd(calendar, stamp, format)
	if e != nil {
		err = e
		return
	}
	duration, hasDuration, e := retrieveDuration(calendar, stamp, format)
	if e != nil {
		err = e
		return
	}

	if hasStart && hasEnd {
		instance = Calends{
			startTime: start,
			duration:  end.Sub(start).Float(),
			endTime:   end,
		}
	} else if hasStart && hasDuration {
		instance = Calends{
			startTime: start,
			duration:  &duration,
			endTime:   start.Add(calendars.TAI64NARUXTimeFromFloat(duration)),
		}
	} else if hasEnd && hasDuration {
		instance = Calends{
			startTime: end.Sub(calendars.TAI64NARUXTimeFromFloat(duration)),
			duration:  &duration,
			endTime:   end,
		}
	} else {
		var internal calendars.TAI64NARUXTime

		internal, err = calendars.ToInternal(calendar, stamp, format)
		instance = Calends{
			startTime: internal,
			duration:  big.NewFloat(0.),
			endTime:   internal,
		}
	}

	return
}

func retrieveStart(calendar string, stamp map[string]interface{}, format string) (start calendars.TAI64NARUXTime, hasStart bool, err error) {
	var rawStart interface{}
	rawStart, hasStart = stamp["start"]
	if hasStart {
		start, err = calendars.ToInternal(calendar, rawStart, format)
	}

	return
}

func retrieveEnd(calendar string, stamp map[string]interface{}, format string) (end calendars.TAI64NARUXTime, hasEnd bool, err error) {
	var rawEnd interface{}
	rawEnd, hasEnd = stamp["end"]
	if hasEnd {
		end, err = calendars.ToInternal(calendar, rawEnd, format)
	}

	return
}

func retrieveDuration(calendar string, stamp map[string]interface{}, format string) (duration big.Float, hasDuration bool, err error) {
	var rawDuration interface{}
	rawDuration, hasDuration = stamp["duration"]
	if hasDuration {
		switch rawDuration := rawDuration.(type) {
		case []byte:
			var tmp *big.Float
			tmp, _, err = big.ParseFloat(string(rawDuration), 10, 176, big.ToNearestAway)
			if err == nil {
				duration = *tmp
			} else if err.Error() == "EOF" {
				duration = *big.NewFloat(0)
				err = nil
			}
		case string:
			var tmp *big.Float
			tmp, _, err = big.ParseFloat(rawDuration, 10, 176, big.ToNearestAway)
			if err == nil {
				duration = *tmp
			} else if err.Error() == "EOF" {
				duration = *big.NewFloat(0)
				err = nil
			}
		case int:
			duration = *big.NewFloat(float64(rawDuration))
		case float64:
			duration = *big.NewFloat(rawDuration)
		case *big.Float:
			tmp := rawDuration
			duration = *tmp
		case big.Float:
			duration = rawDuration
		default:
			err = errors.New("Invalid Duration Type")
		}
	}

	return
}

// Date is used to retrieve the value of an instant in a given calendar system.
func (c Calends) Date(calendar, format string) (string, error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return "", err
	}

	if format == "" {
		format = calendars.DefaultFormat(calendar)
	}

	return calendars.FromInternal(calendar, c.startTime, format)
}

// Duration retrieves the number of seconds between the start and end instants.
func (c Calends) Duration() *big.Float {
	return c.duration
}

// EndDate retrieves the value of the end instant in a given calendar system.
func (c Calends) EndDate(calendar, format string) (string, error) {
	if calendar == "" {
		calendar = "unix"
	}
	if !calendars.Registered(calendar) {
		err := errors.Wrap(calendars.ErrUnknownCalendar(calendar), 1)
		return "", err
	}

	if format == "" {
		format = calendars.DefaultFormat(calendar)
	}

	return calendars.FromInternal(calendar, c.endTime, format)
}

// String implements the fmt.Stringer interface.
func (c Calends) String() string {
	if tmp, _ := c.duration.Int64(); tmp == 0 {
		out, _ := c.startTime.MarshalText()
		return string(out)
	}

	start, _ := c.startTime.MarshalText()
	end, _ := c.endTime.MarshalText()

	return fmt.Sprintf("from %s to %s", start, end)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (c Calends) MarshalText() ([]byte, error) {
	if tmp, _ := c.duration.Int64(); tmp == 0 {
		return c.startTime.MarshalText()
	}

	start, err := c.startTime.MarshalText()
	if err != nil {
		return []byte{}, err
	}
	end, err := c.endTime.MarshalText()
	if err != nil {
		return []byte{}, err
	}

	return []byte(fmt.Sprintf("%s::%s", start, end)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (c *Calends) UnmarshalText(text []byte) error {
	var startTime, endTime calendars.TAI64NARUXTime
	var start, end string

	n, err := fmt.Sscanf(string(text), "%56s::%56s", &start, &end)
	if err != nil && err.Error() != io.EOF.Error() && err.Error() != "input does not match format" && err.Error() != "unexpected EOF" {
		return err
	}

	if n < 2 {
		end = start
	}

	startTime.UnmarshalText([]byte(start))
	endTime.UnmarshalText([]byte(end))

	tmp, err := Create(map[string]interface{}{
		"start": startTime,
		"end":   endTime,
	}, "tai64", "tai64narux")

	*c = tmp

	return err
}

// MarshalJSON implements the encoding/json.Marshaler interface.
func (c Calends) MarshalJSON() ([]byte, error) {
	if tmp, _ := c.duration.Int64(); tmp == 0 {
		tmp, err := c.startTime.MarshalText()
		return append(append([]byte{'"'}, tmp...), '"'), err
	}

	start, err := c.startTime.MarshalText()
	if err != nil {
		return []byte{}, err
	}
	end, err := c.endTime.MarshalText()
	if err != nil {
		return []byte{}, err
	}

	return []byte(fmt.Sprintf(`{"start":"%s","end":"%s"}`, start, end)), nil
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (c *Calends) UnmarshalJSON(text []byte) error {
	var startTime, endTime calendars.TAI64NARUXTime

	parsed := make(map[string]string)
	err := json.Unmarshal(text, &parsed)
	if err == nil {
		err = startTime.UnmarshalText([]byte(parsed["start"]))
		if err != nil {
			return errors.New("JSON decode failure while parsing start time [" + parsed["start"] + "]: " + err.Error())
		}

		err = endTime.UnmarshalText([]byte(parsed["end"]))
		if err != nil {
			return errors.New("JSON decode failure while parsing end time [" + parsed["end"] + "]: " + err.Error())
		}
	} else {
		err = startTime.UnmarshalText([]byte(strings.Trim(string(text), `"`)))
		if err != nil {
			return errors.New("JSON decode failure while parsing time [" + strings.Trim(string(text), `"`) + "]: " + err.Error())
		}

		endTime = startTime
	}

	temp, err := Create(map[string]interface{}{
		"start": startTime,
		"end":   endTime,
	}, "tai64", "")

	*c = temp

	if err != nil {
		err = errors.New("JSON decode failure while setting values: " + err.Error())
	}

	return err
}
