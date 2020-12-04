package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func validNum(str string, min, max int) bool {
	n, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return n >= min && n <= max
}

func validByr(p passport) bool {
	if byr, ok := p["byr"]; ok && validNum(byr, 1920, 2002) {
		return true
	}
	fmt.Printf("%v - invalid byr\n", p)
	return false
}

func validIyr(p passport) bool {
	if byr, ok := p["iyr"]; ok && validNum(byr, 2010, 2020) {
		return true
	}
	fmt.Printf("%v - invalid iyr\n", p)
	return false
}

func validEyr(p passport) bool {
	if byr, ok := p["eyr"]; ok && validNum(byr, 2020, 2030) {
		return true
	}
	fmt.Printf("%v - invalid eyr\n", p)
	return false
}

var hgtRegx = regexp.MustCompile(`^(\d+)(cm|in)$`)

func validHgt(p passport) bool {
	if hgt, ok := p["hgt"]; ok && hgtRegx.MatchString(hgt) {
		m := hgtRegx.FindStringSubmatch(hgt)
		fmt.Printf("%+v\n", m)
		switch m[2] {
		case "cm":
			return validNum(m[1], 150, 193)
		case "in":
			return validNum(m[1], 59, 76)
		}
	}
	fmt.Printf("%v - invalid hgt\n", p)
	return false
}

var hclRegx = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func validHcl(p passport) bool {
	if hcl, ok := p["hcl"]; ok && hclRegx.MatchString(hcl) {
		return true
	}
	return false
}

var eclRegex = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)

func validEcl(p passport) bool {
	if ecl, ok := p["ecl"]; ok && eclRegex.MatchString(ecl) {
		return true
	}
	return false
}

var pidRegex = regexp.MustCompile(`^[0-9]{9}$`)

func validPid(p passport) bool {
	if pid, ok := p["pid"]; ok && pidRegex.MatchString(pid) {
		return true
	}
	return false
}

func valid(p passport) bool {
	for _, fname := range requiredFields {
		if k, ok := p[fname]; !ok || k == "" {
			return false
		}
	}
	return validByr(p) &&
		validEyr(p) &&
		validIyr(p) &&
		validHgt(p) &&
		validHcl(p) &&
		validEcl(p) &&
		validPid(p)
}

func main() {

	in := readInput()
	count := 0

	for _, p := range in {
		if valid(p) {
			count += 1
			fmt.Printf("%v - VALID\n", p)
		}
	}
	fmt.Println(count)
}
