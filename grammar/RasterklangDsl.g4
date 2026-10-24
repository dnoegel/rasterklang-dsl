// The DSL v2 grammar -- see the repo root README.md's "Grammar" section
// for the regeneration workflow. This file is what a human edits; the
// generated parsers (src/dsl/generated/, pkg/dsl/generated/) are derived,
// never hand-edited.
//
// Soft keywords, no reserved words: `tempo`, `wave`, `adsr`, `fast`, etc.
// are all just identifier-shaped text the parser/AST-transform-visitor
// dispatches on, so the same word can be used as a plain identifier
// elsewhere (e.g. `wave:pulse` short-mode syntax) without colliding with
// `wave` the keyword. ANTLR auto-promotes any string literal referenced
// directly in a parser rule into its own reserved lexer token that wins
// over IDENT -- naively writing `'wave' tableRow` in a parser rule would
// silently break that. This grammar avoids it two ways:
//   - For every keyword/field name EXCEPT tempo/voice/end (see below), the
//     parser rules are written generically (any IDENT-shaped token in that
//     position), and the AST-transform visitor dispatches on the token's
//     text after the fact -- see the "soft keywords" pattern documented at
//     https://github.com/antlr/antlr4/blob/master/doc/predicates.md.
//     Structural shape differences (token type, not text) are what the
//     grammar itself uses to disambiguate one production from another.
//   - The 8 statement-leading keywords (tempo/speed/voice/pattern/table/
//     instrument/filter/target) are the exception: each needs its own dedicated
//     lexer token so ANTLR can commit to a specific production early
//     enough to produce a precise, position-anchored parse error, and (for
//     tempo/voice specifically) because `tempo N` / `voice N` share an
//     unbounded-lookahead-ambiguous prefix that adaptive LL parsing can't
//     safely resolve if both stay generic IDENT. The `anyIdent: IDENT |
//     TEMPO | SPEED | VOICE | PATTERN | TABLE | INSTRUMENT | FILTER` rule
//     below is used everywhere a plain identifier is still syntactically
//     legal, so all 8 keep working as ordinary identifiers (instrument names,
//     pattern names, etc.) too -- that's what makes them *soft* keywords.
//   - `end` (the block terminator) is also its own dedicated token but is
//     deliberately NOT in `anyIdent` -- it's the one genuinely reserved
//     word in this grammar, since it's a structural block-closer rather
//     than a field/statement name a user would plausibly want to reuse as
//     an identifier.
grammar RasterklangDsl;

// ===== Parser ================================================================

program: (statement NEWLINE)* EOF;

statement
  : tempoStmt
  | targetStmt
  | chipStmt
  | voiceStmt
  | patternDecl
  | tableDecl
  | instrumentBlock
  | filterBlock
  | unknownStmt
  ;

// Catch-all so a bogus leading keyword (one that doesn't even structurally
// match pattern/table/instrument/filter's shape, e.g. `bogus 1 2 3`) still
// parses as A statement -- deferring to the AST-transform visitor's
// leading-keyword check to report "unknown statement keyword 'bogus'"
// (matching parser.js's own behavior) instead of ANTLR's much less
// specific "no viable alternative" syntax error. Must be the LAST
// alternative so a well-formed statement always prefers its real,
// specific production.
unknownStmt: IDENT unknownTail* ;
unknownTail: ~NEWLINE ;

anyIdent: IDENT | TEMPO | SPEED | VOICE | PATTERN | TABLE | INSTRUMENT | FILTER | TARGET | CHIP ;

str: DQSTRING | BTSTRING ;

// --- tempo / speed (the song clock) ---
// The trailing (anyIdent NUMBER) pairs are the song-level settings that
// belong to the clock rather than to any one voice: `beats N` (how many
// beats one BAR is - a pure time signature since 3.0.0, see ADR 0001's D2)
// and `offset N` (how many beats the whole song's pattern clock is shifted
// by). Kept as a repeated generic name/number pair rather than dedicated
// syntax per setting so a future third one costs a visitor case, not a
// grammar change - the AST-transform visitor is what validates the names,
// same soft-keyword pattern as every instrument field.
//
// TEMPO and SPEED are two spellings of the same statement: `tempo 132` is
// BPM (musical, lossy at frame resolution), `speed 6` is frames per beat
// (exact, what GoatTracker natively counts). The grammar accepts either or
// both; "exactly one of them must be present" is deliberately NOT encoded
// as a grammar-level XOR, so the semantic pass (interpret.js) can say which
// of the two failure cases applies (ADR 0001 decision 11).
tempoStmt: (TEMPO | SPEED) tempoValue tempoSetting* ;
tempoValue: NUMBER | str ;
// ADR 0003 D2: `meter N/D` needs a value shape `beats N`/`offset N` can't
// hold (a fraction, not a bare number), so the trailing settings gained an
// optional `/NUMBER` tail. `beats`/`offset` never use it; the AST-transform
// visitor is what enforces that, same soft-keyword split as every other
// setting name here.
tempoSetting: anyIdent NUMBER (SLASH NUMBER)? ;

// --- target (the DRIVE rate: how often the player is invoked) ---
// ADR 0002. Song-level, optional, at most one. Where `tempo`/`speed`/`beats`
// answer "how long is a beat, counted in frames", this answers "how long is
// a frame, counted in CPU cycles" - the question DSL v2 previously had one
// hard-coded answer to (PAL, 19705 cycles = 50.00 Hz).
//
//   target { video PAL }               // 985248/19705  = 49.99990 Hz (the default)
//   target { video NTSC }              // 1022727/17045 = 60.00158 Hz
//   target { video PAL multispeed 3 }  // 3 play calls per video frame
//   target { driveCycles 12801 }       // exact: CPU cycles per play call
//
// `driveCycles` is the one primitive (an exact integer count of CPU cycles
// per play call); `video` and `multispeed` are named macros over it.
// Settings are generic name/value pairs for the same reason tempo's trailing
// `beats N`/`offset N` pairs are: a future fourth setting should cost a
// visitor case, not a grammar change. A value is a NUMBER (`driveCycles
// 12801`, `multispeed 3`) or a bare word (`video PAL`); the AST-transform
// visitor is what validates names, values and their combinations.
//
// Braces rather than a NEWLINE...END block on purpose: it keeps the whole
// declaration on one line, the way a header line should be.
targetStmt: TARGET LBRACE targetSetting* RBRACE ;

// --- chip (ADR 0003 D15: which SID model this song was written for) ---
// Song-level, optional, at most one. Absent means today's default (see
// liveEngine.js's DEFAULT_MODEL) - the only thing this decision adds is
// that a song can now RECORD its own choice instead of it living only in
// transient UI state, lost on save/re-open/export.
chipStmt: CHIP NUMBER ;
targetSetting: anyIdent (NUMBER | anyIdent) ;

// --- pattern / table decl ---
// Each of the 6 statement keywords gets its own dedicated lexer token (see
// this file's own header comment) -- not just for keyword/IDENT
// collision safety, but because committing to a specific token up front is
// what lets ANTLR produce a precise, position-anchored parse error ("found
// STRING, expected '='") the moment something after the keyword doesn't
// fit, matching parser.js's own recursive-descent error style, instead of
// a generic "no viable alternative" if six different shapes had to be
// disambiguated by lookahead alone.
// The optional NUMBER between a pattern's name and its `=` is its DECLARED
// LENGTH IN BEATS (`pattern bass 4 = "c4 ~ c4 ~"`), the opt-in to metric
// behaviour: an unweighted token is then one beat, and the summed content
// weight must equal the declared length or the semantic pass reports a
// mismatch (ADR 0001's D1). Decimals are legal (`pattern lead 7.5 = ...`),
// since the check is a plain comparison against a weight sum that has
// always been fractional. No number means today's proportional pattern,
// unchanged.
//
// The trailing pipelineStep* is NOT a supported form: transforms belong on
// a voice, not on a pattern declaration (see ADR 0001's Deferred section).
// It is matched here purely so the AST-transform visitor can replace
// ANTLR's raw "mismatched input '' expecting NEWLINE" with a sentence that
// says what to do instead (ADR 0001 decision 8 / known defect 3).
// `each 1/D` (ADR 0003 D3) is the step-sequencer spelling: every token in
// the pattern is one 1/D note, so the length is derived rather than
// declared. It sits between the optional declared length and `=`, e.g.
// `pattern hats each 1/16 = "..."` or `pattern hats 2 each 1/16 = "..."`
// (a declared length alongside `each` is legal and checked against the
// derived one, same as any other declared length). `each` is a soft
// keyword like every other field name, validated by the AST-transform
// visitor, not the grammar.
// The second alternative is an ARRANGEMENT decl (ADR 0003 D4): a song part
// built from other patterns by reference, sequence, repeat and nesting -
// `pattern verse = $riffA $riffA $riffB $riffA`, no quotes, no declared
// length (its length is always the sum of what it references). The two
// alternatives are structurally unambiguous: whatever follows `=` is
// either a `str` or a SIGIL/`[`, never both.
patternDecl
  : PATTERN anyIdent NUMBER? patternEach? EQ str pipelineStep*
  | PATTERN anyIdent EQ arrangement
  ;
patternEach: anyIdent NUMBER SLASH NUMBER ;

// --- arrangements (ADR 0003 D4) ---
// A sequence of pattern references, concatenated in BEAT TIME (as opposed
// to `<...>` alternation, which is untouched and still means "one element
// per bar"). This is what fixes F3: today a voice/pattern can only ever
// be ONE proportional or ONE declared-length pattern, so an intro/verse/
// chorus arrangement of DIFFERENT lengths cannot be sequenced at all.
// Deliberately not a mini-notation bracket: `{...}` is already polymetre
// (miniNotation.js's parsePolymeter), so this lives here, at statement
// level, where `$` is already unambiguous and lengths are already
// computed (see the ADR's section 5.2/section 2 D4 "Deliberately not a
// mini-notation bracket"). Mini-notation itself does not change at all
// for this decision.
//
// A voiceBody/patternDecl consisting of exactly one bare SIGIL (no
// brackets, no `*`) is ALSO matched by this rule, but is intentionally
// listed second everywhere it could apply (see voiceBody below) so ANTLR
// resolves that genuine ambiguity toward the pre-existing single-sigil
// alternative - "a sequence of exactly one sigil stays byte-identical to
// today".
arrangement: arrangementElement+ ;
arrangementElement: (SIGIL | LBRACKET arrangementElement+ RBRACKET) (STAR NUMBER)? ;
tableDecl: TABLE anyIdent EQ tableRow ;

// --- instrument ---
instrumentBlock: INSTRUMENT anyIdent NEWLINE instrumentField* END ;
// ADR 0003 D13: a table field can EITHER stay a bare token row on its own
// line (unchanged), OR open its own `end`-terminated block of verb lines
// (`set`/`add .. repeat N`/`wait N`/`loop`/`stop`) - an alternative
// spelling of the same per-frame array, for tables too long to read as a
// row of bare numbers. The two alternatives diverge on the very next
// token after anyIdent (NEWLINE vs. a tableToken), so this needs no
// lookahead beyond ANTLR's default.
instrumentField
  : anyIdent tableRow NEWLINE
  | anyIdent NEWLINE tableVerbLine+ END NEWLINE
  ;
tableVerbLine: anyIdent tableRow? NEWLINE ;

// --- table row / tokens (shared: table_decl, instrument fields, loop pairs) ---
tableRow: tableToken+ ;
// `absNote` (`=60`) is an ABSOLUTE MIDI note number for one table frame,
// as opposed to every other numeric cell in an `arp` row, which is a
// RELATIVE semitone offset from whatever note is currently sounding. It
// exists because GoatTracker wave-table rows whose note column is
// $80-$DF pin that frame to a fixed pitch regardless of the played note
// (gplay.c's WAVEEXEC: `if (note < 0x80) note += cptr->note;` - i.e. only
// the sub-$80 half is note-relative), the classic fixed-pitch drum/
// percussion technique, which a relative-only `arp` cannot express at all.
// Spelled as a plain MIDI NUMBER rather than a note name on purpose: this
// grammar's lexer has no '#' character, so a table cell could never spell
// `c#4` the way mini-notation does, and inventing a second, subtly
// different note-name convention (`cs4`) just for table rows would be a
// worse wart than one explicit number.
tableToken
  : rampGenerator
  | SIGIL
  | absNote
  | anyIdent (PLUS anyIdent)* (BANG NUMBER)?
  | NUMBER (BANG NUMBER)?
  ;
absNote: EQ NUMBER (BANG NUMBER)? ;
// The trailing bare anyIdent (ADR 0003 D9) is `pingpong`/`loop`: a sweep
// SHAPE terminator, structurally unambiguous from the `step`/`over`
// name/NUMBER pairs before it since neither takes a NUMBER of its own.
rampGenerator: anyIdent NUMBER DOTDOT NUMBER (anyIdent NUMBER)? (anyIdent NUMBER)? anyIdent? ;

// --- voice ---
voiceStmt: VOICE NUMBER anyIdent? voiceBody ;
voiceBody
  : (str | SIGIL) voiceShortTail
  | arrangement voiceShortTail
  | NEWLINE voiceField* END
  ;
voiceShortTail: (NEWLINE? inlineAttr)* (NEWLINE? pipelineStep)* ;
inlineAttr: anyIdent COLON (str | valueList) ;
valueList: tableToken (COMMA tableToken)* ;

voiceField
  : (pipelineStep | anyIdent voiceFieldValue) NEWLINE
  ;
voiceFieldValue
  : str
  | SIGIL
  | (anyIdent | HEX) str
  | NUMBER NUMBER?
  ;

// --- pipeline / transform ---
pipelineStep: PIPE transform ;
transform: anyIdent transformArg? transformRef? ;
transformArg: NUMBER | str ;
transformRef: transform | LPAREN transform RPAREN ;

// --- filter ---
filterBlock: FILTER NEWLINE filterField* END ;
filterField: anyIdent filterFieldValue NEWLINE ;
filterFieldValue
  : anyIdent anyIdent NEWLINE switchCase (NEWLINE switchCase)*
  | anyIdent signalRangeOpt signalSpeedMod*
  | numberList
  ;
// A switch-case's value is normally an identifier-led signal expression
// (a ramp/wave name for cutoff/resonance, or a bare mode name like
// `bandpass`). `route`'s case values are fundamentally a list of voice
// indices (e.g. `1 2`), which can't be spelled as anyIdent+range+speed, so
// numberList is accepted as an alternative leaf shape here too.
switchCase: switchCond COLON (anyIdent signalRangeOpt signalSpeedMod* | numberList) ;
switchCond: LT NUMBER | GE NUMBER | anyIdent ;
signalRangeOpt: (NUMBER DOTDOT NUMBER)? ;
signalSpeedMod: (STAR | SLASH) NUMBER | SHIFT NUMBER ;
numberList: NUMBER (COMMA? NUMBER)* ;

// ===== Lexer =================================================================

TEMPO: 'tempo';
SPEED: 'speed';
VOICE: 'voice';
PATTERN: 'pattern';
TABLE: 'table';
INSTRUMENT: 'instrument';
FILTER: 'filter';
TARGET: 'target';
CHIP: 'chip';
END: 'end';

HEX: '0' [xX] [0-9a-fA-F]+;
// A trailing '%' (ADR 0003 D11: 'gate 80%') is folded into NUMBER's own
// lexeme rather than given its own token, so its value resolves to the
// same fraction 'N / 100' at the one place every NUMBER's text already
// gets parsed (lexer.js's convertToken / lexer.go's convertToken) -
// nothing downstream has to know '%' syntax exists, the same way '80%'
// and '0.8' become indistinguishable numbers by the time any consumer
// sees them.
NUMBER: '-'? [0-9]+ ('.' [0-9]+)? '%'?;

DQSTRING: '"' (ESC_D | ~["\\\r\n])*? '"' ;
fragment ESC_D: '\\' ["\\];

BTSTRING: '`' (ESC_B | ~[`\\])*? '`' ;
fragment ESC_B: '\\' [`\\];

SIGIL: '$' [a-zA-Z_] [a-zA-Z0-9_]*;

IDENT: [a-zA-Z_] [a-zA-Z0-9_]*;

PIPE: '|>';
DOTDOT: '..';
SHIFT: '>>';
GE: '>=';
COLON: ':';
COMMA: ',';
STAR: '*';
SLASH: '/';
LT: '<';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
EQ: '=';
PLUS: '+';
BANG: '!';

NEWLINE: '\r'? '\n';
WS: [ \t]+ -> skip;
COMMENT: ';' ~[\r\n]* -> skip;
