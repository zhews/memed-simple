#!/bin/bash

for cmd_dir in ./cmd/*
do
  cmd="${cmd_dir/\.\/cmd\/}"
  go build -o ./bin/"$cmd" "$cmd_dir"
done
