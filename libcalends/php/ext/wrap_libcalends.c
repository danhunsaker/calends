#ifdef HAVE_CONFIG_H
#include "config.h"
#endif

#include "php.h"
#include "php_ext.h"
#include "kernel/array.h"
#include "kernel/exception.h"
#include "kernel/fcall.h"
#include "kernel/object.h"
#include "kernel/memory.h"
#include "calends/calends.zep.h"
#include "calends/calendsexception.zep.h"
#include "calends/taitime.zep.h"
#include "wrap_libcalends.h"

double ext_Calends_create_string(zval *date_param, zval *calendar_param, zval *format_param) {
  char *date, *calendar, *format;

  date = ZSTR_VAL(Z_STR(*date_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_string(date, calendar, format);
}

double ext_Calends_create_string_range(zval *start_param, zval *end_param, zval *calendar_param, zval *format_param) {
  char *start, *end, *calendar, *format;

  start = ZSTR_VAL(Z_STR(*start_param));
  end = ZSTR_VAL(Z_STR(*end_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_string_range(start, end, calendar, format);
}

double ext_Calends_create_string_start_period(zval *start_param, zval *duration_param, zval *calendar_param, zval *format_param) {
  char *start, *duration, *calendar, *format;

  start = ZSTR_VAL(Z_STR(*start_param));
  duration = ZSTR_VAL(Z_STR(*duration_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_string_start_period(start, duration, calendar, format);
}

double ext_Calends_create_string_end_period(zval *duration_param, zval *end_param, zval *calendar_param, zval *format_param) {
  char *duration, *end, *calendar, *format;

  duration = ZSTR_VAL(Z_STR(*duration_param));
  end = ZSTR_VAL(Z_STR(*end_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_string_end_period(duration, end, calendar, format);
}

double ext_Calends_create_long_long(zval *date_param, zval *calendar_param, zval *format_param) {
  long long int date;
  char *calendar, *format;

  date = Z_LVAL(*date_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_long_long(date, calendar, format);
}

double ext_Calends_create_long_long_range(zval *start_param, zval *end_param, zval *calendar_param, zval *format_param) {
  long long int start, end;
  char *calendar, *format;

  start = Z_LVAL(*start_param);
  end = Z_LVAL(*end_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_long_long_range(start, end, calendar, format);
}

double ext_Calends_create_long_long_start_period(zval *start_param, zval *duration_param, zval *calendar_param, zval *format_param) {
  long long int start, duration;
  char *calendar, *format;

  start = Z_LVAL(*start_param);
  duration = Z_LVAL(*duration_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_long_long_start_period(start, duration, calendar, format);
}

double ext_Calends_create_long_long_end_period(zval *duration_param, zval *end_param, zval *calendar_param, zval *format_param) {
  long long int duration, end;
  char *calendar, *format;

  duration = Z_LVAL(*duration_param);
  end = Z_LVAL(*end_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_long_long_end_period(duration, end, calendar, format);
}

double ext_Calends_create_double(zval *date_param, zval *calendar_param, zval *format_param) {
  double date;
  char *calendar, *format;

  date = Z_DVAL(*date_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_double(date, calendar, format);
}

double ext_Calends_create_double_range(zval *start_param, zval *end_param, zval *calendar_param, zval *format_param) {
  double start, end;
  char *calendar, *format;

  start = Z_DVAL(*start_param);
  end = Z_DVAL(*end_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_double_range(start, end, calendar, format);
}

double ext_Calends_create_double_start_period(zval *start_param, zval *duration_param, zval *calendar_param, zval *format_param) {
  double start, duration;
  char *calendar, *format;

  start = Z_DVAL(*start_param);
  duration = Z_DVAL(*duration_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_double_start_period(start, duration, calendar, format);
}

double ext_Calends_create_double_end_period(zval *duration_param, zval *end_param, zval *calendar_param, zval *format_param) {
  double duration, end;
  char *calendar, *format;

  duration = Z_DVAL(*duration_param);
  end = Z_DVAL(*end_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_create_double_end_period(duration, end, calendar, format);
}

zval ext_Calends_date(zval *p_param, zval *calendar_param, zval *format_param) {
  double p;
  char *calendar, *format;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));
	ZVAL_STRING(&out, Calends_date((long long unsigned int)p, calendar, format));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_duration(zval *p_param) {
  double p;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
	ZVAL_STRING(&out, Calends_duration((long long unsigned int)p));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_end_date(zval *p_param, zval *calendar_param, zval *format_param) {
  double p;
  char *calendar, *format;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));
	ZVAL_STRING(&out, Calends_end_date((long long unsigned int)p, calendar, format));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_string(zval *p_param) {
  double p;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
	ZVAL_STRING(&out, Calends_string((long long unsigned int)p));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_encode_text(zval *p_param) {
  double p;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
	ZVAL_STRING(&out, Calends_encode_text((long long unsigned int)p));

  ZEPHIR_MM_RESTORE();

  return out;
}

double ext_Calends_decode_text(zval *encoded_param) {
  char *encoded;

  encoded = ZSTR_VAL(Z_STR(*encoded_param));

	return Calends_decode_text(encoded);
}

zval ext_Calends_encode_json(zval *p_param) {
  double p;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
	ZVAL_STRING(&out, Calends_encode_json((long long unsigned int)p));

  ZEPHIR_MM_RESTORE();

  return out;
}

double ext_Calends_decode_json(zval *encoded_param) {
  char *encoded;

  encoded = ZSTR_VAL(Z_STR(*encoded_param));

	return Calends_decode_json(encoded);
}

zval ext_Calends_difference(zval *p_param, zval *z_param, zval *mode_param) {
  double p, z;
  char *mode;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  mode = ZSTR_VAL(Z_STR(*mode_param));
  ZVAL_STRING(&out, Calends_difference((long long unsigned int)p, (long long unsigned int)z, mode));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_compare(zval *p_param, zval *z_param, zval *mode_param) {
  double p, z;
  char *mode;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  mode = ZSTR_VAL(Z_STR(*mode_param));
  ZVAL_LONG(&out, Calends_compare((long long unsigned int)p, (long long unsigned int)z, mode));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_is_same(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_is_same((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_is_same_duration(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_is_same_duration((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_is_shorter(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_is_shorter((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_is_longer(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_is_longer((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_contains(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_contains((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_overlaps(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_overlaps((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_abuts(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_abuts((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_is_before(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_is_before((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_starts_before(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_starts_before((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_ends_before(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_ends_before((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_is_during(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_is_during((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_starts_during(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_starts_during((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_ends_during(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_ends_during((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_is_after(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_is_after((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_starts_after(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
  ZVAL_BOOL(&out, Calends_starts_after((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_ends_after(zval *p_param, zval *z_param) {
  double p, z;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);
	ZVAL_BOOL(&out, Calends_ends_after((long long unsigned int)p, (long long unsigned int)z));

  ZEPHIR_MM_RESTORE();

  return out;
}

double ext_Calends_add_string(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  char *offset, *calendar;

  p = Z_DVAL(*p_param);
  offset = ZSTR_VAL(Z_STR(*offset_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_add_string((long long unsigned int)p, offset, calendar);
}

double ext_Calends_add_long_long(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  long long int offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_LVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_add_long_long((long long unsigned int)p, offset, calendar);
}

double ext_Calends_add_double(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p, offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_DVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_add_double((long long unsigned int)p, offset, calendar);
}

double ext_Calends_subtract_string(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  char *offset, *calendar;

  p = Z_DVAL(*p_param);
  offset = ZSTR_VAL(Z_STR(*offset_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_subtract_string((long long unsigned int)p, offset, calendar);
}

double ext_Calends_subtract_long_long(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  long long int offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_LVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_subtract_long_long((long long unsigned int)p, offset, calendar);
}

double ext_Calends_subtract_double(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p, offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_DVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_subtract_double((long long unsigned int)p, offset, calendar);
}

double ext_Calends_add_from_end_string(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  char *offset, *calendar;

  p = Z_DVAL(*p_param);
  offset = ZSTR_VAL(Z_STR(*offset_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_add_from_end_string((long long unsigned int)p, offset, calendar);
}

double ext_Calends_add_from_end_long_long(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  long long int offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_LVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_add_from_end_long_long((long long unsigned int)p, offset, calendar);
}

double ext_Calends_add_from_end_double(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p, offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_DVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_add_from_end_double((long long unsigned int)p, offset, calendar);
}

double ext_Calends_subtract_from_end_string(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  char *offset, *calendar;

  p = Z_DVAL(*p_param);
  offset = ZSTR_VAL(Z_STR(*offset_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_subtract_from_end_string((long long unsigned int)p, offset, calendar);
}

double ext_Calends_subtract_from_end_long_long(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  long long int offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_LVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_subtract_from_end_long_long((long long unsigned int)p, offset, calendar);
}

double ext_Calends_subtract_from_end_double(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p, offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_DVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_subtract_from_end_double((long long unsigned int)p, offset, calendar);
}

double ext_Calends_next_string(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  char *offset, *calendar;

  p = Z_DVAL(*p_param);
  offset = ZSTR_VAL(Z_STR(*offset_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_next_string((long long unsigned int)p, offset, calendar);
}

double ext_Calends_next_long_long(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  long long int offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_LVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_next_long_long((long long unsigned int)p, offset, calendar);
}

double ext_Calends_next_double(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p, offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_DVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_next_double((long long unsigned int)p, offset, calendar);
}

double ext_Calends_previous_string(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  char *offset, *calendar;

  p = Z_DVAL(*p_param);
  offset = ZSTR_VAL(Z_STR(*offset_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_previous_string((long long unsigned int)p, offset, calendar);
}

double ext_Calends_previous_long_long(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p;
  long long int offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_LVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_previous_long_long((long long unsigned int)p, offset, calendar);
}

double ext_Calends_previous_double(zval *p_param, zval *offset_param, zval *calendar_param) {
  double p, offset;
  char *calendar;

  p = Z_DVAL(*p_param);
  offset = Z_DVAL(*offset_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_previous_double((long long unsigned int)p, offset, calendar);
}

double ext_Calends_with_date_string(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param) {
  double p;
  char *date, *calendar, *format;

  p = Z_DVAL(*p_param);
  date = ZSTR_VAL(Z_STR(*date_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_with_date_string((long long unsigned int)p, date, calendar, format);
}

double ext_Calends_with_date_long_long(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param) {
  double p;
  long long int date;
  char *calendar, *format;

  p = Z_DVAL(*p_param);
  date = Z_LVAL(*date_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_with_date_long_long((long long unsigned int)p, date, calendar, format);
}

double ext_Calends_with_date_double(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param) {
  double p, date;
  char *calendar, *format;

  p = Z_DVAL(*p_param);
  date = Z_DVAL(*date_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_with_date_double((long long unsigned int)p, date, calendar, format);
}

double ext_Calends_with_end_date_string(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param) {
  double p;
  char *date, *calendar, *format;

  p = Z_DVAL(*p_param);
  date = ZSTR_VAL(Z_STR(*date_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_with_end_date_string((long long unsigned int)p, date, calendar, format);
}

double ext_Calends_with_end_date_long_long(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param) {
  double p;
  long long int date;
  char *calendar, *format;

  p = Z_DVAL(*p_param);
  date = Z_LVAL(*date_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_with_end_date_long_long((long long unsigned int)p, date, calendar, format);
}

double ext_Calends_with_end_date_double(zval *p_param, zval *date_param, zval *calendar_param, zval *format_param) {
  double p, date;
  char *calendar, *format;

  p = Z_DVAL(*p_param);
  date = Z_DVAL(*date_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  format = ZSTR_VAL(Z_STR(*format_param));

	return (double)Calends_with_end_date_double((long long unsigned int)p, date, calendar, format);
}

double ext_Calends_with_duration_string(zval *p_param, zval *duration_param, zval *calendar_param) {
  double p;
  char *duration, *calendar;

  p = Z_DVAL(*p_param);
  duration = ZSTR_VAL(Z_STR(*duration_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_with_duration_string((long long unsigned int)p, duration, calendar);
}

double ext_Calends_with_duration_long_long(zval *p_param, zval *duration_param, zval *calendar_param) {
  double p;
  long long int duration;
  char *calendar;

  p = Z_DVAL(*p_param);
  duration = Z_LVAL(*duration_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_with_duration_long_long((long long unsigned int)p, duration, calendar);
}

double ext_Calends_with_duration_double(zval *p_param, zval *duration_param, zval *calendar_param) {
  double p, duration;
  char *calendar;

  p = Z_DVAL(*p_param);
  duration = Z_DVAL(*duration_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_with_duration_double((long long unsigned int)p, duration, calendar);
}

double ext_Calends_with_duration_from_end_string(zval *p_param, zval *duration_param, zval *calendar_param) {
  double p;
  char *duration, *calendar;

  p = Z_DVAL(*p_param);
  duration = ZSTR_VAL(Z_STR(*duration_param));
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_with_duration_from_end_string((long long unsigned int)p, duration, calendar);
}

double ext_Calends_with_duration_from_end_long_long(zval *p_param, zval *duration_param, zval *calendar_param) {
  double p;
  long long int duration;
  char *calendar;

  p = Z_DVAL(*p_param);
  duration = Z_LVAL(*duration_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_with_duration_from_end_long_long((long long unsigned int)p, duration, calendar);
}

double ext_Calends_with_duration_from_end_double(zval *p_param, zval *duration_param, zval *calendar_param) {
  double p, duration;
  char *calendar;

  p = Z_DVAL(*p_param);
  duration = Z_DVAL(*duration_param);
  calendar = ZSTR_VAL(Z_STR(*calendar_param));

	return (double)Calends_with_duration_from_end_double((long long unsigned int)p, duration, calendar);
}

double ext_Calends_merge(zval *p_param, zval *z_param) {
  double p, z;

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);

	return (double)Calends_merge((long long unsigned int)p, (long long unsigned int)z);
}

double ext_Calends_intersect(zval *p_param, zval *z_param) {
  double p, z;

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);

	return (double)Calends_intersect((long long unsigned int)p, (long long unsigned int)z);
}

double ext_Calends_gap(zval *p_param, zval *z_param) {
  double p, z;

  p = Z_DVAL(*p_param);
  z = Z_DVAL(*z_param);

	return (double)Calends_gap((long long unsigned int)p, (long long unsigned int)z);
}


void ext_Calends_release(zval *p_param) {
  double p;

  p = Z_DVAL(*p_param);

	Calends_release((long long unsigned int)p);
}


TAI64Time ext_Calends_calendar_to_internal_string(char* name_param, char* date_param, char* format_param) {
  TAI64Time out;
  zval out_zval, name, date, format;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&date);
  ZVAL_UNDEF(&format);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_STRING(&date, date_param);
  ZVAL_STRING(&format, format_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 5 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("toInternal"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 5 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &date);
  zephir_array_fast_append(&args, &format);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}

TAI64Time ext_Calends_calendar_to_internal_long_long(char* name_param, long long int date_param, char* format_param) {
  TAI64Time out;
  zval out_zval, name, date, format;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&date);
  ZVAL_UNDEF(&format);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_DOUBLE(&date, date_param);
  ZVAL_STRING(&format, format_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 5 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("toInternal"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 5 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &date);
  zephir_array_fast_append(&args, &format);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}

TAI64Time ext_Calends_calendar_to_internal_double(char* name_param, double date_param, char* format_param) {
  TAI64Time out;
  zval out_zval, name, date, format;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&date);
  ZVAL_UNDEF(&format);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_DOUBLE(&date, date_param);
  ZVAL_STRING(&format, format_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 5 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("toInternal"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 5 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &date);
  zephir_array_fast_append(&args, &format);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}

TAI64Time ext_Calends_calendar_to_internal_tai(char* name_param, TAI64Time date_param) {
  TAI64Time out;
  zval out_zval, name, date, format;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&date);
  ZVAL_UNDEF(&format);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_TAITIME(&date, date_param);
  ZVAL_STRING(&format, "");

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 5 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("toInternal"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 5 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &date);
  zephir_array_fast_append(&args, &format);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}

char* ext_Calends_calendar_from_internal(char* name_param, TAI64Time stamp_param, char* format_param) {
  char *out;
  zval out_zval, name, stamp, format;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&stamp);
  ZVAL_UNDEF(&format);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_TAITIME(&stamp, stamp_param);
  ZVAL_STRING(&format, format_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 7 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("fromInternal"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 7 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &stamp);
  zephir_array_fast_append(&args, &format);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = ZSTR_VAL(Z_STR(out_zval));
  ZEPHIR_MM_RESTORE();
  return out;
}

TAI64Time ext_Calends_calendar_offset_string(char* name_param, TAI64Time stamp_param, char* offset_param) {
  TAI64Time out;
  zval out_zval, name, stamp, offset;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&stamp);
  ZVAL_UNDEF(&offset);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_TAITIME(&stamp, stamp_param);
  ZVAL_STRING(&offset, offset_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 9 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("offset"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 9 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &stamp);
  zephir_array_fast_append(&args, &offset);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}

TAI64Time ext_Calends_calendar_offset_long_long(char* name_param, TAI64Time stamp_param, long long int offset_param) {
  TAI64Time out;
  zval out_zval, name, stamp, offset;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&stamp);
  ZVAL_UNDEF(&offset);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_TAITIME(&stamp, stamp_param);
  ZVAL_DOUBLE(&offset, offset_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 9 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("offset"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 9 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &stamp);
  zephir_array_fast_append(&args, &offset);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}

TAI64Time ext_Calends_calendar_offset_double(char* name_param, TAI64Time stamp_param, double offset_param) {
  TAI64Time out;
  zval out_zval, name, stamp, offset;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&stamp);
  ZVAL_UNDEF(&offset);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_TAITIME(&stamp, stamp_param);
  ZVAL_DOUBLE(&offset, offset_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 9 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("offset"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 9 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &stamp);
  zephir_array_fast_append(&args, &offset);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

	out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}

TAI64Time ext_Calends_calendar_offset_tai(char* name_param, TAI64Time stamp_param, TAI64Time offset_param) {
  TAI64Time out;
  zval out_zval, name, stamp, offset;
  zval calendar_list, func_list, func, args;
  zend_long ZEPHIR_LAST_CALL_STATUS;

  ZVAL_UNDEF(&out_zval);
  ZVAL_UNDEF(&name);
  ZVAL_UNDEF(&stamp);
  ZVAL_UNDEF(&offset);
  ZVAL_UNDEF(&calendar_list);
  ZVAL_UNDEF(&func_list);
  ZVAL_UNDEF(&func);
  ZVAL_UNDEF(&args);

  ZEPHIR_MM_GROW();

  ZVAL_STRING(&name, name_param);
  ZVAL_TAITIME(&stamp, stamp_param);
  ZVAL_TAITIME(&offset, offset_param);

  zephir_read_static_property_ce(&calendar_list, calends_calends_ce, SL("calendars"), PH_NOISY_CC | PH_READONLY);
	zephir_array_fetch(&func_list, &calendar_list, &name, PH_NOISY, "calends/calendar[object]interface.zep{calendar}", 9 TSRMLS_CC);
	zephir_array_fetch_string(&func, &func_list, SL("offset"), PH_NOISY, "calends/calendar[object]interface.zep{method}", 9 TSRMLS_CC);

  zephir_create_array(&args, 0, 0);
  zephir_array_fast_append(&args, &stamp);
  zephir_array_fast_append(&args, &offset);
  ZEPHIR_CALL_USER_FUNC_ARRAY(&out_zval, &func, &args);

  out = Z_TAITIME(out_zval);
  ZEPHIR_MM_RESTORE();
  return out;
}


void ext_Calends_calendar_register(zval *calendar_param, zval *default_format_param) {
  char *calendar, *default_format;

  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  default_format = ZSTR_VAL(Z_STR(*default_format_param));

  Calends_calendar_register(calendar, default_format,
    ext_Calends_calendar_to_internal_string,
    ext_Calends_calendar_to_internal_long_long,
    ext_Calends_calendar_to_internal_double,
    ext_Calends_calendar_to_internal_tai,
    ext_Calends_calendar_from_internal,
    ext_Calends_calendar_offset_string,
    ext_Calends_calendar_offset_long_long,
    ext_Calends_calendar_offset_double,
    ext_Calends_calendar_offset_tai
  );
}

void ext_Calends_calendar_unregister(zval *calendar_param) {
  char *calendar;

  calendar = ZSTR_VAL(Z_STR(*calendar_param));

  Calends_calendar_unregister(calendar);
}

zval ext_Calends_calendar_registered(zval *calendar_param) {
  char *calendar;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  calendar = ZSTR_VAL(Z_STR(*calendar_param));
  ZVAL_BOOL(&out, Calends_calendar_registered(calendar));

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_Calends_calendar_list_registered() {
  char *result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  // We'll let Zephir split the string. C's not great at it...
  result = Calends_calendar_list_registered();
  ZVAL_STRING(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}


zval ext_TAI64Time_add(zval *p_param, zval *z_param) {
  TAI64Time p, z, result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
  z = Z_TAITIME(*z_param);
	result = TAI64Time_add(p, z);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_sub(zval *p_param, zval *z_param) {
  TAI64Time p, z, result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
  z = Z_TAITIME(*z_param);
	result = TAI64Time_sub(p, z);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_string(zval *p_param) {
  TAI64Time p;
  char *result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
	result = TAI64Time_string(p);
  ZVAL_STRING(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_from_string(zval *in_param) {
  char *in;
  TAI64Time result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  in = ZSTR_VAL(Z_STR(*in_param));
	result = TAI64Time_from_string(in);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_hex_string(zval *p_param) {
  TAI64Time p;
  char *result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
	result = TAI64Time_hex_string(p);
  ZVAL_STRING(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_from_hex_string(zval *in_param) {
  char *in;
  TAI64Time result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  in = ZSTR_VAL(Z_STR(*in_param));
	result = TAI64Time_from_hex_string(in);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_double(zval *p_param) {
  TAI64Time p;
  double result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
	result = TAI64Time_double(p);
  ZVAL_DOUBLE(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_from_double(zval *in_param) {
  double in;
  TAI64Time result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  in = Z_DVAL(*in_param);
	result = TAI64Time_from_double(in);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_encode_text(zval *p_param) {
  TAI64Time p;
  char *result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
	result = TAI64Time_encode_text(p);
  ZVAL_STRING(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_decode_text(zval *encoded_param) {
  char *encoded;
  TAI64Time result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  encoded = ZSTR_VAL(Z_STR(*encoded_param));
	result = TAI64Time_from_string(encoded);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_encode_binary(zval *p_param) {
  //TODO...
}

zval ext_TAI64Time_decode_binary(zval *encoded_param, zval *len_param) {
  //TODO...
}

zval ext_TAI64Time_utc_to_tai(zval *p_param) {
  TAI64Time p, result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
	result = TAI64Time_utc_to_tai(p);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}

zval ext_TAI64Time_tai_to_utc(zval *p_param) {
  TAI64Time p, result;
  zval out;

  ZVAL_UNDEF(&out);

  ZEPHIR_MM_GROW();

  p = Z_TAITIME(*p_param);
	result = TAI64Time_tai_to_utc(p);
  ZVAL_TAITIME(&out, result);

  ZEPHIR_MM_RESTORE();

  return out;
}


void ZVAL_TAITIME(zval *out, TAI64Time tai_val) {
  zval _sec, _nano, _atto, _xicto, _ucto, _rocto;

  ZVAL_UNDEF(&_sec);
  ZVAL_UNDEF(&_nano);
  ZVAL_UNDEF(&_atto);
  ZVAL_UNDEF(&_xicto);
  ZVAL_UNDEF(&_ucto);
  ZVAL_UNDEF(&_rocto);
  ZVAL_UNDEF(out);

  ZEPHIR_MM_GROW();

  object_init_ex(out, calends_taitime_ce);
  ZVAL_DOUBLE(&_sec, tai_val.Seconds);
  zephir_update_property_zval(out, SL("seconds"), &_sec);
  ZVAL_LONG(&_nano, tai_val.Nano);
  zephir_update_property_zval(out, SL("nano"), &_nano);
  ZVAL_LONG(&_atto, tai_val.Atto);
  zephir_update_property_zval(out, SL("atto"), &_atto);
  ZVAL_LONG(&_xicto, tai_val.Xicto);
  zephir_update_property_zval(out, SL("xicto"), &_xicto);
  ZVAL_LONG(&_ucto, tai_val.Ucto);
  zephir_update_property_zval(out, SL("ucto"), &_ucto);
  ZVAL_LONG(&_rocto, tai_val.Rocto);
  zephir_update_property_zval(out, SL("rocto"), &_rocto);

  ZEPHIR_MM_RESTORE();
}

TAI64Time Z_TAITIME(zval in) {
  zval _sec, _nano, _atto, _xicto, _ucto, _rocto;

  ZVAL_UNDEF(&_sec);
  ZVAL_UNDEF(&_nano);
  ZVAL_UNDEF(&_atto);
  ZVAL_UNDEF(&_xicto);
  ZVAL_UNDEF(&_ucto);
  ZVAL_UNDEF(&_rocto);

  ZEPHIR_MM_GROW();

  TAI64Time out = {0, 0, 0, 0, 0, 0, 0};

  zephir_read_property(&_sec, &in, SL("seconds"), PH_NOISY_CC | PH_READONLY);
  out.Seconds = Z_DVAL(_sec);
  zephir_read_property(&_nano, &in, SL("nano"), PH_NOISY_CC | PH_READONLY);
  out.Nano = Z_LVAL(_nano);
  zephir_read_property(&_atto, &in, SL("atto"), PH_NOISY_CC | PH_READONLY);
  out.Atto = Z_LVAL(_atto);
  zephir_read_property(&_xicto, &in, SL("xicto"), PH_NOISY_CC | PH_READONLY);
  out.Xicto = Z_LVAL(_xicto);
  zephir_read_property(&_ucto, &in, SL("ucto"), PH_NOISY_CC | PH_READONLY);
  out.Ucto = Z_LVAL(_ucto);
  zephir_read_property(&_rocto, &in, SL("rocto"), PH_NOISY_CC | PH_READONLY);
  out.Rocto = Z_LVAL(_rocto);

  ZEPHIR_MM_RESTORE();

  return out;
}


void ext_panic_handler(char *error_message) {
  ZEPHIR_THROW_EXCEPTION_STRW(calends_calendsexception_ce, error_message);
}

void ext_unregister_all_calendars() {
  zval _wildcard;
	zend_long ZEPHIR_LAST_CALL_STATUS;
	zephir_fcall_cache_entry *_null = NULL;

	ZVAL_UNDEF(&_wildcard);

	ZEPHIR_MM_GROW();

	ZEPHIR_INIT_VAR(&_wildcard);
	ZVAL_STRING(&_wildcard, "*");
	ZEPHIR_CALL_CE_STATIC(NULL, calends_calends_ce, "calendarunregister", &_null, 0, &_wildcard);
	ZEPHIR_MM_RESTORE();
}
