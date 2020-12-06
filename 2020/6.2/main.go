package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var group *uint = nil
	sum := 0
	for s.Scan() {
		t := s.Text()
		switch t {
		case "":
			sum += bits.OnesCount(*group)
			group = nil
		default:
			var current uint = 0
			for _, r := range t {
				current |= 1 << (r - 97)
			}
			if group == nil {
				group = &current
			} else {
				*group = *group & current
			}
		}
	}
	sum += bits.OnesCount(*group)
	fmt.Println(sum)
}
