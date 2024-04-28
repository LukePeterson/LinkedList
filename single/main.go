package main

import "fmt"

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func main() {
	var list LinkedList
	list.toList([]int{1, 2, 3})
	fmt.Printf("list.toSlice(): %v\n", list.toSlice())
}

func (list *LinkedList) toList(input []int) {
	for i := range input {
		list.append(&Node{value: input[i]})
	}
}

func (list *LinkedList) toSlice() []int {
	if list.head == nil {
		return []int{}
	}

	result := []int{}
	currentNode := list.head
	for currentNode != nil {
		result = append(result, currentNode.value)
		currentNode = currentNode.next
	}

	return result
}

func (list *LinkedList) append(newNode *Node) {
	if newNode == nil {
		return
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}
