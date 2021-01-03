package main

import (
	"fmt"
	"prime"
)

func main() {
	fmt.Println(3, " = ", prime.LargestPrimeFactor(3))
	fmt.Println(5, " = ", prime.LargestPrimeFactor(5))
	fmt.Println(11, " = ", prime.LargestPrimeFactor(11))
	fmt.Println(13195, " = ", prime.LargestPrimeFactor(13195))
	fmt.Println(600851475143, " = ", prime.LargestPrimeFactor(600851475143))
}
