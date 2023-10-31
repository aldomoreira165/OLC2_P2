package lenguaje

type TablaSimbolos struct {
	Tabla map[string]Simbolo
}

type Simbolo struct {
	ID 	 string
	TipoSimbolo string
	TipoDato string
	Linea int
	Columna int
}

func NewTablaSimbolos() *TablaSimbolos {
	return &TablaSimbolos{
		Tabla: make(map[string]Simbolo),
	}
}

func (ts *TablaSimbolos) InsertarSimbolo(id string, tipoSimbolo string, tipoDato string, linea int, columna int) {
	ts.Tabla[id] = Simbolo{id, tipoSimbolo, tipoDato, linea, columna}
}