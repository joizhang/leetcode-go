package q16

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
	two int
}

type ans struct {
	one int
}

type question struct {
	p param
	a ans
}

func Test_threeSumClosest(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: param{
				one: []int{-1, 2, 1, -4},
				two: 1,
			},
			a: ans{2},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, threeSumClosest(p.one, p.two), "输入:%v", p)
	}
}
