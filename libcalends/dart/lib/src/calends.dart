import 'dart:ffi';

import 'package:calends/src/calends_exception.dart';

import 'generated_bindings.dart';
import 'util.dart';

class Calends implements Finalizable {
  static void initialize() {
    loadLib().register_panic_handler(Pointer.fromFunction(panicHandler));
  }

  Calends(dynamic stamp, String calendarString, String formatString) {
    final calendar = strToCharPtr(calendarString);
    final format = strToCharPtr(formatString);

    _bound = loadLib();

    if (stamp is String) {
      _reference = _bound.create_string(
        strToCharPtr(stamp),
        calendar,
        format,
      );
    } else if (stamp is int) {
      _reference = _bound.create_long_long(stamp, calendar, format);
    } else if (stamp is double) {
      _reference = _bound.create_double(stamp, calendar, format);
    } else if (stamp is Map<String, String>) {
      if (stamp.containsKey('start') && stamp.containsKey('end')) {
        _reference = _bound.create_string_range(
          strToCharPtr(stamp['start'] as String),
          strToCharPtr(stamp['end'] as String),
          calendar,
          format,
        );
      } else if (stamp.containsKey('start') && stamp.containsKey('duration')) {
        _reference = _bound.create_string_start_period(
          strToCharPtr(stamp['start'] as String),
          strToCharPtr(stamp['duration'] as String),
          calendar,
          format,
        );
      } else if (stamp.containsKey('duration') && stamp.containsKey('end')) {
        _reference = _bound.create_string_end_period(
          strToCharPtr(stamp['duration'] as String),
          strToCharPtr(stamp['end'] as String),
          calendar,
          format,
        );
      }
    } else if (stamp is Map<String, int>) {
      if (stamp.containsKey('start') && stamp.containsKey('end')) {
        _reference = _bound.create_long_long_range(
          stamp['start'] as int,
          stamp['end'] as int,
          calendar,
          format,
        );
      } else if (stamp.containsKey('start') && stamp.containsKey('duration')) {
        _reference = _bound.create_long_long_start_period(
          stamp['start'] as int,
          stamp['duration'] as int,
          calendar,
          format,
        );
      } else if (stamp.containsKey('duration') && stamp.containsKey('end')) {
        _reference = _bound.create_long_long_end_period(
          stamp['duration'] as int,
          stamp['end'] as int,
          calendar,
          format,
        );
      }
    } else if (stamp is Map<String, double>) {
      if (stamp.containsKey('start') && stamp.containsKey('end')) {
        _reference = _bound.create_double_range(
          stamp['start'] as double,
          stamp['end'] as double,
          calendar,
          format,
        );
      } else if (stamp.containsKey('start') && stamp.containsKey('duration')) {
        _reference = _bound.create_double_start_period(
          stamp['start'] as double,
          stamp['duration'] as double,
          calendar,
          format,
        );
      } else if (stamp.containsKey('duration') && stamp.containsKey('end')) {
        _reference = _bound.create_double_end_period(
          stamp['duration'] as double,
          stamp['end'] as double,
          calendar,
          format,
        );
      }
    } else {
      throw CalendsException('Unusable timestamp type');
    }

    _finalizer.attach(this, _reference);
  }

  Calends.decodeText(String textString) {
    _bound = loadLib();
    _reference = _bound.decode_text(strToCharPtr(textString));
    _finalizer.attach(this, _reference);
  }

  Calends.decodeJson(String jsonString) {
    _bound = loadLib();
    _reference = _bound.decode_json(strToCharPtr(jsonString));
    _finalizer.attach(this, _reference);
  }

  Calends._(this._bound, this._reference) {
    _finalizer.attach(this, _reference);
  }

  static final Finalizer<int> _finalizer =
      Finalizer((reference) => loadLib().release(reference));

  late CalendsBindings _bound;

  late int _reference;

  String date(String calendarString, String formatString) {
    final result = _bound.date(
      _reference,
      strToCharPtr(calendarString),
      strToCharPtr(formatString),
    );

    return charPtrToStr(result);
  }

  String duration() {
    final result = _bound.duration(_reference);

    return charPtrToStr(result);
  }

  String endDate(String calendarString, String formatString) {
    final result = _bound.end_date(
      _reference,
      strToCharPtr(calendarString),
      strToCharPtr(formatString),
    );

    return charPtrToStr(result);
  }

  Calends add(dynamic offset, String calendar) {
    Calends retVal;

    switch (offset.runtimeType) {
      case String:
        retVal = Calends._(
          _bound,
          _bound.add_string(
            _reference,
            strToCharPtr(offset as String),
            strToCharPtr(calendar),
          ),
        );
        break;
      case int:
        retVal = Calends._(
          _bound,
          _bound.add_long_long(
            _reference,
            offset as int,
            strToCharPtr(calendar),
          ),
        );
        break;
      case double:
        retVal = Calends._(
          _bound,
          _bound.add_double(
            _reference,
            offset as double,
            strToCharPtr(calendar),
          ),
        );
        break;
      default:
        throw FormatException('offset must be String, int, or double');
    }

    return retVal;
  }

  Calends subtract(dynamic offset, String calendar) {
    Calends retVal;

    switch (offset.runtimeType) {
      case String:
        retVal = Calends._(
          _bound,
          _bound.subtract_string(
            _reference,
            strToCharPtr(offset as String),
            strToCharPtr(calendar),
          ),
        );
        break;
      case int:
        retVal = Calends._(
          _bound,
          _bound.subtract_long_long(
            _reference,
            offset as int,
            strToCharPtr(calendar),
          ),
        );
        break;
      case double:
        retVal = Calends._(
          _bound,
          _bound.subtract_double(
            _reference,
            offset as double,
            strToCharPtr(calendar),
          ),
        );
        break;
      default:
        throw FormatException('offset must be String, int, or double');
    }

    return retVal;
  }

  Calends addFromEnd(dynamic offset, String calendar) {
    Calends retVal;

    switch (offset.runtimeType) {
      case String:
        retVal = Calends._(
          _bound,
          _bound.add_from_end_string(
            _reference,
            strToCharPtr(offset as String),
            strToCharPtr(calendar),
          ),
        );
        break;
      case int:
        retVal = Calends._(
          _bound,
          _bound.add_from_end_long_long(
            _reference,
            offset as int,
            strToCharPtr(calendar),
          ),
        );
        break;
      case double:
        retVal = Calends._(
          _bound,
          _bound.add_from_end_double(
            _reference,
            offset as double,
            strToCharPtr(calendar),
          ),
        );
        break;
      default:
        throw FormatException('offset must be String, int, or double');
    }

    return retVal;
  }

  Calends subtractFromEnd(dynamic offset, String calendar) {
    Calends retVal;

    switch (offset.runtimeType) {
      case String:
        retVal = Calends._(
          _bound,
          _bound.subtract_from_end_string(
            _reference,
            strToCharPtr(offset as String),
            strToCharPtr(calendar),
          ),
        );
        break;
      case int:
        retVal = Calends._(
          _bound,
          _bound.subtract_from_end_long_long(
            _reference,
            offset as int,
            strToCharPtr(calendar),
          ),
        );
        break;
      case double:
        retVal = Calends._(
          _bound,
          _bound.subtract_from_end_double(
            _reference,
            offset as double,
            strToCharPtr(calendar),
          ),
        );
        break;
      default:
        throw FormatException('offset must be String, int, or double');
    }

    return retVal;
  }

  Calends next(dynamic offset, String calendar) {
    Calends retVal;

    switch (offset.runtimeType) {
      case String:
        retVal = Calends._(
          _bound,
          _bound.next_string(
            _reference,
            strToCharPtr(offset as String),
            strToCharPtr(calendar),
          ),
        );
        break;
      case int:
        retVal = Calends._(
          _bound,
          _bound.next_long_long(
            _reference,
            offset as int,
            strToCharPtr(calendar),
          ),
        );
        break;
      case double:
        retVal = Calends._(
          _bound,
          _bound.next_double(
            _reference,
            offset as double,
            strToCharPtr(calendar),
          ),
        );
        break;
      default:
        throw FormatException('offset must be String, int, or double');
    }

    return retVal;
  }

  Calends previous(dynamic offset, String calendar) {
    Calends retVal;

    switch (offset.runtimeType) {
      case String:
        retVal = Calends._(
          _bound,
          _bound.previous_string(
            _reference,
            strToCharPtr(offset as String),
            strToCharPtr(calendar),
          ),
        );
        break;
      case int:
        retVal = Calends._(
          _bound,
          _bound.previous_long_long(
            _reference,
            offset as int,
            strToCharPtr(calendar),
          ),
        );
        break;
      case double:
        retVal = Calends._(
          _bound,
          _bound.previous_double(
            _reference,
            offset as double,
            strToCharPtr(calendar),
          ),
        );
        break;
      default:
        throw FormatException('offset must be String, int, or double');
    }

    return retVal;
  }

  Calends withDate(dynamic stamp, String calendar, String format) {
    int retVal;

    switch (stamp.runtimeType) {
      case String:
        retVal = _bound.with_date_string(
          _reference,
          strToCharPtr(stamp as String),
          strToCharPtr(calendar),
          strToCharPtr(format),
        );
        break;
      case int:
        retVal = _bound.with_date_long_long(
          _reference,
          stamp as int,
          strToCharPtr(calendar),
          strToCharPtr(format),
        );
        break;
      case double:
        retVal = _bound.with_date_double(
          _reference,
          stamp as double,
          strToCharPtr(calendar),
          strToCharPtr(format),
        );
        break;
      default:
        throw FormatException('stamp must be String, int, or double');
    }

    return Calends._(_bound, retVal);
  }

  Calends withEndDate(dynamic stamp, String calendar, String format) {
    int retVal;

    switch (stamp.runtimeType) {
      case String:
        retVal = _bound.with_end_date_string(
          _reference,
          strToCharPtr(stamp as String),
          strToCharPtr(calendar),
          strToCharPtr(format),
        );
        break;
      case int:
        retVal = _bound.with_end_date_long_long(
          _reference,
          stamp as int,
          strToCharPtr(calendar),
          strToCharPtr(format),
        );
        break;
      case double:
        retVal = _bound.with_end_date_double(
          _reference,
          stamp as double,
          strToCharPtr(calendar),
          strToCharPtr(format),
        );
        break;
      default:
        throw FormatException('stamp must be String, int, or double');
    }

    return Calends._(_bound, retVal);
  }

  Calends withDuration(dynamic duration, String calendar) {
    int retVal;

    switch (duration.runtimeType) {
      case String:
        retVal = _bound.with_duration_string(
          _reference,
          strToCharPtr(duration as String),
          strToCharPtr(calendar),
        );
        break;
      case int:
        retVal = _bound.with_duration_long_long(
          _reference,
          duration as int,
          strToCharPtr(calendar),
        );
        break;
      case double:
        retVal = _bound.with_duration_double(
          _reference,
          duration as double,
          strToCharPtr(calendar),
        );
        break;
      default:
        throw FormatException('stamp must be String, int, or double');
    }

    return Calends._(_bound, retVal);
  }

  Calends withDurationFromEnd(dynamic duration, String calendar) {
    int retVal;

    switch (duration.runtimeType) {
      case String:
        retVal = _bound.with_duration_from_end_string(
          _reference,
          strToCharPtr(duration as String),
          strToCharPtr(calendar),
        );
        break;
      case int:
        retVal = _bound.with_duration_from_end_long_long(
          _reference,
          duration as int,
          strToCharPtr(calendar),
        );
        break;
      case double:
        retVal = _bound.with_duration_from_end_double(
          _reference,
          duration as double,
          strToCharPtr(calendar),
        );
        break;
      default:
        throw FormatException('stamp must be String, int, or double');
    }

    return Calends._(_bound, retVal);
  }

  Calends merge(Calends other) {
    return Calends._(
      _bound,
      _bound.merge(_reference, other._reference),
    );
  }

  Calends intersect(Calends other) {
    return Calends._(
      _bound,
      _bound.intersect(_reference, other._reference),
    );
  }

  Calends gap(Calends other) {
    return Calends._(
      _bound,
      _bound.gap(_reference, other._reference),
    );
  }

  String difference(Calends other, String modeString) {
    return charPtrToStr(_bound.difference(
      _reference,
      other._reference,
      strToCharPtr(modeString),
    ));
  }

  int compare(Calends other, String modeString) {
    return _bound.compare(
      _reference,
      other._reference,
      strToCharPtr(modeString),
    );
  }

  bool isSame(Calends other) {
    return _bound.is_same(_reference, other._reference) != 0;
  }

  bool contains(Calends other) {
    return _bound.contains(_reference, other._reference) != 0;
  }

  bool overlaps(Calends other) {
    return _bound.overlaps(_reference, other._reference) != 0;
  }

  bool abuts(Calends other) {
    return _bound.abuts(_reference, other._reference) != 0;
  }

  bool isBefore(Calends other) {
    return _bound.is_before(_reference, other._reference) != 0;
  }

  bool startsBefore(Calends other) {
    return _bound.starts_before(_reference, other._reference) != 0;
  }

  bool endsBefore(Calends other) {
    return _bound.ends_before(_reference, other._reference) != 0;
  }

  bool isDuring(Calends other) {
    return _bound.is_during(_reference, other._reference) != 0;
  }

  bool startsDuring(Calends other) {
    return _bound.starts_during(_reference, other._reference) != 0;
  }

  bool endsDuring(Calends other) {
    return _bound.ends_during(_reference, other._reference) != 0;
  }

  bool isAfter(Calends other) {
    return _bound.is_after(_reference, other._reference) != 0;
  }

  bool startsAfter(Calends other) {
    return _bound.starts_after(_reference, other._reference) != 0;
  }

  bool endsAfter(Calends other) {
    return _bound.ends_after(_reference, other._reference) != 0;
  }

  bool isShorter(Calends other) {
    return _bound.is_shorter(_reference, other._reference) != 0;
  }

  bool isSameDuration(Calends other) {
    return _bound.is_same_duration(_reference, other._reference) != 0;
  }

  bool isLonger(Calends other) {
    return _bound.is_longer(_reference, other._reference) != 0;
  }

  @override
  String toString() {
    final result = _bound.string(_reference);

    return charPtrToStr(result);
  }

  String encodeText() {
    final result = _bound.encode_text(_reference);

    return charPtrToStr(result);
  }

  String encodeJson() {
    final result = _bound.encode_json(_reference);

    return charPtrToStr(result);
  }

  @override
  int get hashCode => _reference.hashCode;

  @override
  bool operator ==(Object other) {
    return other is Calends && isSame(other);
  }

  bool operator <(Calends other) {
    return isBefore(other);
  }

  bool operator <=(Calends other) {
    return isBefore(other) || isSame(other);
  }

  bool operator >=(Calends other) {
    return isAfter(other) || isSame(other);
  }

  bool operator >(Calends other) {
    return isAfter(other);
  }
}
