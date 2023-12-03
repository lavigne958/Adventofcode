package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

const (
	gamePrefix = "Game "
	maxUint    = ^uint64(0)
)

func main() {
	aoc.Harness(run)
}

type cubes struct {
	red      uint64
	maxRed   uint64
	green    uint64
	maxGreen uint64
	bleu     uint64
	maxBleu  uint64
	game     uint64
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	nrColorRe := regexp.MustCompile(`^(\d+) (red|green|blue)$`)
	maxRed := uint64(12)
	maxGreen := uint64(13)
	maxBlue := uint64(14)

	winGames := []uint64{}
	powers := []uint64{}

	res := uint64(0)

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		// solve part 1 here
		cubes := cubes{}

		colon_index := strings.Index(line, ":")
		strGameId := line[len(gamePrefix):colon_index]
		cubes.game, _ = strconv.ParseUint(strGameId, 10, 64)

		winGames = append(winGames, cubes.game)

		colon_index++
		line := line[colon_index:]

		for _, draw := range strings.Split(line, ";") {
			cubes.bleu = 0
			cubes.green = 0
			cubes.red = 0

			draw = strings.TrimSpace(draw)

			for _, nrColor := range strings.Split(draw, ",") {
				nrColor = strings.TrimSpace(nrColor)

				matches := nrColorRe.FindStringSubmatch(nrColor)
				if matches == nil {
					fmt.Println("No color found in: ", nrColor)
				}

				if len(matches) < 3 {
					fmt.Println("not enough color found in: ", nrColor, "matches: ", matches)
				}

				nr, _ := strconv.ParseUint(matches[1], 10, 64)

				switch matches[2] {
				case "red":
					cubes.red = nr
					if nr > cubes.maxRed {
						cubes.maxRed = nr
					}
				case "green":
					cubes.green = nr
					if nr > cubes.maxGreen {
						cubes.maxGreen = nr
					}
				case "blue":
					cubes.bleu = nr
					if nr > cubes.maxBleu {
						cubes.maxBleu = nr
					}
				default:
					panic("no matching color for " + matches[2])
				}
			}

			// when you're ready to do part 2, remove this "not implemented" block
			if !part2 {
				if cubes.red > maxRed || cubes.green > maxGreen || cubes.bleu > maxBlue {

					position := slices.Index[[]uint64, uint64](winGames, cubes.game)
					if position != -1 {
						winGames = winGames[:position]
					}

				}
			}
		}

		if part2 {
			var val uint64 = 1
			if cubes.maxRed > 0 {
				val = val * cubes.maxRed
			}

			if cubes.maxGreen > 0 {
				val = val * cubes.maxGreen
			}

			if cubes.maxBleu > 0 {
				val = val * cubes.maxBleu
			}

			powers = append(powers, val)
		}
	}

	fmt.Println(winGames)

	if part2 {
		for _, val := range powers {
			res += val
		}
	} else {
		for _, game := range winGames {
			// fmt.Println("res: ", res, "game: ", game, "new res: ", res+game)
			res += game
		}
	}

	return res
}
