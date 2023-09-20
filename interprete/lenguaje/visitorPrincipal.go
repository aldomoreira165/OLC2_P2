package lenguaje

import (
	"interprete/Parser"
	"log"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	antlr.ParseTreeVisitor
	shouldBreak        bool
	shouldContinue	 	bool
}

func NewVisitor() parser.SwiftGrammarVisitor {
	return &Visitor{
		shouldBreak:        false,
		shouldContinue:		false,
	}
}

func (l *Visitor) VisitS(ctx *parser.SContext) interface{} {
	return l.Visit(ctx.Block())
}

// visit de block
func (l *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	var out string
	for i := 0; ctx.Stmt(i) != nil; i++ {
		if !l.shouldBreak && !l.shouldContinue{
			stmtResult := l.Visit(ctx.Stmt(i))
			switch stmtResult.(type) {
			case int64:
				out += strconv.FormatInt(stmtResult.(int64), 10)
			case float64:
				out += strconv.FormatFloat(stmtResult.(float64), 'f', -1, 64)
			case string:
				out += stmtResult.(string)
			}
		}else{
			break
		}
	}
	return out
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