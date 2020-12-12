package git

import (
	"log"
	"os/exec"
	"strings"
)

func executeGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	bytes, err := cmd.Output()
	if err != nil {
		log.Printf("error occurred in git command %v: %v\n", args, err)
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

// Tag executes git tag command and read all version.
func Tag() (string, error) {
	return executeGit("tag")
}

// DeleteTag deletes target version tag.
func DeleteTag(version string) (string, error) {
	return executeGit("tag", "-d", version)
}

// PushTag executes git push command and sync remote version and local version.
func PushTag(version string) (string, error) {
	return executeGit("git", "push", "origin", version)
}

// Parse parses version string.
func Parse(versions string) []string {
	return strings.Split(strings.TrimSuffix(versions, "\n"), "\n")
}
