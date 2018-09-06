package q13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	p param
	a ans
}

func Test_romanToInt(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: param{"III"},
			a: ans{3},
		},
		{
			p: param{"IV"},
			a: ans{4},
		},
		{
			p: param{"IX"},
			a: ans{9},
		},
		{
			p: param{"LVIII"},
			a: ans{58},
		},
		{
			p: param{"MCMXCIV"},
			a: ans{1994},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, romanToInt(p.one), "输入:%v", p)
	}
}
