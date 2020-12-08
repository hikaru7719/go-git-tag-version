package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	major, minor, patch bool
	cmd                 = &cobra.Command{
		Use:   "go-git-tag-version",
		Short: "go-git-tag-version suports versioning of git tag.",
		Long: `You can use git-tag-version with --major, --minor, --patch flags like yarn version. 
git-tag-version increments your git tag version depending on your flags.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(major, minor, patch)
		},
	}
)

func init() {
	cmd.PersistentFlags().BoolVarP(&major, "major", "a", false, "increment major version")
	cmd.PersistentFlags().BoolVarP(&minor, "minor", "b", false, "increment minor version")
	cmd.PersistentFlags().BoolVarP(&patch, "patch", "c", false, "increment patch version")
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
