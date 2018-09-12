package q19

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	param
	ans
}

func Test_removeNthFromEnd(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			param{[]int{1, 2}, 2},
			ans{[]int{2}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, 1},
			ans{[]int{1, 2, 3, 4}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, 2},
			ans{[]int{1, 2, 3, 5}},
		},
		{
			param{[]int{1}, 1},
			ans{nil},
		},
	}
	for _, q := range qs {
		a, p := q.ans, q.param
		ast.Equal(a.one, l2s(removeNthFromEnd(s2l(p.one), p.two)), "输入:%v", p)
	}
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}
