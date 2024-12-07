package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type operator int

type equation struct {
	result    uint64
	equation  []uint64
	operators []operator
}

const (
	ADD operator = iota
	MUL
	CAT
)

var (
	parseLineRe  = regexp.MustCompile(`(\d+): ([\d+ ]+)`)
	knownResults map[int]bool
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

	inputs := []equation{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		matches := parseLineRe.FindStringSubmatch(line)

		if len(matches) != 3 {
			panic("failed to identify 2 matching groups")
		}

		result, _ := strconv.ParseUint(matches[1], 10, 64)

		eq := equation{
			result: result,
		}

		for _, val := range strings.Split(matches[2], " ") {
			uintVal, _ := strconv.ParseUint(val, 10, 64)
			eq.equation = append(eq.equation, uintVal)
		}

		for i := 0; i < len(eq.equation)-1; i++ {
			eq.operators = append(eq.operators, ADD)
		}

		inputs = append(inputs, eq)
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return exo1(inputs)
}

func computeEquation(eq equation, operators []operator) uint64 {
	res := eq.equation[0]

	for i, val := range eq.equation[1:] {
		switch operators[i] {
		case ADD:
			res += val
		case MUL:
			res *= val
		}
	}

	return res
}

func computeEquationHash(operators []operator) int {
	res := 0

	for i, op := range operators {
		res += (i + 1) * int(op)
	}

	return res
}

func solvEquation(eq equation, operators []operator, index int) bool {
	//fmt.Printf("[%d] -> %v -> %v\n", eq.result, eq.equation, operators)

	operatorHash := computeEquationHash(operators)

	if res, exists := knownResults[operatorHash]; exists {
		if res {
			return true
		}
	} else {

		addRes := computeEquation(eq, operators)

		if addRes == eq.result {
			return true
		}

		knownResults[computeEquationHash(operators)] = false
	}

	operators[index] = MUL
	mulRes := computeEquation(eq, operators)

	if mulRes == eq.result {
		return true
	}

	knownResults[computeEquationHash(operators)] = false

	if index == len(operators)-1 {
		return false
	}

	if solvEquation(eq, slices.Clone(operators), index+1) {
		return true
	}

	// revert operator back to ADD
	operators[index] = ADD

	return solvEquation(eq, operators, index+1)
}

func exo1(equations []equation) uint64 {
	res := uint64(0)

	for _, eq := range equations {
		knownResults = map[int]bool{}

		if solvEquation(eq, eq.operators, 0) {
			fmt.Printf("[%d] -> %v success -> %v\n", eq.result, eq.equation, eq.operators)
			res += eq.result
		}
	}

	return res
}
