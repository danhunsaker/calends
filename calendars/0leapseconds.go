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
	"encoding/base64"
	"encoding/gob"
	"math/big"
	"strings"
	"time"
)

type leapSecondDate struct {
	Year  int
	Month time.Month
	Day   int
}

func (a leapSecondDate) Compare(b leapSecondDate) int {
	if a.Year > b.Year {
		return 1
	} else if a.Year < b.Year {
		return -1
	}

	// a.Year == b.Year
	if a.Month > b.Month {
		return 1
	} else if a.Month < b.Month {
		return -1
	}

	// a.Year == b.Year && a.Month == b.Month
	if a.Day > b.Day {
		return 1
	} else if a.Day < b.Day {
		return -1
	}

	// a.Year == b.Year && a.Month == b.Month && a.Day == b.Day
	return 0
}

type leapSecondOffsetModifier struct {
	Difference int
	Multiplier float64
}

func (m leapSecondOffsetModifier) Calculate(mjd int) float64 {
	return float64(mjd-m.Difference) * m.Multiplier
}

type leapSecondOffset struct {
	Start    leapSecondDate
	End      leapSecondDate
	Offset   float64
	Modifier leapSecondOffsetModifier
}

func (o leapSecondOffset) Within(d leapSecondDate) bool {
	if d.Compare(o.Start) >= 0 && (d.Compare(o.End) < 0 || o.End == o.Start) {
		return true
	}

	return false
}

var leapSeconds []leapSecondOffset

func init() {
	encoded := "Df+HAgEC/4gAAf+CAABB/4EDAQEGT2Zmc2V0Af+CAAEEAQVTdGFydAH/hAABA0V" +
		"uZAH/hAABBk9mZnNldAEIAAEITW9kaWZpZXIB/4YAAAAt/4MDAQEERGF0ZQH/hAABAwEEWWV" +
		"hcgEEAAEFTW9udGgBBAABA0RheQEEAAAAOv+FAwEBDk9mZnNldE1vZGlmaWVyAf+GAAECAQp" +
		"EaWZmZXJlbmNlAQQAAQpNdWx0aXBsaWVyAQgAAAD+BWv/iAApAQH+D1IBAgECAAEB/g9SARA" +
		"BAgAB+OKPos7cw/Y/AQH9ASNoAfinQGZn0TtVPwAAAQH+D1IBEAECAAEB/g9UAQIBAgAB+Bb" +
		"D1QEQ9/U/AQH9ASNoAfinQGZn0TtVPwAAAQH+D1QBAgECAAEB/g9WARYBAgAB+KH18GWiiP0" +
		"/AQH9ASZCAfgqWhTRCmdSPwAAAQH+D1YBFgECAAEB/g9YAQIBAgAB+DuPiv87Iv8/AQH9ASZ" +
		"CAfgqWhTRCmdSPwAAAQH+D1gBAgECAAEB/g9YAQgBAgAB+M9OBkfJ6wlAAQH9AS7SAfinQGZ" +
		"n0TtVPwAAAQH+D1gBCAECAAEB/g9YARIBAgAB+Jsb0xOWuApAAQH9AS7SAfinQGZn0TtVPwA" +
		"AAQH+D1gBEgECAAEB/g9aAQIBAgAB+Gjon+BihQtAAQH9AS7SAfinQGZn0TtVPwAAAQH+D1o" +
		"BAgECAAEB/g9aAQYBAgAB+DW1bK0vUgxAAQH9AS7SAfinQGZn0TtVPwAAAQH+D1oBBgECAAE" +
		"B/g9aAQ4BAgAB+AKCOXr8Hg1AAQH9AS7SAfinQGZn0TtVPwAAAQH+D1oBDgECAAEB/g9aARI" +
		"BAgAB+M9OBkfJ6w1AAQH9AS7SAfinQGZn0TtVPwAAAQH+D1oBEgECAAEB/g9cAQIBAgAB+Js" +
		"b0xOWuA5AAQH9AS7SAfinQGZn0TtVPwAAAQH+D1wBAgECAAEB/g9gAQQBAgAB+HFa8KKvQBF" +
		"AAQH9ATGsAfinQGZn0TtlPwAAAQH+D2ABBAECAAEB/g9oAQIBAgAB+Ar0iTxJ2hBAAQH9ATG" +
		"sAfinQGZn0TtlPwAAAQH+D2gBAgECAAEB/g9oAQ4BAgAB/iRAAQAAAQH+D2gBDgECAAEB/g9" +
		"qAQIBAgAB/iZAAQAAAQH+D2oBAgECAAEB/g9sAQIBAgAB/ihAAQAAAQH+D2wBAgECAAEB/g9" +
		"uAQIBAgAB/ipAAQAAAQH+D24BAgECAAEB/g9wAQIBAgAB/ixAAQAAAQH+D3ABAgECAAEB/g9" +
		"yAQIBAgAB/i5AAQAAAQH+D3IBAgECAAEB/g90AQIBAgAB/jBAAQAAAQH+D3QBAgECAAEB/g9" +
		"2AQIBAgAB/jFAAQAAAQH+D3YBAgECAAEB/g94AQIBAgAB/jJAAQAAAQH+D3gBAgECAAEB/g9" +
		"6AQ4BAgAB/jNAAQAAAQH+D3oBDgECAAEB/g98AQ4BAgAB/jRAAQAAAQH+D3wBDgECAAEB/g9" +
		"+AQ4BAgAB/jVAAQAAAQH+D34BDgECAAEB/g+CAQ4BAgAB/jZAAQAAAQH+D4IBDgECAAEB/g+" +
		"IAQIBAgAB/jdAAQAAAQH+D4gBAgECAAEB/g+MAQIBAgAB/jhAAQAAAQH+D4wBAgECAAEB/g+" +
		"OAQIBAgAB/jlAAQAAAQH+D44BAgECAAEB/g+QAQ4BAgAB/jpAAQAAAQH+D5ABDgECAAEB/g+" +
		"SAQ4BAgAB/jtAAQAAAQH+D5IBDgECAAEB/g+UAQ4BAgAB/jxAAQAAAQH+D5QBDgECAAEB/g+" +
		"YAQIBAgAB/j1AAQAAAQH+D5gBAgECAAEB/g+aAQ4BAgAB/j5AAQAAAQH+D5oBDgECAAEB/g+" +
		"eAQIBAgAB/j9AAQAAAQH+D54BAgECAAEB/g+sAQIBAgAB/kBAAQAAAQH+D6wBAgECAAEB/g+" +
		"yAQIBAgAB/YBAQAEAAAEB/g+yAQIBAgABAf4PuAEOAQIAAf5BQAEAAAEB/g+4AQ4BAgABAf4" +
		"PvgEOAQIAAf2AQUABAAABAf4PvgEOAQIAAQH+D8IBAgECAAH+QkABAAABAf4PwgECAQIAAQH" +
		"+D8IBAgECAAH9gEJAAQAA"

	gob.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(encoded))).Decode(&leapSeconds)
}

// UTCtoTAI removes the UTC leap second offset from a TAI64NAXURTime value.
func UTCtoTAI(utc TAI64NAXURTime) (tai TAI64NAXURTime) {
	// Calculate year, month, day
	oldYear, oldMonth, oldDay := time.Unix(utc.Seconds, int64(utc.Nano)).Date()
	// Remove the leap second offset
	tai = utc.Sub(getTAIOffset(oldYear, oldMonth, oldDay))

	// Ensure we used the correct offset
	newYear, newMonth, newDay := time.Unix(tai.Seconds, int64(tai.Nano)).Date()
	for newYear != oldYear || newMonth != oldMonth || newDay != oldDay {
		tai = utc.Sub(getTAIOffset(newYear, newMonth, newDay))
		oldYear, oldMonth, oldDay = newYear, newMonth, newDay
		newYear, newMonth, newDay = time.Unix(tai.Seconds, int64(tai.Nano)).Date()
	}

	return
}

// TAItoUTC adds the UTC leap second offset to a TAI64NAXURTime value.
func TAItoUTC(tai TAI64NAXURTime) (utc TAI64NAXURTime) {
	// Calculate year, month, day
	year, month, day := time.Unix(tai.Seconds, int64(tai.Nano)).Date()
	// Add the leap second offset
	utc = tai.Add(getTAIOffset(year, month, day))

	return
}

func getTAIOffset(year int, month time.Month, day int) (offset TAI64NAXURTime) {
	const baseDay = 40587 //Modified Julian Date at January 1, 1970

	var entry *leapSecondOffset

	// Look up the offset
	for _, ls := range leapSeconds {
		if ls.Within(leapSecondDate{year, month, day}) {
			entry = &ls
			break
		}
	}
	if entry == nil {
		return
	}

	timestamp := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
	mjd := int(timestamp/86400 + baseDay)

	offset = TAI64NAXURTimeFromFloat(*big.NewFloat(entry.Offset + entry.Modifier.Calculate(mjd)))

	return
}
