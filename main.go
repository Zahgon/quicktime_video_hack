package main

import (
	"fmt"
	stdlog "log"
	"os"
	"time"

	"github.com/danielpaulus/quicktime_video_hack/screencapture"
	"github.com/danielpaulus/quicktime_video_hack/screencapture/gstadapter"
	"github.com/docopt/docopt-go"
	log "github.com/sirupsen/logrus"
)

const version = "v0.6-beta"

func main() {
	usage := fmt.Sprintf(`Q.uickTime V.ideo H.ack (qvh) %s

Usage:
  qvh devices [-v]
  qvh activate [--udid=<udid>] [-v]
  qvh deactivate [--udid=<udid>] [-v]
  qvh record <h264file> <wavfile> [--udid=<udid>] [-v]
  qvh audio <outfile> (--mp3 | --ogg | --wav) [--udid=<udid>] [-v]
  qvh gstreamer [--pipeline=<pipeline>] [--examples] [--udid=<udid>] [-v]
  qvh diagnostics <outfile> [--dump=<dumpfile>] [--udid=<udid>]
  qvh --version | version


Options:
  -h --help       Show this screen.
  -v              Enable verbose mode (debug logging).
  --version       Show version.
  --udid=<udid>   UDID of the device. If not specified, the first found device will be used automatically.

The commands work as following:
	devices		lists iOS devices attached to this host and tells you if video streaming was activated for them
	
	activate	enables the video streaming config for the device specified by --udid

	deactivate	disables the video streaming config for the device specified by --udid (in case it is stuck on streaming config)

	record		will start video&audio recording. Video will be saved in a raw h264 file playable by VLC. 
	            	Audio will be saved in a uncompressed wav file. Run like: "qvh record /home/yourname/out.h264 /home/yourname/out.wav"

	audio		Records only audio from the device. It does not change the status bar like the video recording mode does.
			The recorded audio will be saved in <outfile> with the selected format. Currently (--mp3 | --ogg | --wav) are supported.
			Adding more formats is trivial though so create an issue or a PR if you need something :-)

	gstreamer	If no additional param is provided, qvh will open a new window and push AV data to gstreamer.
			If "qvh gstreamer --examples" is provided, qvh will print some common gstreamer pipeline examples.
			If --pipeline is provided, qvh will use the provided gstreamer pipeline instead of 
			displaying audio and video in a window. 

	diagnostics	The diagnostics mode is added for running longterm tests to debug and ensure stability. 
			It will log several metrics and debug logs. Optionally specify a dump file with the --dump option that
			will store raw bytes of all messages. Be aware though, that this file will grow quite large over time. 
  `, version)
	arguments, _ := docopt.ParseDoc(usage)
	log.SetFormatter(&log.JSONFormatter{})

	verboseLoggingEnabled, _ := arguments.Bool("-v")
	if verboseLoggingEnabled {
		log.Info("Set Debug mode")
		log.SetLevel(log.DebugLevel)
	}
	stdlog.SetOutput(new(LogrusWriter))
	shouldPrintVersionNoDashes, _ := arguments.Bool("version")
	shouldPrintVersion, _ := arguments.Bool("--version")
	if shouldPrintVersionNoDashes || shouldPrintVersion {
		printVersion()
		return
	}

	devicesCommand, _ := arguments.Bool("devices")
	if devicesCommand {
		devices()
		return
	}

	udid, _ := arguments.String("--udid")
	device, err := findDevice(udid)
	if err != nil {
		printErrJSON(err, "no device found to use")
	}
	checkDeviceIsPaired(device)

	activateCommand, _ := arguments.Bool("activate")
	if activateCommand {
		activate(device)
		return
	}

	deactivateCommand, _ := arguments.Bool("deactivate")
	if deactivateCommand {
		deactivate(device)
		return
	}

	audioCommand, _ := arguments.Bool("audio")
	if audioCommand {
		outfile, err := arguments.String("<outfile>")
		if err != nil {
			printErrJSON(err, "Missing <outfile> parameter. Please specify a valid path like '/home/me/out.h264'")
			return
		}
		log.Infof("Recording audio only to file: %s", outfile)
		mp3, _ := arguments.Bool("--mp3")
		ogg, _ := arguments.Bool("--ogg")
		wav, _ := arguments.Bool("--wav")
		log.Debugf("recording audio only format mp3:%t ogg: %t wav:%t to file: %s", mp3, ogg, wav, outfile)
		if wav {
			recordAudioWav(outfile, device)
			return
		}
		if ogg {
			recordAudioGst(outfile, device, gstadapter.OGG)
			return
		}
		recordAudioGst(outfile, device, gstadapter.MP3)
		return
	}

	diagnostics, _ := arguments.Bool("diagnostics")
	if diagnostics {
		log.SetLevel(log.DebugLevel)
		logfileName := fmt.Sprintf("logfile-%d.log", time.Now().Unix())
		logfile, err := os.OpenFile(logfileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			println("logging to", logfileName, " execute: 'tail -f ", logfileName, "' for logs. Press CTRL+C to stop recording.")
			log.SetOutput(logfile)
		} else {
			log.Info("Failed to log to file, using default stderr")
		}

		outfile, err := arguments.String("<outfile>")
		if err != nil {
			printErrJSON(err, "Missing <outfile> parameter. Please specify a valid path like '/home/me/out.json'")
			return
		}
		dump, _ := arguments.String("--dump")
		runDiagnostics(outfile, dump != "", dump, device)
		return
	}

	recordCommand, _ := arguments.Bool("record")
	if recordCommand {
		h264FilePath, err := arguments.String("<h264file>")
		if err != nil {
			printErrJSON(err, "Missing <h264file> parameter. Please specify a valid path like '/home/me/out.h264'")
			return
		}
		waveFilePath, err := arguments.String("<wavfile>")
		if err != nil {
			printErrJSON(err, "Missing <wavfile> parameter. Please specify a valid path like '/home/me/out.raw'")
			return
		}
		record(h264FilePath, waveFilePath, device)
	}
	gstreamerCommand, _ := arguments.Bool("gstreamer")
	if gstreamerCommand {
		shouldPrintExamples, _ := arguments.Bool("--examples")
		if shouldPrintExamples {
			printExamples()
			return
		}
		gstPipeline, _ := arguments.String("--pipeline")
		if gstPipeline == "" {
			startGStreamer(device)
			return
		}
		startGStreamerWithCustomPipeline(device, gstPipeline)
	}
}

// findDevice grabs the first device on the host for a empty --udid
// or tries to find the provided device otherwise
func findDevice(udid string) (screencapture.IosDevice, error) {
	_ = "STUB: not implemented"
	return *new(screencapture.IosDevice), nil
}

func printVersion() { _ = "STUB: not implemented"; return }

func printExamples() { _ = "STUB: not implemented"; return }

func recordAudioGst(outfile string, device screencapture.IosDevice, audiotype string) {
	_ = "STUB: not implemented"
	return
}

func runDiagnostics(outfile string, dump bool, dumpFile string, device screencapture.IosDevice) {
	_ = "STUB: not implemented"
	return
}

func recordAudioWav(outfile string, device screencapture.IosDevice) {
	_ = "STUB: not implemented"
	return
}

func startGStreamerWithCustomPipeline(device screencapture.IosDevice, pipelineString string) {
	_ = "STUB: not implemented"
	return
}

func startGStreamer(device screencapture.IosDevice) { _ = "STUB: not implemented"; return }

// Just dump a list of what was discovered to the console
func devices() { _ = "STUB: not implemented"; return }

// This command is for testing if we can enable the hidden Quicktime device config
func activate(device screencapture.IosDevice) { _ = "STUB: not implemented"; return }

func deactivate(device screencapture.IosDevice) { _ = "STUB: not implemented"; return }

func record(h264FilePath string, wavFilePath string, device screencapture.IosDevice) {
	_ = "STUB: not implemented"
	return
}

func startWithConsumer(consumer screencapture.CmSampleBufConsumer, device screencapture.IosDevice, audioOnly bool) {
	_ = "STUB: not implemented"
	return
}

func startWithConsumerDump(consumer screencapture.CmSampleBufConsumer, device screencapture.IosDevice, dumpPath string) {
	_ = "STUB: not implemented"
	return
}

func waitForSigInt(stopSignalChannel chan interface{}) { _ = "STUB: not implemented"; return }

func checkDeviceIsPaired(device screencapture.IosDevice) { _ = "STUB: not implemented"; return }

func printErrJSON(err error, msg string) { _ = "STUB: not implemented"; return }

func printJSON(output map[string]interface{}) { _ = "STUB: not implemented"; return }

// this is to ban these irritating "2021/04/29 14:27:59 handle_events: error: libusb: interrupted [code -10]" libusb messages
type LogrusWriter int

const interruptedError = "interrupted [code -10]"

func (LogrusWriter) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
