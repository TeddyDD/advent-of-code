package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type computer struct {
	op    int
	in1   int
	in2   int
	out   int
	pc    int
	debug bool
	mem   []int
}

func (c *computer) load(program string) error {
	str := strings.Split(program, ",")
	c.mem = make([]int, len(str))
	for i, s := range str {
		val, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		c.mem[i] = val
	}
	return nil
}

func (c *computer) setParams(noun, verb int) {
	c.mem[1] = noun
	c.mem[2] = verb
}

func (c *computer) setRegisters() {
	c.in1 = c.mem[c.pc+1]
	c.in2 = c.mem[c.pc+2]
	c.out = c.mem[c.pc+3]
	c.pc = c.pc + 4
}

func (c *computer) eval() (bool, error) {
	c.op = c.mem[c.pc]
	switch c.op {
	case 1:
		c.setRegisters()
		c.mem[c.out] = c.mem[c.in1] + c.mem[c.in2]
	case 2:
		c.setRegisters()
		c.mem[c.out] = c.mem[c.in1] * c.mem[c.in2]
	case 99:
		return true, nil
	default:
		return false, fmt.Errorf("Unknown opcode, state %+v", c)
	}
	return false, nil
}

func (c computer) run() error {
	for {
		if c.debug {
			fmt.Printf("%+v\n", c)
		}
		done, err := c.eval()
		if err != nil {
			return err
		}
		if done {
			break
		}
	}
	return nil
}

func main() {
	c := &computer{}
	b, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	program := strings.TrimSuffix(string(b), "\n")
	err = c.load(program)
	if err != nil {
		panic(err)
	}

	// fixing broken program state
	// https://adventofcode.com/2019/day/2
	c.setParams(12, 2)

	err = c.run()
	if err != nil {
		panic(err)
	}

	fmt.Println(c.mem[0])

	fmt.Println("Part 2")
	noun, verb, err := findPair(c, &program)
	if err != nil {
		panic(err)
	}
	fmt.Printf("noun: %d, verb: %d, anwser %d\n",
		noun, verb, 100*noun+verb)
}

func findPair(c *computer, program *string) (int, int, error) {
	c.debug = false
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			// fmt.Printf("pair %d %d...\n", noun, verb)
			c.load(*program) // won't fail, we did it before
			c.setParams(noun, verb)
			c.run()
			if c.mem[0] == 19690720 {
				return noun, verb, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("pair not found")
}
