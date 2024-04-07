SHELL := /usr/bin/env bash

lower = $(shell echo '$1' | tr '[:upper:]' '[:lower:]')

CONTRACTS := $(basename $(notdir $(wildcard ./src/*.sol)))
BINDINGS := $(foreach contract,$(CONTRACTS),$(call lower,$(contract)))

all: forge gen
.PHONY: all forge gen

forge:
	forge build

gen:
	go run gen/gen.go -config genconfig.toml -base-path ./

$(BINDINGS:%=bindings/%.go): forge
	abigen --abi <(jq '.["abi"]' "out/$(basename $(notdir $@)).sol/$(basename $(notdir $@)).json") --pkg bindings --out $@

bindings: $(BINDINGS:%=bindings/%.go)
