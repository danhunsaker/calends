package calendars

import (
	"math/big"
	"testing"
)

func TestJDCToInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{1, ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{1., ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{1.0, ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{"1", ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{"1.", ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{"1.0", ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{big.NewFloat(1), ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{*big.NewFloat(1), ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},
		{"in": []interface{}{[]byte("1"), ""}, "want": []interface{}{TAI64NARUXTime{Seconds: -3506630400}, nil}},

		{"in": []interface{}{"2440588.6", "full"}, "want": []interface{}{TAI64NARUXTime{Seconds: 95040}, nil}},
		{"in": []interface{}{"2440588.6", "fullday"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"2440588.6", "fulltime"}, "want": []interface{}{TAI64NARUXTime{Seconds: 8640}, nil}},
		{"in": []interface{}{"40588.1", "modified"}, "want": []interface{}{TAI64NARUXTime{Seconds: 95040}, nil}},
		{"in": []interface{}{"40588.1", "day"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"40588.1", "time"}, "want": []interface{}{TAI64NARUXTime{Seconds: 8640}, nil}},

		{"in": []interface{}{"1.0", "invalid"}, "want": []interface{}{TAI64NARUXTime{}, ErrInvalidFormat}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, ""}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := ToInternal("jdc", c["in"][0], c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("JDCToInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("JDCToInternal(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0], c["in"][1].(string), out, c["want"][0].(TAI64NARUXTime))
		}
	}
}

func TestJDCFromInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, ""}, "want": []interface{}{"40587.100000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, "full"}, "want": []interface{}{"2440587.600000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, "fullday"}, "want": []interface{}{"2440587", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, "fulltime"}, "want": []interface{}{"0.600000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, "modified"}, "want": []interface{}{"40587.100000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, "day"}, "want": []interface{}{"40587", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, "time"}, "want": []interface{}{"0.100000", nil}},

		{"in": []interface{}{TAI64NARUXTime{Seconds: 8640}, "invalid"}, "want": []interface{}{"", ErrInvalidFormat}},
	}

	for _, c := range cases {
		out, err := FromInternal("jdc", c["in"][0].(TAI64NARUXTime), c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("JDCFromInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(string) {
			t.Errorf("JDCFromInternal(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0], c["in"][1].(string), out, c["want"][0].(string))
		}
	}
}

func TestJDCOffset(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{}, 1}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, 1.0}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, "1"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, []byte("1")}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, *big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},

		{"in": []interface{}{TAI64NARUXTime{}, TAI64NARUXTime{Seconds: 1}}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := Offset("jdc", c["in"][0].(TAI64NARUXTime), c["in"][1])
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("JDCOffset(%#v, %#v) gave error %#v; want %#v", c["in"][0].(TAI64NARUXTime), c["in"][1], err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("JDCOffset(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0].(TAI64NARUXTime), c["in"][1], out, c["want"][0].(TAI64NARUXTime))
		}
	}
}
