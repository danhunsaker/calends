--TEST--
Calends\Calends->intersect() Basic test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	$tmp1 = Calends\Calends::create();
	$tmp2 = Calends\Calends::create();
	var_dump($tmp1->intersect($tmp2));
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["start"]=>
  string(%d) "%x"
  ["end"]=>
  string(%d) "%x"
}
