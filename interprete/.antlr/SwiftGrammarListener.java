// Generated from d:/Documentos/USAC/Cursos/Cursos Segundo Semestre 2023/Laboratorio Organización de Lenguajes y Compiladores 2 - copia/Proyectos/Proyecto 2/interprete/SwiftGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SwiftGrammarParser}.
 */
public interface SwiftGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(SwiftGrammarParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(SwiftGrammarParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#blockFunc}.
	 * @param ctx the parse tree
	 */
	void enterBlockFunc(SwiftGrammarParser.BlockFuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#blockFunc}.
	 * @param ctx the parse tree
	 */
	void exitBlockFunc(SwiftGrammarParser.BlockFuncContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterStmt(SwiftGrammarParser.StmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitStmt(SwiftGrammarParser.StmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DefStruct}
	 * labeled alternative in {@link SwiftGrammarParser#defstructstmt}.
	 * @param ctx the parse tree
	 */
	void enterDefStruct(SwiftGrammarParser.DefStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DefStruct}
	 * labeled alternative in {@link SwiftGrammarParser#defstructstmt}.
	 * @param ctx the parse tree
	 */
	void exitDefStruct(SwiftGrammarParser.DefStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AtributoStruct}
	 * labeled alternative in {@link SwiftGrammarParser#lista_atributos}.
	 * @param ctx the parse tree
	 */
	void enterAtributoStruct(SwiftGrammarParser.AtributoStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AtributoStruct}
	 * labeled alternative in {@link SwiftGrammarParser#lista_atributos}.
	 * @param ctx the parse tree
	 */
	void exitAtributoStruct(SwiftGrammarParser.AtributoStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#struct_expr}.
	 * @param ctx the parse tree
	 */
	void enterStructExpr(SwiftGrammarParser.StructExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#struct_expr}.
	 * @param ctx the parse tree
	 */
	void exitStructExpr(SwiftGrammarParser.StructExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructExprID}
	 * labeled alternative in {@link SwiftGrammarParser#struct_expr}.
	 * @param ctx the parse tree
	 */
	void enterStructExprID(SwiftGrammarParser.StructExprIDContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructExprID}
	 * labeled alternative in {@link SwiftGrammarParser#struct_expr}.
	 * @param ctx the parse tree
	 */
	void exitStructExprID(SwiftGrammarParser.StructExprIDContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ValorStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#valor_struct_expr}.
	 * @param ctx the parse tree
	 */
	void enterValorStructExpr(SwiftGrammarParser.ValorStructExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ValorStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#valor_struct_expr}.
	 * @param ctx the parse tree
	 */
	void exitValorStructExpr(SwiftGrammarParser.ValorStructExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Duplastruct}
	 * labeled alternative in {@link SwiftGrammarParser#l_dupla}.
	 * @param ctx the parse tree
	 */
	void enterDuplastruct(SwiftGrammarParser.DuplastructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Duplastruct}
	 * labeled alternative in {@link SwiftGrammarParser#l_dupla}.
	 * @param ctx the parse tree
	 */
	void exitDuplastruct(SwiftGrammarParser.DuplastructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccesoStruct}
	 * labeled alternative in {@link SwiftGrammarParser#accesostructstmt}.
	 * @param ctx the parse tree
	 */
	void enterAccesoStruct(SwiftGrammarParser.AccesoStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccesoStruct}
	 * labeled alternative in {@link SwiftGrammarParser#accesostructstmt}.
	 * @param ctx the parse tree
	 */
	void exitAccesoStruct(SwiftGrammarParser.AccesoStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignStruct}
	 * labeled alternative in {@link SwiftGrammarParser#asignstructstmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignStruct(SwiftGrammarParser.AsignStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignStruct}
	 * labeled alternative in {@link SwiftGrammarParser#asignstructstmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignStruct(SwiftGrammarParser.AsignStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorStruct}
	 * labeled alternative in {@link SwiftGrammarParser#vector_struct_stmt}.
	 * @param ctx the parse tree
	 */
	void enterVectorStruct(SwiftGrammarParser.VectorStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorStruct}
	 * labeled alternative in {@link SwiftGrammarParser#vector_struct_stmt}.
	 * @param ctx the parse tree
	 */
	void exitVectorStruct(SwiftGrammarParser.VectorStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccesoVectorStruct}
	 * labeled alternative in {@link SwiftGrammarParser#acceso_vector_struct_stmt}.
	 * @param ctx the parse tree
	 */
	void enterAccesoVectorStruct(SwiftGrammarParser.AccesoVectorStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccesoVectorStruct}
	 * labeled alternative in {@link SwiftGrammarParser#acceso_vector_struct_stmt}.
	 * @param ctx the parse tree
	 */
	void exitAccesoVectorStruct(SwiftGrammarParser.AccesoVectorStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AppendVectorStruct}
	 * labeled alternative in {@link SwiftGrammarParser#append_vector_struct_stmt}.
	 * @param ctx the parse tree
	 */
	void enterAppendVectorStruct(SwiftGrammarParser.AppendVectorStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AppendVectorStruct}
	 * labeled alternative in {@link SwiftGrammarParser#append_vector_struct_stmt}.
	 * @param ctx the parse tree
	 */
	void exitAppendVectorStruct(SwiftGrammarParser.AppendVectorStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#declvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclvectorstmt(SwiftGrammarParser.DeclvectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#declvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclvectorstmt(SwiftGrammarParser.DeclvectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DefVector}
	 * labeled alternative in {@link SwiftGrammarParser#defvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterDefVector(SwiftGrammarParser.DefVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DefVector}
	 * labeled alternative in {@link SwiftGrammarParser#defvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitDefVector(SwiftGrammarParser.DefVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DefVectorID}
	 * labeled alternative in {@link SwiftGrammarParser#defvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterDefVectorID(SwiftGrammarParser.DefVectorIDContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DefVectorID}
	 * labeled alternative in {@link SwiftGrammarParser#defvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitDefVectorID(SwiftGrammarParser.DefVectorIDContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listaexpresiones}.
	 * @param ctx the parse tree
	 */
	void enterListaexpresiones(SwiftGrammarParser.ListaexpresionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listaexpresiones}.
	 * @param ctx the parse tree
	 */
	void exitListaexpresiones(SwiftGrammarParser.ListaexpresionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#accesovectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterAccesovectorstmt(SwiftGrammarParser.AccesovectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#accesovectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitAccesovectorstmt(SwiftGrammarParser.AccesovectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#asignvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignvectorstmt(SwiftGrammarParser.AsignvectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#asignvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignvectorstmt(SwiftGrammarParser.AsignvectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#appendvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterAppendvectorstmt(SwiftGrammarParser.AppendvectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#appendvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitAppendvectorstmt(SwiftGrammarParser.AppendvectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#removeatvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterRemoveatvectorstmt(SwiftGrammarParser.RemoveatvectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#removeatvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitRemoveatvectorstmt(SwiftGrammarParser.RemoveatvectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#removelastvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterRemovelastvectorstmt(SwiftGrammarParser.RemovelastvectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#removelastvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitRemovelastvectorstmt(SwiftGrammarParser.RemovelastvectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#countvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterCountvectorstmt(SwiftGrammarParser.CountvectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#countvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitCountvectorstmt(SwiftGrammarParser.CountvectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#isemptyvectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterIsemptyvectorstmt(SwiftGrammarParser.IsemptyvectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#isemptyvectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitIsemptyvectorstmt(SwiftGrammarParser.IsemptyvectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code declmatrizstmt2}
	 * labeled alternative in {@link SwiftGrammarParser#declmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclmatrizstmt2(SwiftGrammarParser.Declmatrizstmt2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code declmatrizstmt2}
	 * labeled alternative in {@link SwiftGrammarParser#declmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclmatrizstmt2(SwiftGrammarParser.Declmatrizstmt2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code declmatrizstmt3}
	 * labeled alternative in {@link SwiftGrammarParser#declmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclmatrizstmt3(SwiftGrammarParser.Declmatrizstmt3Context ctx);
	/**
	 * Exit a parse tree produced by the {@code declmatrizstmt3}
	 * labeled alternative in {@link SwiftGrammarParser#declmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclmatrizstmt3(SwiftGrammarParser.Declmatrizstmt3Context ctx);
	/**
	 * Enter a parse tree produced by the {@code tipomatriz3}
	 * labeled alternative in {@link SwiftGrammarParser#tipomatriz}.
	 * @param ctx the parse tree
	 */
	void enterTipomatriz3(SwiftGrammarParser.Tipomatriz3Context ctx);
	/**
	 * Exit a parse tree produced by the {@code tipomatriz3}
	 * labeled alternative in {@link SwiftGrammarParser#tipomatriz}.
	 * @param ctx the parse tree
	 */
	void exitTipomatriz3(SwiftGrammarParser.Tipomatriz3Context ctx);
	/**
	 * Enter a parse tree produced by the {@code tipomatriz2}
	 * labeled alternative in {@link SwiftGrammarParser#tipomatriz}.
	 * @param ctx the parse tree
	 */
	void enterTipomatriz2(SwiftGrammarParser.Tipomatriz2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code tipomatriz2}
	 * labeled alternative in {@link SwiftGrammarParser#tipomatriz}.
	 * @param ctx the parse tree
	 */
	void exitTipomatriz2(SwiftGrammarParser.Tipomatriz2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code listavaloresmatriz2}
	 * labeled alternative in {@link SwiftGrammarParser#listavaloresmatriz}.
	 * @param ctx the parse tree
	 */
	void enterListavaloresmatriz2(SwiftGrammarParser.Listavaloresmatriz2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code listavaloresmatriz2}
	 * labeled alternative in {@link SwiftGrammarParser#listavaloresmatriz}.
	 * @param ctx the parse tree
	 */
	void exitListavaloresmatriz2(SwiftGrammarParser.Listavaloresmatriz2Context ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listavaloresmatriz3}.
	 * @param ctx the parse tree
	 */
	void enterListavaloresmatriz3(SwiftGrammarParser.Listavaloresmatriz3Context ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listavaloresmatriz3}.
	 * @param ctx the parse tree
	 */
	void exitListavaloresmatriz3(SwiftGrammarParser.Listavaloresmatriz3Context ctx);
	/**
	 * Enter a parse tree produced by the {@code accesomatriz2}
	 * labeled alternative in {@link SwiftGrammarParser#accesomatriz}.
	 * @param ctx the parse tree
	 */
	void enterAccesomatriz2(SwiftGrammarParser.Accesomatriz2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code accesomatriz2}
	 * labeled alternative in {@link SwiftGrammarParser#accesomatriz}.
	 * @param ctx the parse tree
	 */
	void exitAccesomatriz2(SwiftGrammarParser.Accesomatriz2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code accesomatriz3}
	 * labeled alternative in {@link SwiftGrammarParser#accesomatriz}.
	 * @param ctx the parse tree
	 */
	void enterAccesomatriz3(SwiftGrammarParser.Accesomatriz3Context ctx);
	/**
	 * Exit a parse tree produced by the {@code accesomatriz3}
	 * labeled alternative in {@link SwiftGrammarParser#accesomatriz}.
	 * @param ctx the parse tree
	 */
	void exitAccesomatriz3(SwiftGrammarParser.Accesomatriz3Context ctx);
	/**
	 * Enter a parse tree produced by the {@code asignmatrizstmt2}
	 * labeled alternative in {@link SwiftGrammarParser#asignmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignmatrizstmt2(SwiftGrammarParser.Asignmatrizstmt2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code asignmatrizstmt2}
	 * labeled alternative in {@link SwiftGrammarParser#asignmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignmatrizstmt2(SwiftGrammarParser.Asignmatrizstmt2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code asignmatrizstmt3}
	 * labeled alternative in {@link SwiftGrammarParser#asignmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignmatrizstmt3(SwiftGrammarParser.Asignmatrizstmt3Context ctx);
	/**
	 * Exit a parse tree produced by the {@code asignmatrizstmt3}
	 * labeled alternative in {@link SwiftGrammarParser#asignmatrizstmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignmatrizstmt3(SwiftGrammarParser.Asignmatrizstmt3Context ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#returnstmt}.
	 * @param ctx the parse tree
	 */
	void enterReturnstmt(SwiftGrammarParser.ReturnstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#returnstmt}.
	 * @param ctx the parse tree
	 */
	void exitReturnstmt(SwiftGrammarParser.ReturnstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#breakstmt}.
	 * @param ctx the parse tree
	 */
	void enterBreakstmt(SwiftGrammarParser.BreakstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#breakstmt}.
	 * @param ctx the parse tree
	 */
	void exitBreakstmt(SwiftGrammarParser.BreakstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#continuestmt}.
	 * @param ctx the parse tree
	 */
	void enterContinuestmt(SwiftGrammarParser.ContinuestmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#continuestmt}.
	 * @param ctx the parse tree
	 */
	void exitContinuestmt(SwiftGrammarParser.ContinuestmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#intstmt}.
	 * @param ctx the parse tree
	 */
	void enterIntstmt(SwiftGrammarParser.IntstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#intstmt}.
	 * @param ctx the parse tree
	 */
	void exitIntstmt(SwiftGrammarParser.IntstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#floatstmt}.
	 * @param ctx the parse tree
	 */
	void enterFloatstmt(SwiftGrammarParser.FloatstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#floatstmt}.
	 * @param ctx the parse tree
	 */
	void exitFloatstmt(SwiftGrammarParser.FloatstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#stringstmt}.
	 * @param ctx the parse tree
	 */
	void enterStringstmt(SwiftGrammarParser.StringstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#stringstmt}.
	 * @param ctx the parse tree
	 */
	void exitStringstmt(SwiftGrammarParser.StringstmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionNormal}
	 * labeled alternative in {@link SwiftGrammarParser#funcdclstmt}.
	 * @param ctx the parse tree
	 */
	void enterFuncionNormal(SwiftGrammarParser.FuncionNormalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionNormal}
	 * labeled alternative in {@link SwiftGrammarParser#funcdclstmt}.
	 * @param ctx the parse tree
	 */
	void exitFuncionNormal(SwiftGrammarParser.FuncionNormalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionRetorno}
	 * labeled alternative in {@link SwiftGrammarParser#funcdclstmt}.
	 * @param ctx the parse tree
	 */
	void enterFuncionRetorno(SwiftGrammarParser.FuncionRetornoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionRetorno}
	 * labeled alternative in {@link SwiftGrammarParser#funcdclstmt}.
	 * @param ctx the parse tree
	 */
	void exitFuncionRetorno(SwiftGrammarParser.FuncionRetornoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#accfuncstm}.
	 * @param ctx the parse tree
	 */
	void enterAccfuncstm(SwiftGrammarParser.AccfuncstmContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#accfuncstm}.
	 * @param ctx the parse tree
	 */
	void exitAccfuncstm(SwiftGrammarParser.AccfuncstmContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParametros(SwiftGrammarParser.ParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParametros(SwiftGrammarParser.ParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#parametroscall}.
	 * @param ctx the parse tree
	 */
	void enterParametroscall(SwiftGrammarParser.ParametroscallContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#parametroscall}.
	 * @param ctx the parse tree
	 */
	void exitParametroscall(SwiftGrammarParser.ParametroscallContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void enterIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void exitIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#elseifstmt}.
	 * @param ctx the parse tree
	 */
	void enterElseifstmt(SwiftGrammarParser.ElseifstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#elseifstmt}.
	 * @param ctx the parse tree
	 */
	void exitElseifstmt(SwiftGrammarParser.ElseifstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#switchstmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchstmt(SwiftGrammarParser.SwitchstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#switchstmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchstmt(SwiftGrammarParser.SwitchstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#caseStmt}.
	 * @param ctx the parse tree
	 */
	void enterCaseStmt(SwiftGrammarParser.CaseStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#caseStmt}.
	 * @param ctx the parse tree
	 */
	void exitCaseStmt(SwiftGrammarParser.CaseStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#defaultCase}.
	 * @param ctx the parse tree
	 */
	void enterDefaultCase(SwiftGrammarParser.DefaultCaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#defaultCase}.
	 * @param ctx the parse tree
	 */
	void exitDefaultCase(SwiftGrammarParser.DefaultCaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#typedDeclstmt}.
	 * @param ctx the parse tree
	 */
	void enterTypedDeclstmt(SwiftGrammarParser.TypedDeclstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#typedDeclstmt}.
	 * @param ctx the parse tree
	 */
	void exitTypedDeclstmt(SwiftGrammarParser.TypedDeclstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#untypedDeclstmt}.
	 * @param ctx the parse tree
	 */
	void enterUntypedDeclstmt(SwiftGrammarParser.UntypedDeclstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#untypedDeclstmt}.
	 * @param ctx the parse tree
	 */
	void exitUntypedDeclstmt(SwiftGrammarParser.UntypedDeclstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#optionalTypedDeclstmt}.
	 * @param ctx the parse tree
	 */
	void enterOptionalTypedDeclstmt(SwiftGrammarParser.OptionalTypedDeclstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#optionalTypedDeclstmt}.
	 * @param ctx the parse tree
	 */
	void exitOptionalTypedDeclstmt(SwiftGrammarParser.OptionalTypedDeclstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#asignstmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignstmt(SwiftGrammarParser.AsignstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#asignstmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignstmt(SwiftGrammarParser.AsignstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#whilestmt}.
	 * @param ctx the parse tree
	 */
	void enterWhilestmt(SwiftGrammarParser.WhilestmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#whilestmt}.
	 * @param ctx the parse tree
	 */
	void exitWhilestmt(SwiftGrammarParser.WhilestmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForRange}
	 * labeled alternative in {@link SwiftGrammarParser#forstmt}.
	 * @param ctx the parse tree
	 */
	void enterForRange(SwiftGrammarParser.ForRangeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForRange}
	 * labeled alternative in {@link SwiftGrammarParser#forstmt}.
	 * @param ctx the parse tree
	 */
	void exitForRange(SwiftGrammarParser.ForRangeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#guardstmt}.
	 * @param ctx the parse tree
	 */
	void enterGuardstmt(SwiftGrammarParser.GuardstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#guardstmt}.
	 * @param ctx the parse tree
	 */
	void exitGuardstmt(SwiftGrammarParser.GuardstmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Incremento}
	 * labeled alternative in {@link SwiftGrammarParser#opasignstmt}.
	 * @param ctx the parse tree
	 */
	void enterIncremento(SwiftGrammarParser.IncrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Incremento}
	 * labeled alternative in {@link SwiftGrammarParser#opasignstmt}.
	 * @param ctx the parse tree
	 */
	void exitIncremento(SwiftGrammarParser.IncrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Decremento}
	 * labeled alternative in {@link SwiftGrammarParser#opasignstmt}.
	 * @param ctx the parse tree
	 */
	void enterDecremento(SwiftGrammarParser.DecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Decremento}
	 * labeled alternative in {@link SwiftGrammarParser#opasignstmt}.
	 * @param ctx the parse tree
	 */
	void exitDecremento(SwiftGrammarParser.DecrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BoolExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterBoolExpr(SwiftGrammarParser.BoolExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BoolExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitBoolExpr(SwiftGrammarParser.BoolExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterStringExpr(SwiftGrammarParser.StringExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitStringExpr(SwiftGrammarParser.StringExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NilExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNilExpr(SwiftGrammarParser.NilExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NilExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNilExpr(SwiftGrammarParser.NilExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FloatExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFloatExpr(SwiftGrammarParser.FloatExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FloatExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFloatExpr(SwiftGrammarParser.FloatExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IsEmptyVectorExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIsEmptyVectorExpr(SwiftGrammarParser.IsEmptyVectorExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IsEmptyVectorExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIsEmptyVectorExpr(SwiftGrammarParser.IsEmptyVectorExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIdExpr(SwiftGrammarParser.IdExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIdExpr(SwiftGrammarParser.IdExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OpExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterOpExpr(SwiftGrammarParser.OpExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OpExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitOpExpr(SwiftGrammarParser.OpExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccesoVectorExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAccesoVectorExpr(SwiftGrammarParser.AccesoVectorExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccesoVectorExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAccesoVectorExpr(SwiftGrammarParser.AccesoVectorExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code UnariaExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterUnariaExpr(SwiftGrammarParser.UnariaExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code UnariaExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitUnariaExpr(SwiftGrammarParser.UnariaExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccesoMatrizExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAccesoMatrizExpr(SwiftGrammarParser.AccesoMatrizExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccesoMatrizExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAccesoMatrizExpr(SwiftGrammarParser.AccesoMatrizExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CountVectorExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterCountVectorExpr(SwiftGrammarParser.CountVectorExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CountVectorExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitCountVectorExpr(SwiftGrammarParser.CountVectorExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NumExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNumExpr(SwiftGrammarParser.NumExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NumExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNumExpr(SwiftGrammarParser.NumExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterParExpr(SwiftGrammarParser.ParExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitParExpr(SwiftGrammarParser.ParExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StrExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterStrExpr(SwiftGrammarParser.StrExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StrExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitStrExpr(SwiftGrammarParser.StrExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccFuncExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAccFuncExpr(SwiftGrammarParser.AccFuncExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccFuncExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAccFuncExpr(SwiftGrammarParser.AccFuncExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccesoValorStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAccesoValorStructExpr(SwiftGrammarParser.AccesoValorStructExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccesoValorStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAccesoValorStructExpr(SwiftGrammarParser.AccesoValorStructExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIntExpr(SwiftGrammarParser.IntExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIntExpr(SwiftGrammarParser.IntExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccesoVectorStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAccesoVectorStructExpr(SwiftGrammarParser.AccesoVectorStructExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccesoVectorStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAccesoVectorStructExpr(SwiftGrammarParser.AccesoVectorStructExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AccesoStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAccesoStructExpr(SwiftGrammarParser.AccesoStructExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AccesoStructExpr}
	 * labeled alternative in {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAccesoStructExpr(SwiftGrammarParser.AccesoStructExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo(SwiftGrammarParser.TipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo(SwiftGrammarParser.TipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#tipo_vector}.
	 * @param ctx the parse tree
	 */
	void enterTipo_vector(SwiftGrammarParser.Tipo_vectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#tipo_vector}.
	 * @param ctx the parse tree
	 */
	void exitTipo_vector(SwiftGrammarParser.Tipo_vectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#tipo_matriz2}.
	 * @param ctx the parse tree
	 */
	void enterTipo_matriz2(SwiftGrammarParser.Tipo_matriz2Context ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#tipo_matriz2}.
	 * @param ctx the parse tree
	 */
	void exitTipo_matriz2(SwiftGrammarParser.Tipo_matriz2Context ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#tipo_matriz3}.
	 * @param ctx the parse tree
	 */
	void enterTipo_matriz3(SwiftGrammarParser.Tipo_matriz3Context ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#tipo_matriz3}.
	 * @param ctx the parse tree
	 */
	void exitTipo_matriz3(SwiftGrammarParser.Tipo_matriz3Context ctx);
}