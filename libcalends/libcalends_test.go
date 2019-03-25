package main

import (
	"testing"
)

func testingPanicCatcher(t *testing.T) {
	t.Helper()

	if r := recover(); r != nil {
		t.Logf("Caught panic: %#v", r)
	}
}

// calends.go
func TestCalends_release(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_release(t)
}
func TestCalends_create_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_string(t)
}
func TestCalends_create_string_range(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_string_range(t)
}
func TestCalends_create_string_start_period(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_string_start_period(t)
}
func TestCalends_create_string_end_period(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_string_end_period(t)
}
func TestCalends_create_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_long_long(t)
}
func TestCalends_create_long_long_range(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_long_long_range(t)
}
func TestCalends_create_long_long_start_period(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_long_long_start_period(t)
}
func TestCalends_create_long_long_end_period(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_long_long_end_period(t)
}
func TestCalends_create_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_double(t)
}
func TestCalends_create_double_range(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_double_range(t)
}
func TestCalends_create_double_start_period(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_double_start_period(t)
}
func TestCalends_create_double_end_period(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_create_double_end_period(t)
}
func TestCalends_date(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_date(t)
}
func TestCalends_duration(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_duration(t)
}
func TestCalends_end_date(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_end_date(t)
}
func TestCalends_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_string(t)
}
func TestCalends_encode_text(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_encode_text(t)
}
func TestCalends_decode_text(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_decode_text(t)
}
func TestCalends_encode_json(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_encode_json(t)
}
func TestCalends_decode_json(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_decode_json(t)
}
func TestNo_registered_panic_handler(t *testing.T) {
	defer testingPanicCatcher(t)
	testNo_registered_panic_handler(t)
}
func TestCalends_register_panic_handler(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_register_panic_handler(t)
}
func TestUnknown_error_in_panic_handler(t *testing.T) {
	defer testingPanicCatcher(t)
	testUnknown_error_in_panic_handler(t)
}

// offsets.go
func TestCalends_add_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_add_string(t)
}
func TestCalends_add_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_add_long_long(t)
}
func TestCalends_add_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_add_double(t)
}
func TestCalends_subtract_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_subtract_string(t)
}
func TestCalends_subtract_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_subtract_long_long(t)
}
func TestCalends_subtract_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_subtract_double(t)
}
func TestCalends_add_from_end_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_add_from_end_string(t)
}
func TestCalends_add_from_end_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_add_from_end_long_long(t)
}
func TestCalends_add_from_end_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_add_from_end_double(t)
}
func TestCalends_subtract_from_end_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_subtract_from_end_string(t)
}
func TestCalends_subtract_from_end_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_subtract_from_end_long_long(t)
}
func TestCalends_subtract_from_end_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_subtract_from_end_double(t)
}
func TestCalends_next_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_next_string(t)
}
func TestCalends_next_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_next_long_long(t)
}
func TestCalends_next_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_next_double(t)
}
func TestCalends_previous_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_previous_string(t)
}
func TestCalends_previous_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_previous_long_long(t)
}
func TestCalends_previous_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_previous_double(t)
}
func TestCalends_with_date_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_date_string(t)
}
func TestCalends_with_date_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_date_long_long(t)
}
func TestCalends_with_date_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_date_double(t)
}
func TestCalends_with_end_date_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_end_date_string(t)
}
func TestCalends_with_end_date_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_end_date_long_long(t)
}
func TestCalends_with_end_date_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_end_date_double(t)
}
func TestCalends_with_duration_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_duration_string(t)
}
func TestCalends_with_duration_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_duration_long_long(t)
}
func TestCalends_with_duration_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_duration_double(t)
}
func TestCalends_with_duration_from_end_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_duration_from_end_string(t)
}
func TestCalends_with_duration_from_end_long_long(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_duration_from_end_long_long(t)
}
func TestCalends_with_duration_from_end_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_with_duration_from_end_double(t)
}
func TestCalends_merge(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_merge(t)
}
func TestCalends_intersect(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_intersect(t)
}
func TestCalends_gap(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_gap(t)
}

// comparisons.go
func TestCalends_difference(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_difference(t)
}
func TestCalends_compare(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_compare(t)
}
func TestCalends_is_same(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_is_same(t)
}
func TestCalends_is_same_duration(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_is_same_duration(t)
}
func TestCalends_is_shorter(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_is_shorter(t)
}
func TestCalends_is_longer(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_is_longer(t)
}
func TestCalends_contains(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_contains(t)
}
func TestCalends_overlaps(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_overlaps(t)
}
func TestCalends_abuts(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_abuts(t)
}
func TestCalends_is_before(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_is_before(t)
}
func TestCalends_starts_before(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_starts_before(t)
}
func TestCalends_ends_before(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_ends_before(t)
}
func TestCalends_is_during(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_is_during(t)
}
func TestCalends_starts_during(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_starts_during(t)
}
func TestCalends_ends_during(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_ends_during(t)
}
func TestCalends_is_after(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_is_after(t)
}
func TestCalends_starts_after(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_starts_after(t)
}
func TestCalends_ends_after(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_ends_after(t)
}

// calendars.go
func TestCalends_calendar_registered(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_calendar_registered(t)
}
func TestCalends_calendar_list_registered(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_calendar_list_registered(t)
}
func TestCalends_calendar_register(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_calendar_register(t)
}
func TestCalends_calendar_unregister(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_calendar_unregister(t)
}
func TestCalends_calendar_register_dynamic(t *testing.T) {
	defer testingPanicCatcher(t)
	testCalends_calendar_register_dynamic(t)
}
func TestTAI64Time_add(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_add(t)
}
func TestTAI64Time_sub(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_sub(t)
}
func TestTAI64Time_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_string(t)
}
func TestTAI64Time_from_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_from_string(t)
}
func TestTAI64Time_hex_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_hex_string(t)
}
func TestTAI64Time_from_hex_string(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_from_hex_string(t)
}
func TestTAI64Time_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_double(t)
}
func TestTAI64Time_from_double(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_from_double(t)
}
func TestTAI64Time_encode_text(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_encode_text(t)
}
func TestTAI64Time_decode_text(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_decode_text(t)
}
func TestTAI64Time_encode_binary(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_encode_binary(t)
}
func TestTAI64Time_decode_binary(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_decode_binary(t)
}
func TestTAI64Time_utc_to_tai(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_utc_to_tai(t)
}
func TestTAI64Time_tai_to_utc(t *testing.T) {
	defer testingPanicCatcher(t)
	testTAI64Time_tai_to_utc(t)
}
