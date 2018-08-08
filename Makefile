.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run test deps

run: test

test:
	go test -count=1 -v

deps:
	go get -u -v ./...
