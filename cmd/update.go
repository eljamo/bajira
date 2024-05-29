package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   command.CommandUpdate,
	Short: strings.UpdateDescription,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateWorkspaceCmd)
}
