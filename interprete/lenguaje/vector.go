package lenguaje

import (
	"interprete/Parser"
)

type Vector struct {
	Id      string
	Tipo    string
	Valores []interface{}
}

// append
func (l *Visitor) VisitAppendvectorstmt(ctx *parser.AppendvectorstmtContext) interface{} {
	return nil
}

// asignacion o modificacion de posicion de vector
func (l *Visitor) VisitAsignvectorstmt(ctx *parser.AsignvectorstmtContext) interface{} {
	return nil
}

// remove last
func (l *Visitor) VisitRemovelastvectorstmt(ctx *parser.RemovelastvectorstmtContext) interface{} {
	return nil
}

// remove at
func (l *Visitor) VisitRemoveatvectorstmt(ctx *parser.RemoveatvectorstmtContext) interface{} {
	return nil
}

// count
func (l *Visitor) VisitCountvectorstmt(ctx *parser.CountvectorstmtContext) interface{} {
	return nil
}

// isEmpty
func (l *Visitor) VisitIsemptyvectorstmt(ctx *parser.IsemptyvectorstmtContext) interface{} {
	return nil
}
