package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourpwnguy/strix/pkg/elf"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show ELF header information",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Provide a file to inspect")
		}

		filepath := args[0]

		hdr, cleanup, err := elf.ReadELFHeader(filepath)
		defer cleanup()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		fmt.Printf("%+v", hdr)
	},
}
