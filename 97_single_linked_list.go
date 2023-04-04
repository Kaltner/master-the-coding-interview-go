package main

import (
	"errors"
	"math/rand"
)

/*
* 1 -> 2 -> 10 -> null
* Is this null terminated?
* Any kind of data in the node?
* DO we want the Insert to enclose Append and Prepend?

*
* GET -> BY INDEX OF THE NODE, NOT THE MEMORY
* SEARCH ->  bY VALUE
* INSERT
* DELETE

 */

var nodeNotExist error = errors.New("Node does not exist")
var maxNumberOfNodesErr error = errors.New("Reached max number of nodes")
var invalidPostionErr error = errors.New("Linked List is not big enough for the position chosen")

const MAX_NODES = 100

// I want this to return value if someone requests a Int
type SingleLinkedList struct {
	firstNode  int
	lastNode   int
	linkedList map[int]Node
}

type Node struct {
	value int
	link  int
}

func (n *Node) Int() int {
	return n.value
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		firstNode:  0,
		linkedList: map[int]Node{},
	}
}

func (s *SingleLinkedList) Append(value int) error {
	key, err := s.genNewRandKey()
	if err != nil {
		return err
	}

	s.linkedList[key] = Node{
		value: value,
		link:  0,
	}

	if s.lastNode > 0 {
		s.linkedList[s.lastNode] = Node{
			value: s.linkedList[s.lastNode].value,
			link:  key,
		}
	}
	s.lastNode = key

	if s.firstNode == 0 {
		s.firstNode = key
	}

	return nil
}

func (s *SingleLinkedList) Prepend(value int) error {
	key, err := s.genNewRandKey()
	if err != nil {
		return err
	}

	s.linkedList[key] = Node{
		value: value,
		link:  s.firstNode,
	}

	s.firstNode = key

	if s.lastNode == 0 {
		s.lastNode = key
	}
	return nil
}

func (s *SingleLinkedList) Get(key int) (int, error) {
	node, ok := s.linkedList[key]
	if !ok {
		return 0, nodeNotExist
	}
	return node.value, nil
}

func (s *SingleLinkedList) GetByPosition(position int) (int, error) {
	if position > len(s.linkedList) {
		return 0, invalidPostionErr
	}

	nextNodeKey := s.firstNode
	for i := 0; i <= position; i++ {
		if i == position {
			return s.linkedList[nextNodeKey].value, nil
		}
		nextNodeKey = s.linkedList[nextNodeKey].link
	}

	return 0, invalidPostionErr
}

func (s *SingleLinkedList) Insert(position, value int) error {
	if position > len(s.linkedList) {
		return invalidPostionErr
	}

	nextNodeKey := s.firstNode
	for i := 0; i <= position; i++ {
		if i == position {
			currentNode := nextNodeKey
			nextNodeKey, err := s.genNewRandKey()
			nextNode := s.linkedList[currentNode]
			if err != nil {
				return err
			}
			s.linkedList[currentNode] = Node{value, nextNodeKey}
			s.linkedList[nextNodeKey] = nextNode

			return nil
		}
		nextNodeKey = s.linkedList[nextNodeKey].link
	}

	return invalidPostionErr
}

func (s *SingleLinkedList) GetKeyByValue(value int) (int, error) {
	nothingFound := false
	currentNode := s.firstNode
	for !nothingFound {
		node := s.linkedList[currentNode]
		if node.value == value {
			return currentNode, nil
		}
		if node.link == 0 {
			nothingFound = true
		}
		currentNode = node.link
	}

	return 0, nodeNotExist
}

func (s *SingleLinkedList) genNewRandKey() (int, error) {
	if len(s.linkedList) == MAX_NODES+1 {
		return 0, maxNumberOfNodesErr
	}
	randKey := rand.Intn(MAX_NODES) + 1
	_, ok := s.linkedList[randKey]
	if ok {
		return s.genNewRandKey()
	}
	return randKey, nil
}
