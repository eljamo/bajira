package cmd

import (
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   command.CommandList,
	Short: strings.ListDescription,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listWorkspacesCmd)
}
