package elf

import (
	"fmt"

	"github.com/yourpwnguy/strix/internal/elf/parser"
	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/reader"
)

// Parser handles ELF64 file parsing with cached results for repeated access.
type ELFParser struct {
	reader reader.BinaryReader

	// Raw data (mmap'd or in-memory)
	data []byte

	// ELF specific information
	ehdr *types.Elf64_Ehdr
	phdr []types.Elf64_Phdr
	shdr []types.Elf64_Shdr
}

// New creates a parser from a BinaryReader.
func NewELFParser(r reader.BinaryReader) (*ELFParser, error) {
	if r == nil {
		return nil, fmt.Errorf("reader is nil")
	}

	data := r.Bytes()

	return &ELFParser{
		reader: r,
		data:   data,
	}, nil
}

// Open opens a file and creates a parser.
func Open(path string) (*ELFParser, error) {
	r, err := reader.Open(path)
	if err != nil {
		return nil, err
	}

	p, err := NewELFParser(r)
	if err != nil {
		r.Close()
		return nil, err
	}

	return p, nil
}

// Close releases all resources held by the parser, including the underlying reader.
func (p *ELFParser) Close() error {
	if p.reader != nil {
		return p.reader.Close()
	}
	return nil
}

// Header returns the ELF64 header. Results are cached after the first call.
func (p *ELFParser) ELFHeader() (*types.Elf64_Ehdr, error) {
	if p.ehdr != nil {
		return p.ehdr, nil
	}

	hdr, err := parser.ParseELFHeader(p.data)
	if err != nil {
		return nil, err
	}

	p.ehdr = hdr
	return hdr, err
}

// ProgramHeaders returns all program headers. Results are cached after the first call.
func (p *ELFParser) ProgramHeaders() ([]types.Elf64_Phdr, error) {
	if p.phdr != nil {
		return p.phdr, nil
	}

	ehdr, err := p.ELFHeader()
	if err != nil {
		return nil, err
	}

	phdr, err := parser.ParseProgramHeaders(p.data, ehdr)
	if err != nil {
		return nil, err
	}

	p.phdr = phdr
	return phdr, err
}

// SectionHeaders returns all section headers. Results are cached after the first call.
func (p *ELFParser) SectionHeaders() ([]types.Elf64_Shdr, error) {
	if p.shdr != nil {
		return p.shdr, nil
	}

	ehdr, err := p.ELFHeader()
	if err != nil {
		return nil, err
	}

	shdr, err := parser.ParseSectionHeaders(p.data, ehdr)
	if err != nil {
		return nil, err
	}

	p.shdr = shdr
	return shdr, nil
}

// Data returns the raw file bytes loaded by the ELFParser.
func (p *ELFParser) Data() []byte {
	return p.data
}
