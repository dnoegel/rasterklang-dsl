// Suite for ADR 0005 (docs/adr/0005-player-idioms-from-unitracker.md): the
// five player idioms taken from comparing unitracker with this language's
// converter gaps. D1 table ops (wtab/atab/ptab/ftab), D2 filter control
// (res/route/fmode/fctl), D3 sweeps (pws/pwm/fsw), D4 hardRestart clauses,
// D5 groove.
//
// Like fxPulseWidth.test.js, assertions are on the compiled register state
// per frame, since what each construct promises is which bytes reach the
// chip on which frame.
import { describe, expect, it } from 'vitest';
import { compileDslV2Source } from '../../src/dsl/index.js';
import { parseFxToken } from '../../src/dsl/fx.js';
import { compile } from 'rasterklang-score/compiler/compile.js';
import { VOICE_REGISTER_OFFSETS, FILTER_REGISTERS, voiceBase, noteToFreqRegister } from 'rasterklang-score/compiler/registerMap.js';

// registerFrames(src, frameCount) -> one 32-byte register image per frame,
// the state after that frame's pokes (compile.js emits deltas only).
function registerFrames(src, frameCount) {
  const frameSource = compile(compileDslV2Source(src));
  const regs = new Array(32).fill(0);
  const out = [];
  for (let f = 0; f < frameCount; f += 1) {
    for (const { addr, value } of frameSource(f)) regs[addr] = value;
    out.push(regs.slice());
  }
  return out;
}

const pw = (frame, voice = 0) => {
  const base = voiceBase(voice);
  return frame[base + VOICE_REGISTER_OFFSETS.pwLo] | ((frame[base + VOICE_REGISTER_OFFSETS.pwHi] & 0x0f) << 8);
};
const freq = (frame, voice = 0) => {
  const base = voiceBase(voice);
  return frame[base + VOICE_REGISTER_OFFSETS.freqLo] | (frame[base + VOICE_REGISTER_OFFSETS.freqHi] << 8);
};
const reg = (frame, offset, voice = 0) => frame[voiceBase(voice) + offset];
const cutoff = (frame) => (frame[FILTER_REGISTERS.cutoffHi] << 3) | (frame[FILTER_REGISTERS.cutoffLo] & 0x07);

// A 4-frame beat. A proportional pattern fills one bar, so each test sets
// `beats` to its step count and every step is exactly 4 frames long.
const clock = (steps) => `speed 4 beats ${steps}`;
const CLOCK = clock(2);

describe('ADR 0005 D1: fx table ops start a named table from a row', () => {
  it('ptab replaces the pw table from the fx frame, on its own clock, and wraps to 0', () => {
    const src = `${CLOCK}
table sweep = 100 200 300
instrument lead
  wave pulse
  pw 1000 1000
end
voice 0 lead
  notes "c4 _"
  fx "~ ptab.sweep"
end
`;
    const frames = registerFrames(src, 8);
    expect(frames.slice(0, 4).map((f) => pw(f))).toEqual([1000, 1000, 1000, 1000]);
    expect(frames.slice(4, 8).map((f) => pw(f))).toEqual([100, 200, 300, 100]);
  });

  it('an explicit loop index wins over the table default', () => {
    const src = `${CLOCK}
table sweep = 100 200 300
instrument lead
  wave pulse
end
voice 0 lead
  notes "c4"
  fx "ptab.sweep.2"
end
`;
    const frames = registerFrames(src, 6);
    expect(frames.map((f) => pw(f))).toEqual([100, 200, 300, 300, 300, 300]);
  });

  it('the next note onset restores the instrument table; a legato tp transition does not', () => {
    const src = `${CLOCK}
table sweep = 100
instrument lead
  wave pulse
  pw 1000
end
voice 0 lead
  notes "c4 e4"
  fx "ptab.sweep ~"
end
voice 1 lead
  notes "c4 e4"
  fx "ptab.sweep tp.0"
end
`;
    const frames = registerFrames(src, 8);
    expect(pw(frames[2], 0)).toBe(100);
    expect(pw(frames[5], 0)).toBe(1000);
    expect(pw(frames[5], 1)).toBe(100);
  });

  it('wtab, atab and ftab drive the waveform, the pitch and the shared cutoff', () => {
    const src = `${clock(3)}
table w = noise pulse
table a = 0 12
table f = 800 400
instrument lead
  wave triangle
end
voice 0 lead
  notes "c4 c4 c4"
  fx "wtab.w atab.a ftab.f"
end
`;
    const frames = registerFrames(src, 12);
    const wave = (f) => reg(f, VOICE_REGISTER_OFFSETS.control) & 0xf0;
    expect([wave(frames[0]), wave(frames[1])]).toEqual([0x80, 0x40]);
    expect(wave(frames[4])).toBe(0x10);
    expect([freq(frames[4]), freq(frames[5])]).toEqual([noteToFreqRegister(60), noteToFreqRegister(72)]);
    expect([cutoff(frames[8]), cutoff(frames[9])]).toEqual([800, 400]);
  });

  it('rejects an unknown table and a word cell in a pw table', () => {
    // fx tokens resolve when the pattern is first queried, as every fx parse
    // error does (see fxPulseWidth.test.js), hence the first frame.
    const base = `${CLOCK}\ntable w = pulse noise\nvoice 0 "c4" wave:pulse fx:`;
    expect(() => compile(compileDslV2Source(`${base}"ptab.nope"\n`))(0)).toThrow(/unknown table 'nope'/);
    expect(() => compile(compileDslV2Source(`${base}"ptab.w"\n`))(0)).toThrow(/pw table takes numbers only/);
  });
});

describe('ADR 0005 D2: per-row filter control', () => {
  it('res, route, fmode and fctl override the filter block from their frame on', () => {
    const src = `${clock(4)}
voice 0
  notes "c4 c4 c4 c4"
  wave "pulse"
  fx "res.9 route.5 fmode.9 fctl.3.2"
end
filter
  cutoff sine 300..300
  resonance 1
  mode lowpass
  route 0
end
`;
    const frames = registerFrames(src, 16);
    const rr = (f) => f[FILTER_REGISTERS.resonanceRouting];
    const mv = (f) => f[FILTER_REGISTERS.modeVolume];
    expect(rr(frames[0])).toBe(0x91);
    expect(rr(frames[4])).toBe(0x95);
    expect(mv(frames[8])).toBe(0x9f);
    expect(rr(frames[12])).toBe(0x32);
    expect(mv(frames[12])).toBe(0x9f);
  });

  it('rejects out-of-range nibbles instead of clamping', () => {
    expect(() => parseFxToken('res.16')).toThrow(/must be 0-15/);
    expect(() => parseFxToken('fctl.1.16')).toThrow(/parameter 2 of 'fctl'/);
  });
});

describe('ADR 0005 D3: register sweeps', () => {
  // (An inline `pw:` attr would be an instrument pw table, which outranks a
  // sweep, so the start value comes from fx pw.N here.)
  it('pws adds per frame and wraps at 12 bits', () => {
    const src = `${CLOCK}
voice 0 "c4 _" wave:pulse fx:"pw.4090 pws.4"
`;
    const frames = registerFrames(src, 8);
    expect(frames.slice(4).map((f) => pw(f))).toEqual([4094, 2, 6, 10]);
  });

  it('pwm bounces between LO and HI', () => {
    const src = `${clock(3)}
voice 0 "c4 _ _" wave:pulse fx:"pw.100 pwm.100.130.10@2"
`;
    const frames = registerFrames(src, 12);
    expect(frames.slice(4).map((f) => pw(f))).toEqual([110, 120, 130, 120, 110, 100, 110, 120]);
  });

  it('a sweep spanning a note onset keeps running; the value then sticks until the next onset', () => {
    const src = `${clock(4)}
voice 0 "c4 e4 g4 ~" wave:pulse fx:"pws.10@2 ~ ~"
`;
    const frames = registerFrames(src, 12);
    expect(pw(frames[3])).toBe(2088);
    expect(pw(frames[4])).toBe(2098);
    expect(pw(frames[7])).toBe(2128);
    expect(pw(frames[8])).toBe(2048);
  });

  it('fsw sweeps the cutoff in 8-bit units with an 8-bit wrap', () => {
    const src = `${CLOCK}
voice 0 "c4" wave:pulse fx:"fcut.254 fsw.1"
`;
    const frames = registerFrames(src, 7);
    expect(cutoff(frames[3])).toBe(254 * 8);
    expect([cutoff(frames[4]), cutoff(frames[5]), cutoff(frames[6])]).toEqual([255 * 8, 0, 8]);
  });
});

describe('ADR 0005 D4: hardRestart clauses', () => {
  const song = (hr, notes) => `${CLOCK}
instrument lead
  wave pulse
  adsr 1 2 3 4
  ${hr}
end
voice 0 lead "${notes}"
`;

  it('writes the restart envelope and test bit in the window before the next note', () => {
    const frames = registerFrames(song('hardRestart 2 ad 0 0 sr 15 0 test', 'c4 e4'), 6);
    const ad = (f) => reg(f, VOICE_REGISTER_OFFSETS.attackDecay);
    const sr = (f) => reg(f, VOICE_REGISTER_OFFSETS.sustainRelease);
    const ctl = (f) => reg(f, VOICE_REGISTER_OFFSETS.control);
    expect([ad(frames[1]), sr(frames[1])]).toEqual([0x12, 0x34]);
    expect([ad(frames[2]), sr(frames[2]), ctl(frames[2]) & 0x09]).toEqual([0x00, 0xf0, 0x08]);
    expect([ad(frames[4]), sr(frames[4]), ctl(frames[4]) & 0x09]).toEqual([0x12, 0x34, 0x01]);
  });

  it('leaves the release alone before a rest', () => {
    const frames = registerFrames(song('hardRestart 2 ad 0 0', 'c4 ~'), 4);
    expect(reg(frames[2], VOICE_REGISTER_OFFSETS.attackDecay)).toBe(0x12);
  });

  it('a plain hardRestart N compiles exactly as gate -N', () => {
    const a = registerFrames(song('hardRestart 2', 'c4 e4 g4'), 12);
    const b = registerFrames(song('gate -2', 'c4 e4 g4'), 12);
    expect(a).toEqual(b);
  });
});

describe('ADR 0005 D5: groove', () => {
  it('alternates beat lengths from the song start, across bar lines', () => {
    const src = `speed 6 groove 5 beats 3
voice 0 "c4 d4 e4" wave:pulse
`;
    const frames = registerFrames(src, 40);
    const onsets = [];
    for (let f = 1; f < frames.length; f += 1) if (freq(frames[f]) !== freq(frames[f - 1])) onsets.push(f);
    // Beats: 6 5 6 | 5 6 5 | 6 5 ... -> bars start at 0, 17, 33, so the
    // pitch changes at 6, 11 | 17, 22, 28 | 33, 39.
    expect(onsets).toEqual([6, 11, 17, 22, 28, 33, 39]);
  });

  it('is rejected with tempo (BPM) and with a speed pattern', () => {
    expect(() => compileDslV2Source('tempo 120 groove 5\nvoice 0 "c4"\n')).toThrow(/needs 'speed N'/);
    expect(() => compileDslV2Source('speed "<6 5>" groove 5\nvoice 0 "c4"\n')).toThrow(/constant 'speed N'/);
  });
});
