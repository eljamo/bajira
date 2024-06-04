package cmd

import (
	"github.com/eljamo/bajira/cmd/boardcmd"
	"github.com/eljamo/bajira/cmd/workspacecmd"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:          command.CommandDelete,
	Short:        strings.DeleteDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(workspacecmd.DeleteWorkspace)
	deleteCmd.AddCommand(boardcmd.DeleteBoard)
}
