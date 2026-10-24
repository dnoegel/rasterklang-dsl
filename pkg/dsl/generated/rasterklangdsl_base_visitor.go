// Code generated from RasterklangDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // RasterklangDsl
import "github.com/antlr4-go/antlr/v4"

type BaseRasterklangDslVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRasterklangDslVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitUnknownStmt(ctx *UnknownStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitUnknownTail(ctx *UnknownTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitAnyIdent(ctx *AnyIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTempoStmt(ctx *TempoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTempoValue(ctx *TempoValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTempoSetting(ctx *TempoSettingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTargetStmt(ctx *TargetStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitChipStmt(ctx *ChipStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTargetSetting(ctx *TargetSettingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitPatternDecl(ctx *PatternDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitPatternEach(ctx *PatternEachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitArrangement(ctx *ArrangementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitArrangementElement(ctx *ArrangementElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTableDecl(ctx *TableDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitInstrumentBlock(ctx *InstrumentBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitInstrumentField(ctx *InstrumentFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTableVerbLine(ctx *TableVerbLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTableRow(ctx *TableRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTableToken(ctx *TableTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitAbsNote(ctx *AbsNoteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitRampGenerator(ctx *RampGeneratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitVoiceStmt(ctx *VoiceStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitVoiceBody(ctx *VoiceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitVoiceShortTail(ctx *VoiceShortTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitInlineAttr(ctx *InlineAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitValueList(ctx *ValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitVoiceField(ctx *VoiceFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitVoiceFieldValue(ctx *VoiceFieldValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitPipelineStep(ctx *PipelineStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTransform(ctx *TransformContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTransformArg(ctx *TransformArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitTransformRef(ctx *TransformRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitFilterBlock(ctx *FilterBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitFilterField(ctx *FilterFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitFilterFieldValue(ctx *FilterFieldValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitSwitchCond(ctx *SwitchCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitSignalRangeOpt(ctx *SignalRangeOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitSignalSpeedMod(ctx *SignalSpeedModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRasterklangDslVisitor) VisitNumberList(ctx *NumberListContext) interface{} {
	return v.VisitChildren(ctx)
}
