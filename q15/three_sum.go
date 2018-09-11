package q15

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo, hi, sum := i+1, len(nums)-1, 0-nums[i]
			for lo < hi {
				if nums[lo]+nums[hi] == sum {
					res = append(res, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return res
}
