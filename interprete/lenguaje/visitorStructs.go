package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitDefStruct(ctx *parser.DefStructContext) interface{} {
	return nil
}

func (l *Visitor) VisitAtributoStruct(ctx *parser.AtributoStructContext) interface{} {
	return nil
}

func (l *Visitor) VisitStructExpr(ctx *parser.StructExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitValorStructExpr(ctx *parser.ValorStructExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitStructExprID(ctx *parser.StructExprIDContext) interface{} {
	return nil
}

func (l *Visitor) VisitDuplastruct(ctx *parser.DuplastructContext) interface{} {
	return nil
}

func (l *Visitor) VisitAccesoStruct(ctx *parser.AccesoStructContext) interface{} {
	return nil
}

func (l *Visitor) VisitAsignStruct(ctx *parser.AsignStructContext) interface{} {
	return nil
}

func (l *Visitor) VisitVectorStruct(ctx *parser.VectorStructContext) interface{} {
	return nil
}

func (l *Visitor) VisitAccesoVectorStruct(ctx *parser.AccesoVectorStructContext) interface{} {
	return nil
}

func (l *Visitor) VisitAppendVectorStruct(ctx *parser.AppendVectorStructContext) interface{} {
	return nil
}
