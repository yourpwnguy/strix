package reader

// BinaryReader abstracts binary data sources for testing and flexibility.
type BinaryReader interface {
	Read(path string) (data []byte, err error)
	Close()
}
