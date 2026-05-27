package gstadapter

import (
	"github.com/danielpaulus/gst"
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
)

// GstAdapter contains the AppSrc for accessing Gstreamer.
type GstAdapter struct {
	videoAppSrc      *gst.AppSrc
	audioAppSrc      *gst.AppSrc
	pipeline         *gst.Pipeline
	firstAudioSample bool
}

const audioAppSrcTargetElementName = "audio_target"
const videoAppSrcTargetElementName = "video_target"

const MP3 = "mp3"
const OGG = "ogg"

// New creates a new MAC OSX compatible gstreamer pipeline that will play device video and audio
// in a nice little window :-D
func New() *GstAdapter { _ = "STUB: not implemented"; return nil }

func NewWithAudioPipeline(outfile string, audiotype string) (*GstAdapter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWithCustomPipeline will parse the given pipelineString, connect the videoAppSrc to whatever element has the name "video_target" and the audioAppSrc to "audio_target"
// see also: https://gstreamer.freedesktop.org/documentation/application-development/appendix/programs.html?gi-language=c
func NewWithCustomPipeline(pipelineString string) (*GstAdapter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//runGlibMainLoop()

// Stop sends an EOS (end of stream) event downstream the gstreamer pipeline.
// Some Elements need this to correctly finish. F.ex. writing mp4 video without
// sending EOS will result in a broken mp4 file
func (gsta GstAdapter) Stop() { _ = "STUB: not implemented"; return }

//I hope those are 60 seconds

// runGlibMainLoop starts the glib Mainloop necessary for the video player to work on MAC OS X.
func runGlibMainLoop() {
	_ = "STUB: not implemented"

	// See: https://golang.org/pkg/runtime/#LockOSThread
	return
}

func setUpAudioPipelineBase(pl *gst.Pipeline) *gst.AppSrc { _ = "STUB: not implemented"; return nil }

func setupVorbis(pl *gst.Pipeline, filepath string) {
	_ = "STUB: not implemented"
	// vorbisenc ! oggmux ! filesink location=alsasrc.ogg
	return
}

func setupMp3(pl *gst.Pipeline, filepath string) {
	_ = "STUB: not implemented"
	// lamemp3enc ! filesink location=sine.mp3
	return
}

func checkElem(e *gst.Element, name string) { _ = "STUB: not implemented"; return }

// Consume will transfer AV data into a Gstreamer AppSrc
func (gsta *GstAdapter) Consume(buf coremedia.CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

//FIXME: ugly hack I added to prevent gstreamer from receiving decreasing timestamps
//I might have messed something up while sending times to the device as my first
//buffer will have this weird, large timestamp. So I hack it to be equal to zero here

//see above comment

func (gsta GstAdapter) sendWavHeader() { _ = "STUB: not implemented"; return }

//TODO: create CGO function that provides offsets so we can delete prependMarker again

func (gsta GstAdapter) sendAudioSample(buf coremedia.CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

//TODO: create CGO function that provides offsets so we can delete prependMarker again

func (gsta GstAdapter) writeNalus(bytes coremedia.CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (gsta GstAdapter) writeNalu(naluBytes []byte, buf coremedia.CMSampleBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

//TODO: create CGO function that provides offsets so we can delete prependMarker again

var naluAnnexBMarkerBytes = []byte{0, 0, 0, 1}

func prependMarker(nalu []byte, length uint32) []byte { _ = "STUB: not implemented"; return nil }
