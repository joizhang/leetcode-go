package q9

// 9. Palindrome Number
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	highestPos := 1
	for x/highestPos >= 10 {
		highestPos *= 10
	}
	for {
		if highestPos == 1 || x == 0 {
			break
		}
		if x/highestPos != x%10 {
			return false
		}
		x %= highestPos
		x /= 10
		highestPos /= 100
	}
	return true
}
