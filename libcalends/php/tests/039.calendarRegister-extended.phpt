--TEST--
new TestCalendar() Extended Interface test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	if (Calends\CalendarDefinition::isRegistered('test')) {
		echo "Calendar already registered!";
		exit;
	}

	class TestCalendar extends Calends\CalendarDefinition {
		public function toInternal($date, string $format = ''): Calends\TAITime {
			return Calends\TAITime::fromNumber(0);
		}

		public function fromInternal(Calends\TAITime $stamp, string $format = ''): string {
			return "{$this->name}::{$stamp}::{$format}";
		}

		public function offset($stamp, string $offset): Calends\TAITime {
			return $stamp;
		}
	}

	new TestCalendar('test', 'default');

	if (!Calends\CalendarDefinition::isRegistered('test')) {
		echo "Calendar not registered!";
		exit;
	}

	echo "testing toInternal():\n";
	$ret = Calends\Calends::create('', 'test');
	var_dump($ret);

	echo "\ntesting fromInternal():\n";
	var_dump($ret->date('test'));

	echo "\ntesting offset():\n";
	var_dump($ret->add('', 'test'));
	var_dump($ret->add('', 'test')->date('test'));
?>
--EXPECTF--
testing toInternal():
object(Calends\Calends)#%d (%d) {
  ["start"]=>
  string(%d) "%x"
  ["end"]=>
  string(%d) "%x"
}

testing fromInternal():
string(16) "test::0::default"

testing offset():
object(Calends\Calends)#%d (%d) {
  ["start"]=>
  string(%d) "%x"
  ["end"]=>
  string(%d) "%x"
}
string(16) "test::0::default"
