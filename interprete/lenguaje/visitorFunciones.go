package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
)

// funcion sin retorno y sin parametros
func (l *Visitor) VisitFuncionNormal(ctx *parser.FuncionNormalContext) interface{} {
	idFunction := ctx.ID().GetText()
	var result Value
	l.generator.SetMainFlag(false)
	l.generator.AddComment("******** Funcion " + idFunction)
	l.generator.AddTittle(idFunction)

	//entorno
	var newEnv Environment
	newEnv = NewEnvironment(l.entorno, "funcion "+idFunction)
	newEnv.Size["size"] = l.entorno.Size["size"] + 1
	l.entorno = newEnv

	//en caso existan parametros
	if ctx.Parametros() != nil {
		listaParametros := l.Visit(ctx.Parametros()).([]interface{})

		for _, param := range listaParametros {
			l.entorno.SaveVariable(param.(Value).Value, param.(Value).Type, false, 0, 0, param.(Value).StringValue)
		}
	}

	//instrucciones de funcion
	instruccionesFuncion := ctx.BlockFunc()

	//recorrer instrucciones de funcion
	for _, StamentsCtx := range instruccionesFuncion.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}

	l.generator.AddComment("******** Fin Funcion " + idFunction)
	l.generator.AddEnd()
	l.generator.SetMainFlag(true)
	l.simbolos.InsertarSimbolo(idFunction, "Funcion", "void", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

	//regresar al entorno anterior
	l.entorno = newEnv.Anterior.(Environment)

	return result
}

func (l *Visitor) VisitFuncionRetorno(ctx *parser.FuncionRetornoContext) interface{} {
	idFunction := ctx.ID().GetText()
	tipoRetorno := getTypeParam(ctx.Tipo().GetText())
	var result, resultVal Value
	l.generator.SetMainFlag(false)
	l.generator.AddComment("******** Funcion " + idFunction)
	l.generator.AddTittle(idFunction)

	//entorno
	var newEnv Environment
	newEnv = NewEnvironment(l.entorno, "funcion "+idFunction)
	newEnv.Size["size"] = l.entorno.Size["size"] + 1
	l.entorno = newEnv

	//en caso existan parametros
	if ctx.Parametros() != nil {
		listaParametros := l.Visit(ctx.Parametros()).([]interface{})

		for _, param := range listaParametros {
			l.entorno.SaveVariable(param.(Value).Value, param.(Value).Type, false, 0, 0, param.(Value).StringValue)
		}
	}

	//instrucciones de funcion
	instruccionesFuncion := ctx.BlockFunc()

	//recorrer instrucciones de funcion
	for _, StamentsCtx := range instruccionesFuncion.AllStmt() {
		result = l.Visit(StamentsCtx).(Value)

		if result.ReturnFlag {
			if result.Type == tipoRetorno {
				result.ReturnFlag = false
				resultVal = result
			} else {
				fmt.Println("[ERROR] El tipo de retorno no coincide con el tipo de la funcion")
				l.errores.InsertarError("El tipo de retorno no coincide con el tipo de la funcion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			}

			for _, lvl := range result.OutLabel {
				l.generator.AddLabel(lvl.(string))
			}
		}
	}

	l.generator.AddComment("******** Fin Funcion " + idFunction)
	l.generator.AddEnd()
	l.generator.SetMainFlag(true)
	l.simbolos.InsertarSimbolo(idFunction, "Funcion", ctx.Tipo().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

	//regresar al entorno anterior
	l.entorno = newEnv.Anterior.(Environment)
	l.entorno.outFunctions[idFunction] = result.Value

	fmt.Println("RESULTADO DE FUNCION: ", resultVal)
	return resultVal
}

func (l *Visitor) VisitAccfuncstm(ctx *parser.AccfuncstmContext) interface{} {
	var result Value
	count := 0
	size := l.entorno.Size["size"]
	idFunction := ctx.ID().GetText()
	l.generator.AddComment("Inicio de llamada")

	if ctx.Parametroscall() != nil {
		tmp1 := l.generator.NewTemp()
		l.generator.AddExpression(tmp1, "P", strconv.Itoa((size + 1)), "+")

		//recorrer parametros
		funcParams := l.Visit(ctx.Parametroscall()).([]interface{})

		for _, param := range funcParams {
			fmt.Println("PARAMETRO: ", param.(Value).StringValue)

			if param.(Value).Type == STRING || param.(Value).Type == BOOLEAN || param.(Value).Type == CHAR {
				l.generator.AddSetStack("(int)"+tmp1, param.(Value).Value)
			} else {
				l.generator.AddSetStack("(int)"+tmp1, param.(Value).StringValue)
			}

			if len(funcParams)-1 != count {
				l.generator.AddExpression(tmp1, tmp1, "1", "+")
			}
			count++
		}

		l.generator.AddExpression("P", "P", strconv.Itoa(size), "+")
		l.generator.AddCall(idFunction)
		l.generator.AddGetStack(tmp1, "(int)P")
		l.generator.AddExpression("P", "P", strconv.Itoa(size), "-")
	} else {
		tmp1 := l.generator.NewTemp()
		l.generator.AddExpression("P", "P", strconv.Itoa(size), "+")
		l.generator.AddCall(idFunction)
		l.generator.AddGetStack(tmp1, "(int)P")
		l.generator.AddExpression("P", "P", strconv.Itoa(size), "-")
	}
	l.generator.AddComment("Final de llamada")

	encontrado := false

	if _, ok := l.entorno.outFunctions[idFunction]; ok {
		encontrado = true
	} else {
		encontrado = false
	}

	fmt.Println("ENCONTRADO: ", encontrado)

	if encontrado {
		tempReturn := l.entorno.outFunctions[idFunction]
		result = NewValue(tempReturn, false, INTEGER, tempReturn, false, false, false)
		return result
	}else{
		return result
	}
}

// parametros de la funcion
func (l *Visitor) VisitParametros(ctx *parser.ParametrosContext) interface{} {
	var listaParametros []interface{}
	var result Value
	counter := 0
	paramList := ctx.AllID()
	typeList := ctx.AllTipo()

	for i := 1; i < len(paramList); i += 2 {
		id := paramList[i].GetText()
		typeParam := getTypeParam(typeList[counter].GetText())
		fmt.Println("Type param: ", typeParam)
		result = NewValue(id, false, typeParam, id, false, false, false)
		listaParametros = append(listaParametros, result)
		counter++
	}
	fmt.Println("Parametros declaracion: ", listaParametros)
	return listaParametros
}

func (l *Visitor) VisitParametroscall(ctx *parser.ParametroscallContext) interface{} {
	var listaParametros []interface{}
	var result Value
	//paramIdList := ctx.AllID()
	paramExprList := ctx.AllExpr()

	for i := 0; i < len(paramExprList); i++ {
		//id := paramIdList[i].GetText()
		expr := l.Visit(paramExprList[i]).(Value)
		result = NewValue(expr.Value, false, expr.Type, expr.StringValue, false, false, false)
		l.entorno.SaveVariable(expr.Value, expr.Type, false, 0, 0, expr.StringValue)
		listaParametros = append(listaParametros, result)
	}

	fmt.Println("Parametros llamada: ", listaParametros)
	return listaParametros
}

func getTypeParam(typeStr string) TipoExpresion {
	switch typeStr {
	case "int":
		return INTEGER
	case "float":
		return FLOAT
	case "String":
		return STRING
	case "bool":
		return BOOLEAN
	case "character":
		return CHAR
	default:
		return NIL
	}
}
