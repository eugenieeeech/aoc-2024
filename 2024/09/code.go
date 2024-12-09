package main

import (
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
	result := 0
	// step 1: parse input
	var filesys []string
	for i, num := range strings.Split(input, "") {
		num, _ := strconv.Atoi(num)
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				filesys = append(filesys, strconv.Itoa(i/2))
			}
		} else {
			//append "."
			for k := 0; k < num; k++ {
				filesys = append(filesys, ".")
			}
		}
	}
	// step 2: move file
	for i := 0; i < len(filesys); i++ {
		if filesys[i] == "." {
			// move the last number to replace "."
			for j := len(filesys) - 1; j > i; j-- {
				if filesys[j] != "." {
					filesys[i] = filesys[j]
					filesys[j] = "."
					break
				}
			}
		}
	}

	// step 3: calculate checksum
	result = checksum(filesys)
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return result
}

func checksum(arr []string) int {
	sum := 0
	for i, j := range arr {
		if j != "." {
			j, _ := strconv.Atoi(j)
			sum += i * j
		} else {
			return sum
		}
	}
	return sum
}
