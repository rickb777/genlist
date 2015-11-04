#!/bin/bash -e
rm -f sample/*.go

./internal/list/test/test.sh
./internal/option/test/test.sh
./internal/set/test/test.sh
./internal/plumbing/test/test.sh
./test/test.sh
go test .

./makeSamples.sh
