package cmd

import (
	"fmt"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/spf13/cobra"
)

var unarchiveWorkspaceCmd = &cobra.Command{
	Use:   command.CommandWorkspace,
	Short: "Unarchive a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Add logic for unarchiving a workspace here")

		return nil
	},
}

func init() {
	unarchiveWorkspaceCmd.Flags().StringVarP(
		&workspaceKey,
		flag.FlagKey,
		flag.FlagK,
		"",
		"Key of the workspace to unarchive",
	)
}
