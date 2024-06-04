package boardcmd

import (
	"context"

	"github.com/eljamo/bajira/internal/board"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var DeleteBoard = &cobra.Command{
	Use:          command.CommandBoard,
	Short:        strings.DeleteBoardDescription,
	SilenceUsage: true,
	RunE:         runDeleteBoard,
}

func init() {
	DeleteBoard.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	DeleteBoard.Flags().StringVarP(
		&boardId,
		flag.FlagBoardId,
		flag.FlagB,
		"",
		strings.BoardIdDescription,
	)
	DeleteBoard.Flags().BoolVarP(&all, flag.FlagAll, flag.FlagA, false, strings.ListAllWorkspacesAndBoardsDescription)
	DeleteBoard.Flags().BoolVarP(&archived, flag.FlagArchived, flag.FlagR, false, strings.ListArchivedWorkspacesAndBoardsDescription)
}

func runDeleteBoard(cmd *cobra.Command, args []string) error {
	err := parseDeleteBoardInput(cmd.Context())
	if err != nil {
		return err
	}

	return deleteBoard(cmd)
}

func parseDeleteBoardInput(ctx context.Context) error {
	if strings.StringIsEmpty(workspaceId) {
		form, err := workspace.NewSelectWorkspaceForm(ctx, all, archived)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		err = form.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		workspaceId = workspace.WorkspaceId
	}

	if strings.StringIsEmpty(boardId) {
		form, err := board.NewSelectBoardForm(ctx, workspaceId, all, archived)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		err = form.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		boardId = board.BoardId
	}

	return nil
}

func deleteBoard(cmd *cobra.Command) error {
	msg, err := board.DeleteBoard(cmd.Context(), workspaceId, boardId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
