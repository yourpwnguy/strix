package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourpwnguy/strix/internal/elf"
	"github.com/yourpwnguy/strix/internal/elf/format"
)

// phdrCmd displays ELF program headers with colored output.
// Program headers describe segments for runtime loading (PT_LOAD, PT_DYNAMIC, etc).
var phdrCmd = &cobra.Command{
	Use:     "phdr [file]",
	Short:   "Display program headers from an ELF file",
	Example: "strix phdr /bin/ls",
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

		ehdr, err := elfParser.ELFHeader()
		if err != nil {
			return err
		}

		phdr, err := elfParser.ProgramHeaders()
		if err != nil {
			return err
		}

		shdr, err := elfParser.SectionHeaders()
		if err != nil {
			return err
		}

		// Pretty Print the Program Headers
		// Maybe I shoudn't pass it like this instead bind this into a struct
		// TODO: Make elf a type and implement a one pass parsing function
		// TODO: Can we make a formatter interface, could use text or json
		format.PrintProgramHeaders(ehdr, phdr, shdr, elfParser.Data())
		return nil
	},
}
