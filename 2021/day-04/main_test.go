package main

import (
	"bytes"
	"testing"
)

func TestFailedCaseParsing(t *testing.T) {
	text := []byte(`1,2,3

 0 74 23 35 48
21 34 97 10  2
17 25 80 89 60
15 27 94 68  7
72 75 69 32 85`)

	b := bytes.NewBuffer(text)
	_, boards := getInput(b)
	if len(boards) != 1 {
		t.Fatal("expected 1 board")
	}
	b1 := boards[0]

	if n := b1.numbers[0][0]; n != 0 {
		t.Fatalf("expected 0, got %d", n)
	}
	if n := b1.numbers[0][1]; n != 74 {
		t.Fatalf("expected 74, got %d", n)
	}
}
