package cmd

import (
	"fmt"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var deleteWorkspaceCmd = &cobra.Command{
	Use:   command.CommandWorkspace,
	Short: strings.DeleteWorkspaceDescription,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Add logic for deleting a workspace here")

		return nil
	},
}

func init() {
	deleteWorkspaceCmd.Flags().StringVarP(
		&workspaceKey,
		flag.FlagWorkspaceKey,
		flag.FlagK,
		"",
		strings.WorkspaceKeyDescription,
	)
}
