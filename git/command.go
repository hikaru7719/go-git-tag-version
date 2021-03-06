package git

import (
	"log"
	"os/exec"
)

// New is factory method for IGit.
func New(dryRun bool) IGit {
	if dryRun {
		return ImitationGit{}
	}
	return Git{}
}

// IGit is interface for git commands.
type IGit interface {
	Fetch() (string, error)
	Tag() (string, error)
	TagVersioning(version string) (string, error)
	DeleteTag(version string) (string, error)
	PushTag(version string) (string, error)
}

func executeGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	bytes, err := cmd.Output()
	if err != nil {
		log.Printf("error occurred in git command %v: %v\n", args, err)
		return "", err
	}
	return string(bytes), nil
}

func fetch() (string, error) {
	return executeGit("fetch")
}

func tagVersioning(version string) (string, error) {
	return executeGit("tag", version)
}

func tag() (string, error) {
	return executeGit("tag")
}

func deleteTag(version string) (string, error) {
	return executeGit("tag", "-d", version)
}

func pushTag(version string) (string, error) {
	return executeGit("git", "push", "origin", version)
}

// Git is a implementation of IGit.
type Git struct{}

// Fetch executes git fetch command.
func (g Git) Fetch() (string, error) {
	return fetch()
}

// TagVersioning executes git tag version command.
func (g Git) TagVersioning(version string) (string, error) {
	return tagVersioning(version)
}

// Tag executes git tag command and read all version.
func (g Git) Tag() (string, error) {
	return tag()
}

// DeleteTag deletes target version tag.
func (g Git) DeleteTag(version string) (string, error) {
	return deleteTag(version)
}

// PushTag executes git push command and sync remote version and local version.
func (g Git) PushTag(version string) (string, error) {
	return pushTag(version)
}

// ImitationGit doesn't making changes.
type ImitationGit struct{}

// Fetch executes same with Git.
func (g ImitationGit) Fetch() (string, error) {
	return fetch()
}

// TagVersioning do nothing.
func (g ImitationGit) TagVersioning(version string) (string, error) {
	log.Printf("dry-run TagVersioning method %s", version)
	return "", nil
}

// Tag executes same with Git.
func (g ImitationGit) Tag() (string, error) {
	return tag()
}

// DeleteTag do nothing.
func (g ImitationGit) DeleteTag(version string) (string, error) {
	log.Printf("dry-run DeleteTag method %s", version)
	return "", nil
}

// PushTag do nothing.
func (g ImitationGit) PushTag(version string) (string, error) {
	log.Printf("dry-run PushTag method %s", version)
	return "", nil
}
