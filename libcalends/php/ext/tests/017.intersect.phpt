--TEST--
Calends\Calends->intersect() Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$tmp1 = Calends\Calends::create();
	$tmp2 = Calends\Calends::create();
	var_dump($tmp1->intersect($tmp2));
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["goId":"Calends\Calends":private]=>
  float(%d)
}
