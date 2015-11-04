#!/bin/bash -e
# make a copy of the generated test code for documentation purposes
mkdir -p sample
for f in foo.go foo_list.go num_list.go; do
  cat test/$f | sed 's/package main/package sample/' > sample/$f
done
