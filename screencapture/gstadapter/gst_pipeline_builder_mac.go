//go:build darwin
// +build darwin

package gstadapter

import "github.com/danielpaulus/gst"

func setupLivePlayAudio(pl *gst.Pipeline) { _ = "STUB: not implemented"; return }

func setUpVideoPipeline(pl *gst.Pipeline) *gst.AppSrc { _ = "STUB: not implemented"; return nil }

// setting this to true, creates extremely choppy video
// I probably messed up something regarding the time stamps
