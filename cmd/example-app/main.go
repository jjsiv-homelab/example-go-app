package main

import (
	"time"

	"github.com/jjsiv-homelab/example-go-app/cmd/example-app/commands"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339

	command := &cobra.Command{
		Use:   "example-app",
		Short: "An example application",
	}

	command.AddCommand(commands.ServeCommand())

	if err := command.Execute(); err != nil {
		log.Fatal().Err(err).Msg("command execution failed")
	}
}
