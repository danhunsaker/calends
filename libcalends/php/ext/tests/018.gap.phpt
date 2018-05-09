--TEST--
Calends\Calends->gap() Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$tmp1 = Calends\Calends::create(0);
	$tmp2 = Calends\Calends::create(10);
	var_dump($tmp1->gap($tmp2));
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["goId":"Calends\Calends":private]=>
  float(%d)
}
