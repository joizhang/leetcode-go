package q704

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func Test_search(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: param{
				one: []int{-1, 0, 3, 5, 9, 12},
				two: 9,
			},
			a: ans{4},
		},
		{
			p: param{
				one: []int{5},
				two: 5,
			},
			a: ans{0},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, search(p.one, p.two), "输入:%v", p)
	}
}
