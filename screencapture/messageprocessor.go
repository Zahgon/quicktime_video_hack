package screencapture

import (
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
	"github.com/danielpaulus/quicktime_video_hack/screencapture/packet"
)

// MessageProcessor is used to implement the control flow of USB messages answers and replies.
// It receives readily split byte frames, parses them, responds to them and passes on
// extracted CMSampleBuffers to a consumer
type MessageProcessor struct {
	usbWriter                                UsbWriter
	stopSignal                               chan interface{}
	clock                                    coremedia.CMClock
	localAudioClock                          coremedia.CMClock
	needClockRef                             packet.CFTypeID
	needMessage                              []byte
	audioSamplesReceived                     int
	videoSamplesReceived                     int
	cmSampleBufConsumer                      CmSampleBufConsumer
	clockBuilder                             func(uint64) coremedia.CMClock
	deviceAudioClockRef                      packet.CFTypeID
	releaseWaiter                            chan interface{}
	firstAudioTimeTaken                      bool
	startTimeDeviceAudioClock                coremedia.CMTime
	startTimeLocalAudioClock                 coremedia.CMTime
	lastEatFrameReceivedDeviceAudioClockTime coremedia.CMTime
	lastEatFrameReceivedLocalAudioClockTime  coremedia.CMTime
	audioOnly                                bool
}

// NewMessageProcessor creates a new MessageProcessor that will write answers to the given UsbWriter,
// forward extracted CMSampleBuffers to the CMSampleBufConsumer and wait for the stopSignal.
func NewMessageProcessor(usbWriter UsbWriter, stopSignal chan interface{}, consumer CmSampleBufConsumer, audioOnly bool) MessageProcessor {
	_ = "STUB: not implemented"
	return *new(MessageProcessor)
}

// NewMessageProcessorWithClockBuilder lets you inject a clockBuilder for the sake of testability.
func NewMessageProcessorWithClockBuilder(usbWriter UsbWriter, stopSignal chan interface{}, consumer CmSampleBufConsumer, clockBuilder func(uint64) coremedia.CMClock, audioOnly bool) MessageProcessor {
	_ = "STUB: not implemented"
	return *new(MessageProcessor)
}

// ReceiveData waits for byte frames of the correct length without the length field.
// This function will only accept byte frames starting with the ASYN, SYNC or PING uint32 magic.
func (mp *MessageProcessor) ReceiveData(data []byte) { _ = "STUB: not implemented"; return }

func (mp *MessageProcessor) handleSyncPacket(data []byte) { _ = "STUB: not implemented"; return }

func (mp *MessageProcessor) handleAsyncPacket(data []byte) { _ = "STUB: not implemented"; return }

// CloseSession shuts down the streams on the device by sending HPA0 and HPD0
// messages and waiting for RELS messages.
func (mp *MessageProcessor) CloseSession() { _ = "STUB: not implemented"; return }

func (mp MessageProcessor) stop() { _ = "STUB: not implemented"; return }
