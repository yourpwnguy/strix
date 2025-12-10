package types

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/yourpwnguy/strix/internal/ui"
)

// GetEiClass returns a human-readable string representation of the EI_CLASS field,
// indicating whether the ELF file is a 32-bit or 64-bit executable.
func GetEiClass(ei_class uint8) string {
	switch ei_class {
	case ELFCLASS32:
		return "ELF32"
	case ELFCLASS64:
		return "ELF64"
	default:
		return "<Unknown>"
	}
}

// GetEiData returns a human-readable string representation of the EI_DATA field,
// indicating the data encoding (endianness) of the ELF file's processor-specific data.
func GetEiData(ei_data uint8) string {
	switch ei_data {
	case ELFDATA2LSB:
		return "2's complement, Little endian"
	case ELFDATA2MSB:
		return "2's complement, Big endian"
	default:
		return "<Unknown>"
	}
}

// GetEiOSABI returns a human-readable string representation of the EI_OSABI field,
// identifying the target operating system and ABI for which the ELF file is intended.
func GetEiOSABI(ei_osabi uint8) string {
	switch ei_osabi {
	case ELFOSABI_NONE:
		return "UNIX - System V"
	case ELFOSABI_LINUX:
		return "UNIX - GNU"
	case ELFOSABI_FREEBSD:
		return "UNIX - FREEBSD"
	case ELFOSABI_OPENBSD:
		return "UNIX - OPENBSD"
	case ELFOSABI_STANDALONE:
		return "Standalone"
	default:
		return "<Unknown>"
	}
}

// GetEType returns a human-readable string representation of the e_type field,
// identifying the object file type (executable, relocatable, shared object, etc.).
func GetEType(e_type uint16) string {
	switch e_type {
	case ET_NONE:
		return "NONE (No file type)"
	case ET_REL:
		return "REL (Relocatable file)"
	case ET_EXEC:
		return "EXEC (Executable file)"
	case ET_DYN:
		return "DYN (Postition-Independent Executable file)"
	case ET_CORE:
		return "CORE (Core file)"
	default:
		return "<Unknown>"
	}
}

// GetEMachine returns a human-readable string representation of the e_machine field,
// identifying the target instruction set architecture required by the ELF file.
func GetEMachine(e_machine uint16) string {
	switch e_machine {
	case EM_NONE:
		return "No machine"
	case EM_M32:
		return "AT&T WE 32100"
	case EM_SPARC:
		return "SUN SPARC"
	case EM_386:
		return "Intel 80386"
	case EM_68K:
		return "Motorola m68k family"
	case EM_88K:
		return "Motorola m88k family"
	case EM_IAMCU:
		return "Intel MCU"
	case EM_860:
		return "Intel 80860"
	case EM_MIPS:
		return "MIPS R3000 big-endian"
	case EM_S370:
		return "IBM System/370"
	case EM_MIPS_RS3_LE:
		return "MIPS R3000 little-endian"
	case EM_PARISC:
		return "HPPA"
	case EM_VPP500:
		return "Fujitsu VPP500"
	case EM_SPARC32PLUS:
		return "Sun's v8plus"
	case EM_960:
		return "Intel 80960"
	case EM_PPC:
		return "PowerPC"
	case EM_PPC64:
		return "PowerPC 64-bit"
	case EM_S390:
		return "IBM S390"
	case EM_SPU:
		return "IBM SPU/SPC"
	case EM_V800:
		return "NEC V800 series"
	case EM_FR20:
		return "Fujitsu FR20"
	case EM_RH32:
		return "TRW RH-32"
	case EM_RCE:
		return "Motorola RCE"
	case EM_ARM:
		return "ARM"
	case EM_FAKE_ALPHA:
		return "Digital Alpha"
	case EM_SH:
		return "Hitachi SH"
	case EM_SPARCV9:
		return "SPARC v9 64-bit"
	case EM_TRICORE:
		return "Siemens Tricore"
	case EM_ARC:
		return "Argonaut RISC Core"
	case EM_H8_300:
		return "Hitachi H8/300"
	case EM_H8_300H:
		return "Hitachi H8/300H"
	case EM_H8S:
		return "Hitachi H8S"
	case EM_H8_500:
		return "Hitachi H8/500"
	case EM_IA_64:
		return "Intel Merced"
	case EM_MIPS_X:
		return "Stanford MIPS-X"
	case EM_COLDFIRE:
		return "Motorola Coldfire"
	case EM_68HC12:
		return "Motorola M68HC12"
	case EM_MMA:
		return "Fujitsu MMA Multimedia Accelerator"
	case EM_PCP:
		return "Siemens PCP"
	case EM_NCPU:
		return "Sony nCPU embedded RISC"
	case EM_NDR1:
		return "Denso NDR1 microprocessor"
	case EM_STARCORE:
		return "Motorola Start*Core processor"
	case EM_ME16:
		return "Toyota ME16 processor"
	case EM_ST100:
		return "STMicroelectronic ST100 processor"
	case EM_TINYJ:
		return "Advanced Logic Corp. Tinyj emb.fam"
	case EM_X86_64:
		return "Advanced Micro Devices X86-64"
	case EM_PDSP:
		return "Sony DSP Processor"
	case EM_PDP10:
		return "Digital PDP-10"
	case EM_PDP11:
		return "Digital PDP-11"
	case EM_FX66:
		return "Siemens FX66 microcontroller"
	case EM_ST9PLUS:
		return "STMicroelectronics ST9+ 8/16 mc"
	case EM_ST7:
		return "STmicroelectronics ST7 8 bit mc"
	case EM_68HC16:
		return "Motorola MC68HC16 microcontroller"
	case EM_68HC11:
		return "Motorola MC68HC11 microcontroller"
	case EM_68HC08:
		return "Motorola MC68HC08 microcontroller"
	case EM_68HC05:
		return "Motorola MC68HC05 microcontroller"
	case EM_SVX:
		return "Silicon Graphics SVx"
	case EM_ST19:
		return "STMicroelectronics ST19 8 bit mc"
	case EM_VAX:
		return "Digital VAX"
	case EM_CRIS:
		return "Axis Communications 32-bit emb.proc"
	case EM_JAVELIN:
		return "Infineon Technologies 32-bit emb.proc"
	case EM_FIREPATH:
		return "Element 14 64-bit DSP Processor"
	case EM_ZSP:
		return "LSI Logic 16-bit DSP Processor"
	case EM_MMIX:
		return "Donald Knuth's educational 64-bit proc"
	case EM_HUANY:
		return "Harvard University machine-independent object files"
	case EM_PRISM:
		return "SiTera Prism"
	case EM_AVR:
		return "Atmel AVR 8-bit microcontroller"
	case EM_FR30:
		return "Fujitsu FR30"
	case EM_D10V:
		return "Mitsubishi D10V"
	case EM_D30V:
		return "Mitsubishi D30V"
	case EM_V850:
		return "NEC v850"
	case EM_M32R:
		return "Mitsubishi M32R"
	case EM_MN10300:
		return "Matsushita MN10300"
	case EM_MN10200:
		return "Matsushita MN10200"
	case EM_PJ:
		return "picoJava"
	case EM_OPENRISC:
		return "OpenRISC 32-bit embedded processor"
	case EM_ARC_COMPACT:
		return "ARC International ARCompact"
	case EM_XTENSA:
		return "Tensilica Xtensa Architecture"
	case EM_VIDEOCORE:
		return "Alphamosaic VideoCore"
	case EM_TMM_GPP:
		return "Thompson Multimedia General Purpose Proc"
	case EM_NS32K:
		return "National Semi. 32000"
	case EM_TPC:
		return "Tenor Network TPC"
	case EM_SNP1K:
		return "Trebia SNP 1000"
	case EM_ST200:
		return "STMicroelectronics ST200"
	case EM_IP2K:
		return "Ubicom IP2xxx"
	case EM_MAX:
		return "MAX processor"
	case EM_CR:
		return "National Semi. CompactRISC"
	case EM_F2MC16:
		return "Fujitsu F2MC16"
	case EM_MSP430:
		return "Texas Instruments msp430"
	case EM_BLACKFIN:
		return "Analog Devices Blackfin DSP"
	case EM_SE_C33:
		return "Seiko Epson S1C33 family"
	case EM_SEP:
		return "Sharp embedded microprocessor"
	case EM_ARCA:
		return "Arca RISC"
	case EM_UNICORE:
		return "PKU-Unity & MPRC Peking Uni. mc series"
	case EM_EXCESS:
		return "eXcess configurable cpu"
	case EM_DXP:
		return "Icera Semi. Deep Execution Processor"
	case EM_ALTERA_NIOS2:
		return "Altera Nios II"
	case EM_CRX:
		return "National Semi. CompactRISC CRX"
	case EM_XGATE:
		return "Motorola XGATE"
	case EM_C166:
		return "Infineon C16x/XC16x"
	case EM_M16C:
		return "Renesas M16C"
	case EM_DSPIC30F:
		return "Microchip Technology dsPIC30F"
	case EM_CE:
		return "Freescale Communication Engine RISC"
	case EM_M32C:
		return "Renesas M32C"
	case EM_TSK3000:
		return "Altium TSK3000"
	case EM_RS08:
		return "Freescale RS08"
	case EM_SHARC:
		return "Analog Devices SHARC family"
	case EM_ECOG2:
		return "Cyan Technology eCOG2"
	case EM_SCORE7:
		return "Sunplus S+core7 RISC"
	case EM_DSP24:
		return "New Japan Radio (NJR 24-bit DSP)"
	case EM_VIDEOCORE3:
		return "Broadcom VideoCore III"
	case EM_LATTICEMICO32:
		return "RISC for Lattice FPGA"
	case EM_SE_C17:
		return "Seiko Epson C17"
	case EM_TI_C6000:
		return "Texas Instruments TMS320C6000 DSP"
	case EM_TI_C2000:
		return "Texas Instruments TMS320C2000 DSP"
	case EM_TI_C5500:
		return "Texas Instruments TMS320C55x DSP"
	case EM_TI_ARP32:
		return "Texas Instruments App. Specific RISC"
	case EM_TI_PRU:
		return "Texas Instruments Prog. Realtime Unit"
	case EM_MMDSP_PLUS:
		return "STMicroelectronics 64bit VLIW DSP"
	case EM_CYPRESS_M8C:
		return "Cypress M8C"
	case EM_R32C:
		return "Renesas R32C"
	case EM_TRIMEDIA:
		return "NXP Semi. TriMedia"
	case EM_QDSP6:
		return "QUALCOMM DSP6"
	case EM_8051:
		return "Intel 8051 and variants"
	case EM_STXP7X:
		return "STMicroelectronics STxP7x"
	case EM_NDS32:
		return "Andes Tech. compact code emb. RISC"
	case EM_ECOG1X:
		return "Cyan Technology eCOG1X"
	case EM_MAXQ30:
		return "Dallas Semi. MAXQ30 mc"
	case EM_XIMO16:
		return "New Japan Radio (NJR 16-bit DSP)"
	case EM_MANIK:
		return "M2000 Reconfigurable RISC"
	case EM_CRAYNV2:
		return "Cray NV2 vector architecture"
	case EM_RX:
		return "Renesas RX"
	case EM_METAG:
		return "Imagination Tech. META"
	case EM_MCST_ELBRUS:
		return "MCST Elbrus"
	case EM_ECOG16:
		return "Cyan Technology eCOG16"
	case EM_CR16:
		return "National Semi. CompactRISC CR16"
	case EM_ETPU:
		return "Freescale Extended Time Processing Unit"
	case EM_SLE9X:
		return "Infineon Tech. SLE9X"
	case EM_L10M:
		return "Intel L10M"
	case EM_K10M:
		return "Intel K10M"
	case EM_AARCH64:
		return "ARM AARCH64"
	case EM_AVR32:
		return "Amtel 32-bit microprocessor"
	case EM_STM8:
		return "STMicroelectronics STM8"
	case EM_TILE64:
		return "Tilera TILE64"
	case EM_TILEPRO:
		return "Tilera TILEPro"
	case EM_MICROBLAZE:
		return "Xilinx MicroBlaze"
	case EM_CUDA:
		return "NVIDIA CUDA"
	case EM_TILEGX:
		return "Tilera TILE-Gx"
	case EM_CLOUDSHIELD:
		return "CloudShield"
	case EM_COREA_1ST:
		return "KIPO-KAIST Core-A 1st gen."
	case EM_COREA_2ND:
		return "KIPO-KAIST Core-A 2nd gen."
	case EM_ARCV2:
		return "Synopsys ARCv2 ISA"
	case EM_OPEN8:
		return "Open8 RISC"
	case EM_RL78:
		return "Renesas RL78"
	case EM_VIDEOCORE5:
		return "Broadcom VideoCore V"
	case EM_78KOR:
		return "Renesas 78KOR"
	case EM_56800EX:
		return "Freescale 56800EX DSC"
	case EM_BA1:
		return "Beyond BA1"
	case EM_BA2:
		return "Beyond BA2"
	case EM_XCORE:
		return "XMOS xCORE"
	case EM_MCHP_PIC:
		return "Microchip 8-bit PIC(r)"
	case EM_INTELGT:
		return "Intel Graphics Technology"
	case EM_KM32:
		return "KM211 KM32"
	case EM_KMX32:
		return "KM211 KMX32"
	case EM_EMX16:
		return "KM211 KMX16"
	case EM_EMX8:
		return "KM211 KMX8"
	case EM_KVARC:
		return "KM211 KVARC"
	case EM_CDP:
		return "Paneve CDP"
	case EM_COGE:
		return "Cognitive Smart Memory Processor"
	case EM_COOL:
		return "Bluechip CoolEngine"
	case EM_NORC:
		return "Nanoradio Optimized RISC"
	case EM_CSR_KALIMBA:
		return "CSR Kalimba"
	case EM_Z80:
		return "Zilog Z80"
	case EM_VISIUM:
		return "Controls and Data Services VISIUMcore"
	case EM_FT32:
		return "FTDI Chip FT32"
	case EM_MOXIE:
		return "Moxie processor"
	case EM_AMDGPU:
		return "AMD GPU"
	case EM_RISCV:
		return "RISC-V"
	case EM_BPF:
		return "Linux BPF"
	case EM_CSKY:
		return "C-SKY"
	case EM_LOONGARCH:
		return "LoongArch"
	default:
		return "<Unknown>"
	}
}

// Getter for getting interpreter from ELF File
func GetInterpreter(ehdr *Elf64_Ehdr, phdr []Elf64_Phdr, data []byte) string {
	for i := uint16(0); i < ehdr.E_phnum; i++ {
		if phdr[i].P_type != PT_INTERP {
			continue
		}
		start := phdr[i].P_offset
		end := start + phdr[i].P_filesz

		if end > uint64(len(data)) || start >= end {
			return "<Invalid interpreter offset>"
		}

		// Strip null terminator
		interp := data[start:end]
		if interp[len(interp)-1] == 0 {
			interp = interp[:len(interp)-1]
		}

		return unsafe.String(&interp[0], len(interp))
	}

	return "<Interpreter not available>"
}

// GetPType returns a human-readable string for the program header type (p_type field).
// It handles standard, OS-specific, and processor-specific segment types.
func GetPType(p_type uint32) string {
	// Fast path: common types (0-8)
	if p_type <= PT_NUM {
		switch p_type {
		case PT_NULL:
			return "NULL"
		case PT_LOAD:
			return "LOAD"
		case PT_DYNAMIC:
			return "DYNAMIC"
		case PT_INTERP:
			return "INTERP"
		case PT_NOTE:
			return "NOTE"
		case PT_SHLIB:
			return "SHLIB"
		case PT_PHDR:
			return "PHDR"
		case PT_TLS:
			return "TLS"
		}
	}

	// GNU extensions (most common in practice)
	switch p_type {
	case PT_GNU_EH_FRAME:
		return "GNU_EH_FRAME"
	case PT_GNU_STACK:
		return "GNU_STACK"
	case PT_GNU_RELRO:
		return "GNU_RELRO"
	case PT_GNU_PROPERTY:
		return "GNU_PROPERTY"
	case PT_GNU_SFRAME:
		return "GNU_SFRAME"
	}

	// Range-based detection (order matters due to overlap)
	if p_type >= PT_LOPROC && p_type <= PT_HIPROC {
		return fmt.Sprintf("PROC_SPECIFIC: 0x%x", p_type)
	}

	if p_type >= PT_LOOS && p_type <= PT_HIOS {
		// Check specific Sun values
		switch p_type {
		case PT_SUNWBSS:
			return "SUNWBSS"
		case PT_SUNWSTACK:
			return "SUNWSTACK"
		}
		return fmt.Sprintf("OS_SPECIFIC: 0x%x", p_type)
	}

	return fmt.Sprintf("<Unknown: 0x%x>", p_type)
}

// Getter for retrieving segment flags (p_flags)
func GetPFlags(p_flags uint32) string {
	var sb strings.Builder

	// R flag
	if p_flags&PF_R != 0 {
		sb.WriteString(ui.Blue.Sprint("R "))
	} else {
		sb.WriteString("  ")
	}
	//
	// W flag
	if p_flags&PF_W != 0 {
		sb.WriteString(ui.Yellow.Sprint("W "))
	} else {
		sb.WriteString("  ")
	}

	// E/X flag
	if p_flags&PF_X != 0 {
		sb.WriteString(ui.Red.Sprint("E "))
	} else {
		sb.WriteString("  ")
	}

	return sb.String()
}

// // Getter for sh_type (section type)
// func GetSectionTypeName(shType uint32) string {
// 	switch shType {
// 	case 0:
// 		return "NULL"
// 	case 1:
// 		return "PROGBITS"
// 	case 2:
// 		return "SYMTAB"
// 	case 3:
// 		return "STRTAB"
// 	case 4:
// 		return "RELA"
// 	case 5:
// 		return "HASH"
// 	case 6:
// 		return "DYNAMIC"
// 	case 7:
// 		return "NOTE"
// 	case 8:
// 		return "NOBITS"
// 	case 9:
// 		return "REL"
// 	case 10:
// 		return "SHLIB"
// 	case 11:
// 		return "DYNSYM"
// 	case 14:
// 		return "INIT_ARRAY"
// 	case 15:
// 		return "FINI_ARRAY"
// 	case 16:
// 		return "PREINIT_ARRAY"
// 	case 17:
// 		return "GROUP"
// 	case 18:
// 		return "SYMTAB_SHNDX"
// 	case 19:
// 		return "RELR"
// 	case 0x6ffffff5:
// 		return "GNU_ATTRIBUTES"
// 	case 0x6ffffff6:
// 		return "GNU_HASH"
// 	case 0x6ffffff7:
// 		return "GNU_LIBLIST"
// 	case 0x6ffffff8:
// 		return "CHECKSUM"
// 	case 0x6ffffffa:
// 		return "SUNW_move"
// 	case 0x6ffffffb:
// 		return "SUNW_COMDAT"
// 	case 0x6ffffffc:
// 		return "SUNW_syminfo"
// 	case 0x6ffffffd:
// 		return "GNU_verdef"
// 	case 0x6ffffffe:
// 		return "GNU_verneed"
// 	case 0x6fffffff:
// 		return "GNU_versym"
// 	default:
// 		if shType >= 0x70000000 && shType <= 0x7fffffff {
// 			return "PROC_SPECIFIC"
// 		}
// 		if shType >= 0x80000000 && shType <= 0x8fffffff {
// 			return "USER_SPECIFIC"
// 		}
// 		if shType >= 0x60000000 && shType < 0x70000000 {
// 			return "OS_SPECIFIC"
// 		}
// 		return "UNKNOWN"
// 	}
// }
//
// // Cache the string table data once
// func GetStringTable(data []byte, hdr *Elf64_Ehdr, shdr []*ELF64_Shdr) []byte {
// 	shstrtab := shdr[hdr.e_shstrndx]
// 	return data[shstrtab.sh_offset : shstrtab.sh_offset+shstrtab.sh_size]
// }
//
// // Then use it multiple times
// func GetSectionNameFromStringTable(strtab []byte, section *ELF64_Shdr) string {
// 	return extractString(strtab, section.sh_name)
// }
//
// // Extracting sh_name from string table
// func extractString(data []byte, offset uint32) string {
// 	if offset >= uint32(len(data)) {
// 		return ""
// 	}
//
// 	end := offset
// 	for end < uint32(len(data)) && data[end] != 0 {
// 		end++
// 	}
// 	return string(data[offset:end])
// }
//
// // Extracting sh_flags
// func GetShFlagsString(shFlags uint32) string {
// 	result := make([]rune, 0, 16)
//
// 	if shFlags&SHF_WRITE != 0 {
// 		result = append(result, 'W')
// 	}
//
// 	if shFlags&SHF_ALLOC != 0 {
// 		result = append(result, 'A')
// 	}
// 	if shFlags&SHF_EXECINSTR != 0 {
// 		result = append(result, 'X')
// 	}
// 	if shFlags&SHF_MERGE != 0 {
// 		result = append(result, 'M')
// 	}
// 	if shFlags&SHF_STRINGS != 0 {
// 		result = append(result, 'S')
// 	}
// 	if shFlags&SHF_INFO_LINK != 0 {
// 		result = append(result, 'I')
// 	}
// 	if shFlags&SHF_LINK_ORDER != 0 {
// 		result = append(result, 'L')
// 	}
// 	if shFlags&SHF_OS_NONCONFORMING != 0 {
// 		result = append(result, 'O')
// 	}
// 	if shFlags&SHF_GROUP != 0 {
// 		result = append(result, 'G')
// 	}
// 	if shFlags&SHF_TLS != 0 {
// 		result = append(result, 'T')
// 	}
// 	if shFlags&SHF_COMPRESSED != 0 {
// 		result = append(result, 'C')
// 	}
// 	if shFlags&SHF_GNU_RETAIN != 0 {
// 		result = append(result, 'R')
// 	}
// 	if shFlags&SHF_EXCLUDE != 0 {
// 		result = append(result, 'E')
// 	}
//
// 	// Check for OS-specific shFlags (excluding specific bits)
// 	if (shFlags&SHF_MASKOS) != 0 && (shFlags&SHF_GNU_RETAIN == 0) {
// 		result = append(result, 'o')
// 	}
//
// 	// Check for processor-specific shFlags
// 	if shFlags&SHF_MASKPROC != 0 {
// 		result = append(result, 'p')
// 	}
//
// 	return string(result)
// }
