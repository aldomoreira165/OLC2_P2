// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftGrammarParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#blockFunc.
	VisitBlockFunc(ctx *BlockFuncContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#DefStruct.
	VisitDefStruct(ctx *DefStructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AtributoStruct.
	VisitAtributoStruct(ctx *AtributoStructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#StructExpr.
	VisitStructExpr(ctx *StructExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#StructExprID.
	VisitStructExprID(ctx *StructExprIDContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ValorStructExpr.
	VisitValorStructExpr(ctx *ValorStructExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#Duplastruct.
	VisitDuplastruct(ctx *DuplastructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoStruct.
	VisitAccesoStruct(ctx *AccesoStructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AsignStruct.
	VisitAsignStruct(ctx *AsignStructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#VectorStruct.
	VisitVectorStruct(ctx *VectorStructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoVectorStruct.
	VisitAccesoVectorStruct(ctx *AccesoVectorStructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AppendVectorStruct.
	VisitAppendVectorStruct(ctx *AppendVectorStructContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#declvectorstmt.
	VisitDeclvectorstmt(ctx *DeclvectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#DefVector.
	VisitDefVector(ctx *DefVectorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#DefVectorID.
	VisitDefVectorID(ctx *DefVectorIDContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#listaexpresiones.
	VisitListaexpresiones(ctx *ListaexpresionesContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#accesovectorstmt.
	VisitAccesovectorstmt(ctx *AccesovectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#asignvectorstmt.
	VisitAsignvectorstmt(ctx *AsignvectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#appendvectorstmt.
	VisitAppendvectorstmt(ctx *AppendvectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#removeatvectorstmt.
	VisitRemoveatvectorstmt(ctx *RemoveatvectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#removelastvectorstmt.
	VisitRemovelastvectorstmt(ctx *RemovelastvectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#countvectorstmt.
	VisitCountvectorstmt(ctx *CountvectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#isemptyvectorstmt.
	VisitIsemptyvectorstmt(ctx *IsemptyvectorstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#declmatrizstmt2.
	VisitDeclmatrizstmt2(ctx *Declmatrizstmt2Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#declmatrizstmt3.
	VisitDeclmatrizstmt3(ctx *Declmatrizstmt3Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipomatriz3.
	VisitTipomatriz3(ctx *Tipomatriz3Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipomatriz2.
	VisitTipomatriz2(ctx *Tipomatriz2Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#listavaloresmatriz2.
	VisitListavaloresmatriz2(ctx *Listavaloresmatriz2Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#listavaloresmatriz3.
	VisitListavaloresmatriz3(ctx *Listavaloresmatriz3Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#accesomatriz2.
	VisitAccesomatriz2(ctx *Accesomatriz2Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#accesomatriz3.
	VisitAccesomatriz3(ctx *Accesomatriz3Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#asignmatrizstmt2.
	VisitAsignmatrizstmt2(ctx *Asignmatrizstmt2Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#asignmatrizstmt3.
	VisitAsignmatrizstmt3(ctx *Asignmatrizstmt3Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#returnstmt.
	VisitReturnstmt(ctx *ReturnstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#breakstmt.
	VisitBreakstmt(ctx *BreakstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#continuestmt.
	VisitContinuestmt(ctx *ContinuestmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#printstmt.
	VisitPrintstmt(ctx *PrintstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#intstmt.
	VisitIntstmt(ctx *IntstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#floatstmt.
	VisitFloatstmt(ctx *FloatstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#stringstmt.
	VisitStringstmt(ctx *StringstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#FuncionNormal.
	VisitFuncionNormal(ctx *FuncionNormalContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#FuncionRetorno.
	VisitFuncionRetorno(ctx *FuncionRetornoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#accfuncstm.
	VisitAccfuncstm(ctx *AccfuncstmContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#parametroscall.
	VisitParametroscall(ctx *ParametroscallContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ifstmt.
	VisitIfstmt(ctx *IfstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#elseifstmt.
	VisitElseifstmt(ctx *ElseifstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#switchstmt.
	VisitSwitchstmt(ctx *SwitchstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#caseStmt.
	VisitCaseStmt(ctx *CaseStmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#typedDeclstmt.
	VisitTypedDeclstmt(ctx *TypedDeclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#untypedDeclstmt.
	VisitUntypedDeclstmt(ctx *UntypedDeclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#optionalTypedDeclstmt.
	VisitOptionalTypedDeclstmt(ctx *OptionalTypedDeclstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#asignstmt.
	VisitAsignstmt(ctx *AsignstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#whilestmt.
	VisitWhilestmt(ctx *WhilestmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#forstmt.
	VisitForstmt(ctx *ForstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#guardstmt.
	VisitGuardstmt(ctx *GuardstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#rangostmt.
	VisitRangostmt(ctx *RangostmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#Incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#Decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#StringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#NilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IsEmptyVectorExpr.
	VisitIsEmptyVectorExpr(ctx *IsEmptyVectorExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoVectorExpr.
	VisitAccesoVectorExpr(ctx *AccesoVectorExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#UnariaExpr.
	VisitUnariaExpr(ctx *UnariaExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoMatrizExpr.
	VisitAccesoMatrizExpr(ctx *AccesoMatrizExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#CountVectorExpr.
	VisitCountVectorExpr(ctx *CountVectorExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#NumExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccFuncExpr.
	VisitAccFuncExpr(ctx *AccFuncExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoValorStructExpr.
	VisitAccesoValorStructExpr(ctx *AccesoValorStructExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoVectorStructExpr.
	VisitAccesoVectorStructExpr(ctx *AccesoVectorStructExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoStructExpr.
	VisitAccesoStructExpr(ctx *AccesoStructExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipo_vector.
	VisitTipo_vector(ctx *Tipo_vectorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipo_matriz2.
	VisitTipo_matriz2(ctx *Tipo_matriz2Context) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipo_matriz3.
	VisitTipo_matriz3(ctx *Tipo_matriz3Context) interface{}
}
