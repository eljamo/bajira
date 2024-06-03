package workspace

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var ListWorkspacesCmd = &cobra.Command{
	Use:          command.CommandWorkspaces,
	Short:        strings.ListWorkspaceDescription,
	SilenceUsage: true,
	RunE:         listWorkspaces,
}

func init() {
	ListWorkspacesCmd.Flags().BoolVarP(&all, flag.FlagAll, flag.FlagA, false, strings.ListAllWorkspacesDescription)
	ListWorkspacesCmd.Flags().BoolVarP(&archived, flag.FlagArchived, flag.FlagR, false, strings.ListArchivedWorkspacesDescription)
}

func listWorkspaces(cmd *cobra.Command, args []string) error {
	table, err := workspace.GenerateWorkspaceList(cmd.Context(), all, archived)
	if err != nil {
		return err
	}

	cmd.Println(table)

	return nil
}
