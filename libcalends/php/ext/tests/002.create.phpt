--TEST--
Calends\Calends::create() Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$ret = Calends\Calends::create();
	var_dump($ret);
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["goId":"Calends\Calends":private]=>
  float(%d)
}
