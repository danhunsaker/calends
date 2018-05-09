--TEST--
Calends\Calends->withDuration() Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$tmp = Calends\Calends::create();
	var_dump($ret = $tmp->withDuration(10));
	echo $ret->date(), "\n";
	echo $ret->duration(), "\n";
	echo $ret->endDate(), "\n";
?>
--EXPECTF--
object(Calends\Calends)#%d (%d) {
  ["goId":"Calends\Calends":private]=>
  float(%d)
}
1.994980000
10.002592
11.994980000
