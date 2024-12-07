package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type DIRECTION int

const (
	UP DIRECTION = iota
	RIGHT
	DOWN
	LEFT
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
	area := [][]byte{}
	guardX, guardY := 0, 0

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		area = append(area, make([]byte, len(line)))
		for x := range line {
			if line[x] == '^' {
				guardX = x
				guardY = y

				area[y][x] = '.'

				continue
			}

			area[y][x] = line[x]
		}
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return exo1(area, guardX, guardY)
}

func printMap(area [][]byte) {
	for y := range area {
		for x := range area {
			fmt.Printf("%c", area[y][x])
		}

		fmt.Println()
	}
}

func canIMove(area [][]byte, x, y int) bool {
	return area[y][x] != '#'
}

func isItExit(area [][]byte, x, y int) bool {
	return x < 0 || y < 0 || x >= len(area[0]) || y >= len(area)
}

func nextGuardPosition(x, y int, guardDirection DIRECTION) (int, int) {
	switch guardDirection {
	case UP:
		return x, y - 1
	case RIGHT:
		return x + 1, y
	case DOWN:
		return x, y + 1
	case LEFT:
		return x - 1, y
	}

	panic("bad direction !")
}

func nextGuardDirection(guardDirection DIRECTION) DIRECTION {
	return DIRECTION((guardDirection + 1) % 4)
}

func countUniqSpot(area [][]byte) int {
	res := 0

	for i := range area {
		for j := range area[i] {
			if area[i][j] == 'X' {
				res++
			}
		}
	}

	return res
}

func exo1(area [][]byte, guardX, guardY int) int {
	guardDirection := UP
	nextX, nextY := nextGuardPosition(guardX, guardY, guardDirection)

	for !isItExit(area, nextX, nextY) {
		// fmt.Printf("[%d][%d] - %d\n", guardX, guardY, guardDirection)
		if canIMove(area, nextX, nextY) {
			// move forward
			area[guardY][guardX] = 'X'
			guardX = nextX
			guardY = nextY
			// fmt.Printf("MOVE - next pos [%d][%d] - %d\n", guardX, guardY, guardDirection)
			nextX, nextY = nextGuardPosition(guardX, guardY, guardDirection)
		} else {
			// obstcle
			guardDirection = nextGuardDirection(guardDirection)
			nextX, nextY = nextGuardPosition(guardX, guardY, guardDirection)
			// fmt.Printf("OBSTACLE - STAY [%d][%d] - %d\n", guardX, guardY, guardDirection)
		}
	}

	area[guardY][guardX] = 'X'

	printMap(area)

	return countUniqSpot(area)
}
