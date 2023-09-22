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
				result = NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op_izq.IntValue + op_der.IntValue
				return result
			}else{
				fmt.Println("Error de tipos en la suma")
			}
		}
	case "-":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "-")
				result = NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op_izq.IntValue - op_der.IntValue
				return result
			}
		}
	case "*":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "*")
				result = NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op_izq.IntValue * op_der.IntValue
				return result
			}
		}
	case "/":
		{
			dominante = tabla_dominante[op_izq.Type][op_der.Type]
			if dominante == INTEGER || dominante == FLOAT {
				l.generator.AddExpression(newTemp, op_izq.Value, op_der.Value, "/")
				result = NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op_izq.IntValue / op_der.IntValue
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
		num, err := strconv.ParseFloat(numStr, 64);
		if err != nil {
			fmt.Println("Error al convertir el numero: ", err)
		}
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), num, FLOAT)
		result := NewValue(fmt.Sprintf("%v", primitive.Valor), false, primitive.Tipo, false, false, false)
		result.IntValue = primitive.Valor.(int)
		return result
	} else {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error al convertir el numero: ", err)
		}
		primitive := NewPrimitive(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), num, INTEGER)
		result := NewValue(fmt.Sprintf("%v", primitive.Valor), false, primitive.Tipo, false, false, false)
		return result
	}	
}

func (l *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitUnariaExpr(ctx *parser.UnariaExprContext) interface{} {
	return nil
}

//not expr
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

