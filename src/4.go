package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	m := 0
	for i := 1; i <= 999; i++ {
		for j := 1; j <= 999; j++ {
			k := i * j
			n := strconv.Itoa(k)
			r := []rune(n)
			l := len(r)-1
			x := 0
			for l >= int(len(r)/2) {
				r[x], r[l] = r[l], r[x]
				l--
				x++
			}
			if 0 == strings.Compare(n, string(r)) {
				if m < k {
					m = k
					fmt.Println(i, "*", j, " = ", n, " = ", string(r))
				}
			}
		}
	}
}