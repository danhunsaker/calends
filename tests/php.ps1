#!/usr/bin/env pwsh

param([string]$php, [string]$ext = 'so')

# Ensure everything is built, first, so we're testing the latest code
$start = Get-Location
$env:VSLocation = (Get-VSSetupInstance).InstallationPath
Import-Module "$env:VSLocation\Common7\Tools\Microsoft.VisualStudio.DevShell.dll"
Enter-VsDevShell -VsInstallPath "$env:VSLocation"

go get -v ./...
Set-Location libcalends
go build -v -o libcalends.${ext} -buildmode c-shared .

# Run the actual tests
Set-Location php

Invoke-WebRequest "https://windows.php.net/downloads/releases/php-test-pack-${php}.zip" -OutFile 'ptp.zip'
Expand-Archive 'ptp.zip'

Set-Content -Value "#define FFI_SCOPE `"CALENDS`"" -Path tmp.h
Add-Content -Value "#define FFI_LIB `"../libcalends.${ext}`"" -Path tmp.h
cl /EP /D"__SIZE_TYPE__=unsigned long" ../libcalends.h | Add-Content -Path tmp.h
$text = ""
Get-Content tmp.h | ForEach-Object { $text += $_ + "\r\n" }
$text = $text -Replace "(#|__)pragma( |\()(pack|warning)\(push.+?pragma( |\()(pack|warning)\(pop.+?\\r\\n|__declspec\(dllexport\) ", ""
$text.Replace("\r\n", "`r`n") | Set-Content -Path tmp.h
Get-Content -Path tmp.h | Select-String -NotMatch -Pattern 'complex|^static inline .*\{$|return |^\}$|^\s*$|pragma' | Set-Content -Path Calends.h

composer install

$s = if ([System.Version]$php -ge [System.Version]"8.0.0") { "s" }
$env:NO_INTERACTION = 'true'
php ./ptp/run-test${s}.php -P -d "ffi.preload=Calends.h"

# Remove-Item -Recurse ptp, Calends.h, tmp.h

Set-Location $start
./tests/failed_test_info.ps1
[Environment]::Exit($LastExitCode)
