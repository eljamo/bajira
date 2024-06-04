package workspacecmd

import (
	"context"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var DeleteWorkspace = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.DeleteWorkspaceDescription,
	SilenceUsage: true,
	RunE:         runDeleteBoard,
}

func init() {
	DeleteWorkspace.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	DeleteWorkspace.Flags().BoolVarP(&all, flag.FlagAll, flag.FlagA, false, strings.ListAllWorkspacesDescription)
	DeleteWorkspace.Flags().BoolVarP(&archived, flag.FlagArchived, flag.FlagR, false, strings.ListArchivedWorkspacesDescription)
}

func runDeleteBoard(cmd *cobra.Command, args []string) error {
	err := parseDeleteWorkspaceInput(cmd.Context())
	if err != nil {
		return err
	}

	return deleteWorkspace(cmd)
}

func parseDeleteWorkspaceInput(ctx context.Context) error {
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

	return nil
}

func deleteWorkspace(cmd *cobra.Command) error {
	msg, err := workspace.DeleteWorkspace(cmd.Context(), workspaceId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
