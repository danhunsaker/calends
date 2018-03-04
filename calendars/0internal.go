// TAI ↔ UTC CONVERSIONS
/*

Part of what the helper methods offer is automatic conversion between UTC and
TAI counts. The difference between the two is the inclusion of leap seconds in
UTC, to keep noon synchronized with the solar zenith in the times used in
day-to-day situations. TAI counts do not include such an offset, and as such
there is a difference of a few seconds between the two values, which varies
based on when the instant actually occurs.

In order to perform this conversion, Calends makes use of leap second tables
available from the Centre de la Rotation de la Terre du Service International de
la Rotation Terrestre et des Systèmes de Référence à L’Observatoire de Paris (in
English, the Earth Orientation Center, at the International Earth Rotation and
Reference Systems Service, itself located in the Paris Observatory). The second
of these tables is the full-second offset between TAI and UTC, starting in 1972
(when the leap second system itself was devised). The first, then, is the
calculation used to find the offset from 1961 through 1971, which itself is
based on the actual differences between Earth's rotation-based time and the
actual number of seconds elapsed in that period. Times prior to this period are
unadjusted until a more comprehensive offset can reliably be calculated into the
past. Times after the latest offset (generally applicable to the present moment)
are similarly adjusted only to the latest offset, until a decent approximation
can be devised. These tables are updated periodically, and any updates will be
reflected by subsequent releases of the Calends library.

NOTE: The TAI ↔ UTC conversion feature has not yet been implemented.

*/
package calendars

import (
	"encoding/hex"
	// "fmt"
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

// Float returns the math.big.Float representation of the TAI64NAXURTime value.
func (t TAI64NAXURTime) Float() *big.Float {
	out, _, _ := big.ParseFloat(t.String(), 10, 176, big.ToNearestAway)
	return out
}

// TAI64NAXURTimeFromFloat calculates a TAI64NAXURTime from its math.big.Float
// representation.
func TAI64NAXURTimeFromFloat(in big.Float) TAI64NAXURTime {
	// fmt.Printf("TAI64NAXURTimeFromFloat: %#v\n", in)
	return TAI64NAXURTimeFromDecimalString(in.Text('f', 45))
}

func rollOverAt9(value int32) (roll int32, remain uint32) {
	if value >= 0 {
		roll = value / 1000000000
	} else {
		roll = 1
	}

	remain = uint32(math.Mod(float64(value), 1000000000))

	return
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t TAI64NAXURTime) MarshalText() ([]byte, error) {
	out, err := FromInternal("tai64", t, "tai64naxur")
	return []byte(out), err
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *TAI64NAXURTime) UnmarshalText(in []byte) error {
	tmp, err := ToInternal("tai64", in, "tai64naxur")
	t = &tmp
	return err
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t *TAI64NAXURTime) MarshalBinary() (out []byte, err error) {
	in, err := t.MarshalText()
	if err != nil {
		return
	}

	_, err = hex.Decode(out, in)

	return
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (t *TAI64NAXURTime) UnmarshalBinary(in []byte) error {
	var out []byte

	_ = hex.Encode(out, in)

	return t.UnmarshalText(out)
}

// UTCtoTAI removes the UTC leap second offset from a TAI64NAXURTime value.
func UTCtoTAI(utc TAI64NAXURTime) (tai TAI64NAXURTime, err error) {
	// Calculate year, month, day
	var year, month, day int
	// Remove the leap second offset
	offset, err := getTAIOffset(year, month, day)
	tai = utc.Add(offset)

	return
}

// TAItoUTC adds the UTC leap second offset to a TAI64NAXURTime value.
func TAItoUTC(tai TAI64NAXURTime) (utc TAI64NAXURTime, err error) {
	// Calculate year, month, day
	var year, month, day int
	// Add the leap second offset
	offset, err := getTAIOffset(year, month, day)
	utc = tai.Sub(offset)

	return
}

func getTAIOffset(year, month, day int) (offset TAI64NAXURTime, err error) {
	// Look up the offset
	return
}
