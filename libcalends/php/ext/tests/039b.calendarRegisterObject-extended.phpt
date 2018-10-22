--TEST--
Calends\Calends::calendarRegister() Extended Object test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	if (Calends\Calends::calendarRegistered('test')) {
		echo "Calendar already registered!";
		exit;
	}

	class TestCalendar implements Calends\CalendarObjectInterface {
		function toInternal($date, string $format): Calends\TAITime {
			return new Calends\TAITime();
		}

		function fromInternal(Calends\TAITime $stamp, string $format): string {
			return "{$stamp->toString()}::{$format}";
		}

		function offset(Calends\TAITime $stamp, $offset): Calends\TAITime {
			return $stamp;
		}
	}

	Calends\Calends::calendarRegister('test', 'default', new TestCalendar);

	if (!Calends\Calends::calendarRegistered('test')) {
		echo "Calendar not registered!";
		exit;
	}

	echo "testing toInternal():\n";
	$ret = Calends\Calends::create(null, 'test');
	var_dump($ret);

	echo "\ntesting fromInternal():\n";
	var_dump($ret->date('test'));

	echo "\ntesting offset():\n";
	var_dump($ret->add(null, 'test'));
	var_dump($ret->add(null, 'test')->date('test'));
?>
--EXPECTF--
testing toInternal():
object(Calends\Calends)#%d (%d) {
  ["goId":"Calends\Calends":private]=>
  float(%d)
}

testing fromInternal():
string(10) "0::default"

testing offset():
object(Calends\Calends)#%d (%d) {
  ["goId":"Calends\Calends":private]=>
  float(%d)
}
string(10) "0::default"
