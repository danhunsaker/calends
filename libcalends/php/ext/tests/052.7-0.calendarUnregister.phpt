--TEST--
Calends\Calends::calendarUnregister() Basic test (PHP 7.0-7.1)
--SKIPIF--
<?php
if (version_compare(PHP_VERSION, '7.2', '>=') || !extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	var_dump(Calends\Calends::calendarRegistered('test'));

	class TestCalendar implements Calends\CalendarObjectInterface {
		function toInternal($date, $format): Calends\TAITime {
			return new Calends\TAITime();
		}

		function fromInternal(Calends\TAITime $stamp, $format): string {
			return "{$stamp->toString()}::{$format}";
		}

		function offset(Calends\TAITime $stamp, $offset): Calends\TAITime {
			return $stamp;
		}
	}

	Calends\Calends::calendarRegister('test', 'default', new TestCalendar);

	var_dump(Calends\Calends::calendarRegistered('test'));

	Calends\Calends::calendarUnregister('test');

	var_dump(Calends\Calends::calendarRegistered('test'));
?>
--EXPECT--
bool(false)
bool(true)
bool(false)
