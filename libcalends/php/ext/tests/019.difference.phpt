--TEST--
Calends\Calends->difference() Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$tmp1 = Calends\Calends::create('0', 'tai64', 'decimal');
	$tmp2 = Calends\Calends::create('10', 'tai64', 'decimal');
	var_dump($tmp1->difference($tmp2));
?>
--EXPECT--
string(3) "-10"
