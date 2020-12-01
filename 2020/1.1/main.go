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

func find(arr []int, sum int) (int, int) {
	l := 0
	r := len(arr) - 1

	for l < r {
		n := arr[l] + arr[r]
		switch {
		case n == sum:
			return arr[l], arr[r]
		case n < sum:
			l += 1
		case n > sum:
			r -= 1
		}
	}
	return 0, 0
}

func main() {
	nums := sort.IntSlice(readInput())
	sort.Sort(nums)

	a, b := find(nums, 2020)
	fmt.Println(a, b)
	fmt.Println(a * b)
}
