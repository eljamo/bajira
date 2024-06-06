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

var ArchiveBoard = &cobra.Command{
	Use:          command.CommandBoard,
	Short:        strings.ArchiveBoardDescription,
	SilenceUsage: true,
	RunE:         runArchiveBoard,
}

func init() {
	ArchiveBoard.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	ArchiveBoard.Flags().StringVarP(
		&boardId,
		flag.FlagBoardId,
		flag.FlagB,
		"",
		strings.BoardIdDescription,
	)
	ArchiveBoard.Flags().BoolVarP(&allWorkspaces, flag.FlagAllWorkspaces, flag.FlagW, false, strings.ListAllWorkspacesDescription)
	ArchiveBoard.Flags().BoolVarP(&archivedWorkspaces, flag.FlagArchivedWorkspaces, flag.FlagU, false, strings.ListArchivedWorkspacesDescription)
	ArchiveBoard.Flags().BoolVarP(&allBoards, flag.FlagAllBoards, flag.FlagA, false, strings.ListAllBoardsDescription)
	ArchiveBoard.Flags().BoolVarP(&archivedBoards, flag.FlagArchivedBoards, flag.FlagR, false, strings.ListArchivedBoardsDescription)
}

func runArchiveBoard(cmd *cobra.Command, args []string) error {
	err := parseArchiveBoardInput(cmd.Context())
	if err != nil {
		return err
	}

	return archiveBoard(cmd)
}

func parseArchiveBoardInput(ctx context.Context) error {
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

func archiveBoard(cmd *cobra.Command) error {
	msg, err := board.ArchiveBoard(cmd.Context(), workspaceId, boardId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
