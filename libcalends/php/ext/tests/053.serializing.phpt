--TEST--
Calends\Calends serializing Basic test
--SKIPIF--
<?php
if (!extension_loaded('calends')) {
	echo 'skip';
}
?>
--FILE--
<?php
	$tmp00 = Calends\Calends::create(0);
	$tmp01 = (string)$tmp00;
	echo '__toString(0):  ', $tmp01, "\n";
	$tmp02 = serialize($tmp00);
	echo 'serialize(0):   ', $tmp02, "\n";
	$tmp03 = unserialize($tmp02);
	echo 'unserialize(0): ', (string)$tmp03, "\n";
	$tmp04 = json_encode($tmp00);
	echo 'json_encode(0): ', $tmp04, "\n";
	$tmp05 = Calends\Calends::fromJson($tmp04);
	echo 'fromJson(0):    ', (string)$tmp05, "\n";

	$tmp10 = Calends\Calends::create(['start' => 0, 'end' => 10]);
	$tmp11 = (string)$tmp10;
	echo '__toString(1):  ', $tmp11, "\n";
	$tmp12 = serialize($tmp10);
	echo 'serialize(1):   ', $tmp12, "\n";
	$tmp13 = unserialize($tmp12);
	echo 'unserialize(1): ', (string)$tmp13, "\n";
	$tmp14 = json_encode($tmp10);
	echo 'json_encode(1): ', $tmp14, "\n";
	$tmp15 = Calends\Calends::fromJson($tmp14);
	echo 'fromJson(1):    ', (string)$tmp15, "\n";
?>
--EXPECT--
__toString(0):  3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046
serialize(0):   C:15:"Calends\Calends":56:{3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046}
unserialize(0): 3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046
json_encode(0): "3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046"
fromJson(0):    3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046
__toString(1):  from 3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046 to 40000000000000020001404F3B9AC633119BB28532DB8E0D3399E120
serialize(1):   C:15:"Calends\Calends":114:{3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046::40000000000000020001404F3B9AC633119BB28532DB8E0D3399E120}
unserialize(1): from 3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046 to 40000000000000020001404F3B9AC633119BB28532DB8E0D3399E120
json_encode(1): {"start":"3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046","end":"40000000000000020001404F3B9AC633119BB28532DB8E0D3399E120"}
fromJson(1):    from 3FFFFFFFFFFFFFF93B747D4F3B9AC9F32E697BF03735DCC300FE4046 to 40000000000000020001404F3B9AC633119BB28532DB8E0D3399E120
