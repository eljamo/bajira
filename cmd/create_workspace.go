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

var createWorkspaceCmd = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.CreateWorkspaceDescription,
	SilenceUsage: true,
	RunE:         createWorkspace,
}

func init() {
	createWorkspaceCmd.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagK,
		"",
		strings.WorkspaceIdDescription,
	)
	createWorkspaceCmd.Flags().StringVarP(
		&workspaceName,
		flag.FlagWorkspaceName,
		flag.FlagN,
		"",
		strings.WorkspaceNameDescription,
	)
}

func createWorkspace(cmd *cobra.Command, args []string) error {
	err := parseCreateWorkspaceInput(cmd.Context())
	if err != nil {
		return err
	}

	return setupWorkspace(cmd)
}

func parseCreateWorkspaceInput(ctx context.Context) error {
	if strings.CheckIfStringIsEmpty(workspaceName) {
		form, err := workspace.NewCreateWorkspaceForm(ctx)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		err = form.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		workspaceName = workspace.WorkspaceName
		workspaceId = workspace.WorkspaceId
	}

	return nil
}

func setupWorkspace(cmd *cobra.Command) error {
	msg, err := workspace.CreateWorkspace(cmd.Context(), workspaceName, workspaceId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
