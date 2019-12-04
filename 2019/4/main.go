package main

import (
	"fmt"
	"sort"
	"strconv"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

func repeatCheck(n int) bool {
	prev := -1
	for n > 0 {
		c := n % 10
		if prev == c {
			return true
		}
		prev = c
		n /= 10
	}
	return false
}

func main() {
    // first part only
	count := 0
	for i := 158126; i <= 624574; i++ {
		s := strconv.Itoa(i)
		if repeatCheck(i) && sorted(s) == s && len(s) == 6 {
			count += 1
		}
	}
	fmt.Println(count)
}
