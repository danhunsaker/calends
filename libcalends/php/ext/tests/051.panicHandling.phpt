--TEST--
Panic handling test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	try {
		$tmp = new Calends\Calends;
	} catch (Error $e) {
		echo "Caught Error: {$e->getMessage()}\n";
	}

	try {
		Calends\Calends::create(null, 'invalid');
	} catch (Calends\CalendsException $e) {
		echo "Caught Calends\CalendsException: {$e->getMessage()}\n";
	}
?>
--EXPECTF--
Caught Error: Call to private Calends\Calends::__construct() from invalid context
Caught Calends\CalendsException: Unknown Calendar: invalid
