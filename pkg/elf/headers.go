package elf

import (
	"fmt"
	"unsafe"

	"github.com/yourpwnguy/strix/pkg/utils"
)

// Read ElF Header into Elf64 Ehdr struct
func ReadELFHeader(filepath string) (*Elf64_Ehdr, func(), error) {
	data, cleanup, err := utils.MmapFile(filepath)
	if err != nil {
		return nil, func() {}, fmt.Errorf("%s Mmap failed: %w", utils.ErrPrefix, err)
	}

	header := (*Elf64_Ehdr)(unsafe.Pointer(&data[0]))

	if len(data) < int(unsafe.Sizeof(Elf64_Ehdr{})) {
		return nil, cleanup, fmt.Errorf("%s File too small to be ELF: %d bytes", utils.ErrPrefix, len(data))
	}

	if err := validateheader(header); err != nil {
		return nil, cleanup, err
	}

	return header, cleanup, nil
}

// Validate ELF MAGIC, CLASS, DATA, VERSION
func validateheader(header *Elf64_Ehdr) error {
	ident := &header.e_ident

	magic := *(*uint32)(unsafe.Pointer(&header.e_ident.fileIdn[0]))

	// Check if it's an ELF
	if magic != ELF_MAGIC {
		return fmt.Errorf("%s Not an ELF file - it has the wrong magic bytes at the start: 0x%x", utils.ErrPrefix, magic)
	}

	// Class check
	if ident.ei_class != ELFCLASS32 && ident.ei_class != ELFCLASS64 {
		return fmt.Errorf("%s Invalid ELF class: %d", utils.ErrPrefix, ident.ei_class)
	}

	// Data encoding check
	if ident.ei_data != ELFDATA2LSB && ident.ei_data != ELFDATA2MSB {
		return fmt.Errorf("%s Invalid data encoding: %d", utils.ErrPrefix, ident.ei_data)
	}

	// Version check (EV_CURRENT is always 1)
	if ident.ei_version != 1 {
		return fmt.Errorf("%s Invalid ELF version: %d", utils.ErrPrefix, ident.ei_version)
	}

	return nil
}
