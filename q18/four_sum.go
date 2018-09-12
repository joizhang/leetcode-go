package q18

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return findNSum(nums, target, 4, 0)
}

func findNSum(nums []int, target int, n int, index int) [][]int {
	var res [][]int
	if index >= len(nums) {
		return res
	}
	if n == 2 {
		i, j := index, len(nums)-1
		for i < j {
			if nums[j]+nums[i] == target {
				res = append(res, []int{nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j-1] == nums[j] {
					j--
				}
				i++
				j--
			} else if nums[j]+nums[i] < target {
				i++
			} else {
				j--
			}
		}
	} else {
		for i := index; i < len(nums)-n+1; i++ {
			temp := findNSum(nums, target-nums[i], n-1, i+1)
			if len(temp) != 0 {
				for _, arr := range temp {
					res = append(res, append([]int{nums[i]}, arr...))
				}
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}
