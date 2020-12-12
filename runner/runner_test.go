package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
