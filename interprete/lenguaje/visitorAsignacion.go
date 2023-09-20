package lenguaje

import (
	"interprete/Parser"
)

// visit de asignaciones
func (l *Visitor) VisitAsignstmt(ctx *parser.AsignstmtContext) interface{} {
	return nil
}

//operadores de asignacion += -=

func (l *Visitor) VisitIncremento(ctx *parser.IncrementoContext) interface{} {
	return nil
}

func (l *Visitor) VisitDecremento(ctx *parser.DecrementoContext) interface{} {
	return nil
}
