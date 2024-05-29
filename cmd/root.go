package cmd

import (
	"os"

	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.BajiraApplicationName,
	Short: strings.BajiraApplicationDescription,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
