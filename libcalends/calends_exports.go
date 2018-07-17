package main

//typedef void (*Calends_panic_handler) (char*);
import "C"

import (
	"github.com/danhunsaker/calends"
)

//export Calends_release
func Calends_release(p C.ulonglong) {
	instances.Delete(uint64(p))
}

//export Calends_create_string
func Calends_create_string(stamp, calendar, format *C.char) C.ulonglong {
	return calends_create(C.GoString(stamp), calendar, format)
}

//export Calends_create_string_range
func Calends_create_string_range(start, end, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start": C.GoString(start),
		"end":   C.GoString(end),
	}, calendar, format)
}

//export Calends_create_string_start_period
func Calends_create_string_start_period(start, duration, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start":    C.GoString(start),
		"duration": C.GoString(duration),
	}, calendar, format)
}

//export Calends_create_string_end_period
func Calends_create_string_end_period(duration, end, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"duration": C.GoString(duration),
		"end":      C.GoString(end),
	}, calendar, format)
}

//export Calends_create_long_long
func Calends_create_long_long(stamp C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(int(stamp), calendar, format)
}

//export Calends_create_long_long_range
func Calends_create_long_long_range(start, end C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start": int(start),
		"end":   int(end),
	}, calendar, format)
}

//export Calends_create_long_long_start_period
func Calends_create_long_long_start_period(start, duration C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start":    int(start),
		"duration": int(duration),
	}, calendar, format)
}

//export Calends_create_long_long_end_period
func Calends_create_long_long_end_period(duration, end C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"duration": int(duration),
		"end":      int(end),
	}, calendar, format)
}

//export Calends_create_double
func Calends_create_double(stamp C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(float64(stamp), calendar, format)
}

//export Calends_create_double_range
func Calends_create_double_range(start, end C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start": float64(start),
		"end":   float64(end),
	}, calendar, format)
}

//export Calends_create_double_start_period
func Calends_create_double_start_period(start, duration C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start":    float64(start),
		"duration": float64(duration),
	}, calendar, format)
}

//export Calends_create_double_end_period
func Calends_create_double_end_period(duration, end C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"duration": float64(duration),
		"end":      float64(end),
	}, calendar, format)
}

//export Calends_date
func Calends_date(p C.ulonglong, calendar, format *C.char) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.Date(C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return C.CString(out)
}

//export Calends_duration
func Calends_duration(p C.ulonglong) *C.char {
	c := instGet(p)
	return C.CString(c.Duration().String())
}

//export Calends_end_date
func Calends_end_date(p C.ulonglong, calendar, format *C.char) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.EndDate(C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return C.CString(out)
}

//export Calends_string
func Calends_string(p C.ulonglong) *C.char {
	defer handlePanic()
	c := instGet(p)
	return C.CString(c.String())
}

//export Calends_encode_text
func Calends_encode_text(p C.ulonglong) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.MarshalText()
	if err != nil {
		panic(err)
	}
	return C.CString(string(out))
}

//export Calends_decode_text
func Calends_decode_text(in *C.char) C.ulonglong {
	defer handlePanic()
	c := calends.Calends{}
	err := c.UnmarshalText([]byte(C.GoString(in)))
	if err != nil {
		panic(err)
	}
	return instNum(c)
}

//export Calends_encode_json
func Calends_encode_json(p C.ulonglong) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return C.CString(string(out))
}

//export Calends_decode_json
func Calends_decode_json(in *C.char) C.ulonglong {
	defer handlePanic()
	c := calends.Calends{}
	err := c.UnmarshalJSON([]byte(C.GoString(in)))
	if err != nil {
		panic(err)
	}
	return instNum(c)
}

//export Calends_register_panic_handler
func Calends_register_panic_handler(handler C.Calends_panic_handler) {
	id := nextPanHandle.Id()
	panicHandlers.Store(id, handler)
}
