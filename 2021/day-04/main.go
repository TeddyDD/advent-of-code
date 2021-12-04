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

const BoardSize = 5

type bingoBoard struct {
	numbers    [BoardSize][BoardSize]int
	marks      [BoardSize][BoardSize]bool
	alreadyWon bool
}

func NewBoard() *bingoBoard {
	return &bingoBoard{}
}

func (b *bingoBoard) String() string {
	s := &strings.Builder{}
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			if b.marks[y][x] {
				s.WriteString("✔️")
			} else {
				s.WriteString(" ")
			}
			fmt.Fprintf(s, "%2d ", b.numbers[y][x])
		}
		s.WriteRune('\n')
	}
	return s.String()
}

var space = regexp.MustCompile(`\s+`)

func (b *bingoBoard) LoadRow(row int, input string) {
	input = space.ReplaceAllString(input, " ")
	input = strings.TrimSpace(input)

	n := LoadNums(input, " ")
	for i := 0; i < BoardSize; i++ {
		b.numbers[row][i] = n[i]
	}
}

func (b *bingoBoard) Mark(draw int) {
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			if b.numbers[y][x] == draw {
				b.marks[y][x] = true
			}
		}
	}
}

func (b *bingoBoard) SumUnmarked() int {
	sum := 0
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			if !b.marks[y][x] {
				sum += b.numbers[y][x]
			}
		}
	}
	return sum
}

func (b *bingoBoard) checkRow(row int) bool {
	for i := 0; i < BoardSize; i++ {
		if !b.marks[row][i] {
			return false
		}
	}
	return true
}

func (b *bingoBoard) checkColumn(column int) bool {
	for i := 0; i < BoardSize; i++ {
		if !b.marks[i][column] {
			return false
		}
	}
	return true
}

func (b *bingoBoard) AlreadyWon() bool {
	return b.alreadyWon
}

func (b *bingoBoard) CheckIfWon() bool {
	for y := 0; y < BoardSize; y++ {
		if b.checkRow(y) {
			b.alreadyWon = true
			return true
		}
	}
	for x := 0; x < BoardSize; x++ {
		if b.checkColumn(x) {
			b.alreadyWon = true
			return true
		}
	}
	return false
}

func LoadNums(in, sep string) []int {
	nums := strings.Split(in, sep)
	res := make([]int, len(nums))
	for i := range nums {
		j, _ := strconv.Atoi(nums[i])
		res[i] = j
	}
	return res
}

func getInput(in io.Reader) (draws []int, boards []*bingoBoard) {
	scanner := bufio.NewScanner(in)

	scanner.Scan()
	draws = LoadNums(scanner.Text(), ",")
	scanner.Scan() // empty line

	for {
		b := NewBoard()

		for i := 0; i < 5; i++ {
			scanner.Scan()
			l := scanner.Text()
			b.LoadRow(i, l)
		}
		boards = append(boards, b)

		ok := scanner.Scan() // empty line
		if !ok {
			break
		}
	}
	return
}

type result struct {
	*bingoBoard
	draw int
}

func (r result) String() string {
	return fmt.Sprintf("%s\n%d", r.bingoBoard, r.SumUnmarked()*r.draw)
}

func main() {
	draws, boards := getInput(os.Stdin)

	// run game

	var first result
	var last result

	for _, draw := range draws {
		for i := range boards {
			if boards[i].AlreadyWon() {
				continue
			}
			boards[i].Mark(draw)
			if boards[i].CheckIfWon() {
				if first.bingoBoard == nil {
					first.bingoBoard = boards[i]
					first.draw = draw
				}
				last.bingoBoard = boards[i]
				last.draw = draw
			}
		}
	}

	fmt.Println(first)
	fmt.Println(last)
}
