package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteWorkspaceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing workspace",
	Long:  `Delete an existing workspace by its name or ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add your logic for deleting a workspace here")
	},
}
