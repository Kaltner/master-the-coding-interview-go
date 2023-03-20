package main

import "errors"

// This is an approch of creating an Array type using build in arrays from Go and generics

func throwStackOverflowErr() error {
	return errors.New("Value added that goes beyond this array capacity")
}

func throwStackUnderflowErr() error {
	return errors.New("Value added that goes below this array capacity")
}

type StaticArrayTypes interface {
	int64 | string
}

func setStaticArray[T StaticArrayTypes](length uint) StaticArray[T] {
	var emptyValue T
	var initialItems []T
	for i := 0; i < int(length); i++ {
		initialItems = append(initialItems, emptyValue)
	}
	a := StaticArray[T]{
		items:        initialItems,
		length:       length,
		currentIndex: 0,
	}
	return a
}

type StaticArray[T StaticArrayTypes] struct {
	items        []T
	length       uint
	currentIndex uint
}

func (arr *StaticArray[T]) Push(newItem T) error {
	if arr.currentIndex == arr.length {
		return throwStackOverflowErr()
	}
	arr.items[arr.currentIndex] = newItem
	arr.currentIndex++
	return nil
}

func (arr *StaticArray[T]) Get(index int) T {
	return arr.items[index]
}

func (arr *StaticArray[T]) Pop() (T, error) {
	var emptyValue T
	if arr.currentIndex-1 < 0 {
		return emptyValue, throwStackUnderflowErr()
	}
	arr.currentIndex--
	poppedValue := arr.items[arr.currentIndex]
	arr.items[arr.currentIndex] = emptyValue
	return poppedValue, nil
}

func (arr *StaticArray[T]) Shift()  {}
func (arr *StaticArray[T]) Delete() {}
func (arr *StaticArray[T]) Insert() {}
