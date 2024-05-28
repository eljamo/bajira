package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listWorkspaceCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces",
	Long:  `List all workspaces with their details.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add your logic for listing all workspaces here")
	},
}
