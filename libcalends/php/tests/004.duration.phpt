--TEST--
Calends\Calends->duration() Basic test
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
	echo $tmp->duration(), "\n";
	$tmp = Calends\Calends::create('10', 'tai64', 'decimal');
	echo $tmp->duration(), "\n";
	$tmp = Calends\Calends::create(['start' => '0', 'end' => '10'], 'tai64', 'decimal');
	echo $tmp->duration(), "\n";
?>
--EXPECT--
0
0
10
