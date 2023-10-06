package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// visit del print
func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	expresiones := l.Visit(ctx.Listaexpresiones()).([]interface{})
	var result Value
	for _, exp := range expresiones {
		result = exp.(Value)
		if result.Type == INTEGER {
			l.generator.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
			l.generator.AddPrintf("c", "(char)10")
			l.generator.AddBr()
		} else if result.Type == FLOAT {
			l.generator.AddPrintf("f", "(float)"+fmt.Sprintf("%v", result.Value))
			l.generator.AddPrintf("c", "(char)10")
			l.generator.AddBr()
		} else if result.Type == STRING {
			//llamar a generar printstring
			l.generator.GeneratePrintString()
			//agregar codigo en el main
			newTemp1 := l.generator.NewTemp()
			newTemp2 := l.generator.NewTemp()
			size := "0"
			l.generator.AddExpression(newTemp1, "P", size, "+")            //nuevo temporal en pos vacia
			l.generator.AddExpression(newTemp1, newTemp1, "1", "+")        //se deja espacio de retorno
			l.generator.AddSetStack("(int)"+newTemp1, result.Value) //se coloca string en parametro que se manda
			l.generator.AddExpression("P", "P", size, "+")                 // cambio de entorno
			l.generator.AddCall("printString")                             //Llamada
			l.generator.AddGetStack(newTemp2, "(int)P")                    //obtencion retorno
			l.generator.AddExpression("P", "P", size, "-")                 //regreso del entorno
			l.generator.AddPrintf("c", "(char)10")                         //salto de linea
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
			l.generator.AddPrintf("c", "(char)10")
			l.generator.AddBr()
		} else if result.Type == NIL {
			l.generator.AddPrintf("c", "(char)110")
			l.generator.AddPrintf("c", "(char)105")
			l.generator.AddPrintf("c", "(char)108")
			l.generator.AddPrintf("c", "(char)10")
			l.generator.AddBr()
		} else {
			fmt.Println("Error en print")
		}
	}
	return result
}

// visit del int
func (l *Visitor) VisitIntstmt(ctx *parser.IntstmtContext) interface{} {
	/*
	Esta función propia del lenguaje T-Swift permite convertir una expresión ya sea de tipo
	String o Float en una expresión de tipo Int. Si la cadena que recibe como parámetro no
	se puede convertir a un valor numérico se debe desplegar un mensaje de error y devolver
	un valor nil, en caso de recibir un valor de tipo Float, debe realizar un truncamiento.
	*/
	return nil
}

// visit del float
func (l *Visitor) VisitFloatstmt(ctx *parser.FloatstmtContext) interface{} {
	return nil
}

// visit del string
func (l *Visitor) VisitStringstmt(ctx *parser.StringstmtContext) interface{} {
	return nil
}
