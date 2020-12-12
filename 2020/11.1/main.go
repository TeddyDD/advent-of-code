package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"
	"strings"
)

type ferry struct {
	arr  []byte
	w, h int
}

func (f ferry) Sum() uint32 {
	return crc32.ChecksumIEEE(f.arr)
}

func (f ferry) get(x, y int) (byte, error) {
	idx := y*f.w + x
	if idx < 0 || idx > len(f.arr) {
		return 0, fmt.Errorf("out of bounds")
	}
	return f.arr[idx], nil
}

func (f *ferry) set(x, y int, ch byte) {
	f.arr[y*f.w+x] = ch
}

func (f ferry) adjacent(x, y int) (occupied, empty int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			checkX := x + i
			checkY := y + j
			if !(i == 0 && j == 0) &&
				checkX >= 0 &&
				checkY >= 0 &&
				checkX < f.w &&
				checkY < f.h {
				c, err := f.get(x+i, y+j)
				if err != nil {
					continue
				}
				switch c {
				case 'L':
					empty++
				case '#':
					occupied++
				}
			}
		}
	}
	return
}

func (f ferry) Count() (occupied int) {
	for _, i := range f.arr {
		if i == '#' {
			occupied++
		}
	}
	return
}

func (f ferry) Next() ferry {
	n := &ferry{
		arr: make([]byte, len(f.arr)),
		w:   f.w,
		h:   f.h,
	}
	copy(n.arr, f.arr)

	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			o, _ := f.adjacent(x, y)
			this, _ := f.get(x, y)
			switch this {
			case 'L':
				if o == 0 {
					n.set(x, y, '#')
				}
			case '#':
				if o >= 4 {
					n.set(x, y, 'L')
				}
				// default:
				// 	n.set(x, y, this)
			}
		}
	}
	return *n
}

func (f ferry) String() string {
	b := &strings.Builder{}
	fmt.Fprintln(b)
	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			c, _ := f.get(x, y)
			fmt.Fprintf(b, "%s", string(c))
		}
		fmt.Fprintf(b, "\n")
	}
	fmt.Fprintln(b)
	return b.String()
}

func readInput() ferry {
	f := ferry{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		b := s.Bytes()
		if f.w == 0 {
			f.w = len(b)
		}
		f.arr = append(f.arr, b...)
		f.h++
	}
	return f
}

func main() {
	f := readInput()
	prev := f.Sum()
	var newSum uint32
	for {
		n := f.Next()
		newSum = n.Sum()
		if newSum == prev {
			break
		}
		prev = newSum
		f = n
	}
	fmt.Println(f)
	fmt.Println(f.Count())
}
