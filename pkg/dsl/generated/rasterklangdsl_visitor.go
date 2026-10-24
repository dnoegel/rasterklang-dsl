// Code generated from RasterklangDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // RasterklangDsl
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RasterklangDslParser.
type RasterklangDslVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RasterklangDslParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#unknownStmt.
	VisitUnknownStmt(ctx *UnknownStmtContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#unknownTail.
	VisitUnknownTail(ctx *UnknownTailContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#anyIdent.
	VisitAnyIdent(ctx *AnyIdentContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#tempoStmt.
	VisitTempoStmt(ctx *TempoStmtContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#tempoValue.
	VisitTempoValue(ctx *TempoValueContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#tempoSetting.
	VisitTempoSetting(ctx *TempoSettingContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#targetStmt.
	VisitTargetStmt(ctx *TargetStmtContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#chipStmt.
	VisitChipStmt(ctx *ChipStmtContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#targetSetting.
	VisitTargetSetting(ctx *TargetSettingContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#patternDecl.
	VisitPatternDecl(ctx *PatternDeclContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#patternEach.
	VisitPatternEach(ctx *PatternEachContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#arrangement.
	VisitArrangement(ctx *ArrangementContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#arrangementElement.
	VisitArrangementElement(ctx *ArrangementElementContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#tableDecl.
	VisitTableDecl(ctx *TableDeclContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#instrumentBlock.
	VisitInstrumentBlock(ctx *InstrumentBlockContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#instrumentField.
	VisitInstrumentField(ctx *InstrumentFieldContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#tableVerbLine.
	VisitTableVerbLine(ctx *TableVerbLineContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#tableRow.
	VisitTableRow(ctx *TableRowContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#tableToken.
	VisitTableToken(ctx *TableTokenContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#absNote.
	VisitAbsNote(ctx *AbsNoteContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#rampGenerator.
	VisitRampGenerator(ctx *RampGeneratorContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#voiceStmt.
	VisitVoiceStmt(ctx *VoiceStmtContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#voiceBody.
	VisitVoiceBody(ctx *VoiceBodyContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#voiceShortTail.
	VisitVoiceShortTail(ctx *VoiceShortTailContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#inlineAttr.
	VisitInlineAttr(ctx *InlineAttrContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#voiceField.
	VisitVoiceField(ctx *VoiceFieldContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#voiceFieldValue.
	VisitVoiceFieldValue(ctx *VoiceFieldValueContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#pipelineStep.
	VisitPipelineStep(ctx *PipelineStepContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#transform.
	VisitTransform(ctx *TransformContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#transformArg.
	VisitTransformArg(ctx *TransformArgContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#transformRef.
	VisitTransformRef(ctx *TransformRefContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#filterBlock.
	VisitFilterBlock(ctx *FilterBlockContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#filterField.
	VisitFilterField(ctx *FilterFieldContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#filterFieldValue.
	VisitFilterFieldValue(ctx *FilterFieldValueContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#switchCond.
	VisitSwitchCond(ctx *SwitchCondContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#signalRangeOpt.
	VisitSignalRangeOpt(ctx *SignalRangeOptContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#signalSpeedMod.
	VisitSignalSpeedMod(ctx *SignalSpeedModContext) interface{}

	// Visit a parse tree produced by RasterklangDslParser#numberList.
	VisitNumberList(ctx *NumberListContext) interface{}
}
