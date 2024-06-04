package boardcmd

import (
	"github.com/eljamo/bajira/internal/board"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var ListBoardsCmd = &cobra.Command{
	Use:          command.CommandBoards,
	Short:        strings.ListBoardsDescription,
	SilenceUsage: true,
	RunE:         listBoards,
}

func init() {
	ListBoardsCmd.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	ListBoardsCmd.Flags().BoolVarP(&all, flag.FlagAll, flag.FlagA, false, strings.ListAllWorkspacesDescription)
	ListBoardsCmd.Flags().BoolVarP(&archived, flag.FlagArchived, flag.FlagR, false, strings.ListArchivedWorkspacesDescription)
}

func listBoards(cmd *cobra.Command, args []string) error {
	wsid, err := workspace.GetWorkspaceId(cmd)
	if err != nil {
		return err
	}

	table, err := board.ListBoards(cmd.Context(), wsid, all, archived)
	if err != nil {
		return err
	}

	cmd.Println(table)

	return nil
}
