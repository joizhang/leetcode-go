package q704

// 704. Binary Search
func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, lo int, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	if target < nums[mid] {
		return binarySearch(nums, target, lo, mid-1)
	} else if target > nums[mid] {
		return binarySearch(nums, target, mid+1, hi)
	} else {
		return mid
	}
}
