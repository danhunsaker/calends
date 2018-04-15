package main

//#include "calendars.h"
import "C"

import (
	"math/big"
	"unsafe"

	"github.com/danhunsaker/calends/calendars"
)

//export Calends_register_calendar
func Calends_register_calendar(
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
		wrapToInternal(toInternalString, toInternalLongLong, toInternalDouble, toInternalTai),
		wrapFromInternal(fromInternal),
		wrapOffset(offsetString, offsetLongLong, offsetDouble, offsetTai),
		C.GoString(defaultFormat),
	)
}

//export Calends_registered
func Calends_registered(calendar *C.char) bool {
	return calendars.Registered(C.GoString(calendar))
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

//export TAI64Time_hex_string
func TAI64Time_hex_string(t C.TAI64Time) *C.char {
	base := taiCToGo(t)
	return C.CString(base.HexString())
}

//export TAI64Time_double
func TAI64Time_double(t C.TAI64Time) C.double {
	base := taiCToGo(t)
	out, _ := base.Float().Float64()
	return C.double(out)
}

//export TAI64Time_encode_text
func TAI64Time_encode_text(t C.TAI64Time) *C.char {
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

//export TAI64Time_from_decimal_string
func TAI64Time_from_decimal_string(in *C.char) C.TAI64Time {
	return taiGoToC(calendars.TAI64NAXURTimeFromDecimalString(C.GoString(in)))
}

//export TAI64Time_from_hex_string
func TAI64Time_from_hex_string(in *C.char) C.TAI64Time {
	return taiGoToC(calendars.TAI64NAXURTimeFromHexString(C.GoString(in)))
}

//export TAI64Time_from_double
func TAI64Time_from_double(in C.double) C.TAI64Time {
	return taiGoToC(calendars.TAI64NAXURTimeFromFloat(*big.NewFloat(float64(in))))
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
	toInternalString C.Calends_calendar_to_internal_string,
	toInternalLongLong C.Calends_calendar_to_internal_long_long,
	toInternalDouble C.Calends_calendar_to_internal_double,
	toInternalTai C.Calends_calendar_to_internal_tai,
) func(interface{}, string) (calendars.TAI64NAXURTime, error) {
	return func(in interface{}, format string) (out calendars.TAI64NAXURTime, err error) {
		var raw *C.TAI64Time
		switch in := in.(type) {
		case string:
			raw, err = C.wrap_Calends_calendar_to_internal_string(toInternalString, C.CString(in), C.CString(format))
		case int:
			raw, err = C.wrap_Calends_calendar_to_internal_long_long(toInternalLongLong, C.longlong(in), C.CString(format))
		case float64:
			raw, err = C.wrap_Calends_calendar_to_internal_double(toInternalDouble, C.double(in), C.CString(format))
		case calendars.TAI64NAXURTime:
			pass := taiGoToC(in)
			raw, err = C.wrap_Calends_calendar_to_internal_tai(toInternalTai, &pass)
		default:
			err = calendars.ErrUnsupportedInput
		}
		if err != nil {
			return
		}
		out = taiCToGo(*raw)
		return
	}
}

func wrapFromInternal(fromInternal C.Calends_calendar_from_internal) func(calendars.TAI64NAXURTime, string) (string, error) {
	return func(in calendars.TAI64NAXURTime, format string) (out string, err error) {
		var raw *C.char
		pass := taiGoToC(in)
		raw, err = C.wrap_Calends_calendar_from_internal(fromInternal, &pass, C.CString(format))
		if err != nil {
			return
		}
		out = C.GoString(raw)
		return
	}
}

func wrapOffset(
	offsetString C.Calends_calendar_offset_string,
	offsetLongLong C.Calends_calendar_offset_long_long,
	offsetDouble C.Calends_calendar_offset_double,
	offsetTai C.Calends_calendar_offset_tai,
) func(calendars.TAI64NAXURTime, interface{}) (calendars.TAI64NAXURTime, error) {
	return func(in calendars.TAI64NAXURTime, offset interface{}) (out calendars.TAI64NAXURTime, err error) {
		var raw *C.TAI64Time
		pass := taiGoToC(in)
		switch offset := offset.(type) {
		case string:
			raw, err = C.wrap_Calends_calendar_offset_string(offsetString, &pass, C.CString(offset))
		case int:
			raw, err = C.wrap_Calends_calendar_offset_long_long(offsetLongLong, &pass, C.longlong(offset))
		case float64:
			raw, err = C.wrap_Calends_calendar_offset_double(offsetDouble, &pass, C.double(offset))
		case calendars.TAI64NAXURTime:
			pass2 := taiGoToC(offset)
			raw, err = C.wrap_Calends_calendar_offset_tai(offsetTai, &pass, &pass2)
		default:
			err = calendars.ErrUnsupportedInput
		}
		if err != nil {
			return
		}
		out = taiCToGo(*raw)
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
		[4]byte{},
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
