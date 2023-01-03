package calendars

import (
	"math/big"
	"testing"
)

func TestStardateToInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{-352996, "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 83289, Nano: 600000000}, nil}},
		{"in": []interface{}{-352995.9024, "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-352996", "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 83289, Nano: 600000000}, nil}},
		{"in": []interface{}{"-352995.9024", "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{big.NewFloat(-352995.9024), "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{*big.NewFloat(-352995.9024), "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{[]byte("-352995.9024"), "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},

		{"in": []interface{}{"[-36]9355", ""}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"[-36]9355", "main"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-518956.8", "kennedy"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-353002.7397", "pugh90s"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-353004.3875", "pughfixed"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-352997.2603", "schmidt"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"23469.5414", "guide-equiv"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-320007.8084", "guide-tng"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-778176.1923", "guide-tos"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-352390.8245", "guide-oldtng"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-352995.9024", "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"-353230.64917", "red-dragon"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"47608.21918", "sto-hynes"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"47636.8359", "sto-academy"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"47608.48756", "sto-tom"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{"47608.1354", "sto-anthodev"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},

		{"in": []interface{}{"[-27]4920", "main"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-436814.4", "kennedy"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-307855.1913", "pugh90s"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-307858.1901", "pughfixed"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-306144.8087", "schmidt"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"70322.38193", "guide-equiv"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-276985.1188", "guide-tng"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-654711.2454", "guide-tos"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-305536.9834", "guide-oldtng"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-306142.0613", "aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"-306345.6497", "red-dragon"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"94460.67071", "sto-hynes"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"94458.0039", "sto-academy"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"94461.3283", "sto-tom"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},
		{"in": []interface{}{"94460.62959", "sto-anthodev"}, "want": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, nil}},

		{"in": []interface{}{"[21]00000", "main"}, "want": []interface{}{TAI64NARUXTime{Seconds: 11139552000}, nil}},
		{"in": []interface{}{"14814.24", "kennedy"}, "want": []interface{}{TAI64NARUXTime{Seconds: 11139552000}, nil}},

		{"in": []interface{}{"[20]5005.82", "main"}, "want": []interface{}{TAI64NARUXTime{Seconds: 11139520896}, nil}},

		{"in": []interface{}{"1.0", "invalid"}, "want": []interface{}{TAI64NARUXTime{}, ErrInvalidFormat}},
		{"in": []interface{}{"1512.90", "guide-oldtos"}, "want": []interface{}{TAI64NARUXTime{}, ErrInvalidFormat}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1}, ""}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := ToInternal("stardate", c["in"][0], c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("stardateToInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("stardateToInternal(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0], c["in"][1].(string), out, c["want"][0].(TAI64NARUXTime))
		}
	}
}

func TestStardateFromInternal(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, ""}, "want": []interface{}{"[-36]9355", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "main"}, "want": []interface{}{"[-36]9355", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "kennedy"}, "want": []interface{}{"-518956.8", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "pugh90s"}, "want": []interface{}{"-353002.7397", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "pughfixed"}, "want": []interface{}{"-353004.3875", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "schmidt"}, "want": []interface{}{"-352997.2603", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "guide-equiv"}, "want": []interface{}{"23469.54141", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "guide-tng"}, "want": []interface{}{"-320007.8084", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "guide-tos"}, "want": []interface{}{"-778176.1923", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "guide-oldtng"}, "want": []interface{}{"-352390.8245", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "aldrich"}, "want": []interface{}{"-352995.9024", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "red-dragon"}, "want": []interface{}{"-353230.6492", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "sto-hynes"}, "want": []interface{}{"47608.21918", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "sto-academy"}, "want": []interface{}{"47636.8359", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "sto-tom"}, "want": []interface{}{"47608.48756", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "sto-anthodev"}, "want": []interface{}{"47608.13541", nil}},

		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "main"}, "want": []interface{}{"[-27]4920", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "kennedy"}, "want": []interface{}{"-436814.4", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "pugh90s"}, "want": []interface{}{"-307855.1913", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "pughfixed"}, "want": []interface{}{"-307858.1901", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "schmidt"}, "want": []interface{}{"-306144.8087", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "guide-equiv"}, "want": []interface{}{"70322.38193", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "guide-tng"}, "want": []interface{}{"-276985.1188", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "guide-tos"}, "want": []interface{}{"-654711.2454", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "guide-oldtng"}, "want": []interface{}{"-305536.9834", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "aldrich"}, "want": []interface{}{"-306142.0613", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "red-dragon"}, "want": []interface{}{"-306345.6497", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "sto-hynes"}, "want": []interface{}{"94460.67071", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "sto-academy"}, "want": []interface{}{"94458.0039", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "sto-tom"}, "want": []interface{}{"94461.3283", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 1478649600}, "sto-anthodev"}, "want": []interface{}{"94460.62959", nil}},

		{"in": []interface{}{TAI64NARUXTime{Seconds: 11139552000}, "main"}, "want": []interface{}{"[21]00000", nil}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 11139552000}, "kennedy"}, "want": []interface{}{"14814.24", nil}},

		{"in": []interface{}{TAI64NARUXTime{Seconds: 11139520896}, "main"}, "want": []interface{}{"[20]5005.82", nil}},

		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "invalid"}, "want": []interface{}{"", ErrInvalidFormat}},
		{"in": []interface{}{TAI64NARUXTime{Seconds: 86400}, "guide-oldtos"}, "want": []interface{}{"", ErrInvalidFormat}},
	}

	for _, c := range cases {
		out, err := FromInternal("stardate", c["in"][0].(TAI64NARUXTime), c["in"][1].(string))
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("stardateFromInternal(%#v, %#v) gave error %#v; want %#v", c["in"][0], c["in"][1].(string), err, c["want"][1])
		}
		if out != c["want"][0].(string) {
			t.Errorf("stardateFromInternal(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0], c["in"][1].(string), out, c["want"][0].(string))
		}
	}
}

func TestStardateOffset(t *testing.T) {
	cases := []map[string][]interface{}{
		{"in": []interface{}{TAI64NARUXTime{}, "-352995.9024 aldrich"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},
		{"in": []interface{}{TAI64NARUXTime{}, []byte("-352995.9024 aldrich")}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},

		{"in": []interface{}{TAI64NARUXTime{}, "aldrich -352995.9024"}, "want": []interface{}{TAI64NARUXTime{Seconds: 86400}, nil}},

		{"in": []interface{}{TAI64NARUXTime{}, 1}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{TAI64NARUXTime{}, 1.0}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{TAI64NARUXTime{}, big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{TAI64NARUXTime{}, *big.NewFloat(1)}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
		{"in": []interface{}{TAI64NARUXTime{}, TAI64NARUXTime{Seconds: 1}}, "want": []interface{}{TAI64NARUXTime{}, ErrUnsupportedInput}},
	}

	for _, c := range cases {
		out, err := Offset("stardate", c["in"][0].(TAI64NARUXTime), c["in"][1])
		if (err != nil && c["want"][1] != nil && err.Error() != c["want"][1].(error).Error()) || (err == nil && err != c["want"][1]) || (c["want"][1] == nil && err != c["want"][1]) {
			t.Errorf("stardateOffset(%#v, %#v) gave error %#v; want %#v", c["in"][0].(TAI64NARUXTime), c["in"][1], err, c["want"][1])
		}
		if out != c["want"][0].(TAI64NARUXTime) {
			t.Errorf("stardateOffset(%#v, %#v)\nreturned %s\nwanted   %s", c["in"][0].(TAI64NARUXTime), c["in"][1], out, c["want"][0].(TAI64NARUXTime))
		}
	}
}
