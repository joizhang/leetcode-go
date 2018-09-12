package q18

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
	two int
}

type ans struct {
	one [][]int
}

type question struct {
	param
	ans
}

func Test_fourSum(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			param{
				one: []int{1, 0, -1, 0, -2, 2},
				two: 0,
			},
			ans{[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			}},
		},
	}
	for _, q := range qs {
		a, p := q.ans, q.param
		ast.Equal(a.one, fourSum(p.one, p.two), "输入:%v", p)
	}
}
