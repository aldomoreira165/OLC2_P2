package lenguaje

type Matriz struct {
	Id      string
	Tipo    string
	Dimension int
	Valores [][]interface{}
}

type Matriz3D struct {
	Id      string
	Tipo    string
	Dimension int
	Valores [][][]interface{}
}