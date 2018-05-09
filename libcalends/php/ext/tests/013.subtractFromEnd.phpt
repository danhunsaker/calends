--TEST--
Calends\Calends->subtractFromEnd() Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$tmp = Calends\Calends::create();
	var_dump($tmp->subtractFromEnd());
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["goId":"Calends\Calends":private]=>
  float(%d)
}
