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

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	groups := strings.Split(string(data), "\n\n")

	// Let's first extract the stacks into a map
	stacks := make(map[string][]string)
	entries := strings.Split(groups[0], "\n")
	for idx, num := range entries[len(entries)-1] {
		if num == ' ' {
			continue
		}
		// We have a number, let's find the stack above it
		vals := make([]string, 0)
		for i := len(entries) - 2; i >= 0; i-- {
			entry := entries[i]
			if idx < len(entry) {
				val := entry[idx]
				if val != ' ' {
					vals = append(vals, string(val))
				}
			}
		}
		stacks[string(num)] = vals
	}
	fmt.Println(stacks)

	// Next we process the actions
	actions := strings.Split(groups[1], "\n")
	actions = actions[:len(actions)-1]
	for _, action := range actions {
		comps := strings.Split(action, " ")

		count, err := strconv.Atoi(comps[1])
		check(err)

		orig_key := comps[3]
		orig := stacks[orig_key]

		dest_key := comps[5]
		dest := stacks[dest_key]

		for i := 0; i < count; i++ {
			popped := orig[len(orig)-1]
			orig = orig[:len(orig)-1]
			dest = append(dest, popped)
		}
		stacks[orig_key] = orig
		stacks[dest_key] = dest
	}
	fmt.Println(stacks)

	// Now walk through the stacks in order
	res := ""
	for i := 1; i <= len(stacks); i++ {
		val := stacks[fmt.Sprint(i)]
		if len(val) > 0 {
			res += val[len(val)-1]
		}
	}
	fmt.Println(res)
}
