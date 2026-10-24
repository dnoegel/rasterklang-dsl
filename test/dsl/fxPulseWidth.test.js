// Suite for the `fx` lane's `pw.N` command (per-row one-shot pulse-width
// overwrite). Pulse width was the only SID register class with no fx
// spelling at all, while being the largest class of register writes a
// decompiler could not otherwise express. See src/dsl/fx.js's `pw` entry for
// the rationale.
//
// `pw.N` is deliberately modelled on `wave.N` at every layer (parse ->
// { op, a } -> a sticky per-voice runtime override cleared at the next note
// onset -> a register write that yields to an instrument's own table), so
// most assertions below are stated as "the same shape wave.N has", and the
// wave-interop test exercises both commands on one voice at once.
//
// Assertions are on the COMPILED POKE STREAM rather than on "it parsed", in
// the style of sharedTableImpliedLoop.test.js: the whole point of the command
// is which bytes reach $D402/$D403, so that is what is checked.
import { describe, expect, it } from 'vitest';
import { compileDslV2Source } from '../../src/dsl/index.js';
import { parseFxToken, FX_OPS } from '../../src/dsl/fx.js';
import { compile } from 'rasterklang-score/compiler/compile.js';
import { VOICE_REGISTER_OFFSETS, voiceBase } from 'rasterklang-score/compiler/registerMap.js';

// pulseWidthTrace(src, frameCount, voice) -> the 12-bit pulse width voice N's
// $02/$03 registers actually hold at every frame. compile.js emits DELTA
// pokes only (just the registers that changed this frame), so the value at
// any frame is the last one written - reconstructed by walking the stream
// forward, exactly as sharedTableImpliedLoop.test.js's cutoffTrace() does for
// the filter cutoff. The high register holds only the top 4 bits (the SID's
// pulse width is 12 bits wide, the upper nibble of $03 is unused), hence the
// `& 0x0f`.
function pulseWidthTrace(src, frameCount, voice = 0) {
  const frameSource = compile(compileDslV2Source(src));
  const base = voiceBase(voice);
  let lo = 0;
  let hi = 0;
  const out = [];
  for (let f = 0; f < frameCount; f += 1) {
    for (const { addr, value } of frameSource(f)) {
      if (addr === base + VOICE_REGISTER_OFFSETS.pwLo) lo = value;
      else if (addr === base + VOICE_REGISTER_OFFSETS.pwHi) hi = value;
    }
    out.push(lo | ((hi & 0x0f) << 8));
  }
  return out;
}

// controlTrace(src, frameCount, voice) -> the same walk-forward reconstruction
// for the control register ($04), used by the wave-interop test to prove the
// two commands' overrides are independent rather than one clobbering the
// other.
function controlTrace(src, frameCount, voice = 0) {
  const frameSource = compile(compileDslV2Source(src));
  const base = voiceBase(voice);
  let control = 0;
  const out = [];
  for (let f = 0; f < frameCount; f += 1) {
    for (const { addr, value } of frameSource(f)) {
      if (addr === base + VOICE_REGISTER_OFFSETS.control) control = value;
    }
    out.push(control);
  }
  return out;
}

// `tempo 1000 beats 4` is 12 frames per cycle, so a 4-step fx string is
// exactly 3 frames per row - short enough that the traces below can be read
// as literal row-by-row expectations.
const FRAMES_PER_ROW = 3;

describe('fx `pw.N` parsing', () => {
  it('is a 1-parameter op in FX_OPS, like `wave`', () => {
    expect(FX_OPS.pw).toBe(1);
    expect(FX_OPS.pw).toBe(FX_OPS.wave);
  });

  it('parses to the same { op, a: { value } } shape `wave.N` does', () => {
    expect(parseFxToken('pw.512')).toEqual({ op: 'pw', a: { value: 512 } });
    // Both ends of the SID's 12-bit domain are legal values, not edge cases.
    expect(parseFxToken('pw.0')).toEqual({ op: 'pw', a: { value: 0 } });
    expect(parseFxToken('pw.4095')).toEqual({ op: 'pw', a: { value: 4095 } });
  });

  it('rejects a value outside the SID\'s 12-bit pulse-width domain, naming the range', () => {
    // The failure mode this guards is a real one for the importer feeding
    // this lane: mixing up an 8-bit tracker pulse value with the 12-bit
    // register one. Clamping would hide it; erroring names it.
    expect(() => parseFxToken('pw.4096')).toThrow(/pulse width 4096 is out of range - must be 0-4095/);
    expect(() => parseFxToken('pw.-1')).toThrow(/pulse width -1 is out of range - must be 0-4095/);
  });

  it('rejects the shift-derived `sN` parameter form, which is frequency-only', () => {
    // `sN` means "one semitone's worth of FREQUENCY at the current note,
    // shifted right by N" (see fx.js's parseParam) - a note-dependent pitch
    // quantity with no meaning for a pulse-width register.
    expect(() => parseFxToken('pw.s3')).toThrow(/shift-derived 'sN' form is frequency-only/);
  });

  it('still enforces the shared arity rule', () => {
    expect(() => parseFxToken('pw')).toThrow(/takes exactly 1 parameter\(s\), got 0/);
    expect(() => parseFxToken('pw.100.200')).toThrow(/takes exactly 1 parameter\(s\), got 2/);
  });

  it('surfaces an out-of-range value through a whole-source compile, with the line and voice, exactly as an unknown effect token does', () => {
    const src = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
  wave pulse
end
voice 0 lead
  notes "c4@4"
  fx "pw.9999"
end
`;
    // Like every other fx token error, this surfaces when the pattern is
    // first sampled (interpret.js maps fx values lazily), so the check runs
    // the first frame rather than only building the song.
    expect(() => compile(compileDslV2Source(src))(0)).toThrow(
      // Line 6 is the `voice` statement itself, which is the line interpret.js
      // attributes every fx token error to (vstmt.line), not the `fx` field's
      // own line - the same attribution an unknown effect token gets.
      /line 6: voice 0: fx token 'pw\.9999': pulse width 9999 is out of range/,
    );
  });
});

describe('fx `pw.N` register effect', () => {
  // A table-less instrument: the pulse width comes from the chain's own
  // default (2048), so every change in the trace below is the fx command's
  // doing and nothing else's.
  const src = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
  wave pulse
end
voice 0 lead
  notes "c4@4"
  fx "~ pw.512 ~ pw.4095"
end
`;

  it('writes the commanded 12-bit value to $02/$03 and holds it until the next command', () => {
    const trace = pulseWidthTrace(src, 4 * FRAMES_PER_ROW);
    expect(trace).toEqual([
      // row 0 (no command): the chain default, 2048.
      2048, 2048, 2048,
      // row 1: pw.512 - the command fires on the frame its hap becomes
      // active and STAYS in effect (rows 2's `~` does not restore anything),
      // the same sticky lifetime fx `wave`/`ad`/`sr` have.
      512, 512, 512,
      512, 512, 512,
      // row 3: pw.4095, the top of the domain.
      4095, 4095, 4095,
    ]);
  });

  it('emits exactly the register bytes the split implies, on the frame the command fires', () => {
    const frameSource = compile(compileDslV2Source(src));
    const frames = [];
    for (let f = 0; f < 4 * FRAMES_PER_ROW; f += 1) frames.push(frameSource(f));
    // 512 = 0x200: low byte 0x00 (unchanged from 2048's own 0x00, so it is
    // correctly NOT re-poked by the delta encoder), high nibble 0x02.
    expect(frames[3]).toEqual([{ addr: 3, value: 2 }]);
    // 4095 = 0x0fff: both bytes change.
    expect(frames[9]).toEqual([
      { addr: 2, value: 255 },
      { addr: 3, value: 15 },
    ]);
    // The rows in between poke nothing at all: a one-shot command, not a
    // per-frame rewrite.
    expect(frames[4]).toEqual([]);
    expect(frames[10]).toEqual([]);
  });

  it('is cleared by the next note ONSET, like fx `wave`/`ad`/`sr`', () => {
    const twoNotes = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
  wave pulse
end
voice 0 lead
  notes "c4@2 e4@2"
  fx "~ pw.512 ~ ~"
end
`;
    const trace = pulseWidthTrace(twoNotes, 4 * FRAMES_PER_ROW);
    expect(trace.slice(0, 3)).toEqual([2048, 2048, 2048]);
    expect(trace.slice(3, 6)).toEqual([512, 512, 512]);
    // The e4 onset at frame 6 re-establishes the instrument's own value -
    // gplay.c's new-note init block writing over a command's register poke.
    expect(trace.slice(6, 12)).toEqual([2048, 2048, 2048, 2048, 2048, 2048]);
  });

  it('yields to an instrument\'s own `pw` TABLE, exactly as `wave.N` yields to a `wave` table', () => {
    const withTable = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
  wave pulse
  pw 1000 1100
end
voice 0 lead
  notes "c4@4"
  fx "~ pw.512 ~ ~"
end
`;
    // A per-frame table rewrites this register every frame it advances, so a
    // one-shot value would be overwritten anyway - the command is ignored
    // outright rather than flickering for one frame. The trace is the bare
    // alternating table, with no 512 anywhere in it.
    const trace = pulseWidthTrace(withTable, 4 * FRAMES_PER_ROW);
    expect(trace).toEqual([1000, 1100, 1000, 1100, 1000, 1100, 1000, 1100, 1000, 1100, 1000, 1100]);
    expect(trace).not.toContain(512);
  });

  it('takes the same 0-4095 value domain as the instrument `pw` column: the same number spelled either way reaches the same registers', () => {
    // The point of pinning `pw.N` to the table column's domain is that the
    // two spellings cannot disagree. A static `pw 3000` table and a
    // `fx "pw.3000"` command must produce identical $02/$03 bytes.
    const asColumn = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
  wave pulse
  pw 3000
end
voice 0 lead
  notes "c4@4"
end
`;
    const asCommand = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
  wave pulse
end
voice 0 lead
  notes "c4@4"
  fx "pw.3000"
end
`;
    expect(pulseWidthTrace(asColumn, 12)).toEqual(new Array(12).fill(3000));
    expect(pulseWidthTrace(asCommand, 12)).toEqual(new Array(12).fill(3000));
  });
});

describe('fx `pw.N` interoperating with the other fx tokens on the same voice', () => {
  // No instrument `wave`/`pw` tables here (the waveform comes from the
  // voice-level `wave "pulse"` PATTERN, chain.waveform), so both the `wave`
  // and the `pw` command are live - the configuration where the two can
  // actually interfere if either override were mis-scoped.
  const src = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
end
voice 0 lead
  notes "c4@4"
  wave "pulse"
  fx "~ pw.512 wave.16 pw.3000"
end
`;

  it('`pw` and `wave` overrides are independent and both sticky', () => {
    const pw = pulseWidthTrace(src, 4 * FRAMES_PER_ROW);
    const control = controlTrace(src, 4 * FRAMES_PER_ROW);
    // Row 1 sets pw.512; row 2's wave.16 (triangle) must NOT disturb it.
    expect(pw).toEqual([2048, 2048, 2048, 512, 512, 512, 512, 512, 512, 3000, 3000, 3000]);
    // Control byte: 0x41 = pulse (0x40) + gate, then 0x11 = triangle (0x10)
    // + gate from row 2 on. Row 3's pw.3000 must NOT disturb the waveform
    // override either - it stays 0x11 to the end.
    expect(control).toEqual([0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11]);
  });

  it('does not disturb the envelope registers a neighbouring `ad`/`sr` row sets', () => {
    const withEnvelope = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
end
voice 0 lead
  notes "c4@4"
  wave "pulse"
  fx "ad.1.2 pw.512 sr.3.4 pw.3000"
end
`;
    const frameSource = compile(compileDslV2Source(withEnvelope));
    let ad = 0;
    let sr = 0;
    const envelope = [];
    for (let f = 0; f < 4 * FRAMES_PER_ROW; f += 1) {
      for (const { addr, value } of frameSource(f)) {
        if (addr === VOICE_REGISTER_OFFSETS.attackDecay) ad = value;
        else if (addr === VOICE_REGISTER_OFFSETS.sustainRelease) sr = value;
      }
      envelope.push([ad, sr]);
    }
    // ad.1.2 packs to 0x12 and holds through both pulse-width rows; sr.3.4
    // packs to 0x34 from row 2 on. The instrument's own `adsr 0 8 8 8` gives
    // 0x08/0x88 before each is commanded.
    expect(envelope[0]).toEqual([0x12, 0x88]);
    expect(envelope[5]).toEqual([0x12, 0x88]);
    expect(envelope[8]).toEqual([0x12, 0x34]);
    expect(envelope[11]).toEqual([0x12, 0x34]);
    // And the pulse width still did what it was told, alongside all of that.
    expect(pulseWidthTrace(withEnvelope, 12)).toEqual([2048, 2048, 2048, 512, 512, 512, 512, 512, 512, 3000, 3000, 3000]);
  });

  it('works on a voice other than 0, writing that voice\'s own $02/$03 pair', () => {
    const src2 = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
  wave pulse
end
voice 2 lead
  notes "c4@4"
  fx "~ pw.777 ~ ~"
end
`;
    // voiceBase(2) = 14, so the pulse registers are $0e/$0f - proof the
    // command is per-voice state (runtime.fxPw lives in voiceRuntime), not a
    // song-global one like `fcut`/`vol`.
    expect(pulseWidthTrace(src2, 12, 2)).toEqual([2048, 2048, 2048, 777, 777, 777, 777, 777, 777, 777, 777, 777]);
    // Voice 0 is unassigned in this song and must be untouched by it.
    expect(pulseWidthTrace(src2, 12, 0)).toEqual(new Array(12).fill(0));
  });
});

describe('fx `pw.N` is purely additive', () => {
  it('a song that uses no `pw.N` compiles to the identical poke stream with the other fx tokens unchanged', () => {
    // This song exercises every other one-shot register command; the hash
    // pins its stream so a change to the fx dispatch that perturbs the other
    // commands fails loudly. (The value below was produced by this suite.)
    const src = `tempo 1000 beats 4
instrument lead
  adsr 0 8 8 8
end
voice 0 lead
  notes "c4@4"
  wave "pulse"
  fx "wave.16 ad.1.2 sr.3.4 vol.7"
end
`;
    const frameSource = compile(compileDslV2Source(src));
    const frames = [];
    for (let f = 0; f < 12; f += 1) frames.push(frameSource(f));
    // Pulse width never leaves the chain default: no `pw.N` in the song, so
    // runtime.fxPw stays null.
    expect(pulseWidthTrace(src, 12)).toEqual(new Array(12).fill(2048));
    // Every command that IS present still lands.
    expect(frames[3]).toEqual([{ addr: 5, value: 0x12 }]);
    expect(frames[6]).toEqual([{ addr: 6, value: 0x34 }]);
    expect(frames[9]).toEqual([{ addr: 24, value: 7 }]);
  });
});
