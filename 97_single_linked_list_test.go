package main

import (
	"testing"
)

func TestSingleLinkedListBasicOperations(t *testing.T) {
	linkedList := NewSingleLinkedList()

	for i := 0; i < 5; i++ {
		err := linkedList.Append(i)
		if err != nil {
			t.Fatalf("Unexpected error in the Append op: \n Got: %v", err)
		}

		getPosiOp, err := linkedList.GetByPosition(i)
		if err != nil {
			t.Fatalf("Unexpected error in the GetByPosition op: \n Got: %v", err)
		}

		if getPosiOp != i {
			t.Fatalf("Unexpected value in the Get Op: \n Expected: %d \n Got: %d", i, getPosiOp)
		}
	}

	getOp, err := linkedList.Get(linkedList.firstNode)
	if err != nil {
		t.Fatalf("Unexpected error in the Get op: \n Got: %v", err)
	}

	if getOp != 0 {
		t.Fatalf("Unexpected value in the Get Op: \n Expected: %d \n Got: %d", 0, getOp)
	}

	getKeyByValueOp, err := linkedList.GetKeyByValue(4)
	if err != nil {
		t.Fatalf("Unexpected error in the getKeyByValue op: \n Got: %v", err)
	}

	if getKeyByValueOp != linkedList.lastNode {
		t.Fatalf("Unexpected value in the getKeyByValue Op: \n Expected: %d \n Got: %d", 4, getKeyByValueOp)
	}

	err = linkedList.Insert(2, 5)
	if err != nil {
		t.Fatalf("Unexpected error in the Insert op: \n Got: %v", err)
	}

	getPosiOp, err := linkedList.GetByPosition(2)
	if err != nil {
		t.Fatalf("Unexpected error in the GetByPosition op: \n Got: %v", err)
	}

	if getPosiOp != 5 {
		t.Fatalf("Unexpected value in the Get Op: \n Expected: %d \n Got: %d", 5, getPosiOp)
	}

	getPosiOp, err = linkedList.GetByPosition(3)
	if err != nil {
		t.Fatalf("Unexpected error in the GetByPosition op: \n Got: %v", err)
	}

	if getPosiOp != 2 {
		t.Fatalf("Unexpected value in the Get Op: \n Expected: %d \n Got: %d", 2, getPosiOp)
	}
}

// func TestNodeIntReturn(t *testing.T) {
// 	node := Node{
// 		value: 5,
// 		link:  0,
// 	}
// 	if int(node) != 5 {
// 		t.Error("")
// 	}
// }
