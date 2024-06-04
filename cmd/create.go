package cmd

import (
	"github.com/eljamo/bajira/cmd/boardcmd"
	"github.com/eljamo/bajira/cmd/workspacecmd"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:          command.CommandCreate,
	Short:        strings.CreateDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(workspacecmd.CreateWorkspace)
	createCmd.AddCommand(boardcmd.CreateBoard)
}
