package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var listWorkspacesCmd = &cobra.Command{
	Use:   command.CommandWorkspaces,
	Short: strings.ListWorkspaceDescription,
	RunE: func(cmd *cobra.Command, args []string) error {
		table, err := workspace.GenerateWorkspaceList()
		if err != nil {
			return err
		}

		if table == nil {
			cmd.Println(strings.NoWorkspacesFound)
			return nil
		}

		cmd.Println(table)

		return nil
	},
}
