--TEST--
Calends\Calends->withDurationFromEnd() Basic test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	$tmp = Calends\Calends::create('10', 'tai64', 'decimal');
	var_dump($ret = $tmp->withDurationFromEnd(10));
	echo $ret->date(), "\n";
	echo $ret->duration(), "\n";
	echo $ret->endDate(), "\n";
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["start"]=>
  string(%d) "%x"
  ["end"]=>
  string(%d) "%x"
}
8.000082000
10
18.000082000
