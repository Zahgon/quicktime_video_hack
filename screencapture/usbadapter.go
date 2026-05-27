package screencapture

import (
	"io"

	"github.com/google/gousb"
)

// UsbAdapter reads and writes from AV Quicktime USB Bulk endpoints
type UsbAdapter struct {
	outEndpoint   *gousb.OutEndpoint
	Dump          bool
	DumpOutWriter io.Writer
	DumpInWriter  io.Writer
}

// WriteDataToUsb implements the UsbWriter interface and sends the byte array to the usb bulk endpoint.
func (usbAdapter *UsbAdapter) WriteDataToUsb(bytes []byte) { _ = "STUB: not implemented"; return }

// StartReading claims the AV Quicktime USB Bulk endpoints and starts reading until a stopSignal is sent.
// Every received data is added to a frameextractor and when it is complete, sent to the UsbDataReceiver.
func (usbAdapter *UsbAdapter) StartReading(device IosDevice, receiver UsbDataReceiver, stopSignal chan interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

//the 4 bytes header are included in the length, so we need to subtract them
//here to know how long the payload will be

func clearFeature(usbDevice *gousb.Device, inboundBulkEndpointAddress gousb.EndpointAddress, outboundBulkEndpointAddress gousb.EndpointAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func findBulkEndpoint(setting gousb.InterfaceSetting, direction gousb.EndpointDirection) (int, gousb.EndpointAddress, error) {
	_ = "STUB: not implemented"
	return 0, *new(gousb.EndpointAddress), nil
}

func findAndClaimQuickTimeInterface(config *gousb.Config) (*gousb.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
