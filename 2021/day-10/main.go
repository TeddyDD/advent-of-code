package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var score = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var matching = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
	'<': '>',
}

type stack []rune

func (s *stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *stack) Pop() (rune, error) {
	l := len(*s) - 1
	if l < 0 {
		return 0, errors.New("underflow")
	}
	r := (*s)[l]
	*s = (*s)[:len(*s)-1]
	return r, nil
}

func (s stack) String() string {
	b := &strings.Builder{}
	for _, c := range s {
		fmt.Fprint(b, string(c))
	}
	fmt.Fprintln(b)
	return b.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Part 1

	incomplete := []string{}

	sum := 0
outer:
	for _, line := range lines {
		stack := &stack{}
		for _, c := range line {
			switch c {
			case '(', '{', '<', '[':
				stack.Push(c)
			default:
				r, e := stack.Pop()
				if e != nil {
					fmt.Println(line)
					incomplete = append(incomplete, line)
					continue outer
				}
				r = matching[r]
				if r != c {
					sum += score[c]
					continue outer
				}
			}
		}
		incomplete = append(incomplete, line)
	}
	fmt.Println(sum)
}

func addScore(score int, char rune) int {
	m := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	return (score * 5) + m[char]
}
