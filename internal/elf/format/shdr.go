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

	// Section headers count and offset
	sb.WriteString(ui.Cyan.Sprintf("\n%s", "Section Headers: "))
	sb.WriteString(ui.Green.Sprintf("%d entries ", ehdr.E_shnum))
	sb.WriteString("(")
	sb.WriteString(ui.Yellow.Sprintf("offset %#x", ehdr.E_shoff))
	sb.WriteString(")\n\n")

	sb.WriteString(
		ui.Magenta.Sprintf("%-7s%-24s%-16s%-10s%-10s%s\n",
			"",
			"Name",
			"Type",
			"Flags",
			"Align",
			"Address",
		))

	strtab := types.GetStringTable(data, ehdr, shdr)

	// Table rows
	for i := range shdr {
		sh := &shdr[i]
		fmt.Fprintf(&sb, "%s%s%s%s%s%s\n",
			ui.Red.Sprintf("%-7s", fmt.Sprintf("[%02d]", i)),
			ui.Cyan.Sprintf("%-24s", types.GetSectionNameFromStringTable(strtab, sh)),
			ui.Yellow.Sprintf("%-16s", types.GetSType(sh.Sh_type)),
			ui.Yellow.Sprintf("%-10s", fmt.Sprintf("%#04x", sh.Sh_flags)),
			ui.Yellow.Sprintf("%-10s", fmt.Sprintf("%#04x", sh.Sh_addralign)),
			ui.Green.Sprintf("%#016x", sh.Sh_addr),
		)
	}
	fmt.Print(sb.String())
}
