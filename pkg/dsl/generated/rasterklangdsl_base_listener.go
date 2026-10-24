// Code generated from RasterklangDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // RasterklangDsl
import "github.com/antlr4-go/antlr/v4"

// BaseRasterklangDslListener is a complete listener for a parse tree produced by RasterklangDslParser.
type BaseRasterklangDslListener struct{}

var _ RasterklangDslListener = &BaseRasterklangDslListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRasterklangDslListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRasterklangDslListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRasterklangDslListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRasterklangDslListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRasterklangDslListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRasterklangDslListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRasterklangDslListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRasterklangDslListener) ExitStatement(ctx *StatementContext) {}

// EnterUnknownStmt is called when production unknownStmt is entered.
func (s *BaseRasterklangDslListener) EnterUnknownStmt(ctx *UnknownStmtContext) {}

// ExitUnknownStmt is called when production unknownStmt is exited.
func (s *BaseRasterklangDslListener) ExitUnknownStmt(ctx *UnknownStmtContext) {}

// EnterUnknownTail is called when production unknownTail is entered.
func (s *BaseRasterklangDslListener) EnterUnknownTail(ctx *UnknownTailContext) {}

// ExitUnknownTail is called when production unknownTail is exited.
func (s *BaseRasterklangDslListener) ExitUnknownTail(ctx *UnknownTailContext) {}

// EnterAnyIdent is called when production anyIdent is entered.
func (s *BaseRasterklangDslListener) EnterAnyIdent(ctx *AnyIdentContext) {}

// ExitAnyIdent is called when production anyIdent is exited.
func (s *BaseRasterklangDslListener) ExitAnyIdent(ctx *AnyIdentContext) {}

// EnterStr is called when production str is entered.
func (s *BaseRasterklangDslListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseRasterklangDslListener) ExitStr(ctx *StrContext) {}

// EnterTempoStmt is called when production tempoStmt is entered.
func (s *BaseRasterklangDslListener) EnterTempoStmt(ctx *TempoStmtContext) {}

// ExitTempoStmt is called when production tempoStmt is exited.
func (s *BaseRasterklangDslListener) ExitTempoStmt(ctx *TempoStmtContext) {}

// EnterTempoValue is called when production tempoValue is entered.
func (s *BaseRasterklangDslListener) EnterTempoValue(ctx *TempoValueContext) {}

// ExitTempoValue is called when production tempoValue is exited.
func (s *BaseRasterklangDslListener) ExitTempoValue(ctx *TempoValueContext) {}

// EnterTempoSetting is called when production tempoSetting is entered.
func (s *BaseRasterklangDslListener) EnterTempoSetting(ctx *TempoSettingContext) {}

// ExitTempoSetting is called when production tempoSetting is exited.
func (s *BaseRasterklangDslListener) ExitTempoSetting(ctx *TempoSettingContext) {}

// EnterTargetStmt is called when production targetStmt is entered.
func (s *BaseRasterklangDslListener) EnterTargetStmt(ctx *TargetStmtContext) {}

// ExitTargetStmt is called when production targetStmt is exited.
func (s *BaseRasterklangDslListener) ExitTargetStmt(ctx *TargetStmtContext) {}

// EnterChipStmt is called when production chipStmt is entered.
func (s *BaseRasterklangDslListener) EnterChipStmt(ctx *ChipStmtContext) {}

// ExitChipStmt is called when production chipStmt is exited.
func (s *BaseRasterklangDslListener) ExitChipStmt(ctx *ChipStmtContext) {}

// EnterTargetSetting is called when production targetSetting is entered.
func (s *BaseRasterklangDslListener) EnterTargetSetting(ctx *TargetSettingContext) {}

// ExitTargetSetting is called when production targetSetting is exited.
func (s *BaseRasterklangDslListener) ExitTargetSetting(ctx *TargetSettingContext) {}

// EnterPatternDecl is called when production patternDecl is entered.
func (s *BaseRasterklangDslListener) EnterPatternDecl(ctx *PatternDeclContext) {}

// ExitPatternDecl is called when production patternDecl is exited.
func (s *BaseRasterklangDslListener) ExitPatternDecl(ctx *PatternDeclContext) {}

// EnterPatternEach is called when production patternEach is entered.
func (s *BaseRasterklangDslListener) EnterPatternEach(ctx *PatternEachContext) {}

// ExitPatternEach is called when production patternEach is exited.
func (s *BaseRasterklangDslListener) ExitPatternEach(ctx *PatternEachContext) {}

// EnterArrangement is called when production arrangement is entered.
func (s *BaseRasterklangDslListener) EnterArrangement(ctx *ArrangementContext) {}

// ExitArrangement is called when production arrangement is exited.
func (s *BaseRasterklangDslListener) ExitArrangement(ctx *ArrangementContext) {}

// EnterArrangementElement is called when production arrangementElement is entered.
func (s *BaseRasterklangDslListener) EnterArrangementElement(ctx *ArrangementElementContext) {}

// ExitArrangementElement is called when production arrangementElement is exited.
func (s *BaseRasterklangDslListener) ExitArrangementElement(ctx *ArrangementElementContext) {}

// EnterTableDecl is called when production tableDecl is entered.
func (s *BaseRasterklangDslListener) EnterTableDecl(ctx *TableDeclContext) {}

// ExitTableDecl is called when production tableDecl is exited.
func (s *BaseRasterklangDslListener) ExitTableDecl(ctx *TableDeclContext) {}

// EnterInstrumentBlock is called when production instrumentBlock is entered.
func (s *BaseRasterklangDslListener) EnterInstrumentBlock(ctx *InstrumentBlockContext) {}

// ExitInstrumentBlock is called when production instrumentBlock is exited.
func (s *BaseRasterklangDslListener) ExitInstrumentBlock(ctx *InstrumentBlockContext) {}

// EnterInstrumentField is called when production instrumentField is entered.
func (s *BaseRasterklangDslListener) EnterInstrumentField(ctx *InstrumentFieldContext) {}

// ExitInstrumentField is called when production instrumentField is exited.
func (s *BaseRasterklangDslListener) ExitInstrumentField(ctx *InstrumentFieldContext) {}

// EnterTableVerbLine is called when production tableVerbLine is entered.
func (s *BaseRasterklangDslListener) EnterTableVerbLine(ctx *TableVerbLineContext) {}

// ExitTableVerbLine is called when production tableVerbLine is exited.
func (s *BaseRasterklangDslListener) ExitTableVerbLine(ctx *TableVerbLineContext) {}

// EnterTableRow is called when production tableRow is entered.
func (s *BaseRasterklangDslListener) EnterTableRow(ctx *TableRowContext) {}

// ExitTableRow is called when production tableRow is exited.
func (s *BaseRasterklangDslListener) ExitTableRow(ctx *TableRowContext) {}

// EnterTableToken is called when production tableToken is entered.
func (s *BaseRasterklangDslListener) EnterTableToken(ctx *TableTokenContext) {}

// ExitTableToken is called when production tableToken is exited.
func (s *BaseRasterklangDslListener) ExitTableToken(ctx *TableTokenContext) {}

// EnterAbsNote is called when production absNote is entered.
func (s *BaseRasterklangDslListener) EnterAbsNote(ctx *AbsNoteContext) {}

// ExitAbsNote is called when production absNote is exited.
func (s *BaseRasterklangDslListener) ExitAbsNote(ctx *AbsNoteContext) {}

// EnterRampGenerator is called when production rampGenerator is entered.
func (s *BaseRasterklangDslListener) EnterRampGenerator(ctx *RampGeneratorContext) {}

// ExitRampGenerator is called when production rampGenerator is exited.
func (s *BaseRasterklangDslListener) ExitRampGenerator(ctx *RampGeneratorContext) {}

// EnterVoiceStmt is called when production voiceStmt is entered.
func (s *BaseRasterklangDslListener) EnterVoiceStmt(ctx *VoiceStmtContext) {}

// ExitVoiceStmt is called when production voiceStmt is exited.
func (s *BaseRasterklangDslListener) ExitVoiceStmt(ctx *VoiceStmtContext) {}

// EnterVoiceBody is called when production voiceBody is entered.
func (s *BaseRasterklangDslListener) EnterVoiceBody(ctx *VoiceBodyContext) {}

// ExitVoiceBody is called when production voiceBody is exited.
func (s *BaseRasterklangDslListener) ExitVoiceBody(ctx *VoiceBodyContext) {}

// EnterVoiceShortTail is called when production voiceShortTail is entered.
func (s *BaseRasterklangDslListener) EnterVoiceShortTail(ctx *VoiceShortTailContext) {}

// ExitVoiceShortTail is called when production voiceShortTail is exited.
func (s *BaseRasterklangDslListener) ExitVoiceShortTail(ctx *VoiceShortTailContext) {}

// EnterInlineAttr is called when production inlineAttr is entered.
func (s *BaseRasterklangDslListener) EnterInlineAttr(ctx *InlineAttrContext) {}

// ExitInlineAttr is called when production inlineAttr is exited.
func (s *BaseRasterklangDslListener) ExitInlineAttr(ctx *InlineAttrContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseRasterklangDslListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseRasterklangDslListener) ExitValueList(ctx *ValueListContext) {}

// EnterVoiceField is called when production voiceField is entered.
func (s *BaseRasterklangDslListener) EnterVoiceField(ctx *VoiceFieldContext) {}

// ExitVoiceField is called when production voiceField is exited.
func (s *BaseRasterklangDslListener) ExitVoiceField(ctx *VoiceFieldContext) {}

// EnterVoiceFieldValue is called when production voiceFieldValue is entered.
func (s *BaseRasterklangDslListener) EnterVoiceFieldValue(ctx *VoiceFieldValueContext) {}

// ExitVoiceFieldValue is called when production voiceFieldValue is exited.
func (s *BaseRasterklangDslListener) ExitVoiceFieldValue(ctx *VoiceFieldValueContext) {}

// EnterPipelineStep is called when production pipelineStep is entered.
func (s *BaseRasterklangDslListener) EnterPipelineStep(ctx *PipelineStepContext) {}

// ExitPipelineStep is called when production pipelineStep is exited.
func (s *BaseRasterklangDslListener) ExitPipelineStep(ctx *PipelineStepContext) {}

// EnterTransform is called when production transform is entered.
func (s *BaseRasterklangDslListener) EnterTransform(ctx *TransformContext) {}

// ExitTransform is called when production transform is exited.
func (s *BaseRasterklangDslListener) ExitTransform(ctx *TransformContext) {}

// EnterTransformArg is called when production transformArg is entered.
func (s *BaseRasterklangDslListener) EnterTransformArg(ctx *TransformArgContext) {}

// ExitTransformArg is called when production transformArg is exited.
func (s *BaseRasterklangDslListener) ExitTransformArg(ctx *TransformArgContext) {}

// EnterTransformRef is called when production transformRef is entered.
func (s *BaseRasterklangDslListener) EnterTransformRef(ctx *TransformRefContext) {}

// ExitTransformRef is called when production transformRef is exited.
func (s *BaseRasterklangDslListener) ExitTransformRef(ctx *TransformRefContext) {}

// EnterFilterBlock is called when production filterBlock is entered.
func (s *BaseRasterklangDslListener) EnterFilterBlock(ctx *FilterBlockContext) {}

// ExitFilterBlock is called when production filterBlock is exited.
func (s *BaseRasterklangDslListener) ExitFilterBlock(ctx *FilterBlockContext) {}

// EnterFilterField is called when production filterField is entered.
func (s *BaseRasterklangDslListener) EnterFilterField(ctx *FilterFieldContext) {}

// ExitFilterField is called when production filterField is exited.
func (s *BaseRasterklangDslListener) ExitFilterField(ctx *FilterFieldContext) {}

// EnterFilterFieldValue is called when production filterFieldValue is entered.
func (s *BaseRasterklangDslListener) EnterFilterFieldValue(ctx *FilterFieldValueContext) {}

// ExitFilterFieldValue is called when production filterFieldValue is exited.
func (s *BaseRasterklangDslListener) ExitFilterFieldValue(ctx *FilterFieldValueContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseRasterklangDslListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseRasterklangDslListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterSwitchCond is called when production switchCond is entered.
func (s *BaseRasterklangDslListener) EnterSwitchCond(ctx *SwitchCondContext) {}

// ExitSwitchCond is called when production switchCond is exited.
func (s *BaseRasterklangDslListener) ExitSwitchCond(ctx *SwitchCondContext) {}

// EnterSignalRangeOpt is called when production signalRangeOpt is entered.
func (s *BaseRasterklangDslListener) EnterSignalRangeOpt(ctx *SignalRangeOptContext) {}

// ExitSignalRangeOpt is called when production signalRangeOpt is exited.
func (s *BaseRasterklangDslListener) ExitSignalRangeOpt(ctx *SignalRangeOptContext) {}

// EnterSignalSpeedMod is called when production signalSpeedMod is entered.
func (s *BaseRasterklangDslListener) EnterSignalSpeedMod(ctx *SignalSpeedModContext) {}

// ExitSignalSpeedMod is called when production signalSpeedMod is exited.
func (s *BaseRasterklangDslListener) ExitSignalSpeedMod(ctx *SignalSpeedModContext) {}

// EnterNumberList is called when production numberList is entered.
func (s *BaseRasterklangDslListener) EnterNumberList(ctx *NumberListContext) {}

// ExitNumberList is called when production numberList is exited.
func (s *BaseRasterklangDslListener) ExitNumberList(ctx *NumberListContext) {}
