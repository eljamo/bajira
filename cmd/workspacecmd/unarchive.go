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

var UnarchiveWorkspace = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.UnarchiveWorkspaceDescription,
	SilenceUsage: true,
	RunE:         runUnarchiveWorkspace,
}

func init() {
	UnarchiveWorkspace.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
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
		form, err := workspace.NewSelectWorkspaceForm(ctx, false, true)
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

func unarchiveWorkspace(cmd *cobra.Command) error {
	msg, err := workspace.UnarchiveWorkspace(cmd.Context(), workspaceId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
