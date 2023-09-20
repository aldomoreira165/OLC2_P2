package lenguaje

import (
	"interprete/Parser"
)

// variables
func (l *Visitor) VisitTypedDeclstmt(ctx *parser.TypedDeclstmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitOptionalTypedDeclstmt(ctx *parser.OptionalTypedDeclstmtContext) interface{} {
	return nil

}

func (l *Visitor) VisitUntypedDeclstmt(ctx *parser.UntypedDeclstmtContext) interface{} {
	return nil
}

// vectores
func (l *Visitor) VisitDeclvectorstmt(ctx *parser.DeclvectorstmtContext) interface{} {
	return nil
}

// definicicion de vector
func (l *Visitor) VisitDefVector(ctx *parser.DefVectorContext) interface{} {
	return nil
}

// asignar a vector los valores de otro vector
func (l *Visitor) VisitDefVectorID(ctx *parser.DefVectorIDContext) interface{} {
	return nil
}

// lista de valores
func (l *Visitor) VisitListaexpresiones(ctx *parser.ListaexpresionesContext) interface{} {
	return nil
}

// acceso vector
func (l *Visitor) VisitAccesovectorstmt(ctx *parser.AccesovectorstmtContext) interface{} {
	return nil
}

// matrices
func (l *Visitor) VisitDeclmatrizstmt2(ctx *parser.Declmatrizstmt2Context) interface{} {
	return nil
}

func (l *Visitor) VisitTipomatriz2(ctx *parser.Tipomatriz2Context) interface{} {
	return nil
}

func (l *Visitor) VisitListavaloresmatriz2(ctx *parser.Listavaloresmatriz2Context) interface{} {
	return nil
}

// acceso matriz
func (l *Visitor) VisitAccesomatriz2(ctx *parser.Accesomatriz2Context) interface{} {
	return nil
}

// asignmatrizstmt
func (l *Visitor) VisitAsignmatrizstmt2(ctx *parser.Asignmatrizstmt2Context) interface{} {
	return nil
}

//matrices de 3 dimensiones

func (l *Visitor) VisitDeclmatrizstmt3(ctx *parser.Declmatrizstmt3Context) interface{} {
	return nil
}

func (l *Visitor) VisitTipomatriz3(ctx *parser.Tipomatriz3Context) interface{} {
	return nil
}

func (l *Visitor) VisitListavaloresmatriz3(ctx *parser.Listavaloresmatriz3Context) interface{} {
	return nil
}

func (l *Visitor) VisitAccesomatriz3(ctx *parser.Accesomatriz3Context) interface{} {
	return nil
}


func (l *Visitor) VisitAsignmatrizstmt3(ctx *parser.Asignmatrizstmt3Context) interface{} {
	return nil
}

