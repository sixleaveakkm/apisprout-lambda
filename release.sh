#!/bin/sh

set -e

govvv install -o apisprout
VERSION=$(apisprout --version | cut -d ' ' -f3)

GOOS=darwin GOARCH=amd64 govvv build -o apisprout
tar -cJf apisprout-$VERSION-mac.tar.xz apisprout

GOOS=linux GOARCH=amd64 govvv build -o apisprout
tar -cJf apisprout-$VERSION-linux.tar.xz apisprout

GOOS=windows GOARCH=amd64 govvv build -o apisprout
zip -r apisprout-$VERSION-win-$GOARCH.zip apisprout.exe

rm -f apisprout apisprout.exe
