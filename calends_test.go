package calends

import (
	"encoding/json"
	"errors"
	"math/big"
	"testing"

	"github.com/danhunsaker/calends/calendars"
)

func testValue(dur float64) Calends {
	out, _ := Create(map[string]interface{}{"start": "0", "duration": dur}, "tai64", "decimal")
	return out
}

func TestCreate(t *testing.T) {
	var got interface{}
	var err error

	successCases := []struct {
		in []interface{}
	}{
		{[]interface{}{0, "", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "end": 0}, "", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": 0}, "", ""}},
		{[]interface{}{map[string]interface{}{"duration": 0, "end": 0}, "", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": "0", "end": 0}, "", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": []byte("0"), "end": 0}, "", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": 0, "end": 0}, "", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": big.NewFloat(0), "end": 0}, "", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": *big.NewFloat(0), "end": 0}, "", ""}},

		{[]interface{}{0, "unix", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "end": 0}, "unix", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": 0}, "unix", ""}},
		{[]interface{}{map[string]interface{}{"duration": 0, "end": 0}, "unix", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": "0", "end": 0}, "unix", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": []byte("0"), "end": 0}, "unix", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": 0, "end": 0}, "unix", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": big.NewFloat(0), "end": 0}, "unix", ""}},
		{[]interface{}{map[string]interface{}{"start": 0, "duration": *big.NewFloat(0), "end": 0}, "unix", ""}},
	}

	failureCases := []struct {
		in   []interface{}
		want error
	}{
		{[]interface{}{0, "invalid", ""}, calendars.ErrUnknownCalendar("invalid")},

		{[]interface{}{map[string]interface{}{"start": calendars.ErrInvalidFormat}, "", ""}, calendars.ErrUnsupportedInput},
		{[]interface{}{map[string]interface{}{"end": calendars.ErrInvalidFormat}, "", ""}, calendars.ErrUnsupportedInput},
		{[]interface{}{map[string]interface{}{"duration": 0}, "", ""}, calendars.ErrUnsupportedInput},

		{[]interface{}{map[string]interface{}{"duration": calendars.ErrInvalidFormat}, "", ""}, errors.New("Invalid Duration Type")},
	}

	got, err = Create(nil, "", "")
	if err != nil {
		t.Errorf("Create(%#v, %#v, %#v) gave error %q", nil, "", "", err)
	}
	switch got := got.(type) {
	case Calends:
		if got.startTime.String() != "0" {
			t.Errorf("Create(%#v, %#v, %#v) returned with startTime %s", nil, "", "", got.startTime)
		}
		if got.duration.String() != "0" {
			t.Errorf("Create(%#v, %#v, %#v) returned with duration %s", nil, "", "", got.duration)
		}
		if got.endTime.String() != "0" {
			t.Errorf("Create(%#v, %#v, %#v) returned with endTime %s", nil, "", "", got.endTime)
		}
	default:
		t.Errorf("Create(%#v, %#v, %#v)\n  ==  %#v,\n want %#v", nil, "", "", got, "Calends")
	}

	for _, c := range successCases {
		got, err = Create(c.in[0], c.in[1].(string), c.in[2].(string))
		if err != nil {
			t.Errorf("Create(%q) gave error %q", c.in, err)
		}
		switch got := got.(type) {
		case Calends:
			if got.startTime.String() != "-7.997489999999999987778664944926276803016662598" {
				t.Errorf("Create(%q) returned with startTime %s", c.in, got.startTime)
			}
			if got.duration.String() != "0" {
				t.Errorf("Create(%q) returned with duration %s", c.in, got.duration)
			}
			if got.endTime.String() != "-7.997489999999999987778664944926276803016662598" {
				t.Errorf("Create(%q) returned with endTime %s", c.in, got.endTime)
			}
		default:
			t.Errorf("Create(%q)\n  ==  %#v,\n want %#v", c.in, got, "Calends")
		}
	}

	for _, c := range failureCases {
		_, err = Create(c.in[0], c.in[1].(string), c.in[2].(string))

		if err == nil {
			t.Errorf("Create(%q) didn't give error; expected %q", c.in, c.want)
		}
		if err.Error() != c.want.Error() {
			t.Errorf("Create(%q) gave error %q; expected %q", c.in, err, c.want)
		}
	}
}

func TestDate(t *testing.T) {
	got, err := testValue(0).Date("unix", "")
	if err != nil {
		t.Errorf("Date(%q) gave error %q", []interface{}{"unix", ""}, err)
	}
	if got != "8.000082000" {
		t.Errorf("Date(%q)\n  ==  %q,\n want %q", []interface{}{"unix", ""}, got, "8.000082000")
	}

	got, err = testValue(0).Date("", "")
	if err != nil {
		t.Errorf("Date(%q) gave error %q", []interface{}{"", ""}, err)
	}
	if got != "8.000082000" {
		t.Errorf("Date(%q)\n  ==  %q,\n want %q", []interface{}{"", ""}, got, "8.000082000")
	}

	_, err = testValue(0).Date("invalid", "")
	if err == nil {
		t.Errorf("Date(%q) didn't give error; expected %q", []interface{}{"invalid", ""}, calendars.ErrUnknownCalendar("invalid"))
	}
	if err.Error() != calendars.ErrUnknownCalendar("invalid").Error() {
		t.Errorf("Date(%q) gave error %q; expected %q", []interface{}{"invalid", ""}, err, calendars.ErrUnknownCalendar("invalid"))
	}
}

func TestDuration(t *testing.T) {
	got := testValue(0).Duration()
	if got.String() != "0" {
		t.Errorf("Duration()\n  ==  %q,\n want %q", got.String(), "0")
	}
}

func TestEndDate(t *testing.T) {
	got, err := testValue(0).EndDate("unix", "")
	if err != nil {
		t.Errorf("EndDate(%q) gave error %q", []interface{}{"unix", ""}, err)
	}
	if got != "8.000082000" {
		t.Errorf("EndDate(%q)\n  ==  %q,\n want %q", []interface{}{"unix", ""}, got, "8.000082000")
	}

	got, err = testValue(0).EndDate("", "")
	if err != nil {
		t.Errorf("EndDate(%q) gave error %q", []interface{}{"", ""}, err)
	}
	if got != "8.000082000" {
		t.Errorf("EndDate(%q)\n  ==  %q,\n want %q", []interface{}{"", ""}, got, "8.000082000")
	}

	_, err = testValue(0).EndDate("invalid", "")
	if err == nil {
		t.Errorf("EndDate(%q) didn't give error; expected %q", []interface{}{"invalid", ""}, calendars.ErrUnknownCalendar("invalid"))
	}
	if err.Error() != calendars.ErrUnknownCalendar("invalid").Error() {
		t.Errorf("EndDate(%q) gave error %q; expected %q", []interface{}{"invalid", ""}, err, calendars.ErrUnknownCalendar("invalid"))
	}
}

func TestString(t *testing.T) {
	zero := testValue(0).String()
	if zero != "40000000000000000000000000000000000000000000000000000000" {
		t.Errorf("zero.String()\n  ==  %q,\n want %q", zero, "40000000000000000000000000000000000000000000000000000000")
	}

	day := testValue(86400).String()
	if day != "from 40000000000000000000000000000000000000000000000000000000 to 40000000000151800000000000000000000000000000000000000000" {
		t.Errorf("day.String()\n  ==  %q,\n want %q", day, "from 40000000000000000000000000000000000000000000000000000000 to 40000000000151800000000000000000000000000000000000000000")
	}
}

func TestMarshalText(t *testing.T) {
	zero, err := testValue(0).MarshalText()
	if err != nil {
		t.Errorf("zero.MarshalText() gave error %q", err)
	}
	if string(zero) != "40000000000000000000000000000000000000000000000000000000" {
		t.Errorf("zero.MarshalText()\n  ==  %q,\n want %q", zero, "40000000000000000000000000000000000000000000000000000000")
	}

	day, err := testValue(86400).MarshalText()
	if err != nil {
		t.Errorf("day.MarshalText() gave error %q", err)
	}
	if string(day) != "40000000000000000000000000000000000000000000000000000000::40000000000151800000000000000000000000000000000000000000" {
		t.Errorf("day.MarshalText()\n  ==  %q,\n want %q", day, "40000000000000000000000000000000000000000000000000000000::40000000000151800000000000000000000000000000000000000000")
	}
}

func TestUnmarshalText(t *testing.T) {
	zero := testValue(0)
	zeroOut, _ := zero.MarshalText()
	err := zero.UnmarshalText(zeroOut)
	if err != nil {
		t.Errorf("zero.UnmarshalText(%q) gave error %q", zeroOut, err)
	}
	if zero.startTime != testValue(0).startTime || zero.endTime != testValue(0).endTime || zero.duration.String() != testValue(0).duration.String() {
		t.Errorf("zero.UnmarshalText(%q)\n  ==  %#v,\n want %#v", zeroOut, zero, testValue(0))
	}

	day := testValue(86400)
	dayOut, _ := day.MarshalText()
	err = day.UnmarshalText(dayOut)
	if err != nil {
		t.Errorf("day.UnmarshalText(%q) gave error %q", dayOut, err)
	}
	if day.startTime != testValue(86400).startTime || day.endTime != testValue(86400).endTime || day.duration.String() != testValue(86400).duration.String() {
		t.Errorf("day.UnmarshalText(%q)\n  ==  %#v,\n want %#v", dayOut, day, testValue(86400))
	}
}

func TestMarshalJSON(t *testing.T) {
	zero, err := testValue(0).MarshalJSON()
	if err != nil {
		t.Errorf("zero.MarshalJSON() gave error %q", err)
	}
	if string(zero) != `"40000000000000000000000000000000000000000000000000000000"` {
		t.Errorf("zero.MarshalJSON()\n  ==  %q,\n want %q", zero, `"40000000000000000000000000000000000000000000000000000000"`)
	}

	day, err := testValue(86400).MarshalJSON()
	if err != nil {
		t.Errorf("day.MarshalJSON() gave error %q", err)
	}
	if string(day) != `{"start":"40000000000000000000000000000000000000000000000000000000","end":"40000000000151800000000000000000000000000000000000000000"}` {
		t.Errorf("day.MarshalJSON()\n  ==  %q,\n want %q", day, `{"start":"40000000000000000000000000000000000000000000000000000000","end":"40000000000151800000000000000000000000000000000000000000"}`)
	}

	both, err := json.Marshal(map[string]interface{}{
		"day":  testValue(86400),
		"zero": testValue(0),
	})
	if err != nil {
		t.Errorf("both.MarshalJSON() gave error %q", err)
	}
	if string(both) != `{"day":{"start":"40000000000000000000000000000000000000000000000000000000","end":"40000000000151800000000000000000000000000000000000000000"},"zero":"40000000000000000000000000000000000000000000000000000000"}` {
		t.Errorf("both.MarshalJSON()\n  ==  %q,\n want %q", both, `{"day":{"start":"40000000000000000000000000000000000000000000000000000000","end":"40000000000151800000000000000000000000000000000000000000"},"zero":"40000000000000000000000000000000000000000000000000000000"}`)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	var both map[string]Calends
	var zero, day Calends

	zeroIn := testValue(0)
	zeroOut, _ := zeroIn.MarshalJSON()
	err := zero.UnmarshalJSON(zeroOut)
	if err != nil {
		t.Errorf("zero.UnmarshalJSON(%q) gave error %q\n", zeroOut, err)
	}
	if zero.startTime != zeroIn.startTime || zero.endTime != zeroIn.endTime || zero.duration.String() != zeroIn.duration.String() {
		t.Errorf("zero.UnmarshalJSON(%q)\n  ==  %#v,\n want %#v\n", zeroOut, zero, zeroIn)
	}

	dayIn := testValue(86400)
	dayOut, _ := dayIn.MarshalJSON()
	err = day.UnmarshalJSON(dayOut)
	if err != nil {
		t.Errorf("day.UnmarshalJSON(%q) gave error %q\n", dayOut, err)
	}
	if day.startTime != dayIn.startTime || day.endTime != dayIn.endTime || day.duration.String() != dayIn.duration.String() {
		t.Errorf("day.UnmarshalJSON(%q)\n  ==  %#v,\n want %#v\n", dayOut, day, dayIn)
	}

	bothIn := map[string]Calends{
		"day":  dayIn,
		"zero": zeroIn,
	}
	bothOut, _ := json.Marshal(bothIn)
	err = json.Unmarshal(bothOut, &both)
	if err != nil {
		t.Errorf("both.UnmarshalJSON(%q) gave error %q\n", bothOut, err)
	}
	if both["zero"].startTime != bothIn["zero"].startTime ||
		both["day"].startTime != bothIn["day"].startTime ||
		both["zero"].endTime != bothIn["zero"].endTime ||
		both["day"].endTime != bothIn["day"].endTime ||
		both["zero"].duration.String() != bothIn["zero"].duration.String() ||
		both["day"].duration.String() != bothIn["day"].duration.String() {

		t.Errorf("both.UnmarshalJSON(%q)\n  ==  %#v,\n want %#v\n", bothOut, both, bothIn)
	}
}
