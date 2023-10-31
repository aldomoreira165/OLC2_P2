package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
)

// visit de asignaciones
func (l *Visitor) VisitAsignstmt(ctx *parser.AsignstmtContext) interface{} {
	var result Value
	idVar := ctx.ID().GetText()
	variable := l.entorno.GetVariable(idVar)

	if !variable.Constante {
		l.generator.AddComment("Asignacion de variable")
		result = l.Visit(ctx.Expr()).(Value)

		if variable.Tipo == result.Type {
			l.generator.AddSetStack(strconv.Itoa(variable.Posicion), result.Value)
			l.generator.AddBr()
			linea := ctx.GetStart().GetLine()
			columna := ctx.GetStart().GetColumn()
			tipo := result.Type
			posicion := variable.Posicion
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.StringValue}
			l.entorno.SetVariable(idVar, value)
			fmt.Println("Asignacion de variable exitosa")
			fmt.Println(l.entorno.Tabla)
		} else if variable.Tipo == FLOAT && result.Type == INTEGER {
			//convertir a float
			l.generator.AddComment("Asignacion de variable")
			l.generator.AddSetStack(strconv.Itoa(variable.Posicion), "(float)"+result.Value)
			l.generator.AddBr()
			linea := ctx.GetStart().GetLine()
			columna := ctx.GetStart().GetColumn()
			tipo := result.Type
			posicion := variable.Posicion
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.StringValue}
			l.entorno.SetVariable(idVar, value)
			fmt.Println("Asignacion de variable exitosa")
		} else {
			//retornar error de tipos y retornar nil
			fmt.Println("error en los tipos en la asignacion")
			l.generator.AddErrorAsignacion()
		}
	} else {
		fmt.Println("La variable es constante")
		l.generator.AddErrorAsignacion()
	}
	return result
}

//operadores de asignacion += -=

func (l *Visitor) VisitIncremento(ctx *parser.IncrementoContext) interface{} {
	var result Value
	idVar := ctx.Expr(0).GetText()
	variable := l.entorno.GetVariable(idVar)

	if !variable.Constante {
		sumando := l.Visit(ctx.Expr(1)).(Value)
		if sumando.Type == variable.Tipo{
			l.generator.AddComment("Incremento de variable")
			newTemp := l.generator.NewTemp()
			valorVariable := l.Visit(ctx.Expr(0)).(Value)
			strVal := strconv.Itoa(sumando.IntValue + valorVariable.IntValue)
			l.generator.AddExpression(newTemp, valorVariable.Value, sumando.Value, "+")
			result = NewValue(newTemp, true, sumando.Type, strVal, false, false, false)
			l.generator.AddSetStack(strconv.Itoa(variable.Posicion), result.Value)
			l.generator.AddBr()
			linea := ctx.GetStart().GetLine()
			columna := ctx.GetStart().GetColumn()
			tipo := result.Type
			posicion := variable.Posicion
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.StringValue}
			l.entorno.SetVariable(idVar, value)
			fmt.Println("Incremento de variable exitosa")
		}else if sumando.Type == 0 && variable.Tipo == 1 {
			l.generator.AddComment("Incremento de variable")
			newTemp := l.generator.NewTemp()
			valorVariable := l.Visit(ctx.Expr(0)).(Value)
			strVal := strconv.Itoa(sumando.IntValue + valorVariable.IntValue)
			l.generator.AddExpression(newTemp, valorVariable.Value, sumando.Value, "+")
			result = NewValue(newTemp, true, FLOAT, strVal, false, false, false)
			l.generator.AddSetStack(strconv.Itoa(variable.Posicion), result.Value)
			l.generator.AddBr()
			linea := ctx.GetStart().GetLine()
			columna := ctx.GetStart().GetColumn()
			tipo := result.Type
			posicion := variable.Posicion
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.StringValue}
			l.entorno.SetVariable(idVar, value)
			fmt.Println("Incremento de variable exitosa")
		} else {
			//retornar error de tipos y retornar nil
			fmt.Println("error en los tipos en el incremento")
			l.errores.InsertarError("Error en los tipos en el incremento", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			l.generator.AddErrorAsignacion()
		}
	}else{
		fmt.Println("La variable es constante")
		l.errores.InsertarError("La variable es constante", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		l.generator.AddErrorAsignacion()
	}
	return result
}

func (l *Visitor) VisitDecremento(ctx *parser.DecrementoContext) interface{} {
	var result Value
	idVar := ctx.Expr(0).GetText()
	variable := l.entorno.GetVariable(idVar)

	if !variable.Constante {
		sumando := l.Visit(ctx.Expr(1)).(Value)
		if sumando.Type == variable.Tipo{
			l.generator.AddComment("Decremento de variable")
			newTemp := l.generator.NewTemp()
			valorVariable := l.Visit(ctx.Expr(0)).(Value)
			strVal := strconv.Itoa(sumando.IntValue - valorVariable.IntValue)
			l.generator.AddExpression(newTemp, valorVariable.Value, sumando.Value, "-")
			result = NewValue(newTemp, true, sumando.Type, strVal, false, false, false)
			l.generator.AddSetStack(strconv.Itoa(variable.Posicion), result.Value)
			l.generator.AddBr()
			linea := ctx.GetStart().GetLine()
			columna := ctx.GetStart().GetColumn()
			tipo := result.Type
			posicion := variable.Posicion
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.StringValue}
			l.entorno.SetVariable(idVar, value)
			fmt.Println("Decremento de variable exitosa")
		}else if sumando.Type == 0 && variable.Tipo == 1 {
			l.generator.AddComment("Decremento de variable")
			newTemp := l.generator.NewTemp()
			valorVariable := l.Visit(ctx.Expr(0)).(Value)
			strVal := strconv.Itoa(sumando.IntValue - valorVariable.IntValue)
			l.generator.AddExpression(newTemp, valorVariable.Value, sumando.Value, "-")
			result = NewValue(newTemp, true, FLOAT, strVal, false, false, false)
			l.generator.AddSetStack(strconv.Itoa(variable.Posicion), result.Value)
			l.generator.AddBr()
			linea := ctx.GetStart().GetLine()
			columna := ctx.GetStart().GetColumn()
			tipo := result.Type
			posicion := variable.Posicion
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.StringValue}
			l.entorno.SetVariable(idVar, value)
			fmt.Println("Decremento de variable exitosa")
		} else {
			//retornar error de tipos y retornar nil
			fmt.Println("error en los tipos en el decremento")
			l.errores.InsertarError("Error en los tipos en el decremento", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			l.generator.AddErrorAsignacion()
		}
	}else{
		fmt.Println("La variable es constante")
		l.errores.InsertarError("La variable es constante", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		l.generator.AddErrorAsignacion()
	}
	return result
}
