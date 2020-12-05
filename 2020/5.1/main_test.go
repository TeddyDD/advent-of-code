package main

import (
	"fmt"
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		row    int
		column int
		index  int
	}{
		{70, 7, 567},
		{14, 7, 119},
		{102, 4, 820},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d:%d-> %d", tc.row, tc.column, tc.index), func(t *testing.T) {
			if i := id(tc.row, tc.column); i != tc.index {
				t.Fatal(i)
			}
		})
	}
}

func TestSearchRow(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{

		{
			"FBFBBFFRLR",
			44,
		},
		{
			"BFFFBBFRRR",
			70,
		},
		{
			"FFFBBBFRRR",
			14,
		},
		{
			"BBFFBBFRLL",
			102,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s -> %d", tc.in, tc.out), func(t *testing.T) {
			if res := search(0, 127, tc.in[0:7]); res != tc.out {
				t.Fatal(res)
			}
			if res := searchRow(tc.in); res != tc.out {
				t.Fatal(res)
			}
		})
	}
}

func TestSearchColumn(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"FBFBBFFRLR",
			5,
		},

		{
			"BFFFBBFRRR",
			7,
		},
		{
			"FFFBBBFRRR",
			7,
		},
		{
			"BBFFBBFRLL",
			4,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s -> %d", tc.in, tc.out), func(t *testing.T) {
			if res := search(0, 7, tc.in[7:]); res != tc.out {
				t.Fatal(res)
			}
			if res := searchColumn(tc.in); res != tc.out {
				t.Fatal(res)
			}
		})
	}
}
