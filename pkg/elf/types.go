package elf

// Types of addresses
type (
	Elf64_Addr = uint64
)

// Types of offsets
type (
	Elf64_Off = uint64
)

/* Type for a 16-bit quantity.  */
type (
	Elf64_Half = uint16
)

/* Types for signed and unsigned 32-bit quantities.  */
type (
	Elf64_Word  = uint32
	Elf64_Sword = int32
)

/* Types for signed and unsigned 64-bit quantities.  */
type (
	Elf64_Xword  = uint64
	Elf64_Sxword = int64
)

/* Type for section indices, which are 16-bit quantities.  */
type (
	Elf64_Section = uint16
)

/* Type for version symbol information.  */
type (
	Elf64_Versym = uint32
)

// Representing ELF Header for 64 Bit Executables
type Elf64_Ehdr struct {
	e_ident     e_ident_t  // Contains magic number and basic information about the ELF
	e_type      Elf64_Half // Object file type
	e_machine   Elf64_Half // Architecture
	e_version   Elf64_Word // Object file version
	e_entry     Elf64_Addr // Entry Point Virtual Address
	e_phoff     Elf64_Off  // Program header table file offset
	e_shoff     Elf64_Off  // Section header table file offset
	e_flags     Elf64_Word // Processor specific flags
	e_ehsize    Elf64_Half // ELF header size in bytes
	e_phentsize Elf64_Half // Program header table entry size
	e_phnum     Elf64_Half // Program header table entry count
	e_shentsize Elf64_Half // Section header table entry size
	e_shnum     Elf64_Half // Section header table entry count
	e_shstrndx  Elf64_Half // Section header string table index
}

type e_ident_t struct {
	fileIdn       [4]uint8 // Magic bytes
	ei_class      uint8    // ELF Class (32/64 bits)
	ei_data       uint8    // ELF Data encoding (LSB/MSB)
	ei_version    uint8    // ELF Version
	ei_osabi      uint8    // OS ABI identification
	ei_abiversion uint8    // OS ABI version
	ei_pad        [7]uint8 // Padding
}
