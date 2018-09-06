package q7

import "math"

// 7. Reverse Integer
func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return math.MaxInt32
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return math.MinInt32
		}
		rev = rev*10 + pop
	}
	return rev
}
