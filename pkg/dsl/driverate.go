package dsl

import "math"

// ADR 0002's drive-rate arithmetic. Line-by-line port of
// src/dsl/driveRate.js - see that file's header for why PAL is 19705
// (machineClock/50, the frame grid the toolchain uses) rather
// than the hardware-true 19656 raster period.

const (
	// PALCPUClockHz / NTSCCPUClockHz are the two C64 machine clocks.
	PALCPUClockHz  = 985248
	NTSCCPUClockHz = 1022727

	// PALDriveCycles / NTSCDriveCycles are cycles per video frame, i.e.
	// cycles per play call for an ordinary single-speed tune.
	PALDriveCycles  = 19705 // 985248 / 50
	NTSCDriveCycles = 17045 // round(1022727 / 60)

	// DefaultDriveCycles / DefaultCPUClockHz are what a song with no
	// `target` block means: PAL, exactly 50.00 play calls per second,
	// so songs written before ADR 0002 keep their meaning.
	DefaultDriveCycles = PALDriveCycles
	DefaultCPUClockHz  = PALCPUClockHz
)

// VideoStandards is the set of names `target { video ... }` accepts.
var VideoStandards = []string{"PAL", "NTSC"}

// DriveRate is a fully resolved `target` block.
type DriveRate struct {
	Video       string
	CPUClockHz  int
	DriveCycles int
	DriveHz     float64
	Multispeed  int
}

// VideoDefaults returns the CPU clock and per-video-frame cycle count for
// "PAL" or "NTSC"; ok is false for any other name.
func VideoDefaults(name string) (clockHz, driveCycles int, ok bool) {
	switch name {
	case "PAL":
		return PALCPUClockHz, PALDriveCycles, true
	case "NTSC":
		return NTSCCPUClockHz, NTSCDriveCycles, true
	}
	return 0, 0, false
}

// ResolveTarget expands a parsed TargetStmt (nil for a song with no target
// block) into the concrete drive rate. See driveRate.js's resolveTarget for
// the resolution order.
func ResolveTarget(t *TargetStmt) DriveRate {
	video := "PAL"
	if t != nil && t.Video != "" {
		video = t.Video
	}
	clockHz, driveCycles, _ := VideoDefaults(video)
	if t != nil && t.DriveCycles != nil {
		driveCycles = *t.DriveCycles
	}
	multispeed := 1
	if t != nil && t.Multispeed != nil {
		multispeed = *t.Multispeed
	}
	if multispeed > 1 {
		driveCycles = int(math.Round(float64(driveCycles) / float64(multispeed)))
		if driveCycles < 1 {
			driveCycles = 1
		}
	}
	return DriveRate{
		Video:       video,
		CPUClockHz:  clockHz,
		DriveCycles: driveCycles,
		DriveHz:     float64(clockHz) / float64(driveCycles),
		Multispeed:  multispeed,
	}
}
