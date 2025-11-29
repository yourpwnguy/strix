package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "strix",
	Short: "A better way to parse ELF files",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(phdrCmd)
}
