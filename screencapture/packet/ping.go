package packet

// Constants for creating a Ping packet
const (
	PingPacketMagic uint32 = 0x70696E67
	PingLength      uint32 = 16
	PingHeader      uint64 = 0x0000000100000000
)

// NewPingPacketAsBytes generates a new default Ping packet
func NewPingPacketAsBytes() []byte { _ = "STUB: not implemented"; return nil }
