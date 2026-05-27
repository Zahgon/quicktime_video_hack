package coremedia

// AudioFormatIDLpcm is the CoreMedia MediaID for LPCM
const AudioFormatIDLpcm uint32 = 0x6C70636D

// AudioStreamBasicDescription represents the struct found here: https://github.com/nu774/MSResampler/blob/master/CoreAudio/CoreAudioTypes.h
type AudioStreamBasicDescription struct {
	SampleRate       float64
	FormatID         uint32
	FormatFlags      uint32
	BytesPerPacket   uint32
	FramesPerPacket  uint32
	BytesPerFrame    uint32
	ChannelsPerFrame uint32
	BitsPerChannel   uint32
	Reserved         uint32
}

// DefaultAudioStreamBasicDescription creates a LPCM AudioStreamBasicDescription with default values I grabbed from the hex dump
func DefaultAudioStreamBasicDescription() AudioStreamBasicDescription {
	_ = "STUB: not implemented"
	return *new(AudioStreamBasicDescription)
}

func (adsb AudioStreamBasicDescription) String() string { _ = "STUB: not implemented"; return "" }

// NewAudioStreamBasicDescriptionFromBytes reads AudioStreamBasicDescription from bytes
func NewAudioStreamBasicDescriptionFromBytes(data []byte) (AudioStreamBasicDescription, error) {
	_ = "STUB: not implemented"
	return *new(AudioStreamBasicDescription), nil
}

// SerializeAudioStreamBasicDescription puts an AudioStreamBasicDescription into the given byte array
func (adsb AudioStreamBasicDescription) SerializeAudioStreamBasicDescription(adsbBytes []byte) {
	_ = "STUB: not implemented"
	return
}
