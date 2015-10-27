#!/bin/bash -e
cd $(dirname $0)
rm -f *_list.go
go get ||:
go run setup.go
touch coverage.out
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
