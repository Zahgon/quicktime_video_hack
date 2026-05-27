package packet

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// AsynCmSampleBufPacket contains a CMSampleBuffer with audio or video data
type AsynCmSampleBufPacket struct {
	ClockRef    CFTypeID
	CMSampleBuf coremedia.CMSampleBuffer
}

// NewAsynCmSampleBufPacketFromBytes parses a new AsynCmSampleBufPacket from bytes
func NewAsynCmSampleBufPacketFromBytes(data []byte) (AsynCmSampleBufPacket, error) {
	_ = "STUB: not implemented"
	return *new(AsynCmSampleBufPacket), nil
}

func newAsynCmSampleBufferPacketFromBytes(data []byte) (CFTypeID, coremedia.CMSampleBuffer, error) {
	_ = "STUB: not implemented"
	return *new(CFTypeID), *new(coremedia.CMSampleBuffer), nil
}

func (sp AsynCmSampleBufPacket) String() string { _ = "STUB: not implemented"; return "" }
