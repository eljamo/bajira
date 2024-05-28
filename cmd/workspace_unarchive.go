package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var unarchiveWorkspaceCmd = &cobra.Command{
	Use:   "unarchive",
	Short: "Unarchive a workspace",
	Long:  `Unarchive a previously archived workspace by its name or ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add your logic for unarchiving a workspace here")
	},
}
