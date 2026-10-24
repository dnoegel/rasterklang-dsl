// ADR 0002's drive-rate arithmetic, in one place so the JS and Go sides
// (pkg/dsl/driverate.go is the line-by-line port) cannot drift, and so every
// consumer agrees on what a `target { ... }` block means in cycles.
//
// The primitive is driveCycles: CPU cycles between two play calls, exact.
// Everything else in this module is a named macro over it, which is the
// stratification the ADR asks for.
//
// Why 19705 and not the hardware-true 19656 (312 rasterlines * 63 cycles):
// the frame grid across the toolchain is defined as machineClock/50 for PAL
// and machineClock/60 for NTSC, not as the true raster period. Adopting 19656
// here would make `target { video PAL }` mean something 0.25% different from
// the no-target-block default (exactly 50.00 fps), which ADR 0002 section 6
// requires to stay bit-identical. So PAL is round(985248/50) = 19705, and the
// default with no `target` block is that same number. Note 985248/50 is
// 19704.96, so 19705 cycles is 49.99990 fps rather than an exact 50: a 2e-5
// relative difference, far too small to move a rounded frames-per-bar count
// at any real tempo, and round(44100 * 19705 / 985248) is still exactly 882
// samples per frame.

// PAL_CPU_CLOCK_HZ / NTSC_CPU_CLOCK_HZ are the two C64 machine clocks.
export const PAL_CPU_CLOCK_HZ = 985248;
export const NTSC_CPU_CLOCK_HZ = 1022727;

// Cycles per video frame, i.e. cycles per play call for an ordinary
// single-speed tune on that machine.
export const PAL_DRIVE_CYCLES = 19705; // 985248 / 50
export const NTSC_DRIVE_CYCLES = 17045; // round(1022727 / 60)

// The default when a song declares no `target` block at all: PAL, exactly
// 50.00 play calls per second. This is DSL v2's only behaviour before ADR
// 0002, so an existing .dsl file's meaning is unchanged.
export const DEFAULT_DRIVE_CYCLES = PAL_DRIVE_CYCLES;
export const DEFAULT_CPU_CLOCK_HZ = PAL_CPU_CLOCK_HZ;

export const VIDEO_STANDARDS = ['PAL', 'NTSC'];

// videoDefaults(name) -> { cpuClockHz, driveCycles } for 'PAL' | 'NTSC',
// or null for anything else.
export function videoDefaults(name) {
  if (name === 'PAL') return { cpuClockHz: PAL_CPU_CLOCK_HZ, driveCycles: PAL_DRIVE_CYCLES };
  if (name === 'NTSC') return { cpuClockHz: NTSC_CPU_CLOCK_HZ, driveCycles: NTSC_DRIVE_CYCLES };
  return null;
}

// resolveTarget(target) -> { video, cpuClockHz, driveCycles, driveHz,
// multispeed }, the fully-expanded drive rate for a parsed TargetStmt node
// (or null/undefined for a song with no target block, which yields the
// default above).
//
// `target` is { video, driveCycles, multispeed }, each null when the
// corresponding setting was not written. Resolution order:
//   1. video (default PAL) fixes the CPU clock and the per-video-frame
//      cycle count,
//   2. an explicit driveCycles overrides that cycle count outright,
//   3. multispeed N divides the cycle count by N (N play calls per video
//      frame). It is rejected together with an explicit driveCycles by the
//      AST-transform visitor, so the two never compose here.
export function resolveTarget(target) {
  const video = (target && target.video) || 'PAL';
  const base = videoDefaults(video);
  let driveCycles = target && target.driveCycles != null ? target.driveCycles : base.driveCycles;
  const multispeed = target && target.multispeed != null ? target.multispeed : 1;
  if (multispeed > 1) driveCycles = Math.max(1, Math.round(driveCycles / multispeed));
  return {
    video,
    cpuClockHz: base.cpuClockHz,
    driveCycles,
    driveHz: base.cpuClockHz / driveCycles,
    multispeed,
  };
}
