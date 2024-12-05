package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	sumNumbersEasyWay(9999999999)
	timeEnd := time.Now()
	fmt.Println("time easy", timeEnd.Sub(timeStart).Milliseconds())

	timeStartHard := time.Now()
	fmt.Println(sumNumbersPerformaticWay(9999999999))
	timeEndHard := time.Now()
	fmt.Println("time performatic", timeEndHard.Sub(timeStartHard).Milliseconds())
}

func sumNumbersEasyWay(n int) {
	var result int = 0
	for i := 1; i <= n; i++ {
		result += i
	}

	fmt.Println(result)
}

func sumNumbersPerformaticWay(n int) int {
	return n * (n + 1) / 2
}
