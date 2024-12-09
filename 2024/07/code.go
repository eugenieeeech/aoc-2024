package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0
	for _, eq := range strings.Split(strings.TrimSpace(input), "\n") {
		temp := strings.Split(eq, ":")
		left := temp[0]
		leftInt, err := strconv.Atoi(left)
		if err != nil {
			continue
		}
		numbers := strings.Split(strings.TrimSpace(temp[1]), " ")
		if match(leftInt, numbers, 0, 0, true) || match(leftInt, numbers, 0, 0, false) {
			sum += leftInt
		}
	}

	if part2 {
		return "not implemented"
	}
	return sum
}

func match(target int, numbers []string, index int, current int, add bool) bool {
	if index == len(numbers) {
		return current == target
	}
	num, _ := strconv.Atoi(numbers[index])
	if add {
		return match(target, numbers, index+1, current+num, true) || match(target, numbers, index+1, current+num, false)
	}
	return match(target, numbers, index+1, current*num, true) || match(target, numbers, index+1, current*num, false)
}
