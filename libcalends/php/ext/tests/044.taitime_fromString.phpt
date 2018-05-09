--TEST--
Calends\TAITime->fromString() Basic test
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
	$t = Calends\TAITime::fromString("0");
	var_dump($t);
?>
--EXPECTF--
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  float(0)
  ["nano"]=>
  int(0)
  ["atto"]=>
  int(0)
  ["xicto"]=>
  int(0)
  ["ucto"]=>
  int(0)
  ["rocto"]=>
  int(0)
}
