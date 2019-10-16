package spell

import (
	"testing"

	. "github.com/otiai10/mint"
)

type testcase struct {
	Label  string
	Given  string
	Expect []string
}

func TestParse(t *testing.T) {

	cases := []testcase{
		{
			Given: `bash -c 'for i in 0 1 2 3; do echo ${i} && sleep 1; done'`,
			Expect: []string{
				"bash",
				"-c",
				"'for i in 0 1 2 3; do echo ${i} && sleep 1; done'",
			},
		},
		{
			Given: `echo "foo baa"`,
			Expect: []string{
				"echo",
				`"foo baa"`,
			},
		},
	}

	for _, c := range cases {
		tokens := Parse(c.Given)
		Expect(t, len(tokens)).ToBe(len(c.Expect))
		Expect(t, tokens).Deeply().ToBe(c.Expect)
	}
}
