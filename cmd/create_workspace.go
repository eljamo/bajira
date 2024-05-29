package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var (
	workspaceKey  string
	workspaceName string
)

var createWorkspaceCmd = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.CreateWorkspaceDescription,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := parseInput()
		if err != nil {
			return err
		}

		return createWorkspace(cmd)
	},
}

func parseInput() error {
	if workspaceName == "" {
		err := workspace.CreateWorkspaceForm.Run()
		if err != nil {
			return errorconc.LocalizedError(err, "failed to run form")
		}

		workspaceName = workspace.CreateWorkspaceName
		workspaceKey = workspace.CreateWorkspaceKey
	}

	return nil
}

func createWorkspace(cmd *cobra.Command) error {
	msg, err := workspace.CreateWorkspace(workspaceName, workspaceKey)
	if err != nil {
		return err
	}

	cmd.Println(msg)

	return nil
}

func init() {
	createWorkspaceCmd.Flags().StringVarP(
		&workspaceKey,
		flag.FlagWorkspaceKey,
		flag.FlagK,
		"",
		strings.WorkspaceKeyDescription,
	)
	createWorkspaceCmd.Flags().StringVarP(
		&workspaceName,
		flag.FlagWorkspaceName,
		flag.FlagN,
		"",
		strings.WorkspaceNameDescription,
	)
}
