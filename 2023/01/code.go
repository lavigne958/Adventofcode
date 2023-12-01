package main

import (
	"fmt"
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
	var res int
	var values []string
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		earliestIndex := len(line)
		strToDigits := [][]string{
			{"zero", "0"},
			{"one", "1"},
			{"two", "2"},
			{"three", "3"},
			{"four", "4"},
			{"five", "5"},
			{"six", "6"},
			{"seven", "7"},
			{"eight", "8"},
			{"nine", "9"},
		}

		//fmt.Printf("before: '%s'\n", line)
		for range strToDigits {
			firstToReplace := -1

			for i, strDigit := range strToDigits {
				index := strings.Index(line, strDigit[0])
				if index != -1 {
					if index < earliestIndex {
						earliestIndex = index
						firstToReplace = i
					}
				}
			}

			if firstToReplace != -1 {
				digit := strToDigits[firstToReplace][1]
				strDigit := strToDigits[firstToReplace][0]
				new_value := fmt.Sprintf("%s%s%s", string(strDigit[0]), digit, string(strDigit[len(strDigit)-1]))
				line = strings.Replace(line, strToDigits[firstToReplace][0], new_value, 1)
			}

			earliestIndex = len(line)
		}
		//fmt.Printf("after: '%s'\n", line)

		b0 := byte('0')
		b9 := byte('9')
		var val1, val2 uint64

		for i := range line {

			if b0 <= line[i] && line[i] <= b9 {
				val1, _ = strconv.ParseUint(string(line[i]), 10, 64)
				val2 = val1

				break
			}
		}

		for i := len(line) - 1; i > 0; i-- {
			if b0 <= line[i] && line[i] <= b9 {
				val2, _ = strconv.ParseUint(string(line[i]), 10, 64)
				break
			}
		}

		values = append(values, fmt.Sprintf("%d%d", val1, val2))
		//fmt.Printf("values: '%v'\n", values)
		// solve part 1 here

	}

	for _, val := range values {
		intVal, _ := strconv.ParseInt(val, 10, 32)
		res += int(intVal)
	}

	return res
}
