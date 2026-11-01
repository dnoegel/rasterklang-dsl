// Flats spelled with 'b' (eb3, ab1, bb2) are notes, the same as the 'f'
// spelling (ef3). Before DSL 4.5.0 they stayed bare words: they compiled
// without a diagnostic and played no pitch of their own. Also: `insts:` and
// `fx:` in short mode reject an unquoted value instead of ignoring it.
import { describe, expect, it } from 'vitest';
import { compileDslV2Source } from '../../src/dsl/index.js';
import { compile } from 'rasterklang-score/compiler/compile.js';
import { VOICE_REGISTER_OFFSETS, noteToFreqRegister } from 'rasterklang-score/compiler/registerMap.js';

function freqAt(src, frame) {
  const frameSource = compile(compileDslV2Source(src));
  const regs = new Array(32).fill(0);
  for (let f = 0; f <= frame; f += 1) for (const { addr, value } of frameSource(f)) regs[addr] = value;
  return regs[VOICE_REGISTER_OFFSETS.freqLo] | (regs[VOICE_REGISTER_OFFSETS.freqHi] << 8);
}

const song = (notes) => `speed 4 beats 2\nvoice 0 "${notes}" wave:pulse adsr:0,0,15,0\n`;

describe('flat spellings', () => {
  it.each([
    ['eb3', 51],
    ['ef3', 51],
    ['ab1', 32],
    ['bb2', 46],
    ['db5', 73],
  ])('%s plays MIDI note %i', (name, midi) => {
    expect(freqAt(song(`${name} ~`), 1)).toBe(noteToFreqRegister(midi));
  });

  it('a note-like word with two accidentals is an error, not a silent string', () => {
    expect(() => compileDslV2Source(song('ebb3 ~'))).toThrow(/not a note name/);
  });
});

describe('unquoted insts:/fx: in short mode', () => {
  const head = 'speed 4 beats 2\ninstrument bd\n  wave noise\nend\npattern kit = "bd ~"\n';
  it('insts:$kit is an error naming the quoted form', () => {
    expect(() => compileDslV2Source(`${head}voice 0 bd $kit insts:$kit\n`)).toThrow(/'insts:' expects a quoted .*insts:"\$kit"/);
  });
  it('insts:"$kit" works', () => {
    expect(() => compileDslV2Source(`${head}voice 0 bd $kit insts:"$kit"\n`)).not.toThrow();
  });
  it('fx:tp is an error too', () => {
    expect(() => compileDslV2Source(`${head}voice 0 bd $kit fx:tp\n`)).toThrow(/'fx:' expects a quoted/);
  });
});

describe('note-name typo check knows the b flat', () => {
  it('hb4 is reported as an invalid note letter', () => {
    expect(() => compileDslV2Source(song('hb4 ~'))).toThrow(/invalid note name "hb4"/);
  });
});
