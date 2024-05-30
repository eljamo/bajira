package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:          command.CommandCreate,
	Short:        strings.CreateDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createWorkspaceCmd)
}
