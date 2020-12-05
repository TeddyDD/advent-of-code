package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() []string {
	out := []string{}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		out = append(out, s.Text())
	}

	return out
}

func main() {
	in := readInput()
	max := 0

	for _, p := range in {
		r := searchRow(p)
		c := searchColumn(p)
		id := id(r, c)
		if id > max {
			max = id
		}
		fmt.Printf("%d\t%s - %d %d\n", id, p, r, c)
	}
	fmt.Fprintln(os.Stderr, max)
}
