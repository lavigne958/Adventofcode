package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func exo1(left, right []int64) int64 {
	res := int64(0)
	for i := range left {
		if left[i] > right[i] {
			res += left[i] - right[i]
		} else {
			res += right[i] - left[i]
		}
	}

	return res
}

func exo2(left, right []int64) int64 {
	res := int64(0)

	for _, l := range left {
		nr := 0

		for _, r := range right {
			if r == l {
				nr += 1
			}
		}

		res += l * int64(nr)
	}

	return res
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	left := []int64{}
	right := []int64{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {

		values := strings.Split(line, " ")

		strLeft := values[0]
		strRigtt := values[3]

		intLeft, _ := strconv.ParseInt(strLeft, 10, 64)
		left = append(left, intLeft)
		intRight, _ := strconv.ParseInt(strRigtt, 10, 64)
		right = append(right, intRight)
	}

	slices.Sort(left)
	slices.Sort(right)

	if part2 {
		return exo2(left, right)
	}

	return exo1(left, right)
}
