--TEST--
Calends\TAITime->fromString() Basic test
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

	$t = Calends\TAITime::fromString("0");
	var_dump($t);
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
