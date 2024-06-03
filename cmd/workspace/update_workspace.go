package workspace

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var UpdateWorkspaceCmd = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.UpdateWorkspaceDescription,
	SilenceUsage: true,
	RunE:         runUpdateWorkspaceCmd,
}

func init() {
	UpdateWorkspaceCmd.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		"",
		"",
		strings.WorkspaceIdDescription,
	)
	UpdateWorkspaceCmd.Flags().StringVarP(
		&newWorkspaceId,
		flag.FlagNewWorkspaceId,
		"",
		"",
		strings.NewWorkspaceIdDescription,
	)
	UpdateWorkspaceCmd.Flags().StringVarP(
		&newWorkspaceName,
		flag.FlagNewWorkspaceName,
		"",
		"",
		strings.NewWorkspaceNameDescription,
	)
}

// steps
// once workspaceId is set, parse rest of the flags
// if all are set, try to update the workspace with new values

// if not prompt the user to enter new values

func runUpdateWorkspaceCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

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

	if strings.StringIsEmpty(newWorkspaceName) {
		cfg, err := workspace.GetWorkspaceConfig(ctx, workspaceId)
		if err != nil {
			return errorconc.LocalizedError(err, "failed to get workspace config")
		}

		newWorkspaceName = cfg.Name
	}

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
