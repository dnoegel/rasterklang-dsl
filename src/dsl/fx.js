// DSL v2's per-note EFFECT lane (`fx` voice field, added in 2.15.0): the
// one-shot, per-row tracker commands GoatTracker patterns carry in their
// command/command-data columns.
//
// Why a parallel mini-notation string, not an inline annotation on `notes`:
// `insts` (2.9.0, GoatTracker's per-row INSTRUMENT column) already solves
// this shape of problem (a second tracker column that lines up with the note
// column but is sampled independently) with a parallel mini-notation
// pattern. Reusing that idiom gives `fx` `<...>` alternation, `!n` repeats,
// `@n` weights, `$name` splices and every transform for free, lets an
// importer reuse its orderlist folding, and keeps note text free of an
// effect-specific micro-syntax.
//
// Semantics differ from `insts` deliberately: an `insts` value is sampled
// only at a note ONSET (a row's instrument column belongs to that row's
// note), whereas a GoatTracker command executes on its row whether or not
// that row has a note, so compile.js samples `fx` on its OWN hap
// boundaries. See rasterklang-score's compile.js fx block for the runtime
// side.
//
// The command semantics below follow GoatTracker 2's player, src/gplay.c
// (the CMD_* switch in its TICK0 block and the TICKNEFFECTS per-tick block).

// FX_OPS - every effect token name, with its parameter arity. Also the
// keyword table src/dsl/keywords.js re-exports for editor completion.
export const FX_OPS = {
  // CMD_PORTAUP (0x01) / CMD_PORTADOWN (0x02): add/subtract a fixed
  // frequency-register delta EVERY frame for as long as the command is in
  // effect. gplay.c reads the delta out of the song's speed table; the
  // importer resolves that indirection at conversion time, so the DSL
  // carries the real number and needs no speed-table concept of its own.
  pup: 1,
  pdn: 1,
  // CMD_TONEPORTA (0x03): glide the frequency register toward THIS note's
  // own pitch at a fixed rate per frame, clamping on arrival (speed 0 =
  // jump there instantly). Also suppresses the note's gate retrigger and
  // instrument re-init, exactly as gplay.c's `if (cptr->newcommand !=
  // CMD_TONEPORTA)` guards do - that legato behavior is half the point of
  // the command.
  tp: 1,
  // CMD_VIBRATO (0x04): gplay.c's own triangle LFO - a vibtime counter
  // stepping by 2, inverted whenever it passes the compare value, with the
  // low bit choosing whether this frame adds or subtracts `speed` from the
  // frequency register. Params: compare value, then speed.
  vib: 2,
  // CMD_SETAD (0x05) / CMD_SETSR (0x06): overwrite this voice's attack/
  // decay or sustain/release register for the rest of the note. Params are
  // the two nibbles, 0-15 each.
  ad: 2,
  sr: 2,
  // CMD_SETWAVE (0x07): overwrite this voice's waveform/control bits.
  // Param is the raw control-byte value gplay.c stores in cptr->wave.
  wave: 1,
  // `pw.N` has NO CMD_* counterpart - GoatTracker 2 reaches pulse width
  // only indirectly, via CMD_SETPULSEPTR (0x09) repointing the instrument's
  // pulse TABLE, so unlike every op above this one is not a GoatTracker
  // command but a DSL-native one-shot register write. Decompilers need it:
  // in converted real-world tunes, pulse width is the largest class of
  // register writes that no tracker structure explains, and it was the only
  // register class with no fx spelling at all (frequency has pup/pdn/tp,
  // the control nibble has wave, the envelope has ad/sr, the global lanes
  // have fcut/vol). Param is the 12-bit pulse
  // width, 0-4095 - deliberately the SAME value domain as an instrument's
  // own `pw` TABLE column (see instrument.js / compile.js's `inst.pw`
  // sampling, which clamps to 0..4095 with a 2048 default), so the two
  // spellings of "this voice's pulse width" can never disagree.
  pw: 1,
  // CMD_SETFILTERCUTOFF (0x0c): set the shared filter's cutoff. Param is
  // GoatTracker's 8-bit cutoff, scaled to the SID's 11-bit register the
  // same way the rest of this codegen scales it (value * 8).
  fcut: 1,
  // CMD_SETMASTERVOL (0x0d): set the master volume nibble, 0-15.
  vol: 1,
};

// A parameter is either a plain integer, or `sN` - "shift-derived": gplay.c
// treats a speed-table entry whose high byte is >= 0x80 (or, for vibrato, a
// compare value >= 0x80) as "one semitone's worth of frequency at the
// current note, shifted right by N", which is note-DEPENDENT and so cannot
// be resolved to a constant at conversion time. Carrying the shift through
// to compile.js keeps those rows exact instead of dropping them.
function parseParam(text, token) {
  if (/^s[0-9]+$/.test(text)) return { shift: Number(text.slice(1)) };
  if (!/^-?[0-9]+$/.test(text)) {
    throw new Error(`fx token '${token}': parameter '${text}' must be an integer or a shift-derived 'sN'.`);
  }
  return { value: Number(text) };
}

// parseFxToken(token) -> a plain effect object compile.js consumes:
// { op, a, b } where a/b are each { value } or { shift }. Throws with the
// offending token quoted for anything unknown or mis-arity'd.
export function parseFxToken(token) {
  if (typeof token !== 'string') {
    throw new Error(`fx pattern: expected an effect token, got ${JSON.stringify(token)}.`);
  }
  const parts = token.split('.');
  const op = parts[0];
  const arity = FX_OPS[op];
  if (arity === undefined) {
    throw new Error(`fx token '${token}': unknown effect '${op}' (expected one of ${Object.keys(FX_OPS).join('/')}).`);
  }
  const args = parts.slice(1);
  if (args.length !== arity) {
    throw new Error(`fx token '${token}': effect '${op}' takes exactly ${arity} parameter(s), got ${args.length}.`);
  }
  const out = { op };
  if (arity >= 1) out.a = parseParam(args[0], token);
  if (arity >= 2) out.b = parseParam(args[1], token);
  // `pw` is the one op whose parameter has a hard, non-negotiable domain:
  // the SID's pulse width is 12 bits ($02/$03 per voice), so anything
  // outside 0..4095 is a scaling mistake, not a value. Rejecting it here
  // rather than letting compile.js clamp is deliberate: an importer feeding
  // this lane could mix up an 8-bit tracker pulse value with the 12-bit
  // register one, and a silent clamp to 4095 would turn that mistake into a
  // plausible-sounding tune instead of an error. The other ops clamp at
  // compile time.
  if (op === 'pw') {
    if (out.a.shift !== undefined) {
      throw new Error(`fx token '${token}': 'pw' takes a plain integer pulse width; the shift-derived 'sN' form is frequency-only.`);
    }
    if (out.a.value < 0 || out.a.value > 4095) {
      throw new Error(`fx token '${token}': pulse width ${out.a.value} is out of range - must be 0-4095 (the SID's 12-bit pulse width, same domain as an instrument's own 'pw' table).`);
    }
  }
  return out;
}
