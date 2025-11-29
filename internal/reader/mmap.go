package reader

import (
	"os"
	"syscall"
)

// Represents the Mmap Reader for reading file data
type MmapReader struct {
	f    *os.File
	data []byte
}

// MmapFile memory-maps the file at path for read-only access with sequential read optimization.
// Returns the mapped byte slice and any error.
func (m *MmapReader) Read(path string) ([]byte, error) {
	// Get the file descriptor
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// File metadata
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	// Size validation
	size := fi.Size()
	if size == 0 {
		f.Close()
		return nil, err
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
		return nil, err
	}

	// Lock pages in memory for critical binaries
	_ = syscall.Mlock(data)

	// Tell kernel we will read sequentially
	_ = syscall.Madvise(data, syscall.MADV_SEQUENTIAL|syscall.MADV_WILLNEED)

	// Storing for cleanup
	m.data = data
	m.f = f

	return data, nil
}

// Close unlocks the pages in memory, unmap the memory, and close the file handle
func (m *MmapReader) Close() {
	if m.data == nil {
		return
	}

	if m.f == nil {
		return
	}

	// Unlock memory
	_ = syscall.Munlock(m.data)

	// Unmap
	_ = syscall.Munmap(m.data)

	// Close file
	_ = m.f.Close()

	// Clear state
	m.data = nil
	m.f = nil
}
