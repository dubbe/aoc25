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

func getSolutionPart1(input string) int {
	result := 0

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			break
		}
		highest := 0
		secondHighest := 0
		for i, r := range row {
			digit, _ := strconv.Atoi(string(r))
			if i < len(row)-1 && digit > highest {
				highest = digit
				secondHighest = 0
			} else if digit > secondHighest {
				secondHighest = digit
			}
		}
		number, _ := strconv.Atoi(fmt.Sprintf("%d%d", highest, secondHighest))
		result += number
	}

	return result
}

type value struct {
	data     int
	position int
}

func getSolutionPart2(input string) int {
	result := 0
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			break
		}

		line := row
		sequence := ""

		window := len(row) - 11

		for i := 0; i < len(row); i++ {

			if len(line) < window {
				break
			}
			possible := line[0:window]

			heighest := 0
			lastPosition := 0
			for i, r := range possible {
				nr, _ := strconv.Atoi(string(r))
				if nr > heighest {
					heighest = nr
					lastPosition = i
				}
				if heighest == 9 {
					break
				}

			}

			window = window - lastPosition
			line = line[lastPosition+1:]
			sequence += fmt.Sprintf("%d", heighest)

		}

		number, _ := strconv.Atoi(sequence)
		result += number

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
