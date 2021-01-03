package main

import (
	"fmt"
	"math"
)

func triangle(n int, tmp *int) int {
	*tmp += n
	return *tmp
}

func pow2(n int) int {
	return int(math.Pow(float64(n), 2))
}

func nfactors(n int) int {
	if n == 1 {
		return 1
	}
	f := 2
	t := n
	for i := 2; i <= t; i++ {
		if 0 == n % i {
			f += 2
			t = n/i
		}
	}
	return f
}

func main() {
	fmt.Println()
	tmp := 0
	for i := 1;; i++ {
		t := triangle(i, &tmp)
		f := nfactors(t)
		fmt.Print("   ", i, " ", t, " = ", f, "    \r")
		if f >= 500 {
			break
		}
	}
	fmt.Println()
}
