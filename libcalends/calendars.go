package main

/*
typedef struct _TAI64Time {
    long long int Seconds; // Seconds since 1970-01-01 00:00:00 TAI
    unsigned int Nano;     // Billionths of a second since the given second
    unsigned int Atto;     // Billionths of a nanosecond since the given nanosecond
    unsigned int Ronto;    // Billionths of an attosecond since the given attosecond
    unsigned int Udecto;   // Billionths of a rontosecond since the given rontosecond
    unsigned int Xindecto; // Billionths of an udectosecond since the given udectosecond
    unsigned int padding;  // round the value out to the nearest 64 bits
} TAI64Time;

typedef TAI64Time (*Calends_calendar_to_internal_string) (char*, char*, char*);
typedef TAI64Time (*Calends_calendar_to_internal_long_long) (char*, long long int, char*);
typedef TAI64Time (*Calends_calendar_to_internal_double) (char*, double, char*);
typedef TAI64Time (*Calends_calendar_to_internal_tai) (char*, TAI64Time);

typedef char* (*Calends_calendar_from_internal) (char*, TAI64Time, char*);

typedef TAI64Time (*Calends_calendar_offset_string) (char*, TAI64Time, char*);
typedef TAI64Time (*Calends_calendar_offset_long_long) (char*, TAI64Time, long long int);
typedef TAI64Time (*Calends_calendar_offset_double) (char*, TAI64Time, double);
typedef TAI64Time (*Calends_calendar_offset_tai) (char*, TAI64Time, TAI64Time);

static inline TAI64Time bridge_Calends_calendar_to_internal_string(Calends_calendar_to_internal_string f, char* name, char* date, char* format) {
    return f(name, date, format);
}
static inline TAI64Time bridge_Calends_calendar_to_internal_long_long(Calends_calendar_to_internal_long_long f, char* name, long long int date, char* format) {
    return f(name, date, format);
}
static inline TAI64Time bridge_Calends_calendar_to_internal_double(Calends_calendar_to_internal_double f, char* name, double date, char* format) {
    return f(name, date, format);
}
static inline TAI64Time bridge_Calends_calendar_to_internal_tai(Calends_calendar_to_internal_tai f, char* name, TAI64Time date) {
    return f(name, date);
}

static inline char* bridge_Calends_calendar_from_internal(Calends_calendar_from_internal f, char* name, TAI64Time stamp, char* format) {
    return f(name, stamp, format);
}

static inline TAI64Time bridge_Calends_calendar_offset_string(Calends_calendar_offset_string f, char* name, TAI64Time stamp, char* offset) {
    return f(name, stamp, offset);
}
static inline TAI64Time bridge_Calends_calendar_offset_long_long(Calends_calendar_offset_long_long f, char* name, TAI64Time stamp, long long int offset) {
    return f(name, stamp, offset);
}
static inline TAI64Time bridge_Calends_calendar_offset_double(Calends_calendar_offset_double f, char* name, TAI64Time stamp, double offset) {
    return f(name, stamp, offset);
}
static inline TAI64Time bridge_Calends_calendar_offset_tai(Calends_calendar_offset_tai f, char* name, TAI64Time stamp, TAI64Time offset) {
    return f(name, stamp, offset);
}
*/
import "C"

import (
	"math/big"
	"strings"
	"syscall"
	"unsafe"

	"github.com/danhunsaker/calends/calendars"
	"github.com/danhunsaker/calends/calendars/dynamic"
	"github.com/go-errors/errors"
)

//export Calends_calendar_register
func Calends_calendar_register(
	name, defaultFormat *C.char,
	toInternalString C.Calends_calendar_to_internal_string,
	toInternalLongLong C.Calends_calendar_to_internal_long_long,
	toInternalDouble C.Calends_calendar_to_internal_double,
	toInternalTai C.Calends_calendar_to_internal_tai,
	fromInternal C.Calends_calendar_from_internal,
	offsetString C.Calends_calendar_offset_string,
	offsetLongLong C.Calends_calendar_offset_long_long,
	offsetDouble C.Calends_calendar_offset_double,
	offsetTai C.Calends_calendar_offset_tai,
) {
	calendars.RegisterElements(
		C.GoString(name),
		wrapToInternal(name, toInternalString, toInternalLongLong, toInternalDouble, toInternalTai),
		wrapFromInternal(name, fromInternal),
		wrapOffset(name, offsetString, offsetLongLong, offsetDouble, offsetTai),
		C.GoString(defaultFormat),
	)
}

//export Calends_calendar_register_dynamic
func Calends_calendar_register_dynamic(json *C.char) {
	defer handlePanic()
	calendar := dynamic.Calendar{}
	err := calendar.UnmarshalJSON([]byte(C.GoString(json)))
	if err != nil {
		panic(err)
	}
	calendars.RegisterDynamic(calendar)
}

//export Calends_calendar_unregister
func Calends_calendar_unregister(name *C.char) {
	calendars.Unregister(C.GoString(name))
}

//export Calends_calendar_registered
func Calends_calendar_registered(calendar *C.char) bool {
	return calendars.Registered(C.GoString(calendar))
}

//export Calends_calendar_list_registered
func Calends_calendar_list_registered() *C.char {
	return C.CString(strings.Join(calendars.ListRegistered(), "\n"))
}

//export TAI64Time_add
func TAI64Time_add(t C.TAI64Time, z C.TAI64Time) C.TAI64Time {
	base := taiCToGo(t)
	addend := taiCToGo(z)
	return taiGoToC(base.Add(addend))
}

//export TAI64Time_sub
func TAI64Time_sub(t C.TAI64Time, z C.TAI64Time) C.TAI64Time {
	base := taiCToGo(t)
	subtend := taiCToGo(z)
	return taiGoToC(base.Sub(subtend))
}

//export TAI64Time_string
func TAI64Time_string(t C.TAI64Time) *C.char {
	base := taiCToGo(t)
	return C.CString(base.String())
}

//export TAI64Time_from_string
func TAI64Time_from_string(in *C.char) C.TAI64Time {
	return taiGoToC(calendars.TAI64NARUXTimeFromDecimalString(C.GoString(in)))
}

//export TAI64Time_hex_string
func TAI64Time_hex_string(t C.TAI64Time) *C.char {
	base := taiCToGo(t)
	return C.CString(base.HexString())
}

//export TAI64Time_from_hex_string
func TAI64Time_from_hex_string(in *C.char) C.TAI64Time {
	return taiGoToC(calendars.TAI64NARUXTimeFromHexString(C.GoString(in)))
}

//export TAI64Time_double
func TAI64Time_double(t C.TAI64Time) C.double {
	base := taiCToGo(t)
	out, _ := base.Float().Float64()
	return C.double(out)
}

//export TAI64Time_from_double
func TAI64Time_from_double(in C.double) C.TAI64Time {
	return taiGoToC(calendars.TAI64NARUXTimeFromFloat(*big.NewFloat(float64(in))))
}

//export TAI64Time_encode_text
func TAI64Time_encode_text(t C.TAI64Time) *C.char {
	defer handlePanic()
	base := taiCToGo(t)
	out, err := base.MarshalText()
	if err != nil {
		panic(err)
	}
	return C.CString(string(out))
}

//export TAI64Time_decode_text
func TAI64Time_decode_text(in *C.char) C.TAI64Time {
	time := calendars.TAI64NARUXTime{}
	time.UnmarshalText([]byte(C.GoString(in)))
	return taiGoToC(time)
}

//export TAI64Time_encode_binary
func TAI64Time_encode_binary(t C.TAI64Time, length *C.int) unsafe.Pointer {
	defer handlePanic()
	base := taiCToGo(t)
	byteStream, err := base.MarshalBinary()
	if err != nil {
		panic(err)
	}
	*length = C.int(len(byteStream))
	return C.CBytes(byteStream)
}

//export TAI64Time_decode_binary
func TAI64Time_decode_binary(in unsafe.Pointer, len C.int) C.TAI64Time {
	time := calendars.TAI64NARUXTime{}
	time.UnmarshalBinary(C.GoBytes(in, len))
	return taiGoToC(time)
}

//export TAI64Time_utc_to_tai
func TAI64Time_utc_to_tai(utc C.TAI64Time) C.TAI64Time {
	return taiGoToC(calendars.UTCtoTAI(taiCToGo(utc)))
}

//export TAI64Time_tai_to_utc
func TAI64Time_tai_to_utc(tai C.TAI64Time) C.TAI64Time {
	return taiGoToC(calendars.TAItoUTC(taiCToGo(tai)))
}

func wrapToInternal(
	name *C.char,
	toInternalString C.Calends_calendar_to_internal_string,
	toInternalLongLong C.Calends_calendar_to_internal_long_long,
	toInternalDouble C.Calends_calendar_to_internal_double,
	toInternalTai C.Calends_calendar_to_internal_tai,
) func(interface{}, string) (calendars.TAI64NARUXTime, error) {
	return func(in interface{}, format string) (out calendars.TAI64NARUXTime, err error) {
		var raw C.TAI64Time
		defer handlePanic()
		switch in := in.(type) {
		case string:
			raw, err = C.bridge_Calends_calendar_to_internal_string(toInternalString, name, C.CString(in), C.CString(format))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case int:
			raw, err = C.bridge_Calends_calendar_to_internal_long_long(toInternalLongLong, name, C.longlong(in), C.CString(format))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case float64:
			raw, err = C.bridge_Calends_calendar_to_internal_double(toInternalDouble, name, C.double(in), C.CString(format))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case calendars.TAI64NARUXTime:
			pass := taiGoToC(in)
			raw, err = C.bridge_Calends_calendar_to_internal_tai(toInternalTai, name, pass)
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		default:
			err = errors.Wrap(calendars.ErrUnsupportedInput, 1)
		}
		if err != nil && !errors.Is(err, syscall.Errno(2)) {
			panic(err)
		}
		out = taiCToGo(raw)
		err = nil
		return
	}
}

func wrapFromInternal(name *C.char, fromInternal C.Calends_calendar_from_internal) func(calendars.TAI64NARUXTime, string) (string, error) {
	return func(in calendars.TAI64NARUXTime, format string) (out string, err error) {
		var raw *C.char
		defer handlePanic()
		pass := taiGoToC(in)
		raw, err = C.bridge_Calends_calendar_from_internal(fromInternal, name, pass, C.CString(format))
		if err != nil && !errors.Is(err, syscall.Errno(2)) {
			panic(errors.Wrap(err, 0))
		}
		out = C.GoString(raw)
		err = nil
		return
	}
}

func wrapOffset(
	name *C.char,
	offsetString C.Calends_calendar_offset_string,
	offsetLongLong C.Calends_calendar_offset_long_long,
	offsetDouble C.Calends_calendar_offset_double,
	offsetTai C.Calends_calendar_offset_tai,
) func(calendars.TAI64NARUXTime, interface{}) (calendars.TAI64NARUXTime, error) {
	return func(in calendars.TAI64NARUXTime, offset interface{}) (out calendars.TAI64NARUXTime, err error) {
		var raw C.TAI64Time
		defer handlePanic()
		pass := taiGoToC(in)
		switch offset := offset.(type) {
		case string:
			raw, err = C.bridge_Calends_calendar_offset_string(offsetString, name, pass, C.CString(offset))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case int:
			raw, err = C.bridge_Calends_calendar_offset_long_long(offsetLongLong, name, pass, C.longlong(offset))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case float64:
			raw, err = C.bridge_Calends_calendar_offset_double(offsetDouble, name, pass, C.double(offset))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case calendars.TAI64NARUXTime:
			pass2 := taiGoToC(offset)
			raw, err = C.bridge_Calends_calendar_offset_tai(offsetTai, name, pass, pass2)
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		default:
			err = errors.Wrap(calendars.ErrUnsupportedInput, 1)
		}
		if err != nil && !errors.Is(err, syscall.Errno(2)) {
			panic(err)
		}
		out = taiCToGo(raw)
		err = nil
		return
	}
}

func taiGoToC(in calendars.TAI64NARUXTime) C.TAI64Time {
	return C.TAI64Time{
		Seconds:  C.longlong(in.Seconds),
		Nano:     C.uint(in.Nano),
		Atto:     C.uint(in.Atto),
		Ronto:    C.uint(in.Ronto),
		Udecto:   C.uint(in.Udecto),
		Xindecto: C.uint(in.Xindecto),
		padding:  C.uint(0),
	}
}

func taiCToGo(in C.TAI64Time) calendars.TAI64NARUXTime {
	return calendars.TAI64NARUXTime{
		Seconds:  int64(in.Seconds),
		Nano:     uint32(in.Nano),
		Atto:     uint32(in.Atto),
		Ronto:    uint32(in.Ronto),
		Udecto:   uint32(in.Udecto),
		Xindecto: uint32(in.Xindecto),
	}
}
