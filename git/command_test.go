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
