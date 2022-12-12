--TEST--
Panic handling test
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
	echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

	try {
		$tmp = new Calends\Calends(0);
	} catch (Error $e) {
		echo "Caught Error: {$e->getMessage()}\n";
	}

	try {
		Calends\Calends::create('', 'invalid');
	} catch (Calends\CalendsException $e) {
		echo "Caught Calends\CalendsException: {$e->getMessage()}\n";
	}
?>
--EXPECTF--
Caught Error: Call to private Calends\Calends::__construct() from %s
Caught Calends\CalendsException: Unknown Calendar: invalid
