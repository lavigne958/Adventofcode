package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	res := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		values := []uint64{}
		inputs := strings.Split(line, " ")

		for _, i := range inputs {
			val, _ := strconv.ParseUint(i, 10, 64)
			values = append(values, val)
		}

		if part2 {
			if exo2(values) {
				res += 1
			}
		} else {

			if exo1(values) {
				res += 1
			}
		}
	}
	return res
}

func delta(left, right uint64) int {
	return int(left - right)
}

func contious(left, right uint64) bool {
	return left > right
}

func isValid(left, right uint64) bool {
	return 1 <= delta(left, right) && delta(left, right) <= 3 && contious(left, right)
}

func exo1Decrease(input []uint64) bool {
	for i := 0; i < len(input)-1; i++ {
		if !isValid(input[i], input[i+1]) {
			return false
		}
	}

	return true
}

func exo1(input []uint64) bool {
	if input[0] > input[1] {
		return exo1Decrease(input)
	} else if input[0] < input[1] {
		copy := slices.Clone(input)
		slices.Reverse(copy)
		return exo1Decrease(copy)
	} else {
		return false

	}
}

func exo2(input []uint64) bool {
	if exo1(input) {
		return true
	}

	for i := 0; i < len(input); i++ {
		temp := append(append(make([]uint64, 0), input[:i]...), input[i+1:]...)
		fmt.Printf("[%d]=%d %v -> %v\n", i, input[i], input, temp)
		if exo1(temp) {
			fmt.Printf("valid  %d -> %v\n", input[i], input)
			return true
		}
	}

	return false
}
