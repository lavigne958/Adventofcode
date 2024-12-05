package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type Order struct {
	parents []uint64
}

var (
	orderRe = regexp.MustCompile(`(\d+)\|(\d+)`)
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

func exo1(inputs []string) uint64 {
	ordering := map[uint64]*Order{}

	splitLine := 0
	for i, line := range inputs {
		if strings.TrimSpace(line) == "" {
			splitLine = i
			break
		}

		strValues := orderRe.FindStringSubmatch(line)

		left, _ := strconv.ParseUint(strValues[1], 10, 64)
		right, _ := strconv.ParseUint(strValues[2], 10, 64)

		_, exists := ordering[right]
		if !exists {
			ordering[right] = &Order{
				parents: []uint64{},
			}
		}

		ordering[right].parents = append(ordering[right].parents, left)
	}

	for _, v := range ordering {
		slices.Sort(v.parents)
	}

	// for k, v := range ordering {
	// 	fmt.Printf("[%d] %v\n", k, v.parents)
	// }

	printLists := [][]uint64{}

	for _, line := range inputs[splitLine+1:] {
		values := strings.Split(line, ",")

		intValues := []uint64{}

		for _, val := range values {
			intVal, _ := strconv.ParseUint(val, 10, 64)
			intValues = append(intValues, intVal)
		}

		printLists = append(printLists, intValues)
	}

	validLists := [][]uint64{}

	for _, printList := range printLists {
		if isPrintListValid(printList, ordering) {
			validLists = append(validLists, printList)
		}
	}

	res := uint64(0)

	for _, valid := range validLists {
		if len(valid)%2 == 0 {
			fmt.Printf("cannot split %d in 1/2\n", len(valid))
		}
		half := len(valid) / 2

		fmt.Printf("%d + %d = %d\n", valid[half], res, res+valid[half])
		res += valid[half]
	}

	return res
}

func isPrintListValid(printList []uint64, orderging map[uint64]*Order) bool {
	for i, print := range printList {
		order, exist := orderging[print]

		if !exist {
			continue
		}

		for _, next := range printList[i+1:] {
			if slices.Contains(order.parents, next) {
				return false
			}
		}
	}

	return true
}
