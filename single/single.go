package single

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
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

func (list *LinkedList) reverse() {
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