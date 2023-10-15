// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockFunc is called when entering the blockFunc production.
	EnterBlockFunc(c *BlockFuncContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDefStruct is called when entering the DefStruct production.
	EnterDefStruct(c *DefStructContext)

	// EnterAtributoStruct is called when entering the AtributoStruct production.
	EnterAtributoStruct(c *AtributoStructContext)

	// EnterStructExpr is called when entering the StructExpr production.
	EnterStructExpr(c *StructExprContext)

	// EnterStructExprID is called when entering the StructExprID production.
	EnterStructExprID(c *StructExprIDContext)

	// EnterValorStructExpr is called when entering the ValorStructExpr production.
	EnterValorStructExpr(c *ValorStructExprContext)

	// EnterDuplastruct is called when entering the Duplastruct production.
	EnterDuplastruct(c *DuplastructContext)

	// EnterAccesoStruct is called when entering the AccesoStruct production.
	EnterAccesoStruct(c *AccesoStructContext)

	// EnterAsignStruct is called when entering the AsignStruct production.
	EnterAsignStruct(c *AsignStructContext)

	// EnterVectorStruct is called when entering the VectorStruct production.
	EnterVectorStruct(c *VectorStructContext)

	// EnterAccesoVectorStruct is called when entering the AccesoVectorStruct production.
	EnterAccesoVectorStruct(c *AccesoVectorStructContext)

	// EnterAppendVectorStruct is called when entering the AppendVectorStruct production.
	EnterAppendVectorStruct(c *AppendVectorStructContext)

	// EnterDeclvectorstmt is called when entering the declvectorstmt production.
	EnterDeclvectorstmt(c *DeclvectorstmtContext)

	// EnterDefVector is called when entering the DefVector production.
	EnterDefVector(c *DefVectorContext)

	// EnterDefVectorID is called when entering the DefVectorID production.
	EnterDefVectorID(c *DefVectorIDContext)

	// EnterListaexpresiones is called when entering the listaexpresiones production.
	EnterListaexpresiones(c *ListaexpresionesContext)

	// EnterAccesovectorstmt is called when entering the accesovectorstmt production.
	EnterAccesovectorstmt(c *AccesovectorstmtContext)

	// EnterAsignvectorstmt is called when entering the asignvectorstmt production.
	EnterAsignvectorstmt(c *AsignvectorstmtContext)

	// EnterAppendvectorstmt is called when entering the appendvectorstmt production.
	EnterAppendvectorstmt(c *AppendvectorstmtContext)

	// EnterRemoveatvectorstmt is called when entering the removeatvectorstmt production.
	EnterRemoveatvectorstmt(c *RemoveatvectorstmtContext)

	// EnterRemovelastvectorstmt is called when entering the removelastvectorstmt production.
	EnterRemovelastvectorstmt(c *RemovelastvectorstmtContext)

	// EnterCountvectorstmt is called when entering the countvectorstmt production.
	EnterCountvectorstmt(c *CountvectorstmtContext)

	// EnterIsemptyvectorstmt is called when entering the isemptyvectorstmt production.
	EnterIsemptyvectorstmt(c *IsemptyvectorstmtContext)

	// EnterDeclmatrizstmt2 is called when entering the declmatrizstmt2 production.
	EnterDeclmatrizstmt2(c *Declmatrizstmt2Context)

	// EnterDeclmatrizstmt3 is called when entering the declmatrizstmt3 production.
	EnterDeclmatrizstmt3(c *Declmatrizstmt3Context)

	// EnterTipomatriz3 is called when entering the tipomatriz3 production.
	EnterTipomatriz3(c *Tipomatriz3Context)

	// EnterTipomatriz2 is called when entering the tipomatriz2 production.
	EnterTipomatriz2(c *Tipomatriz2Context)

	// EnterListavaloresmatriz2 is called when entering the listavaloresmatriz2 production.
	EnterListavaloresmatriz2(c *Listavaloresmatriz2Context)

	// EnterListavaloresmatriz3 is called when entering the listavaloresmatriz3 production.
	EnterListavaloresmatriz3(c *Listavaloresmatriz3Context)

	// EnterAccesomatriz2 is called when entering the accesomatriz2 production.
	EnterAccesomatriz2(c *Accesomatriz2Context)

	// EnterAccesomatriz3 is called when entering the accesomatriz3 production.
	EnterAccesomatriz3(c *Accesomatriz3Context)

	// EnterAsignmatrizstmt2 is called when entering the asignmatrizstmt2 production.
	EnterAsignmatrizstmt2(c *Asignmatrizstmt2Context)

	// EnterAsignmatrizstmt3 is called when entering the asignmatrizstmt3 production.
	EnterAsignmatrizstmt3(c *Asignmatrizstmt3Context)

	// EnterReturnstmt is called when entering the returnstmt production.
	EnterReturnstmt(c *ReturnstmtContext)

	// EnterBreakstmt is called when entering the breakstmt production.
	EnterBreakstmt(c *BreakstmtContext)

	// EnterContinuestmt is called when entering the continuestmt production.
	EnterContinuestmt(c *ContinuestmtContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIntstmt is called when entering the intstmt production.
	EnterIntstmt(c *IntstmtContext)

	// EnterFloatstmt is called when entering the floatstmt production.
	EnterFloatstmt(c *FloatstmtContext)

	// EnterStringstmt is called when entering the stringstmt production.
	EnterStringstmt(c *StringstmtContext)

	// EnterFuncionNormal is called when entering the FuncionNormal production.
	EnterFuncionNormal(c *FuncionNormalContext)

	// EnterFuncionRetorno is called when entering the FuncionRetorno production.
	EnterFuncionRetorno(c *FuncionRetornoContext)

	// EnterAccfuncstm is called when entering the accfuncstm production.
	EnterAccfuncstm(c *AccfuncstmContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterParametroscall is called when entering the parametroscall production.
	EnterParametroscall(c *ParametroscallContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterElseifstmt is called when entering the elseifstmt production.
	EnterElseifstmt(c *ElseifstmtContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterCaseStmt is called when entering the caseStmt production.
	EnterCaseStmt(c *CaseStmtContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterTypedDeclstmt is called when entering the typedDeclstmt production.
	EnterTypedDeclstmt(c *TypedDeclstmtContext)

	// EnterUntypedDeclstmt is called when entering the untypedDeclstmt production.
	EnterUntypedDeclstmt(c *UntypedDeclstmtContext)

	// EnterOptionalTypedDeclstmt is called when entering the optionalTypedDeclstmt production.
	EnterOptionalTypedDeclstmt(c *OptionalTypedDeclstmtContext)

	// EnterAsignstmt is called when entering the asignstmt production.
	EnterAsignstmt(c *AsignstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterForstmt is called when entering the forstmt production.
	EnterForstmt(c *ForstmtContext)

	// EnterGuardstmt is called when entering the guardstmt production.
	EnterGuardstmt(c *GuardstmtContext)

	// EnterRangostmt is called when entering the rangostmt production.
	EnterRangostmt(c *RangostmtContext)

	// EnterIncremento is called when entering the Incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the Decremento production.
	EnterDecremento(c *DecrementoContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterStringExpr is called when entering the StringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterNilExpr is called when entering the NilExpr production.
	EnterNilExpr(c *NilExprContext)

	// EnterFloatExpr is called when entering the FloatExpr production.
	EnterFloatExpr(c *FloatExprContext)

	// EnterIsEmptyVectorExpr is called when entering the IsEmptyVectorExpr production.
	EnterIsEmptyVectorExpr(c *IsEmptyVectorExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterAccesoVectorExpr is called when entering the AccesoVectorExpr production.
	EnterAccesoVectorExpr(c *AccesoVectorExprContext)

	// EnterUnariaExpr is called when entering the UnariaExpr production.
	EnterUnariaExpr(c *UnariaExprContext)

	// EnterAccesoMatrizExpr is called when entering the AccesoMatrizExpr production.
	EnterAccesoMatrizExpr(c *AccesoMatrizExprContext)

	// EnterCountVectorExpr is called when entering the CountVectorExpr production.
	EnterCountVectorExpr(c *CountVectorExprContext)

	// EnterNumExpr is called when entering the NumExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterAccFuncExpr is called when entering the AccFuncExpr production.
	EnterAccFuncExpr(c *AccFuncExprContext)

	// EnterAccesoValorStructExpr is called when entering the AccesoValorStructExpr production.
	EnterAccesoValorStructExpr(c *AccesoValorStructExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterAccesoVectorStructExpr is called when entering the AccesoVectorStructExpr production.
	EnterAccesoVectorStructExpr(c *AccesoVectorStructExprContext)

	// EnterAccesoStructExpr is called when entering the AccesoStructExpr production.
	EnterAccesoStructExpr(c *AccesoStructExprContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterTipo_vector is called when entering the tipo_vector production.
	EnterTipo_vector(c *Tipo_vectorContext)

	// EnterTipo_matriz2 is called when entering the tipo_matriz2 production.
	EnterTipo_matriz2(c *Tipo_matriz2Context)

	// EnterTipo_matriz3 is called when entering the tipo_matriz3 production.
	EnterTipo_matriz3(c *Tipo_matriz3Context)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockFunc is called when exiting the blockFunc production.
	ExitBlockFunc(c *BlockFuncContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDefStruct is called when exiting the DefStruct production.
	ExitDefStruct(c *DefStructContext)

	// ExitAtributoStruct is called when exiting the AtributoStruct production.
	ExitAtributoStruct(c *AtributoStructContext)

	// ExitStructExpr is called when exiting the StructExpr production.
	ExitStructExpr(c *StructExprContext)

	// ExitStructExprID is called when exiting the StructExprID production.
	ExitStructExprID(c *StructExprIDContext)

	// ExitValorStructExpr is called when exiting the ValorStructExpr production.
	ExitValorStructExpr(c *ValorStructExprContext)

	// ExitDuplastruct is called when exiting the Duplastruct production.
	ExitDuplastruct(c *DuplastructContext)

	// ExitAccesoStruct is called when exiting the AccesoStruct production.
	ExitAccesoStruct(c *AccesoStructContext)

	// ExitAsignStruct is called when exiting the AsignStruct production.
	ExitAsignStruct(c *AsignStructContext)

	// ExitVectorStruct is called when exiting the VectorStruct production.
	ExitVectorStruct(c *VectorStructContext)

	// ExitAccesoVectorStruct is called when exiting the AccesoVectorStruct production.
	ExitAccesoVectorStruct(c *AccesoVectorStructContext)

	// ExitAppendVectorStruct is called when exiting the AppendVectorStruct production.
	ExitAppendVectorStruct(c *AppendVectorStructContext)

	// ExitDeclvectorstmt is called when exiting the declvectorstmt production.
	ExitDeclvectorstmt(c *DeclvectorstmtContext)

	// ExitDefVector is called when exiting the DefVector production.
	ExitDefVector(c *DefVectorContext)

	// ExitDefVectorID is called when exiting the DefVectorID production.
	ExitDefVectorID(c *DefVectorIDContext)

	// ExitListaexpresiones is called when exiting the listaexpresiones production.
	ExitListaexpresiones(c *ListaexpresionesContext)

	// ExitAccesovectorstmt is called when exiting the accesovectorstmt production.
	ExitAccesovectorstmt(c *AccesovectorstmtContext)

	// ExitAsignvectorstmt is called when exiting the asignvectorstmt production.
	ExitAsignvectorstmt(c *AsignvectorstmtContext)

	// ExitAppendvectorstmt is called when exiting the appendvectorstmt production.
	ExitAppendvectorstmt(c *AppendvectorstmtContext)

	// ExitRemoveatvectorstmt is called when exiting the removeatvectorstmt production.
	ExitRemoveatvectorstmt(c *RemoveatvectorstmtContext)

	// ExitRemovelastvectorstmt is called when exiting the removelastvectorstmt production.
	ExitRemovelastvectorstmt(c *RemovelastvectorstmtContext)

	// ExitCountvectorstmt is called when exiting the countvectorstmt production.
	ExitCountvectorstmt(c *CountvectorstmtContext)

	// ExitIsemptyvectorstmt is called when exiting the isemptyvectorstmt production.
	ExitIsemptyvectorstmt(c *IsemptyvectorstmtContext)

	// ExitDeclmatrizstmt2 is called when exiting the declmatrizstmt2 production.
	ExitDeclmatrizstmt2(c *Declmatrizstmt2Context)

	// ExitDeclmatrizstmt3 is called when exiting the declmatrizstmt3 production.
	ExitDeclmatrizstmt3(c *Declmatrizstmt3Context)

	// ExitTipomatriz3 is called when exiting the tipomatriz3 production.
	ExitTipomatriz3(c *Tipomatriz3Context)

	// ExitTipomatriz2 is called when exiting the tipomatriz2 production.
	ExitTipomatriz2(c *Tipomatriz2Context)

	// ExitListavaloresmatriz2 is called when exiting the listavaloresmatriz2 production.
	ExitListavaloresmatriz2(c *Listavaloresmatriz2Context)

	// ExitListavaloresmatriz3 is called when exiting the listavaloresmatriz3 production.
	ExitListavaloresmatriz3(c *Listavaloresmatriz3Context)

	// ExitAccesomatriz2 is called when exiting the accesomatriz2 production.
	ExitAccesomatriz2(c *Accesomatriz2Context)

	// ExitAccesomatriz3 is called when exiting the accesomatriz3 production.
	ExitAccesomatriz3(c *Accesomatriz3Context)

	// ExitAsignmatrizstmt2 is called when exiting the asignmatrizstmt2 production.
	ExitAsignmatrizstmt2(c *Asignmatrizstmt2Context)

	// ExitAsignmatrizstmt3 is called when exiting the asignmatrizstmt3 production.
	ExitAsignmatrizstmt3(c *Asignmatrizstmt3Context)

	// ExitReturnstmt is called when exiting the returnstmt production.
	ExitReturnstmt(c *ReturnstmtContext)

	// ExitBreakstmt is called when exiting the breakstmt production.
	ExitBreakstmt(c *BreakstmtContext)

	// ExitContinuestmt is called when exiting the continuestmt production.
	ExitContinuestmt(c *ContinuestmtContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIntstmt is called when exiting the intstmt production.
	ExitIntstmt(c *IntstmtContext)

	// ExitFloatstmt is called when exiting the floatstmt production.
	ExitFloatstmt(c *FloatstmtContext)

	// ExitStringstmt is called when exiting the stringstmt production.
	ExitStringstmt(c *StringstmtContext)

	// ExitFuncionNormal is called when exiting the FuncionNormal production.
	ExitFuncionNormal(c *FuncionNormalContext)

	// ExitFuncionRetorno is called when exiting the FuncionRetorno production.
	ExitFuncionRetorno(c *FuncionRetornoContext)

	// ExitAccfuncstm is called when exiting the accfuncstm production.
	ExitAccfuncstm(c *AccfuncstmContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitParametroscall is called when exiting the parametroscall production.
	ExitParametroscall(c *ParametroscallContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitElseifstmt is called when exiting the elseifstmt production.
	ExitElseifstmt(c *ElseifstmtContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitCaseStmt is called when exiting the caseStmt production.
	ExitCaseStmt(c *CaseStmtContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitTypedDeclstmt is called when exiting the typedDeclstmt production.
	ExitTypedDeclstmt(c *TypedDeclstmtContext)

	// ExitUntypedDeclstmt is called when exiting the untypedDeclstmt production.
	ExitUntypedDeclstmt(c *UntypedDeclstmtContext)

	// ExitOptionalTypedDeclstmt is called when exiting the optionalTypedDeclstmt production.
	ExitOptionalTypedDeclstmt(c *OptionalTypedDeclstmtContext)

	// ExitAsignstmt is called when exiting the asignstmt production.
	ExitAsignstmt(c *AsignstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitForstmt is called when exiting the forstmt production.
	ExitForstmt(c *ForstmtContext)

	// ExitGuardstmt is called when exiting the guardstmt production.
	ExitGuardstmt(c *GuardstmtContext)

	// ExitRangostmt is called when exiting the rangostmt production.
	ExitRangostmt(c *RangostmtContext)

	// ExitIncremento is called when exiting the Incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the Decremento production.
	ExitDecremento(c *DecrementoContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitStringExpr is called when exiting the StringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitNilExpr is called when exiting the NilExpr production.
	ExitNilExpr(c *NilExprContext)

	// ExitFloatExpr is called when exiting the FloatExpr production.
	ExitFloatExpr(c *FloatExprContext)

	// ExitIsEmptyVectorExpr is called when exiting the IsEmptyVectorExpr production.
	ExitIsEmptyVectorExpr(c *IsEmptyVectorExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitAccesoVectorExpr is called when exiting the AccesoVectorExpr production.
	ExitAccesoVectorExpr(c *AccesoVectorExprContext)

	// ExitUnariaExpr is called when exiting the UnariaExpr production.
	ExitUnariaExpr(c *UnariaExprContext)

	// ExitAccesoMatrizExpr is called when exiting the AccesoMatrizExpr production.
	ExitAccesoMatrizExpr(c *AccesoMatrizExprContext)

	// ExitCountVectorExpr is called when exiting the CountVectorExpr production.
	ExitCountVectorExpr(c *CountVectorExprContext)

	// ExitNumExpr is called when exiting the NumExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitAccFuncExpr is called when exiting the AccFuncExpr production.
	ExitAccFuncExpr(c *AccFuncExprContext)

	// ExitAccesoValorStructExpr is called when exiting the AccesoValorStructExpr production.
	ExitAccesoValorStructExpr(c *AccesoValorStructExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitAccesoVectorStructExpr is called when exiting the AccesoVectorStructExpr production.
	ExitAccesoVectorStructExpr(c *AccesoVectorStructExprContext)

	// ExitAccesoStructExpr is called when exiting the AccesoStructExpr production.
	ExitAccesoStructExpr(c *AccesoStructExprContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitTipo_vector is called when exiting the tipo_vector production.
	ExitTipo_vector(c *Tipo_vectorContext)

	// ExitTipo_matriz2 is called when exiting the tipo_matriz2 production.
	ExitTipo_matriz2(c *Tipo_matriz2Context)

	// ExitTipo_matriz3 is called when exiting the tipo_matriz3 production.
	ExitTipo_matriz3(c *Tipo_matriz3Context)
}
