package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newGit() Git {
	return Git{}
}

func TestGitFetch(t *testing.T) {
	cases := map[string]string{
		"execute git fetch": "",
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result, _ := newGit().Fetch()
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
			git := newGit()
			version := "v99.99.99"
			// nolint:errcheck
			git.TagVersioning(version)
			// nolint:errcheck
			defer git.DeleteTag(version)
			result, _ := git.Tag()
			assert.Equal(t, tc, result)
		})
	}
}
