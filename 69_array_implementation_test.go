package main

import "testing"

func TestSetStaticArrayInts_BasicOps(t *testing.T) {
	var cases = []struct {
		expectedLen    uint
		expectedValues []int64
	}{
		{expectedLen: 1, expectedValues: []int64{1}},
		{expectedLen: 5, expectedValues: []int64{1, 2, 3, 4, 5}},
	}

	for i, c := range cases {
		actualArr := setStaticArray[int64](c.expectedLen)
		if uint(len(actualArr.items)) != c.expectedLen {
			t.Errorf("Array did not have the number expected on case Nº%d : /n Got: %d /n Expected: %d", i, len(actualArr.items), c.expectedLen)
		}

		for l, v := range c.expectedValues {
			err := actualArr.Push(v)

			if err != nil {
				t.Errorf(
					"Unexpected error on push operation at case N%d: /n Got: %v",
					i, err,
				)
			}

			actualValue := actualArr.Get(l)
			if actualValue != v {
				t.Errorf("Array did not store the expected value on case N%d: /n Got: %d /n Expected: %d", l, actualValue, v)
			}
		}

		// POP operation
		poppedValue, err := actualArr.Pop()
		if err != nil {
			t.Errorf(
				"Unexpected error on pop operation at case N%d: /n Got: %v",
				i, err,
			)
		}

		expectedPopValue := c.expectedValues[len(c.expectedValues)-1]
		if poppedValue != expectedPopValue {
			t.Errorf(
				"Array did not return the expected value for the pop operation on case N%d: /n Got: %d /n Expected: %d",
				i, poppedValue, expectedPopValue,
			)
		}

		// DELETE Operation

		// SHIFT Operation

	}
}
