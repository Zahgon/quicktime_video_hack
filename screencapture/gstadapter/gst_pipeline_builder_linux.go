//go:build linux
// +build linux

package gstadapter

import "github.com/danielpaulus/gst"

func setupLivePlayAudio(pl *gst.Pipeline) {
	_ = "STUB: not implemented"

	/*hack: I do not know why, but audio on my linux box wont play when using a simple wavpars.
	  On MAC OS it works without any problems though. A hacky workaround to get audio playing that I came up with was
	  to encode audio into ogg/vorbis and directly decode it again.
	*/return
}

//endhack

func setUpVideoPipeline(pl *gst.Pipeline) *gst.AppSrc { _ = "STUB: not implemented"; return nil }

//see gst_adapter_macos comment
