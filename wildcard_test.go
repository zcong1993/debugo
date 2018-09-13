package debugo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	args [2]string
	want bool
}
type fixture []testCase

func TestMatch(t *testing.T) {
	f := fixture{
		testCase{
			[2]string{"", "hello"},
			false,
		},
		testCase{
			[2]string{"hello", "hello"},
			true,
		},
		testCase{
			[2]string{"*llo", "hello"},
			true,
		},
		testCase{
			[2]string{"he*", "hello"},
			true,
		},
		testCase{
			[2]string{"hes*", "hello"},
			false,
		},
		testCase{
			[2]string{"*e*", "hello"},
			true,
		},
	}

	for _, v := range f {
		assert.Equal(t, v.want, Match(v.args[0], v.args[1]))
	}
}
