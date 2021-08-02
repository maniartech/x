#!/bin/bash

echo "Building AST parser ..."

pigeon.exe grammer.peg > grammer.go

echo "The AST parser is ready!"
