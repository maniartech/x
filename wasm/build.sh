#!/usr/bin/bash

echo "Building x.wasm"

# Remove current wasms
rm x.wasm
rm x.wasm.gz

if [[ "$OSTYPE" == "win32" ]]; then
  GOOS=js && GOARCH=wasm && go build -o x.wasm
else
  GOOS=js GOARCH=wasm go build -o x.wasm main.go
fi

# rm x.wasm
# Gzip wasm
gzip -9 -v -c x.wasm > x.wasm.gz

# List output
ls x.wasm* -al --block-size=KB

echo "Done!"
