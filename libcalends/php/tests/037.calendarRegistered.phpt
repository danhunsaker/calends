--TEST--
Calends\CalendarDefinition::calendarRegistered() Basic test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	$ret = Calends\CalendarDefinition::calendarRegistered('invalid');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::calendarRegistered('tai64');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::calendarRegistered('unix');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::calendarRegistered('test');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::calendarRegistered('');
	var_dump($ret);
?>
--EXPECT--
bool(false)
bool(true)
bool(true)
bool(false)
bool(false)
