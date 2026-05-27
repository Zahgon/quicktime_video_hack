package screencapture

import (
	"github.com/google/gousb"
)

// IosDevice contains a gousb.Device pointer for a found device and some additional info like the device usbSerial
type IosDevice struct {
	SerialNumber      string
	ProductName       string
	UsbMuxConfigIndex int
	QTConfigIndex     int
	VID               gousb.ID
	PID               gousb.ID
	UsbInfo           string
}

// OpenDevice finds a gousb.Device by using the provided iosDevice.SerialNumber. It returns an open device handle.
// Opening using VID and PID is not specific enough, as different iOS devices can have identical VID/PID combinations.
func OpenDevice(ctx *gousb.Context, iosDevice IosDevice) (*gousb.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReOpen creates a new Ios device, opening it using VID and PID, using the given context
func (d IosDevice) ReOpen(ctx *gousb.Context) (IosDevice, error) {
	_ = "STUB: not implemented"
	return *new(IosDevice), nil
}

const (
	//UsbMuxSubclass is the subclass used for USBMux USB configuration.
	UsbMuxSubclass = gousb.ClassApplication
	//QuicktimeSubclass is the subclass used for the Quicktime USB configuration.
	QuicktimeSubclass gousb.Class = 0x2A
)

// FindIosDevices finds iOS devices connected on USB ports by looking for their
// USBMux compatible Bulk Endpoints
func FindIosDevices() ([]IosDevice, error) { _ = "STUB: not implemented"; return nil, nil }

func createContext() (*gousb.Context, func()) { _ = "STUB: not implemented"; return nil, nil }

// FindIosDevice finds a iOS device by usbSerial or picks the first one if usbSerial == ""
func FindIosDevice(usbSerial string) (IosDevice, error) {
	_ = "STUB: not implemented"
	return *new(IosDevice), nil
}

func findIosDevices(ctx *gousb.Context, validDeviceChecker func(desc *gousb.DeviceDesc) bool) ([]IosDevice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this function is called for every device present.
// Returning true means the device should be opened.

func mapToIosDevice(devices []*gousb.Device) ([]IosDevice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PrintDeviceDetails returns a list of device details ready to be JSON converted.
func PrintDeviceDetails(devices []IosDevice) []map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func isValidIosDevice(desc *gousb.DeviceDesc) bool { _ = "STUB: not implemented"; return false }

func isValidIosDeviceWithActiveQTConfig(desc *gousb.DeviceDesc) bool {
	_ = "STUB: not implemented"
	return false
}

func findConfigurations(desc *gousb.DeviceDesc) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func isQtConfig(confDesc gousb.ConfigDesc) bool { _ = "STUB: not implemented"; return false }

func isMuxConfig(confDesc gousb.ConfigDesc) bool { _ = "STUB: not implemented"; return false }

func findInterfaceForSubclass(confDesc gousb.ConfigDesc, subClass gousb.Class) (bool, int) {
	_ = "STUB: not implemented"
	return false, 0
}

// IsActivated returns a boolean that is true when this device was enabled for screen mirroring and false otherwise.
func (d *IosDevice) IsActivated() bool { _ = "STUB: not implemented"; return false }

// DetailsMap contains all the info for a device in a map ready to be JSON encoded
func (d *IosDevice) DetailsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// Usually iosDevices have a 40 character USB serial which equals the usbSerial used in usbmuxd, Xcode etc.
// There is an exception, some devices like the Xr and Xs have a 24 character USB serial. Usbmux, Xcode etc.
// however insert a dash after the 8th character in this case. To be compatible with other MacOS X and iOS tools,
// we insert the dash here as well.
func Correct24CharacterSerial(usbSerial string) string { _ = "STUB: not implemented"; return "" }

const sixteenTimesZero = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"

// ValidateUdid checks if a given udid is 25 or 40 characters long.
// 25 character udids must be of format xxxxxxxx-xxxxxxxxxxxxxxxx.
// Serialnumbers on the usb host contain no dashes. As a convenience ValidateUdid
// returns the udid with the dash removed so it can be used
// as a correct USB SerialNumber.
func ValidateUdid(udid string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *IosDevice) String() string { _ = "STUB: not implemented"; return "" }
