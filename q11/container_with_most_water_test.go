package q11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
}

type ans struct {
	one int
}
type question struct {
	p param
	a ans
}

func Test_maxArea(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: param{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			a: ans{49},
		},
	}

	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, maxArea(p.one), "输入:%v", p)
	}
}
