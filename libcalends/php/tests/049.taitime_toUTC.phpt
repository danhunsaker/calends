--TEST--
Calends\TAITime->toUTC() Basic test
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
	var_dump($t->toUTC());
?>
--EXPECTF--
object(Calends\TAITime)#%d (6) {
  ["seconds"]=>
  int(8)
  ["nano"]=>
  int(81999)
  ["atto"]=>
  int(999999027)
  ["xicto"]=>
  int(295416453)
  ["ucto"]=>
  int(853249549)
  ["rocto"]=>
  int(865722656)
}
