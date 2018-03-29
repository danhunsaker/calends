package calends

import (
	"errors"
	"testing"

	"github.com/danhunsaker/calends/calendars"
)

func TestAdd(t *testing.T) {
	test, err := testValue(0).Add("86400", "tai64")
	if err != nil {
		t.Errorf("Add(%#v, %#v) gives error %q", "86400", "tai64", err)
	}
	if test.startTime.String() != "86400" {
		t.Errorf("Add(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test.startTime.String(), "86400")
	}
	if test.duration.String() != "-86400" {
		t.Errorf("Add(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test.duration.String(), "-86400")
	}
	if test.endTime.String() != "0" {
		t.Errorf("Add(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test.endTime.String(), "0")
	}

	test, err = testValue(0).Add("86400", "invalid")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("Add(%#v, %#v) gives error %#v; wanted %#v", "86400", "invalid", err, calendars.ErrUnknownCalendar)
	}

	test, err = testValue(0).Add("invalid", "gregorian")
	if err != calendars.ErrUnsupportedInput {
		t.Errorf("Add(%#v, %#v) gives error %#v; wanted %#v", "invalid", "gregorian", err, calendars.ErrUnsupportedInput)
	}
}

func TestSubtract(t *testing.T) {
	test, err := testValue(0).Subtract("86400", "tai64")
	if err != nil {
		t.Errorf("Subtract(%#v, %#v) gives error %q", "86400", "tai64", err)
	}
	if test.startTime.String() != "-86400" {
		t.Errorf("Subtract(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test.startTime.String(), "-86400")
	}
	if test.duration.String() != "86400" {
		t.Errorf("Subtract(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test.duration.String(), "86400")
	}
	if test.endTime.String() != "0" {
		t.Errorf("Subtract(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test.endTime.String(), "0")
	}

	test, err = testValue(0).Subtract("86400", "invalid")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("Subtract(%#v, %#v) gives error %#v; wanted %#v", "86400", "invalid", err, calendars.ErrUnknownCalendar)
	}

	test, err = testValue(0).Subtract("invalid", "gregorian")
	if err != calendars.ErrUnsupportedInput {
		t.Errorf("Subtract(%#v, %#v) gives error %#v; wanted %#v", "invalid", "gregorian", err, calendars.ErrUnsupportedInput)
	}
}

func TestAddFromEnd(t *testing.T) {
	test, err := testValue(0).AddFromEnd("86400", "tai64")
	if err != nil {
		t.Errorf("AddFromEnd(%#v, %#v) gives error %q", "86400", "tai64", err)
	}
	if test.startTime.String() != "0" {
		t.Errorf("AddFromEnd(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test.startTime.String(), "0")
	}
	if test.duration.String() != "86400" {
		t.Errorf("AddFromEnd(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test.duration.String(), "86400")
	}
	if test.endTime.String() != "86400" {
		t.Errorf("AddFromEnd(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test.endTime.String(), "86400")
	}

	test, err = testValue(0).AddFromEnd("86400", "invalid")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("AddFromEnd(%#v, %#v) gives error %#v; wanted %#v", "86400", "invalid", err, calendars.ErrUnknownCalendar)
	}

	test, err = testValue(0).AddFromEnd("invalid", "gregorian")
	if err != calendars.ErrUnsupportedInput {
		t.Errorf("AddFromEnd(%#v, %#v) gives error %#v; wanted %#v", "invalid", "gregorian", err, calendars.ErrUnsupportedInput)
	}
}

func TestSubtractFromEnd(t *testing.T) {
	test, err := testValue(0).SubtractFromEnd("86400", "tai64")
	if err != nil {
		t.Errorf("SubtractFromEnd(%#v, %#v) gives error %q", "86400", "tai64", err)
	}
	if test.startTime.String() != "0" {
		t.Errorf("SubtractFromEnd(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test.startTime.String(), "0")
	}
	if test.duration.String() != "-86400" {
		t.Errorf("SubtractFromEnd(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test.duration.String(), "-86400")
	}
	if test.endTime.String() != "-86400" {
		t.Errorf("SubtractFromEnd(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test.endTime.String(), "-86400")
	}

	test, err = testValue(0).SubtractFromEnd("86400", "invalid")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("SubtractFromEnd(%#v, %#v) gives error %#v; wanted %#v", "86400", "invalid", err, calendars.ErrUnknownCalendar)
	}

	test, err = testValue(0).SubtractFromEnd("invalid", "gregorian")
	if err != calendars.ErrUnsupportedInput {
		t.Errorf("SubtractFromEnd(%#v, %#v) gives error %#v; wanted %#v", "invalid", "gregorian", err, calendars.ErrUnsupportedInput)
	}
}

func TestNext(t *testing.T) {
	test1, err1 := testValue(0).Next("86400", "tai64")
	test2, err2 := test1.Next("", "")
	test3, err3 := testValue(86400).Next("", "")
	test4, err4 := testValue(0).Next("86400", "")
	_, err5 := testValue(0).Next("86400", "invalid")
	_, err6 := testValue(0).Next("invalid", "gregorian")

	if err1 != nil {
		t.Errorf("1:Next(%#v, %#v) gives error %q", "86400", "tai64", err1)
	}
	if test1.startTime.String() != "0" {
		t.Errorf("1:Next(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test1.startTime.String(), "0")
	}
	if test1.duration.String() != "86400" {
		t.Errorf("1:Next(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test1.duration.String(), "86400")
	}
	if test1.endTime.String() != "86400" {
		t.Errorf("1:Next(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test1.endTime.String(), "86400")
	}

	if err2 != nil {
		t.Errorf("2:Next(%#v, %#v) gives error %q", "", "", err2)
	}
	if test2.startTime.String() != "86400" {
		t.Errorf("2:Next(%#v, %#v) has startTime of %#v\nwant %#v", "", "", test2.startTime.String(), "86400")
	}
	if test2.duration.String() != "86400" {
		t.Errorf("2:Next(%#v, %#v) has duration of %#v\nwant %#v", "", "", test2.duration.String(), "86400")
	}
	if test2.endTime.String() != "172800" {
		t.Errorf("2:Next(%#v, %#v) has endTime of %#v\nwant %#v", "", "", test2.endTime.String(), "172800")
	}

	if err3 != nil {
		t.Errorf("3:Next(%#v, %#v) gives error %q", "", "", err3)
	}
	if test3.startTime.String() != "86400" {
		t.Errorf("3:Next(%#v, %#v) has startTime of %#v\nwant %#v", "", "", test3.startTime.String(), "86400")
	}
	if test3.duration.String() != "86400" {
		t.Errorf("3:Next(%#v, %#v) has duration of %#v\nwant %#v", "", "", test3.duration.String(), "86400")
	}
	if test3.endTime.String() != "172800" {
		t.Errorf("3:Next(%#v, %#v) has endTime of %#v\nwant %#v", "", "", test3.endTime.String(), "172800")
	}

	if err4 != nil {
		t.Errorf("4:Next(%#v, %#v) gives error %q", "86400", "", err4)
	}
	if test4.startTime.String() != "0" {
		t.Errorf("4:Next(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "", test4.startTime.String(), "0")
	}
	if test4.duration.String() != "86400" {
		t.Errorf("4:Next(%#v, %#v) has duration of %#v\nwant %#v", "86400", "", test4.duration.String(), "86400")
	}
	if test4.endTime.String() != "86400" {
		t.Errorf("4:Next(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "", test4.endTime.String(), "86400")
	}

	if err5 != calendars.ErrUnknownCalendar {
		t.Errorf("5:Next(%#v, %#v) gives error %q; want %q", "86400", "invalid", err5, calendars.ErrUnknownCalendar)
	}

	if err6 != calendars.ErrUnsupportedInput {
		t.Errorf("6:Next(%#v, %#v) gives error %q; want %q", "invalid", "gregorian", err6, calendars.ErrUnsupportedInput)
	}
}

func TestPrevious(t *testing.T) {
	test1, err1 := testValue(0).Previous("86400", "tai64")
	test2, err2 := test1.Previous("", "")
	test3, err3 := testValue(86400).Previous("", "")
	test4, err4 := testValue(0).Previous("86400", "")
	_, err5 := testValue(0).Previous("86400", "invalid")
	_, err6 := testValue(0).Previous("invalid", "gregorian")

	if err1 != nil {
		t.Errorf("1:Previous(%#v, %#v) gives error %q", "86400", "tai64", err1)
	}
	if test1.startTime.String() != "-86400" {
		t.Errorf("1:Previous(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test1.startTime.String(), "-86400")
	}
	if test1.duration.String() != "86400" {
		t.Errorf("1:Previous(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test1.duration.String(), "86400")
	}
	if test1.endTime.String() != "0" {
		t.Errorf("1:Previous(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test1.endTime.String(), "0")
	}

	if err2 != nil {
		t.Errorf("2:Previous(%#v, %#v) gives error %q", "", "", err2)
	}
	if test2.startTime.String() != "-172800" {
		t.Errorf("2:Previous(%#v, %#v) has startTime of %#v\nwant %#v", "", "", test2.startTime.String(), "-172800")
	}
	if test2.duration.String() != "86400" {
		t.Errorf("2:Previous(%#v, %#v) has duration of %#v\nwant %#v", "", "", test2.duration.String(), "86400")
	}
	if test2.endTime.String() != "-86400" {
		t.Errorf("2:Previous(%#v, %#v) has endTime of %#v\nwant %#v", "", "", test2.endTime.String(), "-86400")
	}

	if err3 != nil {
		t.Errorf("3:Previous(%#v, %#v) gives error %q", "", "", err3)
	}
	if test3.startTime.String() != "-86400" {
		t.Errorf("3:Previous(%#v, %#v) has startTime of %#v\nwant %#v", "", "", test3.startTime.String(), "-86400")
	}
	if test3.duration.String() != "86400" {
		t.Errorf("3:Previous(%#v, %#v) has duration of %#v\nwant %#v", "", "", test3.duration.String(), "86400")
	}
	if test3.endTime.String() != "0" {
		t.Errorf("3:Previous(%#v, %#v) has endTime of %#v\nwant %#v", "", "", test3.endTime.String(), "0")
	}

	if err4 != nil {
		t.Errorf("4:Previous(%#v, %#v) gives error %q", "86400", "", err4)
	}
	if test4.startTime.String() != "-86400" {
		t.Errorf("4:Previous(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "", test4.startTime.String(), "-86400")
	}
	if test4.duration.String() != "86400" {
		t.Errorf("4:Previous(%#v, %#v) has duration of %#v\nwant %#v", "86400", "", test4.duration.String(), "86400")
	}
	if test4.endTime.String() != "0" {
		t.Errorf("4:Previous(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "", test4.endTime.String(), "0")
	}

	if err5 != calendars.ErrUnknownCalendar {
		t.Errorf("5:Previous(%#v, %#v) gives error %q; want %q", "86400", "invalid", err5, calendars.ErrUnknownCalendar)
	}

	if err6 != calendars.ErrUnsupportedInput {
		t.Errorf("6:Previous(%#v, %#v) gives error %q; want %q", "invalid", "gregorian", err6, calendars.ErrUnsupportedInput)
	}
}

func TestSetDate(t *testing.T) {
	test, err := testValue(0).SetDate("86400", "tai64", "decimal")
	if err != nil {
		t.Errorf("SetDate(%#v, %#v, %#v) gives error %q", "86400", "tai64", "decimal", err)
	}
	if test.startTime.String() != "86400" {
		t.Errorf("SetDate(%#v, %#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", "decimal", test.startTime.String(), "86400")
	}
	if test.duration.String() != "-86400" {
		t.Errorf("SetDate(%#v, %#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", "decimal", test.duration.String(), "-86400")
	}
	if test.endTime.String() != "0" {
		t.Errorf("SetDate(%#v, %#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", "decimal", test.endTime.String(), "0")
	}

	test, err = testValue(0).SetDate("86400", "", "")
	if err != nil {
		t.Errorf("SetDate(%#v, %#v, %#v) gives error %q", "86400", "", "", err)
	}
	if test.startTime.String() != "86392.000081999999999027295416453853249549865722656" {
		t.Errorf("SetDate(%#v, %#v, %#v) has startTime of %#v\nwant %#v", "86400", "", "", test.startTime.String(), "86392.000081999999999027295416453853249549865722656")
	}
	if test.duration.String() != "-86392.00008" {
		t.Errorf("SetDate(%#v, %#v, %#v) has duration of %#v\nwant %#v", "86400", "", "", test.duration.String(), "-86392.00008")
	}
	if test.endTime.String() != "0" {
		t.Errorf("SetDate(%#v, %#v, %#v) has endTime of %#v\nwant %#v", "86400", "", "", test.endTime.String(), "0")
	}

	_, err = testValue(0).SetDate("86400", "invalid", "")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("SetDate(%#v, %#v, %#v) gives error %q; want %q", "86400", "invalid", "", err, calendars.ErrUnknownCalendar)
	}

	_, err = testValue(0).SetDate("invalid", "gregorian", "")
	if err.Error() != `parsing time "invalid" as "Mon, 02 Jan 2006 15:04:05 MST": cannot parse "invalid" as "Mon"` {
		t.Errorf("SetDate(%#v, %#v, %#v) gives error %q; want %q", "invalid", "gregorian", "", err, errors.New(`parsing time "invalid" as "Mon, 02 Jan 2006 15:04:05 MST": cannot parse "invalid" as "Mon"`))
	}
}

func TestSetEndDate(t *testing.T) {
	test, err := testValue(0).SetEndDate("86400", "tai64", "decimal")
	if err != nil {
		t.Errorf("SetEndDate(%#v, %#v, %#v) gives error %q", "86400", "tai64", "decimal", err)
	}
	if test.startTime.String() != "0" {
		t.Errorf("SetEndDate(%#v, %#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", "decimal", test.startTime.String(), "0")
	}
	if test.duration.String() != "86400" {
		t.Errorf("SetEndDate(%#v, %#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", "decimal", test.duration.String(), "86400")
	}
	if test.endTime.String() != "86400" {
		t.Errorf("SetEndDate(%#v, %#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", "decimal", test.endTime.String(), "86400")
	}

	test, err = testValue(0).SetEndDate("86400", "", "")
	if err != nil {
		t.Errorf("SetEndDate(%#v, %#v, %#v) gives error %q", "86400", "", "", err)
	}
	if test.startTime.String() != "0" {
		t.Errorf("SetEndDate(%#v, %#v, %#v) has startTime of %#v\nwant %#v", "86400", "", "", test.startTime.String(), "0")
	}
	if test.duration.String() != "86392.00008" {
		t.Errorf("SetEndDate(%#v, %#v, %#v) has duration of %#v\nwant %#v", "86400", "", "", test.duration.String(), "86392.00008")
	}
	if test.endTime.String() != "86392.000081999999999027295416453853249549865722656" {
		t.Errorf("SetEndDate(%#v, %#v, %#v) has endTime of %#v\nwant %#v", "86400", "", "", test.endTime.String(), "86392.000081999999999027295416453853249549865722656")
	}

	_, err = testValue(0).SetEndDate("86400", "invalid", "")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("SetEndDate(%#v, %#v, %#v) gives error %q; want %q", "86400", "invalid", "", err, calendars.ErrUnknownCalendar)
	}

	_, err = testValue(0).SetEndDate("invalid", "gregorian", "")
	if err.Error() != `parsing time "invalid" as "Mon, 02 Jan 2006 15:04:05 MST": cannot parse "invalid" as "Mon"` {
		t.Errorf("SetEndDate(%#v, %#v, %#v) gives error %q; want %q", "invalid", "gregorian", "", err, errors.New(`parsing time "invalid" as "Mon, 02 Jan 2006 15:04:05 MST": cannot parse "invalid" as "Mon"`))
	}
}

func TestSetDuration(t *testing.T) {
	test, err := testValue(0).SetDuration("86400", "tai64")
	if err != nil {
		t.Errorf("SetDuration(%#v, %#v) gives error %q", "86400", "tai64", err)
	}
	if test.startTime.String() != "0" {
		t.Errorf("SetDuration(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test.startTime.String(), "0")
	}
	if test.duration.String() != "86400" {
		t.Errorf("SetDuration(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test.duration.String(), "86400")
	}
	if test.endTime.String() != "86400" {
		t.Errorf("SetDuration(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test.endTime.String(), "86400")
	}

	_, err = testValue(0).SetDuration("86400", "invalid")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("SetDuration(%#v, %#v) gives error %q; want %q", "86400", "invalid", err, calendars.ErrUnknownCalendar)
	}

	_, err = testValue(0).SetDuration("invalid", "gregorian")
	if err != calendars.ErrUnsupportedInput {
		t.Errorf("SetDuration(%#v, %#v) gives error %q; want %q", "invalid", "gregorian", err, calendars.ErrUnsupportedInput)
	}
}

func TestSetDurationFromEnd(t *testing.T) {
	test, err := testValue(0).SetDurationFromEnd("86400", "tai64")

	if err != nil {
		t.Errorf("SetDurationFromEnd(%#v, %#v) gives error %q", "86400", "tai64", err)
	}
	if test.startTime.String() != "-86400" {
		t.Errorf("SetDurationFromEnd(%#v, %#v) has startTime of %#v\nwant %#v", "86400", "tai64", test.startTime.String(), "-86400")
	}
	if test.duration.String() != "86400" {
		t.Errorf("SetDurationFromEnd(%#v, %#v) has duration of %#v\nwant %#v", "86400", "tai64", test.duration.String(), "86400")
	}
	if test.endTime.String() != "0" {
		t.Errorf("SetDurationFromEnd(%#v, %#v) has endTime of %#v\nwant %#v", "86400", "tai64", test.endTime.String(), "0")
	}

	_, err = testValue(0).SetDurationFromEnd("86400", "invalid")
	if err != calendars.ErrUnknownCalendar {
		t.Errorf("SetDurationFromEnd(%#v, %#v) gives error %q; want %q", "86400", "invalid", err, calendars.ErrUnknownCalendar)
	}

	_, err = testValue(0).SetDurationFromEnd("invalid", "gregorian")
	if err != calendars.ErrUnsupportedInput {
		t.Errorf("SetDurationFromEnd(%#v, %#v) gives error %q; want %q", "invalid", "gregorian", err, calendars.ErrUnsupportedInput)
	}
}

func TestMerge(t *testing.T) {
	temp, _ := testValue(-86400).SetDate("-172800", "tai64", "decimal")
	want, _ := testValue(86400).SetDate("-172800", "tai64", "decimal")

	test, err := temp.Merge(testValue(86400))
	if err != nil {
		t.Errorf("Merge(%s)\ngives error %q", testValue(86400), err)
	}
	if !test.IsSame(want) {
		t.Errorf("Merge(%s)\nreturns %s\nwant %s", testValue(86400), test, want)
	}

	test, err = testValue(86400).Merge(temp)
	if err != nil {
		t.Errorf("Merge(%s)\ngives error %q", temp, err)
	}
	if !test.IsSame(want) {
		t.Errorf("Merge(%s)\nreturns %s\nwant %s", temp, test, want)
	}
}

func TestIntersect(t *testing.T) {
	test1, err1 := testValue(0).Intersect(testValue(86400))
	temp, _ := testValue(-86400).SetEndDate("172800", "tai64", "decimal")
	test2, err2 := temp.Intersect(testValue(86400))
	temp, _ = testValue(-86400).SetDate("-172800", "tai64", "decimal")
	test3, err3 := temp.Intersect(testValue(86400))

	if err1 != nil {
		t.Errorf("1:Intersect(%s)\ngives error %q", testValue(86400), err1)
	}
	if err2 != nil {
		t.Errorf("2:Intersect(%s)\ngives error %q", testValue(86400), err2)
	}
	if err3 == nil {
		t.Errorf("3:Intersect(%s)\nshould give error; got nil", testValue(86400))
	}

	if !test1.IsSame(testValue(0)) {
		t.Errorf("1:Intersect(%s)\nreturns %s\nwant %s", testValue(86400), test1, testValue(0))
	}
	if !test2.IsSame(testValue(86400)) {
		t.Errorf("2:Intersect(%s)\nreturns %s\nwant %s", testValue(86400), test2, testValue(86400))
	}
	if !test3.IsSame(temp) {
		t.Errorf("3:Intersect(%s)\nreturns %s\nwant %s", testValue(86400), test3, temp)
	}
}

func TestGap(t *testing.T) {
	test1, err1 := testValue(0).Gap(testValue(86400))
	temp2, _ := testValue(-86400).SetDate("-172800", "tai64", "decimal")
	test2, err2 := temp2.Gap(testValue(86400))
	temp3, _ := testValue(-86400).SetDate("-172800", "tai64", "decimal")
	test3, err3 := testValue(86400).Gap(temp3)

	if err1 == nil {
		t.Errorf("1:Gap(%s)\nshould give error; got nil", testValue(86400))
	}
	if err2 != nil {
		t.Errorf("2:Gap(%s)\ngives error %q", testValue(86400), err2)
	}
	if err3 != nil {
		t.Errorf("3:Gap(%s)\ngives error %q", testValue(86400), err3)
	}

	if !test1.IsSame(testValue(0)) {
		t.Errorf("1:Gap(%s)\nreturns %s\nwant %s", testValue(86400), test1, testValue(0))
	}
	temp, _ := testValue(0).SetDate("-86400", "tai64", "decimal")
	if !test2.IsSame(temp) {
		t.Errorf("2:Gap(%s)\nreturns %s\nwant %s", testValue(86400), test2, temp)
	}
	temp, _ = testValue(0).SetDate("-86400", "tai64", "decimal")
	if !test3.IsSame(temp) {
		t.Errorf("3:Gap(%s)\nreturns %s\nwant %s", temp3, test3, temp)
	}
}
