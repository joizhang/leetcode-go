package q20

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
}

type ans struct {
	one bool
}

type question struct {
	param
	ans
}

func Test_isValid(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			param{"("},
			ans{false},
		},
		{
			param{")"},
			ans{false},
		},
		{
			param{"()"},
			ans{true},
		},
		{
			param{"()[]{}"},
			ans{true},
		},
		{
			param{"(]"},
			ans{false},
		},
		{
			param{"([)]"},
			ans{false},
		},
		{
			param{"{[]}"},
			ans{true},
		},
	}
	for _, q := range qs {
		a, p := q.ans, q.param
		ast.Equal(a.one, isValid(p.one), "输入:%v", p)
	}
}
