package main

// import (
// 	"fmt"
// 	"testing"
// )

// type hashValueTestExpectedValues[T HashValue] [2]T

// func testHashBasicMethods[T HashValue](t *testing.T, testCases hashValueTestExpectedValues[T]) {
// 	hash := newHash[T]()
// 	testKey := "testKey"

// 	var emptyValue T

// 	actualGetValue := hash.Get(testKey)
// 	if expectedValue := emptyValue; areEqual[T](actualGetValue, expectedValue) {
// 		t.Errorf("Unexpected GET value: \n Got: %v \n Expected: %v", actualGetValue, expectedValue)
// 	}

// 	hash.Insert(testKey, testCases[0])
// 	actualGetValue = hash.Get(testKey)
// 	if expectedValue := testCases[0]; actualGetValue != expectedValue {
// 		t.Errorf("Unexpected GET value: \n Got: %v \n Expected: %v", actualGetValue, expectedValue)
// 	}

// 	hash.Insert(testKey, testCases[1])
// 	actualGetValue = hash.Get(testKey)
// 	if expectedValue := testCases[1]; actualGetValue != expectedValue {
// 		t.Errorf("Unexpected GET value: \n Got: %v \n Expected: %v", actualGetValue, expectedValue)
// 	}

// 	hash.Delete(testKey)
// 	actualGetValue = hash.Get(testKey)
// 	if expectedValue := emptyValue; actualGetValue != expectedValue {
// 		t.Errorf("Unexpected GET value: \n Got: %v \n Expected: %v", actualGetValue, expectedValue)
// 	}
// }

// func areEqual[T HashValue](v1 T, v2 T) bool {
// 	switch v1.(type) {
// 	case []string:
// 	case []int:
// 	case []float64:
// 	}

// 	return v1 == v2
// }

// func TestHashingFunctionDigitalRoot(t *testing.T) {
// 	cases := []struct {
// 		key          string
// 		expectedHash int
// 	}{
// 		{key: "ABC", expectedHash: 9},
// 		{key: "dAdA", expectedHash: 6},
// 	}

// 	for i, c := range cases {
// 		i, c := i, c
// 		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
// 			t.Parallel()
// 			key := c.key
// 			expectedHash := c.expectedHash

// 			var h Hash[string]
// 			actualHash := h.Hash(key)
// 			fmt.Println(actualHash)
// 			if actualHash != expectedHash {
// 				t.Fatalf("Hash don't match on case %d. /n Expected: %v /n Got: %v", 0, expectedHash, actualHash)
// 			}
// 		})
// 	}
// }

// func TestHashGenericValues(t *testing.T) {
// 	stringCase := hashValueTestExpectedValues[string]{"lalulilelo", "rarurirero"}
// 	testHashBasicMethods[string](t, stringCase)
// 	t.SkipNow()
// }

// func TestHashMultipleValues(t *testing.T) {
// 	t.SkipNow()
// }

// func TestHashCollision(t *testing.T) {
// 	t.SkipNow()
// }
