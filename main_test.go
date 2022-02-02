package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPossibleTimeFound(t *testing.T) {
	sampleDigits := []int{1, 2, 3, 4}
	res := possibleTimes(sampleDigits)

	t.Logf("Possible valid 24hr time : %d", res)

	assert.Equal(t, res, 10, "Wrong! the result must be > 0")
}

func TestValidPossibleTimeNotFound(t *testing.T) {
	sampleDigits := []int{10, 2, 3, 4}
	res := possibleTimes(sampleDigits)

	t.Logf("Possible valid 24hr time : %d", res)

	assert.Equal(t, res, 0, "Wrong! the result must be = 0")
}

func TestNonNegativeResponse(t *testing.T) {
	sampleDigits := []int{10, 2, 300, 0}
	res := possibleTimes(sampleDigits)

	t.Logf("Possible valid 24hr time : %d", res)

	assert.Equal(t, res, 0, "Wrong! the result must be > 0")
}
