package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   command.CommandDelete,
	Short: strings.DeleteDescription,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteWorkspaceCmd)
}
