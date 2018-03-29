package calendars

import (
	"testing"
	"time"
)

func TestGregorianToInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{"Thu, 01 Jan 1970 00:00:01 UTC", ""}, "want": []interface{}{TAI64NAXURTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{"1970-01-01 00:00:01 UTC", "2006-01-02 15:04:05 MST"}, "want": []interface{}{TAI64NAXURTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{"1970-01-01 00:00:01 UTC", "%Y-%m-%d %H:%M:%S %Z"}, "want": []interface{}{TAI64NAXURTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{[]byte("Thu, 01 Jan 1970 00:00:01 UTC"), ""}, "want": []interface{}{TAI64NAXURTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{time.Unix(1, 0), ""}, "want": []interface{}{TAI64NAXURTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},

		{"in": []interface{}{1, ""}, "want": []interface{}{TAI64NAXURTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{1., ""}, "want": []interface{}{TAI64NAXURTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{1.0, ""}, "want": []interface{}{TAI64NAXURTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{TAI64NAXURTime{Seconds: 1}, ""}, "want": []interface{}{TAI64NAXURTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := ToInternal("gregorian", c["in"][0], c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("GregorianToInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NAXURTime) {
			t.Errorf("GregorianToInternal(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0], c["in"][1].(string), out, c["want"][0].(TAI64NAXURTime))
		}
	}
}

func TestGregorianFromInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NAXURTime{Seconds: 1}, ""}, "want": []interface{}{"Thu, 01 Jan 1970 00:00:01 UTC", nil}},
		{"in": []interface{}{TAI64NAXURTime{Seconds: 1}, "2006-01-02 15:04:05 MST"}, "want": []interface{}{"1970-01-01 00:00:01 UTC", nil}},
		{"in": []interface{}{TAI64NAXURTime{Seconds: 1}, "%Y-%m-%d %H:%M:%S %Z"}, "want": []interface{}{"1970-01-01 00:00:01 UTC", nil}},
		{"in": []interface{}{TAI64NAXURTime{Seconds: 1}, "%Y-%m-%d %H:%M:%S %z"}, "want": []interface{}{"1970-01-01 00:00:01 +0000", nil}},
	}

	for _, c := range cases {
		out, err := FromInternal("gregorian", c["in"][0].(TAI64NAXURTime), c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("GregorianFromInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(string) {
			t.Errorf("GregorianFromInternal(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0], c["in"][1].(string), out, c["want"][0].(string))
		}
	}
}

func TestGregorianOffset(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NAXURTime{}, "in 1 second"}, "want": []interface{}{TAI64NAXURTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NAXURTime{}, []byte("in 1 second")}, "want": []interface{}{TAI64NAXURTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NAXURTime{}, time.Second}, "want": []interface{}{TAI64NAXURTime{Seconds: 1}, nil}},

		{"in": []interface{}{TAI64NAXURTime{}, "in 17 bloxnards"}, "want": []interface{}{TAI64NAXURTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{TAI64NAXURTime{}, TAI64NAXURTime{Seconds: 1}}, "want": []interface{}{TAI64NAXURTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := Offset("gregorian", c["in"][0].(TAI64NAXURTime), c["in"][1])
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("GregorianOffset(%#v, %#v) gave error %#v; want %#v", c["in"][0].(TAI64NAXURTime), c["in"][1], err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NAXURTime) {
			t.Errorf("GregorianOffset(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0].(TAI64NAXURTime), c["in"][1], out, c["want"][0].(TAI64NAXURTime))
		}
	}
}
