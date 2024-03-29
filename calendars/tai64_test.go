package calendars

import (
	"math/big"
	"testing"
)

func TestTai64ToInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{"1", "decimal"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"1.", "decimal"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"1.0", "decimal"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"40000000000000010000000000000000000000000000000000000000", "tai64narux"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"400000000000000100000000000000000000000000000000", "tai64naru"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"4000000000000001000000000000000000000000", "tai64nar"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"40000000000000010000000000000000", "tai64na"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"400000000000000100000000", "tai64n"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"4000000000000001", "tai64"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{"4000000000000001", "invalid"}, "want": []interface{}{TAI64NARUXTime{}, ErrInvalidFormat}},

		{"in": []interface{}{[]byte("40000000000000010000000000000000000000000000000000000000"), "tai64narux"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "tai64narux"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{1, "tai64narux"}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := ToInternal("tai64", c["in"][0], c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("Tai64ToInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("Tai64ToInternal(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0], c["in"][1].(string), out, c["want"][0].(TAI64NARUXTime))
		}
	}
}

func TestTai64FromInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "decimal"}, "want": []interface{}{"1", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "tai64narux"}, "want": []interface{}{"40000000000000010000000000000000000000000000000000000000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "tai64naru"}, "want": []interface{}{"400000000000000100000000000000000000000000000000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "tai64nar"}, "want": []interface{}{"4000000000000001000000000000000000000000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "tai64na"}, "want": []interface{}{"40000000000000010000000000000000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "tai64n"}, "want": []interface{}{"400000000000000100000000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "tai64"}, "want": []interface{}{"4000000000000001", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, "invalid"}, "want": []interface{}{"", ErrInvalidFormat}},
	}

	for _, c := range cases {
		out, err := FromInternal("tai64", c["in"][0].(TAI64NARUXTime), c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("Tai64FromInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(string) {
			t.Errorf("Tai64FromInternal(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0], c["in"][1].(string), out, c["want"][0].(string))
		}
	}
}

func TestTai64Offset(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{}, "1"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, []byte("1")}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, *big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, TAI64NARUXTime{Seconds: 1}}, "want": []interface{}{TAI64NARUXTime{Seconds: 1}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, 1}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := Offset("tai64", c["in"][0].(TAI64NARUXTime), c["in"][1])
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("Tai64Offset(%#v, %#v) gave error %#v; want %#v", c["in"][0].(TAI64NARUXTime), c["in"][1], err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("Tai64Offset(%#v, %#v)\nreturned %#v\nwanted   %#v", c["in"][0].(TAI64NARUXTime), c["in"][1], out, c["want"][0].(TAI64NARUXTime))
		}
	}
}
