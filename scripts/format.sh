#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
set -eu
echo "******************"
echo "Formatting"
cd $DIR/../gtg
go fmt