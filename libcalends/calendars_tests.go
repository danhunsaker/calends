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

TAI64Time test_Calends_calendar_to_internal_string(char* calendar, char* date, char* format) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}
TAI64Time test_Calends_calendar_to_internal_long_long(char* calendar, long long int date, char* format) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}
TAI64Time test_Calends_calendar_to_internal_double(char* calendar, double date, char* format) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}
TAI64Time test_Calends_calendar_to_internal_tai(char* calendar, TAI64Time date) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}

char* test_Calends_calendar_from_internal(char* calendar, TAI64Time stamp, char* format) {
  return "::";
}

TAI64Time test_Calends_calendar_offset_string(char* calendar, TAI64Time stamp, char* offset) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}
TAI64Time test_Calends_calendar_offset_long_long(char* calendar, TAI64Time stamp, long long int offset) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}
TAI64Time test_Calends_calendar_offset_double(char* calendar, TAI64Time stamp, double offset) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}
TAI64Time test_Calends_calendar_offset_tai(char* calendar, TAI64Time stamp, TAI64Time offset) {
  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};
  return out;
}

*/
import "C"

import (
	"testing"
)

func testCalends_calendar_registered(t *testing.T) {
	t.Helper()

	in := "tai64"
	ret := Calends_calendar_registered(C.CString(in))
	want := true
	if bool(ret) != want {
		t.Errorf("Calends_calendar_registered(%#v) returned %#v; wanted %#v", in, bool(ret), want)
	}

	in = "invalid"
	ret = Calends_calendar_registered(C.CString(in))
	want = false
	if bool(ret) != want {
		t.Errorf("Calends_calendar_registered(%#v) returned %#v; wanted %#v", in, bool(ret), want)
	}
}

func testCalends_calendar_register(t *testing.T) {
	t.Helper()

	in := "test"
	ret := Calends_calendar_registered(C.CString(in))
	want := false
	if bool(ret) != want {
		t.Errorf("Calends_calendar_registered(%#v) returned %#v; wanted %#v", in, bool(ret), want)
	}

	Calends_calendar_register(
		C.CString("test"), C.CString("default"),
		C.Calends_calendar_to_internal_string(C.test_Calends_calendar_to_internal_string),
		C.Calends_calendar_to_internal_long_long(C.test_Calends_calendar_to_internal_long_long),
		C.Calends_calendar_to_internal_double(C.test_Calends_calendar_to_internal_double),
		C.Calends_calendar_to_internal_tai(C.test_Calends_calendar_to_internal_tai),
		C.Calends_calendar_from_internal(C.test_Calends_calendar_from_internal),
		C.Calends_calendar_offset_string(C.test_Calends_calendar_offset_string),
		C.Calends_calendar_offset_long_long(C.test_Calends_calendar_offset_long_long),
		C.Calends_calendar_offset_double(C.test_Calends_calendar_offset_double),
		C.Calends_calendar_offset_tai(C.test_Calends_calendar_offset_tai),
	)

	in = "test"
	ret = Calends_calendar_registered(C.CString(in))
	want = true
	if bool(ret) != want {
		t.Errorf("Calends_calendar_registered(%#v) returned %#v; wanted %#v", in, bool(ret), want)
	}

	calRet := calends_create("", C.CString("test"), C.CString(""))
	if _, ok := instances.Load(uint64(calRet)); uint64(calRet) < 1 || !ok {
		t.Errorf("calends_create(%#v, %#v, %#v) returned invalid Calends object ID %#v", "", "test", "", uint64(calRet))
	}
	Calends_release(calRet)

	calRet = calends_create(0, C.CString("test"), C.CString(""))
	if _, ok := instances.Load(uint64(calRet)); uint64(calRet) < 1 || !ok {
		t.Errorf("calends_create(%#v, %#v, %#v) returned invalid Calends object ID %#v", 0, "test", "", uint64(calRet))
	}
	Calends_release(calRet)

	calRet = calends_create(0.0, C.CString("test"), C.CString(""))
	if _, ok := instances.Load(uint64(calRet)); uint64(calRet) < 1 || !ok {
		t.Errorf("calends_create(%#v, %#v, %#v) returned invalid Calends object ID %#v", 0.0, "test", "", uint64(calRet))
	}
	Calends_release(calRet)

	calRet = calends_create(taiCToGo(C.TAI64Time{0, 0, 0, 0, 0, 0, 0}), C.CString("test"), C.CString(""))
	if _, ok := instances.Load(uint64(calRet)); uint64(calRet) < 1 || !ok {
		t.Errorf("calends_create(%#v, %#v, %#v) returned invalid Calends object ID %#v", taiCToGo(C.TAI64Time{0, 0, 0, 0, 0, 0, 0}), "test", "", uint64(calRet))
	}
	defer Calends_release(calRet)

	str := Calends_date(calRet, C.CString("test"), C.CString(""))
	if C.GoString(str) != "::" {
		t.Errorf("Calends_date(%#v, %#v, %#v) returned %#v; wanted %#v", calRet, "test", "", C.GoString(str), "::")
	}

	working := instGet(calRet)
	wantTime := instGet(calends_create("0", C.CString("tai64"), C.CString("decimal")))

	time, _ := working.Add("", "test")
	if time.String() != wantTime.String() {
		t.Errorf("Calends{%#v}.Add(%#v, %#v) returned %#v; wanted %#v", calRet, "", "test", time, wantTime)
	}

	time, _ = working.Add(0, "test")
	if time.String() != wantTime.String() {
		t.Errorf("Calends{%#v}.Add(%#v, %#v) returned %#v; wanted %#v", calRet, 0, "test", time, wantTime)
	}

	time, _ = working.Add(0.0, "test")
	if time.String() != wantTime.String() {
		t.Errorf("Calends{%#v}.Add(%#v, %#v) returned %#v; wanted %#v", calRet, 0.0, "test", time, wantTime)
	}

	time, _ = working.Add(taiCToGo(C.TAI64Time{0, 0, 0, 0, 0, 0, 0}), "test")
	if time.String() != wantTime.String() {
		t.Errorf("Calends{%#v}.Add(%#v, %#v) returned %#v; wanted %#v", calRet, taiCToGo(C.TAI64Time{0, 0, 0, 0, 0, 0, 0}), "test", time, wantTime)
	}
}

func testTAI64Time_add(t *testing.T) {
	t.Helper()

	in1 := C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	in2 := C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_add(in1, in2)
	want := C.TAI64Time{20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_add(%#v, %#v) returned %#v; wanted %#v", in1, in2, ret, want)
	}
}

func testTAI64Time_sub(t *testing.T) {
	t.Helper()

	in1 := C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	in2 := C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_sub(in1, in2)
	want := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_sub(%#v, %#v) returned %#v; wanted %#v", in1, in2, ret, want)
	}
}

func testTAI64Time_string(t *testing.T) {
	t.Helper()

	in := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_string(in)
	want := "0"
	if C.GoString(ret) != want {
		t.Errorf("TAI64Time_string(%#v) returned %#v; wanted %#v", in, C.GoString(ret), want)
	}

	in = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret = TAI64Time_string(in)
	want = "10"
	if C.GoString(ret) != want {
		t.Errorf("TAI64Time_string(%#v) returned %#v; wanted %#v", in, C.GoString(ret), want)
	}
}

func testTAI64Time_from_string(t *testing.T) {
	t.Helper()

	ret := TAI64Time_from_string(C.CString("0"))
	want := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_from_string(%#v) returned %#v; wanted %#v", "0", ret, want)
	}

	ret = TAI64Time_from_string(C.CString("10"))
	want = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_from_string(%#v) returned %#v; wanted %#v", "10", ret, want)
	}
}

func testTAI64Time_hex_string(t *testing.T) {
	t.Helper()

	in := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_hex_string(in)
	want := "40000000000000000000000000000000000000000000000000000000"
	if C.GoString(ret) != want {
		t.Errorf("TAI64Time_hex_string(%#v) returned %#v; wanted %#v", in, C.GoString(ret), want)
	}

	in = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret = TAI64Time_hex_string(in)
	want = "400000000000000A0000000000000000000000000000000000000000"
	if C.GoString(ret) != want {
		t.Errorf("TAI64Time_hex_string(%#v) returned %#v; wanted %#v", in, C.GoString(ret), want)
	}
}

func testTAI64Time_from_hex_string(t *testing.T) {
	t.Helper()

	ret := TAI64Time_from_hex_string(C.CString("40000000000000000000000000000000000000000000000000000000"))
	want := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_from_hex_string(%#v) returned %#v; wanted %#v", "40000000000000000000000000000000000000000000000000000000", ret, want)
	}

	ret = TAI64Time_from_hex_string(C.CString("400000000000000A0000000000000000000000000000000000000000"))
	want = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_from_hex_string(%#v) returned %#v; wanted %#v", "400000000000000A0000000000000000000000000000000000000000", ret, want)
	}
}

func testTAI64Time_double(t *testing.T) {
	t.Helper()

	in := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_double(in)
	want := 0.0
	if float64(ret) != want {
		t.Errorf("TAI64Time_double(%#v) returned %#v; wanted %#v", in, float64(ret), want)
	}

	in = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret = TAI64Time_double(in)
	want = 10.0
	if float64(ret) != want {
		t.Errorf("TAI64Time_double(%#v) returned %#v; wanted %#v", in, float64(ret), want)
	}
}

func testTAI64Time_from_double(t *testing.T) {
	t.Helper()

	ret := TAI64Time_from_double(C.double(0.0))
	want := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_from_double(%#v) returned %#v; wanted %#v", 0.0, ret, want)
	}

	ret = TAI64Time_from_double(C.double(10.0))
	want = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_from_double(%#v) returned %#v; wanted %#v", 10.0, ret, want)
	}
}

func testTAI64Time_encode_text(t *testing.T) {
	t.Helper()

	in := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_encode_text(in)
	want := "40000000000000000000000000000000000000000000000000000000"
	if C.GoString(ret) != want {
		t.Errorf("TAI64Time_encode_text(%#v) returned %#v; wanted %#v", in, C.GoString(ret), want)
	}

	in = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret = TAI64Time_encode_text(in)
	want = "400000000000000A0000000000000000000000000000000000000000"
	if C.GoString(ret) != want {
		t.Errorf("TAI64Time_encode_text(%#v) returned %#v; wanted %#v", in, C.GoString(ret), want)
	}
}

func testTAI64Time_decode_text(t *testing.T) {
	t.Helper()

	ret := TAI64Time_decode_text(C.CString("40000000000000000000000000000000000000000000000000000000"))
	want := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_decode_text(%#v) returned %#v; wanted %#v", "40000000000000000000000000000000000000000000000000000000", ret, want)
	}

	ret = TAI64Time_decode_text(C.CString("400000000000000A0000000000000000000000000000000000000000"))
	want = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_decode_text(%#v) returned %#v; wanted %#v", "400000000000000A0000000000000000000000000000000000000000", ret, want)
	}
}

func testTAI64Time_encode_binary(t *testing.T) {
	t.Helper()

	in := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_encode_binary(in)
	want := "\x40\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	if string(C.GoBytes(ret, 28)) != want {
		t.Errorf("TAI64Time_encode_binary(%#v) returned %#v; wanted %#v", in, string(C.GoBytes(ret, 28)), want)
	}

	in = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret = TAI64Time_encode_binary(in)
	want = "\x40\x00\x00\x00\x00\x00\x00\x0A\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	if string(C.GoBytes(ret, 28)) != want {
		t.Errorf("TAI64Time_encode_binary(%#v) returned %#v; wanted %#v", in, string(C.GoBytes(ret, 28)), want)
	}
}

func testTAI64Time_decode_binary(t *testing.T) {
	t.Helper()

	ret := TAI64Time_decode_binary(C.CBytes([]byte("\x40\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")), C.int(28))
	want := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_decode_binary(%#v, %#v) returned %#v; wanted %#v", "\x40\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", 28, ret, want)
	}

	ret = TAI64Time_decode_binary(C.CBytes([]byte("\x40\x00\x00\x00\x00\x00\x00\x0A\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")), C.int(28))
	want = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_decode_binary(%#v, %#v) returned %#v; wanted %#v", "\x40\x00\x00\x00\x00\x00\x00\x0A\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", 28, ret, want)
	}
}

func testTAI64Time_utc_to_tai(t *testing.T) {
	t.Helper()

	in := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_utc_to_tai(in)
	want := C.TAI64Time{-7, 0x3b747d4f, 0x3b9ac9f3, 0x2e697bf0, 0x3735dcc3, 0xfe4046, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_utc_to_tai(%#v) returned %#v; wanted %#v", in, ret, want)
	}

	in = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret = TAI64Time_utc_to_tai(in)
	want = C.TAI64Time{2, 0x1404f, 0x3b9ac633, 0x119bb285, 0x32db8e0d, 0x3399e120, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_utc_to_tai(%#v) returned %#v; wanted %#v", in, ret, want)
	}
}

func testTAI64Time_tai_to_utc(t *testing.T) {
	t.Helper()

	in := C.TAI64Time{0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret := TAI64Time_tai_to_utc(in)
	want := C.TAI64Time{8, 0x1404f, 0x3b9ac633, 0x119bb285, 0x32db8e0d, 0x3399e120, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_tai_to_utc(%#v) returned %#v; wanted %#v", in, ret, want)
	}

	in = C.TAI64Time{10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	ret = TAI64Time_tai_to_utc(in)
	want = C.TAI64Time{18, 0x1404f, 0x3b9ac633, 0x119bb285, 0x32db8e0d, 0x3399e120, 0x0}
	if ret != want {
		t.Errorf("TAI64Time_tai_to_utc(%#v) returned %#v; wanted %#v", in, ret, want)
	}
}
