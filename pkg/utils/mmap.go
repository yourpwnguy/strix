package utils

import (
	"os"
	"syscall"
)

func MmapFile(path string) ([]byte, func(), error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, nil, err
	}

	data, err := syscall.Mmap(
		int(f.Fd()),
		0,
		int(fi.Size()),
		syscall.PROT_READ,
		syscall.MAP_SHARED,
	)

	cleanup := func() {
		syscall.Munmap(data)
		f.Close()
	}
	return data, cleanup, nil
}
