package q9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one int
}

type ans struct {
	one bool
}

type question struct {
	p param
	a ans
}

func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: param{121},
			a: ans{true},
		},
		{
			p: param{-121},
			a: ans{false},
		},
		{
			p: param{10},
			a: ans{false},
		},
		{
			p: param{+0},
			a: ans{true},
		},
		{
			p: param{11},
			a: ans{true},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isPalindrome(p.one), "输入:%v", p)
	}
}
