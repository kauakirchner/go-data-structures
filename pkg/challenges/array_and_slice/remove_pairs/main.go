package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%d \t", removePairs(arr))
}

func removePairs(arr []int) (result []int) {
	for _, value := range arr {
		if value%2 != 0 {
			result = append(result, value)
		}
	}
	return
}
