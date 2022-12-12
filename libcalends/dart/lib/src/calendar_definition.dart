import 'dart:ffi';

import 'generated_bindings.dart';
import 'registered.dart';
import 'util.dart';

abstract class CalendarDefinition {
  String get name;

  String get defaultFormat;

  TAI64Time toInternal(dynamic stamp, String? format);

  String fromInternal(TAI64Time internal, String? format);

  TAI64Time offset(TAI64Time internal, dynamic offsetVal);

  static List<String> listRegistered() {
    return charPtrToStr(loadLib().calendar_list_registered()).split('\n');
  }

  void register() {
    registeredCalendars[name] = this;

    loadLib().calendar_register(
      strToCharPtr(name),
      strToCharPtr(defaultFormat),
      Pointer.fromFunction(toInternalString),
      Pointer.fromFunction(toInternalLongLong),
      Pointer.fromFunction(toInternalDouble),
      Pointer.fromFunction(toInternalTai),
      Pointer.fromFunction(fromInternalFunc),
      Pointer.fromFunction(offsetString),
      Pointer.fromFunction(offsetLongLong),
      Pointer.fromFunction(offsetDouble),
      Pointer.fromFunction(offsetTai),
    );
  }

  bool isRegistered() {
    return loadLib().calendar_registered(strToCharPtr(name)) != 0;
  }

  void unregister() {
    loadLib().calendar_unregister(strToCharPtr(name));
  }
}
