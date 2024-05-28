package cmd

import (
	"fmt"

	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/form"
	"github.com/eljamo/bajira/internal/workspace"
	"github.com/spf13/cobra"
)

var (
	workspaceKey  string
	workspaceName string
)

var createWorkspaceCmd = &cobra.Command{
	Use:          "create",
	Short:        "Create a new workspace",
	Long:         `Create a new workspace with specified parameters.`,
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
		err := form.CreateWorkspaceForm.Run()
		if err != nil {
			return fmt.Errorf("failed to run form: %w", err)
		}

		workspaceName = form.CreateWorkspaceName
		workspaceKey = form.CreateWorkspaceKey
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
		flag.FlagKey,
		flag.FlagK,
		"",
		"Custom key for the workspace, if not provided a key will be generated from the name. If provided, the key will be formatted.",
	)
	createWorkspaceCmd.Flags().StringVarP(
		&workspaceName,
		flag.FlagName,
		flag.FlagN,
		"",
		"Name of the workspace",
	)
}
