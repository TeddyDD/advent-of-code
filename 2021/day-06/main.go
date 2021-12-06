package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fish := make([]int, 0, 300)
	for _, n := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(n)
		fish = append(fish, num)
	}

	state := [9]int{}
	for _, f := range fish {
		state[f] += 1
	}

	for i := 0; i < 256; i++ {
		tmp := state[0]
		for j := 1; j <= 8; j++ {
			state[j-1] = state[j]
			state[j] = 0
		}
		state[6] += tmp
		state[8] = tmp
	}

	sum := 0
	for i := range state {
		sum += state[i]
	}
	fmt.Println(state, sum)
}
