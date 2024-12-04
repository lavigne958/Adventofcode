package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var (
	valRe       = regexp.MustCompile(`\d{1,3}`)
	mulRe       = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	doRe        = regexp.MustCompile(`do\(\)`)
	secondMulRe = regexp.MustCompile(`do\(\).*mul\(\d{1,3},\d{1,3}\)`)
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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return exo2(input)
	}
	// solve part 1 here
	return exo1(input)
}

func exo2(input string) uint64 {
	res := uint64(0)

	first := mulRe.FindString(input)

	res += getMulValue(first)

	firstIndex := strings.Index(input, first) + len(first)
	trimmedInput := input[firstIndex+1:]

	trimmedInput = strings.ReplaceAll(trimmedInput, "\n", "")

	for len(trimmedInput) > 8 {
		nextDoLoc := doRe.FindStringIndex(trimmedInput)

		if nextDoLoc == nil {
			break
		}

		fmt.Printf("next do loc: %d:%d\n", nextDoLoc[0], nextDoLoc[1])

		trimmedInput = trimmedInput[nextDoLoc[1]:]

		nextMulLoc := mulRe.FindStringIndex(trimmedInput)
		if nextMulLoc == nil {
			break
		}

		nextMul := trimmedInput[nextMulLoc[0]:nextMulLoc[1]]
		fmt.Printf("next mul: %s\n", nextMul)

		res += getMulValue(nextMul)

		trimmedInput = trimmedInput[nextMulLoc[1]+1:]
	}

	return res
}

func getMulValue(mul string) uint64 {
	vals := valRe.FindAllString(mul, -1)

	left, _ := strconv.ParseUint(vals[0], 10, 64)
	right, _ := strconv.ParseUint(vals[1], 10, 64)

	return (left * right)
}

func exo1(input string) uint64 {

	muls := mulRe.FindAllString(input, -1)

	fmt.Printf("found %d results\n", len(muls))

	res := uint64(0)

	for _, mul := range muls {
		res += getMulValue(mul)
	}

	return res
}
