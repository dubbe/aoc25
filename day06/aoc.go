package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func findNthIndex(sourceString, charSet string, n int) int {
	if n <= 0 {
		return -1
	}

	currentOffset := 0

	for iteration := 1; iteration <= n; iteration++ {
		matchIndexRelative := strings.IndexAny(sourceString[currentOffset:], charSet)

		if matchIndexRelative == -1 {
			return -1
		}

		absoluteIndex := currentOffset + matchIndexRelative

		if iteration == n {
			return absoluteIndex
		}

		currentOffset = absoluteIndex + 1
	}

	return -1
}

func getSolutionPart1(input string) int64 {
	var result int64 = 0

	ops := map[string]func(int64, int64) int64{
		"+": func(a, b int64) int64 { return a + b },
		"-": func(a, b int64) int64 { return a - b },
		"*": func(a, b int64) int64 { return a * b },
		"/": func(a, b int64) int64 { return a / b },
	}

	rows := strings.Split(input, "\n")

	operations := rows[len(rows)-2]
	window := findNthIndex(operations, "+*", 1)
	previousWindow := window
	n := 2
	exit := false
	for {
		window = findNthIndex(operations, "+*", n)

		if window == -1 {
			window = len(operations)
			fmt.Printf("break here,but need the last: '%s'\n", operations[previousWindow:])
			exit = true
		}

		operation := strings.TrimSpace(operations[previousWindow:window])

		var res int64 = 0

		for i, r := range rows {
			if i == len(rows)-2 {
				break
			}

			value, _ := strconv.ParseInt(strings.TrimSpace(r[previousWindow:window]), 10, 64)

			if i == 0 {
				res = value
			} else {
				res = ops[operation](res, value)
			}

		}

		result += res

		previousWindow = window
		n++

		if exit {
			break
		}
	}

	return result
}

func getSolutionPart2(input string) int64 {
	var result int64 = 0
	ops := map[string]func(int64, int64) int64{
		"+": func(a, b int64) int64 { return a + b },
		"-": func(a, b int64) int64 { return a - b },
		"*": func(a, b int64) int64 { return a * b },
		"/": func(a, b int64) int64 { return a / b },
	}

	rows := strings.Split(input, "\n")

	operations := rows[len(rows)-2]
	window := findNthIndex(operations, "+*", 1)
	previousWindow := window
	n := 2
	exit := false
	for {
		window = findNthIndex(operations, "+*", n)

		if window == -1 {
			window = len(operations)
			exit = true
		}

		operation := strings.TrimSpace(operations[previousWindow:window])

		var res int64 = 0

		numbers := []int64{}

		for y := window - 1; y >= previousWindow; y-- {
			number := ""
			for i, r := range rows {
				if i == len(rows)-2 {
					break
				}
				v := string(r[y])

				if v != " " {
					number += v
				}

			}
			numb, _ := strconv.ParseInt(number, 10, 64)
			if numb != 0 {
				numbers = append(numbers, numb)
			}
		}

		for i, n := range numbers {
			if i == 0 {
				res = n
			} else {
				if n != 0 {
					res = ops[operation](res, n)
				}
			}
		}

		result += res

		previousWindow = window
		n++

		if exit {
			break
		}
	}

	return result
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
