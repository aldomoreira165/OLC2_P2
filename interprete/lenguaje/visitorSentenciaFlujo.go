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
	condicion = l.Visit(ctx.Expr()).(Value)
	newLabel := l.generator.NewLabel()

	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	//ejecutando sentencias de condicion principal
	result = l.Visit(ctx.BlockFunc(0)).(Value)
	OutLvls = append(OutLvls, result.OutLabel...)

	//l.generator.AddComment("ultimas etiquetas")
	l.generator.AddGoto(newLabel)
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}
	//*****************************************add false labels

	//else if
	if ctx.AllElseifstmt() != nil {
		//ejecutando sentencias de else if
		elseifStmts := ctx.AllElseifstmt()
		for i := 0; i < len(elseifStmts); i++ {
			result = l.Visit(elseifStmts[i]).(Value)
			OutLvls = append(OutLvls, result.OutLabel...)
		}
	}

	//else
	if ctx.ELSE() != nil {
		//ejecutando sentencias de else
		result = l.Visit(ctx.BlockFunc(1)).(Value)
		OutLvls = append(OutLvls, result.OutLabel...)
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
	result = l.Visit(ctx.BlockFunc()).(Value)
	OutLvls = append(OutLvls, result.OutLabel...)

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
	return nil
}

func (l *Visitor) VisitCaseStmt(ctx *parser.CaseStmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	return nil
}

// visit del while (arreglar) break no jala todavia
func (l *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) interface{} {
	l.generator.AddComment("Inicio de while")
	var condicion, result Value
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
	result = l.Visit(ctx.BlockFunc()).(Value)

	//add goto
	l.generator.AddGoto(RetLvl)
	//add false labels
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}
	return result
}

// for

func (l *Visitor) VisitForRange(ctx *parser.ForRangeContext) interface{} {
	l.generator.AddComment("Inicio de for")
	var result, tmpArr, inicio, final Value
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

	result = l.Visit(ctx.BlockFunc()).(Value)

	for _, lvl := range result.OutLabel {
		l.generator.AddLabel(lvl.(string))
	}

	l.generator.AddLabel(lvl2)
	l.generator.AddExpression(tmp2, tmp2, "1", "+")          //aumento contador
	l.generator.AddExpression(tmp3, tmp2, "1", "+")          //indice
	l.generator.AddExpression(tmp4, tmpArr.Value, tmp3, "+") //arr + indice
	l.generator.AddGetHeap(tmp5, "(int)"+tmp4)               //recuperar valor
	l.generator.AddSetStack("(int)"+tmp6, tmp5)              //guardar i nuevamente
	l.generator.AddGoto(lvl1)
	l.generator.AddLabel(lvl3)
	return result
}

// guard
func (l *Visitor) VisitGuardstmt(ctx *parser.GuardstmtContext) interface{} {
	return nil
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
