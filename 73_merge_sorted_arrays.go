package main

import "fmt"

// Only works for positive results
func mergeAndSortArraysCountMethod(arr1, arr2 []int) []int {
	highestValue := arr1[len(arr1)-1]
	if arr2[len(arr2)-1] > highestValue {
		highestValue = arr2[len(arr2)-1]
	}

	mergedArray := append(arr1, arr2...)

	count := make([]int, highestValue+1)
	output := make([]int, len(mergedArray))

	fmt.Println(count)
	for k, _ := range mergedArray {
		count[k] = count[k] + 1
	}

	fmt.Println(count)

	for i := 1; i <= highestValue; i++ {
		count[i] = count[i] + count[i-1]
	}

	fmt.Println(count)

	fmt.Println(output)
	for i := len(mergedArray) - 1; i > 0; i-- {
		count[i] = count[i] - 1
		output[count[i]] = mergedArray[i]
	}

	fmt.Println(count)
	fmt.Println(output)

	return output
}

func mergeAndSortSimple(arr1, arr2 []int) []int {
	if len(arr2) == 0 {
		return arr1
	}

	if len(arr1) == 0 {
		return arr2
	}

	var mergedArray []int
	i, j := 0, 0

	for i < len(arr1) || j < len(arr2) {
		var nextInt int
		var arr1Ok bool
		arr1Value, arr2Value := 0, 0
		if i < len(arr1) {
			arr1Value = arr1[i]
			arr1Ok = true
		}
		if j < len(arr2) {
			arr2Value = arr2[j]
		}
		if arr1Ok && (arr1Value < arr2Value) {
			nextInt = arr1Value
			i++
		} else {
			nextInt = arr2Value
			j++
		}
		mergedArray = append(mergedArray, nextInt)
	}

	return mergedArray
}
