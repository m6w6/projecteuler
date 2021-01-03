package main

import (
	"fmt"
	"math"
	"os"
)

func pow2(a int) int {
	return int(math.Pow(float64(a),2))
}

func pythTriplet(a, b, c int) bool {
	return a + b + c == 1000 && pow2(a) + pow2(b) == pow2(c)
}

func main() {
	for a:= 1; a < 999; a++ {
		for b := 1; b < 999; b++ {
			if b < a {
				continue
			}
			for c:= 1; c < 999; c++ {
				if c < b {
					continue
				}
				if pythTriplet(a, b ,c) {
					fmt.Println(a, "*", b, "*", c, " = ", a*b*c)
					os.Exit(0)
				}
			}
		}
	}
}
