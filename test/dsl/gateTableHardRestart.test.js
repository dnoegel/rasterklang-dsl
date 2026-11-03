// DSL 4.6.0: a gate table plus `hardRestart N` (or `gate -N`). The table runs
// as before; in the note's last N frames before a following note the gate
// closes and the hard-restart clauses are written, exactly as `hardRestart N`
// does on an instrument without a table. Assertions are on the register state
// per frame, as in playerIdioms2.test.js.
import { describe, expect, it } from 'vitest';
import { compileDslV2Source } from '../../src/dsl/index.js';
import { compile } from 'rasterklang-score/compiler/compile.js';
import { VOICE_REGISTER_OFFSETS } from 'rasterklang-score/compiler/registerMap.js';

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

const gate = (f) => f[VOICE_REGISTER_OFFSETS.control] & 1;
const test = (f) => (f[VOICE_REGISTER_OFFSETS.control] >> 3) & 1;
const ad = (f) => f[VOICE_REGISTER_OFFSETS.attackDecay];
const sr = (f) => f[VOICE_REGISTER_OFFSETS.sustainRelease];
// 4 frames per step, 4 steps per bar
const CLOCK = 'speed 4 beats 4';
const song = (fields, notes) => `${CLOCK}
instrument a
  wave pulse
  adsr 1 2 3 4
${fields.map((f) => `  ${f}`).join('\n')}
end
voice 0 a "${notes}"
`;

describe('gate table plus hardRestart', () => {
  it('a gate table alone still runs unchanged', () => {
    const frames = registerFrames(song(['gate 1 0 1 1'], 'c4 e4 g4 c5'), 8);
    expect(frames.map(gate)).toEqual([1, 0, 1, 1, 1, 0, 1, 1]);
  });

  it('closes the gate in the last N frames before the next note and writes the clauses', () => {
    const frames = registerFrames(song(['gate 1 0 1 1', 'hardRestart 1 ad 0 0 sr 0 0 test'], 'c4 e4 g4 c5'), 8);
    expect(frames.map(gate)).toEqual([1, 0, 1, 0, 1, 0, 1, 0]);
    // the table's own frames keep the instrument's envelope; the window writes the clauses
    expect([ad(frames[2]), sr(frames[2]), test(frames[2])]).toEqual([0x12, 0x34, 0]);
    expect([ad(frames[3]), sr(frames[3]), test(frames[3])]).toEqual([0x00, 0x00, 1]);
    expect([ad(frames[4]), sr(frames[4]), test(frames[4])]).toEqual([0x12, 0x34, 0]);
  });

  it('a two-frame window covers both frames', () => {
    const frames = registerFrames(song(['gate 1!4', 'hardRestart 2 ad 0 0 sr 0 0'], 'c4 e4 g4 c5'), 8);
    expect(frames.map(gate)).toEqual([1, 1, 0, 0, 1, 1, 0, 0]);
    expect(frames.map(sr)).toEqual([0x34, 0x34, 0x00, 0x00, 0x34, 0x34, 0x00, 0x00]);
  });

  it('a tie holds the note: no restart until the note after it', () => {
    const frames = registerFrames(song(['gate 1!4', 'hardRestart 1 sr 0 0'], 'c4 _ e4 g4'), 12);
    expect(frames.map(gate)).toEqual([1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0]);
  });

  it('no restart into a legato transition', () => {
    const frames = registerFrames(song(['gate 1!4', 'hardRestart 1 sr 0 0'], 'c4 slideTo:e4 g4 c5'), 8);
    expect(gate(frames[3])).toBe(1);
    expect(sr(frames[3])).toBe(0x34);
  });

  it('before a rest the gate closes but the clauses are not written', () => {
    const frames = registerFrames(song(['gate 1!4', 'hardRestart 1 sr 0 0'], 'c4 ~ e4 g4'), 4);
    expect(gate(frames[3])).toBe(0);
    expect(sr(frames[3])).toBe(0x34);
  });

  it('with beforeNotes the table runs on into a rest', () => {
    const frames = registerFrames(song(['gate 1!4', 'hardRestart 1 sr 0 0 beforeNotes'], 'c4 ~ e4 g4'), 4);
    expect(frames.map(gate)).toEqual([1, 1, 1, 1]);
  });

  it('the last note of the pattern restarts into the first one of the next bar', () => {
    const frames = registerFrames(song(['gate 1!4', 'hardRestart 1 sr 0 0'], 'c4 e4 g4 c5'), 17);
    expect(gate(frames[15])).toBe(0);
    expect(sr(frames[15])).toBe(0x00);
    expect(gate(frames[16])).toBe(1);
  });

  it('the order of the two fields does not matter, and gate -N is the same spelling', () => {
    const a = registerFrames(song(['gate 1 0 1 1', 'hardRestart 1 ad 0 0'], 'c4 e4 g4'), 16);
    const b = registerFrames(song(['hardRestart 1 ad 0 0', 'gate 1 0 1 1'], 'c4 e4 g4'), 16);
    const c = registerFrames(song(['gate 1 0 1 1', 'gate -1'], 'c4 e4 g4'), 16);
    const d = registerFrames(song(['gate 1 0 1 1', 'hardRestart 1'], 'c4 e4 g4'), 16);
    expect(b).toEqual(a);
    expect(c).toEqual(d);
  });

  it('other gate pairs are still errors', () => {
    expect(() => compileDslV2Source(song(['gate 1 0 1 1', 'gate 0.5'], 'c4'))).toThrow(/only a gate table plus 'hardRestart N' combine/);
    expect(() => compileDslV2Source(song(['gate 1 0', 'gate 1 1 0'], 'c4'))).toThrow(/already has a 'gate'/);
    expect(() => compileDslV2Source(song(['gate legato', 'hardRestart 2'], 'c4'))).toThrow(/already has a 'gate'/);
    expect(() => compileDslV2Source(song(['gate 1 0', 'hardRestart 1', 'hardRestart 2'], 'c4'))).toThrow(/already has a 'gate'/);
  });
});
