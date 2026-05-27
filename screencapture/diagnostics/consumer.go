package diagnostics

import (
	"io"
	"sync"
	"time"

	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// CSVHeader contains the header for the metrics file
const CSVHeader = "audioSamplesRcv, audioBytesRcv, videoSamplesRcv, videoBytesRcv, heapobjects, alloc\n"

// DiagnosticsConsumer periodically logs samples received, bytes received and memory stats to a csv file.
type DiagnosticsConsumer struct {
	outFileWriter   io.Writer
	audioSamplesRcv uint64
	videoSamplesRcv uint64
	audioBytesRcv   uint64
	videoBytesRcv   uint64
	mux             sync.Mutex
	interval        time.Duration
	stop            chan struct{}
	stopDone        chan struct{}
}

// NewDiagnosticsConsumer creates a new DiagnosticsConsumer
func NewDiagnosticsConsumer(outfile io.Writer, interval time.Duration) *DiagnosticsConsumer {
	_ = "STUB: not implemented"
	return nil
}

func fileWriter(d *DiagnosticsConsumer) { _ = "STUB: not implemented"; return }

func getMemStats() (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

func readAndReset(d *DiagnosticsConsumer) (uint64, uint64, uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Consume logs stats
func (d *DiagnosticsConsumer) Consume(buf coremedia.CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiagnosticsConsumer) consumeAudio(buf coremedia.CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiagnosticsConsumer) consumeVideo(buf coremedia.CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop writing to the csv file
func (d *DiagnosticsConsumer) Stop() { _ = "STUB: not implemented"; return }
