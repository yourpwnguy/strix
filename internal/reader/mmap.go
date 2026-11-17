package reader

import (
	"fmt"
	"os"
	"syscall"
)

// MmapReader provides memory-mapped read-only access to files.
type MmapReader struct {
	f    *os.File
	data []byte
	size int64
}

// Open memory-maps a file and returns a ready-to-use reader.
func Open(path string) (*MmapReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	// File metadata
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("stat: %w", err)
	}

	// Size validation
	size := fi.Size()
	if size == 0 {
		f.Close()
		return nil, fmt.Errorf("file is empty")
	}

	// Memory map the file
	data, err := syscall.Mmap(
		int(f.Fd()),
		0,
		int(fi.Size()),
		syscall.PROT_READ,
		syscall.MAP_SHARED|syscall.MAP_POPULATE, // ELF files are usually small, load it all
	)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("mmap: %w", err)
	}

	// Lock pages in memory for critical binaries
	_ = syscall.Mlock(data)

	// Tell kernel we will read sequentially
	_ = syscall.Madvise(data, syscall.MADV_SEQUENTIAL|syscall.MADV_WILLNEED)

	// Storing for cleanup
	return &MmapReader{
		f:    f,
		data: data,
		size: size,
	}, nil
}

// Returns the raw byte data
func (m *MmapReader) Bytes() []byte {
	return m.data
}

// Returns the total byte count
func (m *MmapReader) Size() int64 {
	return m.size
}

// Close unlocks the pages in memory, unmap the memory, and close the file handle
func (m *MmapReader) Close() error {
	if m.data == nil {
		return nil
	}

	_ = syscall.Munlock(m.data)
	_ = syscall.Munmap(m.data)
	m.data = nil

	if m.f != nil {
		m.f.Close()
		m.f = nil
	}

	return nil
}

// Compile-Time Interface Check
var _ BinaryReader = (*MmapReader)(nil)
