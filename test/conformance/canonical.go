// Package conformance holds the Go half of the cross-language
// conformance suite: canonical.go converts pkg/dsl's Go AST into the same
// plain JSON-shaped value (map[string]any / []any / float64 / string /
// bool / nil) that JSON.parse(JSON.stringify(jsAst)) produces on the JS
// side, field-for-field, so conformance_test.go can compare it directly
// against the shared golden files in fixtures/*.expected.json with
// reflect.DeepEqual - no custom diffing logic, and any field-name or
// shape mismatch between the two ports shows up as a normal Go test
// failure with a real diff.
//
// This file is deliberately verbose (one case per AST node kind) rather
// than reflection-based, on the theory that a conformance suite whose own
// translation layer is too clever to trust defeats its own purpose.
package conformance

import "github.com/dnoegel/rasterklang-dsl/pkg/dsl"

// f converts any Go numeric type used in the AST to float64, matching
// how every JS number (int or float) round-trips through JSON.
func f(v float64) float64 { return v }

func canonProgram(p *dsl.Program) map[string]any {
	stmts := make([]any, len(p.Statements))
	for i, s := range p.Statements {
		stmts[i] = canonStatement(s)
	}
	return map[string]any{"type": "Program", "statements": stmts}
}

func canonStatement(s dsl.Statement) any {
	switch n := s.(type) {
	case *dsl.TempoStmt:
		var beats any
		if n.Beats != nil {
			beats = float64(*n.Beats)
		}
		var offset any
		if n.Offset != nil {
			offset = *n.Offset
		}
		var meter any
		if n.Meter != nil {
			meter = map[string]any{"numerator": f(float64(n.Meter.Numerator)), "denominator": f(float64(n.Meter.Denominator))}
		}
		return map[string]any{"type": "TempoStmt", "line": f(float64(n.Line)), "unit": n.Unit, "value": canonTempoValue(n.Value), "beats": beats, "offset": offset, "meter": meter}
	case *dsl.TargetStmt:
		var driveCycles any
		if n.DriveCycles != nil {
			driveCycles = float64(*n.DriveCycles)
		}
		var multispeed any
		if n.Multispeed != nil {
			multispeed = float64(*n.Multispeed)
		}
		var video any
		if n.Video != "" {
			video = n.Video
		}
		return map[string]any{"type": "TargetStmt", "line": f(float64(n.Line)), "video": video, "driveCycles": driveCycles, "multispeed": multispeed}
	case *dsl.ChipStmt:
		return map[string]any{"type": "ChipStmt", "line": f(float64(n.Line)), "model": f(float64(n.Model))}
	case *dsl.PatternDecl:
		var length any
		if n.Length != nil {
			length = *n.Length
		}
		var each any
		if n.Each != nil {
			each = map[string]any{"numerator": f(float64(n.Each.Numerator)), "denominator": f(float64(n.Each.Denominator))}
		}
		var source any = n.Source
		var docStart any = f(float64(n.DocStart))
		var docEnd any = f(float64(n.DocEnd))
		var elements any
		if n.Elements != nil {
			elements = canonArrangement(n.Elements)
			source, docStart, docEnd = nil, nil, nil
		}
		return map[string]any{
			"type": "PatternDecl", "line": f(float64(n.Line)), "name": n.Name, "length": length, "each": each,
			"elements": elements, "source": source, "docStart": docStart, "docEnd": docEnd,
		}
	case *dsl.TableDecl:
		return map[string]any{"type": "TableDecl", "line": f(float64(n.Line)), "name": n.Name, "tokens": canonTableValues(n.Tokens)}
	case *dsl.InstrumentBlock:
		fields := make([]any, len(n.Fields))
		for i, fl := range n.Fields {
			fields[i] = canonInstrumentField(fl)
		}
		return map[string]any{"type": "InstrumentBlock", "line": f(float64(n.Line)), "name": n.Name, "fields": fields}
	case *dsl.VoiceStmt:
		return canonVoiceStmt(n)
	case *dsl.FilterBlock:
		fields := make([]any, len(n.Fields))
		for i, fl := range n.Fields {
			fields[i] = canonFilterField(fl)
		}
		return map[string]any{"type": "FilterBlock", "line": f(float64(n.Line)), "fields": fields}
	default:
		panic("canonStatement: unhandled statement type")
	}
}

func canonTempoValue(v dsl.TempoValue) map[string]any {
	if v.Kind == "number" {
		return map[string]any{"kind": "number", "value": f(v.Number)}
	}
	return map[string]any{"kind": "pattern", "source": v.Source}
}

// canonTableValues mirrors parseTableRow()'s cell shapes: a bare word is a
// plain JS string, a number a plain JS number, a $ref an {ref} object, a
// repeat a {repeat,count} object, and a ramp its own {gen:'ramp',...}
// object.
func canonTableValues(tokens []dsl.TableValue) []any {
	out := make([]any, len(tokens))
	for i, t := range tokens {
		out[i] = canonTableValue(t)
	}
	return out
}

func canonTableValue(t dsl.TableValue) any {
	switch v := t.(type) {
	case dsl.TableWord:
		return v.Value
	case dsl.TableNumber:
		return f(v.Value)
	case dsl.TableRef:
		return map[string]any{"ref": v.Name}
	case dsl.TableAbsNote:
		return map[string]any{"absNote": f(v.Value)}
	case dsl.TableRepeat:
		return map[string]any{"repeat": canonTableValue(v.Value), "count": f(float64(v.Count))}
	case *dsl.TableRamp:
		var step, over any
		if v.Step != nil {
			step = f(*v.Step)
		}
		if v.Over != nil {
			over = f(float64(*v.Over))
		}
		var shape any
		if v.Shape != "" {
			shape = v.Shape
		}
		return map[string]any{"gen": "ramp", "line": f(float64(v.Line)), "from": f(v.From), "to": f(v.To), "step": step, "over": over, "shape": shape}
	default:
		panic("canonTableValue: unhandled table value type")
	}
}

// canonSpeedDelay -> (speed, delay) as `any` (nil unless set), for the ADR
// 0003 D10 fields wave/arp/pw/filterTable/test/gate all carry.
func canonSpeedDelay(speed, delay *int) (any, any) {
	var s, d any
	if speed != nil {
		s = f(float64(*speed))
	}
	if delay != nil {
		d = f(float64(*delay))
	}
	return s, d
}

// canonVerbs (ADR 0003 D13) mirrors astVisitor.js's buildTableVerbLine
// output shape exactly.
func canonVerbs(verbs []dsl.TableVerb) []any {
	out := make([]any, len(verbs))
	for i, v := range verbs {
		switch v.Verb {
		case "set":
			out[i] = map[string]any{"verb": v.Verb, "line": f(float64(v.Line)), "value": f(v.Value)}
		case "add":
			out[i] = map[string]any{"verb": v.Verb, "line": f(float64(v.Line)), "value": f(v.Value), "repeat": f(float64(v.Repeat))}
		case "wait":
			out[i] = map[string]any{"verb": v.Verb, "line": f(float64(v.Line)), "count": f(float64(v.Count))}
		default:
			out[i] = map[string]any{"verb": v.Verb, "line": f(float64(v.Line))}
		}
	}
	return out
}

func canonInstrumentField(fl dsl.InstrumentField) map[string]any {
	switch n := fl.(type) {
	case *dsl.WaveField:
		if n.Verbs != nil {
			return map[string]any{"type": "wave", "line": f(float64(n.Line)), "verbs": canonVerbs(n.Verbs)}
		}
		speed, delay := canonSpeedDelay(n.Speed, n.Delay)
		return map[string]any{"type": "wave", "line": f(float64(n.Line)), "tokens": canonTableValues(n.Tokens), "speed": speed, "delay": delay}
	case *dsl.ArpField:
		if n.Verbs != nil {
			return map[string]any{"type": "arp", "line": f(float64(n.Line)), "verbs": canonVerbs(n.Verbs)}
		}
		speed, delay := canonSpeedDelay(n.Speed, n.Delay)
		return map[string]any{"type": "arp", "line": f(float64(n.Line)), "tokens": canonTableValues(n.Tokens), "speed": speed, "delay": delay}
	case *dsl.PwField:
		if n.Verbs != nil {
			return map[string]any{"type": "pw", "line": f(float64(n.Line)), "verbs": canonVerbs(n.Verbs)}
		}
		speed, delay := canonSpeedDelay(n.Speed, n.Delay)
		return map[string]any{"type": "pw", "line": f(float64(n.Line)), "tokens": canonTableValues(n.Tokens), "speed": speed, "delay": delay}
	case *dsl.FilterTableField:
		if n.Verbs != nil {
			return map[string]any{"type": "filterTable", "line": f(float64(n.Line)), "verbs": canonVerbs(n.Verbs)}
		}
		speed, delay := canonSpeedDelay(n.Speed, n.Delay)
		return map[string]any{"type": "filterTable", "line": f(float64(n.Line)), "tokens": canonTableValues(n.Tokens), "speed": speed, "delay": delay}
	case *dsl.TestField:
		if n.Verbs != nil {
			return map[string]any{"type": "test", "line": f(float64(n.Line)), "verbs": canonVerbs(n.Verbs)}
		}
		speed, delay := canonSpeedDelay(n.Speed, n.Delay)
		return map[string]any{"type": "test", "line": f(float64(n.Line)), "tokens": canonTableValues(n.Tokens), "speed": speed, "delay": delay}
	case *dsl.AdsrField:
		return map[string]any{"type": "adsr", "line": f(float64(n.Line)), "a": f(n.A), "d": f(n.D), "s": f(n.S), "r": f(n.R)}
	case *dsl.LoopField:
		m := map[string]any{"type": "loop", "line": f(float64(n.Line))}
		if n.Value != nil {
			m["value"] = f(float64(*n.Value))
		} else {
			pairs := make([]any, len(n.Pairs))
			for i, pr := range n.Pairs {
				pairs[i] = map[string]any{"table": pr.Table, "value": f(float64(pr.Value))}
			}
			m["pairs"] = pairs
		}
		return m
	case *dsl.InstVibField:
		return map[string]any{"type": "vib", "line": f(float64(n.Line)), "depth": f(n.Depth), "rate": f(n.Rate)}
	case *dsl.InstVibTriField:
		m := map[string]any{"type": "vibTri", "line": f(float64(n.Line)), "cmp": f(n.Cmp), "delay": f(n.Delay)}
		if n.SpeedShift != nil {
			m["speedShift"] = f(float64(*n.SpeedShift))
		} else {
			m["speed"] = f(n.Speed)
		}
		return m
	case *dsl.InstSlideField:
		return map[string]any{"type": "slide", "line": f(float64(n.Line)), "time": f(n.Time)}
	case *dsl.InstGateField:
		m := map[string]any{"type": "gate", "line": f(float64(n.Line))}
		if n.Verbs != nil {
			m["verbs"] = canonVerbs(n.Verbs)
		} else if n.Mode != "" {
			m["mode"] = n.Mode
		} else if n.Value != nil {
			m["value"] = canonTableValue(n.Value)
		} else {
			m["tokens"] = canonTableValues(n.Tokens)
			m["speed"], m["delay"] = canonSpeedDelay(n.Speed, n.Delay)
		}
		return m
	default:
		panic("canonInstrumentField: unhandled instrument field type")
	}
}

func canonNotesValue(nv dsl.NotesValue) map[string]any {
	if nv.IsSigil {
		return map[string]any{"isSigil": true, "sigilName": nv.SigilName}
	}
	return map[string]any{"isSigil": false, "source": nv.Source, "docStart": f(float64(nv.DocStart)), "docEnd": f(float64(nv.DocEnd))}
}

// canonArrangement mirrors astVisitor.js's buildArrangementElements()
// output shape: one map per element, `ref` (a string) xor `group` (a
// nested []any), plus `repeat`.
func canonArrangement(elements []dsl.ArrangementElement) []any {
	out := make([]any, len(elements))
	for i, el := range elements {
		m := map[string]any{"line": f(float64(el.Line)), "repeat": f(float64(el.Repeat))}
		// Only the applicable one of ref/group is present, matching
		// astVisitor.js's buildArrangementElements() object literal, which
		// never sets the other key at all (not even to null).
		if el.Ref != nil {
			m["ref"] = *el.Ref
		} else {
			m["group"] = canonArrangement(el.Group)
		}
		out[i] = m
	}
	return out
}

func canonInlineAttr(a dsl.InlineAttr) map[string]any {
	if a.IsPattern {
		return map[string]any{
			"name": a.Name, "isPattern": true, "source": a.Source,
			"docStart": f(float64(a.DocStart)), "docEnd": f(float64(a.DocEnd)),
		}
	}
	return map[string]any{"name": a.Name, "values": canonTableValues(a.Values)}
}

func canonVoiceStmt(n *dsl.VoiceStmt) map[string]any {
	var instrumentName any
	if n.InstrumentName != nil {
		instrumentName = *n.InstrumentName
	}
	base := map[string]any{
		"type": "VoiceStmt", "mode": n.Mode, "line": f(float64(n.Line)), "index": f(float64(n.Index)), "instrumentName": instrumentName,
	}
	if n.Mode == "short" {
		attrs := make([]any, len(n.InlineAttrs))
		for i, a := range n.InlineAttrs {
			attrs[i] = canonInlineAttr(a)
		}
		transforms := make([]any, len(n.Transforms))
		for i, t := range n.Transforms {
			transforms[i] = canonTransform(t)
		}
		if n.Notes != nil {
			base["notes"] = canonNotesValue(*n.Notes)
		} else {
			base["notes"] = nil
		}
		if n.Arrangement != nil {
			base["arrangement"] = canonArrangement(n.Arrangement)
		} else {
			base["arrangement"] = nil
		}
		base["inlineAttrs"] = attrs
		base["transforms"] = transforms
		return base
	}
	fields := make([]any, len(n.Fields))
	for i, fl := range n.Fields {
		fields[i] = canonVoiceField(fl)
	}
	base["fields"] = fields
	return base
}

func canonVoiceField(fl dsl.VoiceField) map[string]any {
	switch n := fl.(type) {
	case *dsl.PipelineField:
		ts := make([]any, len(n.Transforms))
		for i, t := range n.Transforms {
			ts[i] = canonTransform(t)
		}
		return map[string]any{"type": "pipeline", "line": f(float64(n.Line)), "transforms": ts}
	case *dsl.NotesField:
		m := canonNotesValue(n.NotesValue)
		m["type"] = "notes"
		m["line"] = f(float64(n.Line))
		return m
	case *dsl.InstsField:
		return map[string]any{"type": "insts", "line": f(float64(n.Line)), "source": n.Source, "docStart": f(float64(n.DocStart)), "docEnd": f(float64(n.DocEnd))}
	case *dsl.FxField:
		return map[string]any{"type": "fx", "line": f(float64(n.Line)), "source": n.Source, "docStart": f(float64(n.DocStart)), "docEnd": f(float64(n.DocEnd))}
	case *dsl.RegField:
		return map[string]any{"type": "reg", "line": f(float64(n.Line)), "target": canonRegTarget(n.Target), "source": n.Source}
	case *dsl.VoiceVibField:
		return map[string]any{"type": "vib", "line": f(float64(n.Line)), "depth": f(n.Depth), "rate": f(n.Rate)}
	case *dsl.VoiceSlideField:
		return map[string]any{"type": "slide", "line": f(float64(n.Line)), "time": f(n.Time)}
	case *dsl.VoiceGateField:
		return map[string]any{"type": "gate", "line": f(float64(n.Line)), "value": f(n.Value)}
	case *dsl.SyncField:
		return map[string]any{"type": "sync", "line": f(float64(n.Line)), "value": canonBoolFieldValue(n.Value)}
	case *dsl.RingField:
		return map[string]any{"type": "ring", "line": f(float64(n.Line)), "value": canonBoolFieldValue(n.Value)}
	case *dsl.VoiceWaveField:
		return map[string]any{"type": "wave", "line": f(float64(n.Line)), "source": n.Source, "docStart": f(float64(n.DocStart)), "docEnd": f(float64(n.DocEnd))}
	default:
		panic("canonVoiceField: unhandled voice field type")
	}
}

func canonRegTarget(t dsl.RegTarget) map[string]any {
	if t.Kind == "hex" {
		return map[string]any{"kind": "hex", "value": f(float64(t.Hex))}
	}
	return map[string]any{"kind": "ident", "name": t.Ident}
}

func canonBoolFieldValue(v dsl.BoolFieldValue) map[string]any {
	if v.Kind == "number" {
		return map[string]any{"kind": "number", "value": f(v.Number)}
	}
	return map[string]any{"kind": "pattern", "source": v.Source}
}

func canonArg(a dsl.TransformArg) map[string]any {
	if a.HasSource {
		return map[string]any{"source": a.Source}
	}
	return map[string]any{"value": f(a.Value)}
}

func canonTransform(t dsl.Transform) map[string]any {
	switch n := t.(type) {
	case *dsl.RevTransform:
		return map[string]any{"type": "rev", "line": f(float64(n.Line))}
	case *dsl.FastTransform:
		return map[string]any{"type": "fast", "line": f(float64(n.Line)), "n": canonArg(n.N)}
	case *dsl.SlowTransform:
		return map[string]any{"type": "slow", "line": f(float64(n.Line)), "n": canonArg(n.N)}
	case *dsl.EveryTransform:
		return map[string]any{"type": "every", "line": f(float64(n.Line)), "n": canonArg(n.N), "fn": canonTransform(n.Fn)}
	case *dsl.IterTransform:
		return map[string]any{"type": "iter", "line": f(float64(n.Line)), "n": canonArg(n.N)}
	case *dsl.PlyTransform:
		return map[string]any{"type": "ply", "line": f(float64(n.Line)), "n": canonArg(n.N)}
	case *dsl.OffTransform:
		return map[string]any{"type": "off", "line": f(float64(n.Line)), "time": canonArg(n.Time), "fn": canonTransform(n.Fn)}
	case *dsl.JuxTransform:
		return map[string]any{"type": "jux", "line": f(float64(n.Line)), "fn": canonTransform(n.Fn)}
	case *dsl.DegradeByTransform:
		return map[string]any{"type": "degradeBy", "line": f(float64(n.Line)), "p": canonArg(n.P)}
	case *dsl.SometimesByTransform:
		return map[string]any{"type": "sometimesBy", "line": f(float64(n.Line)), "p": canonArg(n.P), "fn": canonTransform(n.Fn)}
	case *dsl.AddTransform:
		return map[string]any{"type": "add", "line": f(float64(n.Line)), "n": canonArg(n.N)}
	case *dsl.StructTransform:
		return map[string]any{"type": "struct", "line": f(float64(n.Line)), "source": n.Source}
	case *dsl.MaskTransform:
		return map[string]any{"type": "mask", "line": f(float64(n.Line)), "source": n.Source}
	case *dsl.ChunkTransform:
		return map[string]any{"type": "chunk", "line": f(float64(n.Line)), "n": canonArg(n.N), "fn": canonTransform(n.Fn)}
	case *dsl.PalindromeTransform:
		return map[string]any{"type": "palindrome", "line": f(float64(n.Line))}
	case *dsl.LingerTransform:
		return map[string]any{"type": "linger", "line": f(float64(n.Line)), "frac": canonArg(n.Frac)}
	case *dsl.RotTransform:
		return map[string]any{"type": "rot", "line": f(float64(n.Line)), "n": canonArg(n.N)}
	case *dsl.SomeCyclesByTransform:
		return map[string]any{"type": "someCyclesBy", "line": f(float64(n.Line)), "p": canonArg(n.P), "fn": canonTransform(n.Fn)}
	case *dsl.SuperimposeTransform:
		return map[string]any{"type": "superimpose", "line": f(float64(n.Line)), "fn": canonTransform(n.Fn)}
	case *dsl.ScaleTransform:
		return map[string]any{"type": "scale", "line": f(float64(n.Line)), "source": n.Source}
	default:
		panic("canonTransform: unhandled transform type")
	}
}

func canonFilterField(fl dsl.FilterField) map[string]any {
	switch n := fl.(type) {
	case *dsl.CutoffField:
		return map[string]any{"type": "cutoff", "line": f(float64(n.Line)), "expr": canonFilterExpr(n.Expr)}
	case *dsl.ResonanceField:
		return map[string]any{"type": "resonance", "line": f(float64(n.Line)), "expr": canonFilterExpr(n.Expr)}
	case *dsl.ModeField:
		return map[string]any{"type": "mode", "line": f(float64(n.Line)), "expr": canonFilterExpr(n.Expr)}
	case *dsl.RouteField:
		return map[string]any{"type": "route", "line": f(float64(n.Line)), "expr": canonFilterExpr(n.Expr)}
	default:
		panic("canonFilterField: unhandled filter field type")
	}
}

func canonFilterExpr(e dsl.FilterExpr) map[string]any {
	switch n := e.(type) {
	case *dsl.SignalExpr:
		return canonSignalExpr(n)
	case *dsl.NumberListExpr:
		return canonNumberListExpr(n)
	case *dsl.SwitchBlock:
		cases := make([]any, len(n.Cases))
		for i, c := range n.Cases {
			cond := map[string]any{"op": c.Cond.Op}
			if c.Cond.Op != "else" {
				cond["value"] = f(c.Cond.Value)
			}
			cases[i] = map[string]any{"line": f(float64(c.Line)), "cond": cond, "value": canonSwitchCaseValue(c.Value)}
		}
		return map[string]any{"type": "SwitchBlock", "line": f(float64(n.Line)), "cases": cases}
	default:
		panic("canonFilterExpr: unhandled filter expr type")
	}
}

func canonSwitchCaseValue(v dsl.SwitchCaseValue) map[string]any {
	switch n := v.(type) {
	case *dsl.SignalExpr:
		return canonSignalExpr(n)
	case *dsl.NumberListExpr:
		return canonNumberListExpr(n)
	default:
		panic("canonSwitchCaseValue: unhandled switch case value type")
	}
}

func canonNumberListExpr(n *dsl.NumberListExpr) map[string]any {
	values := make([]any, len(n.Values))
	for i, v := range n.Values {
		values[i] = f(v)
	}
	return map[string]any{"type": "NumberListExpr", "line": f(float64(n.Line)), "values": values}
}

func canonSignalExpr(n *dsl.SignalExpr) map[string]any {
	var rng any
	if n.Range != nil {
		rng = map[string]any{"lo": f(n.Range.Lo), "hi": f(n.Range.Hi)}
	}
	speeds := make([]any, len(n.Speeds))
	for i, s := range n.Speeds {
		speeds[i] = map[string]any{"op": s.Op, "value": f(s.Value)}
	}
	return map[string]any{"type": "SignalExpr", "line": f(float64(n.Line)), "signal": n.Signal, "range": rng, "speeds": speeds}
}
