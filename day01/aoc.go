package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"container/ring"
)

//go:embed input.txt
var input string


func getSolutionPart1(input string) int {
	result := 0

	r := ring.New(100)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	r = r.Move(50)
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			break
		}
		direction := row[0:1]
		if steps, err := strconv.Atoi(row[1:len(row)]); err == nil {
			if direction == "L" {
				steps = steps * -1
			}
			r = r.Move(steps)
			if(r.Value == 0) {
				result++
			}
		}
	}
	return result
}

func getSolutionPart2(input string) int {
	result := 0
	
	r := ring.New(100)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	r = r.Move(50)
	rows := strings.Split(input, "\n")
	previous := r.Value
	for _, row := range rows {
		if len(row) == 0 {
			break
		}
		direction := row[0:1]
		stringSteps :=  row[1:len(row)]
		multiple := 0
		if len(stringSteps) >= 3 {
			multiple, _ = strconv.Atoi(stringSteps[0:1])
			stringSteps = stringSteps[1:len(stringSteps)]
		}
		if steps, err := strconv.Atoi(stringSteps); err == nil {
			result += multiple
			if direction == "L" {
				steps = steps * -1
			}
			
			r = r.Move(steps)

			if r.Value == 0 {
				result++
			} else if previous != 0 && direction == "L" && r.Value.(int) > previous.(int) {
				result++
			} else if previous != 0 && direction == "R" && r.Value.(int) < previous.(int) {
				result++
			}

			previous = r.Value
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
