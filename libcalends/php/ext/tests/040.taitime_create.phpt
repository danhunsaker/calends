--TEST--
Calends\TAITime::create() Basic test
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
	$ret = new Calends\TAITime();
	var_dump($ret);
	$ret = new Calends\TAITime("0");
	var_dump($ret);
	$ret = new Calends\TAITime("0x40000000000000000000000000000000000000000000000000000000");
	var_dump($ret);
	$ret = new Calends\TAITime(0);
	var_dump($ret);
	$ret = new Calends\TAITime(0.0);
	var_dump($ret);
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
