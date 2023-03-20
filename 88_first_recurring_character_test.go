package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestFirstRecurringCharacter(t *testing.T) {
	cases := []struct {
		input         []int
		expectedValue int
		expectedErr   error
	}{
		{
			input:         []int{1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9},
			expectedValue: 2,
			expectedErr:   nil,
		},
		{
			input:         []int{},
			expectedValue: 0,
			expectedErr:   NoRecurringValuesErr,
		}, {
			input:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedValue: 0,
			expectedErr:   NoRecurringValuesErr,
		}, {
			input:         []int{1},
			expectedValue: 0,
			expectedErr:   NoRecurringValuesErr,
		},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Test case #%d", i), func(t *testing.T) {
			t.Parallel()
			actualValue, err := FirstRecurringCharacter(c.input)
			if !errors.Is(err, c.expectedErr) {
				t.Fatalf("Unexpected error. \n Got: %v \n Expected: %v", err, c.expectedErr)
			}

			if actualValue != c.expectedValue {
				t.Fatalf("Unexpected value for the recurring character. \n Got: %v \n Expected: %v", actualValue, c.expectedValue)
			}

		})
	}
}
