package main

import (
	"fmt"

	"github.com/LukePeterson/linkedlist/single"
)

func main() {
	var list single.LinkedList

	list.Create([]int{1, 2, 3})
	fmt.Printf("list.Slice(): %v\n", list.Slice())

	list.Reverse()
	fmt.Printf("list.Slice() (after reverse): %v\n", list.Slice())
}
