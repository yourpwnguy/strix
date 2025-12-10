package format

import (
	"fmt"
	"strings"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/ui"
)

// Print ELF Header in colored and formatted way
func PrintELFHeader(ehdr *types.Elf64_Ehdr, phdr []types.Elf64_Phdr) {
	e_ident := &ehdr.E_ident

	var sb strings.Builder
	sb.Grow(2048)

	// Header
	sb.WriteString(ui.Bold.Sprint("ELF Header:\n"))

	// e_ident magic bytes
	sb.WriteString(ui.Cyan.Sprint("  Magic:   "))
	sb.WriteString(ui.Magenta.Sprintf("%02x %02x %02x %02x ",
		e_ident.FileIdn[0],
		e_ident.FileIdn[1],
		e_ident.FileIdn[2],
		e_ident.FileIdn[3],
	))

	// e_ident (rest of the bytes)
	fmt.Fprintf(&sb, "%02x %02x %02x %02x %02x ",
		e_ident.Ei_class,
		e_ident.Ei_data,
		e_ident.Ei_version,
		e_ident.Ei_osabi,
		e_ident.Ei_abiversion,
	)

	// e_ident - ei_pad bytes
	for _, b := range e_ident.Ei_pad {
		fmt.Fprintf(&sb, "%02x ", b)
	}
	sb.WriteByte('\n')

	// ELF Class Information
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Class:"))
	sb.WriteString(ui.Green.Sprint(types.GetEiClass(e_ident.Ei_class)))
	sb.WriteByte('\n')

	// ELF Data
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Data:"))
	sb.WriteString(ui.Green.Sprint(types.GetEiData(e_ident.Ei_class)))
	sb.WriteByte('\n')

	// ELF Version
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Version:"))
	sb.WriteString(ui.Green.Sprintf("%d (current)\n", e_ident.Ei_version))

	// ELF OS/ABI
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "OS/ABI:"))
	sb.WriteString(ui.Green.Sprintf(types.GetEiOSABI(e_ident.Ei_osabi)))
	sb.WriteByte('\n')

	// ELF ABI Version
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "ABI Version:"))
	sb.WriteString(ui.Green.Sprintf("%d\n", e_ident.Ei_abiversion))

	// ELF Object file type
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Type:"))
	sb.WriteString(ui.Yellow.Sprint(types.GetEType(ehdr.E_type, types.HasInterpreter(ehdr, phdr))))
	sb.WriteByte('\n')

	// ELF Machine - Architecture Information
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Machine:"))
	eMachine := types.GetEMachine(ehdr.E_machine)
	if strings.Contains((eMachine), "Advanced Micro Devices") {
		sb.WriteString(ui.Red.Sprint(eMachine))
	} else {
		sb.WriteString(ui.Green.Sprint(eMachine))
	}
	sb.WriteByte('\n')

	// ELF Object file version
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Version:"))
	sb.WriteString(ui.Green.Sprintf("%#x\n", ehdr.E_version))

	// ELF Entry point virtual address
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Entry Point Address:"))
	sb.WriteString(ui.Yellow.Sprintf("%#016x\n", ehdr.E_entry))

	// ELF Program header table file offset
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Start of program headers:"))
	sb.WriteString(ui.Green.Sprintf("%d (bytes into file)\n", ehdr.E_phoff))

	// ELF Section header table file offset
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Start of section headers:"))
	sb.WriteString(ui.Green.Sprintf("%d (bytes into file)\n", ehdr.E_shoff))

	// ELF Processor-specific flags
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Flags:"))
	sb.WriteString(ui.Green.Sprintf("%#x\n", ehdr.E_flags))

	// ELF Header size in bytes
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Size of this header:"))
	sb.WriteString(ui.Blue.Sprintf("%d (bytes)\n", ehdr.E_ehsize))

	// ELF Program header table entry size
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Size of program headers:"))
	sb.WriteString(ui.Blue.Sprintf("%d (bytes)\n", ehdr.E_phentsize))

	// ELF Program header table entry count
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Number of program headers:"))
	sb.WriteString(ui.Green.Sprintf("%d\n", ehdr.E_phnum))

	// ELF Section header table entry size
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Size of section headers:"))
	sb.WriteString(ui.Blue.Sprintf("%d (bytes)\n", ehdr.E_shentsize))

	// ELF Section header table entry count
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Number of section headers:"))
	sb.WriteString(ui.Green.Sprintf("%d\n", ehdr.E_shnum))

	// ELF Section header string table index
	sb.WriteString(ui.Cyan.Sprintf("  %-35s", "Section header string table index:"))
	sb.WriteString(ui.Green.Sprintf("%d\n", ehdr.E_shstrndx))

	// ELF Single output
	fmt.Print(sb.String())
}
