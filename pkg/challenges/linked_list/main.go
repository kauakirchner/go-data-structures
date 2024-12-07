package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data string
	Next *Node
}

func main() {
	linkedList := getALinkedList()
	linkedList.printLinkedListElements()
	linkedList.removeElementByPosition(2)
	linkedList.printLinkedListElements()
	fmt.Println("size", linkedList.getLength())
	linkedList.insertElementAtBeginning("new first element")
	linkedList.printLinkedListElements()
	linkedList.insertElementAtEnd("last element")
	linkedList.printLinkedListElements()
	linkedList.removeFirstElement()
	linkedList.printLinkedListElements()
	linkedList.removeLastElement()
	linkedList.insertElementByPostion("random", 3)
	linkedList.printLinkedListElements()
}

func getALinkedList() *LinkedList {
	nodes := &LinkedList{
		Head: &Node{
			Data: "first element",
			Next: &Node{
				Data: "second element",
				Next: &Node{
					Data: "third element",
					Next: nil,
				},
			},
		},
	}

	return nodes
}

func (ll *LinkedList) removeElementByPosition(position int) {

	if position <= 0 {
		return
	}

	if position == 1 {
		ll.Head = ll.Head.Next
		return
	}

	count := 1
	previous := ll.Head

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	previous.Next = current.Next
	current = nil
}

func (ll *LinkedList) removeFirstElement() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}
	ll.Head = ll.Head.Next
}

func (ll *LinkedList) removeLastElement() {
	if ll.Head == nil {
		return
	}
	previous := ll.Head
	for previous.Next.Next != nil {
		previous = previous.Next
	}

	previous.Next = nil
}

func (ll *LinkedList) insertElementByPostion(data string, position int) {
	newNode := &Node{Data: data}

	if position == 1 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	previous := ll.Head // 1

	count := 1

	// postion = 3

	// while count < 2
	for count < position-1 {
		previous = previous.Next
		// previous = 1
		// previous = 2
		count++
	}

	current := previous.Next // Node that is in the position we want
	newNode.Next = current   // Our node points to current (ex 3 element)
	previous.Next = newNode  // previous that is 2 now points to new node (3 position)

}

func (ll *LinkedList) insertElementAtEnd(data string) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (ll *LinkedList) insertElementAtBeginning(data string) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) getLength() int {
	current := ll.Head
	size := 0
	for current != nil {
		size++
		current = current.Next
	}
	return size
}

func (ll *LinkedList) printLinkedListElements() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Data, " ---> ")
		current = current.Next
	}
	fmt.Println("null")

}
