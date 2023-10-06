package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
)

// variables
func (l *Visitor) VisitTypedDeclstmt(ctx *parser.TypedDeclstmtContext) interface{} {
	return nil
}

func (l *Visitor) VisitOptionalTypedDeclstmt(ctx *parser.OptionalTypedDeclstmtContext) interface{} {
	return nil

}

func (l *Visitor) VisitUntypedDeclstmt(ctx *parser.UntypedDeclstmtContext) interface{} {
	fmt.Println("Declaracion de variable sin tipo")
	idVar := ctx.ID().GetText()
	var result Value
	var newVar Symbol
	result = l.Visit(ctx.Expr()).(Value)
	newVar = l.entorno.SaveVariable(idVar, result.Type, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	
	l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), result.Value)
	l.generator.AddBr()

	return result
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
	var list []interface{}
	for _, exp := range ctx.AllExpr() {
		value := l.Visit(exp).(Value)
		list = append(list, value)
	} 
	return list
}

// acceso vector
func (l *Visitor) VisitAccesovectorstmt(ctx *parser.AccesovectorstmtContext) interface{} {
	return nil
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

