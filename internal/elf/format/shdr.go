package format

import (
	"fmt"
	"strings"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/ui"
)

func PrintSectionHeaders(ehdr *types.Elf64_Ehdr, shdr []types.Elf64_Shdr, data []byte) {
	var sb strings.Builder

	// Taking estimation
	sb.Grow(2048 + int(ehdr.E_shnum)*200)

	// ELF Program headers count and offset
	sb.WriteString(ui.Cyan.Sprintf("\n%s", "Section Headers: "))
	sb.WriteString(ui.Green.Sprintf("%d entries ", ehdr.E_shnum))
	sb.WriteString("(")
	sb.WriteString(ui.Yellow.Sprintf("offset %#x", ehdr.E_shoff))
	sb.WriteString(")\n\n")

	sb.WriteString(
		ui.Magenta.Sprintf("%-7s%-23s%-15s%-9s%-10s%-21s%-21s%s\n%49s%20s\n",
			"",
			"Name",
			"Type",
			"Flags",
			"Align",
			// "Offset",
			"Address",
			"Address",
			"PhysAddr",
			"FileSiz",
			"MemSiz",
		))

	strtab := types.GetStringTable(data, ehdr, shdr)

	// Table rows
	for i := range shdr {
		sh := &shdr[i]
		fmt.Fprintf(&sb, "%-16s%-32s%-24s%-18s%-19s%-30s\n",

			fmt.Sprintf("[%02s]", ui.Red.Sprintf("%02d", i)),

			ui.Cyan.Sprint(types.GetSectionNameFromStringTable(strtab, sh)),

			ui.Yellow.Sprint(types.GetSType(sh.Sh_type)),

			ui.Yellow.Sprintf("%#04x", sh.Sh_flags),
			ui.Yellow.Sprintf("%#04x", sh.Sh_addralign),

			// ui.Green.Sprintf("%#016x", sh.Sh_offset),

			ui.Green.Sprintf("%#016x", sh.Sh_addr),

		// ui.Green.Sprintf("%#016x", sh.P_paddr),

		// ui.Green.Sprintf("%#016x", sh.P_filesz),

		// ui.Green.Sprintf("%#016x", sh.P_memsz))
		)
	}
	fmt.Print(sb.String())
}
