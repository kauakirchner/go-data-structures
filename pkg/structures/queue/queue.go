package main

import (
	"fmt"
)

type Queue struct {
	length int
	front  *ListNode
	rear   *ListNode
}

type ListNode struct {
	data string
	next *ListNode
}

func main() {
	queue := &Queue{}
	validateValueToGenerateBinaryNumber := generateBinaryNumbersMiddleware(queue.generateBinaryNumbers)
	result, err := validateValueToGenerateBinaryNumber(5)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func (ln *Queue) Length() int {
	return ln.length
}

func (ln *Queue) isEmpty() bool {
	return ln.Length() == 0
}

func (q *Queue) enqueue(data string) {
	temp := &ListNode{data: data}
	if q.isEmpty() {
		q.front = temp
	} else {
		q.rear.next = temp
	}

	q.rear = temp
	q.length++
}

func (q *Queue) dequeue() string {
	if q.isEmpty() {
		panic("is not possible to dequeue an empty queue")
	}

	result := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.length--
	return result
}
