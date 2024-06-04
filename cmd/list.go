package cmd

import (
	"github.com/eljamo/bajira/cmd/boardcmd"
	"github.com/eljamo/bajira/cmd/workspacecmd"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:          command.CommandList,
	Short:        strings.ListDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(workspacecmd.ListWorkspacesCmd)
	listCmd.AddCommand(boardcmd.ListBoardsCmd)
}
