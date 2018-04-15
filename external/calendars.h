typedef struct {
        long long int Seconds; // Seconds since 1970-01-01 00:00:00 TAI
        unsigned int Nano;     // Nanoseconds since the given second
        unsigned int Atto;     // Attoseconds since the given nanosecond
        unsigned int Xicto;    // Xictoseconds since the given attosecond
        unsigned int Ucto;     // Uctoseconds since the given xictosecond
        unsigned int Rocto;    // Roctoseconds since the given uctosecond
} TAI64Time;

typedef TAI64Time* (*Calends_calendar_to_internal_string) (char*, char*);
typedef TAI64Time* (*Calends_calendar_to_internal_long_long) (long long int, char*);
typedef TAI64Time* (*Calends_calendar_to_internal_double) (double, char*);
typedef TAI64Time* (*Calends_calendar_to_internal_tai) (TAI64Time*);

typedef char* (*Calends_calendar_from_internal) (TAI64Time*, char*);

typedef TAI64Time* (*Calends_calendar_offset_string) (TAI64Time*, char*);
typedef TAI64Time* (*Calends_calendar_offset_long_long) (TAI64Time*, long long int);
typedef TAI64Time* (*Calends_calendar_offset_double) (TAI64Time*, double);
typedef TAI64Time* (*Calends_calendar_offset_tai) (TAI64Time*, TAI64Time*);

TAI64Time* wrap_Calends_calendar_to_internal_string(Calends_calendar_to_internal_string f, char* date, char* format);
TAI64Time* wrap_Calends_calendar_to_internal_long_long(Calends_calendar_to_internal_long_long f, long long int date, char* format);
TAI64Time* wrap_Calends_calendar_to_internal_double(Calends_calendar_to_internal_double f, double date, char* format);
TAI64Time* wrap_Calends_calendar_to_internal_tai(Calends_calendar_to_internal_tai f, TAI64Time* date);

char* wrap_Calends_calendar_from_internal(Calends_calendar_from_internal f, TAI64Time* stamp, char* format);

TAI64Time* wrap_Calends_calendar_offset_string(Calends_calendar_offset_string f, TAI64Time* stamp, char* offset);
TAI64Time* wrap_Calends_calendar_offset_long_long(Calends_calendar_offset_long_long f, TAI64Time* stamp, long long int offset);
TAI64Time* wrap_Calends_calendar_offset_double(Calends_calendar_offset_double f, TAI64Time* stamp, double offset);
TAI64Time* wrap_Calends_calendar_offset_tai(Calends_calendar_offset_tai f, TAI64Time* stamp, TAI64Time* offset);
