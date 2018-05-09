--TEST--
Calends\Calends->isLonger() Basic test
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
	var_dump($tmp1->isLonger($tmp2));
?>
--EXPECT--
bool(false)
