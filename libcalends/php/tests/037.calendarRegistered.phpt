--TEST--
Calends\CalendarDefinition::isRegistered() Basic test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	$ret = Calends\CalendarDefinition::isRegistered('invalid');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::isRegistered('tai64');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::isRegistered('unix');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::isRegistered('test');
	var_dump($ret);
	$ret = Calends\CalendarDefinition::isRegistered('');
	var_dump($ret);
?>
--EXPECT--
bool(false)
bool(true)
bool(true)
bool(false)
bool(false)
