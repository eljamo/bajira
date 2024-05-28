package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   command.CommandUpdate,
	Short: "Update a workspace, board, task, or subtask",
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateWorkspaceCmd)
}
