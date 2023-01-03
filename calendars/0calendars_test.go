package calendars

import (
	"strings"
	"testing"

	"github.com/go-errors/errors"
)

var testCalendarElements calendarRegistration

type testCalendarClass struct {
	DefaultFormat string
}

func (testCalendarClass) ToInternal(in interface{}, mod string) (out TAI64NARUXTime, err error) {
	err = errors.New("testToInternal")
	return
}

func (testCalendarClass) FromInternal(in TAI64NARUXTime, mod string) (out string, err error) {
	err = errors.New("testFromInternal")
	return
}

func (testCalendarClass) Offset(in TAI64NARUXTime, mod interface{}) (out TAI64NARUXTime, err error) {
	err = errors.New("testOffset")
	return
}

func init() {
	testCalendarElements.ToInternal = func(in interface{}, mod string) (out TAI64NARUXTime, err error) {
		return
	}
	testCalendarElements.FromInternal = func(in TAI64NARUXTime, mod string) (out string, err error) {
		return
	}
	testCalendarElements.Offset = func(in TAI64NARUXTime, mod interface{}) (out TAI64NARUXTime, err error) {
		return
	}
	testCalendarElements.DefaultFormat = "test"
}

func TestRegisterClass(t *testing.T) {
	instance := testCalendarClass{}
	RegisterObject("testCalendarClass", instance, instance.DefaultFormat)

	if !Registered("testCalendarClass") {
		t.Errorf("RegisterObject(%#v, %#v) failed", "testCalendarClass", instance)
	}
}

func TestRegisterElements(t *testing.T) {
	RegisterElements("testCalendarElements", testCalendarElements.ToInternal, testCalendarElements.FromInternal, testCalendarElements.Offset, testCalendarElements.DefaultFormat)

	if !Registered("testCalendarElements") {
		t.Errorf("RegisterElements(%#v, %#v) failed", "testCalendarElements", testCalendarElements)
	}
}

func TestUnregister(t *testing.T) {
	instance := testCalendarClass{}
	RegisterObject("testCalendarClass", instance, instance.DefaultFormat)

	if !Registered("testCalendarClass") {
		t.Errorf("RegisterObject(%#v, %#v) failed", "testCalendarClass", instance)
	}

	Unregister("testCalendarClass")

	if Registered("testCalendarClass") {
		t.Errorf("Unregister(%#v) failed", "testCalendarClass")
	}
}

func TestRegistered(t *testing.T) {
	if Registered("testCalendar") {
		t.Errorf("Registered(testCalendar) failed - the calendar should not be registered before the test starts")
	}

	instance := testCalendarClass{"test"}
	RegisterObject("testCalendar", instance, instance.DefaultFormat)

	if !Registered("testCalendar") {
		t.Errorf("Registered(testCalendar) failed - the calendar should be registered before the test ends")
	}
}

func TestListRegistered(t *testing.T) {
	Unregister("testCalendar")
	Unregister("testCalendarClass")
	Unregister("testCalendarElements")

	got := ListRegistered()
	want := []string{"Gregorian", "Jdc", "Stardate", "Tai64", "Unix"}
	if strings.Join(got, ":") != strings.Join(want, ":") {
		t.Errorf("ListRegistered() failed - the calendar list was incorrect:\n\twant: %v\n\t got: %v", want, got)
	}

	instance := testCalendarClass{"test"}
	RegisterObject("teSt CaLenDar", instance, instance.DefaultFormat)

	got = ListRegistered()
	want = []string{"Gregorian", "Jdc", "Stardate", "Tai64", "TestCalendar", "Unix"}
	if strings.Join(got, ":") != strings.Join(want, ":") {
		t.Errorf("ListRegistered() failed - the calendar list was incorrect:\n\twant: %v\n\t got: %v", want, got)
	}

	Unregister("test calendar")
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

	if err.Error() != ErrUnknownCalendar("invalid").Error() {
		t.Errorf("ToInternal(\"invalid\", \"\", \"\") failed - got %q, but wanted %q", err, ErrUnknownCalendar("invalid"))
	}
}

func TestFromInternal(t *testing.T) {
	if !Registered("testCalendar") {
		TestRegistered(t)
	}

	_, err := FromInternal("testCalendar", TAI64NARUXTime{}, "")

	if err == nil || err.Error() != "testFromInternal" {
		t.Errorf("FromInternal(\"testCalendar\", TAI64NARUXTime{}, \"\") failed - got %#v, but wanted %#v", err.Error(), "testFromInternal")
	}

	_, err = FromInternal("invalid", TAI64NARUXTime{}, "")

	if err.Error() != ErrUnknownCalendar("invalid").Error() {
		t.Errorf("FromInternal(\"invalid\", \"\", \"\") failed - got %q, but wanted %q", err, ErrUnknownCalendar("invalid"))
	}
}

func TestOffset(t *testing.T) {
	if !Registered("testCalendar") {
		TestRegistered(t)
	}

	_, err := Offset("testCalendar", TAI64NARUXTime{}, "")

	if err == nil || err.Error() != "testOffset" {
		t.Errorf("Offset(\"testCalendar\", TAI64NARUXTime{}, \"\") failed - got %#v, but wanted %#v", err.Error(), "testOffset")
	}

	_, err = Offset("invalid", TAI64NARUXTime{}, "")

	if err.Error() != ErrUnknownCalendar("invalid").Error() {
		t.Errorf("Offset(\"invalid\", \"\", \"\") failed - got %q, but wanted %q", err, ErrUnknownCalendar("invalid"))
	}
}

func TestCanonCalendarName(t *testing.T) {
	in := "testCalendar"
	out := canonCalendarName(in)

	if out != "Testcalendar" {
		t.Errorf("canonCalendarName(%#v) failed - got %#v, but wanted %#v", in, out, "Testcalendar")
	}
}
