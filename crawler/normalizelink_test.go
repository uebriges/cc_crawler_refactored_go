package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeLink(t *testing.T) {

	type testCase struct {
		expected string
		base     string
		link     string
	}

	for _, test := range []testCase{
		// {
		// 	expected: "http://test.com/bar",
		// 	base:     "http://test.com/foo",
		// 	link:     "../bar",
		// },
		{
			expected: "http://test.com/foo/bar",
			base:     "http://test.com/foo/baz",
			link:     "../bar",
		},
		// {
		// 	expected: "http://test.com/bar",
		// 	base:     "http://test.com/foo/baz",
		// 	link:     "../../bar",
		// },
		{
			expected: "http://test.com/foo/baz/bar",
			base:     "http://test.com/foo/baz",
			link:     "/bar",
		},
		{
			expected: "http://test.com/index.html",
			base:     "http://test.com/index.html",
			link:     "index.html",
		},
	} {
		actual, err := NormalizeLink(test.base, test.link)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, actual)
	}

}
