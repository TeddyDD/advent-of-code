package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Map [][]int

func (m Map) LowPoint(x, y int) bool {
	n := make([]int, 0, 4)
	if y > 0 {
		n = append(n, m[y-1][x])
	}
	if y < len(m)-1 {
		n = append(n, m[y+1][x])
	}

	if x > 0 {
		n = append(n, m[y][x-1])
	}
	if x < len(m[y])-1 {
		n = append(n, m[y][x+1])
	}

	for i := range n {
		if n[i] <= m[y][x] {
			return false
		}
	}

	if l := len(n); l != 4 && l != 3 && l != 2 {
		panic(fmt.Sprintf("%d %d: %+v", x, y, n))
	}

	return true
}

func (m Map) Iterate(fn func(x, y, v int)) {
	for y := 0; y < len(m); y++ {
		xlen := len(m[y])
		for x := 0; x < xlen; x++ {
			fn(x, y, m[y][x])
		}
	}
}

func LoadMap(in io.Reader) Map {
	res := Map{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		l := make([]int, len(line))
		for i, letter := range line {
			n, _ := strconv.Atoi(string(letter))
			l[i] = n
		}
		res = append(res, l)
	}
	return res
}

func main() {
	input := LoadMap(os.Stdin)
	sum := 0
	input.Iterate(func(x int, y int, v int) {
		if input.LowPoint(x, y) {
			sum = sum + 1 + v
		}
	})
	fmt.Println(sum)
}
