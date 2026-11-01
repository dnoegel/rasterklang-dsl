# ADR 0006: The tracker-conversion backlog (ties, engine fixes, note tables)

Status: implemented (4.4.0)
Date: 2026-11-01
Affects: `rasterklang-dsl` (hold expansion, `target`, `hardRestart`, JS and Go visitors and
mini-length), `rasterklang-score` (pattern parser, compiler), and the converters in
`rasterklang-conv`
Current DSL version: 4.3.0. This ADR targets 4.4.0.

**Compatibility statement.** The new spellings are additive: a leading hold, `target { pitch }`,
`target { tuning }` and `hardRestart N beforeNotes` were all compile errors before 4.4.0. Three
decisions (D2, D3) are engine fixes and do change how existing songs play where they hit the
defect; the owner chose to fix them by default rather than behind a switch. Measured on the same
815-file corpus as ADR 0005 (785 compile), comparing every register on every one of 1500 frames
between the old and the new code:

| change | songs | registers |
| --- | --- | --- |
| `insts` envelope (D2) | 673 | AD/SR only |
| arp through the release (D3) | 9 | frequency, in rests; in one song a following `slide` starts from the new frequency |
| ties under `gate -N` (D2) | 0 | none in this corpus |
| everything else | 0 | gate, pulse width and filter are unchanged in every song |

(666 songs change only in AD/SR, 7 in AD/SR and frequency, 2 only in frequency.) The 30
intentional error fixtures still fail; one message lists the two new `target` settings.

---

## 1. Context

`docs/tracker-features-todo.md` collects player mechanisms that the converters met while lifting
DMC, Music Assembler and other trackers, with the workaround each converter uses today. Testing
every entry against the compiler before deciding showed three kinds of entries:

1. Already expressible, so the work is in the converter, not the language.
2. Real gaps that an additive spelling closes.
3. Engine defects, two of them worse than the backlog describes.

| backlog entry | finding | decision |
| --- | --- | --- |
| per-note row length | `c4@3 e4` already plays a 3-row note, also under `gate -N` | converter (section 8) |
| notes across a pattern boundary | a leading `_` was an error | D1 |
| tie: new pitch without a new gate | `fx tp.0` was already legato, but under `gate -N` the tied note's gate dropped at once and stayed off (silent) | D2 |
| tables run through the release | confirmed: `arp` froze in a rest, `wave`/`pw` ran on | D3 |
| `insts` does not switch AD/SR | confirmed, and worse: a per-note instrument's `adsr` was ignored entirely | D2 |
| release before the next note, not before a key-off | `gate -N` also released early before rests | D5 |
| pulse sweeps that run on, bounce or wrap | `ramp ... pingpong` (instrument) and 4.3.0's `fx pws`/`pwm` (row) cover it | converter (section 8) |
| NTSC SID clock | `video NTSC` changed the drive rate only, never the note table | D4 |
| note tables off A=440 | no spelling | D4 |
| slides with a target and a speed | `fx tp.N` glides toward the row's note at N per frame; blocked only by the D2 defect | converter (section 8) |

## 2. Decision D1: a hold with nothing before it is a tie

A `_` with no step before it in its sequence continues whatever note sounds at that point in time:
the previous bar's last note, the previous `<...>` entry's, or the step before an enclosing group.
Implementation: hold expansion (`miniLength.js`, `minilength.go`) keeps such a hold as a bare `_`
(summing a run of them into one weight) instead of throwing; the pattern parser turns it into a Hap
tagged `tie: true`; `compile.js` replaces a tie Hap with the tones sounding on the previous frame
and treats the transition as legato (no onset, tables keep running), or as silence if nothing
sounds. A hard restart does not release into it (D2).

It continues notes only. In an `fx` or `insts` pattern a leading `_` stays an error, now with a
message saying so. Alternatives considered: a weight that reaches past the bar (it would make a
pattern's length depend on its neighbour), and a dedicated tie token (a second spelling of "this
step continues the sound" next to `_`).

## 3. Decision D2: legato transitions and per-note envelopes (engine fixes)

- **Gate length counts from the note's own start.** A note entered legato (tie, `fx tp`,
  `slideTo`) used the frame counter of the note it was tied from, so `gate -N` and a fractional
  `gate` cut it at once. The compiler now keeps a second counter, frames since the note's own Hap
  began, for gate lengths; the onset counter stays for instrument tables, which a legato note
  deliberately keeps running. For every note entered with an onset the two counters are equal,
  which is why no other song changes.
- **No early release into a legato transition.** `gate -N`/`hardRestart` look at what starts where
  the note ends; for a tie, a `slideTo` or an `fx tp` it keeps the gate open, since dropping it
  would retrigger the very note the tie says keeps sounding.
- **`insts` brings the instrument's `adsr`.** A per-note instrument now supplies its own envelope,
  as a tracker's instrument column does. A voice with a single instrument already had it folded
  into the chain, so only `insts` voices change.

## 4. Decision D3: arp tables run through the release

During a rest the last instrument's `arp` table keeps running from the last note's pitch, as its
`wave` and `pw` tables already did and as a tracker's wave table does after a key-off. The register
written is the arp pitch; the gate stays closed. A `slide` on the next note starts from the
frequency last written, which is now this one.

## 5. Decision D4: the note table's clock and tuning

`target { pitch PAL|NTSC }` computes the note table at that machine's clock, and
`target { tuning CENTS }` (-1200..1200) shifts every note. Without either the table is the PAL
equal-tempered A=440 table every song had, so `video NTSC` keeps meaning what it meant (drive rate
only) and converters that compensate for it keep working. Both live in `target` because they are
facts about the machine, next to the drive rate; `target` settings are generic name/value pairs, so
no grammar change was needed.

A custom 96-entry note table (for players whose table is not equal-tempered at all) is left for
later: no converter needs it yet, and `tuning` covers the measured case (DMC, about +40 cents).

## 6. Decision D5: `hardRestart N beforeNotes`

With `beforeNotes` the early release happens only when another note follows. Before a rest (a
key-off) the gate stays open until the rest begins, as DMC's "release one tick before the next
note" flag does. It combines with ADR 0005's `ad`/`sr`/`test` clauses, which already applied only
before a note.

## 7. Implementation notes

- No grammar change and no parser regeneration.
- Go mirrors the JS visitor (`pitch`, `tuning`, `beforeNotes`) and `minilength.go`'s hold rendering;
  the mini-length golden for `"_ c4"` and `"c4 [_ e4]"` changes from an error to a result, which is
  the point of D1.
- `analyze.test.js`'s "rejects a hold with nothing before it" became "accepts".

## 8. For the converters (no language change needed)

- **Row lengths:** emit `c4@3` instead of `c4 _ _`.
- **Ties and slides:** `fx tp.0` for a tie and `fx tp.N` (or `sN`) for a slide now behave under a
  hard restart; D1's leading `_` replaces the `~` written at pattern starts today.
- **Pulse sweeps:** a bounded sweep is `pw ramp LO..HI step S pingpong` (several ramps plus `loop
  pw K` for a start inside the range); a sweep that runs on across notes or wraps is `fx pws`/`pwm`.
- **NTSC tunes:** emit `target { video NTSC pitch NTSC }` with the musical notes, instead of the
  per-tune register shift.
- **`insts` voices:** the envelope now follows the instrument, so a converter no longer needs to
  approximate it on the voice.
