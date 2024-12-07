package main

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func main() {
	linkedList := getALinkedList()
	linkedList.invertList()
}

func (ll *LinkedList) invertList() {
	current := ll.Head
	var next *Node
	var previous *Node

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	ll.Head = previous
}

func getALinkedList() *LinkedList {
	nodes := &LinkedList{
		Head: &Node{
			Data: 1,
			Next: &Node{
				Data: 2,
				Next: &Node{
					Data: 3,
					Next: nil,
				},
			},
		},
	}

	return nodes
}
