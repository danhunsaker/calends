package main

import "C"

func calends_add(p C.ulonglong, offset interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.Add(offset, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_add_string
func Calends_add_string(p C.ulonglong, offset, calendar *C.char) C.ulonglong {
	return calends_add(p, C.GoString(offset), calendar)
}

//export Calends_add_long_long
func Calends_add_long_long(p C.ulonglong, offset C.longlong, calendar *C.char) C.ulonglong {
	return calends_add(p, int(offset), calendar)
}

//export Calends_add_double
func Calends_add_double(p C.ulonglong, offset C.double, calendar *C.char) C.ulonglong {
	return calends_add(p, float64(offset), calendar)
}

func calends_subtract(p C.ulonglong, offset interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.Subtract(offset, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_subtract_string
func Calends_subtract_string(p C.ulonglong, offset, calendar *C.char) C.ulonglong {
	return calends_subtract(p, C.GoString(offset), calendar)
}

//export Calends_subtract_long_long
func Calends_subtract_long_long(p C.ulonglong, offset C.longlong, calendar *C.char) C.ulonglong {
	return calends_subtract(p, int(offset), calendar)
}

//export Calends_subtract_double
func Calends_subtract_double(p C.ulonglong, offset C.double, calendar *C.char) C.ulonglong {
	return calends_subtract(p, float64(offset), calendar)
}

func calends_add_from_end(p C.ulonglong, offset interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.AddFromEnd(offset, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_add_from_end_string
func Calends_add_from_end_string(p C.ulonglong, offset, calendar *C.char) C.ulonglong {
	return calends_add_from_end(p, C.GoString(offset), calendar)
}

//export Calends_add_from_end_long_long
func Calends_add_from_end_long_long(p C.ulonglong, offset C.longlong, calendar *C.char) C.ulonglong {
	return calends_add_from_end(p, int(offset), calendar)
}

//export Calends_add_from_end_double
func Calends_add_from_end_double(p C.ulonglong, offset C.double, calendar *C.char) C.ulonglong {
	return calends_add_from_end(p, float64(offset), calendar)
}

func calends_subtract_from_end(p C.ulonglong, offset interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.SubtractFromEnd(offset, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_subtract_from_end_string
func Calends_subtract_from_end_string(p C.ulonglong, offset, calendar *C.char) C.ulonglong {
	return calends_subtract_from_end(p, C.GoString(offset), calendar)
}

//export Calends_subtract_from_end_long_long
func Calends_subtract_from_end_long_long(p C.ulonglong, offset C.longlong, calendar *C.char) C.ulonglong {
	return calends_subtract_from_end(p, int(offset), calendar)
}

//export Calends_subtract_from_end_double
func Calends_subtract_from_end_double(p C.ulonglong, offset C.double, calendar *C.char) C.ulonglong {
	return calends_subtract_from_end(p, float64(offset), calendar)
}

func calends_next(p C.ulonglong, offset interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.Next(offset, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_next_string
func Calends_next_string(p C.ulonglong, offset, calendar *C.char) C.ulonglong {
	return calends_next(p, C.GoString(offset), calendar)
}

//export Calends_next_long_long
func Calends_next_long_long(p C.ulonglong, offset C.longlong, calendar *C.char) C.ulonglong {
	return calends_next(p, int(offset), calendar)
}

//export Calends_next_double
func Calends_next_double(p C.ulonglong, offset C.double, calendar *C.char) C.ulonglong {
	return calends_next(p, float64(offset), calendar)
}

func calends_previous(p C.ulonglong, offset interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.Previous(offset, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_previous_string
func Calends_previous_string(p C.ulonglong, offset, calendar *C.char) C.ulonglong {
	return calends_previous(p, C.GoString(offset), calendar)
}

//export Calends_previous_long_long
func Calends_previous_long_long(p C.ulonglong, offset C.longlong, calendar *C.char) C.ulonglong {
	return calends_previous(p, int(offset), calendar)
}

//export Calends_previous_double
func Calends_previous_double(p C.ulonglong, offset C.double, calendar *C.char) C.ulonglong {
	return calends_previous(p, float64(offset), calendar)
}

func calends_with_date(p C.ulonglong, stamp interface{}, calendar, format *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.SetDate(stamp, C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_with_date_string
func Calends_with_date_string(p C.ulonglong, stamp, calendar, format *C.char) C.ulonglong {
	return calends_with_date(p, C.GoString(stamp), calendar, format)
}

//export Calends_with_date_long_long
func Calends_with_date_long_long(p C.ulonglong, stamp C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_with_date(p, int(stamp), calendar, format)
}

//export Calends_with_date_double
func Calends_with_date_double(p C.ulonglong, stamp C.double, calendar, format *C.char) C.ulonglong {
	return calends_with_date(p, float64(stamp), calendar, format)
}

func calends_with_end_date(p C.ulonglong, stamp interface{}, calendar, format *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.SetEndDate(stamp, C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_with_end_date_string
func Calends_with_end_date_string(p C.ulonglong, stamp, calendar, format *C.char) C.ulonglong {
	return calends_with_end_date(p, C.GoString(stamp), calendar, format)
}

//export Calends_with_end_date_long_long
func Calends_with_end_date_long_long(p C.ulonglong, stamp C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_with_end_date(p, int(stamp), calendar, format)
}

//export Calends_with_end_date_double
func Calends_with_end_date_double(p C.ulonglong, stamp C.double, calendar, format *C.char) C.ulonglong {
	return calends_with_end_date(p, float64(stamp), calendar, format)
}

func calends_with_duration(p C.ulonglong, duration interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.SetDuration(duration, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_with_duration_string
func Calends_with_duration_string(p C.ulonglong, duration, calendar *C.char) C.ulonglong {
	return calends_with_duration(p, C.GoString(duration), calendar)
}

//export Calends_with_duration_long_long
func Calends_with_duration_long_long(p C.ulonglong, duration C.longlong, calendar *C.char) C.ulonglong {
	return calends_with_duration(p, int(duration), calendar)
}

//export Calends_with_duration_double
func Calends_with_duration_double(p C.ulonglong, duration C.double, calendar *C.char) C.ulonglong {
	return calends_with_duration(p, float64(duration), calendar)
}

func calends_with_duration_from_end(p C.ulonglong, duration interface{}, calendar *C.char) C.ulonglong {
	defer handlePanic()
	c := instGet(p)
	out, err := c.SetDurationFromEnd(duration, C.GoString(calendar))
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_with_duration_from_end_string
func Calends_with_duration_from_end_string(p C.ulonglong, duration, calendar *C.char) C.ulonglong {
	return calends_with_duration_from_end(p, C.GoString(duration), calendar)
}

//export Calends_with_duration_from_end_long_long
func Calends_with_duration_from_end_long_long(p C.ulonglong, duration C.longlong, calendar *C.char) C.ulonglong {
	return calends_with_duration_from_end(p, int(duration), calendar)
}

//export Calends_with_duration_from_end_double
func Calends_with_duration_from_end_double(p C.ulonglong, duration C.double, calendar *C.char) C.ulonglong {
	return calends_with_duration_from_end(p, float64(duration), calendar)
}

//export Calends_merge
func Calends_merge(p1, p2 C.ulonglong) C.ulonglong {
	defer handlePanic()
	c := instGet(p1)
	z := instGet(p2)
	out, err := c.Merge(*z)
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_intersect
func Calends_intersect(p1, p2 C.ulonglong) C.ulonglong {
	defer handlePanic()
	c := instGet(p1)
	z := instGet(p2)
	out, err := c.Intersect(*z)
	if err != nil {
		panic(err)
	}
	return instNum(out)
}

//export Calends_gap
func Calends_gap(p1, p2 C.ulonglong) C.ulonglong {
	defer handlePanic()
	c := instGet(p1)
	z := instGet(p2)
	out, err := c.Gap(*z)
	if err != nil {
		panic(err)
	}
	return instNum(out)
}
