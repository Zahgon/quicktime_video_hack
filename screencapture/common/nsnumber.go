package common

// NumberValueMagic is vbmn in little endian ascii ==> nmbv
const NumberValueMagic uint32 = 0x6E6D6276

// NSNumber represents a type in the binary protocol used. Type 6 seems to be a float64, type 4 a int64, type 3 a int32.
// I am not sure whether signed or unsigned. They are all in LittleEndian
type NSNumber struct {
	typeSpecifier byte
	//not certain if these are really unsigned
	IntValue   uint32
	LongValue  uint64
	FloatValue float64
}

// NewNSNumberFromUInt32 create NSNumber of type 0x03 with a 4 byte int as value
func NewNSNumberFromUInt32(intValue uint32) NSNumber {
	_ = "STUB: not implemented"
	return *new(NSNumber)
}

// NewNSNumberFromUInt64 create NSNumber of type 0x04 with a 8 byte int as value
func NewNSNumberFromUInt64(longValue uint64) NSNumber {
	_ = "STUB: not implemented"
	return *new(NSNumber)
}

// NewNSNumberFromUFloat64 create NSNumber of type 0x06 with a 8 byte int as value
func NewNSNumberFromUFloat64(floatValue float64) NSNumber {
	_ = "STUB: not implemented"
	return *new(NSNumber)
}

// NewNSNumber reads a NSNumber from bytes.
func NewNSNumber(bytes []byte) (NSNumber, error) {
	_ = "STUB: not implemented"
	return *new(NSNumber), nil
}

// ToBytes serializes a NSNumber into a []byte.
// FIXME: remove allocation of array and use one that is passed in instead
func (n NSNumber) ToBytes() []byte { _ = "STUB: not implemented"; return nil }

//shouldn't happen

func (n NSNumber) String() string { _ = "STUB: not implemented"; return "" }
