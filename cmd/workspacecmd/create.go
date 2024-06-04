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

var CreateWorkspace = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.CreateWorkspaceDescription,
	SilenceUsage: true,
	RunE:         runCreateWorkspace,
}

func init() {
	CreateWorkspace.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	CreateWorkspace.Flags().StringVarP(
		&workspaceName,
		flag.FlagWorkspaceName,
		flag.FlagN,
		"",
		strings.WorkspaceNameDescription,
	)
}

func runCreateWorkspace(cmd *cobra.Command, args []string) error {
	err := parseCreateWorkspaceInput(cmd.Context())
	if err != nil {
		return err
	}

	return createWorkspace(cmd)
}

func parseCreateWorkspaceInput(ctx context.Context) error {
	if strings.StringIsEmpty(workspaceId) || strings.StringIsEmpty(workspaceName) {
		form, err := workspace.NewCreateWorkspaceForm(ctx, workspaceId, workspaceName)
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

func createWorkspace(cmd *cobra.Command) error {
	msg, err := workspace.CreateWorkspace(cmd.Context(), workspaceName, workspaceId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
