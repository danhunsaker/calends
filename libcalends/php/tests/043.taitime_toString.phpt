--TEST--
Calends\TAITime->toString() Basic test
--INI--
precision=20
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	$t = new Calends\TAITime(0);
	var_dump("$t");
?>
--EXPECT--
string(1) "0"
