#include "calendars.h"

TAI64Time* wrap_Calends_calendar_to_internal_string(Calends_calendar_to_internal_string f, char* date, char* format) {
        return f(date, format);
}

TAI64Time* wrap_Calends_calendar_to_internal_long_long(Calends_calendar_to_internal_long_long f, long long int date, char* format) {
        return f(date, format);
}

TAI64Time* wrap_Calends_calendar_to_internal_double(Calends_calendar_to_internal_double f, double date, char* format) {
        return f(date, format);
}

TAI64Time* wrap_Calends_calendar_to_internal_tai(Calends_calendar_to_internal_tai f, TAI64Time* date) {
        return f(date);
}

char* wrap_Calends_calendar_from_internal(Calends_calendar_from_internal f, TAI64Time* stamp, char* format) {
        return f(stamp, format);
}

TAI64Time* wrap_Calends_calendar_offset_string(Calends_calendar_offset_string f, TAI64Time* stamp, char* offset) {
        return f(stamp, offset);
}

TAI64Time* wrap_Calends_calendar_offset_long_long(Calends_calendar_offset_long_long f, TAI64Time* stamp, long long int offset) {
        return f(stamp, offset);
}

TAI64Time* wrap_Calends_calendar_offset_double(Calends_calendar_offset_double f, TAI64Time* stamp, double offset) {
        return f(stamp, offset);
}

TAI64Time* wrap_Calends_calendar_offset_tai(Calends_calendar_offset_tai f, TAI64Time* stamp, TAI64Time* offset) {
        return f(stamp, offset);
}
