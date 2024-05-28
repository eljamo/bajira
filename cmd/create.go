package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   command.CommandCreate,
	Short: "Create a workspace, board, task, or subtask",
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createWorkspaceCmd)
}
