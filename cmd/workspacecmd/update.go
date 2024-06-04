package workspacecmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var UpdateWorkspace = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.UpdateWorkspaceDescription,
	SilenceUsage: true,
	RunE:         runUpdateBoard,
}

func init() {
	UpdateWorkspace.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagI,
		"",
		strings.WorkspaceIdDescription,
	)
	UpdateWorkspace.Flags().StringVarP(
		&newWorkspaceId,
		flag.FlagNewWorkspaceId,
		flag.FlagK,
		"",
		strings.NewWorkspaceIdDescription,
	)
	UpdateWorkspace.Flags().StringVarP(
		&newWorkspaceName,
		flag.FlagNewWorkspaceName,
		flag.FlagN,
		"",
		strings.NewWorkspaceNameDescription,
	)
}

func runUpdateBoard(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// set workspaceId if not set
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

	// if newWorkspaceName is not set, get the current one from workspaceId
	if strings.StringIsEmpty(newWorkspaceName) {
		cfg, err := workspace.GetWorkspaceConfig(ctx, workspaceId)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to get workspace config")
		}

		newWorkspaceName = cfg.Name
	}

	// if newWorkspaceId is not set or newWorkspaceName is not set, prompt the user to enter new values
	if strings.StringIsEmpty(newWorkspaceId) || strings.StringIsEmpty(newWorkspaceName) {
		form, err := workspace.NewUpdateWorkspaceForm(ctx, newWorkspaceId, newWorkspaceName)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		err = form.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		newWorkspaceId = workspace.WorkspaceId
		newWorkspaceName = workspace.WorkspaceName
	}

	return workspace.UpdateWorkspace(ctx, workspaceId, newWorkspaceId, newWorkspaceName)
}
