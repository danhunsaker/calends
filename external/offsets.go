package main

import "C"

import (
	"unsafe"

	"github.com/danhunsaker/calends"
)

//export Calends_add
func Calends_add(p unsafe.Pointer, offset, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.Add(C.GoString(offset), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_subtract
func Calends_subtract(p unsafe.Pointer, offset, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.Subtract(C.GoString(offset), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_add_from_end
func Calends_add_from_end(p unsafe.Pointer, offset, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.AddFromEnd(C.GoString(offset), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_subtract_from_end
func Calends_subtract_from_end(p unsafe.Pointer, offset, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.SubtractFromEnd(C.GoString(offset), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_next
func Calends_next(p unsafe.Pointer, offset, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.Next(C.GoString(offset), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_previous
func Calends_previous(p unsafe.Pointer, offset, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.Previous(C.GoString(offset), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

func Calends_with_date(p unsafe.Pointer, stamp interface{}, calendar, format *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.SetDate(stamp, C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_with_date_string
func Calends_with_date_string(p unsafe.Pointer, stamp, calendar, format *C.char) unsafe.Pointer {
	return Calends_with_date(p, C.GoString(stamp), calendar, format)
}

//export Calends_with_date_long_long
func Calends_with_date_long_long(p unsafe.Pointer, stamp C.longlong, calendar, format *C.char) unsafe.Pointer {
	return Calends_with_date(p, int(stamp), calendar, format)
}

//export Calends_with_date_double
func Calends_with_date_double(p unsafe.Pointer, stamp C.double, calendar, format *C.char) unsafe.Pointer {
	return Calends_with_date(p, float64(stamp), calendar, format)
}

func Calends_with_end_date(p unsafe.Pointer, stamp interface{}, calendar, format *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.SetEndDate(stamp, C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_with_end_date_string
func Calends_with_end_date_string(p unsafe.Pointer, stamp, calendar, format *C.char) unsafe.Pointer {
	return Calends_with_end_date(p, C.GoString(stamp), calendar, format)
}

//export Calends_with_end_date_long_long
func Calends_with_end_date_long_long(p unsafe.Pointer, stamp C.longlong, calendar, format *C.char) unsafe.Pointer {
	return Calends_with_end_date(p, int(stamp), calendar, format)
}

//export Calends_with_end_date_double
func Calends_with_end_date_double(p unsafe.Pointer, stamp C.double, calendar, format *C.char) unsafe.Pointer {
	return Calends_with_end_date(p, float64(stamp), calendar, format)
}

//export Calends_with_duration
func Calends_with_duration(p unsafe.Pointer, duration, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.SetDuration(C.GoString(duration), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_with_duration_from_end
func Calends_with_duration_from_end(p unsafe.Pointer, duration, calendar *C.char) unsafe.Pointer {
	c := (*calends.Calends)(p)
	out, err := c.SetDurationFromEnd(C.GoString(duration), C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_merge
func Calends_merge(p1, p2 unsafe.Pointer) unsafe.Pointer {
	c := (*calends.Calends)(p1)
	z := (*calends.Calends)(p2)
	out, err := c.Merge(*z)
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_intersect
func Calends_intersect(p1, p2 unsafe.Pointer) unsafe.Pointer {
	c := (*calends.Calends)(p1)
	z := (*calends.Calends)(p2)
	out, err := c.Intersect(*z)
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}

//export Calends_gap
func Calends_gap(p1, p2 unsafe.Pointer) unsafe.Pointer {
	c := (*calends.Calends)(p1)
	z := (*calends.Calends)(p2)
	out, err := c.Gap(*z)
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&out)
}
