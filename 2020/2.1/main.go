package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type policy struct {
	letter string
	min    int
	max    int
}

type password struct {
	pwd    string
	policy policy
}

var extract = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<letter>[a-z]): (?P<password>\w+)`)

func matches(exp *regexp.Regexp, str string) map[string]string {
	result := make(map[string]string)
	match := exp.FindStringSubmatch(str)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

func readInput() []password {
	out := []password{}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		m := matches(extract, s.Text())
		min, err := strconv.Atoi(m["min"])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(m["max"])
		if err != nil {
			panic(err)
		}

		res := password{
			pwd: m["password"],
			policy: policy{
				letter: m["letter"],
				min:    min,
				max:    max,
			},
		}

		out = append(out, res)
	}

	return out
}

func isValid(p password) bool {
	c := strings.Count(p.pwd, p.policy.letter)
	if c >= p.policy.min && c <= p.policy.max {
		return true
	}
	return false
}

func main() {
	in := readInput()
	ok := 0
	for _, p := range in {
		if isValid(p) {
			ok += 1
		}
	}
	fmt.Println(ok)
}
