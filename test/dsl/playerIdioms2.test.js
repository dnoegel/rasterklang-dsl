// Suite for ADR 0006 (docs/adr/0006-tracker-conversion-backlog.md): ties
// over bar lines, the three engine fixes (per-note instrument envelopes,
// ties under a hard restart, arp tables through the release), the note
// table's clock and tuning, and `hardRestart N beforeNotes`.
//
// Assertions are on the compiled register state per frame, as in
// playerIdioms.test.js.
import { describe, expect, it } from 'vitest';
import { compileDslV2Source } from '../../src/dsl/index.js';
import { compile } from 'rasterklang-score/compiler/compile.js';
import { VOICE_REGISTER_OFFSETS, noteToFreqRegister } from 'rasterklang-score/compiler/registerMap.js';

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

const freq = (f) => f[VOICE_REGISTER_OFFSETS.freqLo] | (f[VOICE_REGISTER_OFFSETS.freqHi] << 8);
const gate = (f) => f[VOICE_REGISTER_OFFSETS.control] & 1;
const ad = (f) => f[VOICE_REGISTER_OFFSETS.attackDecay];
const sr = (f) => f[VOICE_REGISTER_OFFSETS.sustainRelease];
// Every step is 4 frames: `beats` equals the step count of the patterns.
const clock = (steps) => `speed 4 beats ${steps}`;

describe('ADR 0006 D1: a leading hold ties over the bar line', () => {
  const song = (gateField) => `${clock(4)}
instrument lead
  wave pulse
  ${gateField}
end
pattern a 4 = "c4 _ _ e4"
pattern b 4 = "_ _ g4 _"
voice 0 lead "<$a $b>"
`;

  it('the note keeps sounding into the next bar, without a retrigger', () => {
    const frames = registerFrames(song('adsr 0 9 9 0'), 32);
    // e4 starts at frame 12 and is tied through frames 16-23 of bar 2.
    expect(frames.slice(12, 24).every((f) => freq(f) === noteToFreqRegister(64))).toBe(true);
    expect(frames.slice(12, 24).every((f) => gate(f) === 1)).toBe(true);
    expect(freq(frames[24])).toBe(noteToFreqRegister(67));
  });

  it('a hard restart does not cut the tied note at the bar line', () => {
    const frames = registerFrames(song('hardRestart 1'), 32);
    expect(gate(frames[15])).toBe(1);
    // The tied part releases one frame before g4, like any note.
    expect(gate(frames[23])).toBe(0);
    expect(gate(frames[24])).toBe(1);
  });

  it('is silence when nothing sounds before it', () => {
    const frames = registerFrames(`${clock(2)}\nvoice 0 "_ c4" wave:pulse\n`, 8);
    expect(frames.slice(0, 4).every((f) => gate(f) === 0)).toBe(true);
    expect(gate(frames[4])).toBe(1);
  });

  it('is an error at the start of an fx or insts sequence', () => {
    expect(() => compile(compileDslV2Source(`${clock(2)}\nvoice 0 "c4 e4" wave:pulse fx:"_ tp.0"\n`))(0)).toThrow(/continues a NOTE/);
    const insts = `${clock(2)}\ninstrument a\n  wave pulse\nend\nvoice 0\n  notes "c4 e4"\n  insts "_ a"\nend\n`;
    expect(() => compile(compileDslV2Source(insts))(0)).toThrow(/continues a NOTE/);
  });
});

describe('ADR 0006 D2: engine fixes', () => {
  it('insts switches the envelope with the instrument', () => {
    const src = `${clock(2)}
instrument a
  wave pulse
  adsr 1 2 3 4
end
instrument b
  wave saw
  adsr 9 8 7 6
end
voice 0
  notes "c4 e4"
  insts "a b"
end
`;
    const frames = registerFrames(src, 8);
    expect([ad(frames[1]), sr(frames[1])]).toEqual([0x12, 0x34]);
    expect([ad(frames[5]), sr(frames[5])]).toEqual([0x98, 0x76]);
  });

  it('a tie (fx tp.0) under gate -N keeps the gate open and releases at the tied note\'s end', () => {
    const src = `${clock(2)}
instrument a
  wave pulse
  gate -1
end
voice 0 a
  notes "c4 e4"
  fx "~ tp.0"
end
`;
    const frames = registerFrames(src, 8);
    expect(frames.map(gate)).toEqual([1, 1, 1, 1, 1, 1, 1, 0]);
  });

  it('a slideTo transition under gate -N does not retrigger', () => {
    const src = `${clock(2)}
instrument a
  wave pulse
  gate -1
end
voice 0 a "c4 slideTo:e4"
`;
    const frames = registerFrames(src, 8);
    expect(frames.map(gate)).toEqual([1, 1, 1, 1, 1, 1, 1, 0]);
  });

  it('an arp table keeps running from the last note during a rest', () => {
    const src = `${clock(2)}
instrument a
  wave pulse
  arp 0 12
end
voice 0 a "c4 ~"
`;
    const frames = registerFrames(src, 8);
    const c4 = noteToFreqRegister(60);
    const c5 = noteToFreqRegister(72);
    expect(frames.map(freq)).toEqual([c4, c5, c4, c5, c4, c5, c4, c5]);
    expect(frames.slice(4).every((f) => gate(f) === 0)).toBe(true);
  });
});

describe('ADR 0006 D4: note table clock and tuning', () => {
  const word = (midi, clockHz, cents = 0) => Math.round((440 * 2 ** ((midi - 69) / 12 + cents / 1200) * 16777216) / clockHz) & 0xffff;

  it('pitch NTSC computes the note table at the NTSC clock', () => {
    const frames = registerFrames(`target { video NTSC pitch NTSC }\n${clock(1)}\nvoice 0 "a4" wave:pulse\n`, 2);
    expect(freq(frames[0])).toBe(word(69, 1022727));
    expect(freq(frames[0])).not.toBe(noteToFreqRegister(69));
  });

  it('tuning shifts every note by cents; without either nothing changes', () => {
    const tuned = registerFrames(`target { tuning 40 }\n${clock(1)}\nvoice 0 "a4" wave:pulse\n`, 1);
    expect(freq(tuned[0])).toBe(word(69, 985248, 40));
    const plain = registerFrames(`target { video NTSC }\n${clock(1)}\nvoice 0 "a4" wave:pulse\n`, 1);
    expect(freq(plain[0])).toBe(noteToFreqRegister(69));
  });
});

describe('ADR 0006 D5: hardRestart N beforeNotes', () => {
  it('releases early before a note, but at the rest itself before a key-off', () => {
    const src = `${clock(4)}
instrument a
  wave pulse
  hardRestart 1 beforeNotes
end
voice 0 a "c4 e4 ~ g4"
`;
    const frames = registerFrames(src, 16);
    // g4 releases early too: the bar loops, so c4 follows it.
    expect(frames.map(gate)).toEqual([1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0]);
  });

  it('without beforeNotes the early release also comes before the rest', () => {
    const src = `${clock(4)}
instrument a
  wave pulse
  hardRestart 1
end
voice 0 a "c4 e4 ~ g4"
`;
    const frames = registerFrames(src, 8);
    expect(gate(frames[7])).toBe(0);
  });
});
