package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	arr1           []string
	arr2           []string
	expectedResult bool
}{
	{
		arr1:           []string{},
		arr2:           []string{},
		expectedResult: false,
	},
	{
		arr1:           []string{"a"},
		arr2:           []string{"a"},
		expectedResult: true,
	},
	{
		arr1:           []string{"a"},
		arr2:           []string{"b"},
		expectedResult: false,
	},
	{
		arr1:           []string{"a", "b", "i", "z"},
		arr2:           []string{"x", "c", "v", "z"},
		expectedResult: true,
	},
	{
		arr1:           []string{"a", "b", "i", "z"},
		arr2:           []string{"x", "c", "v", "p", "q", "w", "r", "e"},
		expectedResult: false,
	},
	{
		arr1:           []string{"a", "b", "i", "z"},
		arr2:           []string{"x", "c", "v", "p", "q", "w", "r", "e", "x", "c", "v", "p", "q", "w", "r", "e", "x", "c", "v", "p", "q", "w", "r", "e", "x", "c", "v", "p", "q", "w", "r", "e", "x", "c", "v", "p", "q", "w", "r", "e", "a"},
		expectedResult: true,
	},
}

func TestCompareTwoArraysValuesSlow(t *testing.T) {
	for i, c := range cases {
		actualSlowResult := CompareTwoArraysValuesSlow(c.arr1, c.arr2)
		if actualSlowResult != c.expectedResult {
			t.Errorf(fmt.Sprintf("Failed case #%d on slow approach. Expected: %+v Got: %+v", i, c.expectedResult, actualSlowResult))
		}
	}
}

func TestCompareTwoArraysValuesFast(t *testing.T) {
	for i, c := range cases {
		actualFastResult := CompareTwoArraysValuesFast(c.arr1, c.arr2)
		if actualFastResult != c.expectedResult {
			t.Errorf(fmt.Sprintf("Failed case #%d on fast approach. Expected: %+v Got: %+v", i, c.expectedResult, actualFastResult))
		}
	}
}

func TestCompareTwoArraysValuesStruct(t *testing.T) {
	for i, c := range cases {
		actualStructResult := CompareTwoArraysValuesStruct(c.arr1, c.arr2)
		if actualStructResult != c.expectedResult {
			t.Errorf(fmt.Sprintf("Failed case #%d on struct approach. Expected: %+v Got: %+v", i, c.expectedResult, actualStructResult))
		}
	}
}
