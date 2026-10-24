// Code generated from RasterklangDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // RasterklangDsl
import "github.com/antlr4-go/antlr/v4"

// RasterklangDslListener is a complete listener for a parse tree produced by RasterklangDslParser.
type RasterklangDslListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterUnknownStmt is called when entering the unknownStmt production.
	EnterUnknownStmt(c *UnknownStmtContext)

	// EnterUnknownTail is called when entering the unknownTail production.
	EnterUnknownTail(c *UnknownTailContext)

	// EnterAnyIdent is called when entering the anyIdent production.
	EnterAnyIdent(c *AnyIdentContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterTempoStmt is called when entering the tempoStmt production.
	EnterTempoStmt(c *TempoStmtContext)

	// EnterTempoValue is called when entering the tempoValue production.
	EnterTempoValue(c *TempoValueContext)

	// EnterTempoSetting is called when entering the tempoSetting production.
	EnterTempoSetting(c *TempoSettingContext)

	// EnterTargetStmt is called when entering the targetStmt production.
	EnterTargetStmt(c *TargetStmtContext)

	// EnterChipStmt is called when entering the chipStmt production.
	EnterChipStmt(c *ChipStmtContext)

	// EnterTargetSetting is called when entering the targetSetting production.
	EnterTargetSetting(c *TargetSettingContext)

	// EnterPatternDecl is called when entering the patternDecl production.
	EnterPatternDecl(c *PatternDeclContext)

	// EnterPatternEach is called when entering the patternEach production.
	EnterPatternEach(c *PatternEachContext)

	// EnterArrangement is called when entering the arrangement production.
	EnterArrangement(c *ArrangementContext)

	// EnterArrangementElement is called when entering the arrangementElement production.
	EnterArrangementElement(c *ArrangementElementContext)

	// EnterTableDecl is called when entering the tableDecl production.
	EnterTableDecl(c *TableDeclContext)

	// EnterInstrumentBlock is called when entering the instrumentBlock production.
	EnterInstrumentBlock(c *InstrumentBlockContext)

	// EnterInstrumentField is called when entering the instrumentField production.
	EnterInstrumentField(c *InstrumentFieldContext)

	// EnterTableVerbLine is called when entering the tableVerbLine production.
	EnterTableVerbLine(c *TableVerbLineContext)

	// EnterTableRow is called when entering the tableRow production.
	EnterTableRow(c *TableRowContext)

	// EnterTableToken is called when entering the tableToken production.
	EnterTableToken(c *TableTokenContext)

	// EnterAbsNote is called when entering the absNote production.
	EnterAbsNote(c *AbsNoteContext)

	// EnterRampGenerator is called when entering the rampGenerator production.
	EnterRampGenerator(c *RampGeneratorContext)

	// EnterVoiceStmt is called when entering the voiceStmt production.
	EnterVoiceStmt(c *VoiceStmtContext)

	// EnterVoiceBody is called when entering the voiceBody production.
	EnterVoiceBody(c *VoiceBodyContext)

	// EnterVoiceShortTail is called when entering the voiceShortTail production.
	EnterVoiceShortTail(c *VoiceShortTailContext)

	// EnterInlineAttr is called when entering the inlineAttr production.
	EnterInlineAttr(c *InlineAttrContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterVoiceField is called when entering the voiceField production.
	EnterVoiceField(c *VoiceFieldContext)

	// EnterVoiceFieldValue is called when entering the voiceFieldValue production.
	EnterVoiceFieldValue(c *VoiceFieldValueContext)

	// EnterPipelineStep is called when entering the pipelineStep production.
	EnterPipelineStep(c *PipelineStepContext)

	// EnterTransform is called when entering the transform production.
	EnterTransform(c *TransformContext)

	// EnterTransformArg is called when entering the transformArg production.
	EnterTransformArg(c *TransformArgContext)

	// EnterTransformRef is called when entering the transformRef production.
	EnterTransformRef(c *TransformRefContext)

	// EnterFilterBlock is called when entering the filterBlock production.
	EnterFilterBlock(c *FilterBlockContext)

	// EnterFilterField is called when entering the filterField production.
	EnterFilterField(c *FilterFieldContext)

	// EnterFilterFieldValue is called when entering the filterFieldValue production.
	EnterFilterFieldValue(c *FilterFieldValueContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterSwitchCond is called when entering the switchCond production.
	EnterSwitchCond(c *SwitchCondContext)

	// EnterSignalRangeOpt is called when entering the signalRangeOpt production.
	EnterSignalRangeOpt(c *SignalRangeOptContext)

	// EnterSignalSpeedMod is called when entering the signalSpeedMod production.
	EnterSignalSpeedMod(c *SignalSpeedModContext)

	// EnterNumberList is called when entering the numberList production.
	EnterNumberList(c *NumberListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitUnknownStmt is called when exiting the unknownStmt production.
	ExitUnknownStmt(c *UnknownStmtContext)

	// ExitUnknownTail is called when exiting the unknownTail production.
	ExitUnknownTail(c *UnknownTailContext)

	// ExitAnyIdent is called when exiting the anyIdent production.
	ExitAnyIdent(c *AnyIdentContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitTempoStmt is called when exiting the tempoStmt production.
	ExitTempoStmt(c *TempoStmtContext)

	// ExitTempoValue is called when exiting the tempoValue production.
	ExitTempoValue(c *TempoValueContext)

	// ExitTempoSetting is called when exiting the tempoSetting production.
	ExitTempoSetting(c *TempoSettingContext)

	// ExitTargetStmt is called when exiting the targetStmt production.
	ExitTargetStmt(c *TargetStmtContext)

	// ExitChipStmt is called when exiting the chipStmt production.
	ExitChipStmt(c *ChipStmtContext)

	// ExitTargetSetting is called when exiting the targetSetting production.
	ExitTargetSetting(c *TargetSettingContext)

	// ExitPatternDecl is called when exiting the patternDecl production.
	ExitPatternDecl(c *PatternDeclContext)

	// ExitPatternEach is called when exiting the patternEach production.
	ExitPatternEach(c *PatternEachContext)

	// ExitArrangement is called when exiting the arrangement production.
	ExitArrangement(c *ArrangementContext)

	// ExitArrangementElement is called when exiting the arrangementElement production.
	ExitArrangementElement(c *ArrangementElementContext)

	// ExitTableDecl is called when exiting the tableDecl production.
	ExitTableDecl(c *TableDeclContext)

	// ExitInstrumentBlock is called when exiting the instrumentBlock production.
	ExitInstrumentBlock(c *InstrumentBlockContext)

	// ExitInstrumentField is called when exiting the instrumentField production.
	ExitInstrumentField(c *InstrumentFieldContext)

	// ExitTableVerbLine is called when exiting the tableVerbLine production.
	ExitTableVerbLine(c *TableVerbLineContext)

	// ExitTableRow is called when exiting the tableRow production.
	ExitTableRow(c *TableRowContext)

	// ExitTableToken is called when exiting the tableToken production.
	ExitTableToken(c *TableTokenContext)

	// ExitAbsNote is called when exiting the absNote production.
	ExitAbsNote(c *AbsNoteContext)

	// ExitRampGenerator is called when exiting the rampGenerator production.
	ExitRampGenerator(c *RampGeneratorContext)

	// ExitVoiceStmt is called when exiting the voiceStmt production.
	ExitVoiceStmt(c *VoiceStmtContext)

	// ExitVoiceBody is called when exiting the voiceBody production.
	ExitVoiceBody(c *VoiceBodyContext)

	// ExitVoiceShortTail is called when exiting the voiceShortTail production.
	ExitVoiceShortTail(c *VoiceShortTailContext)

	// ExitInlineAttr is called when exiting the inlineAttr production.
	ExitInlineAttr(c *InlineAttrContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitVoiceField is called when exiting the voiceField production.
	ExitVoiceField(c *VoiceFieldContext)

	// ExitVoiceFieldValue is called when exiting the voiceFieldValue production.
	ExitVoiceFieldValue(c *VoiceFieldValueContext)

	// ExitPipelineStep is called when exiting the pipelineStep production.
	ExitPipelineStep(c *PipelineStepContext)

	// ExitTransform is called when exiting the transform production.
	ExitTransform(c *TransformContext)

	// ExitTransformArg is called when exiting the transformArg production.
	ExitTransformArg(c *TransformArgContext)

	// ExitTransformRef is called when exiting the transformRef production.
	ExitTransformRef(c *TransformRefContext)

	// ExitFilterBlock is called when exiting the filterBlock production.
	ExitFilterBlock(c *FilterBlockContext)

	// ExitFilterField is called when exiting the filterField production.
	ExitFilterField(c *FilterFieldContext)

	// ExitFilterFieldValue is called when exiting the filterFieldValue production.
	ExitFilterFieldValue(c *FilterFieldValueContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitSwitchCond is called when exiting the switchCond production.
	ExitSwitchCond(c *SwitchCondContext)

	// ExitSignalRangeOpt is called when exiting the signalRangeOpt production.
	ExitSignalRangeOpt(c *SignalRangeOptContext)

	// ExitSignalSpeedMod is called when exiting the signalSpeedMod production.
	ExitSignalSpeedMod(c *SignalSpeedModContext)

	// ExitNumberList is called when exiting the numberList production.
	ExitNumberList(c *NumberListContext)
}
