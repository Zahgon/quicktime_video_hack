package packet

// SyncCwpaPacket contains all info from a CWPA packet sent by the device
type SyncCwpaPacket struct {
	ClockRef       CFTypeID
	CorrelationID  uint64
	DeviceClockRef CFTypeID
}

// NewSyncCwpaPacketFromBytes parses a SyncCwpaPacket from a []byte
func NewSyncCwpaPacketFromBytes(data []byte) (SyncCwpaPacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncCwpaPacket), nil
}

// NewReply creates a RPLY packet containing the given clockRef and serializes it to a []byte
func (sp SyncCwpaPacket) NewReply(clockRef CFTypeID) []byte { _ = "STUB: not implemented"; return nil }

func (sp SyncCwpaPacket) String() string { _ = "STUB: not implemented"; return "" }
