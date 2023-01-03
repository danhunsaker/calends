--TEST--
Calends\TAITime::create() Basic test
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
  int(0)
  ["nano"]=>
  int(0)
  ["atto"]=>
  int(0)
  ["ronto"]=>
  int(0)
  ["udecto"]=>
  int(0)
  ["xindecto"]=>
  int(0)
}
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  int(0)
  ["nano"]=>
  int(0)
  ["atto"]=>
  int(0)
  ["ronto"]=>
  int(0)
  ["udecto"]=>
  int(0)
  ["xindecto"]=>
  int(0)
}
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  int(0)
  ["nano"]=>
  int(0)
  ["atto"]=>
  int(0)
  ["ronto"]=>
  int(0)
  ["udecto"]=>
  int(0)
  ["xindecto"]=>
  int(0)
}
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  int(0)
  ["nano"]=>
  int(0)
  ["atto"]=>
  int(0)
  ["ronto"]=>
  int(0)
  ["udecto"]=>
  int(0)
  ["xindecto"]=>
  int(0)
}
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  int(0)
  ["nano"]=>
  int(0)
  ["atto"]=>
  int(0)
  ["ronto"]=>
  int(0)
  ["udecto"]=>
  int(0)
  ["xindecto"]=>
  int(0)
}
