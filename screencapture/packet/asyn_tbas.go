package packet

// AsynTbasPacket contains info about a new Timebase. I do not know what the other reference is used for.
type AsynTbasPacket struct {
	ClockRef     CFTypeID
	SomeOtherRef CFTypeID
}

// NewAsynTbasPacketFromBytes parses a AsynTbasPacket from bytes.
func NewAsynTbasPacketFromBytes(data []byte) (AsynTbasPacket, error) {
	_ = "STUB: not implemented"
	return *new(AsynTbasPacket), nil
}

func (sp AsynTbasPacket) String() string { _ = "STUB: not implemented"; return "" }
