#!/bin/sh
# prepare input for part 2

go run main.go search.go <input | sort -n 2>/dev/null >out
