#!/bin/bash

APPNAME=ntxnicstatus
VERSION=`cat VERSION`

cat /dev/null > "bin/$APPNAME-$VERSION-SHA256SUMS"

export GOOS=linux
export GOARCH=amd64
go build -o bin/$APPNAME
pushd bin/
zip "$APPNAME-$VERSION-$GOOS-$GOARCH.zip" $APPNAME
shasum -a 256 "$APPNAME-$VERSION-$GOOS-$GOARCH.zip" >> "$APPNAME-$VERSION-SHA256SUMS"
rm -f $APPNAME
popd

export GOOS=darwin
export GOARCH=amd64
go build -o bin/$APPNAME
pushd bin/
zip "$APPNAME-$VERSION-$GOOS-$GOARCH.zip" $APPNAME
shasum -a 256 "$APPNAME-$VERSION-$GOOS-$GOARCH.zip" >> "$APPNAME-$VERSION-SHA256SUMS"
rm -f $APPNAME
popd

export GOOS=windows
export GOARCH=amd64
go build -o bin/$APPNAME
pushd bin/
zip "$APPNAME-$VERSION-$GOOS-$GOARCH.zip" $APPNAME
shasum -a 256 "$APPNAME-$VERSION-$GOOS-$GOARCH.zip" >> "$APPNAME-$VERSION-SHA256SUMS"
rm -f $APPNAME
popd
