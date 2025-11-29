package unsafe

import (
	"unsafe"

	"github.com/yourpwnguy/strix/internal/elf/types"
)

// CastHeader performs a zero-copy cast of raw bytes to an Elf64_Ehdr structure.
func CastHeader(data []byte) *types.Elf64_Ehdr {
	return (*types.Elf64_Ehdr)(unsafe.Pointer(&data[0]))
}

// CastProgramHeaders casts raw bytes to []Elf64_Phdr with zero copying.
func CastProgramHeaders(data []byte, count uint16, offset uint64) []types.Elf64_Phdr {
	// Start of ELF program headers
	ptr := unsafe.Pointer(&data[offset])
	return unsafe.Slice((*types.Elf64_Phdr)(ptr), count)
}

// castSectionHeaders casts raw bytes to []Elf64_Shdr with zero copying.
func CastSectionHeaders(data []byte, count uint16, offset uint64) []types.Elf64_Shdr {
	// Start of ELF section headers
	ptr := unsafe.Pointer(&data[offset])
	return unsafe.Slice((*types.Elf64_Shdr)(ptr), count)
}

// HasValidMagic checks if the data starts with the ELF magic number (0x7f 'E' 'L' 'F').
func HasValidMagic(data []byte) bool {
	return len(data) >= 4 &&
		data[0] == 0x7f &&
		data[1] == 'E' &&
		data[2] == 'L' &&
		data[3] == 'F'
}

// HasMinimumSize checks if the data buffer is large enough to contain an ELF64 header.
func HasMinimumSize(data []byte) bool {
	return len(data) >= int(unsafe.Sizeof(types.Elf64_Ehdr{}))
}
