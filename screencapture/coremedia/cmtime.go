package coremedia

// Constants for the CMTime struct
const (
	KCMTimeFlagsValid                 uint32 = 0x0
	KCMTimeFlagsHasBeenRounded        uint32 = 0x1
	KCMTimeFlagsPositiveInfinity      uint32 = 0x2
	KCMTimeFlagsNegativeInfinity      uint32 = 0x4
	KCMTimeFlagsIndefinite            uint32 = 0x8
	KCMTimeFlagsImpliedValueFlagsMask        = KCMTimeFlagsPositiveInfinity | KCMTimeFlagsNegativeInfinity | KCMTimeFlagsIndefinite
	CMTimeLengthInBytes               int    = 24
)

// CMTime is taken from https://github.com/phracker/MacOSX-SDKs/blob/master/MacOSX10.8.sdk/System/Library/Frameworks/CoreMedia.framework/Versions/A/Headers/CMTime.h
type CMTime struct {
	CMTimeValue uint64 /*! @field value The value of the CMTime. value/timescale = seconds. */
	CMTimeScale uint32 /*! @field timescale The timescale of the CMTime. value/timescale = seconds.  */
	CMTimeFlags uint32 /*! @field flags The flags, eg. kCMTimeFlags_Valid, kCMTimeFlags_PositiveInfinity, etc. */
	CMTimeEpoch uint64 /*! @field epoch Differentiates between equal timestamps that are actually different because
	of looping, multi-item sequencing, etc.
	Will be used during comparison: greater epochs happen after lesser ones.
	Additions/subtraction is only possible within a single epoch,
	however, since epoch length may be unknown/variable. */
}

// GetTimeForScale calculates a float64 TimeValue by rescaling this CMTime to the CMTimeScale of the given CMTime
func (time CMTime) GetTimeForScale(newScaleToUse CMTime) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Seconds returns CMTimeValue/CMTimeScale and 0 when all values are 0
func (time CMTime) Seconds() uint64 {
	_ = "STUB: not implemented"
	// prevent division by 0
	return 0
}

// Serialize serializes this CMTime into a given byte slice that needs to be at least of CMTimeLengthInBytes length
func (time CMTime) Serialize(target []byte) error { _ = "STUB: not implemented"; return nil }

// NewCMTimeFromBytes reads a CMTime struct directly from the given byte slice
func NewCMTimeFromBytes(data []byte) (CMTime, error) {
	_ = "STUB: not implemented"
	return *new(CMTime), nil
}

func (time CMTime) String() string { _ = "STUB: not implemented"; return "" }
