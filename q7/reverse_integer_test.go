package q7

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_type(t *testing.T) {
	fmt.Println(reflect.TypeOf('a'))
}

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_reverse(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 123,
			},
			a: ans{
				one: 321,
			},
		},
		{
			p: para{
				one: -123,
			},
			a: ans{
				one: -321,
			},
		},
		{
			p: para{
				one: 1234567899,
			},
			a: ans{
				one: 0,
			},
		},
		{
			p: para{
				one: -1234567899,
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverse(p.one), "输入:%v", p)
	}
}
