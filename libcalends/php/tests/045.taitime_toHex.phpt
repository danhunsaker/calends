--TEST--
Calends\TAITime->toHex() Basic test
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
	var_dump($t->toHex());
?>
--EXPECT--
string(56) "40000000000000000000000000000000000000000000000000000000"
