package parser

import (
	"fmt"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/unsafe"
)

// ParseELFHeader validates and parses the ELF64 header from the provided byte slice.
// Returns an error if the magic number is invalid or the data is too small.
func ParseELFHeader(data []byte) (*types.Elf64_Ehdr, error) {
	if !unsafe.HasMinimumSize(data) {
		return nil, fmt.Errorf("file too small to be ELF: %d bytes", len(data))
	}

	if !unsafe.HasValidMagic(data[:4]) {
		return nil, fmt.Errorf("invalid ELF magic value: %#04x", data[:4])
	}

	return unsafe.CastHeader(data), nil
}
