package calends

import (
	"math/big"

	"github.com/danhunsaker/calends/calendars"
)

func getTimesByMode(a, b Calends, mode string) (x, y calendars.TAI64NAXURTime) {
	switch mode {
	case "end":
		x = a.endTime
		y = b.endTime
	case "end-start":
		x = a.endTime
		y = b.startTime
	case "start-end":
		x = a.startTime
		y = b.endTime
	default:
		x = a.startTime
		y = b.startTime
	}

	return
}

func compareTimesByMode(a, b Calends, mode string) int {
	var x, y *big.Float
	if mode == "duration" {
		x = a.duration
		y = b.duration
	} else {
		tmpX, tmpY := getTimesByMode(a, b, mode)
		x = tmpX.Float()
		y = tmpY.Float()
	}

	return x.Cmp(y)
}

// Difference retrieves the difference between the current Calends object and
// another.
func (a Calends) Difference(b Calends, mode string) (out big.Float) {
	if mode == "duration" {
		out.Sub(a.duration, b.duration)
	} else {
		x, y := getTimesByMode(a, b, mode)
		out = *x.Sub(y).Float()
	}

	return
}

// Compares the start time, end time, or duration of two Calends objects.
func (a Calends) Compare(b Calends, mode string) int {
	return compareTimesByMode(a, b, mode)
}

// IsSame checks whether the current object has the same value(s) as another.
func (a Calends) IsSame(b Calends) bool {
	return a.Compare(b, "start") == 0 && a.Compare(b, "end") == 0
}

// IsDuring checks whether the current object fits entirely within another.
func (a Calends) IsDuring(b Calends) bool {
	return a.Compare(b, "start") >= 0 && a.Compare(b, "end") <= 0
}

// StartsDuring checks whether the current object has its start point between
// another's start and end points.
func (a Calends) StartsDuring(b Calends) bool {
	tmp, _ := a.duration.Int64()
	return a.Compare(b, "start") >= 0 && (a.Compare(b, "start-end") < 0 || (tmp == 0 && a.Compare(b, "start-end") == 0))
}

// EndsDuring checks whether the current object has its end point between
// another's start and end points.
func (a Calends) EndsDuring(b Calends) bool {
	tmp, _ := a.duration.Int64()
	return a.Compare(b, "end") <= 0 && (a.Compare(b, "end-start") > 0 || (tmp == 0 && a.Compare(b, "end-start") == 0))
}

// Contains checks whether another object fits entirely within the current one.
func (a Calends) Contains(b Calends) bool {
	return a.Compare(b, "start") <= 0 && a.Compare(b, "end") >= 0
}

// Overlaps checks whether either of the current object's endpoints occur within
// another object's period, or vice-versa.
func (a Calends) Overlaps(b Calends) bool {
	return a.StartsDuring(b) || b.StartsDuring(a) || a.EndsDuring(b) || b.EndsDuring(a)
}

// Abuts checks whether the current object starts when another object ends, or
// vice-versa, and that neither contains the other.
func (a Calends) Abuts(b Calends) bool {
	return (a.Compare(b, "start-end") == 0 && a.Compare(b, "end-start") == 0) && !(a.Contains(b) || b.Contains(a))
}

// IsBefore checks whether both of the current object's endpoints occur before
// another object's start point.
func (a Calends) IsBefore(b Calends) bool {
	return a.Compare(b, "end-start") <= 0 && a.Compare(b, "start") < 0
}

// StartsBefore checks whether the current object's start point occurs before
// another's.
func (a Calends) StartsBefore(b Calends) bool {
	return a.Compare(b, "start") < 0
}

// EndsBefore checks whether the current object's end point occurs before
// another's.
func (a Calends) EndsBefore(b Calends) bool {
	return a.Compare(b, "end") < 0
}

// IsAfter checks whether both of the current object's endpoints occur after
// another object's end point.
func (a Calends) IsAfter(b Calends) bool {
	return a.Compare(b, "start-end") >= 0 && a.Compare(b, "end") > 0
}

// StartsAfter checks whether the current object's start point occurs after
// another's.
func (a Calends) StartsAfter(b Calends) bool {
	return a.Compare(b, "start") > 0
}

// EndsAfter checks whether the current object's end point occurs after
// another's.
func (a Calends) EndsAfter(b Calends) bool {
	return a.Compare(b, "end") > 0
}

// IsShorter checks whether the current object's duration is less than
// another's.
func (a Calends) IsShorter(b Calends) bool {
	return a.Compare(b, "duration") < 0
}

// IsSameDuration checks whether the current object's duration is equal to
// another's.
func (a Calends) IsSameDuration(b Calends) bool {
	return a.Compare(b, "duration") == 0
}

// IsLonger checks whether the current object's duration is greater than
// another's.
func (a Calends) IsLonger(b Calends) bool {
	return a.Compare(b, "duration") > 0
}
