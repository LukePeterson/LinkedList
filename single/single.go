// Package single implements a singly linked list of nodes
package single

// Node represents a single node in the list
type Node struct {
	next  *Node
	value int
}

// LinkedList represents a single list, with a head and a tail
type LinkedList struct {
	head *Node
	tail *Node
}

// Create creates a new linked list based on an array of ints
func (list *LinkedList) Create(input []int) {
	for i := range input {
		list.Append(&Node{value: input[i]})
	}
}

// Slice converts a linked list to an array of ints
func (list *LinkedList) Slice() []int {
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

// Append adds a new Node to the end of a list
func (list *LinkedList) Append(newNode *Node) {
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

// Reverse reverses a list
func (list *LinkedList) Reverse() {
	var prev *Node = nil
	curr := list.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}
