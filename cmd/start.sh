#!/bin/bash

while getopts "e:" arg; do
  case $arg in
    e) Env=$OPTARG;;
  esac
done

source files/etc/env/env.$Env.sh

# Start the HTTP service locally
go run ./cmd/app-http/*.go -env=$Env