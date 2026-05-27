package coremedia

import (
	"time"
)

// CMClock represents a monotonic Clock that will start counting when created
type CMClock struct {
	ID        uint64
	TimeScale uint32
	factor    float64
	startTime time.Time
}

// NanoSecondScale is the default system clock scale where 1/NanoSecondScale == 1 Nanosecond.
const NanoSecondScale = 1000000000

// NewCMClockWithHostTime creates a new Clock with the given ID with a nanosecond scale.
// Calls to GetTime will measure the time difference since the clock was created.
func NewCMClockWithHostTime(ID uint64) CMClock { _ = "STUB: not implemented"; return *new(CMClock) }

// NewCMClockWithHostTimeAndScale creates a new CMClock with given ID and a custom timeScale
func NewCMClockWithHostTimeAndScale(ID uint64, timeScale uint32) CMClock {
	_ = "STUB: not implemented"
	return *new(CMClock)
}

// GetTime returns a CMTime that gives the time passed since the clock started.
// This is monotonic and does NOT use wallclock time.
func (c CMClock) GetTime() CMTime { _ = "STUB: not implemented"; return *new(CMTime) }

func (c CMClock) calcValue(val int64) uint64 { _ = "STUB: not implemented"; return 0 }

// CalculateSkew calculates the deviation between the frequencies of two given clocks by using time diffs and returns a skew value float64
// scaled to match the second clock.
func CalculateSkew(startTimeClock1 CMTime, endTimeClock1 CMTime, startTimeClock2 CMTime, endTimeClock2 CMTime) float64 {
	_ = "STUB: not implemented"
	return 0
}

//println("scaleddiff:" + scaledDiff)
