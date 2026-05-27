package packet

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// SyncTimePacket contains the data from a decoded Time Packet sent by the device
type SyncTimePacket struct {
	ClockRef      CFTypeID
	CorrelationID uint64
}

// NewSyncTimePacketFromBytes parses a SyncTimePacket from bytes
func NewSyncTimePacketFromBytes(data []byte) (SyncTimePacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncTimePacket), nil
}

// NewReply creates a RPLY packet containing the given CMTime and serializes it to a []byte
func (sp SyncTimePacket) NewReply(time coremedia.CMTime) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sp SyncTimePacket) String() string { _ = "STUB: not implemented"; return "" }
