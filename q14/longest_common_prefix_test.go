package q14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []string
}

type ans struct {
	one string
}

type question struct {
	p param
	a ans
}

func Test_longestCommonPrefix(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: param{[]string{"flower", "flow", "flight"}},
			a: ans{"fl"},
		},
		{
			p: param{[]string{"dog", "racecar", "car"}},
			a: ans{""},
		},
		{
			p: param{[]string{"aa", "a"}},
			a: ans{"a"},
		},
		{
			p: param{[]string{"ca", "a"}},
			a: ans{""},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, longestCommonPrefix(p.one), "输入:%v", p)
	}
}
