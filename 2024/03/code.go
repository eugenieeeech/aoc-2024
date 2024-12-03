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
	regex, err := regexp.Compile(`mul[(][0-9]{1,3},[0-9]{1,3}[)]`)
	if err != nil {
		log.Fatal(err)
	}
	commands := regex.FindAllString(input, -1)
	result := calculateResult(commands)
	removed := 0
	if part2 {
		part2Result := 0
		newInput := input
		// find all string between don't() and do()
		part2Regex, err := regexp.Compile(`don't\(\)(.*?)do\(\)`)
		if err != nil {
			log.Fatal(err)
		}

		for _, part2RegexCommands := range part2Regex.FindAllString(input, -1) {
			removedCommands := regex.FindAllString(part2RegexCommands, -1)
			if len(removedCommands) > 0 {
				removed += calculateResult(removedCommands)
			}
			//newInput = strings.ReplaceAll(newInput, part2RegexCommands, "")

		}
		println("re:", removed)

		newCommands := regex.FindAllString(newInput, -1)
		println(187194524 - removed)
		part2Result = calculateResult(newCommands)
		return part2Result
	}
	return result
}

func calculateResult(commands []string) int {
	result := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, command := range commands {
		// println(command)
		matches := re.FindStringSubmatch(command)
		if len(matches) == 3 {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			result += x * y
		}
	}
	return result
}
