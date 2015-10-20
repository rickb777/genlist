#!/bin/bash -e
./internal/option/test/test.sh
./internal/list/test/test.sh
./test/test.sh
go test .
