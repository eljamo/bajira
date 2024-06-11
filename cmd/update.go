package cmd

import (
	"github.com/eljamo/bajira/cmd/boardcmd"
	"github.com/eljamo/bajira/cmd/workspacecmd"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:          command.CommandUpdate,
	Short:        strings.UpdateDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(workspacecmd.UpdateWorkspace)
	updateCmd.AddCommand(boardcmd.UpdateBoard)
}
