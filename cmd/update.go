package cmd

import (
	"github.com/eljamo/bajira/cmd/workspace"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:          command.CommandUpdate,
	Short:        strings.UpdateDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(workspace.UpdateWorkspaceCmd)
}
