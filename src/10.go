package main

import (
	"fmt"
	"prime"
)

func main() {
	sum := 0
	for i := 2; i < 2000000; i++ {
		if prime.IsPrime(i) {
			if i < 20 {
				fmt.Println(i)
			}
			sum += i
		}
	}
	fmt.Println(sum)
}