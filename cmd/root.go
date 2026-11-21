package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yourpwnguy/strix/internal/ui"
)

// Reepresets the default version
var Version string = "dev"

var rootCmd = &cobra.Command{
	Use:          "strix",
	Short:        "A better and faster way to parse ELF files",
	Version:      Version,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetErrPrefix(ui.ErrPrefix)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(phdrCmd)
	rootCmd.AddCommand(shdrCmd)
}
