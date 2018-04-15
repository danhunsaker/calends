package main

//char* Calends_version;
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/danhunsaker/calends"
)

func init() {
	C.Calends_version = C.CString(calends.Version)
}

func Calends_create(stamp interface{}, calendar, format *C.char) unsafe.Pointer {
	c, err := calends.Create(stamp, C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&c)
}

//export Calends_create_string
func Calends_create_string(stamp, calendar, format *C.char) unsafe.Pointer {
	return Calends_create(C.GoString(stamp), calendar, format)
}

//export Calends_create_long_long
func Calends_create_long_long(stamp C.longlong, calendar, format *C.char) unsafe.Pointer {
	return Calends_create(int(stamp), calendar, format)
}

//export Calends_create_double
func Calends_create_double(stamp C.double, calendar, format *C.char) unsafe.Pointer {
	return Calends_create(float64(stamp), calendar, format)
}

//export Calends_date
func Calends_date(p unsafe.Pointer, calendar, format *C.char) *C.char {
	c := (*calends.Calends)(p)
	out, err := c.Date(C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return C.CString(out)
}

//export Calends_duration
func Calends_duration(p unsafe.Pointer) *C.char {
	c := (*calends.Calends)(p)
	return C.CString(c.Duration().String())
}

//export Calends_end_date
func Calends_end_date(p unsafe.Pointer, calendar, format *C.char) *C.char {
	c := (*calends.Calends)(p)
	out, err := c.EndDate(C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return C.CString(out)
}

//export Calends_string
func Calends_string(p unsafe.Pointer) *C.char {
	c := (*calends.Calends)(p)
	return C.CString(c.String())
}

//export Calends_encode_text
func Calends_encode_text(p unsafe.Pointer) *C.char {
	c := (*calends.Calends)(p)
	out, err := c.MarshalText()
	if err != nil {
		panic(err)
	}
	return C.CString(string(out))
}

//export Calends_decode_text
func Calends_decode_text(in *C.char) unsafe.Pointer {
	c := calends.Calends{}
	err := c.UnmarshalText([]byte(C.GoString(in)))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&c)
}

//export Calends_encode_json
func Calends_encode_json(p unsafe.Pointer) *C.char {
	c := (*calends.Calends)(p)
	out, err := c.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return C.CString(string(out))
}

//export Calends_decode_json
func Calends_decode_json(in *C.char) unsafe.Pointer {
	c := calends.Calends{}
	err := c.UnmarshalJSON([]byte(C.GoString(in)))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&c)
}

func main() {
	panic(fmt.Sprintf("Calends %s\nThis shouldn't ever be called!", C.GoString(C.Calends_version)))
}
