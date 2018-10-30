#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
set -eu
echo "******************"
echo "Running unit tests"
cd $DIR/../gtg
go test
echo "******************"
echo "Using gometalinter with misspell, vet, ineffassign, and gosec"
echo "Testing $DIR/../gtg"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../gtg