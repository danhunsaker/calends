package calends

import (
	"strconv"
	"testing"
)

func TestGetTimesByMode(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	x, y := getTimesByMode(test1, test2, "start")
	if x.String() != "0" && y.String() != "0" {
		t.Errorf("start fail...\ngot  " + x.String() + "," + y.String() + "\nwant 0,0")
	}
	x, y = getTimesByMode(test1, test2, "start-end")
	if x.String() != "0" && y.String() != "86400" {
		t.Errorf("start-end fail...\ngot  " + x.String() + "," + y.String() + "\nwant 0,86400")
	}
	x, y = getTimesByMode(test1, test2, "end-start")
	if x.String() != "0" && y.String() != "0" {
		t.Errorf("end-start fail...\ngot  " + x.String() + "," + y.String() + "\nwant 0,0")
	}
	x, y = getTimesByMode(test1, test2, "end")
	if x.String() != "0" && y.String() != "86400" {
		t.Errorf("end fail...\ngot  " + x.String() + "," + y.String() + "\nwant 0,86400")
	}
}

func TestCompareTimesByMode(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	cmp := compareTimesByMode(test1, test2, "start")
	if cmp != 0 {
		t.Errorf("1,2 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test2, test1, "start")
	if cmp != 0 {
		t.Errorf("2,1 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test1, test1, "start")
	if cmp != 0 {
		t.Errorf("1,1 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test2, test2, "start")
	if cmp != 0 {
		t.Errorf("2,2 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}

	cmp = compareTimesByMode(test1, test2, "start-end")
	if cmp != -1 {
		t.Errorf("1,2 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant -1")
	}
	cmp = compareTimesByMode(test2, test1, "start-end")
	if cmp != 0 {
		t.Errorf("2,1 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test1, test1, "start-end")
	if cmp != 0 {
		t.Errorf("1,1 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test2, test2, "start-end")
	if cmp != -1 {
		t.Errorf("2,2 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant -1")
	}

	cmp = compareTimesByMode(test1, test2, "end-start")
	if cmp != 0 {
		t.Errorf("1,2 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test2, test1, "end-start")
	if cmp != 1 {
		t.Errorf("2,1 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 1")
	}
	cmp = compareTimesByMode(test1, test1, "end-start")
	if cmp != 0 {
		t.Errorf("1,1 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test2, test2, "end-start")
	if cmp != 1 {
		t.Errorf("2,2 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 1")
	}

	cmp = compareTimesByMode(test1, test2, "end")
	if cmp != -1 {
		t.Errorf("1,2 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant -1")
	}
	cmp = compareTimesByMode(test2, test1, "end")
	if cmp != 1 {
		t.Errorf("2,1 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 1")
	}
	cmp = compareTimesByMode(test1, test1, "end")
	if cmp != 0 {
		t.Errorf("1,1 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = compareTimesByMode(test2, test2, "end")
	if cmp != 0 {
		t.Errorf("2,2 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
}

func TestDifference(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	got := test1.Difference(test2, "start")
	if got.String() != "0" {
		t.Errorf("1,2 start fail...\ngot  " + got.String() + "\nwant 0")
	}
	got = test2.Difference(test1, "start")
	if got.String() != "0" {
		t.Errorf("2,1 start fail...\ngot  " + got.String() + "\nwant 0")
	}
	got = test1.Difference(test2, "start-end")
	if got.String() != "-86400" {
		t.Errorf("1,2 start-end fail...\ngot  " + got.String() + "\nwant -86400")
	}
	got = test2.Difference(test1, "start-end")
	if got.String() != "0" {
		t.Errorf("2,1 start-end fail...\ngot  " + got.String() + "\nwant 0")
	}
	got = test1.Difference(test2, "end-start")
	if got.String() != "0" {
		t.Errorf("1,2 end-start fail...\ngot  " + got.String() + "\nwant 0")
	}
	got = test2.Difference(test1, "end-start")
	if got.String() != "86400" {
		t.Errorf("2,1 end-start fail...\ngot  " + got.String() + "\nwant 86400")
	}
	got = test1.Difference(test2, "end")
	if got.String() != "-86400" {
		t.Errorf("1,2 end fail...\ngot  " + got.String() + "\nwant -86400")
	}
	got = test2.Difference(test1, "end")
	if got.String() != "86400" {
		t.Errorf("2,1 end fail...\ngot  " + got.String() + "\nwant 86400")
	}
	got = test1.Difference(test2, "duration")
	if got.String() != "-86400" {
		t.Errorf("1,2 duration fail...\ngot  " + got.String() + "\nwant -86400")
	}
	got = test2.Difference(test1, "duration")
	if got.String() != "86400" {
		t.Errorf("2,1 duration fail...\ngot  " + got.String() + "\nwant 86400")
	}
}

func TestCompare(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	cmp := test1.Compare(test2, "start")
	if cmp != 0 {
		t.Errorf("1,2 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test2.Compare(test1, "start")
	if cmp != 0 {
		t.Errorf("2,1 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test1.Compare(test1, "start")
	if cmp != 0 {
		t.Errorf("1,1 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test2.Compare(test2, "start")
	if cmp != 0 {
		t.Errorf("2,2 start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}

	cmp = test1.Compare(test2, "start-end")
	if cmp != -1 {
		t.Errorf("1,2 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant -1")
	}
	cmp = test2.Compare(test1, "start-end")
	if cmp != 0 {
		t.Errorf("2,1 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test1.Compare(test1, "start-end")
	if cmp != 0 {
		t.Errorf("1,1 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test2.Compare(test2, "start-end")
	if cmp != -1 {
		t.Errorf("2,2 start-end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant -1")
	}

	cmp = test1.Compare(test2, "end-start")
	if cmp != 0 {
		t.Errorf("1,2 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test2.Compare(test1, "end-start")
	if cmp != 1 {
		t.Errorf("2,1 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 1")
	}
	cmp = test1.Compare(test1, "end-start")
	if cmp != 0 {
		t.Errorf("1,1 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test2.Compare(test2, "end-start")
	if cmp != 1 {
		t.Errorf("2,2 end-start fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 1")
	}

	cmp = test1.Compare(test2, "end")
	if cmp != -1 {
		t.Errorf("1,2 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant -1")
	}
	cmp = test2.Compare(test1, "end")
	if cmp != 1 {
		t.Errorf("2,1 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 1")
	}
	cmp = test1.Compare(test1, "end")
	if cmp != 0 {
		t.Errorf("1,1 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
	cmp = test2.Compare(test2, "end")
	if cmp != 0 {
		t.Errorf("2,2 end fail...\ngot  " + strconv.Itoa(cmp) + "\nwant 0")
	}
}

func TestIsSame(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.IsSame(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.IsSame(test1) {
		t.Errorf("2,1 failed")
	}
	if !test1.IsSame(test1) {
		t.Errorf("1,1 failed")
	}
	if !test2.IsSame(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestIsDuring(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if !test1.IsDuring(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.IsDuring(test1) {
		t.Errorf("2,1 failed")
	}
	if !test1.IsDuring(test1) {
		t.Errorf("1,1 failed")
	}
	if !test2.IsDuring(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestStartsDuring(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if !test1.StartsDuring(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.StartsDuring(test1) {
		t.Errorf("2,1 failed")
	}
	if !test1.StartsDuring(test1) {
		t.Errorf("1,1 failed")
	}
	if !test2.StartsDuring(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestEndsDuring(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if !test1.EndsDuring(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.EndsDuring(test1) {
		t.Errorf("2,1 failed")
	}
	if !test1.EndsDuring(test1) {
		t.Errorf("1,1 failed")
	}
	if !test2.EndsDuring(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestContains(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.Contains(test2) {
		t.Errorf("1,2 failed")
	}
	if !test2.Contains(test1) {
		t.Errorf("2,1 failed")
	}
	if !test1.Contains(test1) {
		t.Errorf("1,1 failed")
	}
	if !test2.Contains(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestOverlaps(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if !test1.Overlaps(test2) {
		t.Errorf("1,2 failed")
	}
	if !test2.Overlaps(test1) {
		t.Errorf("2,1 failed")
	}
	if !test1.Overlaps(test1) {
		t.Errorf("1,1 failed")
	}
	if !test2.Overlaps(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestAbuts(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.Abuts(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.Abuts(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.Abuts(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.Abuts(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestIsBefore(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.IsBefore(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.IsBefore(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.IsBefore(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.IsBefore(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestStartsBefore(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.StartsBefore(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.StartsBefore(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.StartsBefore(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.StartsBefore(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestEndsBefore(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if !test1.EndsBefore(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.EndsBefore(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.EndsBefore(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.EndsBefore(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestIsAfter(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.IsAfter(test2) {
		t.Errorf("1,2 failed")
	}
	if !test2.IsAfter(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.IsAfter(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.IsAfter(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestStartsAfter(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.StartsAfter(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.StartsAfter(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.StartsAfter(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.StartsAfter(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestEndsAfter(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.EndsAfter(test2) {
		t.Errorf("1,2 failed")
	}
	if !test2.EndsAfter(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.EndsAfter(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.EndsAfter(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestIsShorter(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if !test1.IsShorter(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.IsShorter(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.IsShorter(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.IsShorter(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestIsSameDuration(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.IsSameDuration(test2) {
		t.Errorf("1,2 failed")
	}
	if test2.IsSameDuration(test1) {
		t.Errorf("2,1 failed")
	}
	if !test1.IsSameDuration(test1) {
		t.Errorf("1,1 failed")
	}
	if !test2.IsSameDuration(test2) {
		t.Errorf("2,2 failed")
	}
}

func TestIsLonger(t *testing.T) {
	test1 := testValue(0)
	test2 := testValue(86400)

	if test1.IsLonger(test2) {
		t.Errorf("1,2 failed")
	}
	if !test2.IsLonger(test1) {
		t.Errorf("2,1 failed")
	}
	if test1.IsLonger(test1) {
		t.Errorf("1,1 failed")
	}
	if test2.IsLonger(test2) {
		t.Errorf("2,2 failed")
	}
}
