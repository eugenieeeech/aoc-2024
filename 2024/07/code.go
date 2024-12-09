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
		if part2 {
			if match2(leftInt, numbers, 0, 0, true, true) || match2(leftInt, numbers, 0, 0, false, true) || match2(leftInt, numbers, 0, 0, true, false) || match2(leftInt, numbers, 0, 0, false, false) {
				sum += leftInt
			}
		} else {
			if match(leftInt, numbers, 0, 0, true) || match(leftInt, numbers, 0, 0, false) {
				sum += leftInt
			}
		}
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

func match2(target int, numbers []string, index int, current int, add bool, concatenation bool) bool {
	if index == len(numbers) {
		return current == target
	}
	num, _ := strconv.Atoi(numbers[index])
	if add {
		return match2(target, numbers, index+1, current+num, true, true) || match2(target, numbers, index+1, current+num, false, false) || match2(target, numbers, index+1, current+num, true, false) || match2(target, numbers, index+1, current+num, false, true)
	}
	if concatenation {
		concatenatedNum, _ := strconv.Atoi(strconv.Itoa(current) + numbers[index])
		return match2(target, numbers, index+1, concatenatedNum, true, true) || match2(target, numbers, index+1, concatenatedNum, false, false) || match2(target, numbers, index+1, concatenatedNum, true, false) || match2(target, numbers, index+1, concatenatedNum, false, true)
	}
	return match2(target, numbers, index+1, current*num, true, true) || match2(target, numbers, index+1, current*num, false, false) || match2(target, numbers, index+1, current*num, true, false) || match2(target, numbers, index+1, current*num, false, true)

}
