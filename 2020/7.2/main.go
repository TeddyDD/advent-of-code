package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func count(color string, b bags) int {
	sum := 1
	for _, in := range b[color] {
		sum += count(in.color, b) * in.ammount
	}
	return sum
}

type inner struct {
	color   string
	ammount int
}

type bags = map[string][]inner

func main() {
	s := bufio.NewScanner(os.Stdin)

	split := regexp.MustCompile(`^(\w+\s\w+) bags contain (.+)$`)
	color := regexp.MustCompile(`(\d) ((?:\w+\s\w+)) bags?`)
	noother := regexp.MustCompile(`no other bags`)

	allBags := make(bags)

	for s.Scan() {
		t := s.Text()
		res := split.FindStringSubmatch(t)
		currentBagColor := res[1]
		contains := strings.Split(res[2], ",")
		for _, desc := range contains {
			if !noother.MatchString(desc) {
				matches := color.FindStringSubmatch(desc)[1:]
				ammount, err := strconv.Atoi(matches[0])
				if err != nil {
					panic(err)
				}
				i := inner{
					color:   matches[1], // name
					ammount: ammount,
				}
				allBags[currentBagColor] = append(allBags[currentBagColor], i)
			}
		}
	}
	fmt.Println(count("shiny gold", allBags) - 1)
}
