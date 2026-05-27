package packet

// AsynTjmpPacket contains the data from a TJMP packet.
// I think this is a notification sent by the device about changing a TimeBase.
// I do not know what the last bytes are for currently.
type AsynTjmpPacket struct {
	ClockRef CFTypeID
	Unknown  []byte
}

// NewAsynTjmpPacketFromBytes parses a new AsynTjmpPacket from byte array
func NewAsynTjmpPacketFromBytes(data []byte) (AsynTjmpPacket, error) {
	_ = "STUB: not implemented"
	return *new(AsynTjmpPacket), nil
}

func (sp AsynTjmpPacket) String() string { _ = "STUB: not implemented"; return "" }
