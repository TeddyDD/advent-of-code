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

type bounds struct {
	cur   int
	start int
}

var mem = make(map[bounds]int)

func count(this, start int, adapters []int) int {
	if res, ok := mem[bounds{this, start}]; ok {
		return res
	}
	if len(adapters) == start {
		return 1
	}
	sum := 0
	for i := start; i < start+3; i++ {
		if i < len(adapters) && adapters[i] <= this+3 {
			sum += count(adapters[i], i+1, adapters)
		}
	}
	mem[bounds{this, start}] = sum
	return sum
}

func main() {
	in := readInput()
	sort.Ints(in)

	count := count(0, 0, in)

	fmt.Printf("%#v\n", in)
	fmt.Println(count)
}
