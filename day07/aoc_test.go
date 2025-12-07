package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string

func TestAOC_getSolutionPart1(t *testing.T) {
	var expectedSolution int64 = 21
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 6
	actualSolution := getSolutionPart2(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}
