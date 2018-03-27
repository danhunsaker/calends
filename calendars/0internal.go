package calendars

import (
	"encoding/hex"
	"math"
	"math/big"
)

// TAI64NAXURTime stores a TAI64NAXUR instant in a reliable, easy-converted
// format.
type TAI64NAXURTime struct {
	Seconds int64  // Seconds since 1970-01-01 00:00:00 TAI
	Nano    uint32 // Nanoseconds since the given second
	Atto    uint32 // Attoseconds since the given nanosecond
	Xicto   uint32 // Xictoseconds since the given attosecond
	Ucto    uint32 // Uctoseconds since the given xictosecond
	Rocto   uint32 // Roctoseconds since the given uctosecond
}

// Add calculates the sum of two TAI64NAXURTime values.
func (a TAI64NAXURTime) Add(b TAI64NAXURTime) TAI64NAXURTime {
	var c TAI64NAXURTime
	var roll int32

	roll, c.Rocto = rollOverAt9(int32(a.Rocto + b.Rocto))
	roll, c.Ucto = rollOverAt9(int32(a.Ucto+b.Ucto) + roll)
	roll, c.Xicto = rollOverAt9(int32(a.Xicto+b.Xicto) + roll)
	roll, c.Atto = rollOverAt9(int32(a.Atto+b.Atto) + roll)
	roll, c.Nano = rollOverAt9(int32(a.Nano+b.Nano) + roll)
	c.Seconds = a.Seconds + b.Seconds + int64(roll)

	return c
}

// Sub calculates the difference of two TAI64NAXURTime values.
func (a TAI64NAXURTime) Sub(b TAI64NAXURTime) TAI64NAXURTime {
	var c TAI64NAXURTime
	var roll int32

	roll, c.Rocto = rollOverAt9(int32(a.Rocto) - int32(b.Rocto))
	roll, c.Ucto = rollOverAt9(int32(a.Ucto) - int32(b.Ucto) - roll)
	roll, c.Xicto = rollOverAt9(int32(a.Xicto) - int32(b.Xicto) - roll)
	roll, c.Atto = rollOverAt9(int32(a.Atto) - int32(b.Atto) - roll)
	roll, c.Nano = rollOverAt9(int32(a.Nano) - int32(b.Nano) - roll)
	c.Seconds = a.Seconds - b.Seconds - int64(roll)

	return c
}

// String returns the decimal string representation of the TAI64NAXURTime value.
func (t TAI64NAXURTime) String() string {
	out, _ := FromInternal("tai64", t, "decimal")
	return out
}

// HexString returns the hex string representation of the TAI64NAXURTime value.
func (t TAI64NAXURTime) HexString() string {
	out, _ := FromInternal("tai64", t, "tai64naxur")
	return out
}

// Float returns the math/big.Float representation of the TAI64NAXURTime value.
func (t TAI64NAXURTime) Float() *big.Float {
	out, _, _ := big.ParseFloat(t.String(), 10, 176, big.ToNearestAway)
	return out
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t TAI64NAXURTime) MarshalText() ([]byte, error) {
	out, err := FromInternal("tai64", t, "tai64naxur")
	return []byte(out), err
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *TAI64NAXURTime) UnmarshalText(in []byte) error {
	tmp, err := ToInternal("tai64", in, "tai64naxur")
	*t = tmp
	return err
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t *TAI64NAXURTime) MarshalBinary() (out []byte, err error) {
	in, err := t.MarshalText()
	if err != nil {
		return
	}

	out, err = hex.DecodeString(string(in))

	return
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (t *TAI64NAXURTime) UnmarshalBinary(in []byte) error {
	out := hex.EncodeToString(in)

	return t.UnmarshalText([]byte(out))
}

// TAI64NAXURTimeFromDecimalString calculates a TAI64NAXURTime from its decimal
// string representation.
func TAI64NAXURTimeFromDecimalString(in string) TAI64NAXURTime {
	out, _ := ToInternal("tai64", in, "decimal")
	// fmt.Printf("TAI64NAXURTimeFromDecimalString: %#v → %#v [%#v]\n", in, out, err)
	return out
}

// TAI64NAXURTimeFromHexString calculates a TAI64NAXURTime from its hexadecimal
// string representation.
func TAI64NAXURTimeFromHexString(in string) TAI64NAXURTime {
	out, _ := ToInternal("tai64", in, "tai64naxur")
	// fmt.Printf("TAI64NAXURTimeFromHexString: %#v → %#v [%#v]\n", in, out, err)
	return out
}

// TAI64NAXURTimeFromFloat calculates a TAI64NAXURTime from its math/big.Float
// representation.
func TAI64NAXURTimeFromFloat(in big.Float) TAI64NAXURTime {
	// fmt.Printf("TAI64NAXURTimeFromFloat: %#v\n", in)
	return TAI64NAXURTimeFromDecimalString(in.Text('f', 45))
}

func rollOverAt9(value int32) (roll int32, remain uint32) {
	working := math.Abs(float64(value))

	roll = int32(working) / 1000000000
	remain = uint32(math.Mod(working, 1000000000))

	return
}
