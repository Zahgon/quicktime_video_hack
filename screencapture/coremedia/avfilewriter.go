package coremedia

import (
	"io"
)

var startCode = []byte{00, 00, 00, 01}

// AVFileWriter writes nalus into a file using 0x00000001 as a separator (h264 ANNEX B) and raw pcm audio into a wav file
// Note that you will have to call WriteWavHeader() on the audiofile when you are done to write a wav header and get a valid file.
type AVFileWriter struct {
	h264FileWriter io.Writer
	wavFileWriter  io.Writer
	outFilePath    string
	audioOnly      bool
}

// NewAVFileWriter binary writes nalus in annex b format to the given writer and audio buffers into a wav file.
// Note that you will have to call WriteWavHeader() on the audiofile when you are done to write a wav header and get a valid file.
func NewAVFileWriter(h264FileWriter io.Writer, wavFileWriter io.Writer) AVFileWriter {
	_ = "STUB: not implemented"
	return *new(AVFileWriter)
}

func NewAVFileWriterAudioOnly(wavFileWriter io.Writer) AVFileWriter {
	_ = "STUB: not implemented"
	return *new(AVFileWriter)
}

// Consume writes PPS and SPS as well as sample bufs into a annex b .h264 file and audio samples into a wav file
// Note that you will have to call WriteWavHeader() on the audiofile when you are done to write a wav header and get a valid file.
func (avfw AVFileWriter) Consume(buf CMSampleBuffer) error { _ = "STUB: not implemented"; return nil }

// Nothing currently
func (avfw AVFileWriter) Stop() { _ = "STUB: not implemented"; return }

func (avfw AVFileWriter) consumeVideo(buf CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (avfw AVFileWriter) writeNalus(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (avfw AVFileWriter) writeNalu(naluBytes []byte) error { _ = "STUB: not implemented"; return nil }

func (avfw AVFileWriter) consumeAudio(buffer CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}
