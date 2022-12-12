--TEST--
new TestCalendar() Basic Interface test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	var_dump(Calends\CalendarDefinition::calendarRegistered('test'));

	class TestCalendar extends Calends\CalendarDefinition {
		function toInternal($date, string $format = ''): Calends\TAITime {
			return new Calends\TAITime();
		}

		function fromInternal(Calends\TAITime $stamp, string $format = ''): string {
			return "{$this->name}::{$stamp}::{$format}";
		}

		function offset($stamp, string $offset): Calends\TAITime {
			return $stamp;
		}
	}

	new TestCalendar('test', 'default');

	var_dump(Calends\CalendarDefinition::calendarRegistered('test'));
	var_dump(Calends\CalendarDefinition::listRegistered());
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
