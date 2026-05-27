package coremedia

import (
	"strings"
)

// Dictionary related magic marker constants
const (
	KeyValuePairMagic uint32 = 0x6B657976 //keyv - vyek
	StringKey         uint32 = 0x7374726B //strk - krts
	IntKey            uint32 = 0x6964786B //idxk - kxdi
	BooleanValueMagic uint32 = 0x62756C76 //bulv - vlub
	DictionaryMagic   uint32 = 0x64696374 //dict - tcid
	DataValueMagic    uint32 = 0x64617476 //datv - vtad
	StringValueMagic  uint32 = 0x73747276 //strv - vrts
)

// StringKeyDict a dictionary that uses strings as keys with an array of StringKeyEntry
type StringKeyDict struct {
	Entries []StringKeyEntry
}

// StringKeyEntry a pair of a string key and an arbitrary value
type StringKeyEntry struct {
	Key   string
	Value interface{}
}

// IndexKeyDict a dictionary that uses uint16 as keys with an array of IndexKeyEntry
type IndexKeyDict struct {
	Entries []IndexKeyEntry
}

// IndexKeyEntry is a pair of a uint16 key and an arbitrary value.
type IndexKeyEntry struct {
	Key   uint16
	Value interface{}
}

// NewIndexDictFromBytes creates a new dictionary assuming the byte array starts with the 4 byte length of the dictionary followed by "dict" as the magic marker
func NewIndexDictFromBytes(data []byte) (IndexKeyDict, error) {
	_ = "STUB: not implemented"
	return *new(IndexKeyDict), nil
}

// NewIndexDictFromBytesWithCustomMarker creates a new dictionary assuming the byte array starts with the 4 byte length of the dictionary followed by magic as the magic marker
func NewIndexDictFromBytesWithCustomMarker(data []byte, magic uint32) (IndexKeyDict, error) {
	_ = "STUB: not implemented"
	return *new(IndexKeyDict), nil
}

// NewStringDictFromBytes creates a new dictionary assuming the byte array starts with the 4 byte length of the dictionary followed by "dict" as the magic marker
func NewStringDictFromBytes(data []byte) (StringKeyDict, error) {
	_ = "STUB: not implemented"
	return *new(StringKeyDict), nil
}

func parseIntDictEntry(bytes []byte) (IndexKeyEntry, error) {
	_ = "STUB: not implemented"
	return *new(IndexKeyEntry), nil
}

// ParseKeyValueEntry parses a byte array into a StringKeyEntry assuming the array starts with a 4 byte length followed by the "keyv" magic
func ParseKeyValueEntry(data []byte) (StringKeyEntry, error) {
	_ = "STUB: not implemented"
	return *new(StringKeyEntry), nil
}

func parseEntry(bytes []byte) (StringKeyEntry, error) {
	_ = "STUB: not implemented"
	return *new(StringKeyEntry), nil
}

func parseKey(bytes []byte) (string, []byte, error) { _ = "STUB: not implemented"; return "", nil, nil }

func parseIntKey(bytes []byte) (uint16, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func parseValue(bytes []byte) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

//FIXME: that is a lazy implementation, improve please

func (dt StringKeyDict) String() string { _ = "STUB: not implemented"; return "" }

func (dt IndexKeyDict) String() string { _ = "STUB: not implemented"; return "" }

func appendIndexEntry(builder *strings.Builder, entry IndexKeyEntry) {
	_ = "STUB: not implemented"
	return
}

func appendEntry(builder *strings.Builder, entry StringKeyEntry) { _ = "STUB: not implemented"; return }

func valueToString(builder *strings.Builder, value interface{}) { _ = "STUB: not implemented"; return }

func (dt IndexKeyDict) getValue(index uint16) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
