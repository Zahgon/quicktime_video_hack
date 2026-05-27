package packet

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// AsynSratPacket is probably related to AVPlayer.SetRate somehow. I dont know exactly what everything means here
type AsynSratPacket struct {
	ClockRef CFTypeID
	Rate1    float32
	Rate2    float32
	Time     coremedia.CMTime
}

// NewAsynSratPacketFromBytes parses a new AsynSratPacket from bytes
func NewAsynSratPacketFromBytes(data []byte) (AsynSratPacket, error) {
	_ = "STUB: not implemented"
	return *new(AsynSratPacket), nil
}

func (sp AsynSratPacket) String() string { _ = "STUB: not implemented"; return "" }
