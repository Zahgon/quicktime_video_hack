package packet

// AsynRelsPacket tells us that a clock was released
type AsynRelsPacket struct {
	ClockRef CFTypeID
}

// NewAsynRelsPacketFromBytes creates a new AsynRelsPacket from bytes
func NewAsynRelsPacketFromBytes(data []byte) (AsynRelsPacket, error) {
	_ = "STUB: not implemented"
	return *new(AsynRelsPacket), nil
}

func (sp AsynRelsPacket) String() string { _ = "STUB: not implemented"; return "" }
