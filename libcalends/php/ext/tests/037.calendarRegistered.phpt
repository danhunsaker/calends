--TEST--
Calends\Calends::calendarRegistered() Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$ret = Calends\Calends::calendarRegistered('invalid');
	var_dump($ret);
	$ret = Calends\Calends::calendarRegistered('tai64');
	var_dump($ret);
	$ret = Calends\Calends::calendarRegistered('unix');
	var_dump($ret);
	$ret = Calends\Calends::calendarRegistered('test');
	var_dump($ret);
	$ret = Calends\Calends::calendarRegistered('');
	var_dump($ret);
?>
--EXPECT--
bool(false)
bool(true)
bool(true)
bool(false)
bool(false)
