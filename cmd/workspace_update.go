package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateWorkspaceCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing workspace",
	Long:  `Update an existing workspace with new parameters.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add your logic for updating a workspace here")
	},
}
