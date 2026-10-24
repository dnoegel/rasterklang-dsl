// A `ramp ... pingpong`/`ramp ... loop` shape terminator's implied loop
// point (ADR 0003 D9) must reach a `filterTable`/`pw`/etc. field the same
// way whether the ramp is written INLINE in that field or spliced through a
// shared `table NAME = ...` declaration and a `$NAME` reference. See
// interpret.js's expandTableRefs/mergeImpliedLoop/TableDecl-case comments.
//
// Why several assertions check structure rather than only register output:
// compile.js's sampleTable()/resolveLoopIndex() default an UNSPECIFIED
// table's loop point to index 0, and expandRamp's loopIndex for both 'loop'
// and 'pingpong' is also 0. So a dropped implied loop point usually compiles
// to the same register stream by coincidence. The STRUCTURAL representation
// differs, though: song.voices[N].instrument.loop should carry an explicit
// { TABLE: 0 } entry for a $ref'd shaped ramp, exactly as the inline form
// does, and that is what is checked directly.
import { describe, expect, it } from 'vitest';
import { compileDslV2Source } from '../../src/dsl/index.js';
import { compile } from 'rasterklang-score/compiler/compile.js';
import { FILTER_REGISTERS } from 'rasterklang-score/compiler/registerMap.js';

// captureFrames(song, frameCount) -> [{addr, value}[], ...], the exact
// per-frame poke arrays compile.js's frameSource(frame) produced - used for
// direct deep-equal comparison between two songs (ref-spliced vs written-out-
// inline), which is a stronger, more direct proof of "these two compile
// identically" than hashing when the two streams are short enough to diff
// directly on a test failure.
function captureFrames(song, frameCount) {
  const frameSource = compile(song);
  const frames = [];
  for (let f = 0; f < frameCount; f += 1) frames.push(frameSource(f));
  return frames;
}

// cutoffTrace(song, frameCount) -> the cumulative $D415/$D416 filter cutoff
// value at every frame - compile.js only emits DELTA pokes (the registers
// that changed THAT frame), so the real value at any frame is the last one
// written, reconstructed here by walking the stream forward. The SID's
// 11-bit cutoff is packed as `(lo & 0x07) | (hi << 3)`.
function cutoffTrace(song, frameCount) {
  const frameSource = compile(song);
  let lo = 0;
  let hi = 0;
  const out = [];
  for (let f = 0; f < frameCount; f += 1) {
    for (const { addr, value } of frameSource(f)) {
      if (addr === FILTER_REGISTERS.cutoffLo) lo = value;
      else if (addr === FILTER_REGISTERS.cutoffHi) hi = value;
    }
    out.push((lo & 0x07) | (hi << 3));
  }
  return out;
}

describe('shared table $ref splicing propagates a ramp shape terminator\'s implied loop point (ADR 0003 D9)', () => {
  it('`loop` shape: a $ref to `table sweep = ramp ... loop` gets the SAME instrument.loop entry as writing the ramp inline would', () => {
    const refSrc = `tempo 500 beats 32
table sweep = ramp 100..0 step -25 loop
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable $sweep
end
pattern p 40 = "c4@40"
voice 0 a $p
`;
    const inlineSrc = `tempo 500 beats 32
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable ramp 100..0 step -25 loop
end
pattern p 40 = "c4@40"
voice 0 a $p
`;
    const refSong = compileDslV2Source(refSrc);
    const inlineSong = compileDslV2Source(inlineSrc);

    // Structural parity: the ref form must produce the per-table map
    // `{ filterTable: 0 }` the inline form produces, not the plain global
    // default `0` (numerically equivalent, see the header comment, but not
    // the same representation).
    expect(refSong.voices[0].instrument.loop).toEqual({ filterTable: 0 });
    expect(refSong.voices[0].instrument.loop).toEqual(inlineSong.voices[0].instrument.loop);
    expect(refSong.voices[0].instrument.filterTable).toEqual(inlineSong.voices[0].instrument.filterTable);
    expect(refSong.voices[0].instrument.filterTable).toEqual([100, 75, 50, 25, 0]);

    // Register-level parity: the two songs must compile to the exact same
    // poke stream, frame for frame - the strongest available proof that
    // splicing by reference really is indistinguishable from writing the
    // ramp out by hand, all the way down to the chip writes.
    expect(captureFrames(refSong, 20)).toEqual(captureFrames(inlineSong, 20));

    // Non-vacuous: the filter cutoff actually saturates at 0 and then jumps
    // back to 100 (loopIndex 0 = back to `from`, a discontinuous sawtooth
    // seam per ADR 0003 D9's own "loop" description) rather than holding at
    // 0 forever - proves the loop point is genuinely being read, not just
    // structurally present but unused.
    const trace = cutoffTrace(refSong, 16);
    expect(trace.slice(0, 5)).toEqual([100, 75, 50, 25, 0]);
    expect(trace.slice(5, 10)).toEqual([100, 75, 50, 25, 0]);
    expect(trace.slice(10, 15)).toEqual([100, 75, 50, 25, 0]);
  });

  it('`pingpong` shape: a $ref\'s extended (forward+backward) values AND its implied loop point both survive the splice, matching inline', () => {
    const refSrc = `tempo 500 beats 32
table sweep = ramp 100..0 step -25 pingpong
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable $sweep
end
pattern p 40 = "c4@40"
voice 0 a $p
`;
    const inlineSrc = `tempo 500 beats 32
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable ramp 100..0 step -25 pingpong
end
pattern p 40 = "c4@40"
voice 0 a $p
`;
    const refSong = compileDslV2Source(refSrc);
    const inlineSong = compileDslV2Source(inlineSrc);

    expect(refSong.voices[0].instrument.loop).toEqual({ filterTable: 0 });
    expect(refSong.voices[0].instrument.loop).toEqual(inlineSong.voices[0].instrument.loop);
    // forward [100,75,50,25,0] + reversed-excluding-both-ends [75,50,25] - see
    // expandRamp's own comment for why both ends are excluded from the
    // appended back leg.
    expect(refSong.voices[0].instrument.filterTable).toEqual([100, 75, 50, 25, 0, 25, 50, 75]);
    expect(refSong.voices[0].instrument.filterTable).toEqual(inlineSong.voices[0].instrument.filterTable);

    expect(captureFrames(refSong, 24)).toEqual(captureFrames(inlineSong, 24));

    // A seamless triangle, not a sawtooth: frame 8 (one past the 8-value
    // table) reads loopIndex 0 again, i.e. value 100 - continuing the
    // up-down-up-down shape rather than a discontinuous jump.
    const trace = cutoffTrace(refSong, 16);
    expect(trace).toEqual([100, 75, 50, 25, 0, 25, 50, 75, 100, 75, 50, 25, 0, 25, 50, 75]);
  });

  it('an UNSHAPED ramp inside a shared table (no `pingpong`/`loop` word) is completely unaffected - the common case, and the one every shipped demo tune actually uses', () => {
    const refSrc = `tempo 500 beats 32
table sweep = ramp 100..0 step -25 over 5
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable $sweep
end
pattern p 40 = "c4@40"
voice 0 a $p
`;
    const song = compileDslV2Source(refSrc);
    // No shape terminator means expandRamp's own loopIndex is null (see
    // interpret.js's expandRamp), so nothing is implied for this table at
    // all - instrument.loop stays the plain global default 0, NOT the
    // per-table map form a shaped ramp produces.
    // This is the structural signal that distinguishes "nothing implied"
    // from "loopIndex 0 implied", even though (per this file's header
    // comment) both happen to resolve to the identical register trace.
    expect(song.voices[0].instrument.loop).toBe(0);
    expect(song.voices[0].instrument.filterTable).toEqual([100, 75, 50, 25, 0]);
    // The saturate-and-hold-then-wrap-to-index-0 trace of a bare default.
    expect(cutoffTrace(song, 16)).toEqual([100, 75, 50, 25, 0, 100, 75, 50, 25, 0, 100, 75, 50, 25, 0, 100]);
  });

  it('an explicit instrument-level `loop TABLE N` still wins over a shaped shared table\'s implied loop point (mergeImpliedLoop\'s "explicit beats implied" precedent)', () => {
    const src = `tempo 500 beats 32
table sweep = ramp 100..0 step -25 loop
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable $sweep
  loop filterTable 2
end
pattern p 40 = "c4@40"
voice 0 a $p
`;
    const song = compileDslV2Source(src);
    // The explicit `loop filterTable 2` field must survive the merge, NOT
    // the ref's own implied 0 - mergeImpliedLoop's `merged = {...impliedLoop,
    // ...explicit}` spreads explicit LAST specifically so it overrides.
    expect(song.voices[0].instrument.loop).toEqual({ filterTable: 2 });
    // A genuinely observable (non-coincidental) register difference: wraps
    // back to index 2 (value 50), never revisiting 100 or 75 again - proof
    // the explicit index, not the ref's implied 0, is what's actually read.
    expect(cutoffTrace(song, 14)).toEqual([100, 75, 50, 25, 0, 50, 25, 0, 50, 25, 0, 50, 25, 0]);
  });

  it('mixed row precedence: the LAST entry in a row that carries an implied loop wins, whether it\'s a $ref or an inline ramp - one rule, not two', () => {
    // $unshaped (no loop of its own) then $shaped (loop): the shaped ref is
    // last AND the only one carrying a loop point, so it wins - loopIndex 0.
    const lastRefWins = compileDslV2Source(`tempo 500 beats 32
table plain = 900 800
table sweep = ramp 100..0 step -25 loop
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable $plain $sweep
end
pattern p 40 = "c4@40"
voice 0 a $p
`);
    expect(lastRefWins.voices[0].instrument.loop).toEqual({ filterTable: 0 });
    expect(lastRefWins.voices[0].instrument.filterTable).toEqual([900, 800, 100, 75, 50, 25, 0]);

    // $shaped (loop) then $unshaped (no loop of its own): a later entry with
    // NO implied loop must not erase the earlier one's - only a LATER entry
    // that itself carries a loop point can replace it (interpret.js's
    // expandTableRefs: `if (resolved.loopIndex != null) loopIndex = ...` -
    // deliberately guarded so a null loopIndex is a no-op, not a reset).
    const earlierRefSurvives = compileDslV2Source(`tempo 500 beats 32
table sweep = ramp 100..0 step -25 loop
table plain = 900 800
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable $sweep $plain
end
pattern p 40 = "c4@40"
voice 0 a $p
`);
    expect(earlierRefSurvives.voices[0].instrument.loop).toEqual({ filterTable: 0 });
    expect(earlierRefSurvives.voices[0].instrument.filterTable).toEqual([100, 75, 50, 25, 0, 900, 800]);

    // A $ref carrying a loop point, followed by an INLINE shaped ramp in the
    // SAME row: the inline ramp is last, so - consistent with "the LAST such
    // ramp in the row wins" for two inline ramps - it wins.
    const inlineLastWins = compileDslV2Source(`tempo 500 beats 32
table sweep = ramp 100..0 step -25 loop
instrument a
  adsr 0 8 8 8
  wave pulse
  pw 2048
  filterTable $sweep ramp 8..0 step -4 loop
end
pattern p 40 = "c4@40"
voice 0 a $p
`);
    expect(inlineLastWins.voices[0].instrument.loop).toEqual({ filterTable: 0 });
    expect(inlineLastWins.voices[0].instrument.filterTable).toEqual([100, 75, 50, 25, 0, 8, 4, 0]);
  });
});
