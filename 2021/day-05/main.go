package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func plot(x, y int) {
	fmt.Println(x, y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Bresenham's algo from Wikipedia ftw
func plotLine(plotFn func(int, int), x0 int, y0 int, x1 int, y1 int) {
	dx := abs(x1 - x0)
	var sx int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	dy := -abs(y1 - y0)

	var sy int
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	err := dx + dy /* error value e_xy */
	for {
		plotFn(x0, y0)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy { /* e_xy+e_x > 0 */
			err += dy
			x0 += sx
		}
		if e2 <= dx { /* e_xy+e_y < 0 */
			err += dx
			y0 += sy
		}
	}
}

var inputRegex = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func parseInputLine(line string) (int, int, int, int) {
	line = strings.TrimSpace(line)
	m := inputRegex.FindStringSubmatch(line)

	x0, _ := strconv.Atoi(m[1])
	y0, _ := strconv.Atoi(m[2])
	x1, _ := strconv.Atoi(m[3])
	y1, _ := strconv.Atoi(m[4])

	return x0, y0, x1, y1
}

func getInput(in io.Reader) (lines []string) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

const MapSize = 1000

type Map struct {
	data   [MapSize][MapSize]int
	xbound int
	ybound int
}

func (m *Map) Put(x, y int) {
	m.data[y][x] += 1
	if m.xbound < x {
		m.xbound = x
	}
	if m.ybound < y {
		m.ybound = y
	}
}

func (m Map) String() string {
	b := &strings.Builder{}
	for y := 0; y <= m.ybound; y++ {
		for x := 0; x <= m.xbound; x++ {
			if m.data[y][x] == 0 {
				b.WriteRune('.')
			} else {
				b.WriteString(strconv.Itoa(m.data[y][x]))
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (m Map) CountOverlaps() (count int) {
	for y := 0; y <= m.ybound; y++ {
		for x := 0; x <= m.xbound; x++ {
			if m.data[y][x] >= 2 {
				count += 1
			}
		}
	}
	return
}

func plotFn(m *Map) func(int, int) {
	return func(x, y int) {
		m.Put(x, y)
	}
}

func main() {
	input := getInput(os.Stdin)

	// part 1
	var map1 Map
	plot1 := plotFn(&map1)

	// part 2
	var map2 Map
	plot2 := plotFn(&map2)

	for i := range input {
		x1, y1, x2, y2 := parseInputLine(input[i])
		plotLine(plot2, x1, y1, x2, y2)
		if !(x1 == x2 || y1 == y2) {
			continue
		}
		plotLine(plot1, x1, y1, x2, y2)
	}
	fmt.Println(map1.CountOverlaps())
	fmt.Println(map2.CountOverlaps())
}
