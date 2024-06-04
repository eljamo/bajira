package boardcmd

import (
	"context"

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
		form, err := workspace.NewSelectWorkspaceForm(ctx, false, false)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		err = form.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		workspaceId = workspace.WorkspaceId
	}

	return nil
}

func archiveBoard(cmd *cobra.Command) error {
	msg, err := workspace.ArchiveWorkspace(cmd.Context(), workspaceId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
