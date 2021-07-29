#!/bin/bash

echo "Building x.wasm ..."
tinygo build -o x.wasm -target wasm ./cmd/main.go

# Gzip wasm
gzip -9 -v -c x.wasm > x.wasm.gz

# List output
ls x.wasm* -al --block-size=KB

echo "WebAssembly is ready!"
