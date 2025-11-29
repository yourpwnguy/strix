package parser

import (
	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/unsafe"
)

// parseSectionHeaders parses section headers with zero-copy.
func parseSectionHeaders(data []byte, ehdr *types.Elf64_Ehdr) ([]types.Elf64_Shdr, error) {
	return unsafe.CastSectionHeaders(data, ehdr.E_shnum, ehdr.E_shoff), nil
}
