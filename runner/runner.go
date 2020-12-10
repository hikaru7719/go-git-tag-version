package runner

import (
	"log"
	"sort"

	"github.com/hikaru7719/go-git-tag-version/git"
	"github.com/hikaru7719/go-git-tag-version/versioning"
	"golang.org/x/xerrors"
)

// Run is cordinator method.
func Run(major, minor, patch bool) error {
	git.Fetch()
	output, err := git.Tag()
	if err != nil {
		return xerrors.Errorf("git tag command failed #w", err)
	}
	semvers := versioning.From(git.Parse(output))
	sort.Sort(semvers)
	incrementer, err := versioning.NewIncrement(major, minor, patch)
	if err != nil {
		return xerrors.Errorf("NewIncrement failed major:%v, minor:%,v patch:%v #w", major, minor, patch, err)
	}
	newVersion := incrementer.Increment(&semvers[semvers.Len()-1])
	_, err = git.TagVersioning(newVersion.ToString())
	if err != nil {
		return xerrors.Errorf("git tab versioning failed %s #w", newVersion.ToString(), err)
	}
	log.Printf("new version is %s\n", newVersion.ToString())
	return nil
}
