package main

import (
	"errors"
)

var NoRecurringValuesErr = errors.New("No recurring values detected")

func FirstRecurringCharacter(charArr []int) (int, error) {
	var recurrentChars = map[int]int{}
	for _, v := range charArr {
		_, ok := recurrentChars[v]
		if ok {
			return v, nil
		}
		recurrentChars[v] = 1
	}
	return 0, NoRecurringValuesErr
}
