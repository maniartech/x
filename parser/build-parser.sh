#!/bin/bash

echo "Building parser ..."

pigeon.exe ./grammer.peg > ./grammer.go

echo "The parser is ready!"
