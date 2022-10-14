#!/bin/bash
antlr -Xexact-output-dir -Dlanguage=Go -o $PWD -visitor -package parser $PWD/JSON.g4
