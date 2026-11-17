package reader

// BinaryReader provides access to raw binary data.
type BinaryReader interface {
	// Bytes returns the complete raw data
	Bytes() []byte

	// Size returns total byte count
	Size() int64

	// Close releases resources
	Close() error
}
