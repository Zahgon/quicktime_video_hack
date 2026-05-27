package packet

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// AsynSprpPacket seems to be a set property packet sent by the device.
type AsynSprpPacket struct {
	ClockRef CFTypeID
	Property coremedia.StringKeyEntry
}

// NewAsynSprpPacketFromBytes creates a new AsynSprpPacket from bytes
func NewAsynSprpPacketFromBytes(data []byte) (AsynSprpPacket, error) {
	_ = "STUB: not implemented"
	return *new(AsynSprpPacket), nil
}

func (sp AsynSprpPacket) String() string { _ = "STUB: not implemented"; return "" }
