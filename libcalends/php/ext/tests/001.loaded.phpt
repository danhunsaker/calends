--TEST--
Check if calends is loaded
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
  echo 'skip';
}
?>
--FILE--
<?php
echo 'The extension "calends" is available';
?>
--EXPECT--
The extension "calends" is available
