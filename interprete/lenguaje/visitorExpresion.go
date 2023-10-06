package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
	"strings"
)

func (l *Visitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	op := ctx.GetOp().GetText()
	op_izq := l.Visit(ctx.GetLeft()).(Value)
	op_der := l.Visit(ctx.GetRight()).(Value)
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
		{NIL, NIL, NIL, NIL, NIL, NIL},
		//BOOLEAN
		{NIL, NIL, NIL, NIL, BOOLEAN, NIL},
		//NIL
		{NIL, NIL, NIL, NIL, NIL, NIL},
	}

	var result Value
	newTemp := l.generator.NewTemp()

	switch op {
	case "+":
		{
			dominante = suma_resta_dominante[op_izq.Type][op_der.Type]
			fmt.Println("dominante: ", dominante)
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "+")
				result = NewValue(newTemp, true, dominante)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				//l.generator.reportarErrorOperacion()
				result = NewValue("", false, NIL)
				return result
			}
		}
	case "-":
		{
			dominante = suma_resta_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "-")
				result = NewValue(newTemp, true, dominante)
				return result
			}else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				//l.generator.reportarErrorOperacion()
				result = NewValue("", false, NIL)
				return result
			}
		}
	case "*":
		{
			dominante = multi_division_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "*")
				result = NewValue(newTemp, true, dominante)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				//l.generator.reportarErrorOperacion()
				result = NewValue("", false, NIL)
				return result
			}
		}
	case "/":
		{
			dominante = multi_division_dominante[op_izq.Type][op_der.Type]
			fmt.Println("dominante: ", dominante)
			if dominante == INTEGER || dominante == FLOAT {
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
				result = NewValue(newTemp, true, dominante)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				//l.generator.reportarErrorOperacion()
				result = NewValue("", false, NIL)
				return result
			}
		}
	case "%":
		{
			dominante = multi_division_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER {
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
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "%")
				l.generator.AddLabel(lvl2)
				result = NewValue(newTemp, true, dominante)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				//l.generator.reportarErrorOperacion()
				result = NewValue("", false, NIL)
				return result
			}
		}
	case "==":
		{
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, "==", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN)
				//result = environment.Value{Value: "", IsTemp: false, Type: environment.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} 
		}
	case "!=":
		{
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT || dominante == STRING || dominante == CHAR || dominante == BOOLEAN || dominante == NIL {
				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, "!=", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN)
				//result = environment.Value{Value: "", IsTemp: false, Type: environment.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			}
		}
	case "<":
		{
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, "<", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN)
				//result = environment.Value{Value: "", IsTemp: false, Type: environment.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result

			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				//l.generator.reportarErrorOperacion()
				result = NewValue("", false, NIL)
				return result
			}
		}
	case ">":
		{
			dominante = relacional_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {

				trueLabel := l.generator.NewLabel()
				falseLabel := l.generator.NewLabel()

				l.generator.AddIf(op_izq.Value, op_der.Value, ">", trueLabel)
				l.generator.AddGoto(falseLabel)

				result = NewValue("", false, BOOLEAN)
				//result = Value{Value: "", IsTemp: false, Type: BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if dominante == NIL {
				fmt.Println("Error: operacion no valida")
				//l.generator.reportarErrorOperacion()
				result = NewValue("", false, NIL)
				return result
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
		}
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), num, FLOAT)
		result := NewValue(fmt.Sprintf("%.1f", primitive.Valor), false, primitive.Tipo)
		return result
	} else {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error al convertir el numero: ", err)
		}
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), num, INTEGER)
		result := NewValue(fmt.Sprintf("%v", primitive.Valor), false, primitive.Tipo)
		return result
	}
}

func (l *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	l.generator.AddComment("Llamando a una variable")
	id := ctx.GetText()
	retSym := l.entorno.GetVariable(id)
	newTemp := l.generator.NewTemp()
	newTemp2 := l.generator.NewTemp()

	if l.generator.MainCode{
		l.generator.AddGetStack(newTemp2, strconv.Itoa(retSym.Posicion))
	}else{
		l.generator.AddExpression(newTemp, "P", strconv.Itoa(retSym.Posicion), "+")
		l.generator.AddGetStack(newTemp2, "(int)"+newTemp)
	}

	result := NewValue(newTemp2, true, retSym.Tipo)
	return result
}

func (l *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	strValue := strings.Trim(ctx.GetText(), "\"")
	fmt.Println("strValue: ", strValue)
	primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), strValue, STRING)
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
	result := NewValue(newTemp, true, primitive.Tipo)
	return result
}

func (l *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	booleanValue := NewValue(ctx.GetText(), false, BOOLEAN)
	trueLabel := l.generator.NewLabel()
	falseLabel := l.generator.NewLabel()
	var value string

	l.generator.AddComment("Primitivo bool")
	if booleanValue.Value == "true" {
		l.generator.AddGoto(trueLabel)
		value = "1"
	} else {
		l.generator.AddGoto(falseLabel)
		value = "0"
	}

	result := NewValue(value, false, BOOLEAN)
	result.TrueLabel = append(result.TrueLabel, trueLabel)
	result.FalseLabel = append(result.FalseLabel, falseLabel)
	return result
}

func (l *Visitor) VisitUnariaExpr(ctx *parser.UnariaExprContext) interface{} {
	expr := l.Visit(ctx.Expr()).(Value)
	

	if expr.Type == FLOAT {
		expresion, _ := strconv.ParseFloat(expr.Value, 64)
		newValue := expresion * -1
		result := NewValue(fmt.Sprintf("%.1f", newValue), false, FLOAT)
		return result
	}else if expr.Type == INTEGER {
		expresion, _ := strconv.Atoi(expr.Value)
		newValue := expresion * -1
		result := NewValue(fmt.Sprintf("%v", newValue), false, INTEGER)
		return result
	}
	return nil
}

// not expr
func (l *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	fmt.Println("entramos a nil")
	valor := NewValue("0", false, NIL)
	return valor
}

func (l *Visitor) VisitAccFuncExpr(ctx *parser.AccFuncExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	return nil
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
