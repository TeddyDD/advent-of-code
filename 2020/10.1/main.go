package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readInput() []int {
	out := []int{}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		out = append(out, n)
	}

	return out
}
func main() {
	in := readInput()
	sort.Ints(in)

	in = append([]int{0}, in...)
	in = append(in, in[len(in)-1]+3)
	fmt.Printf("%#v\n", in)

	d1 := 0
	d3 := 0

	for i := 1; i < len(in); i++ {
		switch c := in[i] - in[i-1]; c {
		case 1:
			d1++
		case 3:
			d3++
		}
	}
	fmt.Println(d1, d3, d1*d3)
}
