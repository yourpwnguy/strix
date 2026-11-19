package parser

// tableFits reports whether a header table with count entries of entrySize
// bytes each, starting at offset, fits inside data. A zero count is always
// fine, even when offset points past the end of the file.
func tableFits(data []byte, offset uint64, count, entrySize uint16) bool {
	if count == 0 {
		return true
	}

	size := uint64(count) * uint64(entrySize)
	return offset <= uint64(len(data)) && size <= uint64(len(data))-offset
}
