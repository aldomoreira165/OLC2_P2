package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// visit del print
func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	value := ctx.Expr().Accept(l)
	if value.(Value).Type == INTEGER || value.(Value).Type == FLOAT{
		l.generator.AddPrintf("d", "(int)"+fmt.Sprintf("%v", value.(Value).Value))
		l.generator.AddPrintf("c", "10")
		l.generator.AddBr()
	}
	return value
}

// visit del int
func (l *Visitor) VisitIntstmt(ctx *parser.IntstmtContext) interface{} {
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
