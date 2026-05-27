package coremedia

// CMItemCount is a simple typedef to int to be a bit closer to MacOS/iOS
type CMItemCount = int

// https://github.com/phracker/MacOSX-SDKs/blob/master/MacOSX10.9.sdk/System/Library/Frameworks/CoreMedia.framework/Versions/A/Headers/CMSampleBuffer.h
const (
	sbuf uint32 = 0x73627566 //the cmsamplebuf and only content of feed asyns
	opts uint32 = 0x6F707473 //output presentation timestamp?
	stia uint32 = 0x73746961 //sampleTimingInfoArray
	sdat uint32 = 0x73646174 //the nalu
	satt uint32 = 0x73617474 //indexkey dict with only number values, CMSampleBufferGetSampleAttachmentsArray
	sary uint32 = 0x73617279 //some dict with index and one boolean
	ssiz uint32 = 0x7373697A //samplesize in bytes, size of what is contained in sdat, sample size array i think
	nsmp uint32 = 0x6E736D70 //numsample so you know how many things are in the arrays

	cmSampleTimingInfoLength = 3 * CMTimeLengthInBytes
)

// CMSampleTimingInfo is a simple struct containing 3 CMtimes: Duration, PresentationTimeStamp and DecodeTimeStamp
type CMSampleTimingInfo struct {
	Duration CMTime /*! @field duration
	The duration of the sample. If a single struct applies to
	each of the samples, they all will have this duration. */
	PresentationTimeStamp CMTime /*! @field presentationTimeStamp
	The time at which the sample will be presented. If a single
	struct applies to each of the samples, this is the presentationTime of the
	first sample. The presentationTime of subsequent samples will be derived by
	repeatedly adding the sample duration. */
	DecodeTimeStamp CMTime /*! @field decodeTimeStamp
	The time at which the sample will be decoded. If the samples
	are in presentation order, this must be set to kCMTimeInvalid. */
}

func (info CMSampleTimingInfo) String() string { _ = "STUB: not implemented"; return "" }

func (buffer CMSampleBuffer) HasSampleData() bool { _ = "STUB: not implemented"; return false }

// CMSampleBuffer represents the CoreMedia class used to exchange AV SampleData and contains meta information like timestamps or
// optional FormatDescriptors
type CMSampleBuffer struct {
	OutputPresentationTimestamp CMTime
	FormatDescription           FormatDescriptor
	HasFormatDescription        bool
	NumSamples                  CMItemCount          //nsmp
	SampleTimingInfoArray       []CMSampleTimingInfo //stia
	SampleData                  []byte
	SampleSizes                 []int
	Attachments                 IndexKeyDict //satt
	Sary                        IndexKeyDict //sary
	MediaType                   uint32
}

func (buffer CMSampleBuffer) String() string { _ = "STUB: not implemented"; return "" }

// NewCMSampleBufferFromBytesAudio parses a CMSampleBuffer containing audio data.
func NewCMSampleBufferFromBytesAudio(data []byte) (CMSampleBuffer, error) {
	_ = "STUB: not implemented"
	return *new(CMSampleBuffer), nil
}

// NewCMSampleBufferFromBytesVideo parses a CMSampleBuffer containing audio video.
func NewCMSampleBufferFromBytesVideo(data []byte) (CMSampleBuffer, error) {
	_ = "STUB: not implemented"
	return *new(CMSampleBuffer), nil
}

// NewCMSampleBufferFromBytes parses a CMSampleBuffer from a []byte assuming it begins with a 4 byte length and the 4byte magic int "sbuf"
func NewCMSampleBufferFromBytes(data []byte, mediaType uint32) (CMSampleBuffer, error) {
	_ = "STUB: not implemented"
	return *new(CMSampleBuffer), nil
}

func parseSampleSizeArray(data []byte) ([]int, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseStia(data []byte) ([]CMSampleTimingInfo, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
