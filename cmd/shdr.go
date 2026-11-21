package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourpwnguy/strix/internal/elf"
	"github.com/yourpwnguy/strix/internal/elf/format"
)

var shdrCmd = &cobra.Command{
	Use:   "shdr [file]",
	Short: "Display section headers from an ELF file",
	Long:  "Parse and display the section headers from a given ELF file in a structured and colorized format.",
	Args:  cobra.ExactArgs(1),
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
		ehdr, err := elfParser.ELFHeader()
		if err != nil {
			return err
		}

		// Getting initiliased ELF Header
		shdr, err := elfParser.SectionHeaders()
		if err != nil {
			return err
		}

		// Pretty Print the ELF Header
		format.PrintSectionHeaders(ehdr, shdr, elfParser.Data())
		return nil
	},
}
