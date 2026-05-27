package packet

// SyncSkewPacket requests us to reply with the current skew value
type SyncSkewPacket struct {
	ClockRef      CFTypeID
	CorrelationID uint64
}

// NewSyncSkewPacketFromBytes parses a SyncSkewPacket from bytes
func NewSyncSkewPacketFromBytes(data []byte) (SyncSkewPacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncSkewPacket), nil
}

// NewReply creates a byte array containing the given skew
func (sp SyncSkewPacket) NewReply(skew float64) []byte { _ = "STUB: not implemented"; return nil }

func (sp SyncSkewPacket) String() string { _ = "STUB: not implemented"; return "" }
