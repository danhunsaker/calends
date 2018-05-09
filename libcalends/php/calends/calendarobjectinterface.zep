namespace Calends;

interface CalendarObjectInterface
{
  public function toInternal(var date, string format) -> <TAITime>;

  public function fromInternal(<TAITime> stamp, string format) -> string;

  public function offset(<TAITime> stamp, var offset) -> <TAITime>;
}
