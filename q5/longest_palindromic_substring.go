package q5

import "math"

// 5. Longest Palindromic Substring
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < length-1; i++ {
		len1 := extendPalindrome(s, i, i)
		len2 := extendPalindrome(s, i, i+1)
		len3 := int(math.Max(float64(len1), float64(len2)))
		if len3 > end+1-start {
			start = i - (len3-1)/2
			end = i + len3/2
		}
	}
	return s[start : end+1]
}

func extendPalindrome(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
