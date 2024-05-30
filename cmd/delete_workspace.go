package cmd

import (
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
	err := parseDeleteWorkspaceInput()
	if err != nil {
		return err
	}

	return removeWorkspace(cmd)
}

func parseDeleteWorkspaceInput() error {
	if workspaceName == "" {
		err := workspace.CreateWorkspaceForm.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to initialize form")
		}

		workspaceId = workspace.CreateWorkspaceId
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
