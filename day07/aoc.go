package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/dubbe/aoc25/helpers/grid"
)

//go:embed input.txt
var input string

var conv map[rune]int = map[rune]int{
	'.': 0,
	'S': 1,
	'^': 2,
	'|': 3,
}

func getSolutionPart1(input string) int64 {
	var result int64 = 0
	rows := strings.Split(input, "\n")

	g := grid.NewGrid(len(rows[0]), len(rows)-1)

	for x, row := range rows {
		for y, r := range row {
			if r != '.' {
				g.Set(x, y, conv[r])
			}
		}
	}

	for c := range g.Iterate() {
		if c.Value == 1 {
			fireBeam(c.Row+1, c.Col, g)
		}
	}

	for c := range g.Iterate() {
		above, _ := g.Get(c.Row-1, c.Col)
		if c.Value == 2 && above == 3 {
			result++
		}
	}

	return result
}

func fireBeam(row, col int, g *grid.Grid) {
	if row >= g.Height {
		return
	}
	g.Set(row, col, 3)
	below, _ := g.Get(row+1, col)
	if below == 0 {
		fireBeam(row+1, col, g)
	} else if below == 2 {
		fireBeam(row+1, col-1, g)
		fireBeam(row+1, col+1, g)
	}
}

func countPaths(row, col int, g *grid.Grid, memo map[grid.Coord]int64) int64 {
	if row == g.Height {
		return 1
	}

	if !g.IsIndexValid(row, col) {
		return 0
	}

	currentCoord := grid.Coord{Row: row, Col: col}
	if count, ok := memo[currentCoord]; ok {
		return count
	}

	var ways int64 = 0

	currentValue, _ := g.Get(row, col)

	switch currentValue {
	case 0:
		ways += countPaths(row+1, col, g, memo)
	case 2:
		ways += countPaths(row+1, col-1, g, memo)
		ways += countPaths(row+1, col+1, g, memo)
	case 1:
		ways += countPaths(row+1, col, g, memo)
	default:
		return 0
	}

	memo[currentCoord] = ways
	return ways
}

func getSolutionPart2(input string) int64 {
	var result int64 = 0

	rows := strings.Split(input, "\n")

	g := grid.NewGrid(len(rows[0]), len(rows)-1)

	for x, row := range rows {
		for y, r := range row {
			if r != '.' {
				g.Set(x, y, conv[r])
			}
		}
	}
	memo := make(map[grid.Coord]int64)
	for c := range g.Iterate() {
		if c.Value == 1 {
			ways := countPaths(c.Row+1, c.Col, g, memo)
			result += ways
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
