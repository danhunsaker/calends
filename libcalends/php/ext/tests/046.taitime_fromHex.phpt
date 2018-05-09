--TEST--
Calends\TAITime->fromHex() Basic test
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
	$t1 = Calends\TAITime::fromHex("40000000000000000000000000000000000000000000000000000000");
	var_dump($t1);
	$t2 = Calends\TAITime::fromHex("0x40000000000000000000000000000000000000000000000000000000");
	var_dump($t2);
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
