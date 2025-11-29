package types

// Types of addresses
type (
	Elf64_Addr = uint64
)

// Types of offsets
type (
	Elf64_Off = uint64
)

// Type for a 16-bit quantity.
type (
	Elf64_Half = uint16
)

// Types for signed and unsigned 32-bit quantities.
type (
	Elf64_Word  = uint32
	Elf64_Sword = int32
)

// Types for signed and unsigned 64-bit quantities.
type (
	Elf64_Xword  = uint64
	Elf64_Sxword = int64
)

// Type for section indices, which are 16-bit quantities.
type (
	Elf64_Section = uint16
)

// Type for version symbol information.
type (
	Elf64_Versym = uint32
)

// Representing ELF identification
type e_ident_t struct {
	FileIdn       [4]uint8 // Magic bytes
	Ei_class      uint8    // ELF Class (32/64 bits)
	Ei_data       uint8    // ELF Data encoding (LSB/MSB)
	Ei_version    uint8    // ELF Version
	Ei_osabi      uint8    // OS ABI identification
	Ei_abiversion uint8    // OS ABI version
	Ei_pad        [7]uint8 // Padding
}

// Representing ELF Header for 64 Bit Executables
type Elf64_Ehdr struct {
	E_ident     e_ident_t  // Contains magic number and basic information about the ELF
	E_type      Elf64_Half // Object file type
	E_machine   Elf64_Half // Architecture
	E_version   Elf64_Word // Object file version
	E_entry     Elf64_Addr // Entry Point Virtual Address
	E_phoff     Elf64_Off  // Program header table file offset
	E_shoff     Elf64_Off  // Section header table file offset
	E_flags     Elf64_Word // Processor specific flags
	E_ehsize    Elf64_Half // ELF header size in bytes
	E_phentsize Elf64_Half // Program header table entry size
	E_phnum     Elf64_Half // Program header table entry count
	E_shentsize Elf64_Half // Section header table entry size
	E_shnum     Elf64_Half // Section header table entry count
	E_shstrndx  Elf64_Half // Section header string table index
}

// Representing ELF program header
type Elf64_Phdr struct {
	P_type   Elf64_Word
	P_flags  Elf64_Word
	P_offset Elf64_Off
	P_vaddr  Elf64_Addr
	P_paddr  Elf64_Addr
	P_filesz Elf64_Xword
	P_memsz  Elf64_Xword
	P_align  Elf64_Xword
}

// Representing ELF Section Header
type Elf64_Shdr struct {
	Sh_name      Elf64_Word  // Section name (string tbl index)
	Sh_type      Elf64_Word  // Section type
	Sh_flags     Elf64_Word  // Section flags
	Sh_addr      Elf64_Addr  // Section virtual addr at execution
	Sh_offset    Elf64_Off   // Section file offset
	Sh_size      Elf64_Xword // Section size in bytes
	Sh_link      Elf64_Word  // Link to another section
	Sh_info      Elf64_Word  // Addtional section information
	Sh_addralign Elf64_Xword // Section alignment
	Sh_entsize   Elf64_Xword // Entry size if section holds table
}
