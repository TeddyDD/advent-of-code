package main

import (
	"bufio"
	"flag"
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

func valid(s []int, current int) bool {
	// nested loops go brrrr
	for i := range s {
		for j := range s {
			if s[i]+s[j] == current && s[i] != s[j] {
				return true
			}
		}
	}
	return false
}

func sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func main() {
	in := readInput()
	target := flag.Int("first", 375054920, "first invalid number in sequence")
	flag.Parse()
	res := []int{}

	for i := 0; i < len(in)-1; i++ {
		for j := i; j < len(in); j++ {
			if sum(in[i:j]) == *target && len(in[i:j]) >= 2 {
				res = in[i:j]
				break
			}
		}
	}
	sort.Ints(res)
	fmt.Printf("%#v\n", res)
	fmt.Println(res[0] + res[len(res)-1])
}
