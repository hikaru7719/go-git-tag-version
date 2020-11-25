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
			TagVersioning(version)
			defer DeleteTag(version)
			result, _ := Tag()
			assert.Equal(t, tc, result)
		})
	}
}
