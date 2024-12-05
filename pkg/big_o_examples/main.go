package main

import "fmt"

func main() {
	printValues(5)
}

// this func bellow is and O(n²);
// it have a bad performance, because the second for loop could run a milion times
// to each first loop iteration, so avoid nested loops is a good pratice
func printValues(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Println(fmt.Sprintf("i = %s, j = %s", i, j))
		}
		fmt.Println("Finished");
	}
	fmt.Println("outside loop")
}
