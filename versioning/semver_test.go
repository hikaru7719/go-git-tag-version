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
		"incremental major version when minor and patch is not zero": {
			before: New(1, 1, 2),
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
		"increment minor version when patch is not zero": {
			before: New(1, 1, 1),
			after:  New(1, 2, 0),
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

func TestSemverListLatest(t *testing.T) {
	cases := map[string]struct {
		list   SemVerList
		expect *SemVer
	}{
		"latest version is v2.0.0 among [v2.0.0, v1.0.0, v1.2.0, v1.1.1, v0.0.1]": {
			list:   From([]string{"v2.0.0", "v1.0.0", "v1.2.0", "v1.1.1", "v0.0.1"}),
			expect: New(2, 0, 0),
		},
		"length of SemVerList is 0": {
			list:   SemVerList{},
			expect: New(0, 0, 0),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			latest := tc.list.Latest()
			assert.Equal(t, tc.expect, latest)
		})
	}
}
