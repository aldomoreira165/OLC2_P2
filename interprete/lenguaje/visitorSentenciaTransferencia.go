package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitBreakstmt(ctx *parser.BreakstmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitReturnstmt(ctx *parser.ReturnstmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitContinuestmt(ctx *parser.ContinuestmtContext) interface{} {
	return nil
}