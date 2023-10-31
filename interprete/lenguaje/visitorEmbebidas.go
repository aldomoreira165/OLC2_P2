package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
	"strings"
)

// visit del print
func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	expresiones := l.Visit(ctx.Listaexpresiones()).([]interface{})
	contador := 1
	var charSalida string
	var result Value
	for _, exp := range expresiones {
		result = exp.(Value)
		
		if contador == len(expresiones) {
			charSalida = "10"
		} else {
			charSalida = "32"
		}

		if result.Type == INTEGER {
			l.generator.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
			l.generator.AddPrintf("c", "(char)"+charSalida)
			l.generator.AddBr()
		} else if result.Type == FLOAT {
			l.generator.AddPrintf("f", "(float)"+fmt.Sprintf("%v", result.Value))
			l.generator.AddPrintf("c", "(char)"+charSalida)
			l.generator.AddBr()
		} else if result.Type == STRING || result.Type == CHAR {
			//llamar a generar printstring
			l.generator.GeneratePrintString()
			//agregar codigo en el main
			newTemp1 := l.generator.NewTemp()
			newTemp2 := l.generator.NewTemp()
			size := strconv.Itoa(l.entorno.Size["size"])
			l.generator.AddExpression(newTemp1, "P", size, "+")     //nuevo temporal en pos vacia
			l.generator.AddExpression(newTemp1, newTemp1, "1", "+") //se deja espacio de retorno
			l.generator.AddSetStack("(int)"+newTemp1, result.Value) //se coloca string en parametro que se manda
			l.generator.AddExpression("P", "P", size, "+")          // cambio de entorno
			l.generator.AddCall("printString")                      //Llamada
			l.generator.AddGetStack(newTemp2, "(int)P")             //obtencion retorno
			l.generator.AddExpression("P", "P", size, "-")          //regreso del entorno
			l.generator.AddPrintf("c","(char)"+charSalida)                  //salto de linea
			l.generator.AddBr()
		} else if result.Type == BOOLEAN {
			if result.IsTemp {
				//cuando es variable
			}
			newLabel := l.generator.NewLabel()
			//add labels
			for _, lvl := range result.TrueLabel {
				l.generator.AddLabel(lvl.(string))
			}
			l.generator.AddPrintf("c", "(char)116")
			l.generator.AddPrintf("c", "(char)114")
			l.generator.AddPrintf("c", "(char)117")
			l.generator.AddPrintf("c", "(char)101")
			l.generator.AddGoto(newLabel)
			//add labels
			for _, lvl := range result.FalseLabel {
				l.generator.AddLabel(lvl.(string))
			}
			l.generator.AddPrintf("c", "(char)102")
			l.generator.AddPrintf("c", "(char)97")
			l.generator.AddPrintf("c", "(char)108")
			l.generator.AddPrintf("c", "(char)115")
			l.generator.AddPrintf("c", "(char)101")
			l.generator.AddLabel(newLabel)
			l.generator.AddPrintf("c", "(char)"+charSalida)
			l.generator.AddBr()
		} else if result.Type == NIL {
			l.generator.AddPrintf("c", "(char)110")
			l.generator.AddPrintf("c", "(char)105")
			l.generator.AddPrintf("c", "(char)108")
			l.generator.AddPrintf("c", "(char)"+charSalida)
			l.generator.AddBr()
		} else {
			fmt.Println("Error en print")
			l.errores.InsertarError("Error en print", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
		contador++
	}
	return result
}

// visit del int
func (l *Visitor) VisitIntstmt(ctx *parser.IntstmtContext) interface{} {
	//convertir a int expresiones string, float en it
	var result Value
	expresion := l.Visit(ctx.Expr()).(Value)

	if expresion.Type == FLOAT {
		//convertir expresion en int
		floatVal, error := strconv.ParseFloat(expresion.Value, 64)

		if error != nil {
			fmt.Println("[Error] expresion no se puede convertir a int")
			l.errores.InsertarError("Error en int", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			result = NewValue("0", false, NIL, "nil", false, false, false)
			return result
		}

		intVal := int(floatVal)
		result = NewValue(strconv.Itoa(intVal), false, INTEGER, expresion.StringValue, false, false, false)
	} else if expresion.Type == STRING {
		cadena := expresion.StringValue

		//convertir cadena a int
		if strings.Contains(cadena, ".") {
			//convertir cadena float a int
			floatVal, error := strconv.ParseFloat(cadena, 64)

			if error != nil {
				fmt.Println("[Error] expresion no se puede convertir a int")
				l.errores.InsertarError("Error en int", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				result = NewValue("0", false, NIL, "nil", false, false, false)
				return result
			}

			intVal := int(floatVal)
			result = NewValue(strconv.Itoa(intVal), false, INTEGER, expresion.StringValue, false, false, false)
		} else {
			//convertir cadena int a int
			intVal, error := strconv.Atoi(cadena)

			if error != nil {
				fmt.Println("[Error] expresion no se puede convertir a int")
				l.errores.InsertarError("Error en int", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				result = NewValue("0", false, NIL, "nil", false, false, false)
				return result
			}

			result = NewValue(strconv.Itoa(intVal), false, INTEGER, expresion.StringValue, false, false, false)
		}
	} else {
		fmt.Println("[Error] expresion no se puede convertir a int")
		l.errores.InsertarError("Error en int", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	return result
}

// visit del float
func (l *Visitor) VisitFloatstmt(ctx *parser.FloatstmtContext) interface{} {
	var result Value
	expresion := l.Visit(ctx.Expr()).(Value)

	if expresion.Type == STRING {
		//obtener valor de la cadena almacenada en el stack con el temporal
		cadena := expresion.StringValue

		//convertir cadena a float
		floatVal, error := strconv.ParseFloat(cadena, 64)
		result = NewValue(strconv.FormatFloat(floatVal, 'f', -1, 64), false, FLOAT, expresion.StringValue, false, false, false)

		if error != nil {
			fmt.Println("[Error] expresion no se puede convertir a float")
			l.errores.InsertarError("Error en float", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			result = NewValue("0", false, NIL, "nil", false, false, false)
			return result
		}
	}
	return result
}

// visit del string
func (l *Visitor) VisitStringstmt(ctx *parser.StringstmtContext) interface{} {
	var result Value
	var tipo TipoExpresion
	expresion := l.Visit(ctx.Expr()).(Value)
	strValue := strings.Trim(expresion.Value, "\"")

	if expresion.Type == INTEGER || expresion.Type == FLOAT{
		strLen := len(strValue)

		if strLen > 1 {
			tipo = STRING
		} else {
			tipo = CHAR
		}

		fmt.Println("String Value: ", strValue)
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), expresion.StringValue, tipo)
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
	} else if expresion.Type == BOOLEAN {
		l.generator.RemoveLastGoto()
		fmt.Println("expresion Value: ", expresion.Value)

		if expresion.Value == "1" {
			strValue = "true"
		} else if expresion.Value == "0" {
			strValue = "false"
		}
		tipo = STRING
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), strValue, tipo)
		//nuevo temporal
		newTemp := l.generator.NewTemp()
		//iguala a heap pointer
		l.generator.AddAssign(newTemp, "H")
		//recorremos string en ascii
		myString := primitive.Valor.(string)
		fmt.Println("myString: ", myString)
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
	} else {
		fmt.Println("[Error] expresion no se puede convertir a string")
		l.errores.InsertarError("Error en string", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	return result
}
