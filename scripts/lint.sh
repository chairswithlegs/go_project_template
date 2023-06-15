#!/bin/sh
SCRIPT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_DIR="$SCRIPT_DIR/.."

export MSYS_NO_PATHCONV=1 # Prevents a volume mounting error when using Git Bash

cd $ROOT_DIR || exit 1
docker run --rm -v "$(pwd)":/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v
