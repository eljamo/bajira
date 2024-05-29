package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var (
	all      bool
	archived bool
)

var listWorkspacesCmd = &cobra.Command{
	Use:   command.CommandWorkspaces,
	Short: strings.ListWorkspaceDescription,
	RunE: func(cmd *cobra.Command, args []string) error {
		table, err := workspace.GenerateWorkspaceList(all, archived)
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

func init() {
	listWorkspacesCmd.Flags().BoolVarP(&all, flag.FlagAll, flag.FlagA, false, strings.ListAllWorkspacesDescription)
	listWorkspacesCmd.Flags().BoolVarP(&archived, flag.FlagArchived, flag.FlagR, false, strings.ListArchivedWorkspacesDescription)
}
