package common

// WriteLengthAndMagic just writes length and magic as uint32 4 byte values into the given array.
func WriteLengthAndMagic(bytes []byte, length int, magic uint32) { _ = "STUB: not implemented"; return }

// ParseLengthAndMagic checks if if the given byte array is longer or equal the uint32 in the first 4 bytes, and if the magic value in the second 4 bytes equals the supplied magic
// and returns the length, a slice of the bytes without length and magic or an error.
func ParseLengthAndMagic(bytes []byte, exptectedMagic uint32) (int, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}
