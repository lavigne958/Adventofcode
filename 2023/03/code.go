package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func isDigit(val byte) bool {
	return '0' <= val && val <= '9'
}

func searchTop(values [][]byte, i int, j int) uint64 {
	// find leftmost start of digit
	for j > 0 && isDigit(values[i][j]) {
		j -= 1
	}

	if j < 0 || !isDigit(values[i][j]) {
		j += 1
	}

	var strNumber string
	// start reading number
	for j < len(values[i]) && isDigit(values[i][j]) {
		strNumber = fmt.Sprintf("%s%c", strNumber, values[i][j])
		values[i][j] = '.'

		j += 1
	}

	number, err := strconv.ParseUint(strNumber, 10, 64)
	if err != nil {
		panic("top search cannot convert " + strNumber)
	}
	return number
}

func searchBottom(values [][]byte, i int, j int) uint64 {

	// find leftmost start of digi
	for j > 0 && isDigit(values[i][j]) {
		j -= 1
	}

	if j < 0 || !isDigit(values[i][j]) {
		j += 1
	}

	// start reading number
	var strNumber string
	for j < len(values[i]) && isDigit(values[i][j]) {
		strNumber = fmt.Sprintf("%s%c", strNumber, values[i][j])
		values[i][j] = '.'

		j += 1
	}
	number, err := strconv.ParseUint(strNumber, 10, 64)
	if err != nil {
		panic("bottom search cannot convert " + strNumber)
	}

	return number
}

func searchNumber(values [][]byte, i int, j int) uint64 {
	var numbers []uint64
	// search top
	if i-1 >= 0 {
		if j+1 <= len(values[i-1]) && isDigit(values[i-1][j+1]) {
			numbers = append(numbers, searchTop(values, i-1, j+1))
		}

		if isDigit(values[i-1][j]) {
			numbers = append(numbers, searchTop(values, i-1, j))
		}

		if j-1 >= 0 && isDigit(values[i-1][j-1]) {
			numbers = append(numbers, searchTop(values, i-1, j-1))
		}
	}

	// search left
	if j-1 >= 0 && isDigit(values[i][j-1]) {
		// build number in a string
		var strNumber string
		endStrNumber := j
		tmpJ := j - 1

		// find leftmost start of digit
		for tmpJ >= 0 && isDigit(values[i][tmpJ]) {
			tmpJ -= 1
		}

		if tmpJ < 0 || !isDigit(values[i][tmpJ]) {
			tmpJ += 1
		}

		strNumber = string(values[i][tmpJ:endStrNumber])

		for x := range values[i][tmpJ:endStrNumber] {
			values[i][x+tmpJ] = '.'
		}

		number, err := strconv.ParseUint(strNumber, 10, 64)
		if err != nil {
			panic("left search cannot convert " + strNumber)
		}

		numbers = append(numbers, number)
	}

	// search right
	if j+1 < len(values[i]) && isDigit(values[i][j+1]) {
		// build number in a string
		var strNumber string
		tmpJ := j + 1
		startStrNumber := tmpJ

		// find leftmost start of digit
		for tmpJ < len(values[i]) && isDigit(values[i][tmpJ]) {
			tmpJ += 1
		}

		strNumber = string(values[i][startStrNumber:tmpJ])

		number, err := strconv.ParseUint(strNumber, 10, 64)
		if err != nil {
			panic("right search cannot convert " + strNumber + " next 4 char " + string(values[i][startStrNumber:startStrNumber+3]))
		}

		for x := range values[i][startStrNumber:tmpJ] {
			values[i][x+startStrNumber] = '.'
		}
		numbers = append(numbers, number)
	}

	// search bottom
	if i+1 < len(values) {
		if j+1 < len(values[i+1]) && isDigit(values[i+1][j+1]) {
			numbers = append(numbers, searchBottom(values, i+1, j+1))
		}

		if isDigit(values[i+1][j]) {
			numbers = append(numbers, searchBottom(values, i+1, j))
		}

		if j-1 >= 0 && isDigit(values[i+1][j-1]) {
			numbers = append(numbers, searchBottom(values, i+1, j-1))
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	} else {
		return 0
	}
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
		return "not implemented"
	}

	firstReturn := strings.Index(input, "\n")

	if firstReturn < 0 {
		panic("could not find first \\n")
	}

	values := make([][]byte, 0)

	for i, line := range strings.Split(input, "\n") {
		values = append(values, []byte{})

		for j := range line {
			values[i] = append(values[i], line[j])
		}
	}

	var res uint64 = 0
	for i, line := range values {
		for j, char := range line {
			if char == '.' {
				continue
			}

			if isDigit(char) {
				continue
			}

			if char == '*' {
				res += searchNumber(values, i, j)
			}
		}
	}

	fmt.Println("After transform")

	path := "/tmp/result.txt"

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		panic("failed to open file" + err.Error())
	}
	defer file.Close()

	file.WriteString("     ")
	for j := range values[0] {
		// fmt.Printf("%3d", j)
		file.WriteString(fmt.Sprintf("%4d", j))
	}

	// fmt.Printf("\n")
	file.WriteString("\n")

	for i := range values {
		// fmt.Printf("[%3d]\t", i)
		file.WriteString(fmt.Sprintf("[%3d]", i))
		for j := range values[i] {
			char := values[i][j]
			if char == '.' {
				file.WriteString("    ")
				continue
			}
			file.WriteString(fmt.Sprintf(" %3c", char))
		}
		// fmt.Printf("\n")
		file.WriteString("\n")

	}

	return res
}
