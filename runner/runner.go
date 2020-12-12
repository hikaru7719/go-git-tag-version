package runner

import (
	"log"
	"strings"

	"github.com/hikaru7719/go-git-tag-version/git"
	"github.com/hikaru7719/go-git-tag-version/versioning"
	"golang.org/x/xerrors"
)

// New is factory method for Runner.
func New(dryRun bool) *Runner {
	return &Runner{
		git: git.New(dryRun),
	}
}

// Runner is cordinator struct.
type Runner struct {
	git git.IGit
}

// Run do versioning up based on cli flag.
func (r *Runner) Run(major, minor, patch bool) error {
	// nolint:errcheck
	r.git.Fetch()
	output, err := r.git.Tag()
	if err != nil {
		return xerrors.Errorf("git tag command failed %w", err)
	}

	semvers := versioning.From(Parse(output))
	incrementer, err := versioning.NewIncrement(major, minor, patch)
	if err != nil {
		return xerrors.Errorf("NewIncrement failed major:%v, minor:%v patch:%v %w", major, minor, patch, err)
	}

	newVersion := incrementer.Increment(semvers.Latest())
	_, err = r.git.TagVersioning(newVersion.ToString())
	if err != nil {
		return xerrors.Errorf("git tab versioning failed %s %w", newVersion.ToString(), err)
	}

	log.Printf("new version is %s\n", newVersion.ToString())
	return nil
}

// Parse parses version string.
func Parse(versions string) []string {
	return strings.Split(strings.TrimSuffix(versions, "\n"), "\n")
}
