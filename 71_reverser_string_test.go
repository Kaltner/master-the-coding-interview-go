package main

import (
	"fmt"
	"testing"
)

func TestStringReversers(t *testing.T) {
	cases := []struct {
		input          string
		expectedOutput string
	}{
		{input: "a", expectedOutput: "a"},
		{input: "ab", expectedOutput: "ba"},
		{input: "Danilo", expectedOutput: "olinaD"},
		{input: "Hello my name is Mr. X!", expectedOutput: "!X .rM si eman ym olleH"},
		{input: "1232 13214 1234151 ", expectedOutput: " 1514321 41231 2321"},
		{input: "      21      ", expectedOutput: "      12      "},
		{input: "こんにちは世界", expectedOutput: "界世はちにんこ"},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
			t.Parallel()
			outputReverse := reverseStringRunes(c.input)
			if outputReverse != c.expectedOutput {
				t.Errorf(
					"Return for the `reverseStringRunes` does not match for case #%d \n Expected: %v \n Got: %v ",
					i, c.expectedOutput, outputReverse,
				)
			}
		})
	}
}
