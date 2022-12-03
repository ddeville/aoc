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

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	score := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			left, right := line[:len(line)/2], line[len(line)/2:]
			items := make(map[rune]int)
			for _, item := range left {
				items[item] = 1
			}
			for _, item := range right {
				if _, ok := items[item]; ok {
					// Lowercase item types a through z have priorities 1 through 26.
					// Uppercase item types A through Z have priorities 27 through 52.
					if item >= a && item <= z {
						score += (int(item) - a + 1)
					} else if item >= A && item <= Z {
						score += (int(item) - A + 27)
					} else {
						panic("unknown value")
					}
					break
				}
			}
		}
	}

	fmt.Println(score)
}
