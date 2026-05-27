package screencapture

import (
	"github.com/google/gousb"
)

// EnableQTConfig enables the hidden QuickTime Device configuration that will expose two new bulk endpoints.
// We will send a control transfer to the device via USB which will cause the device to disconnect and then
// re-connect with a new device configuration. Usually the usbmuxd will automatically enable that new config
// as it will detect it as the device's preferredConfig.
func EnableQTConfig(device IosDevice) (IosDevice, error) {
	_ = "STUB: not implemented"
	return *new(IosDevice), nil
}

func DisableQTConfig(device IosDevice) (IosDevice, error) {
	_ = "STUB: not implemented"
	return *new(IosDevice), nil
}

func sendQTConfigControlRequest(device *gousb.Device) { _ = "STUB: not implemented"; return }

func sendQTDisableConfigControlRequest(device *gousb.Device) { _ = "STUB: not implemented"; return }
