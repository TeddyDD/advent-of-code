package main

// https://adventofcode.com/2019/day/1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func fuel(mass int) int {
	return int((math.Floor(float64(mass) / 3)) - 2)
}

func main() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += fuel(num)
	}
	fmt.Println(sum)
}
