package main

import (
	"fmt"
)

// Hash table exercise requirements:
// Key is a string, but will transform into a idenpotent hash
// Value can be any

// Methods:
// LOOKOUT by key
// INSERT using key
// DELETE on key
// SEARCH
// Hash function

type HashValue interface {
	string | int | float64 | []string | []int | []float64
}

type HashEntry[T HashValue] struct {
	k string
	v T
}

type Hash[T HashValue] [][]HashEntry[T]

func newHash[T HashValue]() Hash[T] {
	return make(Hash[T], 10)
}

func (h Hash[T]) Insert(key string, value T) {
	newEntry := HashEntry[T]{
		k: key,
		v: value,
	}

	hash := h.Hash(key)
	hashEntries := h[hash]
	if len(hashEntries) > 0 {
		for k, v := range hashEntries {
			if v.k == key {
				h[hash][k] = newEntry
				return
			}
		}
	}

	h[hash] = append(h[hash], newEntry)
}

func (h Hash[T]) Get(key string) T {
	hash := h.Hash(key)

	hashEntries := h[hash]
	lenEntries := len(hashEntries)
	switch lenEntries {
	case 0:
		var emptyValue T
		return emptyValue
	case 1:
		return h[hash][0].v
	}

	// Cases if there is more than 1 entry with that hash value
	var returnValue T
	for _, v := range h[hash] {
		if v.k == key {
			returnValue = v.v
			break
		}
	}
	return returnValue
}

func (h Hash[T]) Delete(key string) {
	hash := h.Hash(key)

	hashEntries := h[hash]
	nEntries := len(hashEntries)
	if nEntries == 0 {
		return
	}

	entryN := 0
	if nEntries > 1 {
		for k, v := range hashEntries {
			if v.k == key {
				entryN = k
				break
			}
		}
	}

	hashEntries[entryN] = hashEntries[len(hashEntries)-1]
	h[hash] = hashEntries[:len(hashEntries)-1]
}

func (h Hash[T]) Hash(key string) int {
	return h.hashToDigitalRoot(key)
}

// Extremelly weak hashing function, but good to test some collision cases
func (h Hash[T]) hashToDigitalRoot(key string) int {
	hash := 0

	for _, c := range key {
		hash += int(c)
	}

	digitalRoot := hash
	for digitalRoot > 10 {
		currentDigitalRoot := 0
		for _, c := range fmt.Sprintf("%d", digitalRoot) {
			currentDigitalRoot += int(c) - '0'
		}
		digitalRoot = currentDigitalRoot
	}

	return digitalRoot
}
