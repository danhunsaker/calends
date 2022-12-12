--TEST--
Calends\Calends->next() Basic test
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
	var_dump($tmp->next());
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["start"]=>
  string(%d) "%x"
  ["end"]=>
  string(%d) "%x"
}
