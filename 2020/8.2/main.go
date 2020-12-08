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

func (c cpu) Terminated() bool {
	return c.pc >= len(c.mem)
}

func (c *cpu) Reset() {
	c.pc = 0
	c.acc = 0
	c.cycles = map[int]struct{}{}
}

func (c *cpu) SwapNopJmp(pc int) {
	op, arg := c.Decode(pc)
	switch op {
	case "jmp":
		op = "nop"
	case "nop":
		op = "jmp"
	}
	c.Encode(pc, op, arg)
}

func (c cpu) Decode(pc int) (string, int) {
	instr := c.mem[pc]
	i := strings.Split(instr, " ")
	op := i[0]
	arg, err := strconv.Atoi(i[1])
	if err != nil {
		panic(err)
	}
	return op, arg
}

func (c *cpu) Encode(pc int, op string, arg int) {
	c.mem[pc] = fmt.Sprintf("%s %d", op, arg)
}

func (c *cpu) Tick() {
	c.cycles[c.pc] = struct{}{}

	op, arg := c.Decode(c.pc)
	switch op {
	case "acc":
		c.acc += arg
		c.pc++
	case "nop":
		c.pc++
	case "jmp":
		c.pc += arg
	}
}

func (c *cpu) Run() {
	for !c.IsCycle() && !c.Terminated() {
		c.Tick()
	}
}

func (c cpu) String() string {
	return fmt.Sprintf("ACC:\t%d\nPC:\t%d",
		c.acc, c.pc)
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

	for pc := range c.mem {
		op, _ := c.Decode(pc)
		if op == "jmp" || op == "nop" {
			fmt.Printf("swapping %s at %d\n", op, pc)
			c.SwapNopJmp(pc)
			c.Run()
			if !c.Terminated() {
				fmt.Printf("not terminated, restoring\n")
				c.SwapNopJmp(pc)
				c.Reset()
			} else {
				fmt.Println(c)
				return
			}
		}
	}

}
