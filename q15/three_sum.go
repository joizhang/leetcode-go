package q15

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	if len(nums) < 3 || nums[0] > 0 || nums[len(nums)-1] < 0 {
		return res
	}
	left, right := 0, len(nums)-1
	for right-left >= 1 {
		sum := nums[left] + nums[right]
		index := sort.SearchInts(nums[left+1:right], -sum)
		if index != -1 {
			res = append(res, []int{nums[left], nums[index], nums[right]})
			continue
		}
		if sum < 0 {
			left++
		} else {
			right--
		}
	}
	return res
}
