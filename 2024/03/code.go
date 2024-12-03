package main

import (
	"regexp"
	"strconv"

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
	//Part 1 regex
	regex, _ := regexp.Compile(`mul[(][0-9]{1,3},[0-9]{1,3}[)]`)
	commands := regex.FindAllString(input, -1)
	result := 0
	for _, command := range commands {
		//Part 1
		//command: mul(X,Y))
		// extract X and Y

		// extract X and Y
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindStringSubmatch(command)
		if len(matches) == 3 {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			result += x * y
		}
	}
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return result
}
