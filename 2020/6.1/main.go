package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	group := make(map[rune]struct{})
	sum := 0
	for s.Scan() {
		t := s.Text()
		switch t {
		case "":
			sum += len(group)
			group = make(map[rune]struct{})
		default:
			for _, r := range t {
				group[r] = struct{}{}
			}
		}
	}
	sum += len(group) // final group
	fmt.Println(sum)
}
