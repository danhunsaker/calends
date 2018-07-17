namespace Calends;

class Calends implements \Serializable, \JsonSerializable
{
  private goId = 0.0;

  private static calendars = [];

  private function __construct()
  {
    return;
  }

  public function __destruct()
  {
    Calends_release(this->goId);
  }

  public function __toString() -> string
  {
    string out;

    let out = Calends_string(this->goId);

    return out;
  }

  public function serialize() -> string
  {
    string out;

    let out = Calends_encode_text(this->goId);

    return out;
  }

  public function unserialize(data) -> void
  {
    double result;

    let result = Calends_decode_text((string)data);

    let this->goId = result;
  }

  public function jsonSerialize() -> var
  {
    var out;

    let out = json_decode(Calends_encode_json(this->goId), true);

    return out;
  }

  public static function fromJson(string encoded) -> <Calends>
  {
    double result;
    var out;
    let out = new Calends;

    let result = Calends_decode_json(encoded);

    let out->goId = result;

    return out;
  }

  public static function create(var date = null, string calendar = null, string format = null) -> <Calends>
  {
    var out, start, duration, end;
    double result;
    let out = new Calends;

    if empty date {
      let date = "";
    }
    if empty calendar {
      let calendar = "";
    }
    if empty format {
      let format = "";
    }

    switch typeof date {
      case "string":
        let result = Calends_create_string((string)date, calendar, format);
        break;
      case "integer":
        let result = Calends_create_long_long((long)date, calendar, format);
        break;
      case "double":
        let result = Calends_create_double((double)date, calendar, format);
        break;
      case "array":
        if fetch start, date["start"] && fetch end, date["end"] {
          switch typeof start {
            case "string":
              let result = Calends_create_string_range((string)start, (string)end, calendar, format);
              break;
            case "integer":
              let result = Calends_create_long_long_range((long)start, (long)end, calendar, format);
              break;
            case "double":
              let result = Calends_create_double_range((double)start, (double)end, calendar, format);
              break;
            default:
              throw new CalendsException("Unsupported date type " . typeof start);
          }
        } elseif fetch start, date["start"] && fetch duration, date["duration"] {
          switch typeof start {
            case "string":
              let result = Calends_create_string_start_period((string)start, (string)duration, calendar, format);
              break;
            case "integer":
              let result = Calends_create_long_long_start_period((long)start, (long)duration, calendar, format);
              break;
            case "double":
              let result = Calends_create_double_start_period((double)start, (double)duration, calendar, format);
              break;
            default:
              throw new CalendsException("Unsupported date type " . typeof start);
          }
        } elseif fetch duration, date["duration"] && fetch end, date["end"] {
          switch typeof duration {
            case "string":
              let result = Calends_create_string_end_period((string)duration, (string)end, calendar, format);
              break;
            case "integer":
              let result = Calends_create_long_long_end_period((long)duration, (long)end, calendar, format);
              break;
            case "double":
              let result = Calends_create_double_end_period((double)duration, (double)end, calendar, format);
              break;
            default:
              throw new CalendsException("Unsupported date type " . typeof duration);
          }
        }
        break;
      default:
        throw new CalendsException("Unsupported date type " . typeof date);
    }

    let out->goId = result;
    return out;
  }

  public function date(string calendar = null, string format = null) -> string
  {
    string out;

    if empty calendar {
      let calendar = "";
    }
    if empty format {
      let format = "";
    }

    let out = Calends_date(this->goId, calendar, format);

    return out;
  }

  public function duration() -> string
  {
    string out;

    let out = Calends_duration(this->goId);

    return out;
  }

  public function endDate(string calendar = null, string format = null) -> string
  {
    string out;

    if empty calendar {
      let calendar = "";
    }
    if empty format {
      let format = "";
    }

    let out = Calends_end_date(this->goId, calendar, format);

    return out;
  }

  public function difference(<Calends> z, string mode = null) -> string
  {
    string out;

    if empty mode {
      let mode = "";
    }

    let out = Calends_difference(this->goId, z->goId, mode);

    return out;
  }

  public function compare(<Calends> z, string mode = null) -> int
  {
    var out;

    if empty mode {
      let mode = "";
    }

    let out = Calends_compare(this->goId, z->goId, mode);

    return (int)out;
  }

  public function isSame(<Calends> z) -> bool
  {
    var out;

    let out = Calends_is_same(this->goId, z->goId);

    return (bool)out;
  }

  public function isSameDuration(<Calends> z) -> bool
  {
    var out;

    let out = Calends_is_same_duration(this->goId, z->goId);

    return (bool)out;
  }

  public function isShorter(<Calends> z) -> bool
  {
    var out;

    let out = Calends_is_shorter(this->goId, z->goId);

    return (bool)out;
  }

  public function isLonger(<Calends> z) -> bool
  {
    var out;

    let out = Calends_is_longer(this->goId, z->goId);

    return (bool)out;
  }

  public function contains(<Calends> z) -> bool
  {
    var out;

    let out = Calends_contains(this->goId, z->goId);

    return (bool)out;
  }

  public function overlaps(<Calends> z) -> bool
  {
    var out;

    let out = Calends_overlaps(this->goId, z->goId);

    return (bool)out;
  }

  public function abuts(<Calends> z) -> bool
  {
    var out;

    let out = Calends_abuts(this->goId, z->goId);

    return (bool)out;
  }

  public function isBefore(<Calends> z) -> bool
  {
    var out;

    let out = Calends_is_before(this->goId, z->goId);

    return (bool)out;
  }

  public function startsBefore(<Calends> z) -> bool
  {
    var out;

    let out = Calends_starts_before(this->goId, z->goId);

    return (bool)out;
  }

  public function endsBefore(<Calends> z) -> bool
  {
    var out;

    let out = Calends_ends_before(this->goId, z->goId);

    return (bool)out;
  }

  public function isDuring(<Calends> z) -> bool
  {
    var out;

    let out = Calends_is_during(this->goId, z->goId);

    return (bool)out;
  }

  public function startsDuring(<Calends> z) -> bool
  {
    var out;

    let out = Calends_starts_during(this->goId, z->goId);

    return (bool)out;
  }

  public function endsDuring(<Calends> z) -> bool
  {
    var out;

    let out = Calends_ends_during(this->goId, z->goId);

    return (bool)out;
  }

  public function isAfter(<Calends> z) -> bool
  {
    var out;

    let out = Calends_is_after(this->goId, z->goId);

    return (bool)out;
  }

  public function startsAfter(<Calends> z) -> bool
  {
    var out;

    let out = Calends_starts_after(this->goId, z->goId);

    return (bool)out;
  }

  public function endsAfter(<Calends> z) -> bool
  {
    var out;

    let out = Calends_ends_after(this->goId, z->goId);

    return (bool)out;
  }

  public function add(var offset = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty offset {
      let offset = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof offset {
      case "string":
        let result = Calends_add_string(this->goId, (string)offset, calendar);
        break;
      case "integer":
        let result = Calends_add_long_long(this->goId, (long)offset, calendar);
        break;
      case "double":
        let result = Calends_add_double(this->goId, (double)offset, calendar);
        break;
      default:
        throw new CalendsException("Unsupported offset type " . typeof offset);
    }

    let out->goId = result;
    return out;
  }

  public function subtract(var offset = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty offset {
      let offset = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof offset {
      case "string":
      let result = Calends_subtract_string(this->goId, (string)offset, calendar);
      break;
      case "integer":
      let result = Calends_subtract_long_long(this->goId, (long)offset, calendar);
      break;
      case "double":
      let result = Calends_subtract_double(this->goId, (double)offset, calendar);
      break;
      default:
      throw new CalendsException("Unsupported offset type " . typeof offset);
    }

    let out->goId = result;
    return out;
  }

  public function addFromEnd(var offset = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty offset {
      let offset = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof offset {
      case "string":
        let result = Calends_add_from_end_string(this->goId, (string)offset, calendar);
        break;
      case "integer":
        let result = Calends_add_from_end_long_long(this->goId, (long)offset, calendar);
        break;
      case "double":
        let result = Calends_add_from_end_double(this->goId, (double)offset, calendar);
        break;
      default:
        throw new CalendsException("Unsupported offset type " . typeof offset);
    }

    let out->goId = result;
    return out;
  }

  public function subtractFromEnd(var offset = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty offset {
      let offset = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof offset {
      case "string":
      let result = Calends_subtract_from_end_string(this->goId, (string)offset, calendar);
      break;
      case "integer":
      let result = Calends_subtract_from_end_long_long(this->goId, (long)offset, calendar);
      break;
      case "double":
      let result = Calends_subtract_from_end_double(this->goId, (double)offset, calendar);
      break;
      default:
      throw new CalendsException("Unsupported offset type " . typeof offset);
    }

    let out->goId = result;
    return out;
  }

  public function next(var offset = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty offset {
      let offset = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof offset {
      case "string":
        let result = Calends_next_string(this->goId, (string)offset, calendar);
        break;
      case "integer":
        let result = Calends_next_long_long(this->goId, (long)offset, calendar);
        break;
      case "double":
        let result = Calends_next_double(this->goId, (double)offset, calendar);
        break;
      default:
        throw new CalendsException("Unsupported offset type " . typeof offset);
    }

    let out->goId = result;
    return out;
  }

  public function previous(var offset = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty offset {
      let offset = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof offset {
      case "string":
      let result = Calends_previous_string(this->goId, (string)offset, calendar);
      break;
      case "integer":
      let result = Calends_previous_long_long(this->goId, (long)offset, calendar);
      break;
      case "double":
      let result = Calends_previous_double(this->goId, (double)offset, calendar);
      break;
      default:
      throw new CalendsException("Unsupported offset type " . typeof offset);
    }

    let out->goId = result;
    return out;
  }

  public function withDate(var date = null, string calendar = null, string format = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty date {
      let date = "";
    }
    if empty calendar {
      let calendar = "";
    }
    if empty format {
      let format = "";
    }

    switch typeof date {
      case "string":
        let result = Calends_with_date_string(this->goId, (string)date, calendar, format);
        break;
      case "integer":
        let result = Calends_with_date_long_long(this->goId, (long)date, calendar, format);
        break;
      case "double":
        let result = Calends_with_date_double(this->goId, (double)date, calendar, format);
        break;
      default:
        throw new CalendsException("Unsupported date type " . typeof date);
    }

    let out->goId = result;
    return out;
  }

  public function withDuration(var duration = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty duration {
      let duration = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof duration {
      case "string":
      let result = Calends_with_duration_string(this->goId, (string)duration, calendar);
      break;
      case "integer":
      let result = Calends_with_duration_long_long(this->goId, (long)duration, calendar);
      break;
      case "double":
      let result = Calends_with_duration_double(this->goId, (double)duration, calendar);
      break;
      default:
      throw new CalendsException("Unsupported duration type " . typeof duration);
    }

    let out->goId = result;
    return out;
  }

  public function withDurationFromEnd(var duration = null, string calendar = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty duration {
      let duration = "";
    }
    if empty calendar {
      let calendar = "";
    }

    switch typeof duration {
      case "string":
      let result = Calends_with_duration_from_end_string(this->goId, (string)duration, calendar);
      break;
      case "integer":
      let result = Calends_with_duration_from_end_long_long(this->goId, (long)duration, calendar);
      break;
      case "double":
      let result = Calends_with_duration_from_end_double(this->goId, (double)duration, calendar);
      break;
      default:
      throw new CalendsException("Unsupported duration type " . typeof duration);
    }

    let out->goId = result;
    return out;
  }

  public function withEndDate(var date = null, string calendar = null, string format = null) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    if empty date {
      let date = "";
    }
    if empty calendar {
      let calendar = "";
    }
    if empty format {
      let format = "";
    }

    switch typeof date {
      case "string":
        let result = Calends_with_end_date_string(this->goId, (string)date, calendar, format);
        break;
      case "integer":
        let result = Calends_with_end_date_long_long(this->goId, (long)date, calendar, format);
        break;
      case "double":
        let result = Calends_with_end_date_double(this->goId, (double)date, calendar, format);
        break;
      default:
        throw new CalendsException("Unsupported date type " . typeof date);
    }

    let out->goId = result;
    return out;
  }

  public function merge(<Calends> z) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    let result = Calends_merge(this->goId, z->goId);

    let out->goId = result;
    return out;
  }

  public function intersect(<Calends> z) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    let result = Calends_intersect(this->goId, z->goId);

    let out->goId = result;
    return out;
  }

  public function gap(<Calends> z) -> <Calends>
  {
    var out;
    double result;
    let out = new Calends;

    let result = Calends_gap(this->goId, z->goId);

    let out->goId = result;
    return out;
  }

  public static function calendarRegistered(string name) -> bool
  {
    var out;

    let out = Calends_calendar_registered(name);

    return (bool)out;
  }

  public static function calendarListRegistered() -> array
  {
    var result;

    let result = explode("\n", Calends_calendar_list_registered());

    return (array)result;
  }

  public static function calendarRegister(string name, string defaultFormat, var calendar) -> void
  {
    if !((is_string(calendar) && class_exists(calendar) && is_subclass_of(calendar, "Calends\\CalendarInterface"))
        || calendar instanceof CalendarInterface || calendar instanceof CalendarObjectInterface) {
      if is_string(calendar) {
        throw new CalendsException("Parameter 'calendar' must name or be an implementation of Calends\\CalendarInterface or Calends\\CalendarObjectInterface; got " . calendar . " instead");
      } else {
        throw new CalendsException("Parameter 'calendar' must name or be an implementation of Calends\\CalendarInterface or Calends\\CalendarObjectInterface; got a(n) " . get_class(calendar) . " instead");
      }
    }

    if is_string(calendar) && class_exists(calendar) && is_subclass_of(calendar, "Calends\\CalendarInterface") {
      let self::calendars[name] = [
        "toInternal":   calendar . "::" . "toInternal",
        "fromInternal": calendar . "::" . "fromInternal",
        "offset":       calendar . "::" . "offset"
      ];
    } elseif calendar instanceof CalendarInterface {
      let self::calendars[name] = [
        "toInternal":   typeof calendar . "::" . "toInternal",
        "fromInternal": typeof calendar . "::" . "fromInternal",
        "offset":       typeof calendar . "::" . "offset"
      ];
    } else {
      let self::calendars[name] = [
        "toInternal":   [calendar, "toInternal"],
        "fromInternal": [calendar, "fromInternal"],
        "offset":       [calendar, "offset"]
      ];
    }

    Calends_calendar_register(name, defaultFormat);

    return;
  }

  public static function calendarUnregister(string name) -> void
  {
    var calendar;

    if name == "*" {
      for calendar, _ in self::calendars {
        self::calendarUnregister(calendar);
      }
    } elseif fetch calendar, self::calendars[name] {
      Calends_calendar_unregister(name);

      unset(self::calendars[name]);
    } else {
      // This calendar was never registered, so there's nothing to do, here.
    }

    return;
  }
}
