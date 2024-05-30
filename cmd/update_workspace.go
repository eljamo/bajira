package cmd

import (
	"fmt"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var updateWorkspaceCmd = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.UpdateWorkspaceDescription,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Add logic for updating a workspace here")

		return nil
	},
}

func init() {
	updateWorkspaceCmd.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagK,
		"",
		strings.WorkspaceIdDescription,
	)
}
