package versioning

import "fmt"

// SemVer replesents Semantic Versioning.
type SemVer struct {
	Major uint16
	Minor uint16
	Patch uint16
}

// New is factory of SemVer struct.
func New(major, minor, patch uint16) *SemVer {
	return &SemVer{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

// IncrementMajor increment major version.
func (s *SemVer) IncrementMajor() *SemVer {
	return &SemVer{
		Major: s.Major + 1,
		Minor: s.Minor,
		Patch: s.Patch,
	}
}

// IncrementMinor increment minor version.
func (s *SemVer) IncrementMinor() *SemVer {
	return &SemVer{
		Major: s.Major,
		Minor: s.Minor + 1,
		Patch: s.Patch,
	}
}

// IncrementPatch increment patch version.
func (s *SemVer) IncrementPatch() *SemVer {
	return &SemVer{
		Major: s.Major,
		Minor: s.Minor,
		Patch: s.Patch + 1,
	}
}

// ToString returs semantic versioning string.
func (s *SemVer) ToString() string {
	return fmt.Sprintf("v%d.%d.%d", s.Major, s.Minor, s.Patch)
}

// SemVerList implements sort.Interface.
type SemVerList []SemVer

// Len returs list length.
func (l SemVerList) Len() int {
	return len(l)
}

// Swap swaps value of index i,j.
func (l SemVerList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Less returns bool value whether i is less than j.
func (l SemVerList) Less(i, j int) bool {
	return l[i].Major < l[j].Major && l[i].Minor < l[j].Minor && l[i].Patch < l[j].Patch
}
