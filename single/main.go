package main

import "fmt"

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
}

func main() {
	var list LinkedList
	list.head = nil

	firstNode := Node{next: nil, value: 1}

	fmt.Printf("list: %v\n", list)
	fmt.Printf("firstNode: %v\n", firstNode)
}
