package packet

// SyncStopPacket requests us to stop our clock
type SyncStopPacket struct {
	ClockRef      CFTypeID
	CorrelationID uint64
}

// NewSyncStopPacketFromBytes parses a SyncStopPacket from bytes
func NewSyncStopPacketFromBytes(data []byte) (SyncStopPacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncStopPacket), nil
}

// NewReply creates a byte array containing the given skew
func (sp SyncStopPacket) NewReply() []byte { _ = "STUB: not implemented"; return nil }

func (sp SyncStopPacket) String() string { _ = "STUB: not implemented"; return "" }
