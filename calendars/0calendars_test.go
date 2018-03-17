package calendars

import (
	"errors"
	"testing"
)

var testCalendarElements calendarRegistration

type testCalendarClass struct {
	DefaultFormat string
}

func (testCalendarClass) ToInternal(in interface{}, mod string) (out TAI64NAXURTime, err error) {
	err = errors.New("testToInternal")
	return
}

func (testCalendarClass) FromInternal(in TAI64NAXURTime, mod string) (out string, err error) {
	err = errors.New("testFromInternal")
	return
}

func (testCalendarClass) Offset(in TAI64NAXURTime, mod interface{}) (out TAI64NAXURTime, err error) {
	err = errors.New("testOffset")
	return
}

func init() {
	testCalendarElements.ToInternal = func(in interface{}, mod string) (out TAI64NAXURTime, err error) {
		return
	}
	testCalendarElements.FromInternal = func(in TAI64NAXURTime, mod string) (out string, err error) {
		return
	}
	testCalendarElements.Offset = func(in TAI64NAXURTime, mod interface{}) (out TAI64NAXURTime, err error) {
		return
	}
	testCalendarElements.DefaultFormat = "test"
}

func TestRegisterClass(t *testing.T) {
	instance := testCalendarClass{}
	RegisterClass("testCalendarClass", instance, instance.DefaultFormat)

	if !Registered("testCalendarClass") {
		t.Errorf("RegisterClass(%#v, %#v) failed", "testCalendarClass", instance)
	}
}

func TestRegisterElements(t *testing.T) {
	RegisterElements("testCalendarElements", testCalendarElements.ToInternal, testCalendarElements.FromInternal, testCalendarElements.Offset, testCalendarElements.DefaultFormat)

	if !Registered("testCalendarElements") {
		t.Errorf("RegisterElements(%#v, %#v) failed", "testCalendarElements", testCalendarElements)
	}
}

func TestRegistered(t *testing.T) {
	if Registered("testCalendar") {
		t.Errorf("Registered(testCalendar) failed - the calendar should not be registered before the test starts")
	}

	instance := testCalendarClass{"test"}
	RegisterClass("testCalendar", instance, instance.DefaultFormat)

	if !Registered("testCalendar") {
		t.Errorf("Registered(testCalendar) failed - the calendar should be registered before the test ends")
	}
}

func TestDefaultFormat(t *testing.T) {
	if !Registered("testCalendar") {
		TestRegistered(t)
	}

	if DefaultFormat("testCalendar") != "test" {
		t.Errorf("DefaultFormat(testCalendar) failed - got %#v, but wanted %#v", DefaultFormat("testCalendar"), "test")
	}

	if DefaultFormat("invalid") != "" {
		t.Errorf("DefaultFormat(invalid) failed - got %#v, but wanted %#v", DefaultFormat("invalid"), "")
	}
}

func TestToInternal(t *testing.T) {
	if !Registered("testCalendar") {
		TestRegistered(t)
	}

	_, err := ToInternal("testCalendar", "", "")

	if err == nil || err.Error() != "testToInternal" {
		t.Errorf("ToInternal(\"testCalendar\", \"\", \"\") failed - got %#v, but wanted %#v", err.Error(), "testToInternal")
	}

	_, err = ToInternal("invalid", "", "")

	if err != UnknownCalendarError {
		t.Errorf("ToInternal(\"invalid\", \"\", \"\") failed - got %q, but wanted %q", err, UnknownCalendarError)
	}
}

func TestFromInternal(t *testing.T) {
	if !Registered("testCalendar") {
		TestRegistered(t)
	}

	_, err := FromInternal("testCalendar", TAI64NAXURTime{}, "")

	if err == nil || err.Error() != "testFromInternal" {
		t.Errorf("FromInternal(\"testCalendar\", TAI64NAXURTime{}, \"\") failed - got %#v, but wanted %#v", err.Error(), "testFromInternal")
	}

	_, err = FromInternal("invalid", TAI64NAXURTime{}, "")

	if err != UnknownCalendarError {
		t.Errorf("FromInternal(\"invalid\", \"\", \"\") failed - got %q, but wanted %q", err, UnknownCalendarError)
	}
}

func TestOffset(t *testing.T) {
	if !Registered("testCalendar") {
		TestRegistered(t)
	}

	_, err := Offset("testCalendar", TAI64NAXURTime{}, "")

	if err == nil || err.Error() != "testOffset" {
		t.Errorf("Offset(\"testCalendar\", TAI64NAXURTime{}, \"\") failed - got %#v, but wanted %#v", err.Error(), "testOffset")
	}

	_, err = Offset("invalid", TAI64NAXURTime{}, "")

	if err != UnknownCalendarError {
		t.Errorf("Offset(\"invalid\", \"\", \"\") failed - got %q, but wanted %q", err, UnknownCalendarError)
	}
}

func TestCanonCalendarName(t *testing.T) {
	in := "testCalendar"
	out := canonCalendarName(in)

	if out != "Testcalendar" {
		t.Errorf("canonCalendarName(%#v) failed - got %#v, but wanted %#v", in, out, "Testcalendar")
	}
}
