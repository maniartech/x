#!/bin/bash

echo "Building parser ..."

pigeon.exe parser.peg > parser.go

echo "The parser is ready!"
