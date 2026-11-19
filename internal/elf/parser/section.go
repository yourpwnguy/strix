package parser

import (
	"fmt"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/unsafe"
)

// elf64ShdrSize is sizeof(Elf64_Shdr) from the ELF64 ABI.
const elf64ShdrSize = 64

// ParseSectionHeaders validates and parses the ELF64 section headers.
func ParseSectionHeaders(data []byte, ehdr *types.Elf64_Ehdr) ([]types.Elf64_Shdr, error) {
	if ehdr.E_shnum == 0 {
		return nil, nil
	}

	if ehdr.E_shentsize != elf64ShdrSize {
		return nil, fmt.Errorf("unexpected section header entry size: %d", ehdr.E_shentsize)
	}

	if !tableFits(data, ehdr.E_shoff, ehdr.E_shnum, ehdr.E_shentsize) {
		return nil, fmt.Errorf("section header table out of bounds: offset %#x, %d entries",
			ehdr.E_shoff,
			ehdr.E_shnum,
		)
	}

	return unsafe.CastSectionHeaders(data, ehdr.E_shnum, ehdr.E_shoff), nil
}
