package main

//typedef void (*Calends_panic_handler) (char*);
//
//void test_Calends_panic_handler(char* message) {}
import "C"

import (
	"testing"
)

func testCalends_release(t *testing.T) {
	t.Helper()

	inst := calends_create("", C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(inst)); uint64(inst) < 1 || !ok {
		t.Errorf("Calends_release(%#v) test Calends object ID invalid!", uint64(inst))
	}
	Calends_release(inst)
	if _, ok := instances.Load(uint64(inst)); ok {
		t.Errorf("Calends_release(%#v) failed to release Calends object", uint64(inst))
	}
}

func testCalends_create_string(t *testing.T) {
	t.Helper()

	ret := Calends_create_string(C.CString(""), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", "", "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_string_range(t *testing.T) {
	t.Helper()

	ret := Calends_create_string_range(C.CString(""), C.CString(""), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_string_range(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", "", "", "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_string_start_period(t *testing.T) {
	t.Helper()

	ret := Calends_create_string_start_period(C.CString(""), C.CString(""), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_string_start_period(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", "", "", "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_string_end_period(t *testing.T) {
	t.Helper()

	ret := Calends_create_string_end_period(C.CString(""), C.CString(""), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_string_end_period(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", "", "", "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_long_long(t *testing.T) {
	t.Helper()

	ret := Calends_create_long_long(C.longlong(0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", 0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_long_long_range(t *testing.T) {
	t.Helper()

	ret := Calends_create_long_long_range(C.longlong(0), C.longlong(0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_long_long_range(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", 0, 0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_long_long_start_period(t *testing.T) {
	t.Helper()

	ret := Calends_create_long_long_start_period(C.longlong(0), C.longlong(0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_long_long_start_period(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", 0, 0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_long_long_end_period(t *testing.T) {
	t.Helper()

	ret := Calends_create_long_long_end_period(C.longlong(0), C.longlong(0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_long_long_end_period(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", 0, 0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_double(t *testing.T) {
	t.Helper()

	ret := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", 0.0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_double_range(t *testing.T) {
	t.Helper()

	ret := Calends_create_double_range(C.double(0.0), C.double(0.0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_double_range(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", 0.0, 0.0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_double_start_period(t *testing.T) {
	t.Helper()

	ret := Calends_create_double_start_period(C.double(0.0), C.double(0.0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_double_start_period(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", 0.0, 0.0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_create_double_end_period(t *testing.T) {
	t.Helper()

	ret := Calends_create_double_end_period(C.double(0.0), C.double(0.0), C.CString(""), C.CString(""))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_create_double_end_period(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", 0.0, 0.0, "", "", uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_date(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	ret := Calends_date(inst, C.CString(""), C.CString(""))
	if C.GoString(ret) != "1.994980000" {
		t.Errorf("Calends_date(%#v, %#v, %#v) returned %#v; wanted %#v", uint64(inst), "", "", C.GoString(ret), "1.994980000")
	}
	Calends_release(inst)
}

func testCalends_duration(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	ret := Calends_duration(inst)
	if C.GoString(ret) != "0" {
		t.Errorf("Calends_duration(%#v) returned %#v; wanted %#v", uint64(inst), C.GoString(ret), "0")
	}
	Calends_release(inst)
}

func testCalends_end_date(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	ret := Calends_end_date(inst, C.CString(""), C.CString(""))
	if C.GoString(ret) != "1.994980000" {
		t.Errorf("Calends_end_date(%#v, %#v, %#v) returned %#v; wanted %#v", uint64(inst), "", "", C.GoString(ret), "1.994980000")
	}
	Calends_release(inst)
}

func testCalends_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	ret := Calends_string(inst)
	if C.GoString(ret) != "3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046" {
		t.Errorf("Calends_string(%#v) returned %#v; wanted %#v", uint64(inst), C.GoString(ret), "3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046")
	}
	Calends_release(inst)
}

func testCalends_encode_text(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	ret := Calends_encode_text(inst)
	if C.GoString(ret) != "3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046" {
		t.Errorf("Calends_encode_text(%#v) returned %#v; wanted %#v", uint64(inst), C.GoString(ret), "3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046")
	}
	Calends_release(inst)
}

func testCalends_decode_text(t *testing.T) {
	t.Helper()

	in := "3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046"
	ret := Calends_decode_text(C.CString(in))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_decode_text(%#v) returned invalid Calends object ID %#v", in, uint64(ret))
	}
	Calends_release(ret)
}

func testCalends_encode_json(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	ret := Calends_encode_json(inst)
	if C.GoString(ret) != `"3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046"` {
		t.Errorf("Calends_encode_json(%#v) returned %#v; wanted %#v", uint64(inst), C.GoString(ret), `"3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046"`)
	}
	Calends_release(inst)
}

func testCalends_decode_json(t *testing.T) {
	t.Helper()

	in := `"3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046"`
	ret := Calends_decode_json(C.CString(in))
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_decode_json(%#v) returned invalid Calends object ID %#v", in, uint64(ret))
	}
	Calends_release(ret)
}

func testNo_registered_panic_handler(t *testing.T) {
	t.Helper()

	defer handlePanic()
	main()
}

func testCalends_register_panic_handler(t *testing.T) {
	t.Helper()

	defer handlePanic()
	Calends_register_panic_handler(C.Calends_panic_handler(C.test_Calends_panic_handler))
	main()
}

func testUnknown_error_in_panic_handler(t *testing.T) {
	t.Helper()

	defer handlePanic()
	panic(t)
}
