package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourpwnguy/strix/internal/elf"
	"github.com/yourpwnguy/strix/internal/elf/format"
	"github.com/yourpwnguy/strix/internal/ui"
)

var shdrCmd = &cobra.Command{
	Use:   "shdr [file]",
	Short: "Display section headers from an ELF file",
	Long:  "Parse and display the section headers from a given ELF file in a structured and colorized format.",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(args[0]) == "" {
			fmt.Fprintf(os.Stderr, "%s %s\n",
				ui.ErrPrefix,
				ui.Red.Sprint("Provide an argument !"),
			)
			return
		}

		// Initialises a new parser
		elfParser, err := elf.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s %s\n",
				ui.ErrPrefix,
				err,
			)
			return
		}
		defer elfParser.Close()

		// Getting initiliased ELF Header
		ehdr, err := elfParser.ELFHeader()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		// Getting initiliased ELF Header
		shdr, err := elfParser.SectionHeaders()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		// Pretty Print the ELF Header
		format.PrintSectionHeaders(ehdr, shdr, elfParser.Data())
	},
}
