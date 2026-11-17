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

// phdrCmd displays ELF program headers with colored output.
// Program headers describe segments for runtime loading (PT_LOAD, PT_DYNAMIC, etc).
var phdrCmd = &cobra.Command{
	Use:     "phdr [file]",
	Short:   "Display program headers from an ELF file",
	Example: "strix phdr /bin/ls",
	Args:    cobra.ExactArgs(1),
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

		ehdr, err := elfParser.ELFHeader()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		phdr, err := elfParser.ProgramHeaders()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		shdr, err := elfParser.SectionHeaders()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		// Pretty Print the Program Headers
		// Maybe I shoudn't pass it like this instead bind this into a struct
		// TODO: Make elf a type and implement a one pass parsing function
		// TODO: Can we make a formatter interface, could use text or json
		format.PrintProgramHeaders(ehdr, phdr, shdr, elfParser.Data())
	},
}
