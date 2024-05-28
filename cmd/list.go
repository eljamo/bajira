package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   command.CommandList,
	Short: "List all workspaces, boards, tasks, or subtasks",
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listWorkspacesCmd)
}

var listWorkspacesCmd = &cobra.Command{
	Use:   command.CommandWorkspaces,
	Short: "List all workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		table, err := workspace.GenerateWorkspaceList()
		if err != nil {
			return err
		}

		if table == nil {
			cmd.Println("No workspaces found")
			return nil
		}

		cmd.Println(table)

		return nil
	},
}
