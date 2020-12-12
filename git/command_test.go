package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitFetch(t *testing.T) {
	cases := map[string]string{
		"execute git fetch": "",
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result, _ := Fetch()
			assert.Equal(t, tc, result)
		})
	}
}

func TestGitTag(t *testing.T) {
	cases := map[string]string{
		"execute git tag": "v99.99.99\n",
	}
	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			version := "v99.99.99"
			// nolint:errcheck
			TagVersioning(version)
			// nolint:errcheck
			defer DeleteTag(version)
			result, _ := Tag()
			assert.Equal(t, tc, result)
		})
	}
}

func TestParse(t *testing.T) {
	cases := map[string]struct {
		input       string
		expectArray []string
	}{
		"parse git tag string": {
			input:       "v99.99.99\nv99.99.98\nv99.99.97\n",
			expectArray: []string{"v99.99.99", "v99.99.98", "v99.99.97"},
		},
	}
	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result := Parse(tc.input)
			assert.Equal(t, tc.expectArray, result)
		})
	}
}
