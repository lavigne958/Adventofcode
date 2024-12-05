package main

import (
	"regexp"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var (
	xmasRe      = regexp.MustCompile(`XMAS`)
	xmasReverRe = regexp.MustCompile(`SAMX`)
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
	inputs := strings.Split(strings.TrimSpace(input), "\n")
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return exo1(inputs)
}

func rotateStrs(input []string) []string {
	res := []string{}

	for i := 0; i < len(input[0]); i++ {
		if len(res) < i+1 {
			res = append(res, "")
		}

		for j := 0; j < len(input); j++ {
			res[i] = res[i] + string(input[j][i])
		}
	}

	return res
}

func findXmasLines(inputs []string) int {
	res := 0

	for _, line := range inputs {
		res += len(xmasRe.FindAllString(line, -1))

		res += len(xmasReverRe.FindAllString(line, -1))
	}

	return res
}

func isM(inputs []string, i, j int) bool {
	return inputs[i][j] == 'M'
}

func isA(inputs []string, i, j int) bool {
	return inputs[i][j] == 'A'
}

func isS(inputs []string, i, j int) bool {
	return inputs[i][j] == 'S'
}

func findDiagonales(inputs []string) int {
	res := 0

	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			if inputs[i][j] == 'X' {
				// found a start

				// check upper left diagonal
				if i >= 3 && j >= 3 {
					if isM(inputs, i-1, j-1) && isA(inputs, i-2, j-2) && isS(inputs, i-3, j-3) {
						res++
					}
				}

				// check upper right diagonal
				if i >= 3 && j <= len(inputs[i])-4 {
					if isM(inputs, i-1, j+1) && isA(inputs, i-2, j+2) && isS(inputs, i-3, j+3) {
						res++
					}
				}

				// check lower right diagonal
				if i <= len(inputs)-4 && j <= len(inputs[i])-4 {
					if isM(inputs, i+1, j+1) && isA(inputs, i+2, j+2) && isS(inputs, i+3, j+3) {
						res++
					}
				}

				// check lower left diagonal
				if i <= len(inputs)-4 && j >= 3 {
					if isM(inputs, i+1, j-1) && isA(inputs, i+2, j-2) && isS(inputs, i+3, j-3) {
						res++
					}
				}
			}
		}
	}

	return res
}

func exo1(inputs []string) int {
	res := findXmasLines(inputs)

	rotatedInput := rotateStrs(inputs)

	res += findXmasLines(rotatedInput)

	res += findDiagonales(inputs)

	return res
}
