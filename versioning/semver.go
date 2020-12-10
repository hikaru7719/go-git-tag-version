package versioning

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Incrementer is interface to encapsulate version increment.
type Incrementer interface {
	Increment(*SemVer) *SemVer
}

// MajorIncrement is struct for Incrementer.
type MajorIncrement struct{}

// Increment increments major version.
func (m *MajorIncrement) Increment(semver *SemVer) *SemVer {
	return semver.IncrementMajor()
}

// MinorIncrement is struct for Incrementer.
type MinorIncrement struct{}

// Increment increments minor version.
func (m *MinorIncrement) Increment(semver *SemVer) *SemVer {
	return semver.IncrementMinor()
}

// PatchIncrement is struct for Incrementer.
type PatchIncrement struct{}

// Increment increments patch version.
func (p *PatchIncrement) Increment(semver *SemVer) *SemVer {
	return semver.IncrementPatch()
}

// NewIncrement is factory method of Incrementer interface.
func NewIncrement(major, minor, patch bool) (Incrementer, error) {
	if major {
		return &MajorIncrement{}, nil
	}
	if minor {
		return &MinorIncrement{}, nil
	}
	if patch {
		return &PatchIncrement{}, nil
	}
	return nil, errors.New("all args(major, minor, patch) are false")
}

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
		Minor: 0,
		Patch: 0,
	}
}

// IncrementMinor increment minor version.
func (s *SemVer) IncrementMinor() *SemVer {
	return &SemVer{
		Major: s.Major,
		Minor: s.Minor + 1,
		Patch: 0,
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

// From ceonverts to []Semver from version strings.
func From(versions []string) SemVerList {
	var list SemVerList
	for _, s := range versions {
		result := strings.Split(strings.TrimPrefix(s, "v"), ".")
		if len(result) != 3 {
			continue
		}
		semver, err := from(result)
		if err != nil {
			log.Printf("cannot parse %s: %v", s, err)
			continue
		}
		list = append(list, semver)
	}
	return list
}

func from(list []string) (SemVer, error) {
	var major, minor, patch uint16
	for n, s := range list {
		num, err := strconv.Atoi(s)
		if err != nil {
			return SemVer{}, err
		}
		if num < 0 {
			return SemVer{}, errors.New("version includes negative value")
		}
		switch n {
		case 0:
			major = uint16(num)
		case 1:
			minor = uint16(num)
		case 2:
			patch = uint16(num)
		default:
			return SemVer{}, errors.New("unexpected error")
		}
	}
	return SemVer{Major: major, Minor: minor, Patch: patch}, nil
}

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
	return l[i].Major < l[j].Major || (l[i].Major == l[j].Major &&
		l[i].Minor < l[j].Minor) || (l[i].Major == l[j].Major &&
		l[i].Minor < l[j].Minor && l[i].Patch < l[j].Patch)
}
