package commands

import (
	"fmt"

	"github.com/jjsiv-homelab/example-go-app/internal/version"
	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "version [flags]",
		Short: "Print program version",
		Run: func(c *cobra.Command, args []string) {
			v := version.Version()
			fmt.Printf("Version: %s\nCommit: %s\nGoVersion: %s\n", v.Version, v.CommitSHA, v.GoVersion)
		},
	}

	return command
}
