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
	// TODO: use Map for storing the values
	var leftCol []string
	var rightCol []string
	var distance int = 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		temp := strings.SplitN(line, "   ", 2)
		leftCol = append(leftCol, temp[0])
		rightCol = append(rightCol, temp[1])
		// right := strings.SplitAfter(line, "   ")s

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
		/* Calculate a total similarity score by adding up each number
		in the left list after multiplying it by the number of times
		that number appears in the right list.
		*/
		var score int = 0
		var rightCount int = 0
		for i := 0; i < len(leftCol); i++ {
			leftInt, _ := strconv.Atoi(leftCol[i])
			rightCount = 0
			for j := 0; j < len(rightCol); j++ {
				rightInt, _ := strconv.Atoi(rightCol[j])
				if leftInt == rightInt {
					rightCount++
				}
			}
			score += leftInt * rightCount
		}

		return score
	}
	// solve part 1 here
	return distance
}
