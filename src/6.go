package main

import (
	"fmt"
	"math"
)

func sqSum(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func sumSq(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

func main() {
	fmt.Println(sqSum(100)-sumSq(100))
}