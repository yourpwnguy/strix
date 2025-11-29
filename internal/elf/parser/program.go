package parser

import (
	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/unsafe"
)

// parseProgramHeaders validates and parses the ELF64 Program header from the provided byte slice.
func parseProgramHeaders(data []byte, ehdr *types.Elf64_Ehdr) ([]types.Elf64_Phdr, error) {
	return unsafe.CastProgramHeaders(data, ehdr.E_phnum, ehdr.E_phoff), nil
}
