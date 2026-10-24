package dsl

// Go port of src/dsl/miniLength.js: the structural scanner that answers
// "how many BEATS do this mini-notation string's contents add up to", plus
// the hold-token (`_`) rewriting that goes with it.
//
// It exists on this side because Go generators EMIT declared pattern
// lengths (`pattern bass 4 = "..."`, ADR 0001's D1) and hold tokens, and a
// generator that computes the length it writes with different arithmetic
// than the compiler checks it with would produce files that only fail on
// the other side of the pipeline. test/conformance's mini-length goldens
// are generated from the JS implementation and asserted here, exactly like
// the parser goldens, so any drift between the two shows up as a test
// failure rather than as a mis-timed song.
//
// See miniLength.js's header for the reasoning behind the rules themselves;
// the comments here only cover what is specific to this port.

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// miniWeightEpsilon mirrors miniLength.js's WEIGHT_EPSILON: decimal weights
// (`@0.25`) make exact float equality useless for "same length".
const miniWeightEpsilon = 1e-6

// defaultBeatUnit mirrors miniLength.js's DEFAULT_BEAT_UNIT: the meter's
// denominator D (`meter N/D`, ADR 0003 D2) a `/N` duration token converts
// against, for every one of this port's callers - none of which have a
// song's actual meter to thread through yet. See miniLength.js's own
// comment for the same gap on the JS side.
const defaultBeatUnit = 4

// HoldToken is `_`, the hold: the previous step keeps sounding for one more
// step's worth of time (ADR 0001's D6).
const HoldToken = '_'

// BeatsEqual reports whether two beat counts are the same length.
func BeatsEqual(a, b float64) bool { return math.Abs(a-b) < miniWeightEpsilon }

type miniStepKind int

const (
	stepAtom miniStepKind = iota
	stepRest
	stepHold
	stepGroup
	stepPoly
	stepAlt
	stepChoice
	stepRef
)

type miniStep struct {
	kind      miniStepKind
	text      string
	weight    float64
	replicate int
	suffix    string
	seqs      [][]*miniStep // group/poly: one per ','-stacked sequence
	entries   []*miniStep   // alt/choice elements
}

type miniScanner struct {
	str      string
	pos      int
	err      error
	beatUnit float64
	// defaultWeight is the weight a step gets when it carries no explicit
	// `@n`/`/N` of its own (ADR 0003 D3, `each 1/D`) - 1 unless expanding
	// an `each`-declared pattern, in which case it is beatUnit/D. Mirrors
	// miniLength.js's MiniScanner.defaultWeight.
	defaultWeight float64
	// Position of the first `/N` duration and first `@n` weight seen so
	// far, for the "mixed /N and @n in one pattern" check (D1, D14). -1
	// means "not seen". Mirrors miniLength.js's MiniScanner one for one.
	sawDurationAt int
	sawWeightAt   int
	// refs (ADR 0003 D5): name -> beats, nil/ok=false for unresolved. A
	// $name reference's weight is the referenced pattern's own beats, not
	// the default 1 - mirrors miniLength.js's MiniScanner.refs.
	refs func(name string) (float64, bool)
	// checkBars/barBeats (ADR 0003 D6): a standalone `|` is checked against
	// barBeats when checkBars is true; mirrors miniLength.js's barBeats
	// (nil there, a bool split here since Go floats have no null).
	checkBars bool
	barBeats  float64
}

func newMiniScanner(str string) *miniScanner {
	return &miniScanner{str: str, beatUnit: defaultBeatUnit, defaultWeight: 1, sawDurationAt: -1, sawWeightAt: -1}
}

func (s *miniScanner) fail(format string, args ...interface{}) {
	if s.err == nil {
		s.err = fmt.Errorf("%s in mini-notation %q", fmt.Sprintf(format, args...), s.str)
	}
}

func (s *miniScanner) eof() bool { return s.pos >= len(s.str) || s.err != nil }

func (s *miniScanner) peek() byte {
	if s.pos >= len(s.str) {
		return 0
	}
	return s.str[s.pos]
}

func isSpace(c byte) bool { return c == ' ' || c == '\t' || c == '\n' || c == '\r' }

func (s *miniScanner) skipWs() {
	for s.pos < len(s.str) && isSpace(s.str[s.pos]) {
		s.pos++
	}
}

func (s *miniScanner) expect(ch byte) {
	if s.peek() != ch {
		s.fail("expected '%c' at position %d", ch, s.pos)
		return
	}
	s.pos++
}

func isDigitOrDot(c byte) bool { return (c >= '0' && c <= '9') || c == '.' }

func (s *miniScanner) parseNumberLiteral() (float64, string) {
	start := s.pos
	if s.peek() == '-' {
		s.pos++
	}
	for s.pos < len(s.str) && isDigitOrDot(s.str[s.pos]) {
		s.pos++
	}
	raw := s.str[start:s.pos]
	if raw == "" || raw == "-" {
		s.fail("expected a number at position %d", start)
		return 0, ""
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		s.fail("expected a number at position %d", start)
		return 0, ""
	}
	return v, raw
}

// tryParseNumberLiteral returns ok=false when the next characters are not a
// number, matching miniNotation.js's "bare `!` means 2" handling.
func (s *miniScanner) tryParseNumberLiteral() (float64, string, bool) {
	if s.eof() {
		return 0, "", false
	}
	ch := s.peek()
	var next byte
	if s.pos+1 < len(s.str) {
		next = s.str[s.pos+1]
	}
	if isDigitOrDot(ch) || (ch == '-' && next >= '0' && next <= '9') {
		v, txt := s.parseNumberLiteral()
		return v, txt, true
	}
	return 0, "", false
}

// parseDuration parses a `/` modifier's tail: `INT ('.' | 't')?`. A
// dedicated digit scan, not parseNumberLiteral, because that would swallow
// the dot of `/4.` as a decimal point. Mirrors miniLength.js's
// MiniScanner.parseDuration() exactly.
func (s *miniScanner) parseDuration() float64 {
	start := s.pos
	for s.pos < len(s.str) && s.str[s.pos] >= '0' && s.str[s.pos] <= '9' {
		s.pos++
	}
	digits := s.str[start:s.pos]
	if digits == "" {
		s.fail("expected a duration (e.g. /4, /8., /4t) at position %d", start)
		return 0
	}
	denom, _ := strconv.ParseFloat(digits, 64)
	beats := s.beatUnit / denom
	switch s.peek() {
	case '.':
		s.pos++
		beats *= 1.5
	case 't':
		s.pos++
		beats *= 2.0 / 3.0
	}
	return beats
}

func (s *miniScanner) failMixedDurationWeight() {
	s.fail(
		"mixes /N durations (first at position %d) with @N weights (first at position %d) - "+
			"one is a duration and the other a proportional weight, use one or the other",
		s.sawDurationAt, s.sawWeightAt,
	)
}

func isWordChar(c byte) bool {
	// ':' only ever appears as part of 'slideTo:' (ADR 0003 D12) - mirrors
	// miniNotation.js's scanWordToken exactly.
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '.' || c == '#' || c == ':'
}

func (s *miniScanner) scanWordToken() string {
	start := s.pos
	if s.peek() == '-' && s.pos+1 < len(s.str) && s.str[s.pos+1] >= '0' && s.str[s.pos+1] <= '9' {
		s.pos++
	}
	for s.pos < len(s.str) && isWordChar(s.str[s.pos]) {
		s.pos++
	}
	raw := s.str[start:s.pos]
	if raw == "" {
		s.fail("expected a step at position %d", start)
	}
	return raw
}

func newStep(kind miniStepKind, text string) *miniStep {
	return &miniStep{kind: kind, text: text, weight: 1, replicate: 1}
}

func (s *miniScanner) parseAtomUnit() *miniStep {
	s.skipWs()
	if s.err != nil {
		return newStep(stepAtom, "")
	}
	ch := s.peek()
	switch {
	case ch == '~':
		s.pos++
		return newStep(stepRest, "~")
	case ch == '-' && !(s.pos+1 < len(s.str) && s.str[s.pos+1] >= '0' && s.str[s.pos+1] <= '9'):
		s.pos++
		return newStep(stepRest, "-")
	case ch == HoldToken:
		s.pos++
		return newStep(stepHold, string(HoldToken))
	case ch == '$' && s.refs != nil:
		return s.parseRef()
	case ch == '[' || ch == '{':
		closer := byte(']')
		kind := stepGroup
		if ch == '{' {
			closer = '}'
			kind = stepPoly
		}
		s.pos++
		seqs := [][]*miniStep{s.parseSequence(string([]byte{',', closer}))}
		s.skipWs()
		for s.peek() == ',' {
			s.pos++
			seqs = append(seqs, s.parseSequence(string([]byte{',', closer})))
			s.skipWs()
		}
		s.expect(closer)
		st := newStep(kind, "")
		st.seqs = seqs
		return st
	case ch == '<':
		s.pos++
		var entries []*miniStep
		s.skipWs()
		for !s.eof() && s.peek() != '>' {
			entries = append(entries, s.parseStep())
			s.skipWs()
		}
		s.expect('>')
		st := newStep(stepAlt, "")
		st.entries = entries
		return st
	case ch == 0:
		s.fail("unexpected end of input while parsing a step")
		return newStep(stepAtom, "")
	}
	return newStep(stepAtom, s.scanWordToken())
}

// parseRef parses `$name`, already confirmed present (s.refs != nil) by
// parseAtomUnit. Mirrors miniLength.js's MiniScanner $ handling: a
// reference's weight is the referenced pattern's own beats, not the usual
// 1 - ADR 0003 D5's F1 fix, "a metric pattern referenced anywhere occupies
// its metric length".
func (s *miniScanner) parseRef() *miniStep {
	start := s.pos
	s.pos++ // '$'
	nameStart := s.pos
	for s.pos < len(s.str) && isRefNameChar(s.str[s.pos]) {
		s.pos++
	}
	name := s.str[nameStart:s.pos]
	if name == "" {
		s.fail("expected a pattern name after '$' at position %d", start)
		return newStep(stepAtom, "")
	}
	beats, ok := s.refs(name)
	if !ok {
		s.fail("pattern \"$%s\" is not defined (declare it with `pattern %s = \"...\"` before this use - forward references aren't supported)", name, name)
		return newStep(stepAtom, "")
	}
	st := newStep(stepRef, "$"+name)
	st.weight = beats
	return st
}

func isRefNameChar(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_'
}

func (s *miniScanner) parseChoice() *miniStep {
	first := s.parseAtomUnit()
	if s.peek() != '|' {
		return first
	}
	branches := []*miniStep{first}
	for s.peek() == '|' {
		s.pos++
		branches = append(branches, s.parseAtomUnit())
	}
	st := newStep(stepChoice, "")
	st.entries = branches
	return st
}

func (s *miniScanner) parseStep() *miniStep {
	step := s.parseChoice()
	// each's default (see the struct field comment); '@'/'/' below overwrite
	// it with the step's own explicit weight/duration when present. Not
	// applied to a $name reference (ADR 0003 D5): its weight is already the
	// REFERENCED pattern's own beats (set in parseRef), which `each` has no
	// business overriding - `@n`/`/N` still can, explicitly.
	if step.kind != stepRef {
		step.weight = s.defaultWeight
	}
	for s.err == nil {
		ch := s.peek()
		switch ch {
		case '@':
			modStart := s.pos
			s.pos++
			v, _ := s.parseNumberLiteral()
			step.weight = v
			if s.sawWeightAt == -1 {
				s.sawWeightAt = modStart
			}
			if s.sawDurationAt != -1 {
				s.failMixedDurationWeight()
			}
		case '!':
			s.pos++
			v, _, ok := s.tryParseNumberLiteral()
			if !ok {
				v = 2
			}
			step.replicate = int(v)
		case '/':
			modStart := s.pos
			s.pos++
			step.weight = s.parseDuration()
			if s.sawDurationAt == -1 {
				s.sawDurationAt = modStart
			}
			if s.sawWeightAt != -1 {
				s.failMixedDurationWeight()
			}
		case '*':
			s.pos++
			_, txt := s.parseNumberLiteral()
			step.suffix += "*" + txt
		case '?':
			s.pos++
			_, txt, ok := s.tryParseNumberLiteral()
			if !ok {
				txt = ""
			}
			step.suffix += "?" + txt
		case '(':
			start := s.pos
			rel := strings.IndexByte(s.str[s.pos:], ')')
			if rel == -1 {
				s.fail("unclosed '(' at position %d", start)
				return step
			}
			s.pos += rel + 1
			step.suffix += s.str[start:s.pos]
		default:
			return step
		}
	}
	return step
}

func (s *miniScanner) parseSequence(stopChars string) []*miniStep {
	s.skipWs()
	var steps []*miniStep
	runningBeats := 0.0
	for !s.eof() && strings.IndexByte(stopChars, s.peek()) == -1 {
		if s.peek() == '|' {
			barPos := s.pos
			s.pos++
			if s.checkBars && !BeatsEqual(math.Mod(runningBeats, s.barBeats), 0) && !BeatsEqual(math.Mod(runningBeats, s.barBeats), s.barBeats) {
				s.fail("bar line '|' at position %d is off the grid - it lands at beat %v, not a multiple of the %v-beat bar", barPos, runningBeats, s.barBeats)
			}
			s.skipWs()
			continue
		}
		step := s.parseStep()
		steps = append(steps, step)
		runningBeats += step.weight * float64(step.replicate)
		s.skipWs()
	}
	return steps
}

func (s *miniScanner) parseTop() [][]*miniStep {
	seqs := [][]*miniStep{s.parseSequence(",")}
	for s.peek() == ',' {
		s.pos++
		seqs = append(seqs, s.parseSequence(","))
	}
	if s.err == nil && s.pos < len(s.str) {
		s.fail("unexpected '%c' at position %d", s.peek(), s.pos)
	}
	return seqs
}

func scanMini(source string) ([][]*miniStep, error) {
	return scanMiniWith(source, defaultBeatUnit, 1)
}

func scanMiniWith(source string, beatUnit, defaultWeight float64) ([][]*miniStep, error) {
	s := newMiniScanner(source)
	s.beatUnit = beatUnit
	s.defaultWeight = defaultWeight
	seqs := s.parseTop()
	if s.err != nil {
		return nil, s.err
	}
	return seqs, nil
}

// CheckBarLines mirrors miniLength.js's checkBarLines() (ADR 0003 D6):
// returns an error if any standalone `|` in source does not land on a
// multiple of barBeats beats, naming the beat offset it actually reached.
// Purely a check - it does not change how the pattern plays.
func CheckBarLines(source string, barBeats, beatUnit float64) error {
	s := newMiniScanner(source)
	s.beatUnit = beatUnit
	s.checkBars = true
	s.barBeats = barBeats
	s.parseTop()
	return s.err
}

func sequenceBeats(steps []*miniStep) float64 {
	total := 0.0
	for _, st := range steps {
		total += st.weight * float64(st.replicate)
	}
	return total
}

// MiniContentBeats reports how many beats one pass of a pattern's contents
// occupies, reading every unweighted token as one beat. Stacked (`,`)
// sequences sound together, so a stack is as long as its longest layer.
func MiniContentBeats(source string) (float64, error) {
	seqs, err := scanMini(source)
	if err != nil {
		return 0, err
	}
	max := 0.0
	for _, s := range seqs {
		if b := sequenceBeats(s); b > max {
			max = b
		}
	}
	return max, nil
}

func stepBeats(step *miniStep) float64 {
	if step.kind == stepGroup || step.kind == stepPoly {
		max := 0.0
		for _, s := range step.seqs {
			if b := sequenceBeats(s); b > max {
				max = b
			}
		}
		return max * step.weight
	}
	return step.weight
}

// MiniLengthCandidate is one content length a declared length has to match.
// What names it ("contents", "orderlist element 2", ...) for the error
// message.
type MiniLengthCandidate struct {
	Beats float64
	What  string
}

// MiniDeclaredLengthCandidates mirrors miniLength.js's
// declaredLengthCandidates(), including its orderlist rule: a pattern whose
// whole content is one `<...>` alternation declares the length of EACH
// element, not of their sum (ADR 0001 decision 1).
func MiniDeclaredLengthCandidates(source string) ([]MiniLengthCandidate, error) {
	seqs, err := scanMini(source)
	if err != nil {
		return nil, err
	}
	if len(seqs) == 1 && len(seqs[0]) == 1 && seqs[0][0].kind == stepAlt && seqs[0][0].weight == 1 {
		entries := seqs[0][0].entries
		out := make([]MiniLengthCandidate, 0, len(entries))
		for i, e := range entries {
			out = append(out, MiniLengthCandidate{Beats: stepBeats(e), What: fmt.Sprintf("orderlist element %d", i+1)})
		}
		return out, nil
	}
	out := make([]MiniLengthCandidate, 0, len(seqs))
	for i, s := range seqs {
		what := "contents"
		if len(seqs) > 1 {
			what = fmt.Sprintf("stacked layer %d", i+1)
		}
		out = append(out, MiniLengthCandidate{Beats: sequenceBeats(s), What: what})
	}
	return out, nil
}

// HasHoldToken reports whether a string contains `_` at all - the cheap
// pre-check that lets ExpandHolds return pre-3.0.0 strings untouched.
func HasHoldToken(source string) bool { return strings.IndexByte(source, HoldToken) >= 0 }

// ExpandHolds folds every hold into the weight of the step it extends
// (`"c4 _ _ _"` -> `"c4@4"`), which is how the hold token is expressed to a
// pattern engine that only knows weights. Returns source unchanged when it
// contains no hold.
func ExpandHolds(source string) (string, error) {
	if !HasHoldToken(source) {
		return source, nil
	}
	seqs, err := scanMini(source)
	if err != nil {
		return "", err
	}
	parts := make([]string, 0, len(seqs))
	for _, s := range seqs {
		rendered, err := renderSteps(s, "")
		if err != nil {
			return "", err
		}
		parts = append(parts, rendered)
	}
	return strings.Join(parts, ", "), nil
}

// ExpandEach mirrors miniLength.js's expandEach(): every step's default
// weight becomes beatUnit/denominator instead of 1 (ADR 0003 D3's
// `each 1/D`), rendered out explicitly (`@n`), with any hold tokens folded
// in the same pass. A step already carrying its own `@n`/`/N` is untouched.
func ExpandEach(source string, denominator, beatUnit float64) (string, error) {
	seqs, err := scanMiniWith(source, beatUnit, beatUnit/denominator)
	if err != nil {
		return "", err
	}
	parts := make([]string, 0, len(seqs))
	for _, s := range seqs {
		rendered, err := renderSteps(s, "")
		if err != nil {
			return "", err
		}
		parts = append(parts, rendered)
	}
	return strings.Join(parts, ", "), nil
}

type foldedStep struct {
	step  *miniStep
	extra float64
}

func renderSteps(steps []*miniStep, path string) (string, error) {
	var folded []foldedStep
	for _, step := range steps {
		if step.kind == stepHold {
			if len(folded) == 0 {
				where := ""
				if path != "" {
					where = " (inside " + path + ")"
				}
				return "", fmt.Errorf("the hold token '_' extends the step before it, but there is no step before it%s - start the sequence with a note or a rest.", where)
			}
			folded[len(folded)-1].extra += step.weight * float64(step.replicate)
			continue
		}
		folded = append(folded, foldedStep{step: step})
	}

	out := make([]string, 0, len(folded))
	for _, f := range folded {
		body, err := renderStepBody(f.step)
		if err != nil {
			return "", err
		}
		if f.extra == 0 {
			w := ""
			if f.step.weight != 1 {
				w = "@" + formatMiniNumber(f.step.weight)
			}
			r := ""
			if f.step.replicate != 1 {
				r = "!" + strconv.Itoa(f.step.replicate)
			}
			out = append(out, body+f.step.suffix+w+r)
			continue
		}
		for i := 0; i < f.step.replicate-1; i++ {
			w := ""
			if f.step.weight != 1 {
				w = "@" + formatMiniNumber(f.step.weight)
			}
			out = append(out, body+f.step.suffix+w)
		}
		out = append(out, body+f.step.suffix+"@"+formatMiniNumber(f.step.weight+f.extra))
	}
	return strings.Join(out, " "), nil
}

func renderStepBody(step *miniStep) (string, error) {
	switch step.kind {
	case stepGroup, stepPoly:
		open, closer, path := "[", "]", "a [...] group"
		if step.kind == stepPoly {
			open, closer, path = "{", "}", "a {...} group"
		}
		parts := make([]string, 0, len(step.seqs))
		for _, s := range step.seqs {
			rendered, err := renderSteps(s, path)
			if err != nil {
				return "", err
			}
			parts = append(parts, rendered)
		}
		return open + strings.Join(parts, ", ") + closer, nil
	case stepAlt:
		parts := make([]string, 0, len(step.entries))
		for _, e := range step.entries {
			rendered, err := renderSteps([]*miniStep{e}, "a <...> alternation")
			if err != nil {
				return "", err
			}
			parts = append(parts, rendered)
		}
		return "<" + strings.Join(parts, " ") + ">", nil
	case stepChoice:
		parts := make([]string, 0, len(step.entries))
		for _, e := range step.entries {
			rendered, err := renderSteps([]*miniStep{e}, "a choice")
			if err != nil {
				return "", err
			}
			parts = append(parts, rendered)
		}
		return strings.Join(parts, "|"), nil
	default:
		return step.text, nil
	}
}

// formatMiniNumber matches JS's Number->string spelling for the values a
// folded weight can take: an integer prints without a decimal point, a
// fraction prints with the shortest exact decimal.
func formatMiniNumber(n float64) string {
	rounded := math.Round(n*1e6) / 1e6
	return strconv.FormatFloat(rounded, 'g', -1, 64)
}
