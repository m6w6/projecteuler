package main

import (
	"fmt"
	"os"
)

func main() {
	NEXT:
	for n := 21;; n++ {
		for i := 1; i <= 20; i++ {
			if n % i != 0 {
				continue NEXT
			}
		}
		fmt.Println(n)
		for i := 1; i <= 20; i++ {
			fmt.Println(n, "/", i , "=", n/i)
		}
		os.Exit(0)
	}
}
