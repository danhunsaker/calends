#ifndef WRAP_LIBCALENDS_H
#define WRAP_LIBCALENDS_H

#include "libcalends.h"

double ext_Calends_create_string(zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_string_range(zval *start_param, zval *end_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_string_start_period(zval *start_param, zval *duration_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_string_end_period(zval *duration_param, zval *end_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_long_long(zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_long_long_range(zval *start_param, zval *end_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_long_long_start_period(zval *start_param, zval *duration_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_long_long_end_period(zval *duration_param, zval *end_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_double(zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_double_range(zval *start_param, zval *end_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_double_start_period(zval *start_param, zval *duration_param, zval *calendar_param, zval *format_param);
double ext_Calends_create_double_end_period(zval *duration_param, zval *end_param, zval *calendar_param, zval *format_param);
zval ext_Calends_date(zval *p_param, zval *calendar_param, zval *format_param);
zval ext_Calends_duration(zval *p_param);
zval ext_Calends_end_date(zval *p_param, zval *calendar_param, zval *format_param);
zval ext_Calends_string(zval *p_param);
zval ext_Calends_encode_text(zval *p_param);
double ext_Calends_decode_text(zval *encoded_param);
zval ext_Calends_encode_json(zval *p_param);
double ext_Calends_decode_json(zval *encoded_param);
zval ext_Calends_difference(zval *p_param, zval *z_param, zval *mode_param);
zval ext_Calends_compare(zval *p_param, zval *z_param, zval *mode_param);
zval ext_Calends_is_same(zval *p_param, zval *z_param);
zval ext_Calends_is_same_duration(zval *p_param, zval *z_param);
zval ext_Calends_is_shorter(zval *p_param, zval *z_param);
zval ext_Calends_is_longer(zval *p_param, zval *z_param);
zval ext_Calends_contains(zval *p_param, zval *z_param);
zval ext_Calends_overlaps(zval *p_param, zval *z_param);
zval ext_Calends_abuts(zval *p_param, zval *z_param);
zval ext_Calends_is_before(zval *p_param, zval *z_param);
zval ext_Calends_starts_before(zval *p_param, zval *z_param);
zval ext_Calends_ends_before(zval *p_param, zval *z_param);
zval ext_Calends_is_during(zval *p_param, zval *z_param);
zval ext_Calends_starts_during(zval *p_param, zval *z_param);
zval ext_Calends_ends_during(zval *p_param, zval *z_param);
zval ext_Calends_is_after(zval *p_param, zval *z_param);
zval ext_Calends_starts_after(zval *p_param, zval *z_param);
zval ext_Calends_ends_after(zval *p_param, zval *z_param);
double ext_Calends_add_string(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_add_long_long(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_add_double(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_subtract_string(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_subtract_long_long(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_subtract_double(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_add_from_end_string(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_add_from_end_long_long(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_add_from_end_double(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_subtract_from_end_string(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_subtract_from_end_long_long(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_subtract_from_end_double(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_next_string(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_next_long_long(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_next_double(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_previous_string(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_previous_long_long(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_previous_double(zval *p_param, zval *offset_param, zval *calendar_param);
double ext_Calends_with_date_string(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_with_date_long_long(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_with_date_double(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_with_end_date_string(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_with_end_date_long_long(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_with_end_date_double(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param);
double ext_Calends_with_duration_string(zval *p_param, zval *duration_param, zval *calendar_param);
double ext_Calends_with_duration_long_long(zval *p_param, zval *duration_param, zval *calendar_param);
double ext_Calends_with_duration_double(zval *p_param, zval *duration_param, zval *calendar_param);
double ext_Calends_with_duration_from_end_string(zval *p_param, zval *duration_param, zval *calendar_param);
double ext_Calends_with_duration_from_end_long_long(zval *p_param, zval *duration_param, zval *calendar_param);
double ext_Calends_with_duration_from_end_double(zval *p_param, zval *duration_param, zval *calendar_param);
double ext_Calends_merge(zval *p_param, zval *z_param);
double ext_Calends_intersect(zval *p_param, zval *z_param);
double ext_Calends_gap(zval *p_param, zval *z_param);

void ext_Calends_release(zval *p_param);

TAI64Time ext_Calends_calendar_to_internal_string(char* name, char* date, char* format);
TAI64Time ext_Calends_calendar_to_internal_long_long(char* name, long long int date, char* format);
TAI64Time ext_Calends_calendar_to_internal_double(char* name, double date, char* format);
TAI64Time ext_Calends_calendar_to_internal_tai(char* name, TAI64Time date);

char* ext_Calends_calendar_from_internal(char* name, TAI64Time stamp, char* format);

TAI64Time ext_Calends_calendar_offset_string(char* name, TAI64Time stamp, char* offset);
TAI64Time ext_Calends_calendar_offset_long_long(char* name, TAI64Time stamp, long long int offset);
TAI64Time ext_Calends_calendar_offset_double(char* name, TAI64Time stamp, double offset);
TAI64Time ext_Calends_calendar_offset_tai(char* name, TAI64Time stamp, TAI64Time offset);

void ext_Calends_calendar_register(zval *calendar_param, zval *default_format_param);
void ext_Calends_calendar_unregister(zval *calendar_param);
zval ext_Calends_calendar_registered(zval *calendar_param);
zval ext_Calends_calendar_list_registered();

zval ext_TAI64Time_add(zval *p_param, zval *z_param);
zval ext_TAI64Time_sub(zval *p_param, zval *z_param);
zval ext_TAI64Time_string(zval *p_param);
zval ext_TAI64Time_from_string(zval *in_param);
zval ext_TAI64Time_hex_string(zval *p_param);
zval ext_TAI64Time_from_hex_string(zval *in_param);
zval ext_TAI64Time_double(zval *p_param);
zval ext_TAI64Time_from_double(zval *in_param);
zval ext_TAI64Time_encode_text(zval *p_param);
zval ext_TAI64Time_decode_text(zval *encoded_param);
zval ext_TAI64Time_encode_binary(zval *p_param);
zval ext_TAI64Time_decode_binary(zval *encoded_param, zval *len_param);
zval ext_TAI64Time_utc_to_tai(zval *p_param);
zval ext_TAI64Time_tai_to_utc(zval *p_param);

// helpers, rather than wrappers
void ext_panic_handler(char *error_message);
void ext_unregister_all_calendars();

void ZVAL_TAITIME(zval *out, TAI64Time tai_val);
TAI64Time Z_TAITIME(zval in);

#endif
