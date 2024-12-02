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
	var safeCount int
	// safe is true if the following conditions are both met:
	// The levels are either _all increasing_ or _all decreasing_.
	// Any two adjacent levels differ by _ats least one_ and _at most three_.

	for _, report := range strings.Split(strings.TrimSpace(input), "\n") {
		levels := strings.Split(report, " ")
		if isReportSafe(levels) {
			safeCount++
		}
	}

	if part2 {
		var part2Result int

		for _, report := range strings.Split(strings.TrimSpace(input), "\n") {
			levels := strings.Split(report, " ")
			safe := isReportSafe(levels)

			if !safe {
				for removePos := 0; removePos < 5; removePos++ {
					temp := append([]string{}, levels...)
					newReport := append(temp[:removePos], temp[removePos+1:]...)
					if isReportSafe(newReport) {
						safe = true
						break
					}
				}
			}

			if safe {
				part2Result++
			}
		}
		return part2Result
	}
	return safeCount
}

func isReportSafe(levels []string) bool {

	desc := false
	firstLevel, _ := strconv.Atoi(levels[0])
	secondLevel, _ := strconv.Atoi(levels[1])
	desc = firstLevel > secondLevel

	for pos := 0; pos < len(levels)-1; pos++ {
		levelInt, _ := strconv.Atoi(levels[pos])
		nextLevel, _ := strconv.Atoi(levels[pos+1])
		diff := levelInt - nextLevel

		if diff == 0 || (desc && (diff < 0 || diff > 3)) || (!desc && (diff > 0 || diff < -3)) {
			return false
		}
	}
	return true
}
