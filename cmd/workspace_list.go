package cmd

import (
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var listWorkspaceCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces",
	Long:  `List all workspaces with their details.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		table, err := workspace.GenerateWorkspaceList()
		if err != nil {
			return err
		}

		cmd.Println(table)

		return nil
	},
}
