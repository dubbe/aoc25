package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type freshRange struct {
	from, to int
}

type freshRange64 struct {
	from, to, count int64
}

func getSolutionPart1(input string) int {
	result := 0

	freshRanges := []freshRange{}
	secondPart := false
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			secondPart = true
		}
		if !secondPart {
			r := strings.Split(row, "-")
			from, _ := strconv.Atoi(r[0])
			to, _ := strconv.Atoi(r[1])
			fr := freshRange{from, to}

			freshRanges = append(freshRanges, fr)

		} else {
			ingredient, _ := strconv.Atoi(row)

			for _, freshRange := range freshRanges {
				if ingredient >= freshRange.from && ingredient <= freshRange.to {

					result++
					break
				}
			}
		}
	}

	return result
}

func getSolutionPart2(input string) int64 {
	var result int64
	result = 0

	rows := strings.Split(input, "\n")

	freshIngredients := []freshRange64{}

	for _, row := range rows {
		if len(row) == 0 {
			break
		}

		r := strings.Split(row, "-")
		from, _ := strconv.ParseInt(r[0], 10, 64)
		to, _ := strconv.ParseInt(r[1], 10, 64)
		fr := freshRange64{from, to, to - from + 1}

		freshIngredients = append(freshIngredients, fr)
	}

	slices.SortFunc(freshIngredients, func(a, b freshRange64) int {
		return cmp.Compare(a.from, b.from)
	})

	freshIngredientsMerged := []freshRange64{}
	for _, fi := range freshIngredients {
		changed := false
		for i, fim := range freshIngredientsMerged {
			if fi.from <= fim.to+1 && fi.to >= fim.from {
				if fi.from < fim.from {
					freshIngredientsMerged[i].from = fi.from
				}

				if fi.to > fim.to {
					freshIngredientsMerged[i].to = fi.to
				}

				changed = true
				break
			}
		}
		if !changed {
			freshIngredientsMerged = append(freshIngredientsMerged, fi)
		}
	}

	for _, fim := range freshIngredientsMerged {
		result += fim.to - fim.from + 1
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
