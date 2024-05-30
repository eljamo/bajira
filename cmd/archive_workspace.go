package cmd

import (
	"fmt"

	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/flag"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var archiveWorkspaceCmd = &cobra.Command{
	Use:          command.CommandWorkspace,
	Short:        strings.ArchiveWorkspaceDescription,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Add logic for archiving a workspace here")

		return nil
	},
}

func init() {
	archiveWorkspaceCmd.Flags().StringVarP(
		&workspaceId,
		flag.FlagWorkspaceId,
		flag.FlagK,
		"",
		strings.WorkspaceIdDescription,
	)
}
