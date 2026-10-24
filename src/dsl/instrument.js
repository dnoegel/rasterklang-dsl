// The instrument({...}) runtime representation DSL v2's `instrument NAME ...
// end` blocks (and short-mode inline attrs, which are sugar for an anonymous
// one-off instrument - see docs/dsl-grammar-spec.md's "Voices" section)
// compile down to. A plain, engine-agnostic data object: interpret.js builds
// one per `instrument` block/inline-attrs set and attaches it to a
// VoiceChain via song.js's `.inst()` method; rasterklang-score's compile.js
// reads its fields directly (duck-typed). compile.js never imports from this
// package, so the dependency only points one way and compile.js stays usable
// by songs written directly in JS, with no knowledge of the DSL.
//
// Every table field (wave/arp/pw/filterTable/test) is sampled by
// framesSinceOnset - frames since THIS voice's current note's own onset,
// tracker semantics - not by cycle position, the clock a voice's `notes`
// pattern itself is sampled by. See the grammar spec's "critical semantic
// distinction" callout in the "Instruments" section: get this backwards and
// arpeggios/PWM sweeps phase-drift against the song instead of restarting
// fresh on every note. compile.js's sampleTable() helper owns the actual
// per-frame sampling/loop-wrap logic; this module only normalizes whatever
// the parser handed it into the flat shape compile.js expects.

// LOOP_TABLE_NAMES - the table kinds a `loop` field can independently target
// (see the "Per-table loop points" section below). Shared by interpret.js
// (which builds the `loop` map this file normalizes) and compile.js (which
// resolves a specific table's loop index via resolveLoopIndex() there).
export const LOOP_TABLE_NAMES = ['wave', 'pw', 'arp', 'filterTable', 'test', 'gate'];

// instrument({ wave, adsr, pw, arp, filterTable, loop, vib, vibTri, slide, test, gate,
// gateTable }) -> a plain runtime object. Every field defaults to null/0 when
// omitted - e.g. a percussion instrument that only sets `wave`/`adsr`/`test`
// leaves pw/arp/filterTable/vib/slide/gate all null, and compile.js's
// instrument-table-wins-else-chain-fallback precedence (see its own
// comments) means a null field is simply invisible, never a zero/silent
// override.
// normalizeVibTri: { cmp, speed | speedShift, delay } with sane defaults, or
// null when the field was never set. `speedShift` is the note-dependent
// "one semitone >> N" form (see the DSL's `sN` parameter spelling); exactly
// one of speed/speedShift is ever present.
function normalizeVibTri(v) {
  if (!v || typeof v.cmp !== 'number') return null;
  const delay = typeof v.delay === 'number' ? Math.max(0, Math.round(v.delay)) : 0;
  const cmp = Math.max(0, Math.round(v.cmp)) & 0x7f;
  if (typeof v.speedShift === 'number') return { cmp, speedShift: Math.max(0, Math.round(v.speedShift)), delay };
  if (typeof v.speed !== 'number') return null;
  return { cmp, speed: Math.round(v.speed), delay };
}

export function instrument({
  wave = null,
  adsr = null,
  pw = null,
  arp = null,
  filterTable = null,
  loop = 0,
  vib = null,
  vibTri = null,
  slide = null,
  test = null,
  gate = null,
  gateTable = null,
  speed = null,
  gateMode = null,
} = {}) {
  // normalizeTable: an empty array is treated the same as "field never set"
  // (null) - so every call site downstream (compile.js's `inst.wave` etc.)
  // can use a single truthiness check instead of also checking `.length > 0`
  // everywhere a table might be sampled.
  const normalizeTable = (t) => (Array.isArray(t) && t.length > 0 ? t : null);

  // normalizeLoop: GoatTracker/SID-Wizard give wave/pulse/filter tables each
  // their OWN independent sustain-loop point, which a single global `loop N`
  // can't describe (see docs/dsl-grammar-spec.md's "Per-table loop points"
  // section). A plain integer means that index for every table this
  // instrument has. A `{ TABLE: N, ... }` map (built by interpret.js from
  // `loop wave 1 pw 3 ...`) gives each named table its own loop index; a
  // table not mentioned falls back to 0 (plain wraparound).
  const normalizeLoop = (l) => {
    if (l && typeof l === 'object' && !Array.isArray(l)) {
      const map = {};
      for (const name of LOOP_TABLE_NAMES) {
        if (Number.isInteger(l[name])) map[name] = l[name];
      }
      return map;
    }
    return Number.isInteger(l) ? l : 0;
  };

  // normalizeSpeedDelay (ADR 0003 D10): { TABLE: { speed, delay } }, one
  // entry per table this instrument gave an explicit speed/delay to (built
  // by interpret.js from `arp 0 4 7 speed 2 delay 3`). A table not
  // mentioned is absent from the map entirely - compile.js's
  // resolveSpeedDelay() defaults an unmentioned table to (1, 0).
  const normalizeSpeedDelay = (s) => {
    if (!s || typeof s !== 'object') return null;
    const map = {};
    for (const name of LOOP_TABLE_NAMES) {
      const entry = s[name];
      if (!entry) continue;
      const sp = Number.isInteger(entry.speed) && entry.speed > 0 ? entry.speed : 1;
      const dl = Number.isInteger(entry.delay) && entry.delay >= 0 ? entry.delay : 0;
      map[name] = { speed: sp, delay: dl };
    }
    return Object.keys(map).length > 0 ? map : null;
  };

  return {
    wave: normalizeTable(wave),
    // adsr is 4 fixed positional scalars (attack/decay/sustain/release), not
    // a per-frame table - see the grammar spec's "Fixed positional scalars"
    // callout. Anything other than exactly 4 values is treated as "not set"
    // rather than silently truncating/padding, since a malformed adsr here
    // would otherwise silently fall back to VoiceChain's own defaults with
    // no signal anything was wrong.
    adsr: Array.isArray(adsr) && adsr.length === 4 ? adsr : null,
    pw: normalizeTable(pw),
    arp: normalizeTable(arp),
    filterTable: normalizeTable(filterTable),
    // The table index (0-based) every table wraps back to once past its
    // last entry - a plain integer (applied to every table) or a per-table map
    // (see normalizeLoop above). See compile.js's resolveLoopIndex()/
    // sampleTable() for the actual formula.
    loop: normalizeLoop(loop),
    // { TABLE: { speed, delay } } or null (ADR 0003 D10) - see
    // normalizeSpeedDelay above / compile.js's resolveSpeedDelay().
    speed: normalizeSpeedDelay(speed),
    vib: vib && typeof vib.depth === 'number' && typeof vib.rate === 'number' ? { depth: vib.depth, rate: vib.rate } : null,
    // vibTri - GoatTracker's own vibrato, kept as a SEPARATE field from
    // `vib` rather than a mode flag on it because the two are different
    // shapes in different units: `vib` is a sine of N semitones at N Hz,
    // `vibTri` is gplay.c's integrating triangle counter in raw frequency-
    // register units, with a delay before it starts. Neither can express
    // the other, and an instrument that sets both would be asking for two
    // pitch modulators at once - compile.js gives `vibTri` precedence,
    // since it is the exact one.
    vibTri: normalizeVibTri(vibTri),
    slide: slide && typeof slide.time === 'number' ? { time: slide.time } : null,
    test: normalizeTable(test),
    // A default gate fraction (see song.js's VoiceChain.gate()) - null means
    // "this instrument has no gate default," letting interpret.js tell that
    // apart from "explicitly defaults to holding the whole slot" (which
    // would be the numeric value 1, not null).
    gate: typeof gate === 'number' ? gate : null,
    // gateTable - a per-frame 0/1 gate-bit PROGRAM (frames-since-onset
    // sampled, loopable like wave/arp/pw/test), distinct from `gate` above
    // (a single static fraction-of-slot cutoff). GoatTracker/SID-Wizard
    // songs routinely toggle the gate bit OFF then back ON *within* one
    // note's sustain (a wave-table row with a gate-off marker) to retrigger
    // the ADSR envelope for a percussive pluck without a new note onset -
    // `gate`'s single fraction can only ever turn the gate off once per
    // note, never back on, so it can't express this. When present, this
    // table wins over `gate`/the chain's own .gate() fraction entirely (see
    // compile.js) - same "instrument table wins" precedent as wave/pw/test.
    gateTable: normalizeTable(gateTable),
    // ADR 0003 D11: 'legato' suppresses the forced onset gate-off/gate-on
    // pair between consecutive notes (framesSinceOnset keeps counting, the
    // same suppression toneportaLegato already does for the fx `tp` op -
    // see compile.js); 'retrigger' guarantees a real, frame-observable
    // gate low-to-high edge on every onset, including a repeated pitch, by
    // defaulting gateReleaseFrames to 1 when nothing more specific (a
    // numeric `gate`/`hardRestart` override) already set one. null (the
    // default) is a forced same-frame edge with no real low-to-high
    // guarantee.
    gateMode: gateMode === 'legato' || gateMode === 'retrigger' ? gateMode : null,
  };
}
