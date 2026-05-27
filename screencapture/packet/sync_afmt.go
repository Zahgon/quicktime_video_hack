package packet

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// SyncAfmtPacket contains what I think is information about the audio format
type SyncAfmtPacket struct {
	ClockRef                    CFTypeID
	CorrelationID               uint64
	AudioStreamBasicDescription coremedia.AudioStreamBasicDescription
}

func (sp SyncAfmtPacket) String() string { _ = "STUB: not implemented"; return "" }

// NewSyncAfmtPacketFromBytes parses a new AsynFmtPacket from byte array
func NewSyncAfmtPacketFromBytes(data []byte) (SyncAfmtPacket, error) {
	_ = "STUB: not implemented"
	return *new(SyncAfmtPacket), nil
}

// NewReply returns a []byte containing a correct reploy for afmt
func (sp SyncAfmtPacket) NewReply() []byte { _ = "STUB: not implemented"; return nil }

func createResponseDict() coremedia.StringKeyDict {
	_ = "STUB: not implemented"
	return *new(coremedia.StringKeyDict)
}
