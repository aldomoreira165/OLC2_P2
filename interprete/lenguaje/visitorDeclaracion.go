package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
)

// variables
func (l *Visitor) VisitTypedDeclstmt(ctx *parser.TypedDeclstmtContext) interface{} {
	var result Value
	var newVar Symbol
	var constVar bool
	idVar := ctx.ID().GetText()
	typeVar := ctx.Tipo().GetText()
	result = l.Visit(ctx.Expr()).(Value)
	typeString := ObtenerTipo(result.Type)

	if ctx.LET() != nil {
		constVar = true
	} else {
		constVar = false
	}

	if typeVar == typeString {
		l.generator.AddComment("Declaracion de variable con tipo")
		newVar = l.entorno.SaveVariable(idVar, result.Type, constVar, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), result.StringValue)

		if newVar.Col == 0 && newVar.Lin == 0 && newVar.Tipo == NIL && newVar.Posicion == 0 {
			l.generator.AddErrorDeclaracion()
		} else {
			if result.Type == BOOLEAN {
				//si no es temp (boolean)
				newLabel := l.generator.NewLabel()
				//add labels
				for _, lvl := range result.TrueLabel {
					l.generator.AddLabel(lvl.(string))
				}
				l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), "1")
				l.generator.AddGoto(newLabel)
				//add labels
				for _, lvl := range result.FalseLabel {
					l.generator.AddLabel(lvl.(string))
				}
				l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), "0")
				l.generator.AddGoto(newLabel)
				l.generator.AddLabel(newLabel)
				l.generator.AddBr()
				l.simbolos.InsertarSimbolo(idVar, "Variable", typeVar, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			} else {
				l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), result.Value)
				l.generator.AddBr()
				l.simbolos.InsertarSimbolo(idVar, "Variable", typeVar, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			}
		}
	} else if typeVar == "float" && result.Type == INTEGER {
		//convertir a float
		l.generator.AddComment("Declaracion de variable con tipo")
		newVar = l.entorno.SaveVariable(idVar, FLOAT, constVar, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), result.StringValue)
		if newVar.Col == 0 && newVar.Lin == 0 && newVar.Tipo == NIL && newVar.Posicion == 0 {
			l.generator.AddErrorDeclaracion()
		} else {
			l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), "(float)"+result.Value)
			l.generator.AddBr()
			l.simbolos.InsertarSimbolo(idVar, "Variable", typeVar, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}else {
		//retornar error de tipos y retornar nil
		fmt.Println("error en los tipos de la declaracion")
		l.errores.InsertarError("error en los tipos de la declaracion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		l.generator.AddErrorDeclaracion()
	}
	fmt.Println("Declaracion de variable con tipo")
	fmt.Println(l.entorno.Tabla)
	return result
}

func (l *Visitor) VisitOptionalTypedDeclstmt(ctx *parser.OptionalTypedDeclstmtContext) interface{} {
	l.generator.AddComment("Declaracion de variable con tipo y sin valor")
	idVar := ctx.ID().GetText()
	tipo := ObtenerTipoExpresion(ctx.Tipo().GetText())
	fmt.Println("tipo: ", tipo)
	var result Value
	var newVar Symbol

	if ctx.LET() != nil {
		fmt.Println("let")
		l.errores.InsertarError("error en los tipos de la declaracion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())	
		l.generator.AddErrorDeclaracion()
	} else {
		newVar = l.entorno.SaveVariable(idVar, tipo, false, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), result.StringValue)

		if newVar.Col == 0 && newVar.Lin == 0 && newVar.Tipo == NIL && newVar.Posicion == 0 {
			l.generator.AddErrorDeclaracion()
			l.errores.InsertarError("error en los tipos de la declaracion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		} else {
			l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), "0")
			l.generator.AddBr()
			l.simbolos.InsertarSimbolo(idVar, "Variable", ctx.Tipo().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}
	return result
}

func (l *Visitor) VisitUntypedDeclstmt(ctx *parser.UntypedDeclstmtContext) interface{} {
	l.generator.AddComment("Declaracion de variable sin tipo")
	idVar := ctx.ID().GetText()
	var constVar bool
	var result Value
	var newVar Symbol
	result = l.Visit(ctx.Expr()).(Value)
	fmt.Println("result de var unt: ", result)

	if ctx.LET() != nil {
		constVar = true
	} else {
		constVar = false
	}

	newVar = l.entorno.SaveVariable(idVar, result.Type, constVar, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), result.StringValue)

	if newVar.Col == 0 && newVar.Lin == 0 && newVar.Tipo == NIL && newVar.Posicion == 0 {
		l.generator.AddErrorDeclaracion()
		l.errores.InsertarError("error en los tipos de la declaracion", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	} else {
		if result.Type == BOOLEAN {
			//si no es temp (boolean)
			newLabel := l.generator.NewLabel()
			//add labels
			for _, lvl := range result.TrueLabel {
				l.generator.AddLabel(lvl.(string))
			}
			l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), "1")
			l.generator.AddGoto(newLabel)
			//add labels
			for _, lvl := range result.FalseLabel {
				l.generator.AddLabel(lvl.(string))
			}
			l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), "0")
			l.generator.AddGoto(newLabel)
			l.generator.AddLabel(newLabel)
			l.generator.AddBr()
			l.simbolos.InsertarSimbolo(idVar, "Variable", ObtenerTipo(result.Type), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		} else {
			l.generator.AddSetStack(strconv.Itoa(newVar.Posicion), result.Value)
			l.generator.AddBr()
			l.simbolos.InsertarSimbolo(idVar, "Variable", ObtenerTipo(result.Type), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}
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

func ObtenerTipo(tipo TipoExpresion) string {
	var tipoString string
	if tipo == 0 {
		tipoString = "int"
	} else if tipo == 1 {
		tipoString = "float"
	} else if tipo == 2 {
		tipoString = "String"
	} else if tipo == 3 {
		tipoString = "character"
	} else if tipo == 4 {
		tipoString = "bool"
	} else if tipo == 5 {
		tipoString = "nil"
	}
	return tipoString
}

func ObtenerTipoExpresion(tipo string) TipoExpresion {
	var tipoExp TipoExpresion
	if tipo == "int" {
		tipoExp = INTEGER
	} else if tipo == "float" {
		tipoExp = FLOAT
	} else if tipo == "String" {
		tipoExp = STRING
	} else if tipo == "character" {
		tipoExp = CHAR
	} else if tipo == "bool" {
		tipoExp = BOOLEAN
	} else if tipo == "nil" {
		tipoExp = NIL
	}
	return tipoExp
}
