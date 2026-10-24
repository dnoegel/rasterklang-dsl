package dsl

// AST node types for DSL v2 (docs/dsl-grammar-spec.md's "Grammar
// (EBNF-ish)" section), mirroring parser.js's plain-object node shapes
// 1:1 by field so the spec doc, parser.js, and this file stay easy to
// cross-reference. Every node carries the 1-based source Line it started
// on (see parser.js's header comment on why).
//
// Where the JS produces a dynamically-shaped node (e.g. a table cell that
// is sometimes a bare string, sometimes a {ref}/{repeat}/ramp object),
// this file uses a small sealed interface (a marker method) with one
// concrete struct per shape, instead of `interface{}`/type-switch soup.

// Program is the parse root: one statement per top-level line/block.
type Program struct {
	Statements []Statement
}

// Statement is the sealed interface every top-level statement kind
// implements: TempoStmt, PatternDecl, TableDecl, InstrumentBlock,
// VoiceStmt, FilterBlock.
type Statement interface{ isStatement() }

func (*TempoStmt) isStatement()       {}
func (*TargetStmt) isStatement()      {}
func (*ChipStmt) isStatement()        {}
func (*PatternDecl) isStatement()     {}
func (*TableDecl) isStatement()       {}
func (*InstrumentBlock) isStatement() {}
func (*VoiceStmt) isStatement()       {}
func (*FilterBlock) isStatement()     {}

// === tempo_stmt ==============================================================

type TempoStmt struct {
	Line int
	// Unit is which of the two clock spellings was written: "bpm" for
	// `tempo N` (musical), "framesPerBeat" for `speed N` (exact frames per
	// beat, what GoatTracker natively counts) - ADR 0001's D4. One node
	// type covers both; converting between them is the semantic layer's
	// job, as is enforcing that a song has exactly one clock statement.
	Unit  string
	Value TempoValue
	// Beats is nil when the 'beats' clause was omitted.
	Beats *int
	// Offset is nil when the 'offset' clause was omitted: the whole song's
	// pattern clock shifted by this many BEATS (fractional and negative both
	// allowed), positive delaying the song.
	Offset *float64
	// Meter is nil when the 'meter N/D' clause was omitted (ADR 0003 D2).
	// Mutually exclusive with Beats: the semantic layer rejects both on one
	// statement.
	Meter *Meter
}

// Meter is `meter N/D`: a bar is Numerator beats, and one beat is a
// 1/Denominator note.
type Meter struct {
	Numerator   int
	Denominator int
}

// TempoValue is a NUMBER (Kind "number") or a quoted mini-notation
// pattern (Kind "pattern").
type TempoValue struct {
	Kind   string // "number" | "pattern"
	Number float64
	Source string
}

// === target_stmt =============================================================

// TargetStmt is ADR 0002's `target { video PAL multispeed 3 }` /
// `target { driveCycles 12801 }` block: the song's DRIVE rate, i.e. how
// many CPU cycles pass between two play calls. Every field is nil/"" when
// the corresponding setting was omitted, and a song with no target block at
// all means DefaultDriveCycles (PAL, exactly 50.00 Hz) - DSL v2's only
// pre-0002 behaviour. Use ResolveTarget to expand this into concrete
// cycles/Hz; DriveCycles and Multispeed are never both set (the AST
// transform rejects that pair, since they are two spellings of one number).
type TargetStmt struct {
	Line        int
	Video       string // "PAL" | "NTSC" | "" (omitted)
	DriveCycles *int
	Multispeed  *int
}

// === chip_stmt ================================================================

// ChipStmt is ADR 0003 D15's `chip 6581` / `chip 8580`: which SID model
// this song was written for/against. Optional, at most one per song; a
// song with none gets the player's default model. This only lets a song
// RECORD its choice instead of that choice living in transient player
// state.
type ChipStmt struct {
	Line  int
	Model int // 6581 or 8580, validated by the semantic layer
}

// === pattern_decl ============================================================

type PatternDecl struct {
	Line int
	Name string
	// Length is the declared length in BEATS (`pattern bass 4 = "..."`),
	// nil for a proportional pattern with no declared length. A declared
	// length is the opt-in to metric behaviour (ADR 0001's D1): an
	// unweighted token is then one beat, and the summed content weight
	// must equal this number (see MiniContentBeats).
	Length *float64
	// Each is nil unless the declaration used `each N/D` (ADR 0003 D3): every
	// token is one 1/D note, so the pattern's length is derived rather than
	// declared. Legal alongside Length, which is then checked against the
	// derived total like any other declared length.
	Each *Meter
	// Elements is non-nil for an ARRANGEMENT declaration (ADR 0003 D4):
	// `pattern verse = $riffA $riffA $riffB`, no quotes. Mutually exclusive
	// with Source/Length/Each, which stay zero-valued in that case - an
	// arrangement's length is always the sum of what it references, never
	// declared or derived from mini-notation content.
	Elements []ArrangementElement
	Source   string
	// DocStart/DocEnd: this declaration's own absolute source-document
	// offsets (see lexer.go's Token doc comment) - threaded through so a
	// later semantic layer can anchor a $name-referencing voice's
	// source-highlighting directly at this declaration.
	DocStart int
	DocEnd   int
}

// === table_decl ===============================================================

type TableDecl struct {
	Line   int
	Name   string
	Tokens []TableValue
}

// TableValue is one cell of a table_row: a bare word/waveform-combo
// (TableWord), a NUMBER (TableNumber), a $NAME reference to a table decl
// (TableRef), an absolute-MIDI-note cell (TableAbsNote), a VALUE!N repeat
// (TableRepeat), or a ramp generator (TableRamp).
type TableValue interface{ isTableValue() }

func (TableWord) isTableValue()    {}
func (TableNumber) isTableValue()  {}
func (TableRef) isTableValue()     {}
func (TableAbsNote) isTableValue() {}
func (TableRepeat) isTableValue()  {}
func (*TableRamp) isTableValue()   {}

// TableWord is a bare identifier, or several joined with '+'
// (`noise+pulse+saw`) - kept as the joined source text, not resolved to a
// bitmask (waveform semantics live outside the grammar layer).
type TableWord struct{ Value string }

type TableNumber struct{ Value float64 }

// TableRef is a `$NAME` splice of a `table` declaration.
type TableRef struct{ Name string }

// TableAbsNote is `=60` - an ABSOLUTE MIDI note number for one table
// frame, as opposed to a plain TableNumber cell in an `arp` row (a
// RELATIVE semitone offset from the sounding note). See the grammar's
// absNote rule comment for why it is spelled as a number, not a name.
type TableAbsNote struct{ Value float64 }

// TableRepeat is `VALUE!N` - N copies of Value (itself a TableWord,
// TableNumber or TableAbsNote; never a ref/ramp/another repeat).
type TableRepeat struct {
	Value TableValue
	Count int
}

// TableRamp is `ramp FROM..TO (step S)? (over N)?` - a linear per-frame
// sweep. At least one of Step/Over is always set (the parser rejects
// neither being present).
type TableRamp struct {
	Line int
	From float64
	To   float64
	Step *float64
	Over *int
	// Shape is "" (saturate at the target), "pingpong" or "loop" (ADR
	// 0003 D9) - a terminator on what happens once the ramp reaches its
	// target, expanded into a longer table + an implied per-table loop
	// point at interpret time, not a new engine primitive.
	Shape string
}

// === instrument_block =========================================================

type InstrumentBlock struct {
	Line   int
	Name   string
	Fields []InstrumentField
}

// InstrumentField is the sealed interface for every instrument_field
// kind: WaveField, ArpField, PwField, FilterTableField, TestField,
// AdsrField, LoopField, InstVibField, InstVibTriField, InstSlideField,
// InstGateField.
type InstrumentField interface{ isInstrumentField() }

func (*WaveField) isInstrumentField()        {}
func (*ArpField) isInstrumentField()         {}
func (*PwField) isInstrumentField()          {}
func (*FilterTableField) isInstrumentField() {}
func (*TestField) isInstrumentField()        {}
func (*AdsrField) isInstrumentField()        {}
func (*LoopField) isInstrumentField()        {}
func (*InstVibField) isInstrumentField()     {}
func (*InstVibTriField) isInstrumentField()  {}
func (*InstSlideField) isInstrumentField()   {}
func (*InstGateField) isInstrumentField()    {}

// Speed/Delay (ADR 0003 D10, on WaveField/ArpField/PwField/
// FilterTableField/TestField/InstGateField alike): nil unless this table
// row ended with an explicit `speed N`/`delay N` (astVisitor.js's
// extractSpeedDelay mirror, ExtractSpeedDelay below).
// Verbs (ADR 0003 D13, on WaveField/ArpField/PwField/FilterTableField/
// TestField/InstGateField alike): the verb-block alternative to Tokens -
// exactly one of Tokens/Verbs is set, never both, distinguished the same
// way astVisitor.js's buildInstrumentField does (grammar's tableRow being
// absent means the verb-block alternative was taken).
type WaveField struct {
	Line         int
	Tokens       []TableValue
	Speed, Delay *int
	Verbs        []TableVerb
}
type ArpField struct {
	Line         int
	Tokens       []TableValue
	Speed, Delay *int
	Verbs        []TableVerb
}
type PwField struct {
	Line         int
	Tokens       []TableValue
	Speed, Delay *int
	Verbs        []TableVerb
}
type FilterTableField struct {
	Line         int
	Tokens       []TableValue
	Speed, Delay *int
	Verbs        []TableVerb
}
type TestField struct {
	Line         int
	Tokens       []TableValue
	Speed, Delay *int
	Verbs        []TableVerb
}

// TableVerb is one line of a D13 verb block: Verb is "set"/"add"/"wait"/
// "loop"/"stop". Value is set for set/add; Repeat for add (default 1);
// Count for wait; loop/stop take no arguments.
type TableVerb struct {
	Line   int
	Verb   string
	Value  float64
	Repeat int
	Count  int
}
type AdsrField struct {
	Line       int
	A, D, S, R float64
}

// LoopField is either a single global loop index (Value non-nil, Pairs
// nil) or one-or-more per-table TABLE INT pairs (Pairs non-nil, Value
// nil) - see parser.js's parseLoopField comment for the GoatTracker/
// SID-Wizard rationale.
type LoopField struct {
	Line  int
	Value *int
	Pairs []LoopPair
}
type LoopPair struct {
	Table string
	Value int
}

type InstVibField struct {
	Line        int
	Depth, Rate float64
}

// InstVibTriField is `vibTri CMP SPEED DELAY` - GoatTracker's own triangle
// vibrato (see the JS fieldDocs entry). Exactly one of Speed/SpeedShift is
// set: SpeedShift carries the note-relative `sN` spelling, which cannot be
// resolved to a constant until a note is actually sounding.
type InstVibTriField struct {
	Line       int
	Cmp        float64
	Speed      float64
	SpeedShift *int
	Delay      float64
}
type InstSlideField struct {
	Line int
	Time float64
}

// InstGateField is either a single table cell (Value non-nil, exactly
// one token was given), a full per-frame gate-bit program (Tokens
// non-nil, 2+ tokens), or a bare mode word (Mode non-empty, "legato" or
// "retrigger" - ADR 0003 D11) - see parser.js's parseGateField comment.
type InstGateField struct {
	Line         int
	Value        TableValue
	Tokens       []TableValue
	Speed, Delay *int
	Mode         string
	Verbs        []TableVerb
}

// === voice_stmt ==============================================================

// VoiceStmt covers both grammar modes (Mode "short" | "extended"), same
// as parser.js's single dynamically-shaped node - short-mode fields
// (Notes/InlineAttrs/Transforms) and extended-mode fields (Fields) are
// mutually exclusive, selected by Mode.
type VoiceStmt struct {
	Line           int
	Index          int
	InstrumentName *string
	Mode           string // "short" | "extended"

	// Short mode only. Notes is nil for the arrangement form (Arrangement
	// set instead, ADR 0003 D4) - the two are mutually exclusive.
	Notes       *NotesValue
	InlineAttrs []InlineAttr
	Transforms  []Transform

	// Extended mode only:
	Fields []VoiceField
	// Arrangement is non-nil for short-mode voice text written as an
	// arrangement (ADR 0003 D4) - `voice 0 $a $b` - rather than a quoted
	// string or a single bare sigil. Mutually exclusive with Notes, which
	// stays zero-valued in that case. A voice text of exactly one bare
	// sigil is NOT an arrangement of one element; it stays on the Notes
	// path (the grammar's own ambiguity-resolution-by-order guarantees
	// this - see the .g4 file's comment on arrangement).
	Arrangement []ArrangementElement
}

// ArrangementElement is one entry of an arrangement (ADR 0003 D4):
// either a bare $name reference (Ref set) or a nested [...] group (Group
// set), exactly one of the two, optionally repeated `*n` (Repeat, default
// 1). Nesting to any depth is legal; flattening happens at interpret
// time, not here.
type ArrangementElement struct {
	Line   int
	Ref    *string
	Group  []ArrangementElement
	Repeat int
}

// NotesValue is either a quoted STRING (IsSigil false) or a bare $name
// reference (IsSigil true) - parseNotesValue() accepts both
// interchangeably wherever the grammar calls for a note pattern.
type NotesValue struct {
	IsSigil   bool
	SigilName string
	Source    string
	DocStart  int
	DocEnd    int
}

// InlineAttr is `IDENT ':' value_list`, e.g. `wave:pulse` or
// `adsr:0,4,10,4` - OR, for `wave:` and `insts:` specifically, a quoted
// mini-notation pattern (IsPattern true) instead of a flat value list.
type InlineAttr struct {
	Name      string
	IsPattern bool
	Source    string
	DocStart  int
	DocEnd    int
	Values    []TableValue
}

// VoiceField is the sealed interface for every extended-mode voice_field
// kind: PipelineField, NotesField, InstsField, FxField, RegField, VoiceVibField,
// VoiceSlideField, VoiceGateField, SyncField, RingField, VoiceWaveField.
type VoiceField interface{ isVoiceField() }

func (*PipelineField) isVoiceField()   {}
func (*NotesField) isVoiceField()      {}
func (*InstsField) isVoiceField()      {}
func (*FxField) isVoiceField()         {}
func (*RegField) isVoiceField()        {}
func (*VoiceVibField) isVoiceField()   {}
func (*VoiceSlideField) isVoiceField() {}
func (*VoiceGateField) isVoiceField()  {}
func (*SyncField) isVoiceField()       {}
func (*RingField) isVoiceField()       {}
func (*VoiceWaveField) isVoiceField()  {}

type PipelineField struct {
	Line       int
	Transforms []Transform
}
type NotesField struct {
	Line int
	NotesValue
}
type InstsField struct {
	Line             int
	Source           string
	DocStart, DocEnd int
}

// FxField is the `fx "..."` voice field - the per-row effect-command lane
// (see src/dsl/fx.js), shaped exactly like InstsField.
type FxField struct {
	Line             int
	Source           string
	DocStart, DocEnd int
}

type RegField struct {
	Line   int
	Target RegTarget
	Source string
}

// RegTarget is either a symbolic register name (Kind "ident") or a raw
// 0xNN hex address (Kind "hex").
type RegTarget struct {
	Kind  string // "hex" | "ident"
	Hex   int
	Ident string
}

type VoiceVibField struct {
	Line        int
	Depth, Rate float64
}
type VoiceSlideField struct {
	Line int
	Time float64
}
type VoiceGateField struct {
	Line  int
	Value float64
}

// BoolFieldValue is a sync/ring field's value: a plain 0/1 constant
// (Kind "number") or a full quoted mini-notation pattern (Kind
// "pattern").
type BoolFieldValue struct {
	Kind   string // "number" | "pattern"
	Number float64
	Source string
}

type SyncField struct {
	Line  int
	Value BoolFieldValue
}
type RingField struct {
	Line  int
	Value BoolFieldValue
}
type VoiceWaveField struct {
	Line             int
	Source           string
	DocStart, DocEnd int
}

// === pipeline / transform ====================================================

// Transform is the sealed interface for every transform kind (rev, fast,
// slow, every, iter, ply, off, jux, degradeBy, sometimesBy, add, struct,
// mask, chunk, palindrome, linger, rot, someCyclesBy, superimpose,
// scale).
type Transform interface{ isTransform() }

// TransformArg is a transform's numeric argument: a bare NUMBER
// (HasSource false) or a quoted mini-notation STRING sampled once per
// cycle (HasSource true) - Strudel's "a control isn't a constant, it's a
// pattern" idea, see parser.js's parseArg comment.
type TransformArg struct {
	HasSource bool
	Source    string
	Value     float64
}

type RevTransform struct{ Line int }
type FastTransform struct {
	Line int
	N    TransformArg
}
type SlowTransform struct {
	Line int
	N    TransformArg
}
type EveryTransform struct {
	Line int
	N    TransformArg
	Fn   Transform
}
type IterTransform struct {
	Line int
	N    TransformArg
}
type PlyTransform struct {
	Line int
	N    TransformArg
}
type OffTransform struct {
	Line int
	Time TransformArg
	Fn   Transform
}
type JuxTransform struct {
	Line int
	Fn   Transform
}
type DegradeByTransform struct {
	Line int
	P    TransformArg
}
type SometimesByTransform struct {
	Line int
	P    TransformArg
	Fn   Transform
}
type AddTransform struct {
	Line int
	N    TransformArg
}
type StructTransform struct {
	Line   int
	Source string
}
type MaskTransform struct {
	Line   int
	Source string
}
type ChunkTransform struct {
	Line int
	N    TransformArg
	Fn   Transform
}
type PalindromeTransform struct{ Line int }
type LingerTransform struct {
	Line int
	Frac TransformArg
}
type RotTransform struct {
	Line int
	N    TransformArg
}
type SomeCyclesByTransform struct {
	Line int
	P    TransformArg
	Fn   Transform
}
type SuperimposeTransform struct {
	Line int
	Fn   Transform
}
type ScaleTransform struct {
	Line   int
	Source string
}

func (*RevTransform) isTransform()          {}
func (*FastTransform) isTransform()         {}
func (*SlowTransform) isTransform()         {}
func (*EveryTransform) isTransform()        {}
func (*IterTransform) isTransform()         {}
func (*PlyTransform) isTransform()          {}
func (*OffTransform) isTransform()          {}
func (*JuxTransform) isTransform()          {}
func (*DegradeByTransform) isTransform()    {}
func (*SometimesByTransform) isTransform()  {}
func (*AddTransform) isTransform()          {}
func (*StructTransform) isTransform()       {}
func (*MaskTransform) isTransform()         {}
func (*ChunkTransform) isTransform()        {}
func (*PalindromeTransform) isTransform()   {}
func (*LingerTransform) isTransform()       {}
func (*RotTransform) isTransform()          {}
func (*SomeCyclesByTransform) isTransform() {}
func (*SuperimposeTransform) isTransform()  {}
func (*ScaleTransform) isTransform()        {}

// === filter_block =============================================================

type FilterBlock struct {
	Line   int
	Fields []FilterField
}

// FilterField is the sealed interface for cutoff/resonance/mode/route.
type FilterField interface{ isFilterField() }

func (*CutoffField) isFilterField()    {}
func (*ResonanceField) isFilterField() {}
func (*ModeField) isFilterField()      {}
func (*RouteField) isFilterField()     {}

// All four filter fields (cutoff/resonance/mode/route) share one Expr shape
// - FilterExpr - because the grammar's filterFieldValue rule is itself
// generic across all four (see RasterklangDsl.g4): a plain constant
// (*NumberListExpr for resonance's single number and route's voice-index
// list), a named signal ramp (*SignalExpr - used by cutoff/resonance for
// ramps and by mode for a bare constant mode name, since a mode name is
// just a SignalExpr with no range/speed), or a *SwitchBlock (time-varying,
// any field). Which of the three a given field actually produces in
// practice is enforced by ast_visitor.go/astVisitor.js, not by the type
// system - keeping one shared interface here (rather than one bespoke
// sealed interface per field) avoids four near-identical copies of the same
// three cases.
type CutoffField struct {
	Line int
	Expr FilterExpr
}
type ResonanceField struct {
	Line int
	Expr FilterExpr
}
type ModeField struct {
	Line int
	Expr FilterExpr
}
type RouteField struct {
	Line int
	Expr FilterExpr
}

// FilterExpr is the sealed interface shared by all four filter fields:
// *SignalExpr | *SwitchBlock | *NumberListExpr.
type FilterExpr interface{ isFilterExpr() }

func (*SignalExpr) isFilterExpr()     {}
func (*SwitchBlock) isFilterExpr()    {}
func (*NumberListExpr) isFilterExpr() {}

// NumberListExpr := numberList, i.e. one or more bare numbers - resonance's
// flat constant is always exactly one; route's flat constant (and each of
// its switch-case values) is one or more voice indices.
type NumberListExpr struct {
	Line   int
	Values []float64
}

// SignalExpr := SIGNAL_NAME range? speed* - Range is nil when omitted;
// Speeds may hold more than one chained modifier (the spec's own worked
// example chains two), applied left to right.
type SignalExpr struct {
	Line   int
	Signal string
	Range  *SignalRange
	Speeds []SpeedMod
}
type SignalRange struct{ Lo, Hi float64 }
type SpeedMod struct {
	Op    string // "*" | "/" | ">>"
	Value float64
}

type SwitchBlock struct {
	Line  int
	Cases []SwitchCase
}

// SwitchCase's Cond.Op is "<", ">=", or "else" (Cond.Value unused for
// "else"). Value is normally a *SignalExpr (a ramp for cutoff/resonance, or
// a bare mode name for mode); route's case values are a voice-index list
// and use *NumberListExpr instead (see switchCase's grammar rule).
type SwitchCase struct {
	Line  int
	Cond  SwitchCond
	Value SwitchCaseValue
}

// SwitchCaseValue is the sealed interface for a switch-case's value:
// *SignalExpr | *NumberListExpr.
type SwitchCaseValue interface{ isSwitchCaseValue() }

func (*SignalExpr) isSwitchCaseValue()     {}
func (*NumberListExpr) isSwitchCaseValue() {}

type SwitchCond struct {
	Op    string
	Value float64
}
