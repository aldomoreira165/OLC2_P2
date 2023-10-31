package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

var out = ""
var cont int

type Visitor struct {
	antlr.ParseTreeVisitor
	generator Generator
	entorno Environment
	simbolos *TablaSimbolos
	errores *TablaErrores
}

func NewVisitor() parser.SwiftGrammarVisitor {
	simbolosTabla := NewTablaSimbolos()
	errorTabla := NewTablaErrores()
	return &Visitor{
		generator: NewGenerator(),
		entorno:   NewEnvironment(nil, "global"),	
		simbolos: simbolosTabla,
		errores: errorTabla,
	}
}

func (l *Visitor) VisitS(ctx *parser.SContext) interface{} {
	out = ""
	l.Visit(ctx.Block())
	l.generator.GenerateFinalCode()
	
	for _, val := range l.generator.GetFinalCode(){
		out = out + fmt.Sprintf("%v", val)
	}
	return out
}

// visit de block
func (l *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	l.generator.MainCode = true
	for _, StamentsCtx := range ctx.AllStmt(){
		result:=l.Visit(StamentsCtx).(Value)
		for _, lvl := range result.OutLabel {
			l.generator.AddLabel(lvl.(string))
		}
	}	
	return nil
}

func (l *Visitor) VisitBlockFunc(ctx *parser.BlockFuncContext) interface{} {
	//arreglo de statement
	var arrStm []interface{}

	for _, StamentsCtx := range ctx.AllStmt() {
		arrStm = append(arrStm, StamentsCtx)
	}

	return arrStm
}

// visit de las sentencias
func (l *Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.Printstmt() != nil {
		return l.Visit(ctx.Printstmt())
	}
	if ctx.Ifstmt() != nil {
		return l.Visit(ctx.Ifstmt())
	}
	if ctx.Switchstmt() != nil {
		return l.Visit(ctx.Switchstmt())
	}
	if ctx.TypedDeclstmt() != nil {
		return l.Visit(ctx.TypedDeclstmt())
	}
	if ctx.UntypedDeclstmt() != nil {
		return l.Visit(ctx.UntypedDeclstmt())
	}
	if ctx.OptionalTypedDeclstmt() != nil {
		return l.Visit(ctx.OptionalTypedDeclstmt())
	}
	if ctx.Asignstmt() != nil {
		return l.Visit(ctx.Asignstmt())
	}
	if ctx.Whilestmt() != nil {
		return l.Visit(ctx.Whilestmt())
	}
	if ctx.Forstmt() != nil {
		return l.Visit(ctx.Forstmt())
	}
	if ctx.Guardstmt() != nil {
		return l.Visit(ctx.Guardstmt())
	}
	if ctx.Opasignstmt() != nil {
		return l.Visit(ctx.Opasignstmt())
	}
	if ctx.Breakstmt() != nil {
		return l.Visit(ctx.Breakstmt())
	}
	if ctx.Continuestmt() != nil {
		return l.Visit(ctx.Continuestmt())
	}
	if ctx.Returnstmt() != nil {
		return l.Visit(ctx.Returnstmt())
	}
	if ctx.Funcdclstmt() != nil {
		return l.Visit(ctx.Funcdclstmt())
	}
	if ctx.Accfuncstm() != nil {
		return l.Visit(ctx.Accfuncstm())
	}
	if ctx.Declvectorstmt() != nil {
		return l.Visit(ctx.Declvectorstmt())
	}
	if ctx.Accesovectorstmt() != nil {
		return l.Visit(ctx.Accesovectorstmt())
	}
	if ctx.Appendvectorstmt() != nil {
		return l.Visit(ctx.Appendvectorstmt())
	}
	if ctx.Removelastvectorstmt() != nil {
		return l.Visit(ctx.Removelastvectorstmt())
	}
	if ctx.Removeatvectorstmt() != nil {
		return l.Visit(ctx.Removeatvectorstmt())
	}
	if ctx.Asignvectorstmt() != nil {
		return l.Visit(ctx.Asignvectorstmt())
	}
	if ctx.Declmatrizstmt() != nil {
		return l.Visit(ctx.Declmatrizstmt())
	}
	if ctx.Asignmatrizstmt() != nil {
		return l.Visit(ctx.Asignmatrizstmt())
	}
	if ctx.Defstructstmt() != nil {
		return l.Visit(ctx.Defstructstmt())
	}
	if ctx.Asignstructstmt() != nil {
		return l.Visit(ctx.Asignstructstmt())
	}
	if ctx.Struct_expr() != nil {
		return l.Visit(ctx.Struct_expr())
	}
	if ctx.Vector_struct_stmt() != nil {
		return l.Visit(ctx.Vector_struct_stmt())
	}
	if ctx.Append_vector_struct_stmt() != nil {
		return l.Visit(ctx.Append_vector_struct_stmt())
	}
	return nil
}

// Funcion visitar
func (l *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		nodo := tree.Accept(l)
		return nodo
	}
}

func (l *Visitor) GetSymbolTable() *TablaSimbolos {
	return l.simbolos
}

func (l *Visitor) GetErrorTable() *TablaErrores {
	return l.errores
}
