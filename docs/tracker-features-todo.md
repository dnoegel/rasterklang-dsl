# Additional features from other trackers

A running list of player mechanisms met while converting tunes from other
C64 trackers into DSL v2 (rasterklang-conv `lift`: DMC, Music Assembler, and
next Future Composer, JCH NewPlayer, Soundmonitor, defMON) that DSL v2 or its
engine (rasterklang-score) cannot say yet, or says differently. It is a
backlog of candidates, not a decision to add them. Converters work around
each entry as noted; the workaround is what a feature here would replace.

Format of an entry:

- **Feature** (trackers): what the player does. *Today:* the converter's
  approximation. *Seen in:* example tunes, and reach when measured.

## Timing and structure

- **Per-note row length** (DMC, Music Assembler, Future Composer,
  Soundmonitor): a duration byte sets how many ticks each following row
  lasts, so rows in one pattern have different lengths. *Today:* a row of
  d ticks becomes d GoatTracker rows (note, then holds), compacted by the
  greatest common divisor of all row lengths. Works, but a pattern of long
  notes spells many `_` holds. *Seen in:* every DMC and Music Assembler
  tune.

- **Notes that sound across a pattern boundary** (every tracker: a note
  rings on while the next pattern starts with empty rows): a DSL pattern
  cannot open with `_` ("a hold with nothing before it is an error"), so
  sngcodegen writes `~` there and the note is cut at the cycle boundary; an
  instrument with `gate -N` then releases a note that starts in a cycle's
  last row almost at once. This hits GoatTracker conversions as much as
  lifted ones. *Today:* `lift -from-notes` cuts its entries where no note
  sounds on and uses 64-row entries; decoded patterns are kept as they are.
  A hold that may continue the previous cycle's note (or a note weight that
  reaches past the cycle) would fix it. *Seen in:* DEMOS/G-L/Hey_Jude.sid
  (from notes: gate match 0.82 with 16-row cycles, 0.93 with 64).
- **Tie: a new pitch without a new gate** (DMC slide and glide, v1d legato,
  Music Assembler tied notes, Future Composer row commands, Soundmonitor
  notes without bit 7, JCH duration bit $10): the converters write a
  GoatTracker tone portamento with speed 0, which becomes `fx tp.0`, and
  the engine starts a new note (new gate) there. *Today:* ties retrigger.
  A note modifier that changes pitch without an onset (`gate legato`
  per note) would say it.

## Instruments

- **Instrument tables run through the release** (DMC, Music Assembler,
  GoatTracker itself): after a key-off the tracker keeps executing the wave
  table, so an arpeggio or wave sequence goes on sounding while the envelope
  releases. rasterklang-score samples an instrument's `arp` only while a
  note is active (`tone !== undefined` in compile.js), so during a `~` rest
  the frequency freezes; `pw` and `wave` keep running. *Today:* when an
  instrument's key-off always comes the same number of frames before the
  next note, the note is held and released with `gate -N` instead; other
  key-offs stay rests and lose the running arpeggio. *Seen in:*
  DEMOS/S-Z/Utopia.sid (DMC, octave arpeggio over the release).
- **Per-note instrument switching does not switch AD/SR**: `insts` changes
  an instrument's tables at the onset, but rasterklang-score reads AD/SR
  from the voice chain's own patterns, never from the per-note instrument
  (see the comment in engine/liveAudition.js). A voice that alternates
  instruments keeps the voice instrument's envelope for every note. This hits
  GoatTracker conversions too. *Today:* nothing; the ADSR lane is wrong
  wherever a voice switches instruments. *Seen in:* every multi-instrument
  voice; Utopia.sid voice 1 plays SR $99 notes with the header
  instrument's $C8.
- **Release before the next note, but not before a key-off** (DMC flag
  $10: the gate drops one tick before the next note starts; a row that keys
  off releases at the key-off): `gate -N` counts back from the end of the
  slot, and a `~` ends the slot, so an instrument with both releases early
  before its key-offs too. *Today:* such an instrument keeps its own
  release as `gate -N` and its key-offs as rests, and releases N frames too
  early before each key-off. A gate-off counted from the next ONSET (not the
  slot end) would say it. *Seen in:* MUSICIANS/M/Marcy/Laughter.sid.
- **Pulse sweeps that run on, bounce or wrap** (DMC: sweeps between two
  nibble limits and reverses; GoatTracker: a modulation step adds to the
  current width every frame): a `pw` table is a list of absolute values, so
  `loop pw N` replays values instead of continuing the sweep from where it
  is. *Today:* the converter measures the first frames of the note's pulse
  width and holds the last value after them. A ramp-with-limits (start,
  speed, low, high, reverse or wrap) would say DMC's sweep exactly.
- **NTSC SID clock** (every NTSC tune): `target { video NTSC }` sets the
  60 Hz drive rate, but notes still play from DSL's PAL note table, so the
  same note sounds about 0.65 semitone lower than on the NTSC machine.
  *Today:* converters emit the note whose PAL frequency word matches the
  registers (lift's per-tune register shift, typically -1), so the render
  matches the original's register values but not its musical pitch; the
  musical pitch is kept in the notes corpus. A chip clock that also sets
  the note table would make the two agree. *Seen in:*
  MUSICIANS/A/Abaddon/WIP_Q.sid (JCH, NTSC).
- **Player note tables that are not equal-tempered A=440**: DMC's and
  others' frequency tables sit tens of cents off (DMC around +40 cents in
  Utopia.sid), and some are not even equal-tempered. DSL plays its own
  table. *Today:* notes map to the nearest MIDI note, the detuning is lost.
  A song-level `tuning CENTS` (or a custom 96-entry table) would keep it.

## Effects

- **Slides with a target and a speed** (DMC slide-note `$C0-$CF`/`$E0-$EF`
  with start and target note, glide `$D0-$DF`; Future Composer `$E0-$EF`):
  *Today:* the converter starts the note and ties to the target at the next
  row (`fx` tone portamento with speed 0), so the glide itself is lost.

## Notes for converters (not DSL gaps)

These are limits of the GoatTracker song model the converters route
through, not of DSL v2; listed so they are not mistaken for DSL work.

- GoatTracker's relative note column only goes up ($00-$5F); steps down are
  written as absolute notes when the played note is known.
- Wave-table left values $E0 and up are commands, so control bytes with
  noise plus other waveforms are written without their pulse bit.
- Filter programs are not lifted yet.
