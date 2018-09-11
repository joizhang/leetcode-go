package q16

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	var res int
	sort.Ints(nums)
	lastDistance := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if sum > target {
				hi--
			} else {
				lo++
			}
			distance := distance(sum, target)
			if distance < lastDistance {
				lastDistance = distance
				res = sum
			}
		}
	}
	return res
}

func distance(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
