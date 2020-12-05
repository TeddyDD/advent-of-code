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
	out = sort.IntSlice(out)

	return out
}

func main() {
	in := readInput()
	for i := 0; i < len(in)-1; i++ {
		if in[i+1]-in[i] > 1 {
			fmt.Println(in[i] + 1)
			return
		}
	}
}
