#!/bin/bash -e
rm -f sample/*.go

./internal/list/test/test.sh
./internal/option/test/test.sh
./internal/set/test/test.sh
./internal/plumbing/test/test.sh
./test/test.sh
go test .

# make a copy of the generated test code for documentation purposes
mkdir -p sample
for f in foo.go foo_list.go num_list.go; do
  cat test/$f | sed 's/package main/package sample/' > sample/$f
done
