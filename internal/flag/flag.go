package flag

import (
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

const (
	FlagH             string = "h"
	FlagK             string = "k"
	FlagN             string = "n"
	FlagHelp          string = "help"
	FlagWorkspaceName string = "workspace_name"
	FlagWorkspaceKey  string = "workspace_key"
)

func HelpOverride(cmd *cobra.Command) {
	cmd.Flags().BoolP(FlagHelp, FlagH, false, strings.HelpDescription)
}
