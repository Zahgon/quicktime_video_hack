package packet

// SyncOgPacket represents the OG Message. I do not know what these messages mean.
type SyncOgPacket struct {
	ClockRef      CFTypeID
	CorrelationID uint64
	Unknown       uint32
}

// NewSyncOgPacketFromBytes parses a SyncOgPacket form bytes assuming it starts with SYNC magic and has the correct length.
func NewSyncOgPacketFromBytes(data []byte) (SyncOgPacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncOgPacket), nil
}

// NewReply returns a []byte containing the default reply for a SyncOgPacket
func (sp SyncOgPacket) NewReply() []byte { _ = "STUB: not implemented"; return nil }

func (sp SyncOgPacket) String() string { _ = "STUB: not implemented"; return "" }
