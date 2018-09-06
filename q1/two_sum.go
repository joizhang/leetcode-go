package q1

// 1. Two Sum
func twoSum(nums []int, target int) []int {
	var res []int
	indexMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := indexMap[target-num]; ok {
			res = append(res, j, i)
			break
		} else {
			indexMap[num] = i
		}
	}
	return res
}
