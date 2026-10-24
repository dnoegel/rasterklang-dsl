package dsl

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	gen "github.com/dnoegel/rasterklang-dsl/pkg/dsl/generated"
)

// vibTriShiftRe matches the shift-derived speed spelling `sN` shared with
// the `fx` lane's own parameters.
var vibTriShiftRe = regexp.MustCompile(`^s[0-9]+$`)

var instrumentFieldsList = []string{"wave", "adsr", "pw", "arp", "filterTable", "loop", "vib", "vibTri", "slide", "test", "gate", "hardRestart"}
var gateModesList = []string{"legato", "retrigger"}
var voiceFieldsList = []string{"notes", "insts", "fx", "reg", "vib", "slide", "gate", "sync", "ring", "wave"}
var transformsList = []string{
	"rev", "fast", "slow", "every", "iter", "ply", "off", "jux", "degradeBy", "sometimesBy", "add",
	"struct", "mask", "chunk", "palindrome", "linger", "rot", "someBarsBy", "someCyclesBy", "superimpose", "scale",
}
var filterFieldsList = []string{"cutoff", "resonance", "mode", "route"}
var loopTableNames = []string{"wave", "pw", "arp", "filterTable", "test", "gate"}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

// astErr is a sentinel panic value caught at the top of Build(), letting
// every builder function just call fail() instead of threading (T, error)
// through the whole visitor.
type astErr struct{ err error }

func fail(line int, format string, args ...interface{}) {
	panic(astErr{fmt.Errorf("line %d: %s", line, fmt.Sprintf(format, args...))})
}

type visitor struct {
	tokens []Token
}

func (v *visitor) orig(t antlr.Token) Token {
	return v.tokens[t.GetTokenIndex()]
}

// Build converts an ANTLR parse tree (gen.IProgramContext) to *Program,
// matching parser.js/parser.go's existing node shapes exactly.
func buildProgram(tree gen.IProgramContext, tokens []Token) (prog *Program, err error) {
	v := &visitor{tokens: tokens}
	defer func() {
		if r := recover(); r != nil {
			if ae, ok := r.(astErr); ok {
				err = ae.err
				return
			}
			panic(r)
		}
	}()
	prog = v.buildProgramFromTree(tree)
	return prog, nil
}

func (v *visitor) buildProgramFromTree(ctx gen.IProgramContext) *Program {
	var statements []Statement
	for _, s := range ctx.AllStatement() {
		statements = append(statements, v.buildStatement(s))
	}
	return &Program{Statements: statements}
}

func (v *visitor) buildStatement(ctx gen.IStatementContext) Statement {
	sc := ctx.(*gen.StatementContext)
	if t := sc.TempoStmt(); t != nil {
		return v.buildTempoStmt(t)
	}
	if t := sc.TargetStmt(); t != nil {
		return v.buildTargetStmt(t)
	}
	if t := sc.ChipStmt(); t != nil {
		return v.buildChipStmt(t)
	}
	if t := sc.VoiceStmt(); t != nil {
		return v.buildVoiceStmt(t)
	}
	if t := sc.PatternDecl(); t != nil {
		return v.buildPatternDecl(t)
	}
	if t := sc.TableDecl(); t != nil {
		return v.buildTableDecl(t)
	}
	if t := sc.InstrumentBlock(); t != nil {
		return v.buildInstrumentBlock(t)
	}
	if t := sc.FilterBlock(); t != nil {
		return v.buildFilterBlock(t)
	}
	return v.buildUnknownStmt(sc.UnknownStmt())
}

var statementKeywordsList = []string{"tempo", "speed", "target", "chip", "pattern", "table", "instrument", "voice", "filter"}

func (v *visitor) buildUnknownStmt(ctx gen.IUnknownStmtContext) Statement {
	uc := ctx.(*gen.UnknownStmtContext)
	name := uc.IDENT().GetText()
	fail(ctx.GetStart().GetLine(), "unknown statement keyword '%s' (expected %v)", name, statementKeywordsList)
	return nil
}

func identText(ctx gen.IAnyIdentContext) string {
	return ctx.GetText()
}

// strInfo reads DocStart/DocEnd/unescaped value straight off the original
// Token (see Tokenize/unquoteString) rather than re-deriving them from
// the replayed token's text, which sidesteps any escape/offset fidelity
// concerns entirely.
func (v *visitor) strInfo(ctx gen.IStrContext) Token {
	sc := ctx.(*gen.StrContext)
	if dq := sc.DQSTRING(); dq != nil {
		return v.orig(dq.GetSymbol())
	}
	return v.orig(sc.BTSTRING().GetSymbol())
}

func (v *visitor) numOf(t antlr.TerminalNode) float64 {
	return v.orig(t.GetSymbol()).Num
}

func (v *visitor) buildRamp(ctx gen.IRampGeneratorContext) *TableRamp {
	rc := ctx.(*gen.RampGeneratorContext)
	line := ctx.GetStart().GetLine()
	idents := rc.AllAnyIdent()
	nums := rc.AllNUMBER()
	from := v.numOf(nums[0])
	to := v.numOf(nums[1])
	var step, over *float64
	var overInt *int
	var shape string
	// ADR 0003 D9: a trailing bare anyIdent (no NUMBER left to pair with)
	// is the shape terminator, pingpong/loop - distinguished from step/over
	// by whether a NUMBER is still available, not by position, since
	// either optional pair may be absent. idents[0] is 'ramp' itself.
	numIdx := 2
	for i := 1; i < len(idents); i++ {
		word := identText(idents[i])
		if numIdx < len(nums) {
			val := v.numOf(nums[numIdx])
			numIdx++
			switch word {
			case "step":
				sv := val
				step = &sv
			case "over":
				ov := val
				over = &ov
			default:
				fail(line, "expected 'step' or 'over', found '%s'", word)
			}
		} else {
			if word != "pingpong" && word != "loop" {
				fail(line, "expected 'pingpong' or 'loop', found '%s'", word)
			}
			shape = word
		}
	}
	if step == nil && over == nil {
		fail(line, "a 'ramp' needs at least one of 'step N' or 'over N'.")
	}
	if step != nil && *step == 0 {
		fail(line, "a 'ramp' step of 0 never reaches its target.")
	}
	if over != nil {
		oi := int(*over)
		if oi < 1 {
			fail(line, "a 'ramp' must run over at least 1 frame, got %v.", *over)
		}
		overInt = &oi
	}
	return &TableRamp{Line: line, From: from, To: to, Step: step, Over: overInt, Shape: shape}
}

// buildAbsNote mirrors astVisitor.js's buildAbsNote(): `=60` (an absolute
// MIDI note for one table frame) or `=60!4` (that cell, 4 frames long).
func (v *visitor) buildAbsNote(ctx gen.IAbsNoteContext) TableValue {
	ac := ctx.(*gen.AbsNoteContext)
	nums := ac.AllNUMBER()
	line := ctx.GetStart().GetLine()
	value := TableAbsNote{Value: v.numOf(nums[0])}
	if ac.BANG() != nil {
		count := int(v.numOf(nums[1]))
		if count < 1 {
			fail(line, "a table repeat needs a count of at least 1, got %d.", count)
		}
		return TableRepeat{Value: value, Count: count}
	}
	return value
}

// buildTableToken returns a TableValue: TableWord/TableNumber (plain
// value kinds, mirroring parser.js's plain-string/plain-number cells),
// TableRef, TableRepeat, or *TableRamp.
func (v *visitor) buildTableToken(ctx gen.ITableTokenContext) TableValue {
	tc := ctx.(*gen.TableTokenContext)
	if rg := tc.RampGenerator(); rg != nil {
		return v.buildRamp(rg)
	}
	if sig := tc.SIGIL(); sig != nil {
		txt := v.orig(sig.GetSymbol()).Str
		return TableRef{Name: txt}
	}
	if an := tc.AbsNote(); an != nil {
		return v.buildAbsNote(an)
	}
	idents := tc.AllAnyIdent()
	nums := tc.AllNUMBER()
	bang := tc.BANG()
	line := ctx.GetStart().GetLine()
	if len(idents) > 0 {
		word := identText(idents[0])
		for i := 1; i < len(idents); i++ {
			word += "+" + identText(idents[i])
		}
		if bang != nil {
			count := int(v.numOf(nums[0]))
			if count < 1 {
				fail(line, "a table repeat needs a count of at least 1, got %d.", count)
			}
			return TableRepeat{Value: TableWord{Value: word}, Count: count}
		}
		return TableWord{Value: word}
	}
	base := v.numOf(nums[0])
	if bang != nil {
		count := int(v.numOf(nums[1]))
		if count < 1 {
			fail(line, "a table repeat needs a count of at least 1, got %d.", count)
		}
		return TableRepeat{Value: TableNumber{Value: base}, Count: count}
	}
	return TableNumber{Value: base}
}

func (v *visitor) buildTableRow(ctx gen.ITableRowContext) []TableValue {
	tc := ctx.(*gen.TableRowContext)
	var out []TableValue
	for _, t := range tc.AllTableToken() {
		out = append(out, v.buildTableToken(t))
	}
	return out
}

// === tempo / pattern / table ===

func (v *visitor) buildTempoStmt(ctx gen.ITempoStmtContext) *TempoStmt {
	tc := ctx.(*gen.TempoStmtContext)
	line := ctx.GetStart().GetLine()
	tv := tc.TempoValue().(*gen.TempoValueContext)
	var value TempoValue
	if n := tv.NUMBER(); n != nil {
		value = TempoValue{Kind: "number", Number: v.numOf(n)}
	} else {
		value = TempoValue{Kind: "pattern", Source: v.strInfo(tv.Str()).Str}
	}
	// The trailing settings are generic name/number(/number) triples
	// (`beats N`, `offset N`, `meter N/D`) - the grammar accepts any
	// identifier there and this visitor is what validates the name, the
	// same soft-keyword split every instrument/voice field uses.
	var beats *int
	var offset *float64
	var meter *Meter
	for _, sc := range tc.AllTempoSetting() {
		stc := sc.(*gen.TempoSettingContext)
		name := identText(stc.AnyIdent())
		numbers := stc.AllNUMBER()
		n := v.numOf(numbers[0])
		switch name {
		case "beats":
			if len(numbers) > 1 {
				fail(line, "'beats' takes a single number, not a fraction - did you mean 'meter %v/%v'?", n, v.numOf(numbers[1]))
			}
			if n != float64(int(n)) {
				fail(line, "expected an integer, got %v", n)
			}
			bv := int(n)
			beats = &bv
		case "offset":
			if len(numbers) > 1 {
				fail(line, "'offset' takes a single number, not a fraction")
			}
			ov := n
			offset = &ov
		case "meter":
			if len(numbers) < 2 {
				fail(line, "'meter' needs a fraction, e.g. 'meter 4/4'")
			}
			d := v.numOf(numbers[1])
			if n != float64(int(n)) || d != float64(int(d)) {
				fail(line, "meter's numerator and denominator must be integers, got %v/%v", n, d)
			}
			meter = &Meter{Numerator: int(n), Denominator: int(d)}
		default:
			fail(line, "expected 'meter', 'beats' or 'offset', found '%s'", name)
		}
	}
	if beats != nil && meter != nil {
		fail(line, "a tempo statement cannot combine 'meter' and 'beats' - 'beats' is deprecated in favor of 'meter N/4', use one or the other")
	}
	// Which of the two clock spellings was written (ADR 0001's D4):
	// "bpm" for `tempo N`, "framesPerBeat" for `speed N`. Same node type
	// either way; the unit conversion and the "exactly one clock
	// statement" rule are both semantic, not grammatical (decision 11).
	unit := "bpm"
	if tc.SPEED() != nil {
		unit = "framesPerBeat"
	}
	return &TempoStmt{Line: line, Unit: unit, Value: value, Beats: beats, Offset: offset, Meter: meter}
}

// buildTargetStmt is ADR 0002's `target { ... }` block. Mirrors
// astVisitor.js's buildTargetStmt exactly, error wording included: the
// grammar only knows the shape (name value)*, this is the layer that knows
// the names and their legal combinations.
func (v *visitor) buildTargetStmt(ctx gen.ITargetStmtContext) *TargetStmt {
	tc := ctx.(*gen.TargetStmtContext)
	line := ctx.GetStart().GetLine()
	out := &TargetStmt{Line: line}
	for _, s := range tc.AllTargetSetting() {
		sc := s.(*gen.TargetSettingContext)
		idents := sc.AllAnyIdent()
		name := identText(idents[0])
		num := sc.NUMBER()
		var word string
		if len(idents) > 1 {
			word = identText(idents[1])
		}
		switch name {
		case "video":
			if word == "" {
				fail(line, "'video' takes a video standard (%s), not a number", strings.Join(VideoStandards, " or "))
			}
			if !contains(VideoStandards, word) {
				fail(line, "unknown video standard '%s' (expected %s)", word, strings.Join(VideoStandards, " or "))
			}
			if out.Video != "" {
				fail(line, "'video' given twice in one target block")
			}
			out.Video = word
		case "driveCycles":
			if num == nil {
				fail(line, "'driveCycles' takes a number of CPU cycles per play call")
			}
			n := v.numOf(num)
			if n != float64(int(n)) || n < 1 {
				fail(line, "driveCycles must be a positive integer number of CPU cycles, got %v", n)
			}
			if out.DriveCycles != nil {
				fail(line, "'driveCycles' given twice in one target block")
			}
			dv := int(n)
			out.DriveCycles = &dv
		case "multispeed":
			if num == nil {
				fail(line, "'multispeed' takes a number of play calls per video frame")
			}
			n := v.numOf(num)
			if n != float64(int(n)) || n < 1 {
				fail(line, "multispeed must be a positive integer, got %v", n)
			}
			if out.Multispeed != nil {
				fail(line, "'multispeed' given twice in one target block")
			}
			mv := int(n)
			out.Multispeed = &mv
		default:
			fail(line, "expected 'video', 'driveCycles' or 'multispeed' in a target block, found '%s'", name)
		}
	}
	if out.DriveCycles != nil && out.Multispeed != nil {
		fail(line, "'driveCycles' and 'multispeed' are two spellings of the same thing; give one, not both")
	}
	return out
}

// ChipModels lists the two SID model numbers ADR 0003 D15's `chip`
// statement accepts, mirrored by astVisitor.js's CHIP_MODELS.
var ChipModels = []int{6581, 8580}

// buildChipStmt is ADR 0003 D15's `chip 6581` / `chip 8580`: a bare
// NUMBER, no braces (unlike target), since there's exactly one setting.
func (v *visitor) buildChipStmt(ctx gen.IChipStmtContext) *ChipStmt {
	cc := ctx.(*gen.ChipStmtContext)
	line := ctx.GetStart().GetLine()
	n := v.numOf(cc.NUMBER())
	model := int(n)
	if n != float64(model) || (model != 6581 && model != 8580) {
		fail(line, "'chip' takes 6581 or 8580, got %v", n)
	}
	return &ChipStmt{Line: line, Model: model}
}

// buildArrangementElements flattens a list of arrangementElement contexts
// (whether from the top-level `arrangement` rule or from inside a nested
// `[...]` group, which is the SAME repeated-element shape one level down)
// into a []ArrangementElement, each carrying its own `*n` repeat (default
// 1). Shared by the arrangement forms of patternDecl and voiceBody - see
// the .g4 file's arrangement rule comment for why this is a statement-
// level construct rather than a mini-notation bracket.
func (v *visitor) buildArrangementElements(ctxs []gen.IArrangementElementContext) []ArrangementElement {
	out := make([]ArrangementElement, 0, len(ctxs))
	for _, ec := range ctxs {
		etc := ec.(*gen.ArrangementElementContext)
		line := etc.GetStart().GetLine()
		repeat := 1
		if n := etc.NUMBER(); n != nil {
			rv := v.numOf(n)
			if rv != float64(int(rv)) || rv <= 0 {
				fail(line, "an arrangement element's '*' repeat must be a positive integer, got %v", rv)
			}
			repeat = int(rv)
		}
		if sig := etc.SIGIL(); sig != nil {
			name := v.orig(sig.GetSymbol()).Str
			out = append(out, ArrangementElement{Line: line, Ref: &name, Repeat: repeat})
			continue
		}
		group := v.buildArrangementElements(etc.AllArrangementElement())
		out = append(out, ArrangementElement{Line: line, Group: group, Repeat: repeat})
	}
	return out
}

func (v *visitor) buildArrangement(ctx gen.IArrangementContext) []ArrangementElement {
	ac := ctx.(*gen.ArrangementContext)
	return v.buildArrangementElements(ac.AllArrangementElement())
}

func (v *visitor) buildPatternDecl(ctx gen.IPatternDeclContext) *PatternDecl {
	pc := ctx.(*gen.PatternDeclContext)
	line := ctx.GetStart().GetLine()
	name := identText(pc.AnyIdent())
	// Arrangement form (ADR 0003 D4): `pattern verse = $riffA $riffA $riffB`,
	// no quotes, no declared length - its length is always the sum of what
	// it references.
	if arr := pc.Arrangement(); arr != nil {
		return &PatternDecl{Line: line, Name: name, Elements: v.buildArrangement(arr)}
	}
	s := v.strInfo(pc.Str())
	// `|>` on a pattern declaration parses (see the grammar's patternDecl
	// rule) purely so this check can replace ANTLR's raw "mismatched
	// input" with a sentence that says where the transform belongs
	// instead - ADR 0001 decision 8 / known defect 3.
	if len(pc.AllPipelineStep()) > 0 {
		fail(line, "pipeline transforms ('|>') are not allowed on pattern declarations; add the transform on the voice that plays this pattern instead (e.g. `voice 0 $%s |> slow 4`).", name)
	}
	// The declared length in beats (`pattern bass 4 = \"...\"`), nil for a
	// proportional pattern - see the grammar's patternDecl comment and ADR
	// 0001's D1. Whether it MATCHES the pattern's own content is a
	// semantic question, checked by the JS analyze.js pass (and by
	// MiniContentBeats here in Go), not by this layer.
	var length *float64
	if n := pc.NUMBER(); n != nil {
		lv := v.numOf(n)
		if !(lv > 0) {
			fail(line, "a pattern's declared length must be greater than 0 beats, got %v", lv)
		}
		length = &lv
	}
	// `each N/D` (ADR 0003 D3): every token is one 1/D note. Always
	// written with numerator 1 in this ADR's own examples; enforced here
	// rather than in the grammar so a numerator other than 1 gets a
	// specific message instead of a generic parse failure.
	var each *Meter
	if ec := pc.PatternEach(); ec != nil {
		etc := ec.(*gen.PatternEachContext)
		eachName := identText(etc.AnyIdent())
		if eachName != "each" {
			fail(line, "expected 'each', found '%s'", eachName)
		}
		nums := etc.AllNUMBER()
		num := v.numOf(nums[0])
		den := v.numOf(nums[1])
		if num != 1 {
			fail(line, "'each' takes a fraction with numerator 1 (e.g. 'each 1/16'), got %v/%v", num, den)
		}
		if den != float64(int(den)) || den <= 0 {
			fail(line, "'each' needs a positive integer denominator, got %v", den)
		}
		each = &Meter{Numerator: int(num), Denominator: int(den)}
	}
	return &PatternDecl{Line: line, Name: name, Length: length, Each: each, Source: s.Str, DocStart: s.DocStart, DocEnd: s.DocEnd}
}

func (v *visitor) buildTableDecl(ctx gen.ITableDeclContext) *TableDecl {
	tc := ctx.(*gen.TableDeclContext)
	line := ctx.GetStart().GetLine()
	name := identText(tc.AnyIdent())
	return &TableDecl{Line: line, Name: name, Tokens: v.buildTableRow(tc.TableRow())}
}

// === instrument_block ===

func (v *visitor) buildInstrumentBlock(ctx gen.IInstrumentBlockContext) *InstrumentBlock {
	ic := ctx.(*gen.InstrumentBlockContext)
	line := ctx.GetStart().GetLine()
	name := identText(ic.AnyIdent())
	var fields []InstrumentField
	for _, f := range ic.AllInstrumentField() {
		fields = append(fields, v.buildInstrumentField(f))
	}
	return &InstrumentBlock{Line: line, Name: name, Fields: fields}
}

func tableValueAsNumber(t TableValue) (float64, bool) {
	n, ok := t.(TableNumber)
	return n.Value, ok
}

func tableValueAsWord(t TableValue) (string, bool) {
	w, ok := t.(TableWord)
	return w.Value, ok
}

// extractSpeedDelay mirrors astVisitor.js's extractSpeedDelay() (ADR 0003
// D10): peels trailing `speed N`/`delay N` pairs (either order, either or
// both absent) off the END of a table row's token list.
func extractSpeedDelay(tokens []TableValue, line int) (rest []TableValue, speed, delay *int) {
	out := append([]TableValue(nil), tokens...)
	for len(out) >= 2 {
		name, okName := tableValueAsWord(out[len(out)-2])
		val, okVal := tableValueAsNumber(out[len(out)-1])
		if !okName || !okVal || (name != "speed" && name != "delay") {
			break
		}
		if val != math.Trunc(val) {
			fail(line, "'%s' must be an integer, got %v", name, val)
		}
		if name == "speed" {
			if speed != nil {
				break
			}
			if val <= 0 {
				fail(line, "'speed' must be a positive integer, got %v", val)
			}
			iv := int(val)
			speed = &iv
		} else {
			if delay != nil {
				break
			}
			if val < 0 {
				fail(line, "'delay' must be a non-negative integer, got %v", val)
			}
			iv := int(val)
			delay = &iv
		}
		out = out[:len(out)-2]
	}
	return out, speed, delay
}

// tableVerbFieldsList / verbNamesList (ADR 0003 D13) mirror
// astVisitor.js's TABLE_VERB_FIELDS/VERB_NAMES exactly.
var tableVerbFieldsList = []string{"wave", "arp", "pw", "filterTable", "test", "gate"}
var verbNamesList = []string{"set", "add", "wait", "loop", "stop"}

func containsStr(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func (v *visitor) buildTableVerbLine(ctx gen.ITableVerbLineContext) TableVerb {
	lc := ctx.(*gen.TableVerbLineContext)
	line := ctx.GetStart().GetLine()
	verbName := identText(lc.AnyIdent())
	if !containsStr(verbNamesList, verbName) {
		fail(line, "unknown table verb '%s' (expected %v)", verbName, verbNamesList)
	}
	var tokens []TableValue
	if lc.TableRow() != nil {
		tokens = v.buildTableRow(lc.TableRow())
	}
	switch verbName {
	case "set":
		if len(tokens) != 1 {
			fail(line, "'set' expects 1 number")
		}
		val, ok := tableValueAsNumber(tokens[0])
		if !ok {
			fail(line, "'set' expects 1 number")
		}
		return TableVerb{Line: line, Verb: verbName, Value: val}
	case "add":
		if len(tokens) != 1 && len(tokens) != 3 {
			fail(line, "'add' expects a number, optionally followed by 'repeat N'")
		}
		val, ok := tableValueAsNumber(tokens[0])
		if !ok {
			fail(line, "'add' expects a number")
		}
		repeat := 1
		if len(tokens) == 3 {
			word, okWord := tableValueAsWord(tokens[1])
			n, okNum := tableValueAsNumber(tokens[2])
			if !okWord || word != "repeat" || !okNum {
				fail(line, "'add' expects 'repeat N' after its value")
			}
			if n != math.Trunc(n) || n < 1 {
				fail(line, "'repeat' must be a positive integer, got %v", n)
			}
			repeat = int(n)
		}
		return TableVerb{Line: line, Verb: verbName, Value: val, Repeat: repeat}
	case "wait":
		if len(tokens) != 1 {
			fail(line, "'wait' expects 1 number")
		}
		n, ok := tableValueAsNumber(tokens[0])
		if !ok || n != math.Trunc(n) || n < 1 {
			fail(line, "'wait' must be a positive integer, got %v", tokens[0])
		}
		return TableVerb{Line: line, Verb: verbName, Count: int(n)}
	case "loop", "stop":
		if len(tokens) != 0 {
			fail(line, "'%s' takes no arguments", verbName)
		}
		return TableVerb{Line: line, Verb: verbName}
	default:
		fail(line, "unknown table verb '%s'", verbName)
		return TableVerb{}
	}
}

func (v *visitor) buildInstrumentField(ctx gen.IInstrumentFieldContext) InstrumentField {
	fc := ctx.(*gen.InstrumentFieldContext)
	line := ctx.GetStart().GetLine()
	name := identText(fc.AnyIdent())
	if fc.TableRow() == nil {
		// ADR 0003 D13: the verb-block form - a nil TableRow is exactly how
		// the two alternatives of instrumentField are told apart here,
		// mirroring astVisitor.js's own `!ctx.tableRow()` check.
		if !containsStr(tableVerbFieldsList, name) {
			fail(line, "'%s' can't use the verb-block form (only %v can)", name, tableVerbFieldsList)
		}
		lines := fc.AllTableVerbLine()
		verbs := make([]TableVerb, len(lines))
		for i, l := range lines {
			verbs[i] = v.buildTableVerbLine(l)
		}
		switch name {
		case "wave":
			return &WaveField{Line: line, Verbs: verbs}
		case "arp":
			return &ArpField{Line: line, Verbs: verbs}
		case "pw":
			return &PwField{Line: line, Verbs: verbs}
		case "filterTable":
			return &FilterTableField{Line: line, Verbs: verbs}
		case "test":
			return &TestField{Line: line, Verbs: verbs}
		case "gate":
			return &InstGateField{Line: line, Verbs: verbs}
		}
	}
	tokens := v.buildTableRow(fc.TableRow())
	// ADR 0003 D10: speed/delay peeled off the tail of these six fields'
	// token lists specifically (mirrors astVisitor.js's buildInstrumentField).
	var speed, delay *int
	switch name {
	case "wave", "arp", "pw", "filterTable", "test":
		tokens, speed, delay = extractSpeedDelay(tokens, line)
	case "gate":
		if len(tokens) > 1 {
			tokens, speed, delay = extractSpeedDelay(tokens, line)
		}
	}
	switch name {
	case "wave":
		return &WaveField{Line: line, Tokens: tokens, Speed: speed, Delay: delay}
	case "arp":
		return &ArpField{Line: line, Tokens: tokens, Speed: speed, Delay: delay}
	case "pw":
		return &PwField{Line: line, Tokens: tokens, Speed: speed, Delay: delay}
	case "filterTable":
		return &FilterTableField{Line: line, Tokens: tokens, Speed: speed, Delay: delay}
	case "test":
		return &TestField{Line: line, Tokens: tokens, Speed: speed, Delay: delay}
	case "adsr":
		if len(tokens) != 4 {
			fail(line, "adsr expects 4 numbers")
		}
		a, ok1 := tableValueAsNumber(tokens[0])
		d, ok2 := tableValueAsNumber(tokens[1])
		s, ok3 := tableValueAsNumber(tokens[2])
		r, ok4 := tableValueAsNumber(tokens[3])
		if !ok1 || !ok2 || !ok3 || !ok4 {
			fail(line, "adsr expects 4 numbers")
		}
		return &AdsrField{Line: line, A: a, D: d, S: s, R: r}
	case "vib":
		if len(tokens) != 2 {
			fail(line, "vib expects 2 numbers")
		}
		depth, ok1 := tableValueAsNumber(tokens[0])
		rate, ok2 := tableValueAsNumber(tokens[1])
		if !ok1 || !ok2 {
			fail(line, "vib expects 2 numbers")
		}
		return &InstVibField{Line: line, Depth: depth, Rate: rate}
	case "vibTri":
		if len(tokens) != 3 {
			fail(line, "vibTri expects 3 values (CMP SPEED DELAY)")
		}
		cmp, okCmp := tableValueAsNumber(tokens[0])
		delay, okDelay := tableValueAsNumber(tokens[2])
		if !okCmp || !okDelay {
			fail(line, "vibTri expects numbers for CMP and DELAY")
		}
		field := &InstVibTriField{Line: line, Cmp: cmp, Delay: delay}
		if speed, ok := tableValueAsNumber(tokens[1]); ok {
			field.Speed = speed
		} else if name, ok := tableValueAsWord(tokens[1]); ok && vibTriShiftRe.MatchString(name) {
			n, _ := strconv.Atoi(name[1:])
			field.SpeedShift = &n
		} else {
			fail(line, "vibTri's SPEED must be a number or a shift-derived 'sN'")
		}
		return field
	case "slide":
		if len(tokens) != 1 {
			fail(line, "slide expects 1 number")
		}
		t, ok := tableValueAsNumber(tokens[0])
		if !ok {
			fail(line, "slide expects 1 number")
		}
		return &InstSlideField{Line: line, Time: t}
	case "loop":
		return v.buildLoopField(line, tokens)
	case "gate":
		if len(tokens) == 1 {
			if word, ok := tableValueAsWord(tokens[0]); ok {
				if word != "legato" && word != "retrigger" {
					fail(line, "'gate' expects a number, or one of %v, found '%s'", gateModesList, word)
				}
				return &InstGateField{Line: line, Mode: word}
			}
			return &InstGateField{Line: line, Value: tokens[0]}
		}
		return &InstGateField{Line: line, Tokens: tokens, Speed: speed, Delay: delay}
	case "hardRestart": {
		// ADR 0003 D11: the readable spelling of 'gate -N' - compiles to
		// exactly the same AST shape a numeric 'gate -N' field would.
		if len(tokens) != 1 {
			fail(line, "'hardRestart' expects 1 number of frames")
		}
		n, ok := tableValueAsNumber(tokens[0])
		if !ok || n != math.Trunc(n) || n <= 0 {
			fail(line, "'hardRestart' must be a positive integer number of frames, got %v", tokens[0])
		}
		return &InstGateField{Line: line, Value: TableNumber{Value: -n}}
	}
	default:
		fail(line, "unknown instrument field '%s' (expected %v)", name, instrumentFieldsList)
		return nil
	}
}

func (v *visitor) buildLoopField(line int, tokens []TableValue) *LoopField {
	if len(tokens) == 1 {
		if n, ok := tableValueAsNumber(tokens[0]); ok {
			if n != float64(int(n)) {
				fail(line, "expected an integer, got %v", n)
			}
			iv := int(n)
			return &LoopField{Line: line, Value: &iv}
		}
	}
	var pairs []LoopPair
	for i := 0; i+1 < len(tokens); i += 2 {
		table, ok1 := tableValueAsWord(tokens[i])
		val, ok2 := tableValueAsNumber(tokens[i+1])
		if !ok1 || !ok2 {
			fail(line, "malformed loop pairs")
		}
		if !contains(loopTableNames, table) {
			fail(line, "unknown loop table '%s' (expected %v)", table, loopTableNames)
		}
		pairs = append(pairs, LoopPair{Table: table, Value: int(val)})
	}
	if len(pairs) == 0 {
		fail(line, "expected an integer, or one or more TABLE INT pairs (%v)", loopTableNames)
	}
	return &LoopField{Line: line, Pairs: pairs}
}

// === voice_stmt ===

func (v *visitor) buildVoiceStmt(ctx gen.IVoiceStmtContext) *VoiceStmt {
	vc := ctx.(*gen.VoiceStmtContext)
	line := ctx.GetStart().GetLine()
	index := int(v.numOf(vc.NUMBER()))
	var instrumentName *string
	if ai := vc.AnyIdent(); ai != nil {
		n := identText(ai)
		instrumentName = &n
	}
	body := vc.VoiceBody().(*gen.VoiceBodyContext)
	if body.Str() != nil || body.SIGIL() != nil {
		notes := v.buildNotesValueFromBody(body)
		tail := body.VoiceShortTail().(*gen.VoiceShortTailContext)
		var inlineAttrs []InlineAttr
		for _, a := range tail.AllInlineAttr() {
			inlineAttrs = append(inlineAttrs, v.buildInlineAttr(a))
		}
		var transforms []Transform
		for _, p := range tail.AllPipelineStep() {
			pc := p.(*gen.PipelineStepContext)
			transforms = append(transforms, v.buildTransform(pc.Transform()))
		}
		return &VoiceStmt{Line: line, Index: index, InstrumentName: instrumentName, Mode: "short", Notes: notes, InlineAttrs: inlineAttrs, Transforms: transforms}
	}
	if arr := body.Arrangement(); arr != nil {
		elements := v.buildArrangement(arr)
		tail := body.VoiceShortTail().(*gen.VoiceShortTailContext)
		var inlineAttrs []InlineAttr
		for _, a := range tail.AllInlineAttr() {
			inlineAttrs = append(inlineAttrs, v.buildInlineAttr(a))
		}
		var transforms []Transform
		for _, p := range tail.AllPipelineStep() {
			pc := p.(*gen.PipelineStepContext)
			transforms = append(transforms, v.buildTransform(pc.Transform()))
		}
		return &VoiceStmt{Line: line, Index: index, InstrumentName: instrumentName, Mode: "short", Arrangement: elements, InlineAttrs: inlineAttrs, Transforms: transforms}
	}
	var fields []VoiceField
	for _, f := range body.AllVoiceField() {
		fields = append(fields, v.buildVoiceField(f))
	}
	return &VoiceStmt{Line: line, Index: index, InstrumentName: instrumentName, Mode: "extended", Fields: fields}
}

func (v *visitor) buildNotesValueFromBody(body *gen.VoiceBodyContext) *NotesValue {
	if s := body.Str(); s != nil {
		info := v.strInfo(s)
		return &NotesValue{IsSigil: false, Source: info.Str, DocStart: info.DocStart, DocEnd: info.DocEnd}
	}
	sig := v.orig(body.SIGIL().GetSymbol())
	return &NotesValue{IsSigil: true, SigilName: sig.Str}
}

func (v *visitor) buildInlineAttr(ctx gen.IInlineAttrContext) InlineAttr {
	ac := ctx.(*gen.InlineAttrContext)
	name := identText(ac.AnyIdent())
	if s := ac.Str(); s != nil {
		info := v.strInfo(s)
		return InlineAttr{Name: name, IsPattern: true, Source: info.Str, DocStart: info.DocStart, DocEnd: info.DocEnd}
	}
	vl := ac.ValueList().(*gen.ValueListContext)
	var values []TableValue
	for _, t := range vl.AllTableToken() {
		values = append(values, v.buildTableToken(t))
	}
	return InlineAttr{Name: name, Values: values}
}

func (v *visitor) buildVoiceField(ctx gen.IVoiceFieldContext) VoiceField {
	fc := ctx.(*gen.VoiceFieldContext)
	line := ctx.GetStart().GetLine()
	if p := fc.PipelineStep(); p != nil {
		pc := p.(*gen.PipelineStepContext)
		return &PipelineField{Line: line, Transforms: []Transform{v.buildTransform(pc.Transform())}}
	}
	name := identText(fc.AnyIdent())
	vv := fc.VoiceFieldValue().(*gen.VoiceFieldValueContext)
	switch name {
	case "notes":
		if vv.Str() != nil {
			info := v.strInfo(vv.Str())
			return &NotesField{Line: line, NotesValue: NotesValue{IsSigil: false, Source: info.Str, DocStart: info.DocStart, DocEnd: info.DocEnd}}
		}
		if vv.SIGIL() != nil {
			sig := v.orig(vv.SIGIL().GetSymbol())
			return &NotesField{Line: line, NotesValue: NotesValue{IsSigil: true, SigilName: sig.Str}}
		}
		fail(line, "'notes' expects a quoted note pattern or a $patternName reference")
	case "insts":
		if vv.Str() == nil {
			fail(line, "'insts' expects a quoted mini-notation pattern")
		}
		info := v.strInfo(vv.Str())
		return &InstsField{Line: line, Source: info.Str, DocStart: info.DocStart, DocEnd: info.DocEnd}
	case "fx":
		// The per-row effect-command lane - shaped exactly like `insts`
		// (see src/dsl/fx.js for why a parallel mini-notation string).
		if vv.Str() == nil {
			fail(line, "'fx' expects a quoted mini-notation pattern")
		}
		fxInfo := v.strInfo(vv.Str())
		return &FxField{Line: line, Source: fxInfo.Str, DocStart: fxInfo.DocStart, DocEnd: fxInfo.DocEnd}
	case "wave":
		if vv.Str() == nil {
			fail(line, "'wave' expects a quoted mini-notation pattern")
		}
		info := v.strInfo(vv.Str())
		return &VoiceWaveField{Line: line, Source: info.Str, DocStart: info.DocStart, DocEnd: info.DocEnd}
	case "reg":
		var target RegTarget
		if hex := vv.HEX(); hex != nil {
			target = RegTarget{Kind: "hex", Hex: v.orig(hex.GetSymbol()).HexVal}
		} else if ai := vv.AnyIdent(); ai != nil {
			target = RegTarget{Kind: "ident", Ident: identText(ai)}
		} else {
			fail(line, "expected a register name or a 0xNN hex address")
		}
		if vv.Str() == nil {
			fail(line, "'reg' expects a quoted register-value mini-notation pattern")
		}
		return &RegField{Line: line, Target: target, Source: v.strInfo(vv.Str()).Str}
	case "vib":
		nums := vv.AllNUMBER()
		if len(nums) != 2 {
			fail(line, "'vib' expects 2 numbers")
		}
		return &VoiceVibField{Line: line, Depth: v.numOf(nums[0]), Rate: v.numOf(nums[1])}
	case "slide":
		nums := vv.AllNUMBER()
		if len(nums) != 1 {
			fail(line, "'slide' expects 1 number")
		}
		return &VoiceSlideField{Line: line, Time: v.numOf(nums[0])}
	case "gate":
		nums := vv.AllNUMBER()
		if len(nums) != 1 {
			fail(line, "'gate' expects 1 number")
		}
		return &VoiceGateField{Line: line, Value: v.numOf(nums[0])}
	case "sync":
		return &SyncField{Line: line, Value: v.buildBoolFieldValue(vv, line)}
	case "ring":
		return &RingField{Line: line, Value: v.buildBoolFieldValue(vv, line)}
	default:
		fail(line, "unknown voice field '%s' (expected %v, or a '|>' pipeline)", name, voiceFieldsList)
	}
	return nil
}

func (v *visitor) buildBoolFieldValue(vv *gen.VoiceFieldValueContext, line int) BoolFieldValue {
	if vv.Str() != nil {
		return BoolFieldValue{Kind: "pattern", Source: v.strInfo(vv.Str()).Str}
	}
	nums := vv.AllNUMBER()
	if len(nums) == 1 {
		return BoolFieldValue{Kind: "number", Number: v.numOf(nums[0])}
	}
	fail(line, "expected a number (0/1) or a quoted mini-notation pattern")
	return BoolFieldValue{}
}

// === pipeline / transform ===

func (v *visitor) buildTransform(ctx gen.ITransformContext) Transform {
	tc := ctx.(*gen.TransformContext)
	line := ctx.GetStart().GetLine()
	name := identText(tc.AnyIdent())
	argCtx, _ := tc.TransformArg().(*gen.TransformArgContext)
	refCtx := tc.TransformRef()

	var arg *TransformArg
	if argCtx != nil {
		a := v.buildTransformArg(argCtx)
		arg = &a
	}
	var fn Transform
	if refCtx != nil {
		fn = v.buildTransformRef(refCtx)
	}

	needArg := func() TransformArg {
		if arg == nil {
			fail(line, "'%s' expects an argument", name)
		}
		return *arg
	}
	needIntArg := func() TransformArg {
		a := needArg()
		if !a.HasSource && a.Value != float64(int(a.Value)) {
			fail(line, "expected an integer, got %v", a.Value)
		}
		return a
	}
	needFn := func() Transform {
		if fn == nil {
			fail(line, "'%s' expects a transform (or parenthesized transform)", name)
		}
		return fn
	}
	needStringArg := func(label string) string {
		if argCtx == nil || argCtx.Str() == nil {
			fail(line, "expected %s", label)
		}
		return v.strInfo(argCtx.Str()).Str
	}

	switch name {
	case "rev":
		return &RevTransform{Line: line}
	case "fast":
		return &FastTransform{Line: line, N: needArg()}
	case "slow":
		return &SlowTransform{Line: line, N: needArg()}
	case "every":
		return &EveryTransform{Line: line, N: needIntArg(), Fn: needFn()}
	case "iter":
		return &IterTransform{Line: line, N: needIntArg()}
	case "ply":
		return &PlyTransform{Line: line, N: needIntArg()}
	case "off":
		return &OffTransform{Line: line, Time: needArg(), Fn: needFn()}
	case "jux":
		return &JuxTransform{Line: line, Fn: needFn()}
	case "degradeBy":
		return &DegradeByTransform{Line: line, P: needArg()}
	case "sometimesBy":
		return &SometimesByTransform{Line: line, P: needArg(), Fn: needFn()}
	case "add":
		return &AddTransform{Line: line, N: needArg()}
	case "struct":
		return &StructTransform{Line: line, Source: needStringArg("a quoted boolean rhythm pattern")}
	case "mask":
		return &MaskTransform{Line: line, Source: needStringArg("a quoted boolean mask pattern")}
	case "chunk":
		return &ChunkTransform{Line: line, N: needIntArg(), Fn: needFn()}
	case "palindrome":
		return &PalindromeTransform{Line: line}
	case "linger":
		return &LingerTransform{Line: line, Frac: needArg()}
	case "rot":
		return &RotTransform{Line: line, N: needIntArg()}
	// `someBarsBy` is 3.0.0's name for `someCyclesBy` (ADR 0001's D3:
	// transforms count bars). Both spellings build the SAME node type.
	case "someBarsBy", "someCyclesBy":
		return &SomeCyclesByTransform{Line: line, P: needArg(), Fn: needFn()}
	case "superimpose":
		return &SuperimposeTransform{Line: line, Fn: needFn()}
	case "scale":
		return &ScaleTransform{Line: line, Source: needStringArg("a quoted scale spec")}
	default:
		fail(line, "unknown transform '%s' (expected %v)", name, transformsList)
	}
	return nil
}

func (v *visitor) buildTransformArg(ctx *gen.TransformArgContext) TransformArg {
	if s := ctx.Str(); s != nil {
		return TransformArg{HasSource: true, Source: v.strInfo(s).Str}
	}
	return TransformArg{Value: v.numOf(ctx.NUMBER())}
}

func (v *visitor) buildTransformRef(ctx gen.ITransformRefContext) Transform {
	rc := ctx.(*gen.TransformRefContext)
	return v.buildTransform(rc.Transform())
}

// === filter_block ===

func (v *visitor) buildFilterBlock(ctx gen.IFilterBlockContext) *FilterBlock {
	fc := ctx.(*gen.FilterBlockContext)
	line := ctx.GetStart().GetLine()
	var fields []FilterField
	for _, f := range fc.AllFilterField() {
		fields = append(fields, v.buildFilterField(f))
	}
	return &FilterBlock{Line: line, Fields: fields}
}

func (v *visitor) buildFilterField(ctx gen.IFilterFieldContext) FilterField {
	fc := ctx.(*gen.FilterFieldContext)
	line := ctx.GetStart().GetLine()
	name := identText(fc.AnyIdent())
	fv := fc.FilterFieldValue().(*gen.FilterFieldValueContext)
	switch name {
	case "cutoff":
		return &CutoffField{Line: line, Expr: v.buildFilterExpr(fv, "cutoff expects a signal expression or 'switch bar' block")}
	case "resonance":
		if fv.NumberList() == nil && len(fv.AllAnyIdent()) == 0 {
			fail(line, "resonance expects a number, a signal expression, or 'switch bar' block")
		}
		return &ResonanceField{Line: line, Expr: v.buildFilterExpr(fv, "resonance expects a number, a signal expression, or 'switch bar' block")}
	case "mode":
		if len(fv.AllAnyIdent()) == 0 {
			fail(line, "mode expects an identifier or 'switch bar' block")
		}
		return &ModeField{Line: line, Expr: v.buildFilterExpr(fv, "mode expects an identifier or 'switch bar' block")}
	case "route":
		if fv.NumberList() == nil && len(fv.AllAnyIdent()) == 0 {
			fail(line, "route expects one or more integers or 'switch bar' block")
		}
		return &RouteField{Line: line, Expr: v.buildFilterExpr(fv, "route expects one or more integers or 'switch bar' block")}
	default:
		fail(line, "unknown filter field '%s' (expected %v)", name, filterFieldsList)
	}
	return nil
}

// buildFilterExpr handles all three filterFieldValue shapes uniformly
// (switch-cycle block / identifier-led signal expression / bare number
// list) - which of the three is legal for a given field name is enforced
// by buildFilterField's callers above, not here.
func (v *visitor) buildFilterExpr(fv *gen.FilterFieldValueContext, errMsg string) FilterExpr {
	ai := fv.AllAnyIdent()
	line := fv.GetStart().GetLine()
	if len(ai) == 2 {
		if identText(ai[0]) != "switch" {
			fail(line, "expected 'switch'")
		}
		// `switch bar` is the 3.0.0 spelling; `switch cycle` still
		// parses so no existing song breaks (ADR 0001's D3).
		if w := identText(ai[1]); w != "bar" && w != "cycle" {
			fail(line, "expected 'bar' (or the pre-3.0.0 spelling 'cycle')")
		}
		var cases []SwitchCase
		for _, c := range fv.AllSwitchCase() {
			cases = append(cases, v.buildSwitchCase(c))
		}
		return &SwitchBlock{Line: line, Cases: cases}
	}
	if len(ai) == 1 {
		return v.buildSignalExprFromParts(line, ai[0], fv.SignalRangeOpt(), fv.AllSignalSpeedMod())
	}
	if nl := fv.NumberList(); nl != nil {
		return v.buildNumberListExpr(nl.(*gen.NumberListContext))
	}
	fail(line, "%s", errMsg)
	return nil
}

func (v *visitor) buildSignalExprFromParts(line int, signalIdent gen.IAnyIdentContext, rangeCtx gen.ISignalRangeOptContext, speedCtxs []gen.ISignalSpeedModContext) *SignalExpr {
	signal := identText(signalIdent)
	var rng *SignalRange
	if rc, ok := rangeCtx.(*gen.SignalRangeOptContext); ok && rc != nil {
		nums := rc.AllNUMBER()
		if len(nums) == 2 {
			rng = &SignalRange{Lo: v.numOf(nums[0]), Hi: v.numOf(nums[1])}
		}
	}
	var speeds []SpeedMod
	for _, s := range speedCtxs {
		sc := s.(*gen.SignalSpeedModContext)
		num := sc.NUMBER()
		var op string
		switch {
		case sc.STAR() != nil:
			op = "*"
		case sc.SLASH() != nil:
			op = "/"
		default:
			op = ">>"
		}
		speeds = append(speeds, SpeedMod{Op: op, Value: v.numOf(num)})
	}
	return &SignalExpr{Line: line, Signal: signal, Range: rng, Speeds: speeds}
}

func (v *visitor) buildSwitchCase(ctx gen.ISwitchCaseContext) SwitchCase {
	cc := ctx.(*gen.SwitchCaseContext)
	line := ctx.GetStart().GetLine()
	sc := cc.SwitchCond().(*gen.SwitchCondContext)
	var cond SwitchCond
	switch {
	case sc.LT() != nil:
		cond = SwitchCond{Op: "<", Value: v.numOf(sc.NUMBER())}
	case sc.GE() != nil:
		cond = SwitchCond{Op: ">=", Value: v.numOf(sc.NUMBER())}
	case sc.AnyIdent() != nil:
		if identText(sc.AnyIdent()) != "else" {
			fail(line, "expected a switch case ('< N', '>= N', or 'else')")
		}
		cond = SwitchCond{Op: "else"}
	default:
		fail(line, "expected a switch case ('< N', '>= N', or 'else')")
	}
	var value SwitchCaseValue
	if nl := cc.NumberList(); nl != nil {
		value = v.buildNumberListExpr(nl.(*gen.NumberListContext))
	} else {
		value = v.buildSignalExprFromParts(line, cc.AnyIdent(), cc.SignalRangeOpt(), cc.AllSignalSpeedMod())
	}
	return SwitchCase{Line: line, Cond: cond, Value: value}
}

func (v *visitor) buildNumberListExpr(nlc *gen.NumberListContext) *NumberListExpr {
	line := nlc.GetStart().GetLine()
	var values []float64
	for _, n := range nlc.AllNUMBER() {
		values = append(values, v.numOf(n))
	}
	return &NumberListExpr{Line: line, Values: values}
}
