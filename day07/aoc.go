package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed input.txt
var input string

func getSolutionPart1(input string) int64 {
	var result int64 = 0

	return result
}

func getSolutionPart2(input string) int64 {
	var result int64 = 0

	return result
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
