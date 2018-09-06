package q15

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
	"fmt"
)

type param struct {
	one []int
}

type ans struct {
	one [][]int
}

type question struct {
	p param
	a ans
}

func Test_threeSum(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: param{[]int{-1, 0, 1, 2, -1, -4}},
			a: ans{[][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			}},
		},
	}
	for _,q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, threeSum(p.one), "输入:%v", p)
	}
}

func Test_search(t *testing.T)  {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	fmt.Println(sort.SearchInts(nums[1:5], -1))
}