package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func main() {
	list1 := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: &Node{
					Data: 4,
					Next: &Node{
						Data: 5,
						Next: nil,
					},
				},
			},
		},
	}
	list2 := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: &Node{
					Data: 4,
					Next: &Node{
						Data: 5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println("merged lists:", mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	mergedList := &Node{}
	tail := mergedList

	for list1 != nil && list2 != nil {
		if list1.Data <= list2.Data {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}

	return mergedList.Next
}
