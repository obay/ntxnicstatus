#!/bin/bash

APPNAME=ntxnicstatus
VERSION=`cat VERSION`

export GOOS=linux
export GOARCH=amd64
go build -o bin/$VERSION/linux/amd64/$VERSION/$APPNAME

export GOOS=darwin
export GOARCH=amd64
go build -o bin/$VERSION/darwin/amd64/$APPNAME

export GOOS=windows
export GOARCH=amd64
go build -o bin/$VERSION/windows/amd64/$APPNAME.exe