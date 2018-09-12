package q17

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
}

type ans struct {
	one []string
}

type question struct {
	param
	ans
}

func Test_letterCombinations(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			param{"23"},
			ans{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
	}
	for _, q := range qs {
		a, p := q.ans, q.param
		ast.Equal(a.one, letterCombinations(p.one), "输入:%v", p)
	}
}
