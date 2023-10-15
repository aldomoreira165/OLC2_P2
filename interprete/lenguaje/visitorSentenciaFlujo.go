package lenguaje

import (
	"interprete/Parser"
)

// Visit del if
func (l *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
	l.generator.AddComment("Inicio de if")
	var condicion, result Value
	var OutLvls []interface{}
	condicion = l.Visit(ctx.Expr()).(Value)
	newLabel := l.generator.NewLabel()

	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	//ejecutando sentencias de condicion principal
	result = l.Visit(ctx.BlockFunc(0)).(Value)
	OutLvls = append(OutLvls, result.OutLabel...)

	//l.generator.AddComment("ultimas etiquetas")
	l.generator.AddGoto(newLabel)
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}
	//*****************************************add false labels

	//else if 
	if ctx.Elseifstmt(0) != nil {
		//ejecutando sentencias de else if
		result = l.Visit(ctx.Elseifstmt(0)).(Value)
		OutLvls = append(OutLvls, result.OutLabel...)
	}

	//else
	if ctx.ELSE() != nil {
		//ejecutando sentencias de else
		result = l.Visit(ctx.BlockFunc(1)).(Value)
		OutLvls = append(OutLvls, result.OutLabel...)
	}

	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls
	return result
}

func (l *Visitor) VisitElseifstmt(ctx *parser.ElseifstmtContext) interface{} {
	l.generator.AddComment("Inicio de else if")
	var condicion, result Value
	var OutLvls []interface{}
	condicion = l.Visit(ctx.Expr()).(Value)
	newLabel := l.generator.NewLabel()

	for _, lvl := range condicion.TrueLabel {
		l.generator.AddLabel(lvl.(string))
	}

	//ejecutando sentencias de condicion principal
	result = l.Visit(ctx.BlockFunc()).(Value)
	OutLvls = append(OutLvls, result.OutLabel...)

	//l.generator.AddComment("ultimas etiquetas")
	l.generator.AddGoto(newLabel)
	for _, lvl := range condicion.FalseLabel {
		l.generator.AddLabel(lvl.(string))
	}
	//*****************************************add false labels

	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls
	return result
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

// for
func (l *Visitor) VisitForstmt(ctx *parser.ForstmtContext) interface{} {
	return nil
}

// guard
func (l *Visitor) VisitGuardstmt(ctx *parser.GuardstmtContext) interface{} {
	return nil
}

// rango
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
