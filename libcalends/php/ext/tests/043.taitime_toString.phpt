--TEST--
Calends\TAITime->toString() Basic test
--INI--
precision=20
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$t = new Calends\TAITime(0);
	var_dump($t->toString());
?>
--EXPECT--
string(1) "0"
