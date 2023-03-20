package main

import (
	"strings"
)

func reverseStringRunes(input string) string {
	runeString := []rune(input)
	reversedRune := make([]rune, len(runeString))
	l := 0
	for i := len(runeString) - 1; i >= 0; i-- {
		reversedRune[l] = runeString[i]
		l++
	}
	return string(reversedRune)
}

// NOT UNICODE SAFE
func reverseStringBytes(input string) string {
	var b strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		b.WriteByte(input[i])
	}
	return b.String()
}
