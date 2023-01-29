<?php

// This script looks up and grabs the latest release for the current OS/architecture

$version = json_decode(file_get_contents('https://api.github.com/repos/danhunsaker/calends/releases/latest'), true)['tag_name'];

switch (strtolower(PHP_OS_FAMILY)) {
    case 'linux':
        $os = 'linux';
        $ext = 'tgz';
        break;

    case 'windows':
        $os = 'windows';
        $ext = 'zip';
        break;

    case 'darwin':
        $os = 'darwin';
        $ext = 'tgz';
        break;

    default:
        echo 'Unsupported OS';
        exit(1);
}

switch (strtolower(php_uname('m'))) {
    case 'i86pc':
    case 'x86pc':
    case 'i386':
    case 'i686':
    case 'x86':
        $arch = '386';
        break;

    case 'x86_64':
    case 'amd64':
        $arch = 'amd64';
        break;

    case 'arm32v5':
    case 'arm/v5':
        $arch = 'arm-5';
        break;

    case 'armv6l':
    case 'arm32v6':
    case 'arm/v6':
        $arch = 'arm-6';
        break;

    case 'armv7l':
    case 'arm32v7':
    case 'arm/v7':
        $arch = 'arm-7';
        break;

    case 'aarch64':
    case 'arm64v8':
    case 'arm64':
        $arch = 'arm64';
        break;

    default:
        echo 'Unknown or unsupported processor';
        exit(1);
}

$url = "https://github.com/danhunsaker/calends/releases/download/{$version}/libcalends-{$version}-{$os}-{$arch}.{$ext}";

$ch = curl_init($url);
$root = getcwd();
$dir = "{$root}/vendor/";
$file_name = basename($url);
$save_file_loc = "{$dir}{$file_name}";
$fp = fopen($save_file_loc, 'wb');
curl_setopt($ch, CURLOPT_FILE, $fp);
curl_setopt($ch, CURLOPT_HEADER, 0);
curl_exec($ch);
curl_close($ch);
fclose($fp);

try {
    $phar = new PharData($save_file_loc);
    $phar->extractTo("{$dir}lib", null, true);
} catch (Exception $e) {
    echo "Failed to extract binary dependencies for Calends: {$e->getMessage()}";
    exit(2);
}
unlink($save_file_loc);
