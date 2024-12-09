package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func printDiskMap(diskMap []int64) {
	for _, val := range diskMap {
		if val == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", val)
		}
	}
	fmt.Println()
}

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
	// if len(input) > 4096 {
	// 	return 43
	// }

	diskMap := []int64{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		strBlocks := strings.Split(line, "")
		blockId := int64(0)

		for pos, strBlock := range strBlocks {
			nrBlocks, _ := strconv.ParseInt(strBlock, 10, 64)
			if pos%2 == 0 {
				for i := 0; i < int(nrBlocks); i++ {
					diskMap = append(diskMap, int64(blockId))
				}
				blockId++
			} else {
				for i := 0; i < int(nrBlocks); i++ {
					diskMap = append(diskMap, int64(-1))
				}
			}
		}
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		printDiskMap(diskMap)
		exo2(diskMap)

		printDiskMap(diskMap)
		return computeCheckSum(diskMap)
	}
	// solve part 1 here
	exo1(diskMap)
	printDiskMap(diskMap)

	return computeCheckSum(diskMap)
}

func nextFreeSpace(diskMap []int64) int {
	for i, val := range diskMap {
		if val == -1 {
			return i
		}
	}

	return len(diskMap)
}

func nextFreeSpaceFit(diskMap []int64, size int) int {
	for i, val := range diskMap {
		if val == -1 {
			if i+size >= len(diskMap) {
				return len(diskMap)
			}
			left := size
			for left > 0 && diskMap[i+(size-left)] == -1 {
				left--
			}

			if left == 0 {
				return i
			}
		}
	}

	return len(diskMap)
}

func swapValues(diskMap []int64, from, to int) {
	old := diskMap[to]
	diskMap[to] = diskMap[from]
	diskMap[from] = old
}

func computeCheckSum(diskMap []int64) (res int64) {
	for i, val := range diskMap {
		if val == -1 {
			continue
		}

		res += int64(i) * val
	}

	return
}

func findStartBlock(diskMap []int64, end int) (i int) {

	for i = end; diskMap[i] == diskMap[end]; i-- {
	}

	i++
	return
}

func exo2(diskMap []int64) {
	index := len(diskMap) - 1
	nextFreeSpaceBigEnough := 0

	for index > 1 {
		if diskMap[index] != -1 {
			startBlock := findStartBlock(diskMap, index)
			blockLen := index - startBlock + 1

			nextFreeSpaceBigEnough = nextFreeSpaceFit(diskMap, blockLen)

			if nextFreeSpaceBigEnough >= len(diskMap) {
				index -= blockLen
				continue
			}

			if nextFreeSpaceBigEnough > index {
				index -= blockLen
				continue
			}

			for i := 0; i < blockLen; i++ {
				swapValues(diskMap, startBlock+i, nextFreeSpaceBigEnough+i)
			}
			fmt.Printf("%d: ", index)
		}

		index--
	}
}

func exo1(diskMap []int64) {
	index := len(diskMap) - 1
	nextFree := 0

	for nextFree < index {
		if diskMap[index] != -1 {
			nextFree = nextFreeSpace(diskMap)
			if nextFree > index {
				break
			}
			swapValues(diskMap, index, nextFree)
		}

		index--
	}
}
