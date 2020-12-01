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

func findThree(arr []int, sum int) {
	for i := 0; i < len(arr)-2; i++ {
		l := i + 1
		r := len(arr) - 1
		for l < r {
			n := arr[i] + arr[l] + arr[r]
			switch {
			case n == sum:
				fmt.Println(arr[i], arr[l], arr[r])
				fmt.Println(arr[i] * arr[l] * arr[r])
				return
			case n < sum:
				l += 1
			case n > sum:
				r -= 1
			}
		}
	}
}

func main() {
	nums := sort.IntSlice(readInput())
	sort.Sort(nums)

	findThree(nums, 2020)
}
