package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 5, 4, 7, 6}
	fmt.Printf("%d \t", invertList(arr))
}

func invertList(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
