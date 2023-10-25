package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitBreakstmt(ctx *parser.BreakstmtContext) interface{} {
	result := NewValue("", false, NIL, "nil", false, true, false)
	return result
}

func (l *Visitor) VisitReturnstmt(ctx *parser.ReturnstmtContext) interface{} {
	valRet := l.Visit(ctx.Expr()).(Value)
	result := NewValue(valRet.Value, valRet.IsTemp, valRet.Type, valRet.StringValue, true, false, false)
	l.generator.AddSetStack("(int)P", valRet.Value)
	return result
}

func (l *Visitor) VisitContinuestmt(ctx *parser.ContinuestmtContext) interface{} {
	result := NewValue("", false, NIL, "nil", false, false, true)
	return result
}