package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	rigth = 3
	down  = 1
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	count := 0
	y := 0
	x := 0
	for s.Scan() {
		switch {
		case x%down == 0:
			t := s.Text()
			if t[y%len(t)] == '#' {
				count += 1
			}
		}
		x += 1
		y += rigth
	}
	fmt.Println(count)
}
