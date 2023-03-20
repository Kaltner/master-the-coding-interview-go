package main

// Given 2 arrays, create a function that let's a user know if they contain common itens

// Inputs = Arrays
// Output = Boolean

func CompareTwoArraysValuesSlow(arr1, arr2 []string) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}
	return false
}

func CompareTwoArraysValuesFast(arr1, arr2 []string) bool {
	hashTable := make(map[string]bool)
	for i := 0; i < len(arr2); i++ {
		_, ok := hashTable[arr2[i]]
		if ok == false {
			hashTable[arr2[i]] = true
		}

	}

	for _, v := range arr1 {
		if hashTable[v] {
			return true
		}
	}

	return false
}

type set struct {
	m map[string]bool
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]bool)
	return s
}

func (s *set) Add(value string) {
	_, ok := s.m[value]
	if !ok {
		s.m[value] = true
	}
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, ok := s.m[value]
	return ok
}

func CompareTwoArraysValuesStruct(arr1, arr2 []string) bool {
	arrSet := NewSet()
	for _, v := range arr2 {
		arrSet.Add(v)
	}

	for _, v := range arr1 {
		if arrSet.Contains(v) {
			return true
		}
	}
	return false
}

// Nested loops are O(n^2), so not the best
// Creating a set first should also create two loops, so it's quicker, but worst case scenario, we are still working on two loops
// Apparently the answer above = solution since it will be quicker for the second loop.
// Hash table
