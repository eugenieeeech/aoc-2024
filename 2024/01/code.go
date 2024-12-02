package main

import (
	"math"
	"sort"
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
	var leftCol []string
	var rightCol []string
	var distance int = 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		temp := strings.SplitN(line, "   ", 2)
		leftCol = append(leftCol, temp[0])
		rightCol = append(rightCol, temp[1])
	}

	sort.Slice(leftCol, func(i, j int) bool {
		return leftCol[i] < leftCol[j]
	})
	sort.Slice(rightCol, func(i, j int) bool {
		return rightCol[i] < rightCol[j]
	})

	for i := 0; i < len(leftCol); i++ {
		leftInt, _ := strconv.Atoi(leftCol[i])
		rightInt, _ := strconv.Atoi(rightCol[i])
		distance += int(math.Abs(float64(leftInt - rightInt)))
	}

	if part2 {
		// similarity score
		var score int = 0
		rightCountMap := make(map[int]int)
		for _, val := range rightCol {
			rightInt, _ := strconv.Atoi(val)
			rightCountMap[rightInt]++
		}

		for _, val := range leftCol {
			leftInt, _ := strconv.Atoi(val)
			score += leftInt * rightCountMap[leftInt]
		}

		return score
	}
	return distance
}
