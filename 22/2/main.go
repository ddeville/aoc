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

	res_1 := 0
	res_2 := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			vals := strings.Split(line, " ")

			// Meaning:
			//  A for Rock, B for Paper, and C for Scissors
			//  X for Rock, Y for Paper, and Z for Scissors
			// Points:
			//	1 for Rock, 2 for Paper, and 3 for Scissors
			//	0 if you lost, 3 if the round was a draw, and 6 if you won
			opponent, mine := vals[0], vals[1]

			switch mine {
			case "X":
				res_1 += 1
				switch opponent {
				case "A":
					res_1 += 3
				case "B":
					res_1 += 0
				case "C":
					res_1 += 6
				}
			case "Y":
				res_1 += 2
				switch opponent {
				case "A":
					res_1 += 6
				case "B":
					res_1 += 3
				case "C":
					res_1 += 0
				}
			case "Z":
				res_1 += 3
				switch opponent {
				case "A":
					res_1 += 0
				case "B":
					res_1 += 6
				case "C":
					res_1 += 3
				}
			}

			// Meaning:
			//  A for Rock, B for Paper, and C for Scissors
			//  X to lose, Y to draw, and Z to win
			// Points:
			//  1 for Rock, 2 for Paper, and 3 for Scissors
			//  0 if you lost, 3 if the round was a draw, and 6 if you won
			opponent, strategy := vals[0], vals[1]

			switch strategy {
			case "X":
				res_2 += 0
				switch opponent {
				case "A":
					res_2 += 3
				case "B":
					res_2 += 1
				case "C":
					res_2 += 2
				}
			case "Y":
				res_2 += 3
				switch opponent {
				case "A":
					res_2 += 1
				case "B":
					res_2 += 2
				case "C":
					res_2 += 3
				}
			case "Z":
				res_2 += 6
				switch opponent {
				case "A":
					res_2 += 2
				case "B":
					res_2 += 3
				case "C":
					res_2 += 1
				}
			}
		}
	}
	fmt.Println(res_1)
	fmt.Println(res_2)
}
