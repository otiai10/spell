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
		{
			Given: `foobaa`,
			Expect: []string{
				"foobaa",
			},
		},
		{
			Given: "'unclosed-quote",
			Expect: []string{
				"'unclosed-quote",
			},
		},
	}

	for _, c := range cases {
		words := Parse(c.Given)
		Expect(t, len(words)).ToBe(len(c.Expect))
		Expect(t, []string(words)).Deeply().ToBe(c.Expect)
	}

	words := Parse("@otiai10 hey yo, what's up?")
	Expect(t, words.Flag("yo,")).ToBe(true)
	Expect(t, words.Flag("hat")).ToBe(false)
	Expect(t, words.Arg("hey")).ToBe("yo,")
	Expect(t, words.Arg("up?")).ToBe("")
	words.Remove("hey", 1)
	Expect(t, words.Flag("@otiai10")).ToBe(true)
	Expect(t, words.Flag("hey")).ToBe(false)
	Expect(t, words.Flag("yo,")).ToBe(false)
	Expect(t, words.Flag("up?")).ToBe(true)

	words = Parse("@otiai10 hey yo, what's up?")
	words.Remove("hey", 0)
	Expect(t, words.Flag("hey")).ToBe(false)
	Expect(t, words.Flag("yo,")).ToBe(true)
}
