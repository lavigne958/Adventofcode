package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func Wins(line string) int {
	numberRe := regexp.MustCompile(`Card +\d+: ([\d ]+) \| ([\d ]+)$`)
	matches := numberRe.FindStringSubmatch(line)
	if matches == nil {
		panic("can't find values in: " + line)
	}

	if len(matches) < 3 {
		panic("can't find enough values in " + line)
	}

	winValuesStr := strings.TrimSpace(matches[1])
	myValuesStr := strings.TrimSpace(matches[2])

	var winValues = []uint64{}
	for _, val := range strings.Split(winValuesStr, " ") {
		if len(val) <= 0 {
			continue
		}

		x, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			panic("can't convert win value '" + val + "'")
		}

		winValues = append(winValues, x)
	}

	var myValues = []uint64{}
	for _, val := range strings.Split(myValuesStr, " ") {
		if len(val) <= 0 {
			continue
		}

		x, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			panic("can't convert my value '" + val + "'")
		}

		myValues = append(myValues, x)
	}

	var found = 0
	for _, val := range myValues {
		if slices.Contains[[]uint64, uint64](winValues, val) {
			found += 1
		}
	}

	return found
}

func solve2(input string) int {
	var res = 0
	lines := strings.Split(input, "\n")

	var cards []int = make([]int, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		cards[i] += 1

		nr_winners := Wins(line)

		for j := 1; j <= nr_winners; j++ {
			cards[i+j] += cards[i]
		}
	}

	for _, val := range cards {
		res += val
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
	var res = 0

	if part2 {
		return solve2(input)
	}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) <= 0 {
			continue
		}
	}

	return res
}
