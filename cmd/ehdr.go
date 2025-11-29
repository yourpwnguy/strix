package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourpwnguy/strix/internal/elf/format"
	"github.com/yourpwnguy/strix/internal/elf/parser"
	"github.com/yourpwnguy/strix/internal/reader"
	"github.com/yourpwnguy/strix/internal/ui"
)

// infoCmd displays ELF headers with colored output.
var infoCmd = &cobra.Command{
	Use:     "ehdr <file>",
	Short:   "Display complete ELF information from an ELF file",
	Example: "strix ehdr /bin/ls",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(args[0]) == "" {
			fmt.Fprintf(os.Stderr, "%s %s\n",
				ui.ErrPrefix,
				ui.Red.Sprint("Provide an argument !"),
			)
			return
		}

		elfParser := parser.NewParser(&reader.MmapReader{})

		if err := elfParser.Load(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n",
				err,
			)
			return
		}

		defer elfParser.Close()

		// Getting initiliased ELF Header
		hdr, err := elfParser.ELFHeader()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n",
				err,
			)
			return
		}

		// Pretty Print the ELF Header
		format.PrintELFHeader(hdr)
	},
}
