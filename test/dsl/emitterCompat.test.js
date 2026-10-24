// Compatibility suite for two shapes that machine-generated DSL (a SID-to-DSL
// converter's output) relies on: a single-token `gate 1!N` table and `$pN`
// pattern references that look like note names. Both are legal DSL v2 by
// the grammar's own rules, and both were once broken by stricter checks
// (ADR 0003's gate modes and the D14 note-name check) without any in-repo
// test noticing, because nothing in this repo emits either shape.
//
// The cases are written in the generator's spelling, not a minimal
// reduction, so that the next person to touch either rule sees what
// actually consumes this grammar.
import { describe, expect, it } from 'vitest';
import { compileDslV2Source } from '../../src/dsl/index.js';

const CLOCK = 'tempo 120 beats 4\n';

describe('gate: a single `!N` repeat token is a one-entry table, not a bad scalar', () => {
  // The emitter writes a held gate column as `gate 1!128` whenever an
  // instrument's whole gate table is one repeated value, which is the common
  // case for a sustained instrument. The single-token form must not be
  // mistaken for a malformed scalar by the scalar/mode fast path.
  it('accepts `gate 1!128` on its own', () => {
    const src = `${CLOCK}instrument lead
  wave pulse
  gate 1!128
end
voice 0 lead "c4 g4"
`;
    expect(() => compileDslV2Source(src)).not.toThrow();
  });

  it('accepts a single repeat token beside other single-token forms', () => {
    for (const field of ['gate 1!2', 'gate 0!16', 'gate 1!128']) {
      const src = `${CLOCK}instrument lead\n  wave pulse\n  ${field}\nend\nvoice 0 lead "c4"\n`;
      expect(() => compileDslV2Source(src), field).not.toThrow();
    }
  });

  it('still parses the multi-token table form the fast path never touched', () => {
    const src = `${CLOCK}instrument lead
  wave pulse
  gate 1!128 0!4
end
voice 0 lead "c4"
`;
    expect(() => compileDslV2Source(src)).not.toThrow();
  });

  it('still rejects a bare word that is neither a number nor a gate mode', () => {
    // The point of the fast path. Narrowing it must not blunt it.
    const src = `${CLOCK}instrument lead\n  wave pulse\n  gate bogus\nend\nvoice 0 lead "c4"\n`;
    expect(() => compileDslV2Source(src)).toThrow(/'gate' expects a number, or one of legato\/retrigger/);
  });

  it('still accepts the gate modes themselves', () => {
    for (const mode of ['legato', 'retrigger']) {
      const src = `${CLOCK}instrument lead\n  wave pulse\n  gate ${mode}\nend\nvoice 0 lead "c4"\n`;
      expect(() => compileDslV2Source(src), mode).not.toThrow();
    }
  });
});

describe('note-name check: a `$name` reference is a pattern name, never a note', () => {
  // ADR 0003 D14's checkNoteNames looks for "letter, optional accidental,
  // digits" as a note-name typo (German `h4` for `b4` is the flagship case).
  // ADR 0003 D5's `$name` structural references match that shape exactly when
  // the pattern is named `p0`, `i0`, `h1` and so on, which is precisely how
  // the emitter names them.
  it('accepts the emitter arrangement spelling with $pN references', () => {
    const src = `${CLOCK}instrument lead
  wave pulse
end
pattern p0 4 = "c4 e4 g4 c5"
pattern p1 4 = "g4 e4 c4 ~"
pattern p2 4 = "~ ~ ~ ~"
voice 0 lead "<[$p0] [$p1] [$p1] [$p2]>"
`;
    expect(() => compileDslV2Source(src)).not.toThrow();
  });

  it('accepts reference names whose letter is not a note letter', () => {
    // `i0`, `p0`, `x1` and `h1` all match the typo shape; `h1` is the one the
    // check exists for, which is exactly why the sigil has to win here.
    for (const name of ['p0', 'i0', 'x1', 'h1']) {
      const src = `${CLOCK}instrument lead\n  wave pulse\nend\npattern ${name} 2 = "c4 e4"\nvoice 0 lead "<[$${name}]>"\n`;
      expect(() => compileDslV2Source(src), name).not.toThrow();
    }
  });

  it('still rejects a real note-name typo written without a sigil', () => {
    // The check's whole purpose. Excluding '$' must not exclude bare words.
    const src = `${CLOCK}instrument lead\n  wave pulse\nend\nvoice 0 lead "c4 h4"\n`;
    expect(() => compileDslV2Source(src)).toThrow(/invalid note name "h4"/);
  });

  it('still suggests the German-notation fix for a bare h', () => {
    const src = `${CLOCK}instrument lead\n  wave pulse\nend\nvoice 0 lead "h4"\n`;
    expect(() => compileDslV2Source(src)).toThrow(/did you mean "b4"/);
  });
});
