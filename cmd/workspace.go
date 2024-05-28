package cmd

import (
	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",
	Long:  `Create, delete, update, list, archive, and unarchive workspaces.`,
}

func init() {
	rootCmd.AddCommand(workspaceCmd)
	workspaceCmd.AddCommand(createWorkspaceCmd)
	workspaceCmd.AddCommand(deleteWorkspaceCmd)
	workspaceCmd.AddCommand(updateWorkspaceCmd)
	workspaceCmd.AddCommand(listWorkspaceCmd)
	workspaceCmd.AddCommand(archiveWorkspaceCmd)
	workspaceCmd.AddCommand(unarchiveWorkspaceCmd)
}
