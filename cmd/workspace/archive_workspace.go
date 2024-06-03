package workspace

import (
	"context"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var ArchiveWorkspaceCmd = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.ArchiveWorkspaceDescription,
	SilenceUsage: true,
	RunE:         runArchiveWorkspace,
}

func init() {
	ArchiveWorkspaceCmd.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagK,
		"",
		strings.WorkspaceIdDescription,
	)
}

func runArchiveWorkspace(cmd *cobra.Command, args []string) error {
	err := parseArchiveWorkspaceInput(cmd.Context())
	if err != nil {
		return err
	}

	return archiveWorkspace(cmd)
}

func parseArchiveWorkspaceInput(ctx context.Context) error {
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

func archiveWorkspace(cmd *cobra.Command) error {
	msg, err := workspace.ArchiveWorkspace(cmd.Context(), workspaceId)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}
