package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	r := strings.NewReplacer(
		"F", "0",
		"L", "0",
		"B", "1",
		"R", "1",
	)
	for s.Scan() {
		t := s.Text()
		t = r.Replace(t)
		i, err := strconv.ParseInt(t, 2, 32)
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
	}
}
