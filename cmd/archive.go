package cmd

import (
	"github.com/eljamo/bajira/cmd/boardcmd"
	"github.com/eljamo/bajira/cmd/workspacecmd"
	"github.com/eljamo/bajira/internal/command"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:          command.CommandArchive,
	Short:        strings.ArchiveDescription,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(archiveCmd)
	archiveCmd.AddCommand(workspacecmd.ArchiveWorkspace)
	archiveCmd.AddCommand(boardcmd.ArchiveBoard)
}
