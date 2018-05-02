package main

import "C"

import (
	"testing"
)

func testCalends_add_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_add_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_add_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_add_string(inst, C.CString("86400"), C.CString("invalid"))
}

func testCalends_add_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_add_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_add_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_add_long_long(inst, C.longlong(86400), C.CString("invalid"))
}

func testCalends_add_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_add_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_add_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_add_double(inst, C.double(86400.0), C.CString("invalid"))
}

func testCalends_subtract_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_subtract_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_subtract_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_subtract_string(inst, C.CString("86400"), C.CString("invalid"))
}

func testCalends_subtract_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_subtract_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_subtract_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_subtract_long_long(inst, C.longlong(86400), C.CString("invalid"))
}

func testCalends_subtract_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_subtract_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_subtract_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_subtract_double(inst, C.double(86400.0), C.CString("invalid"))
}

func testCalends_add_from_end_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_add_from_end_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_add_from_end_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_add_from_end_string(inst, C.CString("86400"), C.CString("invalid"))
}

func testCalends_add_from_end_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_add_from_end_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_add_from_end_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_add_from_end_long_long(inst, C.longlong(86400), C.CString("invalid"))
}

func testCalends_add_from_end_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_add_from_end_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_add_from_end_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_add_from_end_double(inst, C.double(86400.0), C.CString("invalid"))
}

func testCalends_subtract_from_end_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_subtract_from_end_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_subtract_from_end_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_subtract_from_end_string(inst, C.CString("86400"), C.CString("invalid"))
}

func testCalends_subtract_from_end_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_subtract_from_end_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_subtract_from_end_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_subtract_from_end_long_long(inst, C.longlong(86400), C.CString("invalid"))
}

func testCalends_subtract_from_end_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_subtract_from_end_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_subtract_from_end_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_subtract_from_end_double(inst, C.double(86400.0), C.CString("invalid"))
}

func testCalends_next_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_next_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_next_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_next_string(inst, C.CString("86400"), C.CString("invalid"))
}

func testCalends_next_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_next_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_next_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_next_long_long(inst, C.longlong(86400), C.CString("invalid"))
}

func testCalends_next_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_next_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_next_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_next_double(inst, C.double(86400.0), C.CString("invalid"))
}

func testCalends_previous_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_previous_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_previous_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_previous_string(inst, C.CString("86400"), C.CString("invalid"))
}

func testCalends_previous_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_previous_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_previous_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_previous_long_long(inst, C.longlong(86400), C.CString("invalid"))
}

func testCalends_previous_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_previous_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_previous_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_previous_double(inst, C.double(86400.0), C.CString("invalid"))
}

func testCalends_with_date_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_date_string(inst, C.CString(""), C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_date_string(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", "", uint64(ret))
	}
	Calends_with_date_string(inst, C.CString(""), C.CString("invalid"), C.CString(""))
}

func testCalends_with_date_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_date_long_long(inst, C.longlong(0), C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_date_long_long(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", "", uint64(ret))
	}
	Calends_with_date_long_long(inst, C.longlong(0), C.CString("invalid"), C.CString(""))
}

func testCalends_with_date_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_date_double(inst, C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_date_double(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", "", uint64(ret))
	}
	Calends_with_date_double(inst, C.double(0.0), C.CString("invalid"), C.CString(""))
}

func testCalends_with_end_date_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_end_date_string(inst, C.CString(""), C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_end_date_string(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", "", uint64(ret))
	}
	Calends_with_end_date_string(inst, C.CString(""), C.CString("invalid"), C.CString(""))
}

func testCalends_with_end_date_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_end_date_long_long(inst, C.longlong(0), C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_end_date_long_long(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", "", uint64(ret))
	}
	Calends_with_end_date_long_long(inst, C.longlong(0), C.CString("invalid"), C.CString(""))
}

func testCalends_with_end_date_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_end_date_double(inst, C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_end_date_double(%#v, %#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", "", uint64(ret))
	}
	Calends_with_end_date_double(inst, C.double(0.0), C.CString("invalid"), C.CString(""))
}

func testCalends_with_duration_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_duration_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_duration_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_with_duration_string(inst, C.CString(""), C.CString("invalid"))
}

func testCalends_with_duration_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_duration_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_duration_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_with_duration_long_long(inst, C.longlong(0), C.CString("invalid"))
}

func testCalends_with_duration_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_duration_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_duration_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_with_duration_double(inst, C.double(0.0), C.CString("invalid"))
}

func testCalends_with_duration_from_end_string(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_duration_from_end_string(inst, C.CString(""), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_duration_from_end_string(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), "", "", uint64(ret))
	}
	Calends_with_duration_from_end_string(inst, C.CString(""), C.CString("invalid"))
}

func testCalends_with_duration_from_end_long_long(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_duration_from_end_long_long(inst, C.longlong(0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_duration_from_end_long_long(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0, "", uint64(ret))
	}
	Calends_with_duration_from_end_long_long(inst, C.longlong(0), C.CString("invalid"))
}

func testCalends_with_duration_from_end_double(t *testing.T) {
	t.Helper()

	inst := Calends_create_double(C.double(0.0), C.CString(""), C.CString(""))
	defer Calends_release(inst)
	ret := Calends_with_duration_from_end_double(inst, C.double(0.0), C.CString(""))
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_with_duration_from_end_double(%#v, %#v, %#v) returned invalid Calends object ID %#v", uint64(inst), 0.0, "", uint64(ret))
	}
	Calends_with_duration_from_end_double(inst, C.double(0.0), C.CString("invalid"))
}

func testCalends_merge(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	defer Calends_release(inst1)
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	defer Calends_release(inst2)
	ret := Calends_merge(inst1, inst2)
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_merge(%#v, %#v) returned invalid Calends object ID %#v", uint64(inst1), uint64(inst2), uint64(ret))
	}
}

func testCalends_intersect(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.1), C.CString(""), C.CString(""))
	defer Calends_release(inst1)
	inst2 := Calends_create_double_range(C.double(5.0), C.double(10.0), C.CString(""), C.CString(""))
	defer Calends_release(inst2)
	inst3 := Calends_create_double_range(C.double(0.0), C.double(4.9), C.CString(""), C.CString(""))
	defer Calends_release(inst3)
	ret := Calends_intersect(inst1, inst2)
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_intersect(%#v, %#v) returned invalid Calends object ID %#v", uint64(inst1), uint64(inst2), uint64(ret))
	}
	Calends_intersect(inst3, inst2)
}

func testCalends_gap(t *testing.T) {
	t.Helper()

	inst1 := Calends_create_double_range(C.double(0.0), C.double(5.0), C.CString(""), C.CString(""))
	defer Calends_release(inst1)
	inst2 := Calends_create_double_range(C.double(5.1), C.double(10.0), C.CString(""), C.CString(""))
	defer Calends_release(inst2)
	inst3 := Calends_create_double_range(C.double(4.9), C.double(10.0), C.CString(""), C.CString(""))
	defer Calends_release(inst3)
	ret := Calends_gap(inst1, inst2)
	defer Calends_release(ret)
	if _, ok := instances.Load(uint64(ret)); uint64(ret) < 1 || !ok {
		t.Errorf("Calends_gap(%#v, %#v) returned invalid Calends object ID %#v", uint64(inst1), uint64(inst2), uint64(ret))
	}
	Calends_gap(inst1, inst3)
}
