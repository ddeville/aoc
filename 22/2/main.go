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

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	res := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			vals := strings.Split(line, " ")
			opponent, mine := vals[0], vals[1]

			// A for Rock, B for Paper, and C for Scissors
			// X for Rock, Y for Paper, and Z for Scissors
			// 1 for Rock, 2 for Paper, and 3 for Scissors
			// 0 if you lost, 3 if the round was a draw, and 6 if you won

			switch mine {
			case "X":
				res += 1
				switch opponent {
				case "A":
					res += 3
				case "B":
					res += 0
				case "C":
					res += 6
				}
			case "Y":
				res += 2
				switch opponent {
				case "A":
					res += 6
				case "B":
					res += 3
				case "C":
					res += 0
				}
			case "Z":
				res += 3
				switch opponent {
				case "A":
					res += 0
				case "B":
					res += 6
				case "C":
					res += 3
				}
			}
		}
	}
	fmt.Println(res)
}
