package cmd

import (
	"context"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var deleteWorkspaceCmd = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.DeleteWorkspaceDescription,
	SilenceUsage: true,
	RunE:         deleteWorkspace,
}

func init() {
	deleteWorkspaceCmd.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagK,
		"",
		strings.WorkspaceIdDescription,
	)
}

func deleteWorkspace(cmd *cobra.Command, args []string) error {
	err := parseDeleteWorkspaceInput(cmd.Context())
	if err != nil {
		return err
	}

	return removeWorkspace(cmd)
}

func parseDeleteWorkspaceInput(ctx context.Context) error {
	if workspaceId == "" {
		form, err := workspace.NewDeleteWorkspaceForm(ctx)
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

func removeWorkspace(cmd *cobra.Command) error {
	msg, err := workspace.DeleteWorkspace(cmd.Context(), workspaceId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
