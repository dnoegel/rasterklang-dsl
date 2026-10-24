// Package dslgen holds engine-agnostic helpers for EMITTING valid DSL v2
// mini-notation source text from already-decided token/value sequences -
// "given these tokens/numbers, what's the shortest valid DSL v2 spelling of
// them" - as opposed to pkg/dsl, which is about PARSING/representing the
// grammar itself.
//
// It lives next to the grammar because any tool that emits DSL text (a
// tracker-file importer, a register-stream analyser) needs the same "fold
// repeated mini-notation tokens" and "collapse a numeric table into
// ramps/repeats" logic, and should share one tested implementation.
//
// The functions here are Go ports of the token-folding helpers in this
// project's JavaScript GoatTracker importer (foldRepeatingGroups,
// foldPeriodicRuns, collapseRepeats, formatNumericTable) and must produce
// identical output; each function's comment notes the correctness subtlety
// it preserves.
package dslgen

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// SwitchCycleCase is one case of a filter_block field's `switch cycle`
// value (see dsl-grammar-spec.md's "Filter" section): Cond is the
// pre-formatted condition text ("< N", ">= N", or "else"), Value is the
// pre-formatted DSL value text for that case - a bare number/number-list
// (resonance/route), a bare identifier (mode), or a signal_expr (cutoff/
// resonance ramps). Formatting the value itself is the caller's job (it
// depends on which of the four filter fields this is for); this type only
// carries the two already-decided text fragments FormatSwitchCycle joins.
type SwitchCycleCase struct {
	Cond  string
	Value string
}

// FormatSwitchCycle renders a filter field's `switch cycle` value, e.g.
// (for fieldIndent = "  ", the indent every filter_block field line
// already sits at):
//
//	switch cycle
//	    < 4 : ...
//	    else : ...
//
// The returned text starts with "switch cycle" (no leading indent/field
// name - the caller already wrote "  cutoff " or equivalent before this),
// so `fmt.Sprintf("  %s %s\n", fieldName, FormatSwitchCycle(cases, "  "))`
// produces a complete filter_field line.
func FormatSwitchCycle(cases []SwitchCycleCase, fieldIndent string) string {
	caseIndent := fieldIndent + "  "
	var sb strings.Builder
	sb.WriteString("switch cycle")
	for _, c := range cases {
		sb.WriteString("\n")
		sb.WriteString(caseIndent)
		sb.WriteString(c.Cond)
		sb.WriteString(" : ")
		sb.WriteString(c.Value)
	}
	return sb.String()
}

// FoldSingleRuns folds consecutive identical tokens into "X!n". Always safe:
// no bracketing, so every step keeps its own weight.
func FoldSingleRuns(tokens []string) []string {
	out := make([]string, 0, len(tokens))
	for i := 0; i < len(tokens); {
		repeats := 1
		for i+repeats < len(tokens) && tokens[i+repeats] == tokens[i] {
			repeats++
		}
		if repeats > 1 {
			out = append(out, fmt.Sprintf("%s!%d", tokens[i], repeats))
		} else {
			out = append(out, tokens[i])
		}
		i += repeats
	}
	return out
}

var atomWeightRe = regexp.MustCompile(`@(\d+(?:\.\d+)?)$`)

// AtomWeight returns the weight suffix of a mini-notation atom ("c4@3" -> 3);
// a bare atom is weight 1.
func AtomWeight(atom string) float64 {
	m := atomWeightRe.FindStringSubmatch(atom)
	if m == nil {
		return 1
	}
	v, _ := strconv.ParseFloat(m[1], 64)
	return v
}

// FoldRepeatingGroups folds repeated runs of tokens into mini-notation's
// "!n" replicate operator, for GROUPS as well as single tokens: "c2 c2 c3 c2"
// four times over becomes "[c2 c2 c3 c2]!4".
//
// Two different safety situations, so two different strategies:
//
// A bare "X@w!n" is ALWAYS safe: it produces n top-level steps each of weight
// w, exactly as if X@w had been written n times. Nothing is bracketed, so
// nothing is re-timed.
//
// A bracketed "[...]!n" is NOT: at the top level mini-notation divides the
// cycle in proportion to each step's weight, and a GROUP counts as one step
// of weight 1 no matter what it contains. Folding a run whose internal weight
// differs from its neighbours' therefore re-times the pattern. So group
// folding only runs at a chunk size that divides the token list into
// equal-weight chunks, which makes every resulting step equal-width by
// construction. Whichever such chunk size yields the shortest text wins; if
// none compresses, fall back to the always-safe single-token folding.
func FoldRepeatingGroups(tokens []string) []string {
	type candidate struct {
		out  []string
		cost int
	}
	var best *candidate

	for g := 2; g <= len(tokens)/2; g++ {
		if len(tokens)%g != 0 {
			continue
		}
		var chunks []string
		var weights []float64
		for i := 0; i < len(tokens); i += g {
			chunk := tokens[i : i+g]
			chunks = append(chunks, strings.Join(chunk, " "))
			w := 0.0
			for _, t := range chunk {
				w += AtomWeight(t)
			}
			weights = append(weights, w)
		}
		allEqual := true
		for _, w := range weights {
			if math.Abs(w-weights[0]) >= 1e-9 {
				allEqual = false
				break
			}
		}
		if !allEqual {
			continue
		}

		var out []string
		i := 0
		for i < len(chunks) {
			repeats := 1
			for i+repeats < len(chunks) && chunks[i+repeats] == chunks[i] {
				repeats++
			}
			if repeats > 1 {
				out = append(out, fmt.Sprintf("[%s]!%d", chunks[i], repeats))
			} else {
				out = append(out, fmt.Sprintf("[%s]", chunks[i]))
			}
			i += repeats
		}
		if len(out) >= len(chunks) {
			continue // nothing actually folded
		}
		cost := len(strings.Join(out, " "))
		if best == nil || cost < best.cost {
			best = &candidate{out: out, cost: cost}
		}
	}

	singles := FoldSingleRuns(tokens)
	if best != nil && best.cost < len(strings.Join(singles, " ")) {
		return best.out
	}
	return singles
}

// MaxAlternationPeriod is the longest repeating unit worth searching for in
// an orderlist/alternation. Beyond a few cycles a "period" is more likely
// coincidence than structure, and the search is quadratic in it.
const MaxAlternationPeriod = 8

// FoldPeriodicRuns folds an alternation's steps ("slots", one per DSL cycle)
// so that periodic ALTERNATING runs - not just simple repeats - collapse.
//
// An orderlist rarely just repeats one pattern - it alternates two (verse
// figure, answering figure) for a long stretch. "!n" alone can't express
// that, since it replicates a single step, and grouping ("[a b]!7") is wrong:
// inside an alternation "[...]" means "one cycle containing all of this", so
// it would play both patterns simultaneously rather than in succession.
//
// A NESTED alternation does say it: "<<a b>!14>". The replicated copies are
// all the same inner alternation, and each is queried at the OUTER cycle
// index, so copy i yields the inner element at i % 2 - exactly the
// alternation, continued for as long as the run lasts.
//
// The catch, and why the unit is built the way it is below: the inner
// alternation is phase-locked to the ABSOLUTE cycle index, not to where the
// run happens to start. A run of "b a b a" beginning at cycle 7 needs the
// unit written "<a b>", not "<b a>", because cycle 7 selects element 7 % 2 =
// 1. That's what the `unit[m] = slots[i + (((m-i) % period) + period) %
// period]` phase-alignment step below computes: for slot index i and
// candidate period, unit[m] is "whatever this run plays on cycles where
// cycleIndex % period == m", derived from the run's own start index i, not
// from m directly.
func FoldPeriodicRuns(slots []string) []string {
	var out []string
	i := 0
	for i < len(slots) {
		type best struct {
			text   string
			length int
			saved  int
		}
		var bestRun *best
		maxPeriod := MaxAlternationPeriod
		if len(slots)-i < maxPeriod {
			maxPeriod = len(slots) - i
		}
		for period := 1; period <= maxPeriod; period++ {
			unit := make([]string, period)
			for m := 0; m < period; m++ {
				idx := (((m-i)%period)+period)%period + i
				unit[m] = slots[idx]
			}
			length := 0
			for i+length < len(slots) && slots[i+length] == unit[(i+length)%period] {
				length++
			}
			if length < period*2 {
				continue // fewer than two full turns isn't a run
			}
			var text string
			if period == 1 {
				text = fmt.Sprintf("%s!%d", unit[0], length)
			} else {
				text = fmt.Sprintf("<%s>!%d", strings.Join(unit, " "), length)
			}
			saved := len(strings.Join(slots[i:i+length], " ")) - len(text)
			if saved > 0 && (bestRun == nil || saved > bestRun.saved) {
				bestRun = &best{text: text, length: length, saved: saved}
			}
		}
		if bestRun == nil {
			out = append(out, slots[i])
			i++
			continue
		}
		out = append(out, bestRun.text)
		i += bestRun.length
	}
	return out
}

// RenderAlternation returns the mini-notation text for a set of per-cycle
// slots. Two reductions: consecutive identical slots fold into "!n" (via
// FoldPeriodicRuns), and an alternation whose slots are ALL the same
// collapses to that single slot with no "<...>" at all (a one-element
// alternation just repeats every cycle).
func RenderAlternation(slots []string) string {
	if len(slots) == 0 {
		return ""
	}
	allSame := true
	for _, s := range slots {
		if s != slots[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return slots[0]
	}
	return "<" + strings.Join(FoldPeriodicRuns(slots), " ") + ">"
}

// CollapseRepeats folds runs of an identical WEIGHT-1 token into "token!n".
// Only weight-1 tokens (no "@" suffix) are folded: a token already carrying
// an "@weight" suffix would produce "x@2!3", stacking two suffixes whose
// combined meaning isn't worth relying on here for the little it would save.
func CollapseRepeats(atoms []string) []string {
	out := make([]string, 0, len(atoms))
	for i := 0; i < len(atoms); {
		run := 1
		for i+run < len(atoms) && atoms[i+run] == atoms[i] {
			run++
		}
		foldable := !strings.Contains(atoms[i], "@") && run > 1
		if foldable {
			out = append(out, fmt.Sprintf("%s!%d", atoms[i], run))
			i += run
		} else {
			out = append(out, atoms[i])
			i++
		}
	}
	return out
}

// MinRampRun: below this many frames a "ramp" is longer to read than the
// values it replaces, so short runs stay literal.
const MinRampRun = 4

// MinRepeatRun: "2184!3" is shorter than "2184 2184 2184", so repeats pay
// off one step sooner than ramps do.
const MinRepeatRun = 3

// clampedRamp is what ClampedRampAt returns when a single "ramp" reproduces
// values[start:start+Length].
type clampedRamp struct {
	step   int
	target int
	length int
}

// clampedRampAt finds the shape of a single `ramp` token that would
// reproduce values from `start`, else returns ok=false.
//
// The shape is a constant-delta stretch followed (optionally) by a run
// holding at the value the stretch CLAMPED to - which is not simply "the
// last value repeated": a saturating sweep steps 1088, 1040, ... 32, then 0
// (32-48 saturated at the floor), then 0 forever. The saturated value is a
// partial step past the last full one, so the tail is only part of THIS ramp
// when it lies beyond the stretch's end in the step's own direction.
// "ramp A..T step D over N" clamps to T identically, so the whole row
// round-trips as one token.
func clampedRampAt(values []int, start int) (clampedRamp, bool) {
	if start+1 >= len(values) {
		return clampedRamp{}, false
	}
	step := values[start+1] - values[start]
	if step == 0 {
		return clampedRamp{}, false
	}
	end := start + 1
	for end+1 < len(values) && values[end+1]-values[end] == step {
		end++
	}

	target := values[end]
	length := end - start + 1
	if end+1 < len(values) {
		next := values[end+1]
		if sign(next-values[end]) == sign(step) {
			// The saturated value, and every frame holding at it, belongs to
			// this ramp.
			tail := end + 1
			for tail+1 < len(values) && values[tail+1] == next {
				tail++
			}
			// ...but only if the ramp really does reproduce it. The tail is a
			// SATURATION hold, and `next` past a constant-delta run is only
			// that when the step would have overshot it. A run that simply
			// takes one bigger step and carries on (a GoatTracker SIMPLEPULSE
			// pulse sweep does exactly that whenever its packed byte wraps:
			// ... 706 738 771 803 ...) would otherwise be swallowed and
			// re-expanded one short.
			if rampExpandsTo(values[start], next, step, tail-start+1, values[start:tail+1]) {
				target = next
				length = tail - start + 1
			}
		}
	}
	return clampedRamp{step: step, target: target, length: length}, true
}

// rampExpandsTo reports whether `ramp from..to step S over N` reproduces
// `want` exactly. It mirrors interpret.js's expandRamp() - clamp to the
// [min,max] of the endpoints, no overshoot - so codegen can never emit a
// generator that reads back as different numbers than it was folded from.
func rampExpandsTo(from, to, step, over int, want []int) bool {
	if over != len(want) {
		return false
	}
	lo, hi := from, to
	if lo > hi {
		lo, hi = hi, lo
	}
	for i, w := range want {
		v := from + step*i
		if v < lo {
			v = lo
		}
		if v > hi {
			v = hi
		}
		if v != w {
			return false
		}
	}
	return true
}

func sign(v int) int {
	if v > 0 {
		return 1
	}
	if v < 0 {
		return -1
	}
	return 0
}

// FormatNumericTable returns the DSL source text for a per-frame numeric
// table, collapsing linear runs into "ramp" generators and constant runs
// into "V!n" repeats.
//
// The shape being matched is specifically a CLAMPED ramp - a constant-delta
// run that reaches its target and then holds there for the rest of the row -
// because "ramp A..B step D over N" clamps identically, so one `ramp` token
// reproduces such a run exactly. Note that `ramp` saturates by design; it is
// not a model of GoatTracker's filter modulation step, whose 8-bit cutoff
// accumulator wraps (ADR 0004).
func FormatNumericTable(values []int) string {
	var parts []string
	i := 0
	for i < len(values) {
		ramp, ok := clampedRampAt(values, i)
		if ok && ramp.length >= MinRampRun {
			parts = append(parts, fmt.Sprintf("ramp %d..%d step %d over %d", values[i], ramp.target, ramp.step, ramp.length))
			i += ramp.length
			continue
		}
		end := i
		for end+1 < len(values) && values[end+1] == values[i] {
			end++
		}
		runLength := end - i + 1
		if runLength >= MinRepeatRun {
			parts = append(parts, fmt.Sprintf("%d!%d", values[i], runLength))
			i += runLength
			continue
		}
		parts = append(parts, strconv.Itoa(values[i]))
		i++
	}
	return strings.Join(parts, " ")
}

// ArpCell is one cell of an `arp` table row: either a RELATIVE semitone
// offset from the sounding note (Abs false, DSL v2's plain numeric cell)
// or an ABSOLUTE MIDI note pinning that frame to a fixed pitch (Abs true,
// DSL v2's `=60` cell, added in 2.15.0 for GoatTracker wave-table note
// bytes $80-$DF - see the grammar's absNote rule).
type ArpCell struct {
	Value int
	Abs   bool
}

// FormatArpTable returns the DSL source text for an `arp` row that mixes
// relative and absolute cells. Runs of the same kind are formatted
// together, so a purely relative table produces byte-for-byte what
// FormatNumericTable already produced (ramp/repeat folding included) and
// only the genuinely absolute stretches gain `=N` cells.
func FormatArpTable(cells []ArpCell) string {
	var parts []string
	for i := 0; i < len(cells); {
		j := i
		for j < len(cells) && cells[j].Abs == cells[i].Abs {
			j++
		}
		if !cells[i].Abs {
			vals := make([]int, 0, j-i)
			for _, c := range cells[i:j] {
				vals = append(vals, c.Value)
			}
			parts = append(parts, FormatNumericTable(vals))
		} else {
			// Constant-run folding only: a `ramp` generator has no
			// absolute-note spelling, and a ramp of fixed drum pitches is
			// not a shape real wave tables produce anyway.
			for k := i; k < j; {
				end := k
				for end+1 < j && cells[end+1].Value == cells[k].Value {
					end++
				}
				runLength := end - k + 1
				if runLength >= MinRepeatRun {
					parts = append(parts, fmt.Sprintf("=%d!%d", cells[k].Value, runLength))
				} else {
					for n := 0; n < runLength; n++ {
						parts = append(parts, fmt.Sprintf("=%d", cells[k].Value))
					}
				}
				k += runLength
			}
		}
		i = j
	}
	return strings.Join(parts, " ")
}
