// Semantic-pass tests: the rules ADR 0001 added that are neither
// grammatical nor engine-level (src/dsl/analyze.js, src/dsl/miniLength.js).
//
// These deliberately go through checkSource() rather than
// compileDslV2Source(), so they exercise the whole lexer -> parser ->
// analysis path while staying inside this repo - the semantic layer's
// whole point is being decidable without the pattern engine, and a test
// that needed the engine would quietly undo that.
import { describe, expect, it } from 'vitest';
import { checkSource } from '../../src/dsl/index.js';

function messages(src, severity) {
  return checkSource(src)
    .filter((d) => (severity ? d.severity === severity : true))
    .map((d) => d.message);
}

function errors(src) {
  return messages(src, 'error');
}

const CLOCK = 'tempo 120 beats 4\n';

describe('D1: declared pattern length', () => {
  it('accepts a pattern whose contents match its declared length', () => {
    expect(errors(`${CLOCK}pattern bass 4 = "c2 ~ c2 ~"\nvoice 0 $bass\n`)).toEqual([]);
  });

  it('reads `@n` as n beats in a metric pattern', () => {
    expect(errors(`${CLOCK}pattern lead 8 = "c5@2 e5 g5 c6@4"\nvoice 0 $lead\n`)).toEqual([]);
  });

  it('reports a mismatch with both the declared and the actual length', () => {
    const [msg] = errors(`${CLOCK}pattern lead 32 = "c4@1.25 d4 e4@2"\nvoice 0 $lead\n`);
    expect(msg).toContain('declares 32 beats');
    expect(msg).toContain('add up to 4.25 beats');
  });

  it('allows a decimal declared length', () => {
    expect(errors(`${CLOCK}pattern odd 7.5 = "c4@0.5 d4@3 e4@4"\nvoice 0 $odd\n`)).toEqual([]);
  });

  it('leaves a pattern with no declared length proportional (no check at all)', () => {
    expect(errors(`${CLOCK}pattern hats = "bd*4 sn cp hh oh"\nvoice 0 $hats\n`)).toEqual([]);
  });

  it('checks each orderlist element against the declared length, not their sum', () => {
    expect(errors(`${CLOCK}pattern song 4 = "<[c4 c4 c4 c4] [e4 e4 e4 e4]>"\nvoice 0 $song\n`)).toEqual([]);
    const [msg] = errors(`${CLOCK}pattern song 4 = "<[c4 c4 c4 c4] [e4 e4 e4]>"\nvoice 0 $song\n`);
    expect(msg).toContain('orderlist element 2');
  });
});

describe('D3: each 1/D', () => {
  it('derives a metric length from token count * beatUnit/D, with no declared number', () => {
    const src = `tempo 120 meter 4/4\npattern hats each 1/16 = "c6 d6 e6 f6 g6 a6 b6 c7"\nvoice 0 $hats\n`;
    expect(errors(src)).toEqual([]);
  });

  it('accepts a declared length that matches the derived one', () => {
    const src = `tempo 120 meter 4/4\npattern hats 2 each 1/16 = "c6 d6 e6 f6 g6 a6 b6 c7"\nvoice 0 $hats\n`;
    expect(errors(src)).toEqual([]);
  });

  it('reports a mismatch between a declared length and the derived one', () => {
    const src = `tempo 120 meter 4/4\npattern hats 4 each 1/16 = "c6 d6 e6 f6 g6 a6 b6 c7"\nvoice 0 $hats\n`;
    const [msg] = errors(src);
    expect(msg).toContain('declares 4 beats');
    expect(msg).toContain('add up to 2 beats');
  });

  it('lets an explicit /N token override the each default for that token', () => {
    // c6/8 is 0.5 beats (overriding each's 1/16 default); the other six
    // tokens are 0.25 each - 0.5 + 6*0.25 = 2, not 7*0.25 = 1.75, so this
    // declared length only matches if the override actually took effect.
    const src = `tempo 120 meter 4/4\npattern hats 2 each 1/16 = "c6/8 d6 e6 f6 g6 a6 b6"\nvoice 0 $hats\n`;
    expect(errors(src)).toEqual([]);
  });
});

describe('D4/decision 11: exactly one clock statement', () => {
  it('accepts `speed N` on its own', () => {
    expect(errors('speed 6 beats 8\nvoice 0 "c4"\n')).toEqual([]);
  });

  it('rejects having both, naming that case', () => {
    const [msg] = errors(`${CLOCK}speed 6 beats 4\nvoice 0 "c4"\n`);
    expect(msg).toContain('mutually exclusive');
    expect(msg).toContain('both');
  });

  it('rejects having neither, naming that case', () => {
    const [msg] = errors('voice 0 "c4"\n');
    expect(msg).toContain('none');
  });
});

describe('D5: polymetre across voices', () => {
  it('is not an error, and reports the arrangement period', () => {
    const src = `${CLOCK}pattern a 4 = "c2 ~ c2 ~"\npattern b 3 = "e3 f3 g3"\nvoice 0 $a\nvoice 1 $b\n`;
    expect(errors(src)).toEqual([]);
    expect(messages(src, 'info')[0]).toContain('Full period: 12 beats (3 bars)');
  });

  it('hints (does not fail) when the period is implausibly long', () => {
    const src = `${CLOCK}pattern a 4 = "c2 ~ c2 ~"\npattern b 65 = "e3@65"\nvoice 0 $a\nvoice 1 $b\n`;
    expect(errors(src)).toEqual([]);
    expect(messages(src, 'warning').join(' ')).toContain('only realign after');
  });

  // ADR 0003 D4/D5: a quoted `"<$a $b>"` is an ordinary proportional voice,
  // not an error. Sequencing different-length patterns in beat time is what
  // the arrangement syntax (no quotes) is for: see the D4 tests below.
  it('does not error on a quoted orderlist of differently-declared-length refs', () => {
    const src = `${CLOCK}pattern a 4 = "c2 ~ c2 ~"\npattern b 8 = "e3@8"\nvoice 0 "<$a $b>"\n`;
    expect(errors(src)).toEqual([]);
  });
});

// ADR 0003 D4: arrangements - $a $b sequences different-length patterns in
// beat time (the ADR's F3).
describe('ADR 0003 D4: arrangements', () => {
  it('is not an error to sequence patterns of different declared lengths', () => {
    const src = `${CLOCK}pattern a 4 = "c4 d4 e4 f4"\npattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"\nvoice 0 $a $b\n`;
    expect(errors(src)).toEqual([]);
  });

  it('reports the summed length as the arrangement period', () => {
    const src = `${CLOCK}pattern a 4 = "c4 d4 e4 f4"\npattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"\nvoice 0 $a $b\n`;
    expect(messages(src, 'info')[0]).toContain('Full period: 12 beats (3 bars)');
  });

  it('$a*2 $b and $a $a $b both report the same length', () => {
    const a = `${CLOCK}pattern a 4 = "c4 d4 e4 f4"\npattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"\nvoice 0 $a*2 $b\n`;
    const b = `${CLOCK}pattern a 4 = "c4 d4 e4 f4"\npattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"\nvoice 0 $a $a $b\n`;
    expect(errors(a)).toEqual([]);
    expect(errors(b)).toEqual([]);
    expect(messages(a, 'info')[0]).toEqual(messages(b, 'info')[0]);
  });

  it('[$a $b]*2 and $a $b $a $b both report the same length', () => {
    const a = `${CLOCK}pattern a 4 = "c4 d4 e4 f4"\npattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"\nvoice 0 [$a $b]*2\n`;
    const b = `${CLOCK}pattern a 4 = "c4 d4 e4 f4"\npattern b 8 = "c5 d5 e5 f5 g5 a5 b5 c6"\nvoice 0 $a $b $a $b\n`;
    expect(errors(a)).toEqual([]);
    expect(errors(b)).toEqual([]);
    expect(messages(a, 'info')[0]).toEqual(messages(b, 'info')[0]);
  });

  it('names the cycle in a self-referencing (recursive) arrangement pattern', () => {
    const src = `${CLOCK}pattern verse = $verse $verse\nvoice 0 $verse\n`;
    const [msg] = errors(src);
    expect(msg).toContain('recursive');
  });

  it('a bare $name reference (the pre-existing shape) still reports its declared length', () => {
    const src = `${CLOCK}pattern lead 8 = "c5 d5 e5 f5 g5 a5 b5 c6"\nvoice 0 $lead\n`;
    expect(errors(src)).toEqual([]);
    expect(messages(src, 'info')[0]).toContain('8 beats');
  });
});

describe('D6: the hold token', () => {
  it('counts a hold as a beat of the note it extends', () => {
    expect(errors(`${CLOCK}pattern lead 4 = "c5 _ ~ _"\nvoice 0 $lead\n`)).toEqual([]);
  });

  it('rejects a hold with nothing before it', () => {
    const [msg] = errors(`${CLOCK}voice 0 "_ c4"\n`);
    expect(msg).toContain("hold token '_'");
  });
});

describe('known defects from ADR 0001 section 6', () => {
  // ADR 0001 known defect 2 cannot occur since ADR 0003 D5: $name is a
  // structural reference (not text substitution), so `$a!3` replicates the
  // WHOLE referenced pattern, bracketed or not.
  it('2: no longer warns about $name!n - it replicates the whole pattern, not the last token', () => {
    const bare = messages(`${CLOCK}pattern a = "c4 e4"\nvoice 0 "<$a!3 c5>"\n`, 'warning');
    const bracketed = messages(`${CLOCK}pattern a = "c4 e4"\nvoice 0 "<[$a]!3 c5>"\n`, 'warning');
    expect(bare.join(' ')).not.toContain('repeats only');
    expect(bracketed.join(' ')).not.toContain('repeats only');
  });

  it('3: explains where a pipeline transform belongs', () => {
    const [msg] = errors(`${CLOCK}pattern a = "c4 e4" |> slow 4\n`);
    expect(msg).toContain('not allowed on pattern declarations');
  });

  it('4: warns about a cutoff above 2047 without changing the clipping', () => {
    const warnings = messages(`${CLOCK}voice 0 "c4"\nfilter\n  cutoff sine 900..2600\nend\n`, 'warning');
    expect(warnings.join(' ')).toContain('clipped to 2047');
  });
});

describe('decision 10: the migration checker', () => {
  it('flags a `beats` value that looks like a melody-stretch lever', () => {
    expect(messages('tempo 132 beats 28\nvoice 0 "c4"\n', 'warning').join(' ')).toContain('unusual for a time signature');
  });

  it('says nothing about ordinary time signatures', () => {
    for (const beats of [1, 2, 3, 4, 6, 8, 9, 12, 16]) {
      expect(messages(`tempo 120 beats ${beats}\nvoice 0 "c4"\n`, 'warning')).toEqual([]);
    }
  });
});

// ADR 0003 bar lines (D6) and the proportional-fit info (D8). D17's grid
// report/misfit warnings need frame arithmetic analyze.js deliberately
// doesn't have (it must stay usable without the pattern engine); those are
// computed in interpret.js instead.
describe('ADR 0003 D6: checked bar lines', () => {
  it('accepts a `|` that lands exactly on a bar boundary', () => {
    const src = `${CLOCK}pattern lead = "c4/4 d4/4 e4/4 f4/4 | g4/4 a4/4 b4/2"\nvoice 0 $lead\n`;
    expect(errors(src)).toEqual([]);
  });

  it('rejects a `|` off the grid, naming the beat offset it reached', () => {
    const src = `${CLOCK}pattern lead = "c4/4 d4/4 | e4/4 f4/4"\nvoice 0 $lead\n`;
    const [msg] = errors(src);
    expect(msg).toContain('off the grid');
    expect(msg).toContain('beat 2');
  });

  it('checks a bar line in a voice\'s own inline pattern too, not just a named declaration', () => {
    const src = `${CLOCK}voice 0 "c4/4 d4/4 | e4/4 f4/4"\n`;
    const [msg] = errors(src);
    expect(msg).toContain('off the grid');
  });
});

describe('ADR 0003 D8: proportional auto-fit is reported, not silent', () => {
  it('reports the derived grid for a pattern with no /N, no each and no declared number', () => {
    const src = `${CLOCK}pattern lead = "c4 d4 e4 f4 g4 a4 b4 c5"\nvoice 0 $lead\n`;
    const [info] = messages(src, 'info').filter((m) => m.includes('proportional'));
    expect(info).toContain('8 tokens over 1 bar = 1/8 note each');
  });

  it('says nothing about a metric pattern (declared length, each, or /N present)', () => {
    const declared = messages(`${CLOCK}pattern a 4 = "c4 d4 e4 f4"\nvoice 0 $a\n`, 'info').join(' ');
    const each = messages(`${CLOCK}pattern a each 1/8 = "c4 d4 e4 f4"\nvoice 0 $a\n`, 'info').join(' ');
    const duration = messages(`${CLOCK}pattern a = "c4/4 d4/4"\nvoice 0 $a\n`, 'info').join(' ');
    expect(declared).not.toContain('proportional');
    expect(each).not.toContain('proportional');
    expect(duration).not.toContain('proportional');
  });
});
