#!/bin/bash
clc -s
cat Version.dat
go mod tidy
go fmt .
staticcheck . | grep -v actions.go.*ST1005
go vet .
golangci-lint run
git st
