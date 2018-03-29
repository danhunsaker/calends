// Package calendars provides the underlying framework that makes Calends
// function across multiple calendar systems.
/*

It provides an interface for custom calendar systems to implement, and a data
type for storing instants in the internal TAI64NAXUR format, so that the various
date/time manipulation functions of Calends can easily operate on them. It also
provides a handful of utility functions and methods to simplify the conversion
process between the calendar system and the internal format.

You generally won't use any of the types, functions, or methods defined here
unless you're implementing a calendar system class yourself.

*/
package calendars

import (
	"errors"
	// "fmt"
	"strings"
)

// CalendarDefinition is the primary interface for defining calendar systems.
type CalendarDefinition interface {
	// Convert a date representation to an internal TAI64NAXURTime
	ToInternal(interface{}, string) (TAI64NAXURTime, error)

	// Convert an internal TAI64NAXURTime to a date representation
	FromInternal(TAI64NAXURTime, string) (string, error)

	// Calculate the TAI64NAXURTime at a given offset from another TAI64NAXURTime
	Offset(TAI64NAXURTime, interface{}) (TAI64NAXURTime, error)
}

type calendarRegistration struct {
	ToInternal    func(interface{}, string) (TAI64NAXURTime, error)
	FromInternal  func(TAI64NAXURTime, string) (string, error)
	Offset        func(TAI64NAXURTime, interface{}) (TAI64NAXURTime, error)
	DefaultFormat string
}

// Error messages returned by Calends operations.
var (
	ErrUnknownCalendar  = errors.New("Unknown Calendar")
	ErrUnsupportedInput = errors.New("Unsupported Value")
	ErrInvalidFormat    = errors.New("Invalid Format")
)

var registeredCalendars = make(map[string]calendarRegistration)

// RegisterClass registers a calendar system class.
/*

Registers `definition` as `name`, and saves `defaultFormat` for later use while
parsing or formatting.

*/
func RegisterClass(name string, definition CalendarDefinition, defaultFormat string) {
	RegisterElements(name, definition.ToInternal, definition.FromInternal, definition.Offset, defaultFormat)
}

// RegisterElements registers a calendar system from its distinct functions.
/*

Registers `toInternal`, `fromInternal`, and 'offset' as the elements of `name`,
and saves `defaultFormat` for later use while parsing or formatting.

*/
func RegisterElements(
	name string,
	toInternal func(interface{}, string) (TAI64NAXURTime, error),
	fromInternal func(TAI64NAXURTime, string) (string, error),
	offset func(TAI64NAXURTime, interface{}) (TAI64NAXURTime, error),
	defaultFormat string,
) {
	registeredCalendars[canonCalendarName(name)] = calendarRegistration{
		ToInternal:    toInternal,
		FromInternal:  fromInternal,
		Offset:        offset,
		DefaultFormat: defaultFormat,
	}
}

// Registered returns whether or not a calendar system has been registered, yet.
func Registered(calendar string) bool {
	_, set := registeredCalendars[canonCalendarName(calendar)]
	return set
}

// DefaultFormat returns the associated value from a registered calendar system.
func DefaultFormat(calendar string) string {
	if !Registered(calendar) {
		return ""
	}

	return registeredCalendars[canonCalendarName(calendar)].DefaultFormat
}

// ToInternal returns the associated value from a registered calendar system.
func ToInternal(calendar string, date interface{}, format string) (TAI64NAXURTime, error) {
	if !Registered(calendar) {
		return TAI64NAXURTime{}, ErrUnknownCalendar
	}

	if format == "" {
		format = DefaultFormat(calendar)
	}

	// fmt.Printf("ToInternal: %#v (%#v) [%#v]\n", canonCalendarName(calendar), format, date)

	return registeredCalendars[canonCalendarName(calendar)].ToInternal(date, format)
}

// FromInternal returns the associated value from a registered calendar system.
func FromInternal(calendar string, stamp TAI64NAXURTime, format string) (string, error) {
	if !Registered(calendar) {
		return "", ErrUnknownCalendar
	}

	if format == "" {
		format = DefaultFormat(calendar)
	}

	// fmt.Printf("FromInternal: %#v (%#v) [%#v]\n", canonCalendarName(calendar), format, stamp)

	return registeredCalendars[canonCalendarName(calendar)].FromInternal(stamp, format)
}

// Offset returns the associated value from a registered calendar system.
func Offset(calendar string, stamp TAI64NAXURTime, offset interface{}) (TAI64NAXURTime, error) {
	if !Registered(calendar) {
		return TAI64NAXURTime{}, ErrUnknownCalendar
	}

	return registeredCalendars[canonCalendarName(calendar)].Offset(stamp, offset)
}

func canonCalendarName(calendar string) string {
	return strings.Join(strings.Fields(strings.Title(strings.ToLower(calendar))), "")
}
