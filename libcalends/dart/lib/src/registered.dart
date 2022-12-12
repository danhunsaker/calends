import 'dart:ffi';

import 'calendar_definition.dart';
import 'generated_bindings.dart';
import 'util.dart';

Map<String, CalendarDefinition> registeredCalendars = {};

TAI64Time toInternalString(
  Pointer<Char> calendar,
  Pointer<Char> stamp,
  Pointer<Char> format,
) {
  return registeredCalendars[charPtrToStr(calendar)]!
      .toInternal(charPtrToStr(stamp), charPtrToStr(format));
}

TAI64Time toInternalLongLong(
  Pointer<Char> calendar,
  int stamp,
  Pointer<Char> format,
) {
  return registeredCalendars[charPtrToStr(calendar)]!
      .toInternal(stamp, charPtrToStr(format));
}

TAI64Time toInternalDouble(
  Pointer<Char> calendar,
  double stamp,
  Pointer<Char> format,
) {
  return registeredCalendars[charPtrToStr(calendar)]!
      .toInternal(stamp, charPtrToStr(format));
}

TAI64Time toInternalTai(
  Pointer<Char> calendar,
  TAI64Time stamp,
) {
  return registeredCalendars[charPtrToStr(calendar)]!.toInternal(stamp, null);
}

Pointer<Char> fromInternalFunc(
  Pointer<Char> calendar,
  TAI64Time internal,
  Pointer<Char> format,
) {
  return strToCharPtr(registeredCalendars[charPtrToStr(calendar)]!
      .fromInternal(internal, charPtrToStr(format)));
}

TAI64Time offsetString(
  Pointer<Char> calendar,
  TAI64Time internal,
  Pointer<Char> offsetStr,
) {
  return registeredCalendars[charPtrToStr(calendar)]!
      .offset(internal, charPtrToStr(offsetStr));
}

TAI64Time offsetLongLong(
  Pointer<Char> calendar,
  TAI64Time internal,
  int offsetInt,
) {
  return registeredCalendars[charPtrToStr(calendar)]!
      .offset(internal, offsetInt);
}

TAI64Time offsetDouble(
  Pointer<Char> calendar,
  TAI64Time internal,
  double offsetDbl,
) {
  return registeredCalendars[charPtrToStr(calendar)]!
      .offset(internal, offsetDbl);
}

TAI64Time offsetTai(
  Pointer<Char> calendar,
  TAI64Time internal,
  TAI64Time offsetTai,
) {
  return registeredCalendars[charPtrToStr(calendar)]!
      .offset(internal, offsetTai);
}
