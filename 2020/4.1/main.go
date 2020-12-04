package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type passport = map[string]string

var requiredFields = []string{"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid",
}

func parseKv(line string, p *passport) {
	fields := strings.Split(line, " ")
	for _, f := range fields {
		kv := strings.Split(f, ":")
		(*p)[kv[0]] = kv[1]
	}
}

func readInput() []passport {
	res := []passport{}

	s := bufio.NewScanner(os.Stdin)
	cur := make(map[string]string)
	for s.Scan() {
		switch t := s.Text(); t {
		case "":
			res = append(res, cur)
			cur = make(map[string]string)
		default:
			parseKv(t, &cur)
		}
	}
	res = append(res, cur)

	return res
}

func valid(p passport) bool {
	for _, fname := range requiredFields {
		if k, ok := p[fname]; !ok || k == "" {
			return false
		}
	}
	return true
}

func main() {

	in := readInput()
	count := 0

	for _, p := range in {
		fmt.Printf("%v ", p)
		if valid(p) {
			fmt.Printf(" valid")
			count += 1
		}
		fmt.Println()
	}
	fmt.Println(count)
}
