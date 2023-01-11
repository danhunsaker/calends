package calendars

import (
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {
	in := TAI64NARUXTime{Seconds: 1}
	and := TAI64NARUXTime{Seconds: 2}
	want := TAI64NARUXTime{Seconds: 3}
	got := in.Add(and)

	if got != want {
		t.Errorf("%#v.Add(%#v) failed\ngot  %#v\nwant %#v", in, and, got, want)
	}
}

func TestSub(t *testing.T) {
	in := TAI64NARUXTime{Seconds: 1}
	and := TAI64NARUXTime{Seconds: 2}
	want := TAI64NARUXTime{Seconds: -1}
	got := in.Sub(and)

	if got != want {
		t.Errorf("%#v.Sub(%#v) failed\ngot  %#v\nwant %#v", in, and, got, want)
	}
}

func TestString(t *testing.T) {
	in := TAI64NARUXTime{Seconds: 1}
	want := "1"
	got := in.String()

	if got != want {
		t.Errorf("%#v.String() failed\ngot  %#v\nwant %#v", in, got, want)
	}
}

func TestHexString(t *testing.T) {
	in := TAI64NARUXTime{Seconds: 1}
	want := "40000000000000010000000000000000000000000000000000000000"
	got := in.HexString()

	if got != want {
		t.Errorf("%#v.HexString() failed\ngot  %#v\nwant %#v", in, got, want)
	}
}

func TestFloat(t *testing.T) {
	in := TAI64NARUXTime{Seconds: 1}
	want := big.NewFloat(1)
	got := in.Float()

	if got.String() != want.String() {
		t.Errorf("%#v.Float() failed\ngot  %#v\nwant %#v", in, got, want)
	}
}

func TestMarshalText(t *testing.T) {
	in := TAI64NARUXTime{Seconds: 1}
	want := "40000000000000010000000000000000000000000000000000000000"
	got, _ := in.MarshalText()

	if string(got) != want {
		t.Errorf("%#v.MarshalText() failed\ngot  %#v\nwant %#v", in, string(got), want)
	}
}

func TestUnmarshalText(t *testing.T) {
	var got TAI64NARUXTime
	in := []byte("40000000000000010000000000000000000000000000000000000000")
	want := TAI64NARUXTime{Seconds: 1}
	got.UnmarshalText(in)

	if got != want {
		t.Errorf("empty.UnmarshalText(%#v) failed\ngot  %#v\nwant %#v", string(in), got, want)
	}
}

func TestMarshalBinary(t *testing.T) {
	in := TAI64NARUXTime{Seconds: 1}
	want := "@\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	got, _ := in.MarshalBinary()

	if string(got) != want {
		t.Errorf("%#v.MarshalBinary() failed\ngot  %#v\nwant %#v", in, string(got), want)
	}
}

func TestUnmarshalBinary(t *testing.T) {
	var got TAI64NARUXTime
	in := []byte("@\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	want := TAI64NARUXTime{Seconds: 1}
	got.UnmarshalBinary(in)

	if got != want {
		t.Errorf("empty.UnmarshalBinary(%#v) failed\ngot  %#v\nwant %#v", string(in), got, want)
	}
}

func TestTAI64NARUXTimeFromDecimalString(t *testing.T) {
	in := "1"
	want := TAI64NARUXTime{Seconds: 1}
	got := TAI64NARUXTimeFromDecimalString(in)

	if got != want {
		t.Errorf("TAI64NARUXTimeFromDecimalString(%#v) failed\ngot  %#v\nwant %#v", in, got, want)
	}
}

func TestTAI64NARUXTimeFromHexString(t *testing.T) {
	in := "40000000000000010000000000000000000000000000000000000000"
	want := TAI64NARUXTime{Seconds: 1}
	got := TAI64NARUXTimeFromHexString(in)

	if got != want {
		t.Errorf("TAI64NARUXTimeFromHexString(%#v) failed\ngot  %#v\nwant %#v", in, got, want)
	}
}

func TestTAI64NARUXTimeFromFloat(t *testing.T) {
	in := *big.NewFloat(1)
	want := TAI64NARUXTime{Seconds: 1}
	got := TAI64NARUXTimeFromFloat(in)

	if got != want {
		t.Errorf("TAI64NARUXTimeFromFloat(%#v) failed\ngot  %#v\nwant %#v", in, got, want)
	}
}

func TestRollOverAt9(t *testing.T) {
	type result struct {
		Roll   int32
		Remain uint32
	}
	var got result

	in := int32(1000000001)
	want := result{1, 1}
	got.Roll, got.Remain = rollOverAt9(in)

	if got != want {
		t.Errorf("rollOverAt9(%#v) failed\ngot  %#v\nwant %#v", in, got, want)
	}

	in = int32(-1000000000)
	want = result{1, 0}
	got.Roll, got.Remain = rollOverAt9(in)

	if got != want {
		t.Errorf("rollOverAt9(%#v) failed\ngot  %#v\nwant %#v", in, got, want)
	}
}
