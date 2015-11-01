#!/bin/bash -e
./internal/list/test/test.sh
./internal/option/test/test.sh
./internal/set/test/test.sh
./test/test.sh
go test .
