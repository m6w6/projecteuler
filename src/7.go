package main

import (
	"fmt"
	"prime"
)

func main() {
	n := 0 // 2
	for i := 1; n < 10001; i++ {
		if prime.IsPrime(i) {
			n++
			fmt.Println(n, " = ", i)
		}
	}
}
