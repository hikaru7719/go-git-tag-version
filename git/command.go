package git

import (
	"log"
	"os/exec"
)

func executeGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	bytes, err := cmd.Output()
	if err != nil {
		log.Printf("error occured in git command %v: %v", args, err)
		return "", err
	}
	return string(bytes), nil
}

// Fetch executes git fetch command.
func Fetch() (string, error) {
	return executeGit("fetch")
}

// TagVersioning executes git tag version command.
func TagVersioning(version string) (string, error) {
	return executeGit("tag", version)
}

// Tag execute git tag command and read all version.
func Tag() (string, error) {
	return executeGit("tag")
}

// DeleteTag deletes target version tag.
func DeleteTag(version string) (string, error) {
	return executeGit("tag", "-d", version)
}

// PushTag execute git push command and sync remote version and local version.
func PushTag(version string) (string, error) {
	return executeGit("git", "push", "origin", version)
}
