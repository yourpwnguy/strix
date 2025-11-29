package parser

import (
	"fmt"

	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/ui"
	"github.com/yourpwnguy/strix/internal/unsafe"
)

// parseHeader validates and parses the ELF64 header from the provided byte slice.
// Returns an error if the magic number is invalid or the data is too small.
func parseELFHeader(data []byte) (*types.Elf64_Ehdr, error) {
	if !unsafe.HasMinimumSize(data) {
		return nil, fmt.Errorf("%s File too small to be ELF: %d bytes",
			ui.ErrPrefix,
			len(data),
		)
	}

	if !unsafe.HasValidMagic(data[:4]) {
		return nil, fmt.Errorf("%s Invalid ELF magic value: %#04x",
			ui.ErrPrefix,
			data[:4],
		)
	}

	return unsafe.CastHeader(data), nil
}
