package format

import (
	"fmt"
	"strings"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/ui"
)

// PrintProgramHeaders displays the program header table in a formatted layout.
func PrintProgramHeaders(ehdr *types.Elf64_Ehdr, phdr []types.Elf64_Phdr) {
	var sb strings.Builder
	sb.Grow(2048 + int(ehdr.E_phnum)*200)

	// ELF Object Type
	sb.WriteString(ui.Cyan.Sprintf("\n%s", "ELF Type: "))
	sb.WriteString(ui.Green.Sprint(types.GetEType(ehdr.E_type)))
	sb.WriteByte('\n')

	// ELF Entry point virtual address
	sb.WriteString(ui.Cyan.Sprintf("%s", "Entry: "))
	sb.WriteString(ui.Green.Sprintf("%#016x\n", ehdr.E_entry))

	// ELF Program headers count and offset
	sb.WriteString(ui.Cyan.Sprintf("%s", "Program Headers: "))
	sb.WriteString(ui.Green.Sprintf("%d entries ", ehdr.E_phnum))
	sb.WriteString("(")
	sb.WriteString(ui.Yellow.Sprintf("offset %#x", ehdr.E_phoff))
	sb.WriteString(")\n\n")

	sb.WriteString(
		ui.Magenta.Sprintf("%-7s%-16s%-8s%-11s%-21s%-21s%s\n%49s%20s\n",
			"",
			"Type",
			"Flags",
			"Align",
			"Offset",
			"VirtAddr",
			"PhysAddr",
			"FileSiz",
			"MemSiz",
		))

	// Table rows
	for i := uint16(0); i < ehdr.E_phnum; i++ {
		ph := &phdr[i]
		sb.WriteString(
			fmt.Sprintf("%-16s%-25s%s  %-20s%-30s%-30s%s\n%69s%30s\n",
				// Count
				fmt.Sprintf("[%02s]", ui.Red.Sprintf("%02d", i)),
				// Type
				ui.Cyan.Sprint(types.GetPType(ph.P_type)),
				// Flags
				types.GetPFlags(ph.P_flags),
				// Align
				ui.Yellow.Sprintf("%#04x", ph.P_align),
				// Offset
				ui.Green.Sprintf("%#016x", ph.P_offset),
				// VirtAddr
				ui.Green.Sprintf("%#016x", ph.P_vaddr),
				// PhysAddr
				ui.Green.Sprintf("%#016x", ph.P_paddr),
				// File Size
				ui.Green.Sprintf("%#016x", ph.P_filesz),
				// Memory Size
				ui.Green.Sprintf("%#016x", ph.P_memsz),
			))
	}
	fmt.Print(sb.String())
}
