--TEST--
Check if ffi is loaded
--SKIPIF--
<?php
if (!extension_loaded('ffi')) {
  echo 'skip';
}
?>
--FILE--
<?php
require_once('vendor/autoload.php');

if (class_exists('Calends\\Calends')) {
  echo 'The "Calends\\Calends" class is available';
}
?>
--EXPECT--
The "Calends\Calends" class is available
