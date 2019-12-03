package main

// only 1st part cuz I'm tired today

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	X float64
	Y float64
}

func (p point) Add(other point) point {
	return point{p.X + other.X, p.Y + other.Y}
}

type set struct {
	list map[point]struct{}
}

func NewSet() *set {
	s := &set{}
	s.list = make(map[point]struct{})
	return s
}

func (s *set) Add(v point) {
	s.list[v] = struct{}{}
}

func (s set) Has(v point) bool {
	_, ok := s.list[v]
	return ok
}

func (s *set) Intersect(s2 *set) *set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

func distance(a, b point) float64 {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}

func offset(direction rune) point {
	switch direction {
	case 'U':
		return point{0, 1}
	case 'D':
		return point{0, -1}
	case 'R':
		return point{1, 0}
	case 'L':
		return point{-1, 0}
	}
	panic("unreachable")
}

func follow(path string) *set {
	res := NewSet()
	cur := point{0, 0}
	steps := strings.Split(path, ",")
	for _, s := range steps {
		num, _ := strconv.Atoi(s[1:])
		off := offset(rune(s[0]))
		for i := 0; i < num; i++ {
			cur = cur.Add(off)
			res.Add(cur)
		}
	}
	return res
}

func main() {
	byt, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	tasks := strings.Split(string(byt), "\n")
	a := follow(tasks[0])
	b := follow(tasks[1])
	c := a.Intersect(b)
	res := make([]float64, 1)
	for k, _ := range c.list {
		d := distance(point{0, 0}, k)
		res = append(res, d)
		fmt.Printf("%+v : %f\n", k, d)
	}
	sort.Sort(sort.Float64Slice(res))
	fmt.Println(res)
}
