--TEST--
Calends\Calends->endDate() Basic test
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
	echo $tmp->endDate(), "\n";
	$tmp = Calends\Calends::create('', 'tai64', 'decimal');
	echo $tmp->endDate(), "\n";
	$tmp = Calends\Calends::create('10', 'tai64', 'decimal');
	echo $tmp->endDate(), "\n";
?>
--EXPECT--
1.994980000
8.000082000
18.000082000
