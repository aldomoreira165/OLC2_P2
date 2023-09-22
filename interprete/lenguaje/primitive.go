package lenguaje

type Primitive struct {
	Lin   int
	Col   int
	Valor interface{}
	Tipo  TipoExpresion
}

func NewPrimitive(lin int, col int, valor interface{}, tipo TipoExpresion) Primitive {
	exp := Primitive{lin, col, valor, tipo}
	return exp
}
