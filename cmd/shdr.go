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

		elfParser := parser.NewParser(&reader.MmapReader{})

		if err := elfParser.Load(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "%s %s\n",
				ui.ErrPrefix,
				err,
			)
			return
		}
		defer elfParser.Close()

		fmt.Printf("%+v", elfParser)
		// Getting initiliased ELF Header
		hdr, err := elfParser.ELFHeader()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		// phdr, err := elfParser.ProgramHeaders()
		// if err != nil {
		// 	fmt.Fprint(os.Stderr, err)
		// 	return
		// }

		fmt.Printf("%+v", elfParser)
		// Pretty Print the ELF Header
		format.PrintELFHeader(hdr)
	},
}
