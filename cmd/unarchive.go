package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var unarchiveCmd = &cobra.Command{
	Use:   command.CommandUnarchive,
	Short: strings.UnarchiveDescription,
}

func init() {
	rootCmd.AddCommand(unarchiveCmd)
	unarchiveCmd.AddCommand(unarchiveWorkspaceCmd)
}
