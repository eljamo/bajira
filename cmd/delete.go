package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   command.CommandDelete,
	Short: "Delete a workspace, board, task, or subtask",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteWorkspaceCmd)
}
