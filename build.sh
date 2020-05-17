#!/bin/bash

APPNAME=ntxnicstatus

export GOOS=linux
export GOARCH=amd64
go build -o bin/linux/amd64/$APPNAME

export GOOS=darwin
export GOARCH=amd64
go build -o bin/darwin/amd64/$APPNAME

export GOOS=windows
export GOARCH=amd64
go build -o bin/windows/amd64/$APPNAME.exe