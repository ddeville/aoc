package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Range struct {
	start int
	end   int
}

func makeRange(value string) Range {
	r := strings.Split(value, "-")
	start, err := strconv.Atoi(r[0])
	check(err)
	end, err := strconv.Atoi(r[1])
	check(err)
	return Range{start, end}
}

func (r1 Range) Contains(r2 Range) bool {
	return r1.start <= r2.start && r1.end >= r2.end
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	count := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")

		r1 := makeRange(pairs[0])
		r2 := makeRange(pairs[1])
		if r1.Contains(r2) || r2.Contains(r1) {
			count += 1
		}
	}
	fmt.Println(count)
}
