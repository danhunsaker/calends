package calendars

import (
	"math/big"
	"testing"
)

func TestUnixToInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{1, ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{1., ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{1.0, ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{"1", ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{"1.", ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{"1.0", ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{big.NewFloat(1), ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{*big.NewFloat(1), ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},
		{"in": []interface{}{[]byte("1"), ""}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("-6.997489999999999987778664944926276803016662598"), nil}},

		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, ""}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := ToInternal("unix", c["in"][0], c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("UnixToInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("UnixToInternal(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0], c["in"][1].(string), out, c["want"][0].(TAI64NARUXTime))
		}
	}
}

func TestUnixFromInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, ""}, "want": []interface{}{"9.000082000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "%f"}, "want": []interface{}{"9.000082", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "%.45f"}, "want": []interface{}{"9.000081999999999027295416453853249549865722656", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "%.0f"}, "want": []interface{}{"9", nil}},
	}

	for _, c := range cases {
		out, err := FromInternal("unix", c["in"][0].(TAI64NARUXTime), c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("UnixFromInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(string) {
			t.Errorf("UnixFromInternal(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0], c["in"][1].(string), out, c["want"][0].(string))
		}
	}
}

func TestUnixOffset(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{}, 1}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("1"), nil}},
		{"in": []interface{}{TAI64NARUXTime{}, 1.0}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("1"), nil}},
		{"in": []interface{}{TAI64NARUXTime{}, "1"}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("1"), nil}},
		{"in": []interface{}{TAI64NARUXTime{}, []byte("1")}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("1"), nil}},
		{"in": []interface{}{TAI64NARUXTime{}, big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("1"), nil}},
		{"in": []interface{}{TAI64NARUXTime{}, *big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTimeFromDecimalString("1"), nil}},
		{"in": []interface{}{TAI64NARUXTime{}, TAI64NARUXTime{Seconds: 1}}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := Offset("unix", c["in"][0].(TAI64NARUXTime), c["in"][1])
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("UnixOffset(%#v, %#v) gave error %#v; want %#v", c["in"][0].(TAI64NARUXTime), c["in"][1], err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("UnixOffset(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0].(TAI64NARUXTime), c["in"][1], out, c["want"][0].(TAI64NARUXTime))
		}
	}
}
