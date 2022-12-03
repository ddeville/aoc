package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const a = 97
const z = 122
const A = 65
const Z = 90

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func characterValue(item rune) int {
	if item >= a && item <= z {
		return (int(item) - a + 1)
	} else if item >= A && item <= Z {
		return (int(item) - A + 27)
	} else {
		panic("unknown value")
	}
}

func firstPart(lines []string) {
	score := 0
	for _, line := range lines {
		left, right := line[:len(line)/2], line[len(line)/2:]
		items := make(map[rune]int)
		for _, item := range left {
			items[item] = 1
		}
		for _, item := range right {
			if _, ok := items[item]; ok {
				score += characterValue(item)
				break
			}
		}
	}
	fmt.Println(score)
}

func secondPart(lines []string) {
	score := 0
	for idx := 0; idx < len(lines); idx += 3 {
		searchMap := make(map[rune]int)
		resMap := make(map[rune]int)

		// First pass through the first list to build the lookup table
		for _, item := range lines[idx] {
			searchMap[item] += 1
		}

		// Second pass to match the duplicates
		for _, item := range lines[idx+1] {
			if _, ok := searchMap[item]; ok {
				resMap[item] += 1
			}
		}
		searchMap = resMap
		resMap = make(map[rune]int)

		// Third pass to match duplicates against the lookup table from the second pass
		for _, item := range lines[idx+2] {
			if _, ok := searchMap[item]; ok {
				resMap[item] += 1
			}
		}

		if len(resMap) != 1 {
			panic("bad input")
		}
		var item rune
		for key := range resMap {
			item = key
		}

		score += characterValue(item)
	}
	fmt.Println(score)
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	firstPart(lines)
	secondPart(lines)
}
