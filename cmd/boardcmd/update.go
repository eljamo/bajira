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

var UpdateBoard = &cobra.Command{
	Use:          command.CommandBoard,
	Short:        strings.UpdateBoardDescription,
	SilenceUsage: true,
	RunE:         runUpdateBoard,
}

func init() {
	UpdateBoard.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	UpdateBoard.Flags().StringVarP(
		&boardId,
		flag.FlagBoardId,
		flag.FlagB,
		"",
		strings.BoardIdDescription,
	)
	UpdateBoard.Flags().StringVarP(
		&newBoardId,
		flag.FlagNewBoardId,
		flag.FlagK,
		"",
		strings.NewBoardIdDescription,
	)
	UpdateBoard.Flags().StringVarP(
		&newBoardName,
		flag.FlagNewBoardName,
		flag.FlagN,
		"",
		strings.NewBoardNameDescription,
	)
	UpdateBoard.Flags().BoolVarP(&allWorkspaces, flag.FlagAllWorkspaces, flag.FlagW, false, strings.ListAllWorkspacesDescription)
	UpdateBoard.Flags().BoolVarP(&archivedWorkspaces, flag.FlagArchivedWorkspaces, flag.FlagC, false, strings.ListArchivedWorkspacesDescription)
	UpdateBoard.Flags().BoolVarP(&allBoards, flag.FlagAllBoards, flag.FlagA, false, strings.ListAllBoardsDescription)
	UpdateBoard.Flags().BoolVarP(&archivedBoards, flag.FlagArchivedBoards, flag.FlagR, false, strings.ListArchivedBoardsDescription)
}

func parseUpdateWorkspaceInput(ctx context.Context) error {
	// set workspaceId if not set
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

	// set boardId if not set
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

	// newBoardName is not set, get the current one from workspaceId and boardId
	if strings.StringIsEmpty(newBoardName) {
		cfg, err := board.GetBoardConfig(ctx, workspaceId, boardId)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to get board config")
		}

		newBoardName = cfg.Name
	}

	// newBoardId is not set or newBoardName is not set, prompt the user to enter new values
	if strings.StringIsEmpty(newBoardId) || strings.StringIsEmpty(newBoardName) {
		form, err := board.NewUpdateBoardForm(ctx, newBoardId, newBoardName)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		err = form.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		newBoardId = board.BoardId
		newBoardName = board.BoardName
	}

	return nil
}

func runUpdateBoard(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	err := parseUpdateWorkspaceInput(ctx)
	if err != nil {
		return err
	}

	return board.UpdateBoard(ctx, workspaceId, boardId, newBoardId, newBoardName)
}
