// Package calendars provides the underlying framework that makes Calends
// function across multiple calendar systems.
/*

It provides an interface for custom calendar systems to implement, and a data
type for storing instants in the internal TAI64NARUX format, so that the various
date/time manipulation functions of Calends can easily operate on them. It also
provides a handful of utility functions and methods to simplify the conversion
process between the calendar system and the internal format.

You generally won't use any of the types, functions, or methods defined here
unless you're implementing a calendar system class yourself.

*/
package calendars

import (
	// "fmt"
	"sort"
	"strings"

	"github.com/go-errors/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CalendarDefinition is the primary interface for defining calendar systems.
type CalendarDefinition interface {
	// Convert a date representation to an internal TAI64NARUXTime
	ToInternal(interface{}, string) (TAI64NARUXTime, error)

	// Convert an internal TAI64NARUXTime to a date representation
	FromInternal(TAI64NARUXTime, string) (string, error)

	// Calculate the TAI64NARUXTime at a given offset from another TAI64NARUXTime
	Offset(TAI64NARUXTime, interface{}) (TAI64NARUXTime, error)
}

type calendarRegistration struct {
	ToInternal    func(interface{}, string) (TAI64NARUXTime, error)
	FromInternal  func(TAI64NARUXTime, string) (string, error)
	Offset        func(TAI64NARUXTime, interface{}) (TAI64NARUXTime, error)
	DefaultFormat string
}

// Error messages returned by Calends operations.
var (
	ErrUnsupportedInput = errors.Errorf("Unsupported Value")
	ErrInvalidFormat    = errors.Errorf("Invalid Format")
)

// ErrUnknownCalendar generates a "calendar not registered" error including the
// calendar's actual name in the error message
func ErrUnknownCalendar(calendar string) error {
	return errors.Errorf("Unknown Calendar: " + calendar)
}

var registeredCalendars = make(map[string]calendarRegistration)

// RegisterObject registers a calendar system object.
/*

Registers `definition` as `name`, and saves `defaultFormat` for later use while
parsing or formatting.

*/
func RegisterObject(name string, definition CalendarDefinition, defaultFormat string) {
	RegisterElements(name, definition.ToInternal, definition.FromInternal, definition.Offset, defaultFormat)
}

// RegisterElements registers a calendar system from its distinct functions.
/*

Registers `toInternal`, `fromInternal`, and 'offset' as the elements of `name`,
and saves `defaultFormat` for later use while parsing or formatting.

*/
func RegisterElements(
	name string,
	toInternal func(interface{}, string) (TAI64NARUXTime, error),
	fromInternal func(TAI64NARUXTime, string) (string, error),
	offset func(TAI64NARUXTime, interface{}) (TAI64NARUXTime, error),
	defaultFormat string,
) {
	registeredCalendars[canonCalendarName(name)] = calendarRegistration{
		ToInternal:    toInternal,
		FromInternal:  fromInternal,
		Offset:        offset,
		DefaultFormat: defaultFormat,
	}
}

// Unregister removes a calendar system from the callback list.
func Unregister(name string) {
	if Registered(name) {
		delete(registeredCalendars, canonCalendarName(name))
	}
}

// Registered returns whether or not a calendar system has been registered, yet.
func Registered(calendar string) bool {
	_, set := registeredCalendars[canonCalendarName(calendar)]
	return set
}

// ListRegistered returns the list of calendar systems currently registered.
func ListRegistered() []string {
	var out []string

	for name := range registeredCalendars {
		out = append(out, name)
	}
	sort.Strings(out)

	return out
}

// DefaultFormat returns the associated value from a registered calendar system.
func DefaultFormat(calendar string) string {
	if !Registered(calendar) {
		return ""
	}

	return registeredCalendars[canonCalendarName(calendar)].DefaultFormat
}

// ToInternal returns the associated value from a registered calendar system.
func ToInternal(calendar string, date interface{}, format string) (TAI64NARUXTime, error) {
	if !Registered(calendar) {
		return TAI64NARUXTime{}, errors.Wrap(ErrUnknownCalendar(calendar), 1)
	}

	if format == "" {
		format = DefaultFormat(calendar)
	}

	// fmt.Printf("ToInternal: %#v (%#v) [%#v]\n", canonCalendarName(calendar), format, date)

	out, err := registeredCalendars[canonCalendarName(calendar)].ToInternal(date, format)
	if err != nil {
		return out, errors.Wrap(err, 1)
	}

	return out, nil
}

// FromInternal returns the associated value from a registered calendar system.
func FromInternal(calendar string, stamp TAI64NARUXTime, format string) (string, error) {
	if !Registered(calendar) {
		return "", errors.Wrap(ErrUnknownCalendar(calendar), 1)
	}

	if format == "" {
		format = DefaultFormat(calendar)
	}

	// fmt.Printf("FromInternal: %#v (%#v) [%#v]\n", canonCalendarName(calendar), format, stamp)

	out, err := registeredCalendars[canonCalendarName(calendar)].FromInternal(stamp, format)
	if err != nil {
		return out, errors.Wrap(err, 1)
	}

	return out, nil
}

// Offset returns the associated value from a registered calendar system.
func Offset(calendar string, stamp TAI64NARUXTime, offset interface{}) (TAI64NARUXTime, error) {
	if !Registered(calendar) {
		return TAI64NARUXTime{}, errors.Wrap(ErrUnknownCalendar(calendar), 1)
	}

	out, err := registeredCalendars[canonCalendarName(calendar)].Offset(stamp, offset)
	if err != nil {
		return out, errors.Wrap(err, 1)
	}

	return out, nil
}

func canonCalendarName(calendar string) string {
	return strings.Join(strings.Fields(cases.Title(language.Und).String(strings.ToLower(calendar))), "")
}
