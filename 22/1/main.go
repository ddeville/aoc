package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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
	groups := strings.Split(string(data), "\n\n")

	elves := make([]int, 0)
	for _, group := range groups {
		cur := 0
		for _, val_str := range strings.Split(group, "\n") {
			if len(val_str) > 0 {
				val, err := strconv.Atoi(val_str)
				check(err)
				cur += val
			}
		}
		if cur != 0 {
			elves = append(elves, cur)
		}
	}

	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1])

	res := 0
	for _, val := range elves[len(elves)-3:] {
		res += val
	}
	fmt.Println(res)
}
