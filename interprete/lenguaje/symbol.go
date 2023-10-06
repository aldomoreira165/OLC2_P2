package lenguaje

import "fmt"

type Symbol struct {
	Lin      int
	Col      int
	Id       string
	Tipo     TipoExpresion
	TipoArr  TipoExpresion
	Posicion int
}

func getDataSymbol(s Symbol) {
	fmt.Println("--- Symbol ---")
	fmt.Println("Lin: ", s.Lin)
	fmt.Println("Col: ", s.Col)
	fmt.Println("Id: ", s.Id)
	fmt.Println("Tipo: ", s.Tipo)
	fmt.Println("TipoArr: ", s.TipoArr)
	fmt.Println("Posicion: ", s.Posicion)
	fmt.Println("--------------")
}