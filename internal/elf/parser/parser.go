package parser

import (
	"github.com/yourpwnguy/strix/internal/elf/types"
	"github.com/yourpwnguy/strix/internal/reader"
)

// Parser handles ELF64 file parsing with cached results for repeated access.
type Parser struct {
	reader reader.BinaryReader

	// Raw data (mmap'd or in-memory)
	data []byte

	// ELF specific information
	ehdr *types.Elf64_Ehdr
	phdr []types.Elf64_Phdr
	shdr []types.Elf64_Shdr
}

// NewParser creates a new Parser instance with the specified binary reader.
func NewParser(reader reader.BinaryReader) *Parser {
	return &Parser{
		reader: reader,
	}
}

// Load reads the ELF binary file into memory using the configured reader.
func (p *Parser) Load(path string) error {
	data, err := p.reader.Read(path)
	if err != nil {
		return err
	}
	p.data = data
	return nil
}

// Close releases all resources held by the parser, including the underlying reader.
func (p *Parser) Close() {
	if p.reader != nil {
		p.reader.Close()
	}
}

// Header returns the ELF64 header. Results are cached after the first call.
func (p *Parser) ELFHeader() (*types.Elf64_Ehdr, error) {
	if p.ehdr != nil {
		return p.ehdr, nil
	}

	hdr, err := parseELFHeader(p.data)
	if err != nil {
		return nil, err
	}

	p.ehdr = hdr
	return hdr, err
}

// ProgramHeaders returns all program headers. Results are cached after the first call.
func (p *Parser) ProgramHeaders() ([]types.Elf64_Phdr, error) {
	if p.phdr != nil {
		return p.phdr, nil
	}

	ehdr, err := p.ELFHeader()
	if err != nil {
		return nil, err
	}

	phdr, err := parseProgramHeaders(p.data, ehdr)
	if err != nil {
		return nil, err
	}

	p.phdr = phdr
	return phdr, err
}

// SectionHeaders returns all section headers. Results are cached after the first call.
func (p *Parser) SectionHeaders() ([]types.Elf64_Shdr, error) {
	if p.shdr != nil {
		return p.shdr, nil
	}

	ehdr, err := p.ELFHeader()
	if err != nil {
		return nil, err
	}

	shdr, err := parseSectionHeaders(p.data, ehdr)
	if err != nil {
		return nil, err
	}

	p.shdr = shdr
	return shdr, nil
}

// Data returns the raw file bytes loaded by the parser.
func (p *Parser) Data() []byte {
	return p.data
}
