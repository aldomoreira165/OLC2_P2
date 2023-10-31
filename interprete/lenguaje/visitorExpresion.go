package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
	"strings"
)

func (l *Visitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	op := ctx.GetOp().GetText()
	var dominante TipoExpresion

	suma_resta_dominante := [6][6]TipoExpresion{
		//		INTEGER FLOAT STRING CHAR BOOLEAN NIL
		//INTEGER
		{INTEGER, FLOAT, NIL, NIL, NIL, NIL},
		//FLOAT
		{FLOAT, FLOAT, NIL, NIL, NIL, NIL},
		//STRING
		{NIL, NIL, STRING, NIL, NIL, NIL},
		//CHAR
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//BOOLEAN
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//NIL
		{NIL, NIL, NIL, NIL, NIL, NIL},
	}
	multi_division_dominante := [6][6]TipoExpresion{
		//		INTEGER FLOAT STRING CHAR BOOLEAN NIL
		//INTEGER
		{INTEGER, FLOAT, NIL, NIL, NIL, NIL},
		//FLOAT
		{FLOAT, FLOAT, NIL, NIL, NIL, NIL},
		//STRING
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//CHAR
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//BOOLEAN
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//NIL
		{NIL, NIL, NIL, NIL, NIL, NIL},
	}
	relacional_dominante := [6][6]TipoExpresion{
		//		INTEGER	FLOAT STRING CHAR BOOLEAN NIL
		//INTEGER
		{INTEGER, FLOAT, NIL, NIL, NIL, NIL},
		//FLOAT
		{FLOAT, FLOAT, NIL, NIL, NIL, NIL},
		//STRING
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//CHAR
		{NIL, NIL, NIL, CHAR, NIL, NIL},
		//BOOLEAN
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//NIL
		{NIL, NIL, NIL, NIL, NIL, NIL},
	}

	var result Value
	newTemp := l.generator.NewTemp()

	switch op {
	case "+":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)

			dominante = suma_resta_dominante[op_izq.Type][op_der.Type]
			fmt.Println("dominante: ", dominante)
			if dominante == INTEGER {
				num1, _ := strconv.Atoi(op_izq.StringValue)
				num2, _ := strconv.Atoi(op_der.StringValue)
				strVal := strconv.Itoa(num1 + num2)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "+")
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				result.IntValue = num1 + num2
				return result
			} else if dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.StringValue, 64)
				num2, _ := strconv.ParseFloat(op_der.StringValue, 64)
				strVal := strconv.FormatFloat(num1+num2, 'f', -1, 64)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "+")
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				return result
			} else if dominante == STRING {
				strVal := op_izq.StringValue + op_der.StringValue
				l.generator.GenerateConcatString()
				envSize := strconv.Itoa(l.entorno.Size["size"])
				tmp1 := l.generator.NewTemp()
				tmp2 := l.generator.NewTemp()
				l.generator.AddExpression(tmp1, "P", envSize, "+")
				l.generator.AddExpression(tmp1, tmp1, "1", "+")
				l.generator.AddSetStack("(int)"+tmp1, op_izq.Value)
				l.generator.AddExpression(tmp1, tmp1, "1", "+")
				l.generator.AddSetStack("(int)"+tmp1, op_der.Value)
				l.generator.AddExpression("P", "P", envSize, "+")
				l.generator.AddCall("concatString")
				l.generator.AddGetStack(tmp2, "(int)P")
				l.generator.AddExpression("P", "P", envSize, "-")
				l.generator.AddBr()
				result := NewValue(tmp2, true, STRING, strVal, false, false, false)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case "-":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)

			dominante = suma_resta_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER {
				num1, _ := strconv.Atoi(op_izq.StringValue)
				num2, _ := strconv.Atoi(op_der.StringValue)
				strVal := strconv.Itoa(num1 - num2)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "-")
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				result.IntValue = num1 - num2
				return result
			} else if dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.StringValue, 64)
				num2, _ := strconv.ParseFloat(op_der.StringValue, 64)
				strVal := strconv.FormatFloat(num1-num2, 'f', -1, 64)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "-")
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case "*":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)

			dominante = multi_division_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER {
				num1, _ := strconv.Atoi(op_izq.Value)
				num2, _ := strconv.Atoi(op_der.Value)
				strVal := strconv.Itoa(num1 * num2)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "*")
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				result.IntValue = num1 * num2
				return result
			} else if dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatFloat(num1*num2, 'f', -1, 64)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "*")
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case "/":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)

			dominante = multi_division_dominante[op_izq.Type][op_der.Type]
			fmt.Println("dominante: ", dominante)
			if dominante == INTEGER {
				num1, _ := strconv.Atoi(op_izq.Value)
				num2, _ := strconv.Atoi(op_der.Value)
				strVal := strconv.Itoa(num1 / num2)
				fmt.Println(num1, "/", num2, "=", strVal)

				lvl1 := l.generator.NewLabel()
				lvl2 := l.generator.NewLabel()

				l.generator.AddIf(op_der.Value, "0", "!=", lvl1)
				l.generator.AddPrintf("c", "(char)77")
				l.generator.AddPrintf("c", "(char)97")
				l.generator.AddPrintf("c", "(char)116")
				l.generator.AddPrintf("c", "(char)104")
				l.generator.AddPrintf("c", "(char)69")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)111")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)10")
				l.generator.AddExpression(newTemp, "0", "", "")
				l.generator.AddGoto(lvl2)
				l.generator.AddLabel(lvl1)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "/")
				l.generator.AddLabel(lvl2)
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				return result
			} else if dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatFloat(num1/num2, 'f', -1, 64)

				lvl1 := l.generator.NewLabel()
				lvl2 := l.generator.NewLabel()

				l.generator.AddIf(op_der.Value, "0", "!=", lvl1)
				l.generator.AddPrintf("c", "(char)77")
				l.generator.AddPrintf("c", "(char)97")
				l.generator.AddPrintf("c", "(char)116")
				l.generator.AddPrintf("c", "(char)104")
				l.generator.AddPrintf("c", "(char)69")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)111")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)10")
				l.generator.AddExpression(newTemp, "0", "", "")
				l.generator.AddGoto(lvl2)
				l.generator.AddLabel(lvl1)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "/")
				l.generator.AddLabel(lvl2)
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case "%":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)

			dominante = multi_division_dominante[op_izq.Type][op_der.Type]
			fmt.Println("dominante: ", dominante)
			if dominante == INTEGER {
				num1, _ := strconv.Atoi(op_izq.Value)
				num2, _ := strconv.Atoi(op_der.Value)
				strVal := strconv.Itoa(num1 % num2)

				lvl1 := l.generator.NewLabel()
				lvl2 := l.generator.NewLabel()

				l.generator.AddIf(op_der.StringValue, "0", "!=", lvl1)
				l.generator.AddPrintf("c", "(char)77")
				l.generator.AddPrintf("c", "(char)97")
				l.generator.AddPrintf("c", "(char)116")
				l.generator.AddPrintf("c", "(char)104")
				l.generator.AddPrintf("c", "(char)69")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)111")
				l.generator.AddPrintf("c", "(char)114")
				l.generator.AddPrintf("c", "(char)10")
				l.generator.AddExpression(newTemp, "0", "", "")
				l.generator.AddGoto(lvl2)
				l.generator.AddLabel(lvl1)
				l.generator.AddExpression(newTemp,"(int)"+op_izq.Value, "(int)"+op_der.Value, "%")
				l.generator.AddLabel(lvl2)
				result = NewValue(newTemp, true, dominante, strVal, false, false, false)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case "==":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatBool(num1 == num2)
				fmt.Println("STRVAL: ", strVal)

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, "==", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN, strVal, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			}
		}
	case "!=":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatBool(num1 != num2)

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, "!=", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN, strVal, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			}
		}
	case "<":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatBool(num1 < num2)

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, "<", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN, strVal, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result

			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case ">":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatBool(num1 > num2)

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, ">", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN, strVal, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case "<=":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatBool(num1 <= num2)

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, "<=", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN, strVal, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case ">=":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			op_der := l.Visit(ctx.GetRight()).(Value)
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				num1, _ := strconv.ParseFloat(op_izq.Value, 64)
				num2, _ := strconv.ParseFloat(op_der.Value, 64)
				strVal := strconv.FormatBool(num1 >= num2)

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, ">=", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN, strVal, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				l.generator.AddErrorOperacion()
				result = NewValue("", false, NIL, "nil", false, false, false)
				return result
			}
		}
	case "&&":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)

			//add op1 labels
			for _, lvl := range op_izq.TrueLabel {
				l.generator.AddLabel(lvl.(string))
			}

			op_der := l.Visit(ctx.GetRight()).(Value)

			result = NewValue("", false, BOOLEAN, "", false, false, false)
			result.TrueLabel = append(op_der.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op_izq.FalseLabel, result.FalseLabel...)
			result.FalseLabel = append(op_der.FalseLabel, result.FalseLabel...)
			return result

		}
	case "||":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)

			//add op1 labels
			for _, lvl := range op_izq.FalseLabel {
				l.generator.AddLabel(lvl.(string))
			}

			op_der := l.Visit(ctx.GetRight()).(Value)

			result = NewValue("", false, BOOLEAN, "", false, false, false)
			result.TrueLabel = append(op_izq.TrueLabel, result.TrueLabel...)
			result.TrueLabel = append(op_der.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op_der.FalseLabel, result.FalseLabel...)
			return result
		}
	case "!":
		{
			op_izq := l.Visit(ctx.GetLeft()).(Value)
			var value string
			var valueString string

			if op_izq.Value == "1" {
				value = "0"
				valueString = "false"
			} else {
				value = "1"
				valueString = "true"
			}

			if op_izq.Type == BOOLEAN {
				result = NewValue(value, false, BOOLEAN, valueString, false, false, false)
				result.TrueLabel = append(op_izq.FalseLabel, result.FalseLabel...)
				result.FalseLabel = append(op_izq.TrueLabel, result.FalseLabel...)
				return result
			} else {
				fmt.Println("ERROR: tipo no compatible !")
				l.errores.InsertarError("Error en la operacion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			}
		}
	}
	return result
}

func (l *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return l.Visit(ctx.Expr())
}

func (l *Visitor) VisitNumExpr(ctx *parser.NumExprContext) interface{} {
	numStr := ctx.GetText()
	if strings.Contains(numStr, ".") {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			fmt.Println("Error al convertir el numero: ", err)
			l.errores.InsertarError("Error al convertir el numero", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), num, FLOAT)
		result := NewValue(fmt.Sprintf("%.4f", primitive.Valor), false, primitive.Tipo, numStr, false, false, false)
		return result
	} else {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error al convertir el numero: ", err)
			l.errores.InsertarError("Error al convertir el numero", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), num, INTEGER)
		result := NewValue(fmt.Sprintf("%v", primitive.Valor), false, primitive.Tipo, numStr, false, false, false)
		result.IntValue = num
		return result
	}
}

func (l *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	l.generator.AddComment("Llamando a una variable")
	var result Value
	id := ctx.GetText()
	retSym := l.entorno.GetVariable(id)
	newTemp := l.generator.NewTemp()
	newTemp2 := l.generator.NewTemp()

	if l.generator.MainCode {
		l.generator.AddGetStack(newTemp2, strconv.Itoa(retSym.Posicion))
	} else {
		l.generator.AddExpression(newTemp, "P", strconv.Itoa(retSym.Posicion), "+")
		l.generator.AddGetStack(newTemp2, "(int)"+newTemp)
	}

	if retSym.Tipo == BOOLEAN {
		trueLabel := l.generator.NewLabel()
		falseLabel := l.generator.NewLabel()
		l.generator.AddIf(newTemp2, "1", "==", trueLabel)
		l.generator.AddGoto(falseLabel)
		result = NewValue("", false, BOOLEAN, retSym.Value, false, false, false)
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)
	} else {
		fmt.Println("retSym.Value: ", retSym.Value)
		result = NewValue(newTemp2, true, retSym.Tipo, retSym.Value, false, false, false)
		//result.ArrType = retSym.TipoArr
	}

	return result
}

func (l *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	var tipo TipoExpresion
	strValue := strings.Trim(ctx.GetText(), "\"")
	fmt.Println("strValue: ", strValue)

	//obtener longitud de cadena para saber si sera string o char
	strLen := len(strValue)

	if strLen > 1 {
		tipo = STRING
	} else {
		tipo = CHAR
	}

	primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), strValue, tipo)
	//nuevo temporal
	newTemp := l.generator.NewTemp()
	//iguala a heap pointer
	l.generator.AddAssign(newTemp, "H")
	//recorremos string en ascii
	myString := primitive.Valor.(string)
	byteArray := []byte(myString)
	for _, asc := range byteArray {
		//se agrega ascii al heap
		l.generator.AddSetHeap("(int)H", strconv.Itoa(int(asc)))
		//suma heap pointer
		l.generator.AddExpression("H", "H", "1", "+")
	}
	//caracteres de escape
	l.generator.AddSetHeap("(int)H", "-1")
	l.generator.AddExpression("H", "H", "1", "+")
	l.generator.AddBr()
	result := NewValue(newTemp, true, primitive.Tipo, strValue, false, false, false)
	return result
}

func (l *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	boolVal := ctx.GetText()
	var valRet string
	var valueString string
	var result Value
	l.generator.AddComment("Primitivo bool")
	trueLabel := l.generator.NewLabel()
	falseLabel := l.generator.NewLabel()
	if boolVal == "true" {
		l.generator.AddGoto(trueLabel)
		valRet = "1"
		valueString = "true"
	} else {
		l.generator.AddGoto(falseLabel)
		valRet = "0"
		valueString = "false"
	}
	result = NewValue(valRet, false, BOOLEAN, valueString, false, false, false)
	result.TrueLabel = append(result.TrueLabel, trueLabel)
	result.FalseLabel = append(result.FalseLabel, falseLabel)
	return result
}

func (l *Visitor) VisitUnariaExpr(ctx *parser.UnariaExprContext) interface{} {
	expr := l.Visit(ctx.Expr()).(Value)

	if expr.Type == FLOAT {
		expresion, _ := strconv.ParseFloat(expr.Value, 64)
		newValue := expresion * -1
		result := NewValue(fmt.Sprintf("%.1f", newValue), false, FLOAT, strconv.FormatFloat(newValue, 'f', 4, 64), false, false, false)
		return result
	} else if expr.Type == INTEGER {
		expresion, _ := strconv.Atoi(expr.Value)
		newValue := expresion * -1
		result := NewValue(fmt.Sprintf("%v", newValue), false, INTEGER, strconv.Itoa(newValue), false, false, false)
		return result
	}
	return nil
}

func (l *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	fmt.Println("entramos a nil")
	valor := NewValue("0", false, NIL, "nil", false, false, false)
	return valor
}

func (l *Visitor) VisitAccFuncExpr(ctx *parser.AccFuncExprContext) interface{} {
	return l.Visit(ctx.Accfuncstm())
}

func (l *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return l.Visit(ctx.Intstmt())
}

func (l *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return l.Visit(ctx.Floatstmt())
}

func (l *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	return l.Visit(ctx.Stringstmt())
}

func (l *Visitor) VisitAccesoVectorExpr(ctx *parser.AccesoVectorExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitCountVectorExpr(ctx *parser.CountVectorExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitIsEmptyVectorExpr(ctx *parser.IsEmptyVectorExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitAccesoMatrizExpr(ctx *parser.AccesoMatrizExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitAccesoStructExpr(ctx *parser.AccesoStructExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitAccesoValorStructExpr(ctx *parser.AccesoValorStructExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitAccesoVectorStructExpr(ctx *parser.AccesoVectorStructExprContext) interface{} {
	return nil
}
