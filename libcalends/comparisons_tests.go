package main

import "C"

import (
	"testing"
)

func testCalends_difference(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_difference(inst1, inst2, C.CString("start"))
	if C.GoString(ret) != "-5" {
		t.Errorf("Calends_difference(%#v, %#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), "start", C.GoString(ret), "-5")
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_compare(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_compare(inst1, inst2, C.CString("start"))
	if int8(ret) != -1 {
		t.Errorf("Calends_compare(%#v, %#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), "start", int8(ret), -1)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_is_same(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_is_same(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_is_same(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_is_same_duration(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_is_same_duration(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_is_same_duration(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_is_shorter(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_is_shorter(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_is_shorter(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_is_longer(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_is_longer(inst1, inst2)
	if bool(ret) != true {
		t.Errorf("Calends_is_longer(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), true)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_contains(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_contains(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_contains(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_overlaps(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_overlaps(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_overlaps(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_abuts(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_abuts(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_abuts(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_is_before(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_is_before(inst1, inst2)
	if bool(ret) != true {
		t.Errorf("Calends_is_before(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), true)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_starts_before(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_starts_before(inst1, inst2)
	if bool(ret) != true {
		t.Errorf("Calends_starts_before(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), true)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_ends_before(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_ends_before(inst1, inst2)
	if bool(ret) != true {
		t.Errorf("Calends_ends_before(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), true)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_is_during(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_is_during(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_is_during(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_starts_during(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_starts_during(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_starts_during(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_ends_during(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_ends_during(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_ends_during(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_is_after(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_is_after(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_is_after(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_starts_after(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_starts_after(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_starts_after(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}

func testCalends_ends_after(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	ret := Calends_ends_after(inst1, inst2)
	if bool(ret) != false {
		t.Errorf("Calends_ends_after(%#v, %#v) returned %#v; wanted %#v", uint64(inst1), uint64(inst2), bool(ret), false)
	}
	Calends_release(inst1)
	Calends_release(inst2)
}
