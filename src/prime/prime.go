package prime

import (
	"math"
)

func IsPrime(n int) bool {
	switch n {
	case 1: return true
	case 2: return true
	case 3: return true
	default:
		if 0 != (n & 1) {
			if 0 != (n % 3) {
				for i := 5; int(math.Pow(float64(i), 2)) <= n; i += 6 {
					if (n % i) == 0 || (n % (i + 2)) == 0 {
						return false
					}
				}
				return true
			}
		}
		return false
	}
}

func LargestPrimeFactor(n int) int {
	if IsPrime(n) {
		return n
	}
	l := 1
	r := math.Sqrt(float64(n))
	for i := 1; i <= int(r); i += 2 {
		if 0 == n % i && IsPrime(i) {
			l = i
		}
	}
	return l
}

