#!/bin/bash -e
rm -f *_slice.go
go get
go run setup.go
touch coverage.out
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
