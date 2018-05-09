#!/bin/bash

go test ./... -race -coverprofile coverage.out
go tool cover -html coverage.out -o coverage.htm
