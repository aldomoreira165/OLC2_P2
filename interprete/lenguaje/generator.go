package lenguaje

import (
	"fmt"
	"strings"
)

type Generator struct {
	Temporal      int
	Label         int
	Code          []interface{}
	FinalCode     []interface{}
	Natives       []interface{}
	FuncCode      []interface{}
	TempList      []interface{}
	PrintStringFlag  bool
	ConcatStringFlag bool
	BreakLabel    string
	ContinueLabel string
	MainCode      bool
}

func NewGenerator() Generator {
	generator := Generator{
		Temporal:      0,
		Label:         0,
		Code:             []interface{}{},
		FinalCode:        []interface{}{},
		Natives:          []interface{}{},
		FuncCode:         []interface{}{},
		TempList:         []interface{}{},
		PrintStringFlag:  true,
		ConcatStringFlag: true,
		BreakLabel:    "",
		ContinueLabel: "",
		MainCode:      true,
	}
	return generator
}

func (g Generator) GetCode() []interface{} {
	return g.Code
}

func (g Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

func (g Generator) GetTemps() []interface{} {
	return g.TempList
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
	if g.MainCode {
		g.Code = append(g.Code, Label+":\n")
	} else {
		g.FuncCode = append(g.FuncCode, Label+":\n")
	}
}

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	}
}
func (g *Generator) RemoveLastGoto() {
	var codeSlice []interface{}
	if g.MainCode {
		codeSlice = g.Code
	} else {
		codeSlice = g.FuncCode
	}

	if len(codeSlice) > 0 && strings.HasPrefix(fmt.Sprintf("%v", codeSlice[len(codeSlice)-1]), "goto") {
		codeSlice = codeSlice[:len(codeSlice)-1]
	}

	if g.MainCode {
		g.Code = codeSlice
	} else {
		g.FuncCode = codeSlice
	}
}

func (g *Generator) AddGoto(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "goto "+Label+";\n")
	}
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	if g.MainCode {	
		g.Code = append(g.Code, target+" = "+left+" "+operator+" "+right+";\n")
	} else {	
		g.FuncCode = append(g.FuncCode, target+" = "+left+" "+operator+" "+right+";\n")
	}
}

func (g *Generator) AddAssign(target, val string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+val+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+val+";\n")
	}
}

func (g *Generator) AddSetHeap(index string, value string) {
	if g.MainCode {
		g.Code=append(g.Code,"heap[" + index + "] = " + value + ";\n")
	} else {
		g.FuncCode=append(g.FuncCode,"heap[" + index + "] = " + value + ";\n")
	}
}

func (g *Generator) AddGetHeap(target string, index string) {
	if g.MainCode {
		g.Code=append(g.Code,target + " = heap[" + index + "];\n")
	} else {
		g.FuncCode=append(g.FuncCode,target + " = heap[" + index + "];\n")
	}
}

func (g *Generator) AddSetStack(index string, value string) {
	if g.MainCode {
		g.Code=append(g.Code,"stack[" + index + "] = " + value + ";\n")
	} else {
		g.FuncCode=append(g.FuncCode,"stack[" + index + "] = " + value + ";\n")
	}
}

func (g *Generator) AddGetStack(target string, index string) {
	if g.MainCode {
		g.Code=append(g.Code,target + " = stack[" + index + "];\n")
	} else {
		g.FuncCode=append(g.FuncCode,target + " = stack[" + index + "];\n")
	}
}

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code=append(g.Code,target + "();\n")
	} else {
		g.FuncCode=append(g.FuncCode,target + "();\n")
	}
}
func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "//"+target+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "//"+target+"\n")
	}
}

func (g *Generator) AddTittle(target string) {
	if g.MainCode {
		g.Code = append(g.Code,"void " + target + "() {\n") 
	} else {
		g.FuncCode = append(g.FuncCode,"void " + target + "() {\n")
	}
}

func (g *Generator) AddEnd() {
	if g.MainCode {
		g.Code = append(g.Code,"\treturn;\n")
		g.Code = append(g.Code,"}\n\n")
	} else {
		g.FuncCode = append(g.FuncCode,"\treturn;\n")
		g.FuncCode = append(g.FuncCode,"}\n\n")
	}
}

func (g *Generator) SetMainFlag(newVal bool) {
	g.MainCode = newVal
}

//agregar headers
func (g *Generator) GenerateFinalCode() {
	//****************** add head
	g.FinalCode = append(g.FinalCode,"/*- Area de header -*/\n")
	g.FinalCode = append(g.FinalCode,"#include <stdio.h>\n")
	g.FinalCode= append(g.FinalCode,"#include <math.h>\n")
	g.FinalCode= append(g.FinalCode,"double heap[30101999];\n")
	g.FinalCode= append(g.FinalCode,"double stack[30101999];\n")
	g.FinalCode= append(g.FinalCode,"double P;\n")
	g.FinalCode= append(g.FinalCode,"double H;\n")
	
	if len(g.GetTemps()) > 0 {
		//****************** add temporal declaration
		g.FinalCode= append(g.FinalCode,"double ")
		for i, s := range g.GetTemps() {
			g.FinalCode= append(g.FinalCode,fmt.Sprintf("%v", s))
			if i < len(g.GetTemps())-1{
				g.FinalCode= append(g.FinalCode,", ")
			}
		}
		g.FinalCode= append(g.FinalCode,";\n")
	}
	
	g.FinalCode= append(g.FinalCode,"\n")
	//****************** add natives functions
	if len(g.Natives) > 0 {
		g.FinalCode = append(g.FinalCode,"/*- Area de Natives -*/\n")
		g.FinalCode = append(g.FinalCode, g.Natives...)
	}
	//****************** add functions
	if len(g.FuncCode) > 0 {
		g.FinalCode = append(g.FinalCode,"/*- Area de funciones -*/\n")
		g.FinalCode = append(g.FinalCode, g.FuncCode...)
	}
	//****************** add main
	g.FinalCode = append(g.FinalCode,"/*- Area de main -*/\n")
	g.FinalCode= append(g.FinalCode,"int main() {\n")
	g.FinalCode= append(g.FinalCode,"\tP = 0; H = 0;\n\n")
	for _, s := range g.Code {
		g.FinalCode= append(g.FinalCode,"\t" + s.(string))
	}
	g.FinalCode= append(g.FinalCode,"\n\treturn 0;\n}\n")
}

func (g *Generator) GeneratePrintString() {
	if g.PrintStringFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives=append(g.Natives,"void printString() {\n")
		g.Natives=append(g.Natives,"\t" + newTemp1 + " = P + 1;\n")
		g.Natives=append(g.Natives,"\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives=append(g.Natives,"\t" + newLvl2 + ":\n")
		g.Natives=append(g.Natives,"\t" + newTemp3 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives=append(g.Natives,"\tif(" + newTemp3 + " == -1) goto " + newLvl1 + ";\n")
		g.Natives=append(g.Natives,"\tprintf(\"%c\", (char)" + newTemp3 + ");\n")
		g.Natives=append(g.Natives,"\t" + newTemp2 + " = " + newTemp2 + " + 1;\n")
		g.Natives=append(g.Natives,"\tgoto " + newLvl2 + ";\n")
		g.Natives=append(g.Natives,"\t" + newLvl1 + ":\n")
		g.Natives=append(g.Natives,"\treturn;\n")
		g.Natives=append(g.Natives,"}\n\n")
		g.PrintStringFlag = false
	}
}

func (g *Generator) GenerateConcatString() {
	if g.ConcatStringFlag {
		//generando temporales y etiquetas
		tmp1 := g.NewTemp()
		tmp2 := g.NewTemp()
		tmp3 := g.NewTemp()
		tmp4 := g.NewTemp()
		tmp5 := g.NewTemp()
		lvl1 := g.NewLabel()
		lvl2 := g.NewLabel()
		lvl3 := g.NewLabel()
		lvl4 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives = append(g.Natives,"void concatString() {\n") 
		g.Natives = append(g.Natives,"\t" + tmp1 + " = H;" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp2 + " = P + 1;" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp4 + " = stack[(int)" + tmp2 + "];" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp3 + " = P + 2;" + "\n")
		g.Natives = append(g.Natives,"\t" + lvl2 + ":" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp5 + " = heap[(int)" + tmp4 + "];" + "\n")
		g.Natives = append(g.Natives,"\t" + "if(" + tmp5 + " == -1) goto " + lvl3 + ";" + "\n")
		g.Natives = append(g.Natives,"\t" + "heap[(int)H] = " + tmp5 + ";" + "\n")
		g.Natives = append(g.Natives,"\t" + "H = H + 1;" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp4 + " = " + tmp4 + " + 1;" + "\n")
		g.Natives = append(g.Natives,"\t" + "goto " + lvl2 + ";" + "\n")
		g.Natives = append(g.Natives,"\t" + lvl3 + ":" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp4 + " = stack[(int)" + tmp3 + "];" + "\n")
		g.Natives = append(g.Natives,"\t" + lvl4 + ":" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp5 + " = heap[(int)" + tmp4 + "];" + "\n")
		g.Natives = append(g.Natives,"\t" + "if(" + tmp5 + " == -1) goto " + lvl1 + ";" + "\n")
		g.Natives = append(g.Natives,"\t" + "heap[(int)H] = " + tmp5 + ";" + "\n")
		g.Natives = append(g.Natives,"\t" + "H = H + 1;" + "\n")
		g.Natives = append(g.Natives,"\t" + tmp4 + " = " + tmp4 + " + 1;" + "\n")
		g.Natives = append(g.Natives,"\t" + "goto " + lvl4 + ";" + "\n")
		g.Natives = append(g.Natives,"\t" + lvl1 + ":" + "\n")
		g.Natives = append(g.Natives,"\t" + "heap[(int)H] = -1;" + "\n")
		g.Natives = append(g.Natives,"\t" + "H =	 H + 1;" + "\n")
		g.Natives = append(g.Natives,"\t" + "stack[(int)P] = " + tmp1 + ";" + "\n")
		g.Natives = append(g.Natives,"\treturn;\n")
		g.Natives = append(g.Natives,"}\n\n")
		g.ConcatStringFlag = false
	}
}
func (g *Generator) Reset() {
    g.Temporal = 0
    g.Label = 0
    g.Code = []interface{}{}
    g.FinalCode = []interface{}{}
    g.Natives = []interface{}{}
    g.FuncCode = []interface{}{}
    g.TempList = []interface{}{}
    g.PrintStringFlag = true
    g.ConcatStringFlag = true
    g.BreakLabel = ""
    g.ContinueLabel = ""
    g.MainCode = true
}

//errores

func (g *Generator) AddErrorAsignacion(){
	//escribir "error de asignacion"
	g.Code = append(g.Code, "printf(\"%c\", (char)101);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)111);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)32);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)100);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)101);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)32);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)97);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)115);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)105);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)103);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)110);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)97);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)99);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)105);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)111);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)110);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)10);\n")
}

func (g *Generator) AddErrorDeclaracion(){
	//escribir "error de declaracion"
	g.Code = append(g.Code, "printf(\"%c\", (char)101);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)111);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)32);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)100);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)101);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)32);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)97);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)115);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)105);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)103);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)110);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)97);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)99);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)105);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)111);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)110);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)10);\n")
}

func (g *Generator) AddErrorOperacion(){
	//escribir "error de operacion"
	g.Code = append(g.Code, "printf(\"%c\", (char)101);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)111);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)32);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)100);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)101);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)32);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)111);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)112);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)101);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)114);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)97);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)99);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)105);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)111);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)110);\n")
	g.Code = append(g.Code, "printf(\"%c\", (char)10);\n")
}