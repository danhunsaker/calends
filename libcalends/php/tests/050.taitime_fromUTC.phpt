--TEST--
Calends\TAITime->fromUTC() Basic test
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
	var_dump($t->fromUTC());
?>
--EXPECTF--
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  int(-7)
  ["nano"]=>
  int(997489999)
  ["atto"]=>
  int(999999987)
  ["ronto"]=>
  int(778664944)
  ["udecto"]=>
  int(926276803)
  ["xindecto"]=>
  int(16662598)
}
