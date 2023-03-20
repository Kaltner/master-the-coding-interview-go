package main

import (
	"fmt"
	"testing"
)

func TestMergeAndSortSimple(t *testing.T) {
	// Are the inputs going to be always sorted?]
	// Could the arrays be the same?
	// Are the inputs always going to be numbers?
	// If they are not numbers, how is expected then to be sorted?
	// Could the numbers repeat on each side?

	// Test case structures:
	// One two dimensional array for both the inputs
	// One for the outputs

	// Cases of tests:
	// Simple test
	// Intercalated values
	// Negative values
	// First has all the biggest numbers
	// One of the arrays empty
	cases := []struct {
		inputs         [][]int
		expectedOutput []int
	}{
		{
			inputs:         [][]int{{0, 2, 4}, {}},
			expectedOutput: []int{0, 2, 4},
		},
		{
			inputs:         [][]int{{}, {1, 3, 5}},
			expectedOutput: []int{1, 3, 5},
		},
		{
			inputs:         [][]int{{0, 2, 4, 6, 8, 10}, {1, 3, 5, 7, 9, 11}},
			expectedOutput: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			inputs:         [][]int{{0, 2, 4, 6, 8, 10}, {0, 2, 4, 6, 8, 10}},
			expectedOutput: []int{0, 2, 4, 6, 8, 10},
		},
		{
			inputs:         [][]int{{0, 4, 6, 8, 10}, {0, 2, 4, 6, 8, 10}},
			expectedOutput: []int{0, 2, 4, 6, 8, 10},
		},
		// {
		// inputs:         [][]int{{0, 0, 2, 4}, {1, 2, 3}},
		// expectedOutput: []int{0, 1, 2, 3, 4},
		// },
	}

	for i, c := range cases {
		i, c := i+1, c
		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
			t.Parallel()
			output := mergeAndSortSimple(c.inputs[0], c.inputs[1])
			if !areArrayEqual(output, c.expectedOutput) {
				t.Fatalf("Unexpected output in case #%d. \n Expected: %v \n Got: %v", i, c.expectedOutput, output)
			}
		})
	}
}

func areArrayEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
