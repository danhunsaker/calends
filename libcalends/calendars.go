package main

/*
typedef struct _TAI64Time {
        long long int Seconds; // Seconds since 1970-01-01 00:00:00 TAI
        unsigned int Nano;     // Nanoseconds since the given second
        unsigned int Atto;     // Attoseconds since the given nanosecond
        unsigned int Xicto;    // Xictoseconds since the given attosecond
        unsigned int Ucto;     // Uctoseconds since the given xictosecond
        unsigned int Rocto;    // Roctoseconds since the given uctosecond
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

static inline TAI64Time wrap_Calends_calendar_to_internal_string(Calends_calendar_to_internal_string f, char* name, char* date, char* format) {
  return f(name, date, format);
}
static inline TAI64Time wrap_Calends_calendar_to_internal_long_long(Calends_calendar_to_internal_long_long f, char* name, long long int date, char* format) {
  return f(name, date, format);
}
static inline TAI64Time wrap_Calends_calendar_to_internal_double(Calends_calendar_to_internal_double f, char* name, double date, char* format) {
  return f(name, date, format);
}
static inline TAI64Time wrap_Calends_calendar_to_internal_tai(Calends_calendar_to_internal_tai f, char* name, TAI64Time date) {
  return f(name, date);
}

static inline char* wrap_Calends_calendar_from_internal(Calends_calendar_from_internal f, char* name, TAI64Time stamp, char* format) {
  return f(name, stamp, format);
}

static inline TAI64Time wrap_Calends_calendar_offset_string(Calends_calendar_offset_string f, char* name, TAI64Time stamp, char* offset) {
  return f(name, stamp, offset);
}
static inline TAI64Time wrap_Calends_calendar_offset_long_long(Calends_calendar_offset_long_long f, char* name, TAI64Time stamp, long long int offset) {
  return f(name, stamp, offset);
}
static inline TAI64Time wrap_Calends_calendar_offset_double(Calends_calendar_offset_double f, char* name, TAI64Time stamp, double offset) {
  return f(name, stamp, offset);
}
static inline TAI64Time wrap_Calends_calendar_offset_tai(Calends_calendar_offset_tai f, char* name, TAI64Time stamp, TAI64Time offset) {
  return f(name, stamp, offset);
}
*/
import "C"

import (
	"math/big"
	"strings"
	"unsafe"

	"github.com/danhunsaker/calends/calendars"
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
	return taiGoToC(calendars.TAI64NAXURTimeFromDecimalString(C.GoString(in)))
}

//export TAI64Time_hex_string
func TAI64Time_hex_string(t C.TAI64Time) *C.char {
	base := taiCToGo(t)
	return C.CString(base.HexString())
}

//export TAI64Time_from_hex_string
func TAI64Time_from_hex_string(in *C.char) C.TAI64Time {
	return taiGoToC(calendars.TAI64NAXURTimeFromHexString(C.GoString(in)))
}

//export TAI64Time_double
func TAI64Time_double(t C.TAI64Time) C.double {
	base := taiCToGo(t)
	out, _ := base.Float().Float64()
	return C.double(out)
}

//export TAI64Time_from_double
func TAI64Time_from_double(in C.double) C.TAI64Time {
	return taiGoToC(calendars.TAI64NAXURTimeFromFloat(*big.NewFloat(float64(in))))
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
	time := calendars.TAI64NAXURTime{}
	time.UnmarshalText([]byte(C.GoString(in)))
	return taiGoToC(time)
}

//export TAI64Time_encode_binary
func TAI64Time_encode_binary(t C.TAI64Time) unsafe.Pointer {
	defer handlePanic()
	base := taiCToGo(t)
	out, err := base.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return C.CBytes(out)
}

//export TAI64Time_decode_binary
func TAI64Time_decode_binary(in unsafe.Pointer, len C.int) C.TAI64Time {
	time := calendars.TAI64NAXURTime{}
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
) func(interface{}, string) (calendars.TAI64NAXURTime, error) {
	return func(in interface{}, format string) (out calendars.TAI64NAXURTime, err error) {
		var raw C.TAI64Time
		defer handlePanic()
		switch in := in.(type) {
		case string:
			raw, err = C.wrap_Calends_calendar_to_internal_string(toInternalString, name, C.CString(in), C.CString(format))
		case int:
			raw, err = C.wrap_Calends_calendar_to_internal_long_long(toInternalLongLong, name, C.longlong(in), C.CString(format))
		case float64:
			raw, err = C.wrap_Calends_calendar_to_internal_double(toInternalDouble, name, C.double(in), C.CString(format))
		case calendars.TAI64NAXURTime:
			pass := taiGoToC(in)
			raw, err = C.wrap_Calends_calendar_to_internal_tai(toInternalTai, name, pass)
		default:
			err = calendars.ErrUnsupportedInput
		}
		if err != nil {
			return
		}
		out = taiCToGo(raw)
		return
	}
}

func wrapFromInternal(name *C.char, fromInternal C.Calends_calendar_from_internal) func(calendars.TAI64NAXURTime, string) (string, error) {
	return func(in calendars.TAI64NAXURTime, format string) (out string, err error) {
		var raw *C.char
		defer handlePanic()
		pass := taiGoToC(in)
		raw, err = C.wrap_Calends_calendar_from_internal(fromInternal, name, pass, C.CString(format))
		if err != nil {
			return
		}
		out = C.GoString(raw)
		return
	}
}

func wrapOffset(
	name *C.char,
	offsetString C.Calends_calendar_offset_string,
	offsetLongLong C.Calends_calendar_offset_long_long,
	offsetDouble C.Calends_calendar_offset_double,
	offsetTai C.Calends_calendar_offset_tai,
) func(calendars.TAI64NAXURTime, interface{}) (calendars.TAI64NAXURTime, error) {
	return func(in calendars.TAI64NAXURTime, offset interface{}) (out calendars.TAI64NAXURTime, err error) {
		var raw C.TAI64Time
		defer handlePanic()
		pass := taiGoToC(in)
		switch offset := offset.(type) {
		case string:
			raw, err = C.wrap_Calends_calendar_offset_string(offsetString, name, pass, C.CString(offset))
		case int:
			raw, err = C.wrap_Calends_calendar_offset_long_long(offsetLongLong, name, pass, C.longlong(offset))
		case float64:
			raw, err = C.wrap_Calends_calendar_offset_double(offsetDouble, name, pass, C.double(offset))
		case calendars.TAI64NAXURTime:
			pass2 := taiGoToC(offset)
			raw, err = C.wrap_Calends_calendar_offset_tai(offsetTai, name, pass, pass2)
		default:
			err = calendars.ErrUnsupportedInput
		}
		if err != nil {
			return
		}
		out = taiCToGo(raw)
		return
	}
}

func taiGoToC(in calendars.TAI64NAXURTime) C.TAI64Time {
	return C.TAI64Time{
		C.longlong(in.Seconds),
		C.uint(in.Nano),
		C.uint(in.Atto),
		C.uint(in.Xicto),
		C.uint(in.Ucto),
		C.uint(in.Rocto),
		C.uint(0),
	}
}

func taiCToGo(in C.TAI64Time) calendars.TAI64NAXURTime {
	return calendars.TAI64NAXURTime{
		int64(in.Seconds),
		uint32(in.Nano),
		uint32(in.Atto),
		uint32(in.Xicto),
		uint32(in.Ucto),
		uint32(in.Rocto),
	}
}
