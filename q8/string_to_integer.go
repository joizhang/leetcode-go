package q8

import (
	"strings"
	"math"
)

func myAtoi(str string) int {
	res := 0
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return res
	}

	var sign uint8 = '+'
	if str[0] == '-' || str[0] == '+' {
		sign = str[0]
		str = str[1:]
	}

	for len(str) > 1 && str[0] == '0' {
		str = str[1:]
	}

	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return res
		}
		pop := int(str[i] - '0')
		if sign == '+' {
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && pop > 7) {
				return math.MaxInt32
			} else {
				res = res*10 + pop
			}
		} else if sign == '-' {
			if res < math.MinInt32/10 || (res == math.MinInt32/10 && -pop < -8) {
				return math.MinInt32
			} else {
				res = res*10 - pop
			}
		}

	}
	return res
}
