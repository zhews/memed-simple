#!/bin/bash

for microservice_dir in ./build/package/*; do
  microservice="${microservice_dir/\.\/build\/package\//}"
  docker build -t "memed-$microservice" -f "$microservice_dir/Dockerfile" .
done
