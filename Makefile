SHELL := /usr/bin/env bash

all: forge gen
.PHONY: all forge gen

forge:
	forge build

gen:
	go run gen/gen.go -config genconfig.toml -base-path ./
