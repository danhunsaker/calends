--TEST--
Calends\TAITime->fromUTC() Basic test
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
	var_dump($t->fromUTC());
?>
--EXPECTF--
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  float(-7)
  ["nano"]=>
  int(997489999)
  ["atto"]=>
  int(999999987)
  ["xicto"]=>
  int(778664944)
  ["ucto"]=>
  int(926276803)
  ["rocto"]=>
  int(16662598)
}
