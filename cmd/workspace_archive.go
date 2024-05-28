package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var archiveWorkspaceCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive an existing workspace",
	Long:  `Archive an existing workspace by its name or ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add your logic for archiving a workspace here")
	},
}
