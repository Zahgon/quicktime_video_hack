package packet

// ParseAsynHeader checks for the ASYN magic and the given messagemagic and then returns
// the remainingBytes starting after the messagemagic so after 16 bytes, the clockRef and an error
// is the packet is not Asyn or if the messagemagic is wrong
func ParseAsynHeader(data []byte, messagemagic uint32) ([]byte, CFTypeID, error) {
	_ = "STUB: not implemented"
	return nil, *new(CFTypeID), nil
}

// ParseSyncHeader checks for the SYNC magic and the given messagemagic and then returns
// the remainingBytes starting after the messagemagic so after 16 bytes, the clockRef, correlationID and an error
// is the packet is not SYNC or if the messagemagic is wrong
func ParseSyncHeader(data []byte, messagemagic uint32) ([]byte, CFTypeID, uint64, error) {
	_ = "STUB: not implemented"
	return nil, *new(CFTypeID), 0, nil
}

func parseHeader(data []byte, packetmagic uint32, messagemagic uint32) ([]byte, CFTypeID, error) {
	_ = "STUB: not implemented"
	return nil, *new(CFTypeID), nil
}
