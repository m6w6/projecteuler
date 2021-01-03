package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return fib(n-1) + fib(n-2)
}

func fibEvenSum() int {
	s := 0
	f := 1
	for {
		e := fib(f)
		if e >= 4000000 {
			break
		}
		fmt.Println(e)
		if 0 == (e & 1) {
			s += e
		}
		f++
	}
	return s
}

func main()  {
	fmt.Println(fibEvenSum())
}