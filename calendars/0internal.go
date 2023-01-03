package calendars

import (
	"encoding/hex"
	"math"
	"math/big"
)

// TAI64NARUXTime stores a TAI64NARUX instant in a reliable, easy-converted
// format.
type TAI64NARUXTime struct {
	Seconds  int64  // Seconds since 1970-01-01 00:00:00 TAI
	Nano     uint32 // billionths of a second since the given second
	Atto     uint32 // billionths of a nanosecond since the given nanosecond
	Ronto    uint32 // billionths of an attosecond since the given attosecond
	Udecto   uint32 // billionths of a rontosecond since the given rontosecond
	Xindecto uint32 // billionths of an udectosecond since the given udectosecond
}

// Add calculates the sum of two TAI64NARUXTime values.
func (t TAI64NARUXTime) Add(z TAI64NARUXTime) TAI64NARUXTime {
	var o TAI64NARUXTime
	var roll int32

	roll, o.Xindecto = rollOverAt9(int32(t.Xindecto + z.Xindecto))
	roll, o.Udecto = rollOverAt9(int32(t.Udecto+z.Udecto) + roll)
	roll, o.Ronto = rollOverAt9(int32(t.Ronto+z.Ronto) + roll)
	roll, o.Atto = rollOverAt9(int32(t.Atto+z.Atto) + roll)
	roll, o.Nano = rollOverAt9(int32(t.Nano+z.Nano) + roll)
	o.Seconds = t.Seconds + z.Seconds + int64(roll)

	return o
}

// Sub calculates the difference of two TAI64NARUXTime values.
func (t TAI64NARUXTime) Sub(z TAI64NARUXTime) TAI64NARUXTime {
	var o TAI64NARUXTime
	var roll int32

	roll, o.Xindecto = rollOverAt9(int32(t.Xindecto) - int32(z.Xindecto))
	roll, o.Udecto = rollOverAt9(int32(t.Udecto) - int32(z.Udecto) - roll)
	roll, o.Ronto = rollOverAt9(int32(t.Ronto) - int32(z.Ronto) - roll)
	roll, o.Atto = rollOverAt9(int32(t.Atto) - int32(z.Atto) - roll)
	roll, o.Nano = rollOverAt9(int32(t.Nano) - int32(z.Nano) - roll)
	o.Seconds = t.Seconds - z.Seconds - int64(roll)

	return o
}

// String returns the decimal string representation of the TAI64NARUXTime value.
func (t TAI64NARUXTime) String() string {
	out, _ := FromInternal("tai64", t, "decimal")
	return out
}

// HexString returns the hex string representation of the TAI64NARUXTime value.
func (t TAI64NARUXTime) HexString() string {
	out, _ := FromInternal("tai64", t, "tai64narux")
	return out
}

// Float returns the math/big.Float representation of the TAI64NARUXTime value.
func (t TAI64NARUXTime) Float() *big.Float {
	out, _, _ := big.ParseFloat(t.String(), 10, 176, big.ToNearestAway)
	return out
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t TAI64NARUXTime) MarshalText() ([]byte, error) {
	out, err := FromInternal("tai64", t, "tai64narux")
	return []byte(out), err
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *TAI64NARUXTime) UnmarshalText(in []byte) error {
	tmp, err := ToInternal("tai64", in, "tai64narux")
	*t = tmp
	return err
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t *TAI64NARUXTime) MarshalBinary() (out []byte, err error) {
	in, err := t.MarshalText()
	if err != nil {
		return
	}

	out, err = hex.DecodeString(string(in))

	return
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (t *TAI64NARUXTime) UnmarshalBinary(in []byte) error {
	out := hex.EncodeToString(in)

	return t.UnmarshalText([]byte(out))
}

// TAI64NARUXTimeFromDecimalString calculates a TAI64NARUXTime from its decimal
// string representation.
func TAI64NARUXTimeFromDecimalString(in string) TAI64NARUXTime {
	out, _ := ToInternal("tai64", in, "decimal")
	// fmt.Printf("TAI64NARUXTimeFromDecimalString: %#v → %#v [%#v]\n", in, out, err)
	return out
}

// TAI64NARUXTimeFromHexString calculates a TAI64NARUXTime from its hexadecimal
// string representation.
func TAI64NARUXTimeFromHexString(in string) TAI64NARUXTime {
	out, _ := ToInternal("tai64", in, "tai64narux")
	// fmt.Printf("TAI64NARUXTimeFromHexString: %#v → %#v [%#v]\n", in, out, err)
	return out
}

// TAI64NARUXTimeFromFloat calculates a TAI64NARUXTime from its math/big.Float
// representation.
func TAI64NARUXTimeFromFloat(in big.Float) TAI64NARUXTime {
	// fmt.Printf("TAI64NARUXTimeFromFloat: %#v\n", in)
	return TAI64NARUXTimeFromDecimalString(in.Text('f', 45))
}

func rollOverAt9(value int32) (roll int32, remain uint32) {
	working := math.Abs(float64(value))

	roll = int32(working) / 1000000000
	remain = uint32(math.Mod(working, 1000000000))

	return
}
