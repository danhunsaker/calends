--TEST--
Calends\TAITime->add() Basic test
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
	$t1 = new Calends\TAITime(0);
	$t2 = new Calends\TAITime(10);
	var_dump($t1->add($t2));
?>
--EXPECTF--
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  float(10)
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
