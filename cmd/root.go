package cmd

import (
	"context"
	"os"

	"github.com/eljamo/bajira/internal/consts"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var (
	all           bool
	archived      bool
	workspaceId   string
	workspaceName string
)

var rootCmd = &cobra.Command{
	Use:   consts.BajiraApplicationName,
	Short: strings.BajiraApplicationDescription,
}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
