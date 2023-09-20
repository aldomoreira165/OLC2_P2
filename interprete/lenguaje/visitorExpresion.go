package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitNumExpr(ctx *parser.NumExprContext) interface{} {
	return nil
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

