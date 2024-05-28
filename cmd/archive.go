package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   command.CommandArchive,
	Short: "Archive a workspace or a board",
}

func init() {
	rootCmd.AddCommand(archiveCmd)
	archiveCmd.AddCommand(archiveWorkspaceCmd)
}
