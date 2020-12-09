package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	in := readInput()
	for i := 25; i < len(in); i++ {
		if !valid(in[i-25:i], in[i]) {
			fmt.Println(in[i])
		}
	}
}
