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
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.Value}
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
			value := Symbol{Lin: linea, Col: columna, Tipo: tipo, Posicion: posicion, Value: result.Value}
			l.entorno.SetVariable(idVar, value)
			fmt.Println("Asignacion de variable exitosa")
		} else {
			//retornar error de tipos y retornar nil
			fmt.Println("error en los tipos en la asignacion")
			l.generator.AddErrorAsignacion()
		}
	}else{
		fmt.Println("La variable es constante")
		l.generator.AddErrorAsignacion()
	}
	return result
}

//operadores de asignacion += -=

func (l *Visitor) VisitIncremento(ctx *parser.IncrementoContext) interface{} {
	return nil
}

func (l *Visitor) VisitDecremento(ctx *parser.DecrementoContext) interface{} {
	return nil
}
