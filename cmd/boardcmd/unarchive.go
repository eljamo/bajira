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

var UnarchiveBoard = &cobra.Command{
	Use:          command.CommandBoard,
	Short:        strings.UnarchiveBoardDescription,
	SilenceUsage: true,
	RunE:         runUnarchiveWorkspace,
}

func init() {
	UnarchiveBoard.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	UnarchiveBoard.Flags().StringVarP(
		&boardId,
		flag.FlagBoardId,
		flag.FlagB,
		"",
		strings.BoardIdDescription,
	)
	UnarchiveBoard.Flags().BoolVarP(&allWorkspaces, flag.FlagAllWorkspaces, flag.FlagW, false, strings.ListAllWorkspacesDescription)
	UnarchiveBoard.Flags().BoolVarP(&archivedWorkspaces, flag.FlagArchivedWorkspaces, flag.FlagU, false, strings.ListArchivedWorkspacesDescription)
	UnarchiveBoard.Flags().BoolVarP(&allBoards, flag.FlagAllBoards, flag.FlagA, false, strings.ListAllBoardsDescription)
	UnarchiveBoard.Flags().BoolVarP(&archivedBoards, flag.FlagArchivedBoards, flag.FlagR, false, strings.ListArchivedBoardsDescription)
}

func runUnarchiveWorkspace(cmd *cobra.Command, args []string) error {
	err := parseUnarchiveWorkspaceInput(cmd.Context())
	if err != nil {
		return err
	}

	return unarchiveWorkspace(cmd)
}

func parseUnarchiveWorkspaceInput(ctx context.Context) error {
	if strings.StringIsEmpty(workspaceId) {
		form, err := workspace.NewSelectWorkspaceForm(ctx, allWorkspaces, archivedWorkspaces)
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
		form, err := board.NewSelectBoardForm(ctx, workspaceId, allBoards, archivedBoards)
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

func unarchiveWorkspace(cmd *cobra.Command) error {
	msg, err := board.UnarchiveBoard(cmd.Context(), workspaceId, boardId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
