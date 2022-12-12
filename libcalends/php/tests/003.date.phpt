--TEST--
Calends\Calends->date() Basic test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	$tmp = Calends\Calends::create();
	echo $tmp->date(), "\n";
	$tmp = Calends\Calends::create('', 'tai64', 'decimal');
	echo $tmp->date(), "\n";
	$tmp = Calends\Calends::create('10', 'tai64', 'decimal');
	echo $tmp->date(), "\n";
?>
--EXPECT--
1.994980000
8.000082000
18.000082000
