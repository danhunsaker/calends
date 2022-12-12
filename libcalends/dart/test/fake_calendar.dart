import 'package:calends/calends.dart';

class FakeCalendar extends CalendarDefinition {
  @override
  String get name => 'fake';

  @override
  String get defaultFormat => 'unknown';

  @override
  TAI64Time toInternal(dynamic stamp, String? format) {
    if (stamp is double) {
      return TAI64TimeMethods.fromDouble(stamp);
    } else if (stamp is int) {
      return TAI64TimeMethods.fromDouble(stamp * 1.0);
    } else if (stamp is String) {
      return TAI64TimeMethods.fromHex(stamp);
    } else if (stamp is TAI64Time) {
      return stamp;
    } else {
      throw CalendsException('Unsupported Type');
    }
  }

  @override
  String fromInternal(TAI64Time internal, String? format) {
    return internal.toHex();
  }

  @override
  TAI64Time offset(TAI64Time internal, dynamic offsetVal) {
    if (offsetVal is double) {
      return internal.add(TAI64TimeMethods.fromDouble(offsetVal));
    } else if (offsetVal is int) {
      return internal.add(TAI64TimeMethods.fromDouble(offsetVal * 1.0));
    } else if (offsetVal is String) {
      return internal.add(TAI64TimeMethods.fromHex(offsetVal));
    } else if (offsetVal is TAI64Time) {
      return internal.add(offsetVal);
    } else {
      throw CalendsException('Unsupported Type');
    }
  }
}
