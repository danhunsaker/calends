#!/usr/bin/env pwsh

Join-Path 'libcalends' 'php' | Set-Location

$diffs = ( Get-ChildItem -Path tests/ -Filter '*.diff' -Recurse | Sort-Object )

if ( ( Write-Output $diffs | Measure-Object ).Count -gt 0 ) { 
  Write-Output $diffs | ForEach-Object {
    $diff = $_.Name
    $php = $_.Name -replace 'diff$','php'

    Write-Output "`n--------`n$_`n========"
    Get-Content "tests/$php"
    Write-Output "--RESULT--"
    Get-Content "tests/$diff"
    Write-Output ''
  }

  [Environment]::Exit(1)
}
