package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
)

// Visit del if
func (l *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
	l.generator.AddComment("Inicio de if")
	var condicion, result Value
	var OutLvls []interface{}
	var newEnv Environment

	//entorno nuevo
	newEnv = NewEnvironment(l.entorno, "if")
	newEnv.Size["size"] = l.entorno.Size["size"] + 1
	l.entorno = newEnv

	condicion = l.Visit(ctx.Expr()).(Value)
	newLabel := l.generator.NewLabel()

	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	//ejecutando sentencias de condicion principal
	bloqueInstrucciones := ctx.BlockFunc(0)

	for _, StamentsCtx := range bloqueInstrucciones.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.BreakFlag {
			l.generator.AddGoto(l.generator.BreakLabel)
			result.BreakFlag = false
		}

		if result.ContinueFlag {
			l.generator.AddGoto(l.generator.ContinueLabel)
			result.ContinueFlag = false
		}

		for _, lvl := range result.OutLabel {
			OutLvls = append(OutLvls, lvl.(string))
		}
	}

	//l.generator.AddComment("ultimas etiquetas")
	l.generator.AddGoto(newLabel)
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}

	l.entorno = newEnv.Anterior.(Environment)
	//*****************************************add false labels

	//else if
	if ctx.AllElseifstmt() != nil {
		//nuevo entorno
		//ento nuevo
		newEnv = NewEnvironment(l.entorno, "elseif")
		newEnv.Size["size"] = l.entorno.Size["size"] + 1
		l.entorno = newEnv

		for _, elseif := range ctx.AllElseifstmt() {
			result = l.Visit(elseif).(Value)
			
			for _, lvl := range result.OutLabel {
				OutLvls = append(OutLvls, lvl.(string))
			}
		}
		l.entorno = newEnv.Anterior.(Environment)
	}

	//else
	if ctx.ELSE() != nil {

		//nuevo entorno
		newEnv = NewEnvironment(l.entorno, "else")
		newEnv.Size["size"] = l.entorno.Size["size"] + 1
		l.entorno = newEnv

		instruccionesElse := ctx.BlockFunc(1)

		for _, StamentsCtx := range instruccionesElse.AllStmt() {
			result = l.Visit(StamentsCtx).(Value)

			if result.BreakFlag {
				l.generator.AddGoto(l.generator.BreakLabel)
				result.BreakFlag = false
			}

			if result.ContinueFlag {
				l.generator.AddGoto(l.generator.ContinueLabel)
				result.ContinueFlag = false
			}

			for _, lvl := range result.OutLabel {
				OutLvls = append(OutLvls, lvl.(string))
			}
		}

		l.entorno = newEnv.Anterior.(Environment)
	}

	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls
	return result
}

func (l *Visitor) VisitElseifstmt(ctx *parser.ElseifstmtContext) interface{} {
	l.generator.AddComment("Inicio de else if")
	var condicion, result Value
	var OutLvls []interface{}
	condicion = l.Visit(ctx.Expr()).(Value)
	newLabel := l.generator.NewLabel()

	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	//ejecutando sentencias de condicion principal
	bloqueInstrucciones := ctx.BlockFunc()

	for _, StamentsCtx := range bloqueInstrucciones.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.BreakFlag {
			l.generator.AddGoto(l.generator.BreakLabel)
			result.BreakFlag = false
		}

		if result.ContinueFlag {
			l.generator.AddGoto(l.generator.ContinueLabel)
			result.ContinueFlag = false
		}
		

		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}

	//l.generator.AddComment("ultimas etiquetas")
	l.generator.AddGoto(newLabel)
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}
	//*****************************************add false labels
	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls
	return result
}

// switch
func (l *Visitor) VisitSwitchstmt(ctx *parser.SwitchstmtContext) interface{} {
	l.generator.AddComment("Inicio de switch")
	var result, caseResult Value
	var OutLvls []interface{}
	switchExpr := l.Visit(ctx.Expr()).(Value)
	endSwitchLabel := l.generator.NewLabel()
	caseStmts := ctx.AllCaseStmt()

	for _, caseStmtCtx := range caseStmts {
		caseExpr := l.Visit(caseStmtCtx.Expr()).(Value)


		if caseExpr.StringValue == switchExpr.StringValue {
			caseResult = l.Visit(caseStmtCtx).(Value)

			for _, lvl := range caseResult.OutLabel {
				l.generator.AddLabel(lvl.(string))
			}

			// Agregar una etiqueta de salto al final del switch
			l.generator.AddGoto(endSwitchLabel)
			break
		}
	}

	// Si no se encontró un caso que coincida con la condición, ejecutar el caso default
	if caseResult.StringValue == "" {
		defaultCase := ctx.DefaultCase()
		if defaultCase != nil {
			caseResult = l.Visit(defaultCase).(Value)

			for _, lvl := range caseResult.OutLabel {
				l.generator.AddLabel(lvl.(string))
			}
		}
		// Agregar una etiqueta de salto al final del switch
		l.generator.AddGoto(endSwitchLabel)
	}
	// Agregar la etiqueta de salida al final del switch
	l.generator.AddLabel(endSwitchLabel)

	result.OutLabel = OutLvls
	return result
}

func (l *Visitor) VisitCaseStmt(ctx *parser.CaseStmtContext) interface{} {
	l.generator.AddComment("Inicio de case")
	var condicion, result Value
	var OutLvls []interface{}
	condicion = l.Visit(ctx.Expr()).(Value)
	newLabel := l.generator.NewLabel()

	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	//ejecutando sentencias de condicion principal
	instruccionesCase := ctx.BlockFunc()

	for _, StamentsCtx := range instruccionesCase.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.BreakFlag {
			l.generator.AddGoto(l.generator.BreakLabel)
			result.BreakFlag = false
		}

		if result.ContinueFlag {
			l.generator.AddGoto(l.generator.ContinueLabel)
			result.ContinueFlag = false
		}

		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}

	//l.generator.AddComment("ultimas etiquetas")
	l.generator.AddGoto(newLabel)
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}
	//*****************************************add false labels

	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls
	return result
}

func (l *Visitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	l.generator.AddComment("Inicio de default")
	var result Value
	var OutLvls []interface{}
	newLabel := l.generator.NewLabel()

	//ejecutando sentencias de condicion principal
	instruccionesDefault := ctx.BlockFunc()

	for _, StamentsCtx := range instruccionesDefault.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.BreakFlag {
			l.generator.AddGoto(l.generator.BreakLabel)
			result.BreakFlag = false
		}

		if result.ContinueFlag {
			l.generator.AddGoto(l.generator.ContinueLabel)
			result.ContinueFlag = false
		}

		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}

	//l.generator.AddComment("ultimas etiquetas")
	l.generator.AddGoto(newLabel)

	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls
	return result
}

// while
func (l *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) interface{} {
	l.generator.AddComment("Inicio de while")
	var condicion, result Value
	var newEnv Environment

	//entorno nuevo
	newEnv = NewEnvironment(l.entorno, "while")
	newEnv.Size["size"] = l.entorno.Size["size"] + 1
	l.entorno = newEnv

	RetLvl := l.generator.NewLabel()
	l.generator.AddLabel(RetLvl)
	condicion = l.Visit(ctx.Expr()).(Value)

	//******************** add break & continue lvls
	l.generator.AddContinue(RetLvl)
	l.generator.AddBreak(condicion.FalseLabel[0].(string))

	//add true labels
	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	//instrucciones de while
	instruccionesWhile := ctx.BlockFunc()

	for _, StamentsCtx := range instruccionesWhile.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.BreakFlag {
			l.generator.AddGoto(l.generator.BreakLabel)
			result.BreakFlag = false
		}

		if result.ContinueFlag {
			l.generator.AddGoto(l.generator.ContinueLabel)
			result.ContinueFlag = false
		}

		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}

	//add goto
	l.generator.AddGoto(RetLvl)
	//add false labels
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}

	l.entorno = newEnv.Anterior.(Environment)
	return result
}

// for

func (l *Visitor) VisitForRange(ctx *parser.ForRangeContext) interface{} {
	l.generator.AddComment("Inicio de for")
	var result, tmpArr, inicio, final Value
	var newEnv Environment

	//entorno nuevo
	newEnv = NewEnvironment(l.entorno, "for")
	newEnv.Size["size"] = l.entorno.Size["size"] + 1
	l.entorno = newEnv

	id := ctx.ID().GetText()
	var arrType TipoExpresion
	var tmpList []Value

	inicio = l.Visit(ctx.Expr(0)).(Value)
	final = l.Visit(ctx.Expr(1)).(Value)

	if inicio.Type == INTEGER && final.Type == INTEGER {
		inicioInt, _ := strconv.Atoi(inicio.StringValue)
		finalInt, _ := strconv.Atoi(final.StringValue)
		fmt.Println("Final int: ", finalInt)
		size := finalInt - inicioInt + 1

		for i := inicioInt; i <= finalInt; i++ {
			res := NewValue(strconv.Itoa(i), false, INTEGER, strconv.Itoa(i), false, false, false)
			fmt.Println("RES: ", res)
			tmpList = append(tmpList, res)
		}
		//generando array
		l.generator.AddComment("Generando array")
		newtmp1 := l.generator.NewTemp()
		newtmp2 := l.generator.NewTemp()
		l.generator.AddAssign(newtmp1, "H")
		l.generator.AddExpression(newtmp2, newtmp1, "1", "+")
		l.generator.AddSetHeap("(int)H", strconv.Itoa(size))
		l.generator.AddExpression("H", "H", strconv.Itoa(size+1), "+")
		//recorriendo lista de expressiones
		for _, s := range tmpList {
			arrType = s.Type
			l.generator.AddSetHeap("(int)"+newtmp2, s.Value)
			l.generator.AddExpression(newtmp2, newtmp2, "1", "+")
		}
		tmpArr = NewValue(newtmp1, true, arrType, "", false, false, false)
		tmpArr.ArrType = arrType
	}
	tmp1 := l.generator.NewTemp()
	tmp2 := l.generator.NewTemp()
	tmp3 := l.generator.NewTemp()
	tmp4 := l.generator.NewTemp()
	tmp5 := l.generator.NewTemp()
	tmp6 := l.generator.NewTemp()
	lvl1 := l.generator.NewLabel()
	lvl2 := l.generator.NewLabel()
	lvl3 := l.generator.NewLabel()
	l.generator.AddGetHeap(tmp1, "(int)"+tmpArr.Value)       //recuperar tamaño arr
	l.generator.AddExpression(tmp2, "0", "", "")             //contador inicial
	l.generator.AddExpression(tmp3, tmp2, "1", "+")          //indice
	l.generator.AddExpression(tmp4, tmpArr.Value, tmp3, "+") //arr + indice
	l.generator.AddGetHeap(tmp5, "(int)"+tmp4)               //recuperar valor
	symSave := l.entorno.SaveVariable(id, INTEGER, false, 0, 0, inicio.StringValue)
	fmt.Println("symSave: ", symSave)
	l.generator.AddExpression(tmp6, "P", strconv.Itoa(symSave.Posicion), "+") //posicion libre stack
	l.generator.AddSetStack("(int)"+tmp6, tmp5)                               //guardar i
	l.generator.AddLabel(lvl1)                                                //retorno
	l.generator.AddIf(tmp2, tmp1, "==", lvl3)

	instruccionesFor := ctx.BlockFunc()

	for _, StamentsCtx := range instruccionesFor.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.BreakFlag {
			l.generator.AddGoto(l.generator.BreakLabel)
			result.BreakFlag = false
		}

		if result.ContinueFlag {
			l.generator.AddGoto(l.generator.ContinueLabel)
			result.ContinueFlag = false
		}

		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}

	l.generator.AddLabel(lvl2)
	l.generator.AddExpression(tmp2, tmp2, "1", "+")          //aumento contador
	l.generator.AddExpression(tmp3, tmp2, "1", "+")          //indice
	l.generator.AddExpression(tmp4, tmpArr.Value, tmp3, "+") //arr + indice
	l.generator.AddGetHeap(tmp5, "(int)"+tmp4)               //recuperar valor
	l.generator.AddSetStack("(int)"+tmp6, tmp5)              //guardar i nuevamente
	l.generator.AddGoto(lvl1)
	l.generator.AddLabel(lvl3)

	l.entorno = newEnv.Anterior.(Environment)
	return result
}

// guard
func (l *Visitor) VisitGuardstmt(ctx *parser.GuardstmtContext) interface{} {
	l.generator.AddComment("Inicio de guard")
	var condicion, result Value
	var OutLvls []interface{}
	var newEnv Environment

	// Entorno nuevo
	newEnv = NewEnvironment(l.entorno, "guard")
	newEnv.Size["size"] = l.entorno.Size["size"] + 1
	l.entorno = newEnv

	condicion = l.Visit(ctx.Expr()).(Value)
	newLabel := l.generator.NewLabel()

	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}

	bloqueInstrucciones := ctx.BlockFunc()

	for _, StamentsCtx := range bloqueInstrucciones.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.BreakFlag {
			l.generator.AddGoto(l.generator.BreakLabel)
			result.BreakFlag = false
		}

		if result.ContinueFlag {
			l.generator.AddGoto(l.generator.ContinueLabel)
			result.ContinueFlag = false
		}

		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}

	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	l.generator.AddGoto(newLabel)
	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls

	l.entorno = newEnv.Anterior.(Environment)
	return result
}

func (l *Visitor) VisitTipo(ctx *parser.TipoContext) interface{} {
	return nil
}

func (l *Visitor) VisitTipo_vector(ctx *parser.Tipo_vectorContext) interface{} {
	return nil
}

func (l *Visitor) VisitTipo_matriz2(ctx *parser.Tipo_matriz2Context) interface{} {
	return nil
}

func (l *Visitor) VisitTipo_matriz3(ctx *parser.Tipo_matriz3Context) interface{} {
	return nil
}
