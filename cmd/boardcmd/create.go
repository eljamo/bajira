package boardcmd

import (
	"github.com/eljamo/bajira/internal/board"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var CreateBoard = &cobra.Command{
	Use:          command.CommandBoard,
	Short:        strings.CreateBoardDescription,
	SilenceUsage: true,
	RunE:         runCreateBoard,
}

func init() {
	CreateBoard.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	CreateBoard.Flags().StringVarP(
		&boardId,
		flag.FlagBoardId,
		flag.FlagB,
		"",
		strings.BoardIdDescription,
	)
	CreateBoard.Flags().StringVarP(
		&boardName,
		flag.FlagBoardName,
		flag.FlagN,
		"",
		strings.BoardNameDescription,
	)
}

func runCreateBoard(cmd *cobra.Command, args []string) error {
	err := parseCreateBoardInput(cmd)
	if err != nil {
		return err
	}

	return createBoard(cmd)
}

func parseCreateBoardInput(cmd *cobra.Command) error {
	ctx := cmd.Context()
	wsid, err := workspace.GetWorkspaceId(cmd)
	if err != nil {
		return err
	}

	if strings.StringIsEmpty(wsid) {
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

	if strings.StringIsEmpty(boardId) || strings.StringIsEmpty(boardName) {
		form, err := board.NewCreateBoardForm(ctx, boardId, boardName)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		err = form.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		boardId = board.BoardId
		boardName = board.BoardName
	}

	return nil
}

func createBoard(cmd *cobra.Command) error {
	msg, err := board.CreateBoard(cmd.Context(), workspaceId, boardId, boardName)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
