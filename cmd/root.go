package cmd

import "github.com/spf13/cobra"

// Reepresets the default version
var Version string = "dev"

var rootCmd = &cobra.Command{
	Use:     "strix",
	Short:   "A better and faster way to parse ELF files",
	Version: Version,
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(phdrCmd)
	rootCmd.AddCommand(shdrCmd)
}
