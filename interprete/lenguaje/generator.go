package lenguaje

import (
	"fmt"
)

type Generator struct {
	Temporal      int
	Label         int
	Code          []interface{}
	FinalCode     []interface{}
	Natives       []interface{}
	FuncCode      []interface{}
	TempList      []interface{}
	BreakLabel    string
	ContinueLabel string
}

func NewGenerator() Generator {
	generator := Generator{
		Temporal:      0,
		Label:         0,
		BreakLabel:    "",
		ContinueLabel: "",
	}
	return generator
}

func (g Generator) GetCode() []interface{} {
	return g.Code
}

func (g *Generator) GetFuncCode() []interface{} {
	return g.FuncCode
}

func (g Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

// add break lvl
func (g *Generator) AddBreak(lvl string) {
	g.BreakLabel = lvl
}

// add continue lvl
func (g *Generator) AddContinue(lvl string) {
	g.ContinueLabel = lvl
}

// Generar un nuevo temporal
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temporal)
	g.Temporal = g.Temporal + 1
	//Lo guardamos para declararlo
	g.TempList = append(g.TempList, temp)
	return temp
}

// Generador de Label
func (g *Generator) NewLabel() string {
	temp := g.Label
	g.Label = g.Label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

// Añade Label al codigo
func (g *Generator) AddLabel(Label string) {
	g.FuncCode = append(g.FuncCode, Label+":\n")
}

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
}

func (g *Generator) AddGoto(Label string) {
	g.FuncCode = append(g.FuncCode, "goto "+Label+";\n")
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	g.FuncCode = append(g.FuncCode, target+" = "+left+" "+operator+" "+right+";\n")
}

func (g *Generator) AddAssign(target, val string) {
	g.FuncCode = append(g.FuncCode, target+" = "+val+";\n")
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
}

func (g *Generator) AddBr() {
	g.FuncCode = append(g.FuncCode, "\n")
}

func (g *Generator) AddComment(target string) {
	g.FuncCode = append(g.FuncCode, "//"+target+"\n")
}

func (g *Generator) addErrorMessage() {
	g.AddPrintf("c", "69")
	g.AddPrintf("c", "114")
	g.AddPrintf("c", "114")
	g.AddPrintf("c", "111")
	g.AddPrintf("c", "114")
	g.AddPrintf("c", "58")
	g.AddPrintf("c", "32")
	g.AddPrintf("c", "99")
	g.AddPrintf("c", "111")
	g.AddPrintf("c", "109")
	g.AddPrintf("c", "98")
	g.AddPrintf("c", "105")
	g.AddPrintf("c", "110")
	g.AddPrintf("c", "97")
	g.AddPrintf("c", "99")
	g.AddPrintf("c", "105")
	g.AddPrintf("c", "111")
	g.AddPrintf("c", "110")
	g.AddPrintf("c", "32")
	g.AddPrintf("c", "100")
	g.AddPrintf("c", "101")
	g.AddPrintf("c", "32")
	g.AddPrintf("c", "116")
	g.AddPrintf("c", "105")
	g.AddPrintf("c", "112")
	g.AddPrintf("c", "111")
	g.AddPrintf("c", "115")
	g.AddPrintf("c", "32")
	g.AddPrintf("c", "105")
	g.AddPrintf("c", "110")
	g.AddPrintf("c", "118")
	g.AddPrintf("c", "97")
	g.AddPrintf("c", "108")
	g.AddPrintf("c", "105")
	g.AddPrintf("c", "100")
	g.AddPrintf("c", "97")
	g.AddPrintf("c", "10")
}
