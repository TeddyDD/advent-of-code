package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	pc  int
	mem []string
	acc int

	cycles map[int]struct{}
}

func (c cpu) IsCycle() bool {
	_, ok := c.cycles[c.pc]
	return ok
}

func (c *cpu) Tick() {
	c.cycles[c.pc] = struct{}{}
	instr := c.mem[c.pc]
	i := strings.Split(instr, " ")
	op := i[0]
	arg, err := strconv.Atoi(i[1])
	if err != nil {
		panic(err)
	}

	switch op {
	case "acc":
		c.acc += arg
		c.pc += 1
	case "nop":
		c.pc += 1
	case "jmp":
		c.pc += arg
	}
}

func (c *cpu) Run() {
	for !c.IsCycle() {
		c.Tick()
	}
}

func (c cpu) String() string {
	return fmt.Sprintf("ACC:\t%d\nPC:\t%d", c.acc, c.pc)
}

func NewCPU(in io.Reader) cpu {
	s := bufio.NewScanner(in)
	mem := []string{}
	for s.Scan() {
		mem = append(mem, s.Text())
	}

	res := cpu{
		pc:     0,
		mem:    mem,
		acc:    0,
		cycles: map[int]struct{}{},
	}
	return res
}

func main() {
	c := NewCPU(os.Stdin)
	c.Run()
	fmt.Println(c)
}
