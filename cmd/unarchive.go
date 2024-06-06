package cmd

import (
	"github.com/eljamo/bajira/cmd/boardcmd"
	"github.com/eljamo/bajira/cmd/workspacecmd"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var unarchiveCmd = &cobra.Command{
	Use:          command.CommandUnarchive,
	Short:        strings.UnarchiveDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(unarchiveCmd)
	unarchiveCmd.AddCommand(workspacecmd.UnarchiveWorkspace)
	unarchiveCmd.AddCommand(boardcmd.UnarchiveBoard)
}
