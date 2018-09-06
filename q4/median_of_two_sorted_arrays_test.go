package q4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one []int
	two []int
}

type ans struct {
	one float64
}

type question struct {
	p para
	a ans
}

func Test_findMedianSortedArrays(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: para{
				one: []int{1, 3},
				two: []int{2},
			},
			a: ans{
				one: 2,
			},
		},
		{
			p: para{
				one: []int{1, 3},
				two: []int{2, 4},
			},
			a: ans{
				one: 2.5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findMedianSortedArrays(p.one, p.two), "输入:%v", p)
	}
}
