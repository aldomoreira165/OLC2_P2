// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockFunc is called when production blockFunc is entered.
func (s *BaseSwiftGrammarListener) EnterBlockFunc(ctx *BlockFuncContext) {}

// ExitBlockFunc is called when production blockFunc is exited.
func (s *BaseSwiftGrammarListener) ExitBlockFunc(ctx *BlockFuncContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseSwiftGrammarListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseSwiftGrammarListener) ExitStmt(ctx *StmtContext) {}

// EnterDefStruct is called when production DefStruct is entered.
func (s *BaseSwiftGrammarListener) EnterDefStruct(ctx *DefStructContext) {}

// ExitDefStruct is called when production DefStruct is exited.
func (s *BaseSwiftGrammarListener) ExitDefStruct(ctx *DefStructContext) {}

// EnterAtributoStruct is called when production AtributoStruct is entered.
func (s *BaseSwiftGrammarListener) EnterAtributoStruct(ctx *AtributoStructContext) {}

// ExitAtributoStruct is called when production AtributoStruct is exited.
func (s *BaseSwiftGrammarListener) ExitAtributoStruct(ctx *AtributoStructContext) {}

// EnterStructExpr is called when production StructExpr is entered.
func (s *BaseSwiftGrammarListener) EnterStructExpr(ctx *StructExprContext) {}

// ExitStructExpr is called when production StructExpr is exited.
func (s *BaseSwiftGrammarListener) ExitStructExpr(ctx *StructExprContext) {}

// EnterStructExprID is called when production StructExprID is entered.
func (s *BaseSwiftGrammarListener) EnterStructExprID(ctx *StructExprIDContext) {}

// ExitStructExprID is called when production StructExprID is exited.
func (s *BaseSwiftGrammarListener) ExitStructExprID(ctx *StructExprIDContext) {}

// EnterValorStructExpr is called when production ValorStructExpr is entered.
func (s *BaseSwiftGrammarListener) EnterValorStructExpr(ctx *ValorStructExprContext) {}

// ExitValorStructExpr is called when production ValorStructExpr is exited.
func (s *BaseSwiftGrammarListener) ExitValorStructExpr(ctx *ValorStructExprContext) {}

// EnterDuplastruct is called when production Duplastruct is entered.
func (s *BaseSwiftGrammarListener) EnterDuplastruct(ctx *DuplastructContext) {}

// ExitDuplastruct is called when production Duplastruct is exited.
func (s *BaseSwiftGrammarListener) ExitDuplastruct(ctx *DuplastructContext) {}

// EnterAccesoStruct is called when production AccesoStruct is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoStruct(ctx *AccesoStructContext) {}

// ExitAccesoStruct is called when production AccesoStruct is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoStruct(ctx *AccesoStructContext) {}

// EnterAsignStruct is called when production AsignStruct is entered.
func (s *BaseSwiftGrammarListener) EnterAsignStruct(ctx *AsignStructContext) {}

// ExitAsignStruct is called when production AsignStruct is exited.
func (s *BaseSwiftGrammarListener) ExitAsignStruct(ctx *AsignStructContext) {}

// EnterVectorStruct is called when production VectorStruct is entered.
func (s *BaseSwiftGrammarListener) EnterVectorStruct(ctx *VectorStructContext) {}

// ExitVectorStruct is called when production VectorStruct is exited.
func (s *BaseSwiftGrammarListener) ExitVectorStruct(ctx *VectorStructContext) {}

// EnterAccesoVectorStruct is called when production AccesoVectorStruct is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoVectorStruct(ctx *AccesoVectorStructContext) {}

// ExitAccesoVectorStruct is called when production AccesoVectorStruct is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoVectorStruct(ctx *AccesoVectorStructContext) {}

// EnterAppendVectorStruct is called when production AppendVectorStruct is entered.
func (s *BaseSwiftGrammarListener) EnterAppendVectorStruct(ctx *AppendVectorStructContext) {}

// ExitAppendVectorStruct is called when production AppendVectorStruct is exited.
func (s *BaseSwiftGrammarListener) ExitAppendVectorStruct(ctx *AppendVectorStructContext) {}

// EnterDeclvectorstmt is called when production declvectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterDeclvectorstmt(ctx *DeclvectorstmtContext) {}

// ExitDeclvectorstmt is called when production declvectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitDeclvectorstmt(ctx *DeclvectorstmtContext) {}

// EnterDefVector is called when production DefVector is entered.
func (s *BaseSwiftGrammarListener) EnterDefVector(ctx *DefVectorContext) {}

// ExitDefVector is called when production DefVector is exited.
func (s *BaseSwiftGrammarListener) ExitDefVector(ctx *DefVectorContext) {}

// EnterDefVectorID is called when production DefVectorID is entered.
func (s *BaseSwiftGrammarListener) EnterDefVectorID(ctx *DefVectorIDContext) {}

// ExitDefVectorID is called when production DefVectorID is exited.
func (s *BaseSwiftGrammarListener) ExitDefVectorID(ctx *DefVectorIDContext) {}

// EnterListaexpresiones is called when production listaexpresiones is entered.
func (s *BaseSwiftGrammarListener) EnterListaexpresiones(ctx *ListaexpresionesContext) {}

// ExitListaexpresiones is called when production listaexpresiones is exited.
func (s *BaseSwiftGrammarListener) ExitListaexpresiones(ctx *ListaexpresionesContext) {}

// EnterAccesovectorstmt is called when production accesovectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAccesovectorstmt(ctx *AccesovectorstmtContext) {}

// ExitAccesovectorstmt is called when production accesovectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAccesovectorstmt(ctx *AccesovectorstmtContext) {}

// EnterAsignvectorstmt is called when production asignvectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAsignvectorstmt(ctx *AsignvectorstmtContext) {}

// ExitAsignvectorstmt is called when production asignvectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAsignvectorstmt(ctx *AsignvectorstmtContext) {}

// EnterAppendvectorstmt is called when production appendvectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAppendvectorstmt(ctx *AppendvectorstmtContext) {}

// ExitAppendvectorstmt is called when production appendvectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAppendvectorstmt(ctx *AppendvectorstmtContext) {}

// EnterRemoveatvectorstmt is called when production removeatvectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterRemoveatvectorstmt(ctx *RemoveatvectorstmtContext) {}

// ExitRemoveatvectorstmt is called when production removeatvectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitRemoveatvectorstmt(ctx *RemoveatvectorstmtContext) {}

// EnterRemovelastvectorstmt is called when production removelastvectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterRemovelastvectorstmt(ctx *RemovelastvectorstmtContext) {}

// ExitRemovelastvectorstmt is called when production removelastvectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitRemovelastvectorstmt(ctx *RemovelastvectorstmtContext) {}

// EnterCountvectorstmt is called when production countvectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterCountvectorstmt(ctx *CountvectorstmtContext) {}

// ExitCountvectorstmt is called when production countvectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitCountvectorstmt(ctx *CountvectorstmtContext) {}

// EnterIsemptyvectorstmt is called when production isemptyvectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIsemptyvectorstmt(ctx *IsemptyvectorstmtContext) {}

// ExitIsemptyvectorstmt is called when production isemptyvectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIsemptyvectorstmt(ctx *IsemptyvectorstmtContext) {}

// EnterDeclmatrizstmt2 is called when production declmatrizstmt2 is entered.
func (s *BaseSwiftGrammarListener) EnterDeclmatrizstmt2(ctx *Declmatrizstmt2Context) {}

// ExitDeclmatrizstmt2 is called when production declmatrizstmt2 is exited.
func (s *BaseSwiftGrammarListener) ExitDeclmatrizstmt2(ctx *Declmatrizstmt2Context) {}

// EnterDeclmatrizstmt3 is called when production declmatrizstmt3 is entered.
func (s *BaseSwiftGrammarListener) EnterDeclmatrizstmt3(ctx *Declmatrizstmt3Context) {}

// ExitDeclmatrizstmt3 is called when production declmatrizstmt3 is exited.
func (s *BaseSwiftGrammarListener) ExitDeclmatrizstmt3(ctx *Declmatrizstmt3Context) {}

// EnterTipomatriz3 is called when production tipomatriz3 is entered.
func (s *BaseSwiftGrammarListener) EnterTipomatriz3(ctx *Tipomatriz3Context) {}

// ExitTipomatriz3 is called when production tipomatriz3 is exited.
func (s *BaseSwiftGrammarListener) ExitTipomatriz3(ctx *Tipomatriz3Context) {}

// EnterTipomatriz2 is called when production tipomatriz2 is entered.
func (s *BaseSwiftGrammarListener) EnterTipomatriz2(ctx *Tipomatriz2Context) {}

// ExitTipomatriz2 is called when production tipomatriz2 is exited.
func (s *BaseSwiftGrammarListener) ExitTipomatriz2(ctx *Tipomatriz2Context) {}

// EnterListavaloresmatriz2 is called when production listavaloresmatriz2 is entered.
func (s *BaseSwiftGrammarListener) EnterListavaloresmatriz2(ctx *Listavaloresmatriz2Context) {}

// ExitListavaloresmatriz2 is called when production listavaloresmatriz2 is exited.
func (s *BaseSwiftGrammarListener) ExitListavaloresmatriz2(ctx *Listavaloresmatriz2Context) {}

// EnterListavaloresmatriz3 is called when production listavaloresmatriz3 is entered.
func (s *BaseSwiftGrammarListener) EnterListavaloresmatriz3(ctx *Listavaloresmatriz3Context) {}

// ExitListavaloresmatriz3 is called when production listavaloresmatriz3 is exited.
func (s *BaseSwiftGrammarListener) ExitListavaloresmatriz3(ctx *Listavaloresmatriz3Context) {}

// EnterAccesomatriz2 is called when production accesomatriz2 is entered.
func (s *BaseSwiftGrammarListener) EnterAccesomatriz2(ctx *Accesomatriz2Context) {}

// ExitAccesomatriz2 is called when production accesomatriz2 is exited.
func (s *BaseSwiftGrammarListener) ExitAccesomatriz2(ctx *Accesomatriz2Context) {}

// EnterAccesomatriz3 is called when production accesomatriz3 is entered.
func (s *BaseSwiftGrammarListener) EnterAccesomatriz3(ctx *Accesomatriz3Context) {}

// ExitAccesomatriz3 is called when production accesomatriz3 is exited.
func (s *BaseSwiftGrammarListener) ExitAccesomatriz3(ctx *Accesomatriz3Context) {}

// EnterAsignmatrizstmt2 is called when production asignmatrizstmt2 is entered.
func (s *BaseSwiftGrammarListener) EnterAsignmatrizstmt2(ctx *Asignmatrizstmt2Context) {}

// ExitAsignmatrizstmt2 is called when production asignmatrizstmt2 is exited.
func (s *BaseSwiftGrammarListener) ExitAsignmatrizstmt2(ctx *Asignmatrizstmt2Context) {}

// EnterAsignmatrizstmt3 is called when production asignmatrizstmt3 is entered.
func (s *BaseSwiftGrammarListener) EnterAsignmatrizstmt3(ctx *Asignmatrizstmt3Context) {}

// ExitAsignmatrizstmt3 is called when production asignmatrizstmt3 is exited.
func (s *BaseSwiftGrammarListener) ExitAsignmatrizstmt3(ctx *Asignmatrizstmt3Context) {}

// EnterReturnstmt is called when production returnstmt is entered.
func (s *BaseSwiftGrammarListener) EnterReturnstmt(ctx *ReturnstmtContext) {}

// ExitReturnstmt is called when production returnstmt is exited.
func (s *BaseSwiftGrammarListener) ExitReturnstmt(ctx *ReturnstmtContext) {}

// EnterBreakstmt is called when production breakstmt is entered.
func (s *BaseSwiftGrammarListener) EnterBreakstmt(ctx *BreakstmtContext) {}

// ExitBreakstmt is called when production breakstmt is exited.
func (s *BaseSwiftGrammarListener) ExitBreakstmt(ctx *BreakstmtContext) {}

// EnterContinuestmt is called when production continuestmt is entered.
func (s *BaseSwiftGrammarListener) EnterContinuestmt(ctx *ContinuestmtContext) {}

// ExitContinuestmt is called when production continuestmt is exited.
func (s *BaseSwiftGrammarListener) ExitContinuestmt(ctx *ContinuestmtContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterIntstmt is called when production intstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIntstmt(ctx *IntstmtContext) {}

// ExitIntstmt is called when production intstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIntstmt(ctx *IntstmtContext) {}

// EnterFloatstmt is called when production floatstmt is entered.
func (s *BaseSwiftGrammarListener) EnterFloatstmt(ctx *FloatstmtContext) {}

// ExitFloatstmt is called when production floatstmt is exited.
func (s *BaseSwiftGrammarListener) ExitFloatstmt(ctx *FloatstmtContext) {}

// EnterStringstmt is called when production stringstmt is entered.
func (s *BaseSwiftGrammarListener) EnterStringstmt(ctx *StringstmtContext) {}

// ExitStringstmt is called when production stringstmt is exited.
func (s *BaseSwiftGrammarListener) ExitStringstmt(ctx *StringstmtContext) {}

// EnterFuncionNormal is called when production FuncionNormal is entered.
func (s *BaseSwiftGrammarListener) EnterFuncionNormal(ctx *FuncionNormalContext) {}

// ExitFuncionNormal is called when production FuncionNormal is exited.
func (s *BaseSwiftGrammarListener) ExitFuncionNormal(ctx *FuncionNormalContext) {}

// EnterFuncionRetorno is called when production FuncionRetorno is entered.
func (s *BaseSwiftGrammarListener) EnterFuncionRetorno(ctx *FuncionRetornoContext) {}

// ExitFuncionRetorno is called when production FuncionRetorno is exited.
func (s *BaseSwiftGrammarListener) ExitFuncionRetorno(ctx *FuncionRetornoContext) {}

// EnterAccfuncstm is called when production accfuncstm is entered.
func (s *BaseSwiftGrammarListener) EnterAccfuncstm(ctx *AccfuncstmContext) {}

// ExitAccfuncstm is called when production accfuncstm is exited.
func (s *BaseSwiftGrammarListener) ExitAccfuncstm(ctx *AccfuncstmContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseSwiftGrammarListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseSwiftGrammarListener) ExitParametros(ctx *ParametrosContext) {}

// EnterParametroscall is called when production parametroscall is entered.
func (s *BaseSwiftGrammarListener) EnterParametroscall(ctx *ParametroscallContext) {}

// ExitParametroscall is called when production parametroscall is exited.
func (s *BaseSwiftGrammarListener) ExitParametroscall(ctx *ParametroscallContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterElseifstmt is called when production elseifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterElseifstmt(ctx *ElseifstmtContext) {}

// ExitElseifstmt is called when production elseifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitElseifstmt(ctx *ElseifstmtContext) {}

// EnterSwitchstmt is called when production switchstmt is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchstmt(ctx *SwitchstmtContext) {}

// ExitSwitchstmt is called when production switchstmt is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchstmt(ctx *SwitchstmtContext) {}

// EnterCaseStmt is called when production caseStmt is entered.
func (s *BaseSwiftGrammarListener) EnterCaseStmt(ctx *CaseStmtContext) {}

// ExitCaseStmt is called when production caseStmt is exited.
func (s *BaseSwiftGrammarListener) ExitCaseStmt(ctx *CaseStmtContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseSwiftGrammarListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseSwiftGrammarListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterTypedDeclstmt is called when production typedDeclstmt is entered.
func (s *BaseSwiftGrammarListener) EnterTypedDeclstmt(ctx *TypedDeclstmtContext) {}

// ExitTypedDeclstmt is called when production typedDeclstmt is exited.
func (s *BaseSwiftGrammarListener) ExitTypedDeclstmt(ctx *TypedDeclstmtContext) {}

// EnterUntypedDeclstmt is called when production untypedDeclstmt is entered.
func (s *BaseSwiftGrammarListener) EnterUntypedDeclstmt(ctx *UntypedDeclstmtContext) {}

// ExitUntypedDeclstmt is called when production untypedDeclstmt is exited.
func (s *BaseSwiftGrammarListener) ExitUntypedDeclstmt(ctx *UntypedDeclstmtContext) {}

// EnterOptionalTypedDeclstmt is called when production optionalTypedDeclstmt is entered.
func (s *BaseSwiftGrammarListener) EnterOptionalTypedDeclstmt(ctx *OptionalTypedDeclstmtContext) {}

// ExitOptionalTypedDeclstmt is called when production optionalTypedDeclstmt is exited.
func (s *BaseSwiftGrammarListener) ExitOptionalTypedDeclstmt(ctx *OptionalTypedDeclstmtContext) {}

// EnterAsignstmt is called when production asignstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAsignstmt(ctx *AsignstmtContext) {}

// ExitAsignstmt is called when production asignstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAsignstmt(ctx *AsignstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseSwiftGrammarListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseSwiftGrammarListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterForstmt is called when production forstmt is entered.
func (s *BaseSwiftGrammarListener) EnterForstmt(ctx *ForstmtContext) {}

// ExitForstmt is called when production forstmt is exited.
func (s *BaseSwiftGrammarListener) ExitForstmt(ctx *ForstmtContext) {}

// EnterGuardstmt is called when production guardstmt is entered.
func (s *BaseSwiftGrammarListener) EnterGuardstmt(ctx *GuardstmtContext) {}

// ExitGuardstmt is called when production guardstmt is exited.
func (s *BaseSwiftGrammarListener) ExitGuardstmt(ctx *GuardstmtContext) {}

// EnterRangostmt is called when production rangostmt is entered.
func (s *BaseSwiftGrammarListener) EnterRangostmt(ctx *RangostmtContext) {}

// ExitRangostmt is called when production rangostmt is exited.
func (s *BaseSwiftGrammarListener) ExitRangostmt(ctx *RangostmtContext) {}

// EnterIncremento is called when production Incremento is entered.
func (s *BaseSwiftGrammarListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production Incremento is exited.
func (s *BaseSwiftGrammarListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production Decremento is entered.
func (s *BaseSwiftGrammarListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production Decremento is exited.
func (s *BaseSwiftGrammarListener) ExitDecremento(ctx *DecrementoContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseSwiftGrammarListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseSwiftGrammarListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterStringExpr is called when production StringExpr is entered.
func (s *BaseSwiftGrammarListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production StringExpr is exited.
func (s *BaseSwiftGrammarListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterNilExpr is called when production NilExpr is entered.
func (s *BaseSwiftGrammarListener) EnterNilExpr(ctx *NilExprContext) {}

// ExitNilExpr is called when production NilExpr is exited.
func (s *BaseSwiftGrammarListener) ExitNilExpr(ctx *NilExprContext) {}

// EnterFloatExpr is called when production FloatExpr is entered.
func (s *BaseSwiftGrammarListener) EnterFloatExpr(ctx *FloatExprContext) {}

// ExitFloatExpr is called when production FloatExpr is exited.
func (s *BaseSwiftGrammarListener) ExitFloatExpr(ctx *FloatExprContext) {}

// EnterIsEmptyVectorExpr is called when production IsEmptyVectorExpr is entered.
func (s *BaseSwiftGrammarListener) EnterIsEmptyVectorExpr(ctx *IsEmptyVectorExprContext) {}

// ExitIsEmptyVectorExpr is called when production IsEmptyVectorExpr is exited.
func (s *BaseSwiftGrammarListener) ExitIsEmptyVectorExpr(ctx *IsEmptyVectorExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseSwiftGrammarListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseSwiftGrammarListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterOpExpr is called when production OpExpr is entered.
func (s *BaseSwiftGrammarListener) EnterOpExpr(ctx *OpExprContext) {}

// ExitOpExpr is called when production OpExpr is exited.
func (s *BaseSwiftGrammarListener) ExitOpExpr(ctx *OpExprContext) {}

// EnterAccesoVectorExpr is called when production AccesoVectorExpr is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoVectorExpr(ctx *AccesoVectorExprContext) {}

// ExitAccesoVectorExpr is called when production AccesoVectorExpr is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoVectorExpr(ctx *AccesoVectorExprContext) {}

// EnterUnariaExpr is called when production UnariaExpr is entered.
func (s *BaseSwiftGrammarListener) EnterUnariaExpr(ctx *UnariaExprContext) {}

// ExitUnariaExpr is called when production UnariaExpr is exited.
func (s *BaseSwiftGrammarListener) ExitUnariaExpr(ctx *UnariaExprContext) {}

// EnterAccesoMatrizExpr is called when production AccesoMatrizExpr is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoMatrizExpr(ctx *AccesoMatrizExprContext) {}

// ExitAccesoMatrizExpr is called when production AccesoMatrizExpr is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoMatrizExpr(ctx *AccesoMatrizExprContext) {}

// EnterCountVectorExpr is called when production CountVectorExpr is entered.
func (s *BaseSwiftGrammarListener) EnterCountVectorExpr(ctx *CountVectorExprContext) {}

// ExitCountVectorExpr is called when production CountVectorExpr is exited.
func (s *BaseSwiftGrammarListener) ExitCountVectorExpr(ctx *CountVectorExprContext) {}

// EnterNumExpr is called when production NumExpr is entered.
func (s *BaseSwiftGrammarListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production NumExpr is exited.
func (s *BaseSwiftGrammarListener) ExitNumExpr(ctx *NumExprContext) {}

// EnterParExpr is called when production ParExpr is entered.
func (s *BaseSwiftGrammarListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production ParExpr is exited.
func (s *BaseSwiftGrammarListener) ExitParExpr(ctx *ParExprContext) {}

// EnterStrExpr is called when production StrExpr is entered.
func (s *BaseSwiftGrammarListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production StrExpr is exited.
func (s *BaseSwiftGrammarListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterAccFuncExpr is called when production AccFuncExpr is entered.
func (s *BaseSwiftGrammarListener) EnterAccFuncExpr(ctx *AccFuncExprContext) {}

// ExitAccFuncExpr is called when production AccFuncExpr is exited.
func (s *BaseSwiftGrammarListener) ExitAccFuncExpr(ctx *AccFuncExprContext) {}

// EnterAccesoValorStructExpr is called when production AccesoValorStructExpr is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoValorStructExpr(ctx *AccesoValorStructExprContext) {}

// ExitAccesoValorStructExpr is called when production AccesoValorStructExpr is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoValorStructExpr(ctx *AccesoValorStructExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseSwiftGrammarListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseSwiftGrammarListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseSwiftGrammarListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseSwiftGrammarListener) ExitIntExpr(ctx *IntExprContext) {}

// EnterAccesoVectorStructExpr is called when production AccesoVectorStructExpr is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoVectorStructExpr(ctx *AccesoVectorStructExprContext) {}

// ExitAccesoVectorStructExpr is called when production AccesoVectorStructExpr is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoVectorStructExpr(ctx *AccesoVectorStructExprContext) {}

// EnterAccesoStructExpr is called when production AccesoStructExpr is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoStructExpr(ctx *AccesoStructExprContext) {}

// ExitAccesoStructExpr is called when production AccesoStructExpr is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoStructExpr(ctx *AccesoStructExprContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseSwiftGrammarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseSwiftGrammarListener) ExitTipo(ctx *TipoContext) {}

// EnterTipo_vector is called when production tipo_vector is entered.
func (s *BaseSwiftGrammarListener) EnterTipo_vector(ctx *Tipo_vectorContext) {}

// ExitTipo_vector is called when production tipo_vector is exited.
func (s *BaseSwiftGrammarListener) ExitTipo_vector(ctx *Tipo_vectorContext) {}

// EnterTipo_matriz2 is called when production tipo_matriz2 is entered.
func (s *BaseSwiftGrammarListener) EnterTipo_matriz2(ctx *Tipo_matriz2Context) {}

// ExitTipo_matriz2 is called when production tipo_matriz2 is exited.
func (s *BaseSwiftGrammarListener) ExitTipo_matriz2(ctx *Tipo_matriz2Context) {}

// EnterTipo_matriz3 is called when production tipo_matriz3 is entered.
func (s *BaseSwiftGrammarListener) EnterTipo_matriz3(ctx *Tipo_matriz3Context) {}

// ExitTipo_matriz3 is called when production tipo_matriz3 is exited.
func (s *BaseSwiftGrammarListener) ExitTipo_matriz3(ctx *Tipo_matriz3Context) {}
