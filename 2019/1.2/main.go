package main

// https://adventofcode.com/2019/day/1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuel(mass int) int {
	return (mass / 3) - 2
}

func stageFuel(mass int) int {
	res := fuel(mass)
	for add := fuel(res); add > 0; add = fuel(add) {
		res += add
	}
	return res
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
		sum += stageFuel(num)
	}
	fmt.Println(sum)
}
