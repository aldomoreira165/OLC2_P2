package lenguaje

import (
	"github.com/antlr4-go/antlr/v4"
)

type Funcion struct {
	Id  string
	Tipo string
	Sentencias antlr.ParserRuleContext
	Parametros []ParametroDef
}

type ParametroDef struct {
	Externo string
	Interno string
	Inout bool
	Tipo    string
}

type ParametroCall struct{
	NameVar string
	Valor interface{}
}

