package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string

func createGrid(input string) *Grid {
	rows := strings.Split(input, "\n")
	grid := NewGrid(len(rows[0]), len(rows)-1)

	for y, row := range rows {
		for x, v := range row {
			if string(v) == "@" {
				grid.Set(y, x, 1)
			}
		}
	}

	grid.Print()
	return grid
}

func getSolutionPart1(input string) int {
	result := 0

	grid := createGrid(input)

	for p := range grid.Iterate() {
		adjacent := grid.GetAdjacent(p.Row, p.Col)
		if len(adjacent) < 4 {
			grid.Set(p.Row, p.Col, 2)
			result++
		}
	}

	grid.Print()

	return result
}

func getSolutionPart2(input string) int {
	result := 0

	grid := createGrid(input)

	for {
		rollsRemoved := 0
		for p := range grid.Iterate() {

			adjacent := grid.GetAdjacent(p.Row, p.Col)
			if len(adjacent) < 4 {
				grid.Set(p.Row, p.Col, 2)
				rollsRemoved++
			}
		}
		fmt.Printf("Removed %d rolls of paper: \n", rollsRemoved)
		grid.Print()
		if rollsRemoved == 0 {
			break
		}
		result += rollsRemoved

		for p := range grid.Iterate() {
			roll, _ := grid.Get(p.Row, p.Col)
			if roll == 2 {
				grid.Set(p.Row, p.Col, 0)
			}
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
