package elf

const (
	// Magic number validation for ELF
	ELF_MAGIC Elf64_Word = 0x464c457f

	// ELF Class (ei_class)
	ELFCLASSNOEN uint8 = 0
	ELFCLASS32   uint8 = 1
	ELFCLASS64   uint8 = 2

	// Data encoding (ei_data)
	ELFDATANONE uint8 = 0
	ELFDATA2LSB uint8 = 1
	ELFDATA2MSB uint8 = 2

	// OS/ABI (ei_osabi)
	ELFOSABI_NONE       uint8 = 0
	ELFOSABI_LINUX      uint8 = 3
	ELFOSABI_FREEBSD    uint8 = 9
	ELFOSABI_OPENBSD    uint8 = 12
	ELFOSABI_STANDALONE uint8 = 255
)
