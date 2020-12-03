package main

import (
	"bufio"
	"fmt"
	"os"
)

func trees(arr []string, rigth, down int) int {
	count := 0

	x := 0
	for y := 0; y < len(arr); y += down {
		if y%down == 0 {
			if l := arr[y]; l[x%len(l)] == '#' {
				count += 1
			}
		}
		x += rigth
	}

	return count
}

func main() {
	lines := []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	slopes := []struct {
		rigth int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	res := 1
	for _, s := range slopes {
		res *= trees(lines, s.rigth, s.down)
	}

	fmt.Println(res)
}
