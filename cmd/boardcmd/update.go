package boardcmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
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
		"",
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
}

func runUpdateBoard(cmd *cobra.Command, args []string) error {
	// ctx := cmd.Context()
	// get workspace id
	/*
		workspaceId, err := workspace.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}
	*/

	/*
		if strings.StringIsEmpty(workspaceId) {
			// select workspace
			// workspaceId =
		}
	*/

	/*
		if strings.StringIsEmpty(boardId) {
			// select board
			// boardId =
		}
	*/

	/*
		if strings.StringIsEmpty(newBoardName) {
			// get current board name
			cfg, err := board.GetBoardConfig(ctx, workspaceId, boardId)
			if err != nil {
				return errorconc.LocalizedError(err, "failed to get workspace config")
			}

			newBoardName = cfg.Name
		}
	*/

	/*
		if strings.StringIsEmpty(newBoardId) || strings.StringIsEmpty(newBoardName) {
			// input new board id and new board name
			// newBoardId =
			// newBoardName =
		}
	*/

	// 	return board.UpdateBoard(ctx, workspaceId, boardId, newBoardId, newBoardName)
	return nil
}
