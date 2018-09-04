#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $BASEDIR
if [ -e build ]; then
	rm -rf build
fi
mkdir -p build

cd $BASEDIR/src
GOOS=darwin   GOARCH=amd64 bash -xec 'go build -o ../build/webserver.$GOOS-$GOARCH main.go'
GOOS=linux    GOARCH=amd64 bash -xec 'go build -o ../build/webserver.$GOOS-$GOARCH main.go'
GOOS=linux    GOARCH=arm   bash -xec 'go build -o ../build/webserver.$GOOS-$GOARCH main.go'
GOOS=windows  GOARCH=amd64 bash -xec 'go build -o ../build/webserver.$GOOS-$GOARCH.exe main.go'

