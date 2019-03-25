package dynamic

import (
	"errors"
	"io/ioutil"
	"math/big"
	"testing"
)

func getCalendar(t *testing.T) (raw []byte, calendar Calendar) {
	t.Helper()

	raw, err := ioutil.ReadFile("../../tests/dynamic.json")
	if err != nil {
		t.Fatalf("Failed to load test dynamic calendar: %#v", err)
	}

	if err := calendar.UnmarshalJSON(raw); err != nil {
		t.Fatalf("Failed to unmarshal test dynamic calendar: %#v", err)
	}

	return
}

func TestCalendarUnmarshalJson(t *testing.T) {
	getCalendar(t)
}

func TestCalendarMarshalJson(t *testing.T) {
	raw, calendar := getCalendar(t)

	testRaw, err:= calendar.MarshalJSON()
	if err != nil {
		t.Fatalf("Failed to marshal test dynamic calendar: %#v", err)
	}

	testCal := Calendar{}
	if err := testCal.UnmarshalJSON(testRaw); err != nil {
		t.Errorf("Failed to re-unmarshal test dynamic calendar: %#v", err)
	}

	if string(append(testRaw, '\n')) != string(raw) {
		t.Errorf("Source JSON doesn't match test JSON: %s != %s", raw, testRaw)
	}
}

func TestCalendarToTimestamp(t *testing.T) {
	_, calendar := getCalendar(t)

	cases := []map[string][]interface{}{
		{"in": []interface{}{*big.NewFloat(0), ""}, "want": []interface{}{big.NewFloat(0), nil}},
		{"in": []interface{}{big.NewFloat(0.1), "default"}, "want": []interface{}{big.NewFloat(0.1), nil}},
		{"in": []interface{}{0.1, ""}, "want": []interface{}{big.NewFloat(0.1), nil}},
		{"in": []interface{}{1, ""}, "want": []interface{}{big.NewFloat(1), nil}},
		{"in": []interface{}{"0-100.000000000", "T-t.n"}, "want": []interface{}{big.NewFloat(100), nil}},
		{"in": []interface{}{[]byte("0-100.000000000"), ""}, "want": []interface{}{big.NewFloat(100), nil}},

		// {"in": []interface{}{1, ""}, "want": []interface{}{big.NewFloat(0), errors.New("Unsupported Value")}},
		// {"in": []interface{}{1., ""}, "want": []interface{}{big.NewFloat(0), errors.New("Unsupported Value")}},
		// {"in": []interface{}{1.0, ""}, "want": []interface{}{big.NewFloat(0), errors.New("Unsupported Value")}},
		{"in": []interface{}{t, ""}, "want": []interface{}{nil, errors.New("Unsupported Value")}},
	}

	for _, c := range cases {
		out, err := calendar.ToTimestamp(c["in"][0], c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("dynamic.ToTimestamp(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if (out == nil && c["want"][0] != nil) || (out != nil && c["want"][0] == nil) || (out != nil && out.Cmp(c["want"][0].(*big.Float)) != 0) {
			t.Errorf("dynamic.ToTimestamp(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0], c["in"][1].(string), out, c["want"][0])
		}
	}
}

func TestCalendarFromTimestamp(t *testing.T) {
	_, calendar := getCalendar(t)

	cases := []map[string][]interface{}{
		{"in": []interface{}{big.NewFloat(0.1), ""}, "want": []interface{}{"0-000.100000000", nil}},
		{"in": []interface{}{big.NewFloat(100), ""}, "want": []interface{}{"0-100.000000000", nil}},
		{"in": []interface{}{big.NewFloat(1000), ""}, "want": []interface{}{"1-000.000000000", nil}},
		{"in": []interface{}{big.NewFloat(1000), "shift"}, "want": []interface{}{"morning, 0 of unix", nil}},
		{"in": []interface{}{big.NewFloat(-1000), "shift"}, "want": []interface{}{"night, 0 of darkness", nil}},
		{"in": []interface{}{big.NewFloat(-1000), "S, d of e (T-t.n)"}, "want": []interface{}{"night, 0 of darkness (-1-000.000000000)", nil}},
	}

	for _, c := range cases {
		out, err := calendar.FromTimestamp(c["in"][0].(*big.Float), c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("dynamic.FromTimestamp(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(string) {
			t.Errorf("dynamic.FromTimestamp(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0], c["in"][1].(string), out, c["want"][0])
		}
	}
}

func TestCalendarOffset(t *testing.T) {
	_, calendar := getCalendar(t)

	cases := []map[string][]interface{}{
		{"in": []interface{}{big.NewFloat(0), *big.NewFloat(0)}, "want": []interface{}{big.NewFloat(0), nil}},
		{"in": []interface{}{big.NewFloat(0), big.NewFloat(0.1)}, "want": []interface{}{big.NewFloat(0.1), nil}},
		{"in": []interface{}{big.NewFloat(0), 0.1}, "want": []interface{}{big.NewFloat(0.1), nil}},
		{"in": []interface{}{big.NewFloat(0), 1}, "want": []interface{}{big.NewFloat(1), nil}},
		{"in": []interface{}{big.NewFloat(0), "100"}, "want": []interface{}{big.NewFloat(100), nil}},
		{"in": []interface{}{big.NewFloat(0), []byte("100")}, "want": []interface{}{big.NewFloat(100), nil}},
		{"in": []interface{}{big.NewFloat(0), "1 tock"}, "want": []interface{}{big.NewFloat(1000), nil}},
		{"in": []interface{}{big.NewFloat(0), "1 cycle"}, "want": []interface{}{big.NewFloat(86400), nil}},

		// {"in": []interface{}{big.NewFloat(0), 1}, "want": []interface{}{big.NewFloat(0), errors.New("Unsupported Value")}},
		// {"in": []interface{}{big.NewFloat(0), 1.}, "want": []interface{}{big.NewFloat(0), errors.New("Unsupported Value")}},
		// {"in": []interface{}{big.NewFloat(0), 1.0}, "want": []interface{}{big.NewFloat(0), errors.New("Unsupported Value")}},
		{"in": []interface{}{big.NewFloat(0), t}, "want": []interface{}{nil, errors.New("Unsupported Value")}},
	}

	for _, c := range cases {
		out, err := calendar.Offset(c["in"][0].(*big.Float), c["in"][1])
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("dynamic.Offset(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1], err, c["want"][1])
		}
		if (out == nil && c["want"][0] != nil) || (out != nil && c["want"][0] == nil) || (out != nil && out.Cmp(c["want"][0].(*big.Float)) != 0) {
			t.Errorf("dynamic.Offset(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0], c["in"][1], out, c["want"][0])
		}
	}
}

func TestCalendarDateToUnits(t *testing.T) {
	// _, calendar := getCalendar(t)

	t.Skip("\x1b[35mTODO\x1b[0m: Implement Calendar.dateToUnits() test")
}

func TestCalendarUnitsToDate(t *testing.T) {
	// _, calendar := getCalendar(t)

	t.Skip("\x1b[35mTODO\x1b[0m: Implement Calendar.unitsToDate() test")
}

func TestCalendarUnitsWithOffset(t *testing.T) {
	// _, calendar := getCalendar(t)

	t.Skip("\x1b[35mTODO\x1b[0m: Implement Calendar.unitsWithOffset() test")
}

func TestCalendarUnitsToTime(t *testing.T) {
	// _, calendar := getCalendar(t)

	t.Skip("\x1b[35mTODO\x1b[0m: Implement Calendar.unitsToTime() test")
}

func TestCalendarTimeToUnits(t *testing.T) {
	// _, calendar := getCalendar(t)

	t.Skip("\x1b[35mTODO\x1b[0m: Implement Calendar.timeToUnits() test")
}

func TestCalendarSumUnits(t *testing.T) {
	// _, calendar := getCalendar(t)

	t.Skip("\x1b[35mTODO\x1b[0m: Implement Calendar.sumUnits() test")
}

func TestCalendarEpochUnits(t *testing.T) {
	// _, calendar := getCalendar(t)

	t.Skip("\x1b[35mTODO\x1b[0m: Implement Calendar.epochUnits() test")
}
