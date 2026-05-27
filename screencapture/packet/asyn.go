package packet

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// Async Packet types
const (
	AsynPacketMagic uint32 = 0x6173796E
	FEED            uint32 = 0x66656564 //These contain CMSampleBufs which contain raw h264 Nalus
	TJMP            uint32 = 0x746A6D70
	SRAT            uint32 = 0x73726174 //CMTimebaseSetRateAndAnchorTime https://developer.apple.com/documentation/coremedia/cmtimebase?language=objc
	SPRP            uint32 = 0x73707270 // Set Property
	TBAS            uint32 = 0x74626173 //TimeBase https://developer.apple.com/library/archive/qa/qa1643/_index.html
	RELS            uint32 = 0x72656C73
	HPD1            uint32 = 0x68706431 //hpd1 - 1dph | For specifying/requesting the video format
	HPA1            uint32 = 0x68706131 //hpa1 - 1aph | For specifying/requesting the audio format
	NEED            uint32 = 0x6E656564 //need - deen
	EAT             uint32 = 0x65617421 //contains audio sbufs
	HPD0            uint32 = 0x68706430
	HPA0            uint32 = 0x68706130
)

// NewAsynHpd1Packet creates a []byte containing a valid ASYN packet with the Hpd1 dictionary
func NewAsynHpd1Packet(stringKeyDict coremedia.StringKeyDict) []byte {
	_ = "STUB: not implemented"
	return nil
}

// NewAsynHpa1Packet creates a []byte containing a valid ASYN packet with the Hpa1 dictionary
func NewAsynHpa1Packet(stringKeyDict coremedia.StringKeyDict, clockRef CFTypeID) []byte {
	_ = "STUB: not implemented"
	return nil
}

func newAsynDictPacket(stringKeyDict coremedia.StringKeyDict, subtypeMarker uint32, asynTypeHeader uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// AsynNeedPacketBytes can be used to create the NEED message as soon as the clockRef from SYNC CVRP has been received.
func AsynNeedPacketBytes(clockRef CFTypeID) []byte { _ = "STUB: not implemented"; return nil }

//need - deen

// CreateHpd1DeviceInfoDict creates a dict.StringKeyDict that needs to be sent to the device before receiving a feed
func CreateHpd1DeviceInfoDict() coremedia.StringKeyDict {
	_ = "STUB: not implemented"
	return *new(coremedia.StringKeyDict)
}

// CreateHpa1DeviceInfoDict creates a dict.StringKeyDict that needs to be sent to the device before receiving a feed
func CreateHpa1DeviceInfoDict() coremedia.StringKeyDict {
	_ = "STUB: not implemented"
	return *new(coremedia.StringKeyDict)
}

// NewAsynHPD0 creates the bytes needed for stopping video streaming
func NewAsynHPD0() []byte { _ = "STUB: not implemented"; return nil }

// NewAsynHPA0 creates the bytes needed for stopping audio streaming
func NewAsynHPA0(clockRef uint64) []byte { _ = "STUB: not implemented"; return nil }
