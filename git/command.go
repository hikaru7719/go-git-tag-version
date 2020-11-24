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

// Fetch function executes git fetch command
func Fetch() (string, error) {
	return executeGit("fetch")
}
