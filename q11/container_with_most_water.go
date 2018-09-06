package q11

// 11. Container With Most Water
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i != j {
		area := (j - i) * min(height[i], height[j])
		if area > res {
			res = area
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
