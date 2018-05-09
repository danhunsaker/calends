namespace Calends;

interface CalendarInterface
{
  public static function toInternal(var date, string format) -> <TAITime>;

  public static function fromInternal(<TAITime> stamp, string format) -> string;

  public static function offset(<TAITime> stamp, var offset) -> <TAITime>;
}
