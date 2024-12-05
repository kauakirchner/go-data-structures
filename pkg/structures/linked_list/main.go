package main

import "fmt"

// A liked list is an structure that save linked nodes, that follow the bellow props:
//  - A node has an information and saves the next node address.
//  - The first node is called "head"
//  - The last node points to null

type Node struct {
	Data string
	Next *Node
}

func main() {
	linkedList := getALinkedList()
	linkedList.printLinkedListElements()
}

func (n *Node) printLinkedListElements() {
	current := n
	for current != nil {
		fmt.Print(current.Data, " ---> ")
		current = current.Next
	}
	fmt.Println("null")
}

func getALinkedList() Node {
	nodes := Node{
		Data: "just a test",
		Next: &Node{
			Data: "More tests",
			Next: &Node{
				Data: "Last test",
				Next: nil,
			},
		},
	}

	return nodes
}
