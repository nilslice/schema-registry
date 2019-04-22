#! /usr/bin/env bash

set -e

PRIMARY_GOPATH=`echo $GOPATH | sed -e 's/:.*//'`
if [ -z $PRIMARY_GOPATH ]; then
	PRIMARY_GOPATH=`echo $GOPATH | sed -e 's/.*://'`
fi
PATH=$PRIMARY_GOPATH/bin:$PATH

case "$OSTYPE" in
    darwin*)  PLATFORM=osx ;;
    linux*)   PLATFORM=linux ;;
    msys*)    PLATFORM=windows ;;
    cygwin*)  PLATFORM=windows ;;
    mingw*)   PLATFORM=windows ;;
    *)        PLATFORM=unknown ;;
esac

PROTOC_VERSION=3.7.1
PROTOC_ZIP=protoc-${PROTOC_VERSION}-${PLATFORM}-x86_64.zip
PROTOC_DIR=protoc-3
PROTOC=$PROTOC_DIR/bin/protoc

if [ ! -f "$PROTOC" ]; then
    echo Setting up protoc...
    /bin/rm -rf $PROTOC_DIR
    mkdir -p $PROTOC_DIR
    cd $PROTOC_DIR
    curl -s -O -L https://github.com/google/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP
    unzip -qq $PROTOC_ZIP
    cd ..
fi

echo Setting up Go support...
export GO111MODULE=on
go get -u google.golang.org/grpc || true
go get -u github.com/golang/protobuf/{proto,protoc-gen-go} || true

echo Clearing out previously generated code...
/bin/rm -rf go

for f in *.proto ; do
    PKG=${f%.*}pb
    echo Generating Go code for $PKG
    mkdir -p go/$PKG
    $PROTOC \
        --go_out=plugins=grpc:go/$PKG $f
done
