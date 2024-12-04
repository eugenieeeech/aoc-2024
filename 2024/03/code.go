package main

import (
	"log"
	"regexp"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	//Part 1 regex
	result := 0
	regex, err := regexp.Compile(`mul[(][0-9]{1,3},[0-9]{1,3}[)]`)
	if err != nil {
		log.Fatal(err)
	}
	commands := regex.FindAllString(input, -1)
	for _, command := range commands {
		result += calculateResult(command)
	}

	if part2 {
		part2Result := 0
		part2Regex, err := regexp.Compile(`don't\(\)|do\(\)|mul\(\d{1,3},\d{1,3}\)`)
		if err != nil {
			log.Fatal(err)
		}
		isEnabled := true
		for _, part2RegexCommands := range part2Regex.FindAllString(input, -1) {

			if part2RegexCommands == "don't()" {
				isEnabled = false
			} else if part2RegexCommands == "do()" {
				isEnabled = true
			} else if isEnabled {

				part2Result += calculateResult(part2RegexCommands)
			}
		}
		return part2Result
	}
	return result
}

func calculateResult(command string) int {
	result := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindStringSubmatch(command)
	if len(matches) == 3 {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		result += x * y
	}
	return result
}
