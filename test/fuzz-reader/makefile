.PHONY: fuzz

fuzz: fuzz-build fuzz-run

fuzz-build:
	GO111MODULE=on go mod vendor
	CGO_ENABLED=1 GO111MODULE=off go-fuzz-build github.com/moov-io/wire/test/fuzz-reader

fuzz-run:
	CGO_ENABLED=1 GO111MODULE=off go-fuzz -bin=./fuzzreader-fuzz.zip -workdir=$(shell pwd)

install:
	CGO_ENABLED=1 GO111MODULE=off go get -u github.com/dvyukov/go-fuzz/...
