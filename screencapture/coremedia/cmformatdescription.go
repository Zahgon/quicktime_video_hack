package coremedia

// Those are the markers found in the hex dumps.
// For convenience I have added the ASCII representation as a comment
// in normal byte order and reverse byteorder (so you can find them in the hex dumps)
// Note: I have just guessed what the names could be from the marker ascii, I could be wrong ;-)
const (
	FormatDescriptorMagic            uint32 = 0x66647363 //fdsc - csdf
	MediaTypeVideo                   uint32 = 0x76696465 //vide - ediv
	MediaTypeSound                   uint32 = 0x736F756E //nuos - soun
	MediaTypeMagic                   uint32 = 0x6D646961 //mdia - aidm
	VideoDimensionMagic              uint32 = 0x7664696D //vdim - midv
	CodecMagic                       uint32 = 0x636F6463 //codc - cdoc
	CodecAvc1                        uint32 = 0x61766331 //avc1 - 1cva
	ExtensionMagic                   uint32 = 0x6578746E //extn - ntxe
	AudioStreamBasicDescriptionMagic uint32 = 0x61736264 //asdb - dbsa
)

// FormatDescriptor is actually a CMFormatDescription
// https://developer.apple.com/documentation/coremedia/cmformatdescription
// https://github.com/phracker/MacOSX-SDKs/blob/master/MacOSX10.9.sdk/System/Library/Frameworks/CoreMedia.framework/Versions/A/Headers/CMFormatDescription.h
type FormatDescriptor struct {
	MediaType            uint32
	VideoDimensionWidth  uint32
	VideoDimensionHeight uint32
	Codec                uint32
	Extensions           IndexKeyDict
	//PPS contains bytes of the Picture Parameter Set h264 NALu
	PPS []byte
	//SPS contains bytes of the Picture Parameter Set h264 NALu
	SPS                         []byte
	AudioStreamBasicDescription AudioStreamBasicDescription
}

// NewFormatDescriptorFromBytes parses a CMFormatDescription from bytes
func NewFormatDescriptorFromBytes(data []byte) (FormatDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(FormatDescriptor), nil
}

func parseSoundFdsc(remainingBytes []byte) (FormatDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(FormatDescriptor), nil
}

func parseVideoFdsc(remainingBytes []byte) (FormatDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(FormatDescriptor), nil
}

//doc on extensions at the bottom of: https://developer.apple.com/documentation/coremedia/cmformatdescription?language=objc

func extractPPS(dict IndexKeyDict) ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func parseCodec(bytes []byte) (uint32, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func parseVideoDimension(bytes []byte) (uint32, uint32, []byte, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

func parseMediaType(bytes []byte) (uint32, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (fdsc FormatDescriptor) String() string { _ = "STUB: not implemented"; return "" }

func readableCodec(codec uint32) string { _ = "STUB: not implemented"; return "" }

func readableMediaType(mediaType uint32) string { _ = "STUB: not implemented"; return "" }
