package main

import (
	"os"
	"testing"
)

func Test0Visible(t *testing.T) {
	file, _ := os.Open("./0visible")
	defer file.Close()
	f := readInput(file)
	o, _ := f.visible(3, 3)
	if 0 != o {
		t.Fatalf("should be 0, not %d", o)
	}
}

func Test8Visible(t *testing.T) {
	file, _ := os.Open("./8visible")
	defer file.Close()
	f := readInput(file)
	o, _ := f.visible(3, 4)
	if 8 != o {
		t.Fatalf("should be 8, not %d", o)
	}
}
