#!/bin/bash

export OLD_DYLD=$DYLD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=$GOPATH/clibs/lib
#GOARCH=amd64 GOOS=linux go build -ldflags -extldflags=-L$DYLD_LIBRARY_PATH
go build -ldflags -extldflags=-L$DYLD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=$OLD_DYLD
