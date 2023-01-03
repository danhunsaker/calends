package calends

import (
	"math/big"

	"github.com/danhunsaker/calends/calendars"
)

func getTimesByMode(c, z Calends, mode string) (x, y calendars.TAI64NARUXTime) {
	switch mode {
	case "end":
		x = c.endTime
		y = z.endTime
	case "end-start":
		x = c.endTime
		y = z.startTime
	case "start-end":
		x = c.startTime
		y = z.endTime
	default:
		x = c.startTime
		y = z.startTime
	}

	return
}

func compareTimesByMode(c, z Calends, mode string) int {
	var x, y *big.Float
	if mode == "duration" {
		x = c.duration
		y = z.duration
	} else {
		tmpX, tmpY := getTimesByMode(c, z, mode)
		x = tmpX.Float()
		y = tmpY.Float()
	}

	return x.Cmp(y)
}

// Difference retrieves the difference between the current Calends object and
// another.
func (c Calends) Difference(z Calends, mode string) (out big.Float) {
	if mode == "duration" {
		out.Sub(c.duration, z.duration)
	} else {
		x, y := getTimesByMode(c, z, mode)
		out = *x.Sub(y).Float()
	}

	return
}

// Compare the start time, end time, some mix of those two, or duration of two
// Calends objects.
func (c Calends) Compare(z Calends, mode string) int {
	return compareTimesByMode(c, z, mode)
}

// IsSame checks whether the current object has the same value(s) as another.
func (c Calends) IsSame(z Calends) bool {
	return c.Compare(z, "start") == 0 && c.Compare(z, "end") == 0
}

// IsDuring checks whether the current object fits entirely within another.
func (c Calends) IsDuring(z Calends) bool {
	return c.Compare(z, "start") >= 0 && c.Compare(z, "end") <= 0
}

// StartsDuring checks whether the current object has its start point between
// another's start and end points.
func (c Calends) StartsDuring(z Calends) bool {
	tmp, _ := c.duration.Int64()
	return c.Compare(z, "start") >= 0 && (c.Compare(z, "start-end") < 0 || (tmp == 0 && c.Compare(z, "start-end") == 0))
}

// EndsDuring checks whether the current object has its end point between
// another's start and end points.
func (c Calends) EndsDuring(z Calends) bool {
	tmp, _ := c.duration.Int64()
	return c.Compare(z, "end") <= 0 && (c.Compare(z, "end-start") > 0 || (tmp == 0 && c.Compare(z, "end-start") == 0))
}

// Contains checks whether another object fits entirely within the current one.
func (c Calends) Contains(z Calends) bool {
	return c.Compare(z, "start") <= 0 && c.Compare(z, "end") >= 0
}

// Overlaps checks whether either of the current object's endpoints occur within
// another object's period, or vice-versa.
func (c Calends) Overlaps(z Calends) bool {
	return c.StartsDuring(z) || z.StartsDuring(c) || c.EndsDuring(z) || z.EndsDuring(c)
}

// Abuts checks whether the current object starts when another object ends, or
// vice-versa, and that neither contains the other.
func (c Calends) Abuts(z Calends) bool {
	return ((c.Compare(z, "start-end") == 0 || c.Compare(z, "end-start") == 0) && !(c.Contains(z) || z.Contains(c)))
}

// IsBefore checks whether both of the current object's endpoints occur before
// another object's start point.
func (c Calends) IsBefore(z Calends) bool {
	return c.Compare(z, "end-start") <= 0 && c.Compare(z, "start") < 0
}

// StartsBefore checks whether the current object's start point occurs before
// another's.
func (c Calends) StartsBefore(z Calends) bool {
	return c.Compare(z, "start") < 0
}

// EndsBefore checks whether the current object's end point occurs before
// another's.
func (c Calends) EndsBefore(z Calends) bool {
	return c.Compare(z, "end") < 0
}

// IsAfter checks whether both of the current object's endpoints occur after
// another object's end point.
func (c Calends) IsAfter(z Calends) bool {
	return c.Compare(z, "start-end") >= 0 && c.Compare(z, "end") > 0
}

// StartsAfter checks whether the current object's start point occurs after
// another's.
func (c Calends) StartsAfter(z Calends) bool {
	return c.Compare(z, "start") > 0
}

// EndsAfter checks whether the current object's end point occurs after
// another's.
func (c Calends) EndsAfter(z Calends) bool {
	return c.Compare(z, "end") > 0
}

// IsShorter checks whether the current object's duration is less than
// another's.
func (c Calends) IsShorter(z Calends) bool {
	return c.Compare(z, "duration") < 0
}

// IsSameDuration checks whether the current object's duration is equal to
// another's.
func (c Calends) IsSameDuration(z Calends) bool {
	return c.Compare(z, "duration") == 0
}

// IsLonger checks whether the current object's duration is greater than
// another's.
func (c Calends) IsLonger(z Calends) bool {
	return c.Compare(z, "duration") > 0
}
