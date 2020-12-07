package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func contains(bag, color string, bags map[string][]string) bool {
	for _, inner := range bags[bag] {
		if inner == "no other" {
			return false
		}
		if inner == color {
			return true
		}
		if contains(inner, color, bags) {
			return true
		}
	}
	return false
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	split := regexp.MustCompile(`^(\w+\s\w+) bags contain (.+)$`)
	color := regexp.MustCompile(`(?:\d )?((?:\w+\s\w+)|(?:no other)) bags?`)

	bags := make(map[string][]string)
	count := 0

	for s.Scan() {
		t := s.Text()
		res := split.FindStringSubmatch(t)
		bag := res[1]
		rest := strings.Split(res[2], ",")
		for _, r := range rest {
			cur := color.FindStringSubmatch(r)[1:]
			for _, c := range cur {
				bags[bag] = append(bags[bag], c)
			}
		}
	}
	fmt.Printf("%#v\n", bags)
	for b := range bags {
		if contains(b, "shiny gold", bags) {
			count += 1
		}
	}
	fmt.Println(count)
}
