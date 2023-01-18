--TEST--
Calends\CalendarDefinition::calendarUnregister() Basic test
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

	$test = new TestCalendar('test', 'default');

	var_dump(Calends\CalendarDefinition::calendarRegistered('test'));

	$test->unregister();

	var_dump(Calends\CalendarDefinition::calendarRegistered('test'));
?>
--EXPECT--
bool(false)
bool(true)
bool(false)
