package main

import (
	"log"
	"os"

	"github.com/hikaru7719/go-git-tag-version/runner"
	"github.com/spf13/cobra"
)

var (
	major, minor, patch bool
	dryRun              bool
	cmd                 = &cobra.Command{
		Use:   "go-git-tag-version",
		Short: "go-git-tag-version suports versioning of git tag.",
		Long: `You can use git-tag-version with --major, --minor, --patch flags like yarn version. 
git-tag-version increments your git tag version depending on your flags.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runner.New(dryRun).Run(major, minor, patch)
		},
	}
)

func init() {
	cmd.PersistentFlags().BoolVarP(&major, "major", "a", false, "increment major version")
	cmd.PersistentFlags().BoolVarP(&minor, "minor", "b", false, "increment minor version")
	cmd.PersistentFlags().BoolVarP(&patch, "patch", "c", false, "increment patch version")
	cmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "execute dry-run mode")
}

// Execute run cli app.
func Execute() {
	if err := cmd.Execute(); err != nil {
		log.Printf("failed command execution: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
