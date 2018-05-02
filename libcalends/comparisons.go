package main

import "C"

import (
// "github.com/danhunsaker/calends"
)

//export Calends_difference
func Calends_difference(p1, p2 C.ulonglong, mode *C.char) *C.char {
	c := instGet(p1)
	z := instGet(p2)
	out := c.Difference(*z, C.GoString(mode))
	return C.CString(out.String())
}

//export Calends_compare
func Calends_compare(p1, p2 C.ulonglong, mode *C.char) C.schar {
	c := instGet(p1)
	z := instGet(p2)
	return C.schar(c.Compare(*z, C.GoString(mode)))
}

//export Calends_is_same
func Calends_is_same(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.IsSame(*z)
}

//export Calends_is_same_duration
func Calends_is_same_duration(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.IsSameDuration(*z)
}

//export Calends_is_shorter
func Calends_is_shorter(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.IsShorter(*z)
}

//export Calends_is_longer
func Calends_is_longer(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.IsLonger(*z)
}

//export Calends_contains
func Calends_contains(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.Contains(*z)
}

//export Calends_overlaps
func Calends_overlaps(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.Overlaps(*z)
}

//export Calends_abuts
func Calends_abuts(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.Abuts(*z)
}

//export Calends_is_before
func Calends_is_before(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.IsBefore(*z)
}

//export Calends_starts_before
func Calends_starts_before(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.StartsBefore(*z)
}

//export Calends_ends_before
func Calends_ends_before(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.EndsBefore(*z)
}

//export Calends_is_during
func Calends_is_during(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.IsDuring(*z)
}

//export Calends_starts_during
func Calends_starts_during(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.StartsDuring(*z)
}

//export Calends_ends_during
func Calends_ends_during(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.EndsDuring(*z)
}

//export Calends_is_after
func Calends_is_after(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.IsAfter(*z)
}

//export Calends_starts_after
func Calends_starts_after(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.StartsAfter(*z)
}

//export Calends_ends_after
func Calends_ends_after(p1, p2 C.ulonglong) bool {
	c := instGet(p1)
	z := instGet(p2)
	return c.EndsAfter(*z)
}
