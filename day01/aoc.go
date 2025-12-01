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
	fmt.Printf("r: %d \n", r.Value)
	r = r.Move(50)
	fmt.Printf("r: %d \n", r.Value)
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
			fmt.Printf("direction: %s, steps: %d, pointAt: %d \n", direction, steps, r.Value)
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
	fmt.Printf("r: %d \n", r.Value)
	r = r.Move(50)
	fmt.Printf("r: %d \n", r.Value)
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
				fmt.Printf("LANDED ON ZERO\n")
				result++
			} else if previous != 0 && direction == "L" && r.Value.(int) > previous.(int) {
				fmt.Printf("MOVED PASS ZERO\n")
				result++
			} else if previous != 0 && direction == "R" && r.Value.(int) < previous.(int) {
				fmt.Printf("MOVED PASS ZERO\n")
				result++
			}

			if steps == 0 {
				fmt.Printf("FÖRBANNAT! row: %s \n", row)
			}
			//fmt.Printf("direction: %s, previous: %d, steps: %d, pointAt: %d, multiple: %d \n", direction, previous, steps, r.Value, multiple)
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
