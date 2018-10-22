--TEST--
Calends\Calends::calendarRegister() Basic Interface test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	var_dump(Calends\Calends::calendarRegistered('test'));

	class TestCalendar implements Calends\CalendarInterface {
		static function toInternal($date, string $format): Calends\TAITime {
			return new Calends\TAITime();
		}

		static function fromInternal(Calends\TAITime $stamp, string $format): string {
			return "{$stamp->toString()}::{$format}";
		}

		static function offset(Calends\TAITime $stamp, $offset): Calends\TAITime {
			return $stamp;
		}
	}

	Calends\Calends::calendarRegister('test', 'default', TestCalendar::class);

	var_dump(Calends\Calends::calendarRegistered('test'));
	var_dump(Calends\Calends::calendarListRegistered());
?>
--EXPECT--
bool(false)
bool(true)
array(6) {
  [0]=>
  string(9) "Gregorian"
  [1]=>
  string(3) "Jdc"
  [2]=>
  string(8) "Stardate"
  [3]=>
  string(5) "Tai64"
  [4]=>
  string(4) "Test"
  [5]=>
  string(4) "Unix"
}
