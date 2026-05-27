package packet

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// SyncCvrpPacket contains all info from a CVRP packet sent by the device
type SyncCvrpPacket struct {
	ClockRef       CFTypeID
	CorrelationID  uint64
	DeviceClockRef CFTypeID
	Payload        coremedia.StringKeyDict
}

// NewSyncCvrpPacketFromBytes parses a SyncCvrpPacket from a []byte
func NewSyncCvrpPacketFromBytes(data []byte) (SyncCvrpPacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncCvrpPacket), nil
}

// NewReply creates a RPLY packet containing the given clockRef and serializes it to a []byte
func (sp SyncCvrpPacket) NewReply(clockRef CFTypeID) []byte { _ = "STUB: not implemented"; return nil }

func (sp SyncCvrpPacket) String() string { _ = "STUB: not implemented"; return "" }
