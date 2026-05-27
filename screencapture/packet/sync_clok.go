package packet

// SyncClokPacket contains a decoded Clok packet from the device
type SyncClokPacket struct {
	ClockRef      CFTypeID
	CorrelationID uint64
}

// NewSyncClokPacketFromBytes parses a SynClokPacket from bytes
func NewSyncClokPacketFromBytes(data []byte) (SyncClokPacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncClokPacket), nil
}

// NewReply creates a RPLY message containing the given clockRef and serializes it into a []byte
func (sp SyncClokPacket) NewReply(clockRef CFTypeID) []byte { _ = "STUB: not implemented"; return nil }

func (sp SyncClokPacket) String() string { _ = "STUB: not implemented"; return "" }
