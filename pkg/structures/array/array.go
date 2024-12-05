package main

import "fmt"

func main() {
	arr := createArray()
	slice := createSlice()
	// addElementsIntoArray()
	addElementsIntoSlice()
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
	}
}

// array x slice
// array we must define the size of this array;
// and array is more performatic than slice, because we already know the size and the capacity
// ex bellow:

func createArray() [5]int {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	otherWayToCreateArray := make([]int, 1) // type, size/capacity (capactiy will always be 2x larger than size)
	fmt.Print(otherWayToCreateArray)
	return arr
}

func createSlice() []int {
	var slice []int = []int{1, 2, 3, 4, 5, 6}
	return slice
}

func addElementsIntoArray() {
	array := [5]int{} // wil start as an array with five zeros [0,0,0,0,0]
	array[0] = 1
	array[1] = 2
	// array[5] = 5 it's gonna give an error, because the size of array is five
	fmt.Println(array)
}

func addElementsIntoSlice() {
	slice := []int{}
	slice = append(slice, 10)
	for _, value := range slice {
		fmt.Println(value)
	}
	fmt.Println(slice)
}
