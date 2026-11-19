package parser

import (
	"fmt"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/unsafe"
)

// elf64PhdrSize is sizeof(Elf64_Phdr) from the ELF64 ABI.
const elf64PhdrSize = 56

// ParseProgramHeaders validates and parses the ELF64 program headers.
func ParseProgramHeaders(data []byte, ehdr *types.Elf64_Ehdr) ([]types.Elf64_Phdr, error) {
	if ehdr.E_phnum == 0 {
		return nil, nil
	}

	if ehdr.E_phentsize != elf64PhdrSize {
		return nil, fmt.Errorf("unexpected program header entry size: %d", ehdr.E_phentsize)
	}

	if !tableFits(data, ehdr.E_phoff, ehdr.E_phnum, ehdr.E_phentsize) {
		return nil, fmt.Errorf("program header table out of bounds: offset %#x, %d entries",
			ehdr.E_phoff,
			ehdr.E_phnum,
		)
	}

	return unsafe.CastProgramHeaders(data, ehdr.E_phnum, ehdr.E_phoff), nil
}
