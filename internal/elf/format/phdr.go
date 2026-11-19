package format

import (
	"fmt"
	"strings"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/ui"
)

// PrintProgramHeaders displays the program header table in a formatted layout.
func PrintProgramHeaders(ehdr *types.Elf64_Ehdr, phdr []types.Elf64_Phdr, shdr []types.Elf64_Shdr, data []byte) {
	var sb strings.Builder

	// Basic estimation
	sb.Grow(2048 + int(ehdr.E_phnum)*200)

	// ELF Object Type
	sb.WriteString(ui.Cyan.Sprintf("\n%s", "ELF Type: "))
	sb.WriteString(ui.Green.Sprint(types.GetEType(ehdr.E_type, types.HasInterpreter(ehdr, phdr))))
	sb.WriteByte('\n')

	// ELF Entry point virtual address
	sb.WriteString(ui.Cyan.Sprintf("%s", "Entry: "))
	sb.WriteString(ui.Green.Sprintf("%#016x\n", ehdr.E_entry))

	// ELF Interpreter
	sb.WriteString(ui.Cyan.Sprintf("%s", "Interpreter: "))
	sb.WriteString(ui.Green.Sprint(types.GetInterpreter(ehdr, phdr, data)))
	sb.WriteString("\n\n")

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
	for i := range phdr {
		ph := &phdr[i]
		fmt.Fprintf(&sb, "%-16s%-25s%s  %-20s%-30s%-30s%s\n%69s%30s\n",

			fmt.Sprintf("[%02s]", ui.Red.Sprintf("%02d", i)),

			ui.Cyan.Sprint(types.GetPType(ph.P_type)),

			types.GetPFlags(ph.P_flags),

			ui.Yellow.Sprintf("%#04x", ph.P_align),

			ui.Green.Sprintf("%#016x", ph.P_offset),

			ui.Green.Sprintf("%#016x", ph.P_vaddr),

			ui.Green.Sprintf("%#016x", ph.P_paddr),

			ui.Green.Sprintf("%#016x", ph.P_filesz),

			ui.Green.Sprintf("%#016x", ph.P_memsz))
	}

	printSectionMapping(
		&sb,
		ehdr,
		phdr,
		shdr,
		data,
	)

	fmt.Print(sb.String())
}

// printSectionMapping prints which sections belong to each program segment.
// This shows how the linker packed sections into loadable segments.
func printSectionMapping(sb *strings.Builder, ehdr *types.Elf64_Ehdr, phdr []types.Elf64_Phdr, shdr []types.Elf64_Shdr, data []byte) {
	// Without a usable section name table there is nothing to map.
	strs := types.GetStringTable(data, ehdr, shdr)
	if len(strs) == 0 {
		return
	}

	sb.WriteString(ui.Cyan.Sprint("\nSection to Segment Mapping:\n"))
	sb.WriteString(ui.Magenta.Sprint(" Segment", " Sections...\n"))

	for i, ph := range phdr {
		sb.WriteString(ui.Yellow.Sprintf("  %02d     ", i))

		for _, sh := range shdr {
			if sh.Sh_type == types.SHT_NULL {
				continue
			}

			if types.SectionInSegment(sh, ph) {
				sb.WriteString(ui.Green.Sprint(types.Cstring(strs, sh.Sh_name)) + " ")
			}
		}
		sb.WriteByte('\n')
	}
}
