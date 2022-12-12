#!/usr/bin/env pwsh

go get -v ./...
go mod vendor

Set-Content -Value "mode: atomic" -Path coverage.txt

go list ./... | Select-String -NotMatch -Pattern 'vendor' | ForEach-Object {
    go test -v -coverprofile profile.out -covermode atomic "$_"
    if (Test-Path profile.out) {
        Select-String -NotMatch -Pattern '^mode:' -Path profile.out -Raw | Add-Content -Path coverage.txt
        Remove-Item profile.out
    }
}
