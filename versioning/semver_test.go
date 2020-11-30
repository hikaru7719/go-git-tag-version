package versioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementMajor(t *testing.T) {
	cases := map[string]struct {
		before *SemVer
		after  *SemVer
	}{
		"increment major version": {
			before: New(1, 0, 0),
			after:  New(2, 0, 0),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result := tc.before.IncrementMajor()
			assert.Equal(t, tc.after, result)
		})
	}
}

func TestIncrementMinor(t *testing.T) {
	cases := map[string]struct {
		before *SemVer
		after  *SemVer
	}{
		"increment minor version": {
			before: New(1, 0, 0),
			after:  New(1, 1, 0),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result := tc.before.IncrementMinor()
			assert.Equal(t, tc.after, result)
		})
	}
}

func TestIncrementPatch(t *testing.T) {
	cases := map[string]struct {
		before *SemVer
		after  *SemVer
	}{
		"increment minor version": {
			before: New(1, 0, 0),
			after:  New(1, 0, 1),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result := tc.before.IncrementPatch()
			assert.Equal(t, tc.after, result)
		})
	}
}

func TestToString(t *testing.T) {
	cases := map[string]struct {
		target *SemVer
		expect string
	}{
		"increment minor version": {
			target: New(1, 0, 0),
			expect: "v1.0.0",
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result := tc.target.ToString()
			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestFrom(t *testing.T) {
	cases := map[string]struct {
		input  []string
		expect SemVerList
	}{
		"semver list from string literal": {
			input: []string{"v99.99.99", "v99.99.98", "v99.99.97"},
			expect: SemVerList{
				SemVer{
					Major: 99,
					Minor: 99,
					Patch: 99,
				},
				SemVer{
					Major: 99,
					Minor: 99,
					Patch: 98,
				},
				SemVer{
					Major: 99,
					Minor: 99,
					Patch: 97,
				},
			},
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			result := From(tc.input)
			assert.Equal(t, tc.expect, result)
		})
	}
}
