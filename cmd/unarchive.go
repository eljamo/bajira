package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/spf13/cobra"
)

var unarchiveCmd = &cobra.Command{
	Use:   command.CommandUnarchive,
	Short: "Unarchive a workspace or a board",
}

func init() {
	rootCmd.AddCommand(unarchiveCmd)
	unarchiveCmd.AddCommand(unarchiveWorkspaceCmd)
}
