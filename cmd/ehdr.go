package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourpwnguy/strix/internal/elf"
	"github.com/yourpwnguy/strix/internal/elf/format"
)

// infoCmd displays ELF headers with colored output.
var infoCmd = &cobra.Command{
	Use:     "ehdr <file>",
	Short:   "Display complete ELF information from an ELF file",
	Example: "strix ehdr /bin/ls",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if strings.TrimSpace(args[0]) == "" {
			return errors.New("provide an argument")
		}

		// Initialises a new parser
		elfParser, err := elf.Open(args[0])
		if err != nil {
			return err
		}
		defer elfParser.Close()

		// Getting initiliased ELF Header
		hdr, err := elfParser.ELFHeader()
		if err != nil {
			return err
		}

		phdr, err := elfParser.ProgramHeaders()
		if err != nil {
			return err
		}

		// Pretty Print the ELF Header
		format.PrintELFHeader(hdr, phdr)
		return nil
	},
}
