package main

import (
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
	// when you're ready to do part 2, remove this "not implemented" block

	// Proccess the input to map
	arr := make([][]string, len(strings.Split(input, "\n")))
	for i := range arr {
		arr[i] = make([]string, len(strings.Split(input, "\n")[0]))
	}
	for col, row := range strings.Split(input, "\n") {
		temp := strings.Split(row, "")
		copy(arr[col], temp)
	}
	result := 0
	// Find XMAS in the map
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == "X" {
				// check adjacent seats
				//row case
				if j+3 < len(arr[i]) && arr[i][j+1] == "M" && arr[i][j+2] == "A" && arr[i][j+3] == "S" {
					result++
				}
				if j-3 >= 0 && arr[i][j-1] == "M" && arr[i][j-2] == "A" && arr[i][j-3] == "S" {
					result++
				}
				//col case
				if i+3 < len(arr) && arr[i+1][j] == "M" && arr[i+2][j] == "A" && arr[i+3][j] == "S" {
					result++
				}
				if i-3 >= 0 && arr[i-1][j] == "M" && arr[i-2][j] == "A" && arr[i-3][j] == "S" {
					result++
				}
				//diagonal case
				if i+3 < len(arr) && j+3 < len(arr[i]) && arr[i+1][j+1] == "M" && arr[i+2][j+2] == "A" && arr[i+3][j+3] == "S" {
					result++
				}
				if i-3 >= 0 && j-3 >= 0 && arr[i-1][j-1] == "M" && arr[i-2][j-2] == "A" && arr[i-3][j-3] == "S" {
					result++
				}
				if i+3 < len(arr) && j-3 >= 0 && arr[i+1][j-1] == "M" && arr[i+2][j-2] == "A" && arr[i+3][j-3] == "S" {
					result++
				}
				if i-3 >= 0 && j+3 < len(arr[i]) && arr[i-1][j+1] == "M" && arr[i-2][j+2] == "A" && arr[i-3][j+3] == "S" {
					result++
				}
			}
		}
	}

	if part2 {
		result = 0

		// Find crossed MAS in the map
		// X-MAS puzzle in which you're supposed to find two MAS in the shape of an X.
		for i := range arr {
			for j := range arr[i] {
				if arr[i][j] == "A" {
					leftCross := false
					rightCross := false
					// check adjacent seats
					if i-1 >= 0 && j+1 < len(arr[i]) && arr[i-1][j+1] == "M" {
						if i+1 < len(arr) && j-1 >= 0 && arr[i+1][j-1] == "S" {
							leftCross = true
						}
					}
					if i-1 >= 0 && j+1 < len(arr[i]) && arr[i-1][j+1] == "S" {
						if i+1 < len(arr) && j-1 >= 0 && arr[i+1][j-1] == "M" {
							leftCross = true
						}
					}
					if i-1 >= 0 && j-1 >= 0 && arr[i-1][j-1] == "M" {
						if i+1 < len(arr) && j+1 < len(arr[i]) && arr[i+1][j+1] == "S" {
							rightCross = true
						}
					}
					if i-1 >= 0 && j-1 >= 0 && arr[i-1][j-1] == "S" {
						if i+1 < len(arr) && j+1 < len(arr[i]) && arr[i+1][j+1] == "M" {
							rightCross = true
						}
					}
					if leftCross && rightCross {
						result++
					}
				}
			}
		}
		return result
	}
	// solve part 1 here
	return result
}
