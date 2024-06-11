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
	if err := setWorkspaceId(ctx); err != nil {
		return err
	}

	if err := setBoardId(ctx); err != nil {
		return err
	}

	if err := setNewBoardName(ctx); err != nil {
		return err
	}

	if err := setNewBoardIdAndName(ctx); err != nil {
		return err
	}

	return nil
}

func setWorkspaceId(ctx context.Context) error {
	if !strings.StringIsEmpty(workspaceId) {
		return nil
	}

	form, err := workspace.NewSelectWorkspaceForm(ctx, allWorkspaces, archivedWorkspaces)
	if err != nil {
		return errorconc.LocalizedError(err, "failed to initialize form")
	}

	err = form.Run()
	if err != nil {
		return errorconc.LocalizedError(err, "failed to run form")
	}

	workspaceId = workspace.WorkspaceId
	return nil
}

func setBoardId(ctx context.Context) error {
	if !strings.StringIsEmpty(boardId) {
		return nil
	}

	form, err := board.NewSelectBoardForm(ctx, workspaceId, allBoards, archivedBoards)
	if err != nil {
		return errorconc.LocalizedError(err, "failed to initialize form")
	}

	err = form.Run()
	if err != nil {
		return errorconc.LocalizedError(err, "failed to run form")
	}

	boardId = board.BoardId
	return nil
}

func setNewBoardName(ctx context.Context) error {
	if !strings.StringIsEmpty(newBoardName) {
		return nil
	}

	cfg, err := board.GetBoardConfig(ctx, workspaceId, boardId)
	if err != nil {
		return errorconc.LocalizedError(err, "failed to get board config")
	}

	newBoardName = cfg.Name
	return nil
}

func setNewBoardIdAndName(ctx context.Context) error {
	if !strings.StringIsEmpty(newBoardId) && !strings.StringIsEmpty(newBoardName) {
		return nil
	}

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
