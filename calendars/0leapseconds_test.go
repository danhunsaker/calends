package calendars

import (
	"testing"
	"time"
)

func TestLeapSecondDateCompare(t *testing.T) {
	cases := []map[string]interface{}{
		{"a": leapSecondDate{1969, time.June, 15}, "b": leapSecondDate{1970, time.June, 15}, "want": -1},
		{"a": leapSecondDate{1970, time.May, 15}, "b": leapSecondDate{1970, time.June, 15}, "want": -1},
		{"a": leapSecondDate{1970, time.June, 14}, "b": leapSecondDate{1970, time.June, 15}, "want": -1},

		{"a": leapSecondDate{1970, time.June, 15}, "b": leapSecondDate{1970, time.June, 15}, "want": 0},

		{"a": leapSecondDate{1971, time.June, 15}, "b": leapSecondDate{1970, time.June, 15}, "want": 1},
		{"a": leapSecondDate{1970, time.July, 15}, "b": leapSecondDate{1970, time.June, 15}, "want": 1},
		{"a": leapSecondDate{1970, time.June, 16}, "b": leapSecondDate{1970, time.June, 15}, "want": 1},
	}

	for _, c := range cases {
		out := c["a"].(leapSecondDate).Compare(c["b"].(leapSecondDate))
		if out != c["want"].(int) {
			t.Errorf("%#v.Compare(%#v)\nreturned %#v\nwanted   %#v", c["a"], c["b"], out, c["want"])
		}
	}
}

func TestUTCtoTAI(t *testing.T) {
	in := TAI64NAXURTime{Seconds: 1}
	want := TAI64NAXURTimeFromDecimalString("-6.997489999999999987778664944926276803016662598")
	got := UTCtoTAI(in)

	if got != want {
		t.Errorf("UTCtoTAI(%#v) failed\ngot  %s\nwant %s", in, got, want)
	}
}

func TestTAItoUTC(t *testing.T) {
	in := TAI64NAXURTime{Seconds: 1}
	want := TAI64NAXURTimeFromDecimalString("9.000081999999999027295416453853249549865722656")
	got := TAItoUTC(in)

	if got != want {
		t.Errorf("TAItoUTC(%#v) failed\ngot  %#v\nwant %#v", in, got, want)
	}
}

func TestGetTAIOffset(t *testing.T) {
	year, month, day := 1970, time.January, 1
	want := TAI64NAXURTimeFromDecimalString("8.000081999999999027295416453853249549865722656")
	got := getTAIOffset(year, month, day)

	if got != want {
		t.Errorf("getTAIOffset(%#v, %#v, %#v) failed\ngot  %#v\nwant %#v", year, month, day, got, want)
	}
}
