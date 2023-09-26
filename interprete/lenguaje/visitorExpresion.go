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

	tabla_dominante := [5][5]TipoExpresion{
		//		INTEGER			FLOAT				STRING				BOOLEAN				NULL
		{INTEGER, FLOAT, STRING, BOOLEAN, NULL},
		//FLOAT
		{FLOAT, FLOAT, STRING, NULL, NULL},
		//STRING
		{STRING, STRING, STRING, STRING, NULL},
		//BOOLEAN
		{BOOLEAN, NULL, STRING, BOOLEAN, NULL},
		//NULL
		{NULL, NULL, NULL, NULL, NULL},
	}

	var result Value
	newTemp := l.generator.NewTemp()

	switch op {
	case "+":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "+")
				result = NewValue(newTemp, true, dominante)
				return result
			}
		}
	case "-":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "-")
				result = NewValue(newTemp, true, dominante)
				return result
			}
		}
	case "*":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "*")
				result = NewValue(newTemp, true, dominante)
				return result
			}
		}
	case "/":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				lvl1 := l.generator.NewLabel()
				lvl2 := l.generator.NewLabel()

				l.generator.AddIf(op_der.Value, "0", "!=", lvl1)
				l.generator.AddPrintf("c", "77")
				l.generator.AddPrintf("c", "97")
				l.generator.AddPrintf("c", "116")
				l.generator.AddPrintf("c", "104")
				l.generator.AddPrintf("c", "69")
				l.generator.AddPrintf("c", "114")
				l.generator.AddPrintf("c", "114")
				l.generator.AddPrintf("c", "111")
				l.generator.AddPrintf("c", "114")
				l.generator.AddExpression(newTemp, "0", "", "")
				l.generator.AddGoto(lvl2)
				l.generator.AddLabel(lvl1)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "/")
				l.generator.AddLabel(lvl2)
				result = NewValue(newTemp, true, dominante)
				return result
			}
		}
	case "%":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER {
				lvl1 := l.generator.NewLabel()
				lvl2 := l.generator.NewLabel()

				l.generator.AddIf(op_der.Value, "0", "!=", lvl1)
				l.generator.AddPrintf("c", "77")
				l.generator.AddPrintf("c", "97")
				l.generator.AddPrintf("c", "116")
				l.generator.AddPrintf("c", "104")
				l.generator.AddPrintf("c", "69")
				l.generator.AddPrintf("c", "114")
				l.generator.AddPrintf("c", "114")
				l.generator.AddPrintf("c", "111")
				l.generator.AddPrintf("c", "114")
				l.generator.AddExpression(newTemp, "0", "", "")
				l.generator.AddGoto(lvl2)
				l.generator.AddLabel(lvl1)
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "%")
				l.generator.AddLabel(lvl2)
				result = NewValue(newTemp, true, dominante)
				return result
			}
		}
	}
	return result
}

func (l *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitNumExpr(ctx *parser.NumExprContext) interface{} {
	numStr := ctx.GetText()
	if strings.Contains(numStr, ".") {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			fmt.Println("Error al convertir el numero: ", err)
		}
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), num, FLOAT)
		result := NewValue(fmt.Sprintf("%v", primitive.Valor), false, primitive.Tipo)
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
	return nil
}

func (l *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	strValue := ctx.GetText()
	primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), strValue, STRING)
	result := NewValue(fmt.Sprintf("%v", primitive.Valor), false, primitive.Tipo)
	return result
}

func (l *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	booleanValue := NewValue(ctx.GetText(), false, BOOLEAN)
	trueLabel := l.generator.NewLabel()
	falseLabel := l.generator.NewLabel()

	if booleanValue.Value == "true" {
		l.generator.AddGoto(trueLabel)
	} else {
		l.generator.AddGoto(falseLabel)
	}

	result := NewValue("", false, BOOLEAN)
	return result
}

func (l *Visitor) VisitUnariaExpr(ctx *parser.UnariaExprContext) interface{} {
	return nil
}

// not expr
func (l *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	return nil
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
