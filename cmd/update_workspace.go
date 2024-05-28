package cmd

import (
	"fmt"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/spf13/cobra"
)

var updateWorkspaceCmd = &cobra.Command{
	Use:   command.CommandWorkspace,
	Short: "Update a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Add logic for updating a workspace here")

		return nil
	},
}

func init() {
	updateWorkspaceCmd.Flags().StringVarP(
		&workspaceKey,
		flag.FlagKey,
		flag.FlagK,
		"",
		"Key of the workspace to update",
	)
}
