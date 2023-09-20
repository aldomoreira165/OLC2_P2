package lenguaje

import (
	"interprete/Parser"
)

// visit del if
func (l *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitElseifstmt(ctx *parser.ElseifstmtContext) interface{} {
	return nil
}

// switch
func (l *Visitor) VisitSwitchstmt(ctx *parser.SwitchstmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitCaseStmt(ctx *parser.CaseStmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	return nil
}

// visit del while (arreglar) break no jala todavia
func (l *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) interface{} {
	return nil
}

//for
func (l *Visitor) VisitForstmt(ctx *parser.ForstmtContext) interface{} {
	return nil
}

//guard
func (l *Visitor) VisitGuardstmt(ctx *parser.GuardstmtContext) interface{} {
	return nil
}

//rango
func (l *Visitor) VisitRangostmt(ctx *parser.RangostmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitTipo(ctx *parser.TipoContext) interface{} {
	return nil
}

func (l *Visitor) VisitTipo_vector(ctx *parser.Tipo_vectorContext) interface{} {
	return nil
}

func (l *Visitor) VisitTipo_matriz2(ctx *parser.Tipo_matriz2Context) interface{} {
	return nil
}

func (l *Visitor) VisitTipo_matriz3(ctx *parser.Tipo_matriz3Context) interface{} {
	return nil
}