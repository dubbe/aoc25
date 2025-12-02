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

func splitHalf(s string) (string, string) {
	runes := []rune(s)

	halfLength := len(runes) / 2

	firstHalf := string(runes[:halfLength])

	secondHalf := string(runes[halfLength:])

	return firstHalf, secondHalf
}

func allEqual[T comparable](s []T) bool {
	if len(s) <= 1 {
		return true
	}

	firstElement := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] != firstElement {
			return false
		}
	}

	return true
}

func chunkString(s string, chunkSize int) []string {
	if chunkSize <= 0 {
		return nil
	}

	runes := []rune(s)

	var chunks []string

	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}

		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}

func splitAndCheck(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		chunks := chunkString(s, i)
		if allEqual(chunks) {
			return true
		}
	}
	return false
}

func getSolutionPart1(input string) int {
	result := 0

	input = strings.TrimSuffix(input, "\n")
	ranges := strings.SplitSeq(input, ",")
	for r := range ranges {
		ra := strings.Split(r, "-")
		start, _ := strconv.Atoi(ra[0])
		stop, _ := strconv.Atoi(ra[1])

		for i := start; i <= stop; i++ {
			// fmt.Printf("i: %d\n", i)
			first, second := splitHalf(strconv.Itoa(i))
			if first == second {
				result += i
			}

		}

	}

	return result
}

func getSolutionPart2(input string) int {
	result := 0

	input = strings.TrimSuffix(input, "\n")
	ranges := strings.SplitSeq(input, ",")
	for r := range ranges {
		ra := strings.Split(r, "-")
		start, _ := strconv.Atoi(ra[0])
		stop, _ := strconv.Atoi(ra[1])

		for i := start; i <= stop; i++ {
			if splitAndCheck(strconv.Itoa(i)) {
				result += i
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
